# Test Infrastructure — terraform-provider-delphix

This file describes the *infrastructure* required to run tests — what services must exist, how they're configured, and what state they need to be in.

## Provisioning fresh test infrastructure with `dc` command (engine_configuration + environment)

> **Applies when `DELPHIX_TEST_RESOURCE` is one of: `engine_configuration`, `environment`.**
>
> - `engine_configuration` tests need a freshly-cloned **engine** VM — `engine_configuration` is a destructive first-boot config and cannot be re-run against an already-configured engine.
> - `environment` tests need a freshly-cloned **target host** VM — fresh hosts ensure the toolkit install path is clean and prior test state doesn't bleed in. They reuse an already-configured engine via `ENGINE_NAME`.
>
> All other resources (`vdb`, `vdb_group`, `appdata_dsource`, `oracle_dsource`, `database_postgresql`, `engine_registration`, `database_plugin`) **reuse** existing infrastructure and **must not** trigger any `dc clone` or `dc destroy`. For those, `ENGINE_NAME` simply points at a long-lived pre-configured engine the user maintains; the golden-image vars stay empty.

Use the dc commands to clone VMs from golden images.

### Prerequisite — SSH to the `dc` host

> **Do NOT install `dc` locally.** The `dc` CLI is already installed and configured on `DCOA_HOST`. Do not run any installer (brew, pip, curl-to-shell, etc.) on the test runner or dev laptop. If `dc` is missing on the remote host, surface that to the user — do not attempt a workaround by installing it locally.

**Before** running any provisioning step below, open an SSH session to that host using the credentials from `.claude/settings.local.json`:

```bash
# Values come from .claude/settings.local.json — never hardcode
ssh "${DCOA_USER}@${DCOA_HOST}"
# password: ${DCOA_PASSWORD}
```

Once the SSH session is established, **log in to `dc`** before any other command call:

```bash
dc login
# username: ${DCOA_USER}
# password: ${DCOA_PASSWORD}
# Authenticator code: <ask the user — 2FA TOTP, rotates every 30s>
```

The authenticator code is a time-based 2FA token that is **not** stored in `settings.local.json` and cannot be derived from any other env var. The assistant must **pause and ask the user** for the current code at the moment `dc login` prompts for it — do not guess, do not retry with stale codes.


If `DCOA_HOST`, `DCOA_USER`, or `DCOA_PASSWORD` is empty in `settings.local.json`, halt and ask the user to populate them before continuing.

### Required env vars in `.claude/settings.local.json`

| Var | Required when | Purpose | Example |
|---|---|---|---|
| `DCOA_HOST` | **Always** (any flow that touches `dc`) | Hostname of the dc host you SSH into to run `dc` commands | `dlpxdc.co` |
| `DCOA_USER` | **Always** | SSH user on the dc host | `user@example.com` |
| `DCOA_PASSWORD` | **Always** | SSH password on the dc host (sensitive — gitignored) | — |
| `ENGINE_NAME` | **Always** (source for `DELPHIX_ENGINE_HOST` template in every test) | The `dc` VM name for the engine. Cloned for `engine_configuration` tests; pre-existing for all others. | `tergcpcc` |
| `ENGINE_GOLDEN_IMAGE` | **engine_configuration only** | Golden image group used by `dc clone-latest` to create a fresh engine VM | `dlpx-develop` |
| `DELPHIX_ENGINE_CLOUD` | **engine_configuration only** | Target cloud passed to `dc clone-latest --cloud` | `GCP` |
| `ENV_NAME` | **environment only** (and optionally for other tests that need a target host) | The `dc` VM name for the target host | `tergcptarget` |
| `ENVIRONMENT_GOLDEN_IMAGE` | **environment only** | Golden image group used by `dc clone-latest` to create a fresh target host | `dlpx-target-linux` |

Required-var checklist by test resource (every row also requires `DCOA_HOST` + `DCOA_USER` + `DCOA_PASSWORD` for the SSH prerequisite):

| `DELPHIX_TEST_RESOURCE` | Must be non-empty | Optional |
|---|---|---|
| `engine_configuration` | `ENGINE_NAME` + `ENGINE_GOLDEN_IMAGE` + `DELPHIX_ENGINE_CLOUD` | — |
| `environment` | `ENGINE_NAME` (pre-configured) + `ENV_NAME` + `ENVIRONMENT_GOLDEN_IMAGE` + `DELPHIX_ENGINE_CLOUD` | — |
| anything else | `ENGINE_NAME` (pre-configured) | resource-specific vars from the matrix above |

If any required-for-this-resource var is empty, the assistant must halt and ask the user to populate `settings.local.json` before invoking `dc` or `go test`.

### Derived value: `DELPHIX_ENGINE_HOST`

`DELPHIX_ENGINE_HOST` is stored as the template `"http://${ENGINE_NAME}.dlpxdc.co"` so that updating `ENGINE_NAME` alone updates the engine URL. The assistant **must expand `${ENGINE_NAME}` at run time** when reading the env from `settings.local.json` — substitute the current value of `ENGINE_NAME` before passing the value to `go test`.

Resolution rule:
1. Read `ENGINE_NAME` from `.claude/settings.local.json`.
2. Read the raw `DELPHIX_ENGINE_HOST` template (e.g. `http://${ENGINE_NAME}.dlpxdc.co`).
3. Substitute `${ENGINE_NAME}` with the value from step 1.
4. Export the resolved value as `DELPHIX_ENGINE_HOST` before running tests.

If `ENGINE_NAME` is empty, the URL is incomplete — surface that to the user and ask for a name.

To change the engine target between runs: **edit `ENGINE_NAME` only** — do not edit `DELPHIX_ENGINE_HOST` directly.

### Provisioning workflow

Tests requiring fresh infrastructure (any `engine_configuration` test, or any test where `DELPHIX_ENGINE_HOST` does not resolve / engine is not reachable) clone VMs from golden images via clone command. VM names, golden images, and cloud target all come from `.claude/settings.local.json` — never hardcoded, never auto-generated.

**Project-side gates (must hold before any `dc` invocation):**

1. Required vars for the selected resource (see table above) are all non-empty.
3. For an `engine_configuration` test: if a VM already exists with name `ENGINE_NAME`, reuse it — do not re-clone, do not destroy.


**Clone + test + teardown shape:**

```bash
# Read these from .claude/settings.local.json — never hardcode
#   ENGINE_NAME, ENV_NAME, ENGINE_GOLDEN_IMAGE, ENVIRONMENT_GOLDEN_IMAGE, DELPHIX_ENGINE_CLOUD

1. Clone the engine VM if it doesn't already exist. Use ENGINE_GOLDEN_IMAGE for image and name should be ENGINE_NAME. if DELPHIX_ENGINE_CLOUD is provided then use that cloud to deploy engine. 

# 4. Run the test
go test -v -timeout 120m -run "<regex resolved from DELPHIX_TEST_RESOURCE/TYPE/CLOUD>" ./internal/provider/

# 5. Teardown — only destroy VMs that were freshly cloned in this run
[ "$CLONED_ENGINE" = "true" ] && dc destroy "$ENGINE_NAME" -w
[ "$CLONED_ENV" = "true" ] && dc destroy "$ENV_NAME" -w
```

**Rules:**
- **Idempotent:** if `dc list` shows the VM exists, skip the clone — engines are expensive to provision (3–5 min each).
- **Don't destroy what you didn't create:** only destroy VMs that this run cloned (tracked via `CLONED_ENGINE` / `CLONED_ENV`).
- **Cloud match:** the engine must be cloned in the cloud matching `DELPHIX_ENGINE_CLOUD` so the test's cloud-specific config (S3/Azure/GCS bucket) is reachable.

### When to clone vs reuse

| Scenario | Action |
|---|---|
| Running `engine_configuration` test | **Clone fresh** — engines can only be first-boot configured once |
| Running `vdb` / `dsource` / `environment` test | **Reuse** an already-configured engine + bucket; clone a fresh target host if needed |
| CI / automated run | **Clone fresh** for every run; `dc expire <days>` to auto-cleanup |
| Local dev iteration on non-config tests | **Reuse** to save quota + time |

### Cost & cleanup

- Default size (`MEDIUM` ≈ t3.medium ≈ $0.04/hr on AWS). GCP equivalents apply.
- Use `dc expire <days> <name>` to auto-destroy after N days as a safety net.
- Always `dc destroy <name> -w` when done to release resources.
- Use `dc list -a` to audit all VMs in your account.

---

## Step 0 — Pick the component(s) to test

Before running any acceptance test, **ask the user which component(s) to test** from the matrix below. Each resource has different infrastructure requirements, environment variables, and side effects. Running "all tests" against a shared DCT is rarely the right move — it provisions real VDBs/dSources/environments and can take 60+ minutes.

### Component Test Matrix

| # | Resource | Test file | Required env vars | Infrastructure needed | Notes |
|---|---|---|---|---|---|
| 1 | `delphix_engine_configuration` | [resource_engine_configuration_test.go](../internal/provider/resource_engine_configuration_test.go) | `DELPHIX_ENGINE_HOST` + one of: `S3_BUCKET_NAME` & `AWS_ACCESS_KEY_ID`/`AWS_SECRET_ACCESS_KEY` (AWS); `AZURE_CONTAINER_NAME` & `AZURE_ACCOUNT_NAME` & `AZURE_ACCESS_KEY` (Azure); `GCP_BUCKET_NAME` (GCP) | Unconfigured Delphix engine (CD or CC); for OBJECT storage, a cloud bucket the engine VM can reach; NTP reachable | First-boot config — destructive. CC variant requires `engine_type=CC` plus `compliance_user`, `compliance_password`, `compliance_new_password`, `compliance_email`. |
| 2 | `delphix_engine_dct_registration` | [resource_engine_registration_test.go](../internal/provider/resource_engine_registration_test.go) | `DELPHIX_ENGINE_HOST`, engine credentials | Already-configured engine reachable from DCT | Registers an existing engine with DCT. |
| 3 | `delphix_environment` | [resource_environment_test.go](../internal/provider/resource_environment_test.go) | Target host hostname, SSH credentials, OS user | A Unix or Windows host SSH-reachable from the CD engine; Delphix toolkit install path writable | Standalone, cluster, and RAC variants supported. |
| 4 | `delphix_vdb` | [resource_vdb_test.go](../internal/provider/resource_vdb_test.go) | Source dSource ID, target environment ID, target group | Existing dSource with a snapshot; target environment registered; storage on CD engine | Slowest test — provision can take several minutes per VDB. |
| 5 | `delphix_vdb_group` | [resource_vdb_group_test.go](../internal/provider/resource_vdb_group_test.go) | VDB IDs to group | Two or more existing VDBs | Manages VDB grouping + tags; cheap. |
| 6 | `delphix_appdata_dsource` | [resource_appdata_dsource_test.go](../internal/provider/resource_appdata_dsource_test.go) | Source host details, AppData plugin reference, link parameters | Source host registered as environment; AppData plugin (e.g. PostgreSQL) uploaded to engine | Plugin must be installed on engine first. |
| 7 | `delphix_oracle_dsource` | [resource_oracle_dsource_test.go](../internal/provider/resource_oracle_dsource_test.go) | Source Oracle host, database name, RMAN connection user | Oracle 11g+ source DB, RMAN access from engine, listener reachable | Most setup-heavy source type. |
| 8 | `delphix_database_postgresql` | [resource_database_postgresql_test.go](../internal/provider/resource_database_postgresql_test.go) | Source PostgreSQL host, port, repository user | PostgreSQL 9.6+, replication user, archive WAL access | Pure PostgreSQL resource — doesn't require AppData plugin. |
| 9 | `delphix_database_plugin` | [resource_engine_plugin_upload_test.go](../internal/provider/resource_engine_plugin_upload_test.go) | Path to plugin .tgz | Engine accepts plugin uploads; valid plugin artifact on disk | Uploads/replaces plugins on the engine. |
| — | Provider config / TLS | [provider_test.go](../internal/provider/provider_test.go), [security_test.go](../internal/provider/security_test.go) | None (unit-only) | None | Always safe to run; runs under `make test`. |

### Selecting the target via `DELPHIX_TEST_RESOURCE` / `DELPHIX_ENGINE_TYPE` / `DELPHIX_ENGINE_CLOUD`

`.claude/settings.local.json` has three env vars that compose to narrow test selection:

| Var | Purpose | Values |
|---|---|---|
| `DELPHIX_TEST_RESOURCE` | Which resource | `engine_configuration`, `engine_registration`, `environment`, `vdb`, `vdb_group`, `appdata_dsource`, `oracle_dsource`, `database_postgresql`, `database_plugin`, `all`, or empty |
| `DELPHIX_ENGINE_TYPE` | Engine type (resource-specific) | `CD`, `CC`, or empty |
| `DELPHIX_ENGINE_CLOUD` | Storage backend (resource-specific) | `BLOCK`, `AWS`, `AZURE`, `GCP`, or empty |

The assistant resolves these to a `go test -run` regex per the tables below. Empty narrowers mean "all variants of the resource". An explicit chat instruction ("run only the GCP CC test") always overrides the env vars.

#### Resolution table — `engine_configuration`

| TYPE | CLOUD | `-run` regex | Test function |
|---|---|---|---|
| `CD` | `BLOCK` (or empty CLOUD) | `^TestAccEngineConfiguration_blockDevice$` | block-device CD |
| `CD` | `AWS` | `^TestAccEngineConfiguration_objectStorageWith(Role\|AccessKey)$` | AWS role + access-key |
| `CD` | `AZURE` | `^TestAccEngineConfiguration_azureObjectStorage` | Azure managed-id + access-key |
| `CD` | `GCP` | `^TestAccEngineConfiguration_gcpObjectStorage$` | GCP CD |
| `CC` | `GCP` | `^TestAccEngineConfiguration_gcpObjectStorage_CC$` | GCP CC |
| `CC` | `AWS` / `AZURE` | — | not implemented |
| empty | empty | `^TestAccEngineConfiguration_` | all engine_configuration tests |

#### Resolution table — other resources

`DELPHIX_ENGINE_TYPE` / `DELPHIX_ENGINE_CLOUD` are mostly engine_configuration-specific. For other resources the regex is just the prefix below; `TYPE`/`CLOUD` are ignored unless that resource adds variant tests later.

| `DELPHIX_TEST_RESOURCE` | `-run` regex |
|---|---|
| `engine_registration` | `^TestAccEngineDctRegistration_` |
| `environment` | `^TestAccEnvironment_` |
| `vdb` | `^TestAccVDB_` |
| `vdb_group` | `^TestAccVDBGroup_` |
| `appdata_dsource` | `^TestAccAppDataDsource_` |
| `oracle_dsource` | `^TestAccOracleDsource_` |
| `database_postgresql` | `^TestAccDatabasePostgresql_` |
| `database_plugin` | `^TestAccDatabasePlugin_` |
| `all` | `.*` (every acceptance test — long, infrastructure-heavy) |
| *empty* | no preselection — ask the user |

#### Example invocations

```bash
# DELPHIX_TEST_RESOURCE=engine_configuration, DELPHIX_ENGINE_TYPE=CC, DELPHIX_ENGINE_CLOUD=GCP
go test -v -timeout 120m -run "^TestAccEngineConfiguration_gcpObjectStorage_CC$" ./internal/provider/

# DELPHIX_TEST_RESOURCE=engine_configuration, DELPHIX_ENGINE_TYPE=CD, DELPHIX_ENGINE_CLOUD=AWS
go test -v -timeout 120m -run "^TestAccEngineConfiguration_objectStorageWith(Role|AccessKey)$" ./internal/provider/

# DELPHIX_TEST_RESOURCE=engine_configuration, narrowers empty
go test -v -timeout 120m -run "^TestAccEngineConfiguration_" ./internal/provider/
```

When a chat message specifies the target ("run only the GCP CC test"), that wins over all three env vars.

### Decision questions to ask the user

When the user says "run tests" or "test X" without specifying scope, ask:

1. **Which resource(s)?** Pick from the matrix above. Default to "engine_configuration" only if no context.
2. **Which cloud / source variant?** For engine_configuration: AWS / Azure / GCP. For dSources: Oracle / PostgreSQL / AppData. For environments: Unix / Windows / RAC.
3. **CD or CC engine?** CD covers all VDB/dSource flows; CC covers compliance-only flows.
4. **Is the infrastructure already provisioned?** Existing engine + bucket = fast. Need to create resources = slow. Confirm reuse vs fresh.
5. **Destructive OK?** Acceptance tests *create real resources*. `engine_configuration` first-boot config is **non-reversible** without re-imaging the engine.

If the user picks multiple components, run them **sequentially** and verify infrastructure for each before starting — a failed dependency wastes the prior runs.

### Common cross-component env vars

> **Source of truth:** All env vars listed in this file (component-specific and common) should be populated from [`.claude/settings.local.json`](settings.local.json) under the `env` block. The file is gitignored, so values stay local to your machine. Do **not** hand-export credentials in your shell ad-hoc when `settings.local.json` exists — keep one source of truth so every test run sees the same values. When adding a new test that needs a new env var, add it to `settings.local.json` (with an empty default) and document it in the matrix above.

These apply to **all** acceptance tests regardless of which component you're testing:

| Var | Required | Default |
|---|---|---|
| `DCT_KEY` | Yes | — |
| `DCT_HOST` | Yes | — |
| `DCT_HOST_SCHEME` | No | `https` |
| `DCT_TLS_INSECURE_SKIP` | No (dev sandboxes only) | `false` |
| `TF_ACC` | Yes (must be `1`) | — |

---

## Tier 1 — Unit Tests (no infrastructure)

Unit tests under `make test` run entirely in-process. No DCT, no engines, no network.

- Pre-reqs: Go 1.25+, `make`, the module's vendored deps via `go mod download`.
- Runs anywhere: dev laptop, CI, ephemeral container.
- Total wall time: well under 30s.

## Tier 2 — Acceptance Tests (live infrastructure required)

Acceptance tests under `TF_ACC=1 make testacc` provision and destroy real resources via DCT. They cannot be mocked.

### Required Components

| Component | Purpose | Min version |
|---|---|---|
| DCT (Delphix Control Tower) | API endpoint the provider talks to | v2025.2.0+ |
| Delphix CD engine | Continuous Data engine for VDB / dSource tests | v29.0.0.0+ |
| Delphix CC engine | Continuous Compliance engine for engine-registration / compliance flows | v29.0.0.0+ |
| Source database host(s) | Oracle, PostgreSQL, AppData sources backing dSource tests | Per source plugin matrix |
| Target environment host(s) | Unix/Windows env(s) for VDB provisioning | Linux x86_64 or Windows Server |
| DCT API key | Authentication for the provider | n/a |

### DCT Setup

The DCT instance must be:

1. **Reachable from the test runner** over HTTPS (or HTTP via `DCT_HOST_SCHEME=http` for non-prod sandboxes).
2. **Registered with at least one CD engine** that has space for VDB provisioning.
3. **Optionally registered with a CC engine** if the test covers `delphix_engine_dct_registration` with `engine_type=CC`.
4. **Provisioned with an API key** issued to a user that has `admin`-equivalent permissions (resource CRUD across all resource types).

Set the key and host before running:

```bash
export DCT_KEY="<api-key>"
export DCT_HOST="dct.example.internal"
export DCT_HOST_SCHEME="https"     # default; set http only for sandboxes
# export DCT_TLS_INSECURE_SKIP=true # only if your sandbox has a self-signed cert
```

### Source Hosts

Acceptance tests for dSource resources need at least one source host per supported source type the change touches:

| Source type | Host requirement |
|---|---|
| Oracle (`delphix_oracle_dsource`) | Oracle DB (11g+) with RMAN access, listener reachable from CD engine |
| PostgreSQL (`delphix_database_postgresql` / appdata) | PostgreSQL 9.6+, replication user, archive WAL access |
| AppData generic (`delphix_appdata_dsource`) | Plugin-specific — see the plugin's source linking docs |

Source hosts must be reachable on the network from the CD engine, and the connection user must have plugin-defined privileges (RMAN for Oracle, replication for PG, etc.).

### Target Environments

VDB provisioning tests need:

- A Unix or Windows environment registered to the CD engine via `delphix_environment`.
- A Delphix toolkit installed at a writable path (default `/var/opt/delphix/toolkit`).
- Sufficient mount points / disk space for the VDB filesystem.

## Tier 3 — Cloud Object Storage Infrastructure

Tests that touch `delphix_engine_configuration` with `device_type=OBJECT` need cloud-specific infrastructure. Required only when running the engine_configuration acceptance suite end-to-end.

| Cloud | Bucket / container | Auth mechanism | Required infra |
|---|---|---|---|
| AWS | S3 bucket | Access key/secret **or** IAM role assumed by engine VM | Bucket created, encryption configured, IAM role / key with `s3:*` on bucket |
| Azure | Blob storage container | Storage account access key | Storage account + container, key in test secrets |
| GCP | GCS bucket | Engine VM's service account (no creds passed to TF) | Bucket created, engine VM's GCE service account has `storage.objectAdmin` on bucket |

NTP server reachable from the engine is mandatory for any object-storage scenario.

## Network Topology

Minimum reachability matrix for acceptance tests:

```
[ test runner ]  ──HTTPS──▶  [ DCT ]  ──HTTPS──▶  [ CD engine ]  ──SSH/DB──▶  [ source host ]
                                                                    │
                                                                    └─SSH───▶  [ target env ]
```

- Test runner → DCT: TCP 443 (or 80 if `DCT_HOST_SCHEME=http`).
- DCT → engine(s): TCP 443.
- Engine → source / target hosts: per plugin (typically SSH 22 plus a DB port).

## Sandbox vs Shared DCT

| Mode | Pros | Cons | When to use |
|---|---|---|---|
| Dedicated dev DCT | Free to break, fast iteration | Cost, maintenance | Active development against a feature branch |
| Shared team DCT | One source of truth, lower cost | Cleanup discipline, scheduling conflicts | PR validation, scheduled runs |
| Ephemeral / provisioned per run | Repeatable, isolated | Slow startup (10–20 min) | CI on protected branches |

Never point acceptance tests at a production DCT. Tests create and delete real resources.

## State Hygiene

Acceptance tests leak state if force-killed. The Terraform SDK's `CheckDestroy` callbacks tear resources down on a clean exit, but a SIGKILL or runner timeout leaves orphans behind. Maintain a periodic sweep on shared DCTs:

- Stale VDBs (older than 24h with `tf-test-` prefix).
- Stale environments / dSources.
- Old plugin uploads.

Use the DCT UI or `delphix_*` resources with `terraform destroy` against a leftover state file when possible.

## CI Notes

When configuring CI:

- **Never** commit `DCT_KEY` or `DCT_HOST` — inject via secret store at job start.
- Run `make test` (unit) on every PR. It's fast and free.
- Run `TF_ACC=1 make testacc` only on `main` / release branches or on-demand against a sandbox DCT — both for cost and for the wall-clock time (up to 120 min).
- Pin the Go toolchain via `go.mod`'s `go` directive plus a `setup-go` action with the same version.
- For the cloud object-storage tests, scope each cloud's credentials so a leaked key cannot harm production buckets.

## Tooling

- `terraform` CLI 1.x — for `examples/` validation (`terraform validate` / `fmt`).
- `dev_overrides` in `~/.terraformrc` — to validate examples against a locally-built provider binary before publishing.
- `make install` — drops the locally built binary into `~/.terraform.d/plugins/delphix.com/dct/delphix/<version>/<os_arch>/` for use with `dev_overrides`.

## When Adding a New Resource

1. List the new infrastructure dependencies (new source plugin? new cloud? new auth path?) in the design doc's `## Affected Components`.
2. Update this file if a new component is now required (e.g., a new source DB type, a new cloud).
3. Make sure CI / sandbox DCTs are updated *before* the acceptance tests land.

---

## Story: DLPXECO-13975 — engine_configuration on GCP (CD + CC + Block regression)

Testing counterpart to the GCP Object Storage feature delivered in DLPXECO-13662 (commit `263cf5e`, provider v4.3.0). The acceptance tests below already exist in [resource_engine_configuration_test.go](../../internal/provider/resource_engine_configuration_test.go); this section captures the *infra* preconditions for running them as a single coordinated suite.

### Test matrix in scope

| # | Engine type | Storage | Auth | Test function | `-run` regex |
|---|---|---|---|---|---|
| 1 | CD | GCS (OBJECT) | Engine VM service account | `TestAccEngineConfiguration_gcpObjectStorage` | `^TestAccEngineConfiguration_gcpObjectStorage$` |
| 2 | CC | GCS (OBJECT) | Engine VM service account | `TestAccEngineConfiguration_gcpObjectStorage_CC` | `^TestAccEngineConfiguration_gcpObjectStorage_CC$` |
| 3 | CD | Block | — | `TestAccEngineConfiguration_blockDevice` | `^TestAccEngineConfiguration_blockDevice$` |

Scenarios 1 and 2 verify the GCS feature itself. Scenario 3 is a **regression check** — running the existing block-storage test against a GCP-hosted engine VM confirms the GCS code paths did not break the non-object-storage flow. All three target a freshly-cloned engine (`engine_configuration` is destructive first-boot config — cannot re-run against a configured engine).

### Engine version requirement (FedEx validation criteria)

For the GCS scenarios, the engine VM must run **Delphix Engine 2025.5.0.2 or newer**. Older engines do not expose the GCP-specific fields on the `engine_configuration` API surface and the tests will fail at the API layer rather than at Terraform schema validation. Set `ENGINE_GOLDEN_IMAGE` to a 2025.5.0.2+ image group such as `dlpx-dose-2026.2.0.0` (current default in [settings.local.json](../settings.local.json)).

The generic provider-level minimum is still Engine v29.0.0.0+ as noted earlier in this file — the 2025.5.0.2+ floor applies *only* to the GCS-on-GCP path covered by this story.

### GCS bucket convention

`GCP_BUCKET_NAME` is stored as the template `"dcoa-prod-${ENGINE_NAME}"` in [settings.local.json](../settings.local.json) so that updating `ENGINE_NAME` alone re-points the bucket. The assistant **must expand `${ENGINE_NAME}` at run time** the same way it does for `DELPHIX_ENGINE_HOST`:

1. Read `ENGINE_NAME` from `.claude/settings.local.json`.
2. Read the raw `GCP_BUCKET_NAME` template (e.g. `dcoa-prod-${ENGINE_NAME}`).
3. Substitute `${ENGINE_NAME}` with the value from step 1.
4. Export the resolved value as `GCP_BUCKET_NAME` before running `go test`.

The bucket itself is provisioned alongside the engine VM by `dc clone-latest` (the golden image's cloud-init grants the engine's GCE service account `storage.objectAdmin` on the matching bucket). No separate Terraform or `gsutil` step is required from the test runner.

### CC variant — extra required fields

The `_CC` test (`TestAccEngineConfiguration_gcpObjectStorage_CC`) configures a Continuous Compliance engine and exercises four additional schema fields beyond the CD path:

| Field | Test value | Notes |
|---|---|---|
| `compliance_user` | `admin` | Initial compliance UI user |
| `compliance_password` | `Admin-12` | Initial compliance UI password |
| `compliance_new_password` | `Admin@45` | New password set on first boot |
| `compliance_email` | `compliance@example.com` | Compliance admin email |

Values above are the literals embedded in `testAccEngineConfigurationGCPObjectStorageCC` — they are not configurable via env vars and live entirely inside the test's Terraform config string. Override by editing the test if a real CC engine needs different first-boot credentials.

The CC engine must be cloned from a CC-flavored golden image (not a CD image). Today the same `dlpx-dose-*` group typically ships both — confirm `dc list-images` shows a CC variant before running scenario 2.

### Required env vars for this story

Beyond the standard `DCT_KEY` / `DCT_HOST` / `TF_ACC=1` baseline:

| Var | Value for this story |
|---|---|
| `DELPHIX_TEST_RESOURCE` | `engine_configuration` |
| `DELPHIX_ENGINE_TYPE` | `CC` for scenario 2; `CD` for scenarios 1 and 3 (or leave empty to run the union) |
| `DELPHIX_ENGINE_CLOUD` | `GCP` (also drives `dc clone-latest --cloud`) |
| `ENGINE_NAME` | A free `dc` VM name, e.g. `sho-gcp` — drives both `DELPHIX_ENGINE_HOST` and `GCP_BUCKET_NAME` via template expansion |
| `ENGINE_GOLDEN_IMAGE` | `dlpx-dose-2026.2.0.0` (or any 2025.5.0.2+ image with both CD and CC variants) |
| `GCP_BUCKET_NAME` | Template `dcoa-prod-${ENGINE_NAME}` — substitute at run time |
| `DELPHIX_ENGINE_HOST` | Template `http://${ENGINE_NAME}.dlpxdc.co` — substitute at run time |

Scenario 3 (block-device regression) reuses the same engine clone but does not need `GCP_BUCKET_NAME` — the test skips the env-var check when running block.

### Run order

Because each scenario triggers destructive first-boot config on the engine, the three scenarios cannot share a single engine VM. Run them sequentially with a fresh `dc clone-latest` per scenario:

```bash
# Scenario 1 — CD + GCS
DELPHIX_ENGINE_TYPE=CD DELPHIX_ENGINE_CLOUD=GCP   go test -v -timeout 120m -run '^TestAccEngineConfiguration_gcpObjectStorage$' ./internal/provider/
dc destroy "$ENGINE_NAME" -w

# Scenario 2 — CC + GCS (re-clone with CC image variant)
dc clone-latest "$ENGINE_GOLDEN_IMAGE" "$ENGINE_NAME" -w --cloud GCP
DELPHIX_ENGINE_TYPE=CC DELPHIX_ENGINE_CLOUD=GCP   go test -v -timeout 120m -run '^TestAccEngineConfiguration_gcpObjectStorage_CC$' ./internal/provider/
dc destroy "$ENGINE_NAME" -w

# Scenario 3 — CD + Block (regression on a GCP-hosted VM)
dc clone-latest "$ENGINE_GOLDEN_IMAGE" "$ENGINE_NAME" -w --cloud GCP
DELPHIX_ENGINE_TYPE=CD DELPHIX_ENGINE_CLOUD=BLOCK   go test -v -timeout 120m -run '^TestAccEngineConfiguration_blockDevice$' ./internal/provider/
dc destroy "$ENGINE_NAME" -w
```

Total wall-clock budget: ~30–45 min (3 × clone @ 3–5 min + 3 × first-boot config @ 5–10 min + teardown).

### Optional — manual round-trip via `make_and_clean.sh`

For an end-to-end manual smoke before/after the automated suite, the [examples/engine_configuration/make_and_clean.sh](../../examples/engine_configuration/make_and_clean.sh) helper runs `terraform init && terraform apply && terraform destroy` against [examples/engine_configuration/main.tf](../../examples/engine_configuration/main.tf), which already contains a GCP Object Storage block (added in commit `263cf5e`). This is a complement to — not a replacement for — the acceptance test suite, and exercises the user-facing example exactly as a customer would.

Before running `make_and_clean.sh`:

1. Export `DCT_KEY`, `DCT_HOST` to the shell.
2. Edit `main.tf` to set `engine_host` to the cloned engine's URL.
3. Confirm only the GCP scenario block is uncommented; comment out AWS / Azure / block scenarios you do not intend to apply.

### Acceptance criteria coverage

| Story criterion | How test-infra supports it |
|---|---|
| All three scenarios pass `go test -v -timeout 120m` | Run-order recipe above; pre-flight checks gate on env vars per the main flow |
| CRUD lifecycle: create → read → (no update on first-boot) → destroy with no leaks | `dc destroy "$ENGINE_NAME" -w` after each scenario; the engine and its bucket are torn down together |
| Test evidence at `docs/DLPXECO-13975-test-evidence.md` | Produced by phase `test` (consumes this infra); see [testing.md](testing.md) Test Evidence section |
| Coverage delta at `docs/DLPXECO-13975-coverage.md` | Produced by phase `test`; this file does not own it |
| Failures triaged with root cause + fix + re-test | Standard testing.md "Things to Avoid" + Phase 6 (test) contract |
| `examples/engine_configuration/main.tf` GCP scenario manually applied + destroyed | The optional `make_and_clean.sh` flow above |

