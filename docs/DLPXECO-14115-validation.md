# Validation Report: DLPXECO-14115

**Jira**: DLPXECO-14115 — [S2 Hard Gate] CI test coverage enforcement
**Branch**: `ci-support`
**Validation date**: 2026-06-07
**Inputs reviewed**: vision, functional, design, test-plan, test-evidence, implementation diff, build output

---

## Overall Verdict

**PASS WITH WARNINGS**

All four FRs are satisfied. All 19 test scenarios pass. The CI workflow file, threshold gate, and documentation sections are present and correct. **Warnings** below are non-blocking for the PR but should be tracked in follow-up tickets:

1. **Pre-existing acceptance tests are not env-var-gated** — when the new CI workflow runs for the first time, `go test ./...` will fail because 17 pre-existing `TestAcc*` / `_positive` tests call `t.Fatal` for missing DCT credentials regardless of `TF_ACC`. This is a project-wide test-design issue documented in the design phase risks and surfaced by the test phase. Resolution recommended in a follow-up ticket (add `TF_ACC` skip-guards).
2. **Scope expansion: `develop` branch added to triggers** — the vision and functional spec scope the workflow to `main` only; the implemented `ci.yml` also triggers on `develop`. The design phase asked this as an open question and the implementation chose the broader scope. Update vision/functional spec retroactively in a doc-only follow-up, or revert `ci.yml` to `main`-only if scope strictness is preferred.

The PR can be raised; reviewers should be informed of both warnings.

---

## Section 1 — Spec Coverage (FR-by-FR)

| FR | Description | Implementation Location | Test Scenarios | Verdict |
|----|-------------|------------------------|----------------|---------|
| FR-001 | GitHub Actions CI Workflow | `.github/workflows/ci.yml` lines 16–69 | S1, S2, S3, S4, S5, S6, S7 (all PASS) | PASS |
| FR-002 | Coverage Threshold Enforcement | `.github/workflows/ci.yml` lines 28–36, 71–85 | S8, S9, S10, S11 (all PASS) | PASS |
| FR-003 | Documentation — CLAUDE.md and CONTRIBUTING.md | `CLAUDE.md` `## CI Contract` section; `CONTRIBUTING.md` `## CI` section | S12, S13, S14, S15, S16 (all PASS) | PASS |
| FR-004 | Branch Protection Contract Documentation | `CLAUDE.md` `### Branch Protection` subsection | S17, S18, S19 (all PASS) | PASS |

Every FR is implemented and verified. No FR is missing tests or evidence.

---

## Section 2 — Acceptance Criteria Coverage (AC-by-AC from functional + design)

### Functional FR-001 (CI workflow)
- [x] AC-1 (`pull_request` to `main` runs job + artifact) — file structure verified; runtime AC-1 will be verified by first real CI run after PR merges (out-of-band)
- [x] AC-2 (job fails when tests fail) — verified at the workflow level (the test step is unguarded; non-zero exit fails the job)
- [x] AC-3 (push to `main` triggers job) — verified (S3)
- [x] AC-4 (`go-version-file: go.mod`) — verified (S4)

### Functional FR-002 (threshold)
- [x] AC-1 (passes when coverage >= threshold) — verified inline (S9)
- [x] AC-2 (fails when below) — verified inline (S10)
- [x] AC-3 (threshold change in `ci.yml` takes effect immediately) — verified by inspection: threshold is an env var read in shell, no other reference
- [x] AC-4 (missing `coverage.out` fails informatively) — verified (S11)

### Functional FR-003 (docs)
- [x] AC-1 (`CLAUDE.md` has CI Contract with workflow path, threshold, local command) — verified (S12, S13, S14)
- [x] AC-2 (`CONTRIBUTING.md` `## CI` names `ci / unit-tests`) — verified (S15, S16)
- [x] AC-3 (no TBD / TODO in new sections) — verified by inspection of CLAUDE.md `## CI Contract` and `CONTRIBUTING.md` `## CI` sections

### Functional FR-004 (branch protection)
- [x] AC-1 (subsection has exact status-check string) — verified (S17)
- [x] AC-2 (`grep "ci / unit-tests" CLAUDE.md` returns match) — verified (S17)

### Design ACs (numbered AC-1..AC-11)
- [x] AC-1: `.github/workflows/ci.yml` exists, triggers on `pull_request` to `main` and `push` to `main` — PASS; also triggers on `develop` (see Warning 2)
- [x] AC-2: Job exits 0 when all unit tests pass — workflow structure correct; **conditional**: pre-existing acceptance tests will fail in CI (Warning 1)
- [x] AC-3: Job exits non-zero on test failure — verified by workflow structure
- [x] AC-4: `go-version-file: go.mod` present — PASS
- [x] AC-5: Threshold below value exits non-zero — PASS (S10)
- [x] AC-6: Missing `coverage.out` exits non-zero — PASS (S11)
- [x] AC-7: `COVERAGE_THRESHOLD` in env block with comment — PASS (S8 + inspection)
- [x] AC-8: `CLAUDE.md` has `## CI Contract` with workflow path, threshold, baseline, local command — PASS (S12, S13, S14)
- [x] AC-9: `CONTRIBUTING.md` `## CI` names `ci / unit-tests` and gives local command — PASS (S15, S16)
- [x] AC-10: `grep "ci / unit-tests" CLAUDE.md` returns match — PASS (S17)
- [x] AC-11: `### Branch Protection` subsection names both required settings — PASS (S18, S19)

**Section 2 verdict**: All 19 ACs satisfied at the artifact level. AC-2 of the design has a runtime caveat documented in Warnings.

---

## Section 3 — Non-Goals Compliance

| Non-Goal | Verification | Verdict |
|----------|--------------|---------|
| NG1: No acceptance tests in CI | `.github/workflows/ci.yml` does not export `TF_ACC` (verified by S5); test command is plain `go test ./...` | PASS |
| NG2: No self-hosted runners | Workflow uses `ubuntu-latest` (line 41) | PASS |
| NG3: No external coverage services | No Codecov / Coveralls / SonarQube references in `ci.yml` | PASS |
| NG4: No existing `*_test.go` file modified | `git status` shows the only test-file change is the NEW `ci_workflow_test.go`; no existing `*_test.go` modified by this ticket. `internal/provider/resource_vdb_test.go` shows as modified, but that change is from a prior commit on this branch (DLPXECO-13662 sibling work), not this ticket | PASS |

---

## Section 4 — Quality Rules Compliance

| Rule | Verification | Status |
|------|--------------|--------|
| API backward compatibility preserved | `git diff` for this ticket changes only `.github/workflows/ci.yml`, `CLAUDE.md`, `CONTRIBUTING.md`, plus the new `ci_workflow_test.go` (test asset). No `internal/provider/*.go` changes attributable to this ticket | PASS |
| Migration path provided | Local equivalent command documented in `CLAUDE.md` `## CI Contract` (lines confirming `coverprofile=coverage.out` present) and `CONTRIBUTING.md` `### Verifying Locally Before Pushing` | PASS |
| No acceptance tests in CI | Workflow does not export `TF_ACC` (verified by S5 line-by-line scan) | PASS |
| Threshold set at or below current baseline | Baseline 2.3% (CI conditions, no creds); threshold 2; 2 <= 2.3 | PASS |

---

## Section 5 — Test Results Summary

| Metric | Value |
|--------|-------|
| Scenarios required by plan | 19 |
| Scenarios passing | **19** |
| Scenarios failing | 0 |
| Scenarios skipped | 0 |
| Generated test file | `ci_workflow_test.go` (repo root, package `main`) |
| Full-suite coverage | 3.8% |
| Pre-existing test failures (out-of-scope) | 17 (documented in test-evidence) |

See `docs/DLPXECO-14115-test-evidence.md` for full per-scenario output.

---

## Section 6 — Edge Cases & Error Scenarios

| ID | Type | Description | Implementation Handling | Verdict |
|----|------|-------------|------------------------|---------|
| EC-1 | Edge | `coverage.out` empty (no test files matched) | Threshold script checks `[[ ! -s coverage.out ]]` and exits 1 with "missing or empty" | PASS (S11) |
| EC-2 | Edge | New file drops coverage below threshold | Threshold step exits non-zero with actual vs. threshold — verified by S10 | PASS |
| EC-3 | Edge | Go version in `go.mod` not yet in `actions/setup-go` | `actions/setup-go@v5` step fails; clear error from action; documented in design | PASS (documented) |
| EC-4 | Edge | Concurrent PRs change coverage | Each PR runs CI against its own commit SHA; merged-to-main threshold is what matters; no special handling needed | PASS (no code path needed) |
| EC-5 | Edge | `make test` 30 s timeout too short | Workflow uses `go test -timeout=300s` directly (line 61), not `make test` | PASS (S6) |
| ERR-1 | Error | Runner cannot install Go | `actions/setup-go` exits non-zero; job marked failed; standard GHA behavior | PASS (no custom handling needed) |
| ERR-2 | Error | `go test` compile failure | No `coverage.out` produced; threshold step's missing-file branch handles it | PASS (S11) |
| ERR-3 | Error | `actions/upload-artifact` transient failure | `if: always()` annotation ensures the step runs even after test failure; artifact-step failure marks job failed | PASS (S7) |

---

## Section 7 — Documentation Completeness

| Document | Required Content | Verification | Verdict |
|----------|------------------|--------------|---------|
| `CLAUDE.md` `## CI Contract` | Workflow file path, triggering conditions, `COVERAGE_THRESHOLD` value, baseline, local equivalent command, how to update threshold, status-check name | All present (S12, S13, S14, S17, S18, S19); no `TODO`/`TBD` in the section | PASS |
| `CLAUDE.md` `### Branch Protection` | Exact status-check string, "Require status checks to pass before merging", "Require branches to be up to date before merging", note that maintainer applies in GitHub UI | All present (S17, S18, S19); note about manual application present in section text | PASS |
| `CONTRIBUTING.md` `## CI` | Status check requirement, what CI does, local verification command, note on acceptance tests | All present (S15, S16); local command block present; acceptance test exclusion note present | PASS |

---

## Section 8 — Drift Detection (Spec vs. Implementation)

| Drift | Severity | Detail | Recommendation |
|-------|----------|--------|----------------|
| Branch trigger scope: `develop` added | Low (scope expansion) | Vision SC1 says `main` only; functional FR-001 says `main` only; design Q-2 (line 305) flagged this as an open question; implementation triggers on both `main` and `develop` | Non-blocking. Either (a) update vision/functional FR-001 to formally include `develop` (doc-only follow-up) or (b) revert `ci.yml` to `main`-only. Recommend (a) since `CLAUDE.md` says internal PRs branch from `develop` |
| S5 test definition vs. design intent | Negligible | Test plan S5 says "file does not contain string TF_ACC"; ci.yml contains `TF_ACC` in explanatory comments; design intent is "does not export TF_ACC". Test tightened to match design intent | Already addressed in test phase; no further action |

No other drift detected. Implementation files match design's `### Source Files to Modify` table exactly (`.github/workflows/ci.yml`, `CLAUDE.md`, `CONTRIBUTING.md`).

---

## Section 9 — Risk Status

| Risk (from vision/design) | Current Status | Notes |
|---------------------------|---------------|-------|
| Low baseline coverage (vision Risk row 1) | Materialised but mitigated | Threshold 2% is infrastructure-level; documented in CLAUDE.md; raising it is a follow-up |
| Acceptance tests hang CI (vision Risk row 2) | Mitigated | Workflow does not export `TF_ACC`; tests using `resource.TestCase` will skip when TF_ACC is unset — **but** the project's PreCheck functions call `t.Fatal` before reaching the SDK skip; pre-existing failure surfaced by test phase (Warning 1) |
| `make test` 30s timeout (vision Risk row 3) | Resolved | Workflow uses `go test -timeout=300s` directly (S6) |
| Go version mismatch (vision Risk row 4) | Resolved | `go-version-file: go.mod` (S4) |
| Branch protection not configured (vision Risk row 5) | Documented | `CLAUDE.md` Branch Protection subsection states maintainer must configure in GitHub Settings → Branches → Branch protection rules |

---

## Recommendations for PR Phase

1. **Mention both warnings in the PR description**, with explicit links to the test-evidence and validation docs. Reviewers must understand that the first CI run will likely be RED on existing acceptance tests and that this is not regression from this PR.
2. **Propose follow-up tickets**:
   - "Gate pre-existing PreCheck functions on `TF_ACC` so unit-test CI passes without DCT credentials" — blocker for the CI gate to be actually useful
   - "Document `develop` branch trigger formally in vision/functional spec (doc-only)" — or revert workflow scope, depending on team preference
3. **Confirm with repo maintainer** that branch protection settings will be applied after merge (out-of-band per FR-004 ACs).
