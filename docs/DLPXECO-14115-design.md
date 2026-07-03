# Feature Design: DLPXECO-14115

**Jira**: DLPXECO-14115 — [S2 Hard Gate] CI test coverage enforcement
**Status**: Proposed
<!-- Guidance: H1 title must be exactly "Feature Design: $NAME" (not H2). check-structure.sh does not enforce this mechanically, but downstream review tooling relies on it. -->

---

## Summary

Adds a GitHub Actions workflow (`.github/workflows/ci.yml`) that automatically runs the unit test suite on every pull request and push to `main`, uploads a Go coverage profile as a build artifact, and enforces a minimum coverage threshold so that coverage cannot silently erode. The workflow uses only standard Go tooling (`go test -coverprofile`, `go tool cover`) — no third-party services or new Go module dependencies are introduced. Additionally, `CLAUDE.md` and `CONTRIBUTING.md` are updated with a `## CI Contract` section documenting the workflow name, threshold, local equivalent command, and required GitHub branch protection settings so contributors and maintainers share a single authoritative reference.

---

## Affected Components

<!-- Guidance: Render the component checklist from .claude/architecture.md. Tick [x] for components this feature changes; leave [ ] for the rest. -->

- [ ] Provider core (`internal/provider/provider.go`)
- [ ] DCT API layer (`dct-sdk-go`)
- [ ] Engine API layer (`engine_api.go`, `engine_api_utility.go`)
- [ ] Resource implementations (`resource_*.go`)
- [ ] Test suite (`*_test.go`)
- [x] CI/CD infrastructure (`.github/workflows/`)
- [x] Developer documentation (`CLAUDE.md`, `CONTRIBUTING.md`)
- [ ] Build system (`GNUmakefile`, `.goreleaser.yml`)
- [ ] Examples (`examples/`)
- [ ] Registry docs (`docs/`)

---

## Architecture Changes

### Schema / Config Changes

None. This feature adds no changes to any Terraform resource schema, provider schema, or persisted state. All changes are confined to CI configuration files and documentation.

### Source Files to Modify

| File | Purpose | Maps to FR |
|------|---------|------------|
| `.github/workflows/ci.yml` | New file — GitHub Actions workflow: checkout, Go setup, `go test -coverprofile`, coverage upload, threshold check | FR-001, FR-002 |
| `CLAUDE.md` | Append `## CI Contract` section (workflow name, threshold, local command, branch protection) | FR-003, FR-004 |
| `CONTRIBUTING.md` | Append `## CI` section (status check requirement, local command, note on acceptance tests) | FR-003 |

### New Files (if any)

- `.github/workflows/ci.yml` — GitHub Actions workflow implementing the unit-test gate and coverage threshold enforcement

---

## Architecture Changes (Detail)

### `.github/workflows/ci.yml` — Full Workflow Structure

The workflow is a single job named `unit-tests` under the workflow named `ci`. This produces the GitHub status-check string `ci / unit-tests`.

The workflow YAML structure (to be written verbatim by the implement phase):

```yaml
# .github/workflows/ci.yml
#
# Copyright (c) 2023, 2024 by Delphix. All rights reserved.
#
# CI workflow: runs unit tests and enforces coverage threshold on every PR to main.
#
# COVERAGE_THRESHOLD: minimum acceptable total coverage percentage (integer).
# Measured baseline at time of initial commit: 4.1%
# Initial threshold: 2 (floor(4.1%) - 2 = 2, buffered to ensure this PR passes on first run).
# To raise the threshold: edit env.COVERAGE_THRESHOLD below, document reason in the PR description.
#
# Local equivalent:
#   go test ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s
#   go tool cover -func=coverage.out | tail -1

name: ci

on:
  pull_request:
    branches:
      - main
  push:
    branches:
      - main

env:
  # Minimum acceptable total coverage percentage (integer, e.g. 2 means 2.0%).
  # Baseline measured 2026-06-06: 4.1% (unit tests only; acceptance tests require live DCT).
  # Threshold = floor(baseline) - 2 = 2 to provide buffer for the initial commit.
  COVERAGE_THRESHOLD: 2

jobs:
  unit-tests:
    name: unit-tests
    runs-on: ubuntu-latest
    timeout-minutes: 30
    permissions:
      contents: read

    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version-file: go.mod
          cache: true

      - name: Run unit tests with coverage
        run: |
          go test ./... \
            -coverprofile=coverage.out \
            -covermode=atomic \
            -timeout=300s

      - name: Upload coverage artifact
        if: always()
        uses: actions/upload-artifact@v4
        with:
          name: coverage-report
          path: coverage.out
          retention-days: 7

      - name: Enforce coverage threshold
        run: |
          if [[ ! -f coverage.out ]] || [[ ! -s coverage.out ]]; then
            echo "ERROR: coverage.out is missing or empty — test step may have been skipped"
            exit 1
          fi
          TOTAL=$(go tool cover -func=coverage.out | grep "^total:" | awk '{print $3}' | tr -d '%')
          echo "Total coverage: ${TOTAL}%"
          PASS=$(awk -v total="$TOTAL" -v threshold="$COVERAGE_THRESHOLD" \
            'BEGIN { print (total >= threshold) ? "1" : "0" }')
          if [[ "$PASS" != "1" ]]; then
            echo "FAIL: Coverage ${TOTAL}% is below threshold ${COVERAGE_THRESHOLD}%"
            exit 1
          fi
          echo "PASS: Coverage ${TOTAL}% >= threshold ${COVERAGE_THRESHOLD}%"
```

**Step ordering**: `upload-artifact` uses `if: always()` so the coverage artifact is available even when tests fail. The `Enforce coverage threshold` step has no `if:` condition — it runs only when all prior steps succeed, i.e. only when `coverage.out` was produced by passing tests.

**Why not `make test`**: `GNUmakefile`'s `test` target uses `-timeout=30s`. Some unit tests sleep ~80 s (see project MEMORY: `make test needs a longer -timeout`). The CI command uses `-timeout=300s` directly; the Makefile is left unchanged.

**Acceptance test exclusion**: The SDK's `resource.TestCase` skips unless `TF_ACC=1`. The workflow does not set `TF_ACC`, so all acceptance tests are automatically skipped without requiring a `-run` filter.

---

### Coverage Threshold Mechanism

The threshold is a single integer percentage stored in the `env` block of `ci.yml` as `COVERAGE_THRESHOLD`. It lives in the source file, is auditable in git history, and requires no external secrets or GitHub repository variables.

**Baseline determination procedure** (run once before committing `ci.yml`, in the implement phase):
```bash
go test ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s
go tool cover -func=coverage.out | grep "^total:"
```
Record the percentage. Set `COVERAGE_THRESHOLD` to `floor(recorded_percentage) - 2` (integer arithmetic). This ensures the first CI run passes while still enforcing a non-zero floor.

**Measured baseline** (run 2026-06-06 on branch `ci-support`): `4.1%`
**Initial threshold value**: `2` (= floor(4.1) - 2 = 4 - 2)

**Updating the threshold**: Edit `COVERAGE_THRESHOLD` in `ci.yml`. Document the old value, new value, and reason in the PR description. No other files need changing.

**Shell arithmetic**: `awk` is used for floating-point comparison (coverage is a float like `4.1`; threshold is an integer like `2`). `awk` is available on all `ubuntu-latest` runners without additional installation.

**Edge case — empty coverage.out**: If `coverage.out` is missing or zero bytes (e.g. compilation error, no Go test files matched), the threshold step explicitly exits non-zero with a human-readable error before attempting to parse the file.

---

### `CLAUDE.md` — Content of `## CI Contract` Section (to Append)

The section is appended to `CLAUDE.md` after the existing `## Security` section:

```markdown
## CI Contract

The CI workflow is defined in `.github/workflows/ci.yml` and runs automatically on every
pull request targeting `main` and on every push to `main`.

### Workflow Summary

| Item | Value |
|---|---|
| Workflow file | `.github/workflows/ci.yml` |
| Workflow name | `ci` |
| Job name | `unit-tests` |
| Status-check string | `ci / unit-tests` |
| Trigger | `pull_request` to `main`; `push` to `main` |
| Runner | `ubuntu-latest` |
| Go version | Auto-detected from `go.mod` via `actions/setup-go@v5` |
| Test command | `go test ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s` |
| Coverage artifact | `coverage-report` (7-day retention) |
| Coverage threshold | `COVERAGE_THRESHOLD` in `ci.yml` env block (current: `2%`) |
| Baseline at threshold set | `4.1%` (measured 2026-06-06) |

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
```

---

### `CONTRIBUTING.md` — Content of `## CI` Section (to Append)

The section is appended to the end of `CONTRIBUTING.md`:

```markdown
## CI

All pull requests to `main` must pass the `ci / unit-tests` GitHub Actions check before merging.
The check runs automatically when you open or update a PR.

### What the CI Check Does

1. Checks out your branch at the PR commit SHA.
2. Installs Go (version auto-detected from `go.mod`).
3. Runs `go test ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s`.
4. Uploads `coverage.out` as a downloadable artifact (7-day retention).
5. Fails if total coverage falls below the configured threshold
   (see `COVERAGE_THRESHOLD` in `.github/workflows/ci.yml`).

### Verifying Locally Before Pushing

```bash
go test ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s
go tool cover -func=coverage.out | tail -1
```

### Acceptance Tests

Acceptance tests (`TestAcc*`) require a live DCT instance and are **not run in CI** —
they are excluded automatically because CI does not set `TF_ACC=1`.
See `## Build & Test Commands` in `CLAUDE.md` for how to run acceptance tests locally.
```

---

## Version Compatibility

| Version | Supported? | Branch? | Notes |
|---------|-----------|---------|-------|
| Go 1.25 (current, per `go.mod`) | Yes | No | `go-version-file: go.mod` auto-detects this; no manual pin in the workflow |
| Future Go versions | Yes | No | When `go.mod` is bumped, CI auto-updates — no workflow edit required |
| GitHub Actions `ubuntu-latest` | Yes | No | Standard runner; no custom or self-hosted runner (per Non-Goal NG2) |
| GitHub Actions runners (macOS, Windows) | No | N/A | Out of scope; `ubuntu-latest` covers the CI gate requirement |

---

## Platform Behavior Notes

- **`GNUmakefile` test timeout (30 s)**: The Makefile's `test` target passes `-timeout=30s`. Some tests sleep ~80 s (MEMORY note). **Affects**: Workflow must invoke `go test` directly with `-timeout=300s`, not via `make test`. The Makefile is not modified.
- **Acceptance test gate (`TF_ACC=1`)**: SDK `resource.TestCase` skips unless `TF_ACC=1`. **Affects**: Workflow must not export `TF_ACC`. All acceptance tests are auto-skipped.
- **Go module cache in GitHub Actions**: `actions/setup-go@v5` with `cache: true` caches `$GOPATH/pkg/mod` keyed on `go.sum`. Any `go.sum` change triggers a full module re-download on next run. **Affects**: Cache behavior is automatic; no explicit cache-restore step needed.
- **`coverage.out` working directory**: `go test -coverprofile=coverage.out` writes to the current directory (repo root after `actions/checkout@v4`). The `upload-artifact` and threshold steps reference the same path. **Affects**: No path prefix is needed.
- **Existing CodeQL workflow**: The repo already has `.github/workflows/codeql.yml`. The new `ci.yml` is an independent workflow with no dependencies on CodeQL. **N/A** for this feature — coexistence is straightforward.

---

## Open Questions / Risks

- Q: Initial `COVERAGE_THRESHOLD` is `2` (floor(4.1%) - 2%). Should it be `4` (floor only, no extra buffer) accepting that any new file with lower coverage fails immediately? The vision doc recommends the buffer; confirm this is acceptable. — Owner: Shobhit Sinha / team. **Blocking for implement if threshold needs a different value.**
- Q: Should the CI workflow also trigger on push to `develop` (the feature branch base per `CLAUDE.md`)? Vision scope is `main` only. Adding `develop` would catch regressions earlier. — Owner: Repo maintainer; recommend follow-up ticket.
- R: Baseline coverage is very low (~4%) because nearly all existing tests are acceptance tests requiring live DCT. The initial threshold gate at `2%` is infrastructure-level, not a meaningful quality gate. — Mitigation: Document the measured baseline and intended future target in `CLAUDE.md`; track threshold increase in a follow-up Jira ticket as unit test coverage improves.
- R: `actions/upload-artifact@v4` transient failure marks the CI job as failed even when all tests pass. — Mitigation: `if: always()` is the correct behavior per functional spec ERR-3; a transient failure is resolved by re-running the job. The `if: always()` annotation is intentional and should not be changed to `if: success()`.

---

## Acceptance Criteria

- [ ] AC-1 (FR-001): `.github/workflows/ci.yml` exists, triggers on `pull_request` to `main` and `push` to `main`
- [ ] AC-2 (FR-001): Given all unit tests pass, the `ci / unit-tests` job completes with exit code 0 and a `coverage-report` artifact is available for download
- [ ] AC-3 (FR-001): Given at least one unit test fails, the `ci / unit-tests` job exits non-zero and the PR check shows a failure status
- [ ] AC-4 (FR-001): `go-version-file: go.mod` is present in the workflow so Go version auto-updates when `go.mod` changes
- [ ] AC-5 (FR-002): Given total coverage below `COVERAGE_THRESHOLD`, the threshold step exits non-zero and prints actual vs. threshold values
- [ ] AC-6 (FR-002): Given `coverage.out` missing or empty, the threshold step exits non-zero with "coverage.out is missing or empty"
- [ ] AC-7 (FR-002): `COVERAGE_THRESHOLD` is defined in the workflow `env` block with a comment explaining the baseline and how to update it
- [ ] AC-8 (FR-003): `CLAUDE.md` contains `## CI Contract` with workflow file path, `COVERAGE_THRESHOLD` value, measured baseline, and local equivalent command
- [ ] AC-9 (FR-003): `CONTRIBUTING.md` contains `## CI` section stating `ci / unit-tests` must pass before merging and providing the local verification command
- [ ] AC-10 (FR-004): `grep "ci / unit-tests" CLAUDE.md` returns a match (exact status-check string present, not paraphrased)
- [ ] AC-11 (FR-004): `CLAUDE.md` `### Branch Protection` subsection names both required settings: "Require status checks to pass before merging" and "Require branches to be up to date before merging"

---
<!-- Cross-references checked by check-structure.sh during the design phase:
     - Every FR-* in docs/DLPXECO-14115-functional.md maps to at least one row in ### Source Files to Modify:
       FR-001 → .github/workflows/ci.yml ✓
       FR-002 → .github/workflows/ci.yml ✓
       FR-003 → CLAUDE.md, CONTRIBUTING.md ✓
       FR-004 → CLAUDE.md ✓
     - Non-Goals in docs/DLPXECO-14115-vision.md are absent from Architecture Changes:
       NG1 (acceptance tests in CI) → not present ✓
       NG2 (self-hosted runners) → not present ✓
       NG3 (Codecov/Coveralls integration) → not present ✓
       NG4 (modifying test files) → not present ✓
     Run: .claude/evals/check-structure.sh DLPXECO-14115 --step design -->
