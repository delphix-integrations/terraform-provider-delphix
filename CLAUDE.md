# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Build
make build          # Produces ./terraform-provider-delphix binary
make install        # Builds and installs to ~/.terraform.d/plugins/delphix.com/dct/delphix/4.2.1/darwin_arm64
make release        # Multi-platform binaries to ./bin/

# Test
make test           # Unit tests (parallel=4, timeout=30s)
make testacc        # Acceptance tests (requires DCT_KEY, DCT_HOST env vars; timeout=120m)

# Run a single test
go test ./internal/provider -run TestProvider -v
TF_ACC=1 go test ./internal/provider -run TestAccVdb_provision_positive -v -timeout 120m
```

Acceptance tests also require resource-specific env vars (e.g., `DATASOURCE_ID`, `ENVIRONMENT_ID`) depending on which resource is being tested.

## Architecture

This is a **Terraform Plugin SDK v2** provider that wraps the **Delphix Control Tower (DCT) API** via the `github.com/delphix/dct-sdk-go/v25` Go SDK.

### Provider Authentication

The provider authenticates to DCT using an API key. Configuration parameters:
- `host` / `DCT_HOST` — DCT hostname
- `key` / `DCT_KEY` — API key (sent as `Authorization: apk <key>`)
- `host_scheme` / `DCT_HOST_SCHEME` — defaults to `https`
- `tls_insecure_skip` / `DCT_TLS_INSECURE_SKIP` — skip TLS verification

The provider validates connectivity on configure by calling `GetRegisteredEngines()`.

### Resources

All provider logic lives in `internal/provider/`. Resources implemented:
- `delphix_vdb` — Virtual Database provisioning (most complex resource)
- `delphix_vdb_group` — VDB group management
- `delphix_environment` — Environment (host) management
- `delphix_appdata_dsource` — Application Data dSource ingestion
- `delphix_oracle_dsource` — Oracle dSource ingestion
- `delphix_database_postgresql` — PostgreSQL source registration
- `delphix_engine_configuration` — Engine configuration
- `delphix_engine_dct_registration` — Register engines with DCT
- `delphix_database_plugin` — Database plugin upload

### Key Files

| File | Purpose |
|------|---------|
| `internal/provider/provider.go` | Provider definition, DCT client setup |
| `internal/provider/commons.go` | Per-resource maps of updatable and destructive fields |
| `internal/provider/utility.go` | `PollJobStatus()` for async job polling (10s interval) |
| `internal/provider/models.go` | Shared data structures |
| `internal/provider/logging.go` | `InfoLog`, `WarnLog`, `ErrorLog` wrappers over `tflog` |
| `internal/provider/engine_api.go` | Direct engine API calls (sessions, login) |
| `internal/provider/security.go` | `SecureClearByteSlice()` for wiping credentials from memory |

### Patterns

**Update strategy**: `commons.go` maintains maps (e.g., `updatableVdbKeys`, `isDestructiveVdbUpdate`) that determine whether a field change triggers an in-place update or forces resource recreation.

**Async operations**: All DCT operations that return a job ID are polled via `PollJobStatus()` until terminal state (COMPLETED/FAILED). Default timeouts: 20 min for Create/Update/Delete.

**Tag handling**: `CustomizeDiffTags` is applied to all resources via `CustomizeDiff`. Resources support `ignore_tag_changes` to suppress tag drift.

**Logging**: Uses `tflog` with `[DELPHIX] [INFO/WARN/ERROR]` prefix convention throughout.
