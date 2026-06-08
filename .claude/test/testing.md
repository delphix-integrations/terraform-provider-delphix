# Testing Rules — terraform-provider-delphix

## Test Tiers

This provider has two distinct test tiers. Both run via the Terraform Plugin SDK v2 testing framework.

| Tier | Command | Scope | Network | Default in CI |
|---|---|---|---|---|
| Unit | `make test` | Schema validation, helper functions, regex/utility code, provider-level checks (no remote calls) | None | Yes |
| Acceptance | `TF_ACC=1 make testacc` | Full CRUD against a live DCT instance — provisions/deletes real resources | Live DCT + Delphix engines | No (requires credentials) |

## CI-Enforced Gate 

PRs to `main` or `develop` are blocked by the `ci / unit-tests` GitHub Actions check.
Three rules apply:

1. **Every PR runs the unit-test suite** — CI executes
   `go test ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s`
   (the local equivalent is `make test`; CI uses a 300 s timeout vs. `make test`'s 30 s).

2. **Coverage must not drop below 2%** — the `COVERAGE_THRESHOLD` env var in
   `.github/workflows/ci.yml` is currently set to `2`. Adding or modifying code
   without a corresponding unit test can lower the total and fail CI. The baseline
   when the gate was set was 2.3% (measured 2026-06-06, unit tests only, no
   `TF_ACC`). Do not lower the threshold without team agreement.

3. **Acceptance tests are NOT run in CI** — `TestAcc*` tests are skipped
   automatically because CI does not export `TF_ACC=1`. Run them locally before
   submitting: `TF_ACC=1 make testacc`.

## Required Environment Variables

> **Source of truth:** All env vars required by tests are declared in [`.claude/settings.local.json`](../settings.local.json) under the `env` block. Populate that file with your sandbox values — the file is gitignored so credentials stay local. The harness loads these into the shell environment for every test run. **Do not commit credentials inline in test files or in CI configs without an injected secret store.**

Acceptance tests fail fast in `testAccPreCheck` if these are unset:

| Variable | Required | Purpose |
|---|---|---|
| `DCT_KEY` | Acceptance | DCT API key (sensitive) |
| `DCT_HOST` | Acceptance | DCT hostname (e.g. `my-dct.example.com`) |
| `DCT_HOST_SCHEME` | Optional | `https` (default) or `http` |
| `DCT_TLS_INSECURE_SKIP` | Optional | `true` to skip TLS verification (dev only — never production) |
| `TF_ACC` | Acceptance | Must be `1` to enable acceptance test runs (SDK gate) |

## Resource Selector — `DELPHIX_TEST_RESOURCE`

`.claude/settings.local.json` includes a `DELPHIX_TEST_RESOURCE` env var that names which component's tests to run. The assistant reads this to pick a `-run` regex automatically — see the "Selecting the target" section in [`test-infra.md`](../test-infra.md).

Common values: `engine_configuration`, `vdb`, `environment`, `oracle_dsource`, `database_postgresql`, `engine_registration`, `appdata_dsource`, `vdb_group`, `database_plugin`, `all`.

Leave it empty to be prompted for the target each run. An explicit chat instruction ("run engine_configuration tests") always overrides this var.

## Run Commands

```bash
# Unit tests — fast, no network. Run on every change.
make test
# Internally: go test -timeout=30s -parallel=4 ./...

# Acceptance tests — slow, requires live DCT.
TF_ACC=1 DCT_KEY=<key> DCT_HOST=<host> make testacc
# Internally: TF_ACC=1 go test -v -timeout 120m ./...

# Run a single resource's acceptance tests
TF_ACC=1 DCT_KEY=<key> DCT_HOST=<host> go test -v -timeout 120m \
  -run TestAccDelphixVDB ./internal/provider/

# Run a single test function
TF_ACC=1 DCT_KEY=<key> DCT_HOST=<host> go test -v -timeout 120m \
  -run TestAccDelphixVDB_basic ./internal/provider/
```

## Timeouts

- Unit tests: **30s** total per package, parallel=4
- Acceptance tests: **120m** total (resources can take many minutes to provision)
- Per-resource Terraform timeouts: 20m create / 20m update / 20m delete (defined in resource schemas)

## Test File Conventions

Test files live alongside the source they cover, in `internal/provider/`, named `<source>_test.go`.

| Test type | Function name pattern | Gate |
|---|---|---|
| Unit | `TestProvider`, `TestProvider_impl`, `Test<Helper>` | Always runs |
| Acceptance | `TestAcc<Resource>_<scenario>` | Skipped unless `TF_ACC=1` |
| Pre-check | `testAccPreCheck(t)` | First call in any acceptance test — verifies env vars |

Existing test files (current count of acceptance test files: 11):
- [provider_test.go](../../internal/provider/provider_test.go) — provider-level config + pre-check helper
- [resource_vdb_test.go](../../internal/provider/resource_vdb_test.go)
- [resource_vdb_group_test.go](../../internal/provider/resource_vdb_group_test.go)
- [resource_environment_test.go](../../internal/provider/resource_environment_test.go)
- [resource_appdata_dsource_test.go](../../internal/provider/resource_appdata_dsource_test.go)
- [resource_oracle_dsource_test.go](../../internal/provider/resource_oracle_dsource_test.go)
- [resource_database_postgresql_test.go](../../internal/provider/resource_database_postgresql_test.go)
- [resource_engine_configuration_test.go](../../internal/provider/resource_engine_configuration_test.go)
- [resource_engine_registration_test.go](../../internal/provider/resource_engine_registration_test.go)
- [resource_engine_plugin_upload_test.go](../../internal/provider/resource_engine_plugin_upload_test.go)
- [security_test.go](../../internal/provider/security_test.go)

## When Adding a New Resource

1. Create `resource_<name>_test.go` in `internal/provider/`.
2. Add at least one unit test that calls `Provider().InternalValidate()` for schema correctness.
3. Add at least one acceptance test (`TestAcc<Resource>_basic`) that exercises full CRUD: Create → Read → Update → Delete.
4. Acceptance tests **must** call `testAccPreCheck(t)` first.
5. Use `resource.TestCase` with `Providers: testAccProviders` from `provider_test.go`.
6. Each acceptance test step needs both `Config` (HCL) and `Check` (assertions).

## When Adding a Helper / Utility

1. If the helper has logic worth testing in isolation (parsers, regex validators, state machines, format converters), write a **unit test** in the closest `*_test.go`.
2. Pure unit tests (no DCT calls) should run under `make test` — must complete in well under 30s.
3. Example: [engine_api_utility.go](../../internal/provider/engine_api_utility.go) `validateStorageSize` regex — add table-driven tests for each accepted/rejected size format.

## When Modifying Examples or Docs

Terraform examples under `examples/` are not part of `make test`. Validate them separately:

```bash
cd examples/<example_dir>
terraform fmt -recursive    # formatting check
terraform init              # requires the provider built locally — see below
terraform validate          # syntax + provider-schema validation
```

For locally-built provider validation, set `dev_overrides` in `~/.terraformrc`:

```hcl
provider_installation {
  dev_overrides {
    "delphix.com/dct/delphix" = "/path/to/your/local/build"
  }
  direct {}
}
```

## Test Evidence for the SDD Workflow

Phase 6 (test) of `feature-implement` writes:

- `docs/$NAME-test-evidence.md` — scenarios run, versions, outcomes (PASS/FAIL/SKIPPED)
- `docs/$NAME-coverage.md` — per-component coverage delta

For each FAIL: capture root cause, fix applied, re-test result. Do not mark a phase complete with failing tests.

## Things to Avoid

- **Never** commit `DCT_KEY`, `DCT_HOST`, passwords, or hostnames to test files. Use `os.Getenv` everywhere.
- **Never** disable TLS verification (`DCT_TLS_INSECURE_SKIP=true`) in committed test fixtures.
- **Never** mock `PollJobStatus` — async job semantics are part of what acceptance tests verify.
- **Don't** run `testacc` in CI without provisioning credentials and a sandbox DCT — it will create real resources.
- **Don't** leave orphaned VDBs/environments behind from a failed acceptance run; the SDK's `CheckDestroy` should handle cleanup, but verify manually if a run is force-killed.
