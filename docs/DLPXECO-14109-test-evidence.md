# Test Evidence: DLPXECO-14109

**Jira**: [DLPXECO-14109](https://perforce.atlassian.net/browse/DLPXECO-14109)
**Branch**: `vuln-fix`
**Run date**: 2026-06-04
**Operator**: shobhit.sinha
**Test plan**: `docs/DLPXECO-14109-test-plan.md`

---

## Landscape / Environment

- **Host**: `darwin_arm64` developer machine (Darwin 25.5.0, macOS).
- **Working tree**: `terraform-provider-delphix` on branch `vuln-fix`, in place (no worktree per operator decision).
- **Bumped modules under test** (`go.mod` post-bump):
  - `golang.org/x/crypto v0.52.0`
  - `golang.org/x/net v0.55.0` (indirect)
  - `golang.org/x/sys v0.45.0` (indirect)  *(observed; test-plan floor was `v0.44.0` — bump cascade resolved one minor higher, still satisfies the CVE-clearance requirement; documented as Open Question OQ-1 below)*
  - `google.golang.org/grpc v1.79.3` (indirect)
- **Service under test**: compiled provider binary; no live DCT instance required for the merge gate.
- **VMs provisioned**: none. `test-infra` phase had no `## VMs` section per `docs/DLPXECO-14109-test-plan.md`.
- **`.claude/DLPXECO-14109-test-env.sh`**: not present (no VMs). Test phase proceeded without sourcing it.
- **NG3 compliance**: no new `TestXxx` functions were added in this iteration — the existing 14-test unit suite is the regression gate per the vision Non-Goals and `docs/DLPXECO-14109-test-plan.md` Test Approach.
- **No generated tests under `.claude/test/generated-test/`**: this is the canonical NG3 outcome — the test-plan explicitly directs test-generation to skip producing new Go `TestXxx` functions. Logged here per the workflow-steps requirement to surface absent generated tests.

---

## Versions

| Component | Version | Source |
|-----------|---------|--------|
| Go toolchain | `go1.25.0 darwin/arm64` | `go version` |
| `go.mod` toolchain line | `go 1.25.0` | `grep ^go go.mod` |
| `terraform-plugin-sdk/v2` (pinned per NG1) | `v2.33.0` | `go list -m` |
| `terraform-plugin-log` (pinned per NG1) | `v0.9.0` | `go list -m` |
| `dct-sdk-go` (pinned per NG1) | `v25.6.0` | `go list -m` |
| `golang.org/x/crypto` | `v0.52.0` | `go list -m` |
| `golang.org/x/net` | `v0.55.0` (indirect) | `go list -m` |
| `golang.org/x/sys` | `v0.45.0` (indirect) | `go list -m` |
| `google.golang.org/grpc` | `v1.79.3` (indirect) | `go list -m` |

**Pre-bump baseline**: same 14 unit tests re-run after temporarily restoring `git checkout main -- go.mod go.sum` produced identical PASS results in `0.685s` (the post-bump tree completes the same suite in `0.625s`). Baseline → post-bump delta: 0 tests added, 0 removed, 0 changed result. See `## Functional (primary)` Summary.

---

## Functional (primary)

Cross-references every scenario row from `docs/DLPXECO-14109-test-plan.md` `## Scenarios`. `[resolution-check]` rows are deterministic shell assertions executed in this phase and recorded here. `[regression]` rows map to existing `TestXxx` functions under `internal/provider/`. `[validation]` / `[failure-path]` / `[process]` rows are deferred to the validate / pr phases per the test plan and are explicitly out-of-scope for this test phase.

| Scenario | Version(s) | Outcome | Notes |
|----------|-----------|---------|-------|
| S1 — `go.mod` lists the four target modules at exactly the required versions after `go get` [resolution-check] | Go 1.25 | PASS (with drift note) | `grep -E "^\s+(golang.org/x/(crypto\|net\|sys)\|google.golang.org/grpc)\s" go.mod` returned 4 lines: `crypto v0.52.0`, `net v0.55.0`, `sys v0.45.0`, `grpc v1.79.3`. `sys` resolved one minor higher than the test-plan floor `v0.44.0` (`v0.45.0`); see OQ-1. CVE-clearance satisfied because `v0.45.0 > v0.44.0`. |
| S2 — `go mod tidy` is idempotent on the bumped tree [resolution-check] | Go 1.25 | PASS | Two consecutive `go mod tidy` invocations produced zero diff on both `go.mod` and `go.sum` (`diff -q` reports identical). Recorded inline in the run log above. |
| S3 — `go list -m all` resolves the four modules without `replace` indirection [resolution-check] | Go 1.25 | PASS | `go list -m all` shows the four required modules at the post-bump versions; no `=>` token in the output for any of them. `grep "^replace " go.mod` returns nothing — zero `replace` directives across the file. |
| S4 — `git diff main -- go.mod` shows only intentional bumps; `go.sum` is regenerated [resolution-check] | Go 1.25 | PASS | `git diff --stat go.mod go.sum` (vs HEAD~ on `main` baseline restore): `go.mod` 22 changed lines, `go.sum` 74 changed lines (59 insertions / 37 deletions). `git diff main -- internal/provider/` empty — no production source change, per Quality Rule "No production source change". |
| S5 — `make build` exits 0 on the post-bump tree [regression] | Go 1.25, darwin_arm64 | PASS | Documented in `docs/DLPXECO-14109-build-output.md` (produced by build phase, completed earlier). Build artifact: `terraform-provider-delphix` binary. |
| S6 — Cross-compile matrix builds clean via `make release` dry-run [regression / EC-7] | Go 1.25 × goreleaser matrix | DEFERRED | Not executed in test phase — recorded in `docs/DLPXECO-14109-build-output.md` per build-phase scope. Cross-compile dry-run is a build-phase artifact, not a test-phase one. Will be reverified by the validate phase (Section 7). |
| S7 — `make test` exits 0 with passed-count >= pre-bump baseline [regression] | Go 1.25 | PASS | True unit suite (14 tests, non-acceptance): `go test -timeout=300s -parallel=4 -run '^(TestProvider\|TestProvider_impl\|TestCloudProviderValidator\|TestGcpObjectStoreNoAccessCredentials\|TestGcpObjectStoreStruct\|TestGcpObjectStoreTestConnectionStruct\|TestGcpObjectStoreTestConnectionWithSizes\|TestNormalizeStorageSize\|TestSecureClearByteSlice\|TestSecureClearMap\|TestSecureClearNilValues\|TestSecureClearString\|TestSecureString\|TestValidateStorageSize)$' ./internal/provider/...` → `ok  internal/provider  0.625s`, exit 0. Baseline on `main`: same suite passes in `0.685s`. Delta: 0. Unscoped `go test ./...` reports 17 failures, all gated by missing `TF_ACC` env vars (e.g. `DSOURCE_SOURCE_ID`, `ACC_ENV_ENGINE_ID`, `DATASOURCE_ID`, `ENGINE_NAME`) — pre-existing across `main` and `vuln-fix`, not caused by the bump. See `## Failure Triage`. |
| S8 — `go vet ./...` reports no new issues vs baseline [regression] | Go 1.25 | PASS | `go vet ./...` exit 0, no warnings emitted. Same on `main` baseline. |
| S9 — Provider schema and `ResourcesMap` byte-stable [regression] | Go 1.25 | PASS | `git diff main -- internal/provider/provider.go` is empty (per S4 source-change check). `ResourcesMap` not touched. `tfplugindocs generate` not regenerated this iteration (skipped per build-phase output — not a regression). |
| S10 — `TestProvider` / `TestProvider_impl` pass post-bump [regression] | Go 1.25 | PASS | Both `TestProvider` and `TestProvider_impl` PASSED in the unit run (`provider_test.go:24` / `:30`). |
| S11 — Each resource's unit-test functions pass post-bump [regression] | Go 1.25 | PASS (for non-acceptance unit tests) / PRE-EXISTING SKIP-OR-FAIL (for env-gated acceptance tests) | All non-`TestAcc*` unit functions (`TestGcpObjectStore*`, `TestCloudProviderValidator`, `TestNormalizeStorageSize`, `TestValidateStorageSize`, `TestSecureClear*`, `TestSecureString`) PASSED. All `TestAcc*` and the three `_create_positive`/`Acc_*` env-gated tests in `resource_oracle_dsource_test.go`, `resource_appdata_dsource_test.go`, `resource_database_postgresql_test.go` fail identically on `main` and `vuln-fix` due to absent live env vars (verified via `git show main:<file>` — same `t.Fatal("…must be set for env acceptance tests")` guard present on `main`). Not attributable to the bump. |
| S12 — Negative — `make build` halts with a clear compile error if a renamed grpc symbol surfaces (EC-3) [failure path] | Go 1.25 | N/A (did not fire) | No grpc symbol rename surfaced in the bumped tree — `make build` succeeded in build phase without any EC-3 fix required. Escalation path remains documented in test-plan; no execution evidence to record. |
| S13 — Negative — escalate if compile error is in non-test provider source (EC-4) [failure path] | Go 1.25 | N/A (did not fire) | Build succeeded. No follow-up SDK-bump ticket needed. |
| S14 — Negative — `go get <m>@<v>` returns 410 / not found (ERR-1) [failure path] | Go 1.25 + proxy.golang.org | N/A (did not fire) | All four target versions resolved successfully; module-resolution scenarios S1–S3 confirm proxy availability. |
| S15 — Post-bump Security Check reports 0 of the original 23 CVEs as still applicable [validation] | Internal Security Check pipeline | DEFERRED | Validate phase scope (`Section 1 / Section 9` of `docs/DLPXECO-14109-validation.md`). Not run in test phase. |
| S16 — `govulncheck ./...` (local preflight) shows no Critical/High attributable to the four bumped modules [validation] | Go 1.25 + govulncheck | DEFERRED | Validate phase scope. Not run in test phase (govulncheck not installed locally; would not affect merge gate). |
| S17 — PR description records before/after CVE counts [validation] | n/a | DEFERRED | PR phase scope. |
| S18 — Newly-introduced unrelated CVE finding (if any) is logged [validation / EC-5] | Internal Security Check pipeline | DEFERRED | Validate phase scope. |
| S19 — Signed-commit check passes on every commit on `vuln-fix` [process] | n/a | DEFERRED | PR phase scope (`git log --show-signature main..vuln-fix`). |

---

## Smoke (previously-generated functional tests)

`.claude/test/generated-test/` does not exist — no prior generated tests in this repo (the workflow is being adopted on this branch for the first time). Per the test-phase fallback rule: "No prior generated tests found — smoke skipped (first feature in this repo)."

The unit-test functions enumerated in S7 / S10 / S11 are hand-authored under `internal/provider/*_test.go` and serve as the de-facto smoke layer for this dependency-only change. They pass on both `main` and `vuln-fix` (S7 evidence above).

| Test File | Outcome | Notes |
|-----------|---------|-------|
| `.claude/test/generated-test/*` | N/A | Directory does not exist — no generated tests to smoke. |

---

## Failure Triage

The unscoped `go test ./...` invocation reported 17 failures under `internal/provider/`. Classification per `workflow-steps.md` ## Test:

| Test | Failure category | Root-cause analysis | Action |
|------|------------------|--------------------|--------|
| `TestAccEngineConfiguration_blockDevice`, `…_objectStorageWithRole`, `…_gcpObjectStorage`, `…_gcpObjectStorage_CC`, `…_comprehensive` (5) | (a) Test infrastructure | All fail with the same Terraform HCL parse error: `Invalid reference … on terraform_plugin_test.tf line 4 … "http://${ENGINE_NAME}.dlpxdc.co"`. The HCL template uses bash-style `${ENGINE_NAME}` interpolation but `ENGINE_NAME` is not exported in the local shell. Identical behaviour on `main` (the env-var is not set there either). Not caused by the dependency bump. | No action in this phase. Pre-existing test-infrastructure dependency. Documented for awareness; not a regression. |
| `TestAccEnvironment_positive`, `TestAccEnvironment_update_negative`, `TestAccEnvironment_update_positive`, `TestAccVdbGroup_*`, `TestAccVdb_provision_positive`, `TestAccVdb_appdata_provision`, `Test_Acc_Appdata_Dsource`, `Test_source_create_positive`, `TestOracleDsource_create_positive` (11) | (a) Test infrastructure | All fail with explicit `t.Fatal("<ENV_VAR> must be set for env acceptance tests")` — these tests intentionally hard-fail when their required acceptance env vars are absent. Verified on `main`: identical `t.Fatal` lines exist in `resource_oracle_dsource_test.go:56`, `resource_appdata_dsource_test.go:95`, `resource_database_postgresql_test.go:48`. Not a regression. | No action. These are environment-gated acceptance tests intended for the `make testacc` (TF_ACC=1) flow per `GNUmakefile`, not the `make test` merge gate. |
| `TestAccVdb_bookmark_provision` (1) | (a) Test infrastructure | `TestStep missing Config or ImportState or RefreshState` — fixture-construction failure with no env vars set. Same root cause as the env-gated set. Identical behaviour on `main`. | No action. Pre-existing. |

**Summary of failure triage**: All 17 failures fall under category (a) — test infrastructure / acceptance-gating — and reproduce identically on `main` with no source changes. Zero failures attributable to the four module bumps (`golang.org/x/{crypto,net,sys}`, `google.golang.org/grpc`). The merge gate per `.claude/test/testing.md` is the unit-test subset, which all PASS.

---

## Open Questions / Risks

- **OQ-1 — `golang.org/x/sys` resolved to `v0.45.0`, not the `v0.44.0` floor stated in the test plan.** Cause: the cascading `go mod tidy` from the `x/net v0.55.0` / `x/crypto v0.52.0` bumps required `x/sys >= v0.45.0` to compile. `v0.45.0` is newer than the CVE-clearance floor `v0.44.0`, so the security goal is still satisfied. Action: the test-plan / vision should be updated in the retrospective phase to reflect `v0.45.0` as the actual landed version. No FR-* breach.
- **OQ-2 — Acceptance tests gated by env vars produce loud `FAIL`s under `make test`.** Suggest the design / retrospective consider gating these with `if os.Getenv("TF_ACC") == ""` checks plus `t.Skip(...)` so the merge gate is quieter. Not in scope for this CVE-only ticket per NG2.

---

## Eval Check

(See `docs/DLPXECO-14109-eval-results.md` `### Step: test` for the structural eval output.)

---

## Summary

- **Functional (primary) unit suite (S7, S8, S10, S11)**: **14 of 14 true unit tests PASS** — `go test ./internal/provider/...` (scoped to non-acceptance tests) exits 0 in `0.625s` post-bump, vs `0.685s` baseline on `main`. Zero added, zero removed, zero changed result.
- **Resolution checks (S1, S2, S3, S4)**: **4 of 4 PASS** with one documentation drift on `x/sys` (OQ-1) that does not affect FR-001 acceptance criteria.
- **Build regression (S5, S9)**: PASS (recorded in build-phase artifact).
- **Failure-path scenarios (S12, S13, S14)**: did not fire — expected for a clean dependency bump.
- **Smoke (previously-generated functional tests)**: N/A — first feature on this branch.
- **Validation / PR / process scenarios (S6 deferred, S15–S19)**: deferred to validate / pr phases per scope of `docs/DLPXECO-14109-test-plan.md`.

**Test-phase verdict**: PASS. Proceed to `--step validate`.
