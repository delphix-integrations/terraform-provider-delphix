# Architecture

## Overview

This is a **HashiCorp Terraform Provider** built with the [Terraform Plugin SDK v2](https://developer.hashicorp.com/terraform/plugin/sdkv2). It wraps the [Delphix Control Tower (DCT) API](https://help.delphix.com/eh/current/content/terraform.htm) via the `dct-sdk-go` Go client library, enabling Infrastructure-as-Code management of Delphix Continuous Data resources.

```
Terraform CLI
     │  schema validation, plan/apply lifecycle
     ▼
terraform-provider-delphix   (this repo)
     │  translates HCL ↔ DCT API calls
     ├─► DCT API  (dct-sdk-go)     ← cloud-plane: VDB, dSource, environment, registration
     └─► Engine API  (raw HTTP)    ← data-plane: storage init, NTP, SMTP, SSO, password setup
```

There are **two distinct API surfaces**:
- **DCT API** — all resource lifecycle operations (create/read/update/delete) go through Delphix Control Tower. Calls are made via the generated `dct-sdk-go` SDK.
- **Engine API** — direct HTTP calls to individual Delphix Engine nodes, used only by `delphix_engine_configuration` for first-time setup operations that DCT does not expose (storage initialization, NTP, SMTP, SSO, DNS, web proxy).

---

## Directory Structure

```
terraform-provider-delphix/
├── main.go                        # Plugin entry point — sets version, enables debug flag
├── go.mod                         # Module: github.com/delphix/dct-sdk-go/v25, terraform-plugin-sdk/v2
├── GNUmakefile                    # Build targets: build, install, release, test, testacc
├── .goreleaser.yml                # Cross-compile + GPG-sign release artifacts
├── internal/provider/             # All provider logic (single package: `provider`)
│   ├── provider.go                # Provider schema, client init, resource registration
│   ├── commons.go                 # Constants, job state strings, updatable field maps
│   ├── models.go                  # Go structs for Engine API JSON payloads
│   ├── utility.go                 # PollJobStatus, flatten helpers, conversion helpers
│   ├── engine_api.go              # Direct Engine API calls (HTTP, not DCT SDK)
│   ├── engine_api_utility.go      # Engine API request/response helpers
│   ├── logging.go                 # Logger initialization (WarnLog, InfoLog, ErrorLog)
│   ├── security.go                # SecureString type, memory-zeroing helpers
│   ├── resource_vdb.go            # delphix_vdb (~2300 lines)
│   ├── resource_vdb_group.go      # delphix_vdb_group
│   ├── resource_environment.go    # delphix_environment
│   ├── resource_appdata_dsource.go # delphix_appdata_dsource
│   ├── resource_oracle_dsource.go  # delphix_oracle_dsource
│   ├── resource_database_postgresql.go # delphix_database_postgresql
│   ├── resource_engine_configuration.go # delphix_engine_configuration
│   ├── resource_engine_registration.go  # delphix_engine_dct_registration
│   ├── resource_engine_plugin_upload.go # delphix_database_plugin
│   └── *_test.go                  # Unit + acceptance tests (one per resource)
├── examples/                      # Runnable Terraform configs per resource
└── docs/                          # Auto-generated Terraform Registry docs
```

---

## Provider Initialization (`provider.go`)

### Schema

The provider block exposes four fields, all readable from environment variables:

| Field | Env Var | Default | Purpose |
|---|---|---|---|
| `key` | `DCT_KEY` | — | DCT API key (sensitive) |
| `host` | `DCT_HOST` | — | DCT hostname |
| `host_scheme` | `DCT_HOST_SCHEME` | `"https"` | HTTP or HTTPS |
| `tls_insecure_skip` | `DCT_TLS_INSECURE_SKIP` | `false` | Skip TLS cert verification |
| `debug` | — | `false` | Log raw HTTP requests |

### Client Initialization (`configure()`)

```
configure()
  │
  ├─ dctapi.NewConfiguration()         set Host, Scheme, UserAgent
  ├─ http.Transport{TLSClientConfig}   TLS config (InsecureSkipVerify if set)
  ├─ Add default headers:
  │     Authorization: apk {KEY}
  │     x-dct-client-name: Terraform
  ├─ dctapi.NewAPIClient(cfg)          create SDK client
  ├─ client.ManagementAPI.GetRegisteredEngines()  ← credential validation test call
  └─ return &apiClient{client}         ← passed as meta to all CRUD functions
```

The `apiClient` wrapper struct is the `meta interface{}` parameter in every resource CRUD function. Cast it with:

```go
client := meta.(*apiClient).client   // *dctapi.APIClient
```

---

## Resource Lifecycle Pattern

Every resource follows the same structural pattern using Terraform Plugin SDK v2 context-aware methods.

### Schema Registration

```go
func resourceVdb() *schema.Resource {
    return &schema.Resource{
        CreateContext: resourceVdbCreate,
        ReadContext:   resourceVdbRead,
        UpdateContext: resourceVdbUpdate,
        DeleteContext: resourceVdbDelete,
        Importer: &schema.ResourceImporter{
            StateContext: schema.ImportStatePassthroughContext,
        },
        Timeouts: &schema.ResourceTimeout{
            Create: schema.DefaultTimeout(20 * time.Minute),
            Update: schema.DefaultTimeout(20 * time.Minute),
            Delete: schema.DefaultTimeout(20 * time.Minute),
        },
        Schema: map[string]*schema.Schema{ ... },
    }
}
```

### CRUD Flow

```
Terraform Apply
     │
     ├─► resourceXxxCreate(ctx, d, meta)
     │       │
     │       ├─ Build SDK parameters struct from d.Get(...)
     │       ├─ ctx = context.WithTimeout(ctx, d.Timeout(schema.TimeoutCreate))
     │       ├─ result, httpRes, err = client.XxxAPI.CreateXxx(ctx, params).Execute()
     │       ├─ jobId = result.GetJob().GetId()
     │       ├─ PollJobStatus(jobId, ctx, client)  ← waits for async completion
     │       ├─ d.SetId(result.GetId())
     │       └─ resourceXxxRead(ctx, d, meta)       ← populate computed fields
     │
     ├─► resourceXxxRead(ctx, d, meta)
     │       │
     │       ├─ result, httpRes, err = client.XxxAPI.GetXxxById(ctx, id).Execute()
     │       ├─ if 404 → d.SetId("") → return nil  (remove from state)
     │       └─ d.Set(...) for all schema fields
     │
     ├─► resourceXxxUpdate(ctx, d, meta)
     │       │
     │       ├─ Identify changed keys via d.HasChange(...)
     │       ├─ Validate against updatableXxxKeys map
     │       ├─ If destructive field changed → disableXxx() → update → enableXxx()
     │       ├─ client.XxxAPI.UpdateXxx(ctx, id, params).Execute()
     │       ├─ PollJobStatus(jobId, ctx, client)
     │       └─ resourceXxxRead(ctx, d, meta)
     │
     └─► resourceXxxDelete(ctx, d, meta)
             │
             ├─ client.XxxAPI.DeleteXxx(ctx, id, params).Execute()
             ├─ PollJobStatus(jobId, ctx, client)
             └─ PollForObjectDeletion(ctx, id, client)  ← wait for 404
```

---

## Async Job Polling (`utility.go`)

All DCT mutating calls are asynchronous — they return a job ID immediately and execute in the background.

### `PollJobStatus(jobId, ctx, client)`

```
loop:
  GET /jobs/{jobId}
  if status ∈ {COMPLETED, FAILED, CANCELED, ABANDONED, TIMEDOUT} → break
  if ctx.Done() → break (timeout or cancel)
  sleep STATUS_POLL_SLEEP_TIME (20s)

returns (status string, errorDetails string)
```

Job states defined in `commons.go`:
```
PENDING → STARTED → COMPLETED   (success)
                  → FAILED       (DCT-side failure)
                  → CANCELED     (user canceled)
                  → ABANDONED    (DCT abandoned)
                  → TIMEDOUT     (DCT timeout)
```

### Timeout Handling

Each CRUD function wraps its context:
```go
ctx, cancel := context.WithTimeout(ctx, d.Timeout(schema.TimeoutCreate))
defer cancel()
```

If the context deadline is exceeded before the job completes, the provider logs the Job ID and instructs the user to check the DCT UI manually — it does **not** cancel the underlying DCT operation.

### Object Existence Polling

For deletes, the provider waits until the DCT API returns 404:
```go
PollForObjectDeletion(ctx, id, client)  // polls until HTTP 404
PollForObjectExistence(ctx, id, client) // polls until HTTP 200
```

---

## Engine API Layer (`engine_api.go`, `engine_api_utility.go`, `models.go`)

Used exclusively by `delphix_engine_configuration`. Bypasses DCT and talks directly to the engine's legacy REST API (`/resources/json/delphix/...`).

### Authentication Flow

```
startSession()   →  POST /resources/json/delphix/session   (sets cookie)
login()          →  POST /resources/json/delphix/login      (authenticates)
```

Session cookie is maintained across subsequent calls in the same HTTP client.

### Initialization Flow

```
initializeSystemAndDevices()
  │
  ├─ device_type == "BLOCK"  →  GET /storage/device (discover disks)
  │                          →  initializeSystem(block storage config)
  │
  └─ device_type == "OBJECT" →  testConnectionForObjectStore()
                             →  initializeSystem(object store config)
                                  │
                                  ├─ AWS:   S3 bucket, region, ROLE or ACCESS_KEY auth
                                  ├─ AZURE: Blob container, account, MANAGED_IDENTITIES or ACCESS_KEY
                                  └─ GCP:   Cloud Storage bucket (no explicit auth_type)
```

### System Services

After initialization, optional service config calls:

| Function | Engine Endpoint |
|---|---|
| `setNtpServers()` | `/resources/json/delphix/service/time` |
| `configureSMTP()` | `/resources/json/delphix/service/smtp` |
| `configureDNS()` | `/resources/json/delphix/service/dns` |
| `configurePhoneHome()` | `/resources/json/delphix/service/phonehome` |
| `configureUserAnalytics()` | `/resources/json/delphix/service/useranalytics` |
| `configureWebProxy()` | `/resources/json/delphix/service/proxy` |
| `configureSSO()` | `/resources/json/delphix/service/saml` |

### Engine API Version

Pinned at `ENGINE_API_VERSION = "1.11.40"` in `models.go`.

---

## State Management Details

### Computed vs Optional Fields

Resources use a deliberate distinction:
- **`Computed: true` only** — fields populated entirely by DCT (e.g., `database_type`, `engine_id`, `ip_address`, `parent_id`, `creation_date`).
- **`Optional: true, Computed: true`** — user may provide a value, but DCT may override or compute it (e.g., `name`, `environment_id`, `repository_id`).
- **`Required: true`** — must be user-provided (e.g., provider `key`, `host`).

### Tag Handling

Tags are tricky because Terraform diffs every computed list element. Two mechanisms prevent spurious diffs:

1. **`ignore_tag_changes`** (default `true`) — when set, the provider skips writing tags to state during Read, preserving whatever was in the config.
2. **`CustomizeDiffTags()`** — a `CustomizeDiff` function that suppresses diffs when tags are only being deleted (handles DCT-managed tags the user didn't add).
3. **`HandleRawConfigReadContext()`** — inspects the raw config (not state) to determine what the user actually specified, avoiding phantom diffs from computed additions.

### Destructive Updates (VDB)

Some VDB fields require the VDB to be disabled, updated, and re-enabled:

```go
if isDestructiveVdbUpdate[changedKey] {
    disableVDB(ctx, vdbId, client)   // POST /vdbs/{id}/disable → poll job
    // ... apply update ...
    enableVDB(ctx, vdbId, client)    // POST /vdbs/{id}/enable  → poll job
}
```

The `isDestructiveVdbUpdate` map in `commons.go` flags which fields trigger this cycle (e.g., `template_id`, `mount_point`).

### Update Field Gating

The `updatableVdbKeys` map (and equivalents for dSources, environments) explicitly lists every field that supports in-place update. If a user changes a field not in this map, the provider returns an error rather than silently ignoring it.

---

## Security (`security.go`)

### Sensitive Field Handling

- API keys and passwords are declared `Sensitive: true` in schema — Terraform redacts them from plan output.
- `SecureString` wraps sensitive values and provides a `Clear(ctx)` method that overwrites memory with random bytes then zeros it (best-effort, since Go strings are immutable).

### TLS Configuration

Built in `configure()`:
```go
tlsConfig := &tls.Config{InsecureSkipVerify: tls_insecure_skip}
transport := &http.Transport{TLSClientConfig: tlsConfig}
httpClient := &http.Client{Transport: transport}
```

Never set `InsecureSkipVerify: true` in production — it disables certificate chain validation.

---

## Logging

Two logging mechanisms coexist:

### `tflog` (preferred for resource code)
```go
import "github.com/hashicorp/terraform-plugin-log/tflog"

tflog.Info(ctx, "[DELPHIX] [INFO] Creating VDB")
tflog.Warn(ctx, "[DELPHIX] [WARN] Field deprecated")
tflog.Error(ctx, "[DELPHIX] [ERROR] Job failed", map[string]interface{}{"job_id": id})
```
Integrated with Terraform's structured logging — visible via `TF_LOG=INFO terraform apply`.

### Standard `log.Logger` (engine_api.go only)
```go
InfoLog.Printf("[DELPHIX] [INFO] ...")
WarnLog.Printf("[DELPHIX] [WARN] ...")
ErrorLog.Printf("[DELPHIX] [ERROR] ...")
```
Initialized in `logging.go`'s `init()` function with appropriate output streams.

---

## Error & Diagnostic Handling

Resources return `diag.Diagnostics` to Terraform (not raw `error`).

### Patterns

```go
// Wrap a Go error
return diag.FromErr(err)

// Create a formatted error
return diag.Errorf("VDB creation failed: %s", jobErr)

// Unified API error handler (logs + wraps HTTP response body)
return apiErrorResponseHelper(ctx, "GetVdbById", httpRes, err)
```

### `apiErrorResponseHelper(ctx, operation, httpRes, err)`
1. Reads `httpRes.Body` to a string (preserves API error detail).
2. Logs via `tflog.Error`.
3. Returns `diag.Errorf` with both the Go error and the response body.

### 404 Handling in Read

```go
if httpRes.StatusCode == http.StatusNotFound {
    d.SetId("")   // remove from state, Terraform will plan a re-create
    return nil
}
```

---

## DCT SDK Usage Pattern

All DCT operations follow the builder pattern from `dct-sdk-go`:

```go
// Execute a request
result, httpRes, err := client.VDBsAPI.
    ProvisionVdbBySnapshot(ctx).
    ProvisionVDBBySnapshotParameters(params).
    Execute()

// Job polling after mutation
jobId := result.GetJob().GetId()
status, jobErr := PollJobStatus(jobId, ctx, client)
if status != COMPLETED {
    return diag.Errorf("provisioning failed: %s", jobErr)
}
```

### API Groups Used

| API Group | Resources That Use It |
|---|---|
| `client.VDBsAPI` | `delphix_vdb` |
| `client.VDBGroupsAPI` | `delphix_vdb_group` |
| `client.EnvironmentsAPI` | `delphix_environment` |
| `client.DSourcesAPI` | `delphix_appdata_dsource`, `delphix_oracle_dsource` |
| `client.SourcesAPI` | `delphix_database_postgresql` |
| `client.ManagementAPI` | `delphix_engine_dct_registration`, provider init |
| `client.JobsAPI` | `utility.go` — `PollJobStatus` |
| `client.TimeflowsAPI` | `resource_vdb.go` — snapshot lookup for import |

---

## Adding a New Resource

1. **Create** `internal/provider/resource_<name>.go` in package `provider`.
2. **Implement** the four context-aware CRUD functions.
3. **Register** in `provider.go` under `ResourcesMap`:
   ```go
   "delphix_<name>": resource<Name>(),
   ```
4. **Add constants** to `commons.go` — updatable field map, any new job state handling.
5. **Add tests** in `resource_<name>_test.go` — both unit and acceptance.
6. **Add example** in `examples/<name>/main.tf`.
7. **Schema descriptions** will auto-generate Registry docs on next `go generate`.

### Checklist

- [ ] All DCT mutations go through `PollJobStatus` — never assume synchronous completion.
- [ ] 404 in Read sets `d.SetId("")` and returns `nil`.
- [ ] Destructive updates disable/re-enable the resource if needed.
- [ ] Sensitive fields marked `Sensitive: true`.
- [ ] Timeouts wired via `context.WithTimeout(ctx, d.Timeout(...))`.
- [ ] Import supported via `schema.ImportStatePassthroughContext`.
- [ ] Logging uses `tflog` with `[DELPHIX]` prefix.

---

## Key Constants (`commons.go`)

```go
// Job terminal states
COMPLETED = "COMPLETED"
FAILED    = "FAILED"
CANCELED  = "CANCELED"
ABANDONED = "ABANDONED"
TIMEDOUT  = "TIMEDOUT"

// Poll intervals
JOB_STATUS_SLEEP_TIME  = 5   // seconds — used in engine API
STATUS_POLL_SLEEP_TIME = 20  // seconds — used in DCT SDK polls

// Log prefixes
DLPX  = "[DELPHIX]"
INFO  = "[INFO]"
WARN  = "[WARN]"
ERROR = "[ERROR]"
```

---

## Data Flow Diagram

```
User writes HCL
      │
      ▼
Terraform parses + validates schema
      │
      ▼
resource<Name>Create/Update/Delete(ctx, d, meta)
      │
      ├─ d.Get("field")           read config values
      ├─ build SDK params struct
      ├─ client.XxxAPI.Xxx().Execute()    ──► DCT API (HTTPS)
      │                                         │
      │                                         ▼
      │                                    Job ID returned
      │                                         │
      ├─ PollJobStatus(jobId, ctx, client) ◄────┘
      │       └─ GET /jobs/{id} every 20s
      │
      ├─ resource<Name>Read(ctx, d, meta)
      │       └─ GET /xxx/{id}
      │       └─ d.Set("field", value)    write to Terraform state
      │
      ▼
Terraform state updated
```
