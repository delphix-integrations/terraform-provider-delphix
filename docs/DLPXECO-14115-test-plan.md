# Test Plan: DLPXECO-14115

**Jira**: DLPXECO-14115 — [S2 Hard Gate] CI test coverage enforcement
**Derived from**: `docs/DLPXECO-14115-design.md` `## Affected Components` and `## Version Compatibility`

---

## Test Approach

This feature delivers CI infrastructure (a YAML workflow file) and documentation updates rather than new Go library code. The appropriate test approach is:
1. **Static validation**: verify the workflow YAML contains the required keys/values (`grep`).
2. **Shell-script logic unit test**: the coverage-threshold enforcement shell snippet is exercised locally with mock `coverage.out` fixtures.
3. **Documentation completeness checks**: verify required strings are present in `CLAUDE.md` and `CONTRIBUTING.md` using `grep`.

Tests are implemented as Go test functions using the standard `testing` package (the project's test framework). No live DCT infrastructure is required.

## Environment / Landscape

- Landscape: local developer environment
- Service under test: none — tests validate file content and shell script logic
- No VMs or external services required

## Versions to Cover

| Version | Why | Required? |
|---------|-----|-----------|
| Go 1.25 (current, from `go.mod`) | Only Go version in use; workflow targets this | Yes |

## Scenarios

| # | Scenario | Maps to FR | Versions | Expected outcome |
|---|----------|-----------|----------|------------------|
| S1 | `.github/workflows/ci.yml` exists and contains `name: ci` at top level | FR-001 | Go 1.25 | File exists; file contains `name: ci` |
| S2 | Workflow triggers on `pull_request` to `main` | FR-001 | Go 1.25 | File contains `pull_request:` and `main` in branches |
| S3 | Workflow triggers on `push` to `main` | FR-001 | Go 1.25 | File contains `push:` and `main` in branches |
| S4 | Workflow uses `go-version-file: go.mod` (not a hardcoded version) | FR-001 | Go 1.25 | File contains `go-version-file: go.mod` |
| S5 | Workflow does NOT export `TF_ACC` (acceptance test exclusion) | FR-001 | Go 1.25 | File does NOT contain the string `TF_ACC` |
| S6 | Workflow test command uses `-timeout=300s` | FR-001 | Go 1.25 | File contains `-timeout=300s` |
| S7 | `upload-artifact` step has `if: always()` | FR-001 | Go 1.25 | File contains `if: always()` |
| S8 | `COVERAGE_THRESHOLD` is defined in `ci.yml` env block | FR-002 | Go 1.25 | File contains `COVERAGE_THRESHOLD:` |
| S9 | Threshold script exits 0 when coverage exceeds threshold (mock: 55%, threshold: 50) | FR-002 | Go 1.25 | Shell snippet runs; exit code 0; output contains "PASS" |
| S10 | Threshold script exits non-zero when coverage is below threshold (mock: 48%, threshold: 50) | FR-002 | Go 1.25 | Shell snippet runs; exit code 1; output contains "FAIL" and both percentages |
| S11 | Threshold script exits non-zero when `coverage.out` is missing | FR-002 | Go 1.25 | Shell snippet runs without coverage.out; exit code 1; output contains "missing or empty" |
| S12 | `CLAUDE.md` contains `## CI Contract` section | FR-003 | N/A | `grep "## CI Contract" CLAUDE.md` returns match |
| S13 | `CLAUDE.md` CI Contract section contains the local equivalent command | FR-003 | N/A | `CLAUDE.md` contains `coverprofile=coverage.out` |
| S14 | `CLAUDE.md` CI Contract section documents `COVERAGE_THRESHOLD` | FR-003 | N/A | `CLAUDE.md` contains `COVERAGE_THRESHOLD` |
| S15 | `CONTRIBUTING.md` contains `## CI` section | FR-003 | N/A | File contains `## CI` heading |
| S16 | `CONTRIBUTING.md` CI section names the status check | FR-003 | N/A | `CONTRIBUTING.md` contains `ci / unit-tests` |
| S17 | `CLAUDE.md` contains the exact status-check string `ci / unit-tests` | FR-004 | N/A | `grep "ci / unit-tests" CLAUDE.md` returns match |
| S18 | `CLAUDE.md` Branch Protection names "Require status checks to pass before merging" | FR-004 | N/A | `CLAUDE.md` contains that exact phrase |
| S19 | `CLAUDE.md` Branch Protection names "Require branches to be up to date before merging" | FR-004 | N/A | `CLAUDE.md` contains that exact phrase |

## Out of Scope

- Live GitHub Actions run verification — tested by the first real PR after merge; not locally testable without a GitHub account/forked repo.
- Acceptance tests (`TF_ACC=1`) — per Non-Goal NG1; require live DCT infrastructure.
- External coverage reporting services (Codecov, Coveralls) — per Non-Goal NG3.
- Modifying existing `*_test.go` files — per Non-Goal NG4.

## Test Data Requirements

- Mock `coverage.out` fixtures (three: 55% above threshold, 48% below threshold, missing file) — created inline in test functions.
- `CLAUDE.md` and `CONTRIBUTING.md` must exist with the new sections after the implement phase runs.
- `.github/workflows/ci.yml` must exist after the implement phase.

## Exit Criteria

- All Required scenarios (S1–S19) PASS
- Smoke suite (`*.test.*` files in `.claude/test/generated-test/` excluding DLPXECO-14115) PASSES or "first feature — smoke skipped" note
- No scenario marked SKIPPED without a documented reason

---
<!-- Cross-references:
     - FR-001 → S1–S7; FR-002 → S8–S11; FR-003 → S12–S16; FR-004 → S17–S19
     - All 4 FRs in docs/DLPXECO-14115-functional.md have at least one scenario here
     - Versions column is a subset of docs/DLPXECO-14115-design.md ## Version Compatibility -->
