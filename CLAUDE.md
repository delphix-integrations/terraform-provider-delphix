# CLAUDE.md — Terraform Provider Delphix

## Project Overview

This is a HashiCorp Terraform provider for [Delphix Control Tower (DCT)](https://help.delphix.com/eh/current/content/terraform.htm). It enables Infrastructure-as-Code management of Delphix Continuous Data & Continuous Compliance resources — virtual databases (VDBs), dSources, environments, engine registration, and engine configuration — by wrapping the DCT REST API via the `dct-sdk-go` SDK.

**Current version:** 4.3.1  
**Required Go version:** 1.25+  
**DCT requirement:** v2025.2.0+  
**Delphix Engine requirement:** v29.0.0.0+

---

## Repository Layout

```
terraform-provider-delphix/
├── main.go                    # Provider entry point; sets version, enables debug mode
├── go.mod / go.sum            # Module definition; Go 1.25
├── GNUmakefile                # Build, install, test targets
├── .goreleaser.yml            # Multi-platform release config (version 4.3.1)
├── internal/provider/         # All provider logic lives here (single package)
│   ├── provider.go            # Schema, resource/datasource registration
│   ├── resource_vdb.go        # delphix_vdb resource
│   ├── resource_vdb_group.go  # delphix_vdb_group resource
│   ├── resource_environment.go
│   ├── resource_appdata_dsource.go
│   ├── resource_oracle_dsource.go
│   ├── resource_database_postgresql.go
│   ├── resource_engine_configuration.go
│   ├── resource_engine_registration.go
│   ├── resource_database_plugin.go
│   ├── commons.go             # Constants: job states, updatable keys, etc.
│   ├── utility.go             # Shared helper functions
│   ├── engine_api.go          # Engine-level API utilities
│   ├── security.go            # TLS/HTTP client configuration
│   └── *_test.go              # Unit and acceptance tests
├── examples/                  # Runnable Terraform examples per resource
├── docs/                      # Auto-generated Terraform Registry docs
└── .github/workflows/         # release.yml, codeql.yml
```

---

## Build & Test Commands

```bash
# Build for current platform (darwin_arm64 by default)
make build

# Install into local Terraform plugin directory
make install

# Cross-compile for all platforms (Linux, Darwin, Windows, FreeBSD, etc.)
make release

# Run unit tests (parallel=4, timeout=30s)
make test

# Run acceptance tests against a live DCT instance (timeout=120m)
TF_ACC=1 make testacc
```

Acceptance tests require a running DCT instance and the environment variables below. Do not run `testacc` without those configured.

---

## Environment Variables

| Variable | Purpose |
|---|---|
| `DCT_KEY` | DCT API key (sensitive) |
| `DCT_HOST` | DCT hostname (e.g. `my-dct.example.com`) |
| `DCT_HOST_SCHEME` | `https` (default) or `http` |
| `DCT_TLS_INSECURE_SKIP` | `true` to skip TLS verification (dev/test only) |

---

## Provider Resources

| Resource | Description |
|---|---|
| `delphix_vdb` | Provision/manage Virtual Databases |
| `delphix_vdb_group` | Group VDBs together with tag management |
| `delphix_environment` | Manage Delphix environments (hosts) |
| `delphix_appdata_dsource` | Application data sources |
| `delphix_oracle_dsource` | Oracle data sources |
| `delphix_database_postgresql` | PostgreSQL database objects |
| `delphix_engine_configuration` | Engine storage (Block + Object: AWS/Azure/GCP), NTP, SMTP, DNS, SSO config |
| `delphix_engine_dct_registration` | Register engines with DCT |
| `delphix_database_plugin` | Upload/manage database plugins |

There are currently **no data sources** — the provider is resource-only.

---

## Key Architectural Patterns

### Async Job Polling
DCT API operations are asynchronous. Every mutating call returns a job ID. The provider polls via `PollJobStatus()` in [utility.go](internal/provider/utility.go) with a configurable sleep interval (5–20 s) until the job reaches `COMPLETED`, `FAILED`, or `CANCELED`.

### Resource CRUD Pattern
All resources use Terraform Plugin SDK v2 context-aware methods:
```go
CreateContext, ReadContext, UpdateContext, DeleteContext
```

### Schema Conventions
- **Optional + Computed** for fields that DCT may populate if omitted.
- **Sensitive: true** for secrets (API keys, passwords).
- Timeouts are 20 minutes for create/update/delete by default.
- `CustomizeDiffTags` is used for tag-aware diff suppression.

### Logging
Use `tflog` from `github.com/hashicorp/terraform-plugin-log/tflog`. Prefix log messages with `[DELPHIX]`:
```go
tflog.Info(ctx, "[DELPHIX] Creating VDB")
tflog.Warn(ctx, "[DELPHIX] ...")
tflog.Error(ctx, "[DELPHIX] ...")
```

### State Constants
Job state constants (`PENDING`, `STARTED`, `FAILED`, `COMPLETED`, `CANCELED`, etc.) are defined in [commons.go](internal/provider/commons.go). Use these — do not hardcode strings.

---

## Adding a New Resource

1. Create `internal/provider/resource_<name>.go` using the `provider` package.
2. Implement `CreateContext`, `ReadContext`, `UpdateContext`, `DeleteContext`.
3. Register the resource in `provider.go` under `ResourcesMap`.
4. Add acceptance tests in `resource_<name>_test.go`.
5. Add an example directory under `examples/<name>/`.
6. Docs are auto-generated from schema descriptions — keep them accurate and complete.

---

## Contribution Notes

- This project is **not accepting external contributions** (per README).
- Internal PRs: branch from `develop` for features, `main` for bugfixes.
- Commits must be **signed** (GPG or SSH).
- A CLA is required via cla-assistant for any external contribution.
- PRs must follow [pull_request_template.md](pull_request_template.md): Context, Problem, Solution, Testing sections are required.

---

## Release Process

Releases are fully automated via GoReleaser triggered by a `v*` tag push:
- Builds for FreeBSD, Windows, Linux, Darwin × amd64/386/arm/arm64.
- GPG-signs checksums.
- Publishes to Terraform Registry via `.goreleaser.yml`.

Do **not** manually edit the version in [GNUmakefile](GNUmakefile) or [.goreleaser.yml](.goreleaser.yml) — these are managed as part of the release workflow.

---

## Security

- Never disable TLS verification (`DCT_TLS_INSECURE_SKIP=true`) in production configs.
- API keys must be passed via `DCT_KEY` env var, never hardcoded in `.tf` files.
- HTTP client security is configured in [security.go](internal/provider/security.go).

---

## Useful Links

- [Terraform Registry Docs](https://registry.terraform.io/providers/delphix-integrations/delphix/latest/docs)
- [Delphix Ecosystem Terraform Docs](https://help.delphix.com/eh/current/content/terraform.htm)
- [DCT SDK Go](https://github.com/delphix/dct-sdk-go)
- [Terraform Plugin SDK v2 Docs](https://developer.hashicorp.com/terraform/plugin/sdkv2)

---

## CI Contract

The CI workflow is defined in `.github/workflows/ci.yml` and runs automatically on every
pull request targeting `main` or `develop`, and on every push to `main` or `develop`.

### Workflow Summary

| Item | Value |
|---|---|
| Workflow file | `.github/workflows/ci.yml` |
| Workflow name | `ci` |
| Job name | `unit-tests` |
| Status-check string | `ci / unit-tests` |
| Trigger | `pull_request` to `main` or `develop`; `push` to `main` or `develop` |
| Runner | `ubuntu-latest` |
| Go version | Auto-detected from `go.mod` via `actions/setup-go@v5` |
| Test command | `go test ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s` |
| Coverage artifact | `coverage-report` (7-day retention) |
| Coverage threshold | `COVERAGE_THRESHOLD` in `ci.yml` env block (current: `2%`) |
| Baseline at threshold set | `2.3%` (measured 2026-06-06; unit tests only, no `TF_ACC`) |

Acceptance tests (`TF_ACC=1`) are **not run in CI** — they require live DCT infrastructure
and are excluded automatically because the workflow does not export `TF_ACC`.

### Running the Equivalent Check Locally

```bash
go test ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s
go tool cover -func=coverage.out | tail -1
```

To see a per-function breakdown:
```bash
go tool cover -func=coverage.out | sort -t$'\t' -k3 -n
```

To see an HTML report in your browser:
```bash
go tool cover -html=coverage.out
```

### Updating the Coverage Threshold

1. Measure current coverage locally using the commands above.
2. Edit `COVERAGE_THRESHOLD` in `.github/workflows/ci.yml`.
3. Document the old value, new value, and reason in the PR description.
4. The change takes effect on the next CI run.

Do not lower the threshold without team agreement.

### Branch Protection

The following settings must be applied to the `main` branch in
**GitHub Settings → Branches → Branch protection rules** by a repo maintainer.
This document describes the required contract; it does not programmatically
configure GitHub.

| Setting | Required Value |
|---|---|
| Require status checks to pass before merging | Enabled |
| Status check to require | `ci / unit-tests` |
| Require branches to be up to date before merging | Enabled |
| Who can bypass | No one (recommended) |

**The exact status-check string to enter in GitHub's branch protection UI is: `ci / unit-tests`**

### CI Gates

Three rules enforced on every PR targeting `main` or `develop`:

1. **Unit-test workflow runs on every PR** — CI executes
   `go test ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s`
   (the local equivalent is `make test`, though CI uses 300 s vs. `make test`'s
   30 s timeout). The workflow must pass before a PR can be merged.

2. **Coverage gate** — if total unit-test coverage drops below **2%**
   (the `COVERAGE_THRESHOLD` in `ci.yml`), the `ci / unit-tests` check fails.
   Adding or modifying a feature without a corresponding unit test will lower
   coverage and block the PR. Raise the threshold — never lower it — when
   adding new coverage.

3. **Acceptance tests are NOT enforced in CI** — `TestAcc*` tests require a
   live DCT instance and are excluded automatically because CI does not set
   `TF_ACC=1`. Run them locally with `TF_ACC=1 make testacc` before marking
   a feature complete.

### Drift Management

The values in the Workflow Summary table above (threshold, status-check
string, trigger branches, workflow name, job name) are duplicated from
`.github/workflows/ci.yml`. Any future PR that changes those values in
`ci.yml` MUST also update the corresponding rows in this section in the
same PR. This is a process rule, not a tooling-enforced check — reviewers
should call out mismatches during PR review.
