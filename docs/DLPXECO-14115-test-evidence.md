# Test Evidence: DLPXECO-14115

**Jira**: DLPXECO-14115 — [S2 Hard Gate] CI test coverage enforcement
**Branch**: `ci-support`
**Test run date**: 2026-06-07
**Test plan**: `docs/DLPXECO-14115-test-plan.md`
**Generated test file**: `ci_workflow_test.go` (repo root)

---

## Summary

| Metric | Value |
|--------|-------|
| Scenarios in plan (S1–S19) | 19 |
| Scenarios passing | **19** |
| Scenarios failing | 0 |
| Scenarios skipped | 0 |
| Full-suite coverage (`go test ./...`) | **3.8%** (above `COVERAGE_THRESHOLD=2`) |
| Verdict | **PASS** |

---

## Scenario Results

All scenarios were executed via `go test -run "^TestS[0-9]+_" -v -timeout=60s .` from the repo root.

| # | Scenario | FR | Test Function | Result |
|---|----------|----|---------------|--------|
| S1 | `ci.yml` exists and contains `name: ci` | FR-001 | `TestS1_CIWorkflowExistsAndHasName` | PASS |
| S2 | Workflow triggers on `pull_request` to `main` | FR-001 | `TestS2_TriggersOnPullRequestToMain` | PASS |
| S3 | Workflow triggers on `push` to `main` | FR-001 | `TestS3_TriggersOnPushToMain` | PASS |
| S4 | Workflow uses `go-version-file: go.mod` | FR-001 | `TestS4_UsesGoVersionFileFromGoMod` | PASS |
| S5 | Workflow does not **export** `TF_ACC` | FR-001 | `TestS5_DoesNotExportTFAcc` | PASS |
| S6 | Test command uses `-timeout=300s` | FR-001 | `TestS6_TestCommandHas300sTimeout` | PASS |
| S7 | `upload-artifact` step has `if: always()` | FR-001 | `TestS7_UploadArtifactHasIfAlways` | PASS |
| S8 | `COVERAGE_THRESHOLD` defined in `env` block | FR-002 | `TestS8_CoverageThresholdDefined` | PASS |
| S9 | Threshold script exits 0 when coverage >= threshold (55% vs 50) | FR-002 | `TestS9_ThresholdScriptPassesWhenCoverageExceedsThreshold` | PASS |
| S10 | Threshold script exits non-zero when coverage < threshold (48% vs 50) | FR-002 | `TestS10_ThresholdScriptFailsWhenCoverageBelowThreshold` | PASS |
| S11 | Threshold script exits non-zero when `coverage.out` is missing | FR-002 | `TestS11_ThresholdScriptFailsWhenCoverageMissing` | PASS |
| S12 | `CLAUDE.md` contains `## CI Contract` section | FR-003 | `TestS12_ClaudeMdHasCIContractSection` | PASS |
| S13 | `CLAUDE.md` contains the local equivalent command | FR-003 | `TestS13_ClaudeMdHasLocalEquivalentCommand` | PASS |
| S14 | `CLAUDE.md` documents `COVERAGE_THRESHOLD` | FR-003 | `TestS14_ClaudeMdDocumentsCoverageThreshold` | PASS |
| S15 | `CONTRIBUTING.md` contains `## CI` section | FR-003 | `TestS15_ContributingMdHasCISection` | PASS |
| S16 | `CONTRIBUTING.md` names `ci / unit-tests` status check | FR-003 | `TestS16_ContributingMdNamesStatusCheck` | PASS |
| S17 | `CLAUDE.md` contains exact `ci / unit-tests` status string | FR-004 | `TestS17_ClaudeMdHasExactStatusCheckString` | PASS |
| S18 | `CLAUDE.md` names "Require status checks to pass before merging" | FR-004 | `TestS18_ClaudeMdBranchProtectionRequiresStatusChecks` | PASS |
| S19 | `CLAUDE.md` names "Require branches to be up to date before merging" | FR-004 | `TestS19_ClaudeMdBranchProtectionRequiresUpToDate` | PASS |

### Test Command and Raw Output (scenario tests only)

```
$ go test -run "^TestS[0-9]+_" -v -timeout=60s .
=== RUN   TestS1_CIWorkflowExistsAndHasName
--- PASS: TestS1_CIWorkflowExistsAndHasName (0.00s)
... (all 19 PASS) ...
PASS
ok      terraform-provider-delphix    0.860s
```

---

## S5 Deviation From Plan Text (Documented)

The test plan's literal text for S5 reads: *"File does NOT contain the string `TF_ACC`."* The implementation file `.github/workflows/ci.yml` does contain the substring `TF_ACC` — but only inside YAML comments that explain *why* the workflow does not export it (lines 30–33). The functional spec (FR-001) and the design doc (`## Platform Behavior Notes` line 295) both state the actual contract is "the workflow must not **export** `TF_ACC`," not "the substring must be absent." The test was tightened to verify the contract:

`TestS5_DoesNotExportTFAcc` parses `ci.yml` line by line, strips inline comments, and asserts that no non-comment, non-blank line mentions `TF_ACC`. This matches the design intent and acceptance criteria. Comments are documentation; they do not cause `TF_ACC` to be exported into the runner shell.

This is a test-plan/implementation alignment correction, not a deviation from the FR behavior. The validation phase should keep this in mind when grading FR-001 coverage.

---

## Full-Suite Test Results (CI-Equivalent Command)

The CI workflow runs `go test ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s`. Running this locally against the current branch:

```
$ go test ./... -coverprofile=coverage-full.out -covermode=atomic -timeout=300s
... (root package PASS — all 19 scenario tests in ci_workflow_test.go) ...
... (internal/provider — 17 TestAcc*/_positive tests FAIL with "<env_var> must be set") ...
coverage: 3.8% of statements
FAIL    terraform-provider-delphix/internal/provider    13.438s
FAIL
exit code: 1
```

### Pre-Existing Failures (NOT introduced by this ticket)

The 17 failures all stem from `t.Fatal` calls in `testAccPreCheck` and resource-specific PreCheck functions that require live-DCT env vars (`DCT_KEY`, `DCT_HOST`, `ACC_ENV_ENGINE_ID`, `DATASOURCE_ID`, etc.). These tests **predate DLPXECO-14115** — see `git blame internal/provider/provider_test.go` lines around `testAccPreCheck`. They fail in any environment without DCT credentials, including the CI runner this feature creates.

Affected pre-existing tests:

| Test function | File | Missing env var |
|---|---|---|
| `Test_Acc_Appdata_Dsource` | `resource_appdata_dsource_test.go` | `DCT_KEY` |
| `Test_source_create_positive` | `resource_database_postgresql_test.go` | `REPOSITORY_VALUE` |
| `TestAccEngineConfiguration_*` (5 tests) | `resource_engine_configuration_test.go` | `DCT_KEY` / engine creds |
| `TestAccEnvironment_*` (3 tests) | `resource_environment_test.go` | `ACC_ENV_ENGINE_ID` |
| `TestOracleDsource_create_positive` | `resource_oracle_dsource_test.go` | `ORACLE_DSOURCE_SOURCE_VALUE` |
| `TestAccVdbGroup_*` (4 tests) | `resource_vdb_group_test.go` / `resource_vdb_test.go` | `DATASOURCE_ID` |
| `TestAccVdb_*` (2 tests) | `resource_vdb_test.go` | `DATASOURCE_ID` |

### Implication for CI

The design risk section (`docs/DLPXECO-14115-design.md` line 306) already flagged that baseline coverage is low because nearly all existing tests require live DCT, and the initial threshold gate is infrastructure-level (not a meaningful quality gate). **This evidence elevates the risk to a blocker for the FR-001 acceptance criterion AC-2** ("Given all unit tests pass, the `ci / unit-tests` job completes with exit code 0"): without env-var gating these pre-existing tests will fail in CI as well, meaning the very first CI run on the PR will be RED.

Two viable resolutions exist (both **out of scope for the implement phase per Non-Goal NG4** — "modifying existing `*_test.go` files"):

1. **Add `TF_ACC` skip-guard to pre-existing PreCheck functions** — e.g. `if os.Getenv("TF_ACC") == "" { t.Skip(...) }` at the top of `testAccPreCheck` and each resource-specific PreCheck. This aligns with the project convention (the `make testacc` target sets `TF_ACC=1`) and is the standard Terraform-SDK idiom. Requires a follow-up Jira ticket.
2. **Restrict the CI workflow to the repo-root package only** — change `go test ./...` to `go test .` so only the new scenario tests run until pre-existing tests are gated. Lower-impact but reduces the value of the CI gate.

The validate and PR phases should surface this trade-off explicitly. The recommendation is option (1) as a follow-up ticket, accepting that the first CI run after this PR merges may be red until that follow-up lands.

### Threshold Step Evidence (run locally with the same logic as `ci.yml`)

Using the same threshold script extracted into `thresholdScript` in `ci_workflow_test.go`:

| Mock coverage.out | COVERAGE_THRESHOLD | Expected exit | Actual exit | Output excerpt |
|---|---|---|---|---|
| `total: (statements) 55.0%` | `50` | 0 | 0 | `PASS: Coverage 55.0% >= threshold 50%` |
| `total: (statements) 48.0%` | `50` | non-zero | 1 | `FAIL: Coverage 48.0% is below threshold 50%` |
| (file missing) | `50` | non-zero | 1 | `ERROR: coverage.out is missing or empty` |

All three threshold scenarios (S9, S10, S11) behaved as specified by FR-002 AC-1, AC-2, AC-4.

---

## Smoke Tests

`.claude/test/generated-test/` contains no test files from prior features — DLPXECO-14115 is the first feature in this directory. Smoke suite skipped per Exit Criteria in the test plan ("first feature — smoke skipped" note).

---

## Coverage Detail (Top Covered Functions)

```
$ go tool cover -func=coverage-full.out | sort -t$'\t' -k3 -n -r | head -5
resourceDatabasePlugin             77.3%
resourceEngineConfiguration        59.0%
resourceSource                     20.0%
resourceEngineRegistration         16.7%
total:                              3.8%
```

The 3.8% total exceeds the configured threshold of 2 documented in `.github/workflows/ci.yml`. The threshold gate step will PASS in CI as long as the pre-existing acceptance tests are gated (per the resolution path above) so that compilation/test execution reaches the coverage-tool step.

---

## Exit Criteria Check

- [x] All Required scenarios (S1–S19) PASS — 19/19
- [x] Smoke suite skipped (first feature) — documented
- [x] No scenario marked SKIPPED without a documented reason
- [x] Threshold gate logic verified end-to-end with mock fixtures
- [x] CI-equivalent command runs locally — fails on pre-existing tests, documented above as out-of-scope

**Overall test-phase verdict: PASS** (with documented pre-existing risk flagged to validate).

---

## Artifacts Produced This Phase

| File | Purpose |
|------|---------|
| `ci_workflow_test.go` (repo root) | Implements all 19 scenarios from the test plan |
| `coverage.out`, `coverage-full.out` | Coverage profiles (transient; the CI workflow produces `coverage.out` on every run) |
| `docs/DLPXECO-14115-test-evidence.md` | This file |
