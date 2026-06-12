# Functional Specification: DLPXECO-14115

**Jira**: DLPXECO-14115 — [S2 Hard Gate] CI test coverage enforcement for terraform-provider-delphix
**Generated from**: Acceptance criteria in Jira ticket and vision doc (docs/DLPXECO-14115-vision.md)

---

## FR-001: GitHub Actions CI Workflow

### Description

Adds a GitHub Actions workflow (`.github/workflows/ci.yml`) that runs the unit test suite (`go test ./... -coverprofile=coverage.out`) on every pull request targeting `main` and on every push to `main`, producing a Go coverage profile as an artifact and failing the job if tests fail.

### Input

- GitHub event: `pull_request` (target branch: `main`) or `push` (branch: `main`)
- Source code in the repository at the triggering commit SHA
- Go version: resolved from `go.mod` via `actions/setup-go@v5` with `go-version-file: go.mod`
- No environment variables required (unit tests do not contact external services)

### Processing

1. Trigger on `pull_request` with `branches: [main]` and `push` with `branches: [main]`
2. Check out repository at the triggering SHA with `actions/checkout@v4`
3. Set up Go using `actions/setup-go@v5` with `go-version-file: go.mod` and caching enabled (`cache: true`)
4. Run `go test ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s` (not via `make test` to avoid the 30 s Makefile timeout and to capture coverage output explicitly)
5. If the command exits non-zero, the job fails immediately — no further steps run
6. Upload `coverage.out` as a workflow artifact named `coverage-report` with a 7-day retention period using `actions/upload-artifact@v4`

### Output

- Success: GitHub Actions check status `ci / unit-tests` shows green on the PR or commit; `coverage.out` artifact available for download
- Failure (test failure): job exits non-zero; GitHub Actions check shows red; merge blocked (once branch protection is configured per FR-004)
- Side effect: coverage profile (`coverage.out`) uploaded as artifact and available as input to FR-002

### Acceptance Criteria

- [ ] AC-1: Given a PR to `main` with all tests passing, when the CI workflow triggers, then the `ci / unit-tests` job completes with exit code 0 and a `coverage-report` artifact is available
- [ ] AC-2: Given a PR to `main` where at least one unit test fails, when the CI workflow triggers, then the `ci / unit-tests` job exits non-zero and the PR check shows a failure status
- [ ] AC-3: Given a push directly to `main`, when the CI workflow triggers, then the same test job runs (not only on pull_request events)
- [ ] AC-4: Given the workflow file references `go-version-file: go.mod`, when a developer updates the Go version in `go.mod`, then CI automatically picks up the new version without editing the workflow file

---

## FR-002: Coverage Threshold Enforcement

### Description

Adds a coverage-check step to the CI workflow that parses the Go coverage profile, computes the total coverage percentage, and fails the job if the percentage falls below a defined minimum threshold stored in the workflow file.

### Input

- `coverage.out`: Go coverage profile produced by FR-001's test step (file must exist in the workspace before this step runs)
- `COVERAGE_THRESHOLD`: a numeric percentage value (e.g. `50`) defined as an environment variable or workflow-level variable in `ci.yml`

### Processing

1. After the test step completes successfully, run `go tool cover -func=coverage.out` to produce a per-function coverage report
2. Extract the total coverage percentage from the last line of the output (format: `total:	(statements)	XX.X%`)
3. Strip the `%` suffix and compare the numeric value against `COVERAGE_THRESHOLD` using shell arithmetic (`awk` or `bc`)
4. If coverage < threshold: print a human-readable message (e.g. "Coverage 47.3% is below threshold 50%") and exit non-zero, failing the job
5. If coverage >= threshold: print the coverage percentage and exit 0
6. The threshold value must be documented as a comment in `ci.yml` explaining how to update it

### Output

- Success: step exits 0; coverage percentage printed to workflow log
- Failure: step exits non-zero; error message printed with actual vs. threshold values; job fails; PR check shows red

### Acceptance Criteria

- [ ] AC-1: Given `coverage.out` with total coverage of 55% and `COVERAGE_THRESHOLD=50`, when the threshold step runs, then it exits 0 and prints the coverage percentage
- [ ] AC-2: Given `coverage.out` with total coverage of 48% and `COVERAGE_THRESHOLD=50`, when the threshold step runs, then it exits non-zero and prints "Coverage 48.0% is below threshold 50%"
- [ ] AC-3: Given `COVERAGE_THRESHOLD` changed to a higher value in `ci.yml`, when the next CI run completes, then the new threshold is enforced without any code change outside `ci.yml`
- [ ] AC-4: Given the `coverage.out` file is missing (e.g. test step was skipped), when the threshold step runs, then it exits non-zero with an informative error rather than silently passing

---

## FR-003: Documentation Update — CLAUDE.md and CONTRIBUTING.md

### Description

Updates `CLAUDE.md` and `CONTRIBUTING.md` to document the CI contract: the workflow name, the current coverage threshold, how to run the equivalent check locally before pushing, and what contributors must do if they intentionally lower coverage.

### Input

- `CLAUDE.md` current content (existing file in the repository root)
- `CONTRIBUTING.md` current content (existing file in the repository root)
- The workflow name and threshold value decided during FR-001 and FR-002 implementation

### Processing

1. In `CLAUDE.md`, add a new `## CI Contract` section containing:
   - Workflow file path (`.github/workflows/ci.yml`) and the triggering conditions
   - The current `COVERAGE_THRESHOLD` value and the measured baseline at the time it was set
   - Local equivalent command: `go test ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s && go tool cover -func=coverage.out`
   - How to read the coverage report and identify uncovered functions
   - How to update the threshold (edit `COVERAGE_THRESHOLD` in `ci.yml`, document the reason in the PR)
   - Which GitHub Actions status check must pass before merge is permitted (links to FR-004)
2. In `CONTRIBUTING.md`, add a `## CI` section after the existing contribution steps, containing:
   - A one-paragraph summary stating that all PRs must pass the `ci / unit-tests` check
   - The local command to verify before pushing
   - A note that acceptance tests (`TF_ACC=1`) are not run in CI and require a live DCT instance

### Output

- `CLAUDE.md`: updated with `## CI Contract` section (non-empty, no placeholders)
- `CONTRIBUTING.md`: updated with `## CI` section (non-empty, no placeholders)

### Acceptance Criteria

- [ ] AC-1: Given `CLAUDE.md` after this change, when a developer searches for "CI Contract", then they find a section that names the workflow file, the current threshold, and the local equivalent command
- [ ] AC-2: Given `CONTRIBUTING.md` after this change, when a new contributor reads it, then they find a `## CI` section stating that the `ci / unit-tests` check must pass before merge
- [ ] AC-3: Given neither file contains `TBD` or `TODO` in the added sections, when the PR is raised, then the documentation is complete as written

---

## FR-004: Branch Protection Contract Documentation

### Description

Documents the required GitHub branch protection settings in `CLAUDE.md` so that repo maintainers can configure the GitHub repository to enforce the CI gate, and so the contract is auditable in the source tree even if GitHub settings are changed.

### Input

- The workflow job name defined in FR-001 (e.g. `unit-tests` under workflow `ci`)
- The GitHub repository's branch protection settings page (configured out-of-band by a maintainer, not by this code change)

### Processing

1. Under the `## CI Contract` section added in FR-003, add a `### Branch Protection` subsection documenting:
   - Required status check name: the exact string that appears in GitHub's branch protection UI (format: `<workflow-name> / <job-name>`, e.g. `ci / unit-tests`)
   - Setting: "Require status checks to pass before merging" must be enabled for the `main` branch
   - Setting: "Require branches to be up to date before merging" should be enabled to prevent stale-branch bypasses
   - Note: this document describes the intended contract; a repo maintainer must apply these settings in GitHub Settings → Branches → Branch protection rules
2. The subsection must be a static document — it does not programmatically configure GitHub (no `gh` CLI calls in the workflow itself)

### Output

- `CLAUDE.md`: `### Branch Protection` subsection added under `## CI Contract`, containing the exact status check name and the required GitHub settings

### Acceptance Criteria

- [ ] AC-1: Given the `### Branch Protection` subsection in `CLAUDE.md`, when a maintainer reads it, then they find the exact status check string to enter in GitHub's branch protection UI
- [ ] AC-2: Given the subsection content, when `grep "ci / unit-tests" CLAUDE.md` is run, then the exact check name appears in the output (the name must be literal, not paraphrased)

---

## Quality Rules

| Rule | Description | Enforcement | Status | Evidence |
|------|-------------|-------------|--------|----------|
| API backward compatibility preserved | No existing resource schema or behavior is changed by this ticket — the CI workflow is additive infrastructure only | `git diff --name-only` on the PR must show only `.github/workflows/ci.yml`, `CLAUDE.md`, and `CONTRIBUTING.md`; no `internal/provider/*.go` changes | | |
| Migration path provided | Contributors must be able to run the same check locally before pushing — the local command must be documented and tested | `CLAUDE.md` contains the local equivalent command; it is verified runnable in the PR author's dev environment | | |
| No acceptance tests in CI | The unit-test command must not accidentally run `TF_ACC=1` tests that require live infrastructure | `grep -v "TF_ACC" .github/workflows/ci.yml` returns the test command line; the workflow does not export `TF_ACC` | | |
| Threshold set at or below current baseline | The initial threshold must not block the first CI run | Baseline measured with `go test ./... -coverprofile=coverage.out && go tool cover -func=coverage.out` before the threshold is committed; threshold value documented in `ci.yml` comment | | |

---

## Edge Cases

- EC-1: `coverage.out` is empty (no test files matched) → `go tool cover -func` prints nothing; the threshold step must detect the empty file and exit non-zero with "no coverage data found" rather than silently reporting 0%
- EC-2: A developer adds a new file with 0% coverage, pulling total below the threshold → CI fails as designed; contributor must either add tests or (with team agreement) lower the threshold in `ci.yml` with a documented reason
- EC-3: Go version in `go.mod` is updated to a version not yet available in `actions/setup-go` → CI fails at setup, not at test; the error message from `actions/setup-go` is clear; the fix is to wait for the action to add support or pin an available version
- EC-4: Two PRs are open simultaneously; one merges and changes coverage; the other PR's CI run used stale coverage from the branch base → each PR runs CI against its own commit; the merged-to-main threshold is what matters; the second PR will re-run CI after rebasing
- EC-5: `make test` timeout (30 s) is too short for some tests in CI → workflow uses `go test -timeout=300s` directly, not `make test`, so the Makefile timeout does not apply

## Error Scenarios

- ERR-1: GitHub Actions runner cannot install Go (network partition, registry unavailable) → `actions/setup-go` step fails with a non-zero exit; the entire job is marked failed; no coverage artifact is produced; re-run the job when the runner recovers
- ERR-2: `go test` fails to compile (syntax error introduced in a PR) → compilation error printed to log; job exits non-zero; no `coverage.out` produced; the threshold step must handle the missing file gracefully (see EC-1)
- ERR-3: `actions/upload-artifact` fails to upload `coverage.out` (storage backend error) → the artifact step fails; the overall job should be marked failed (use `if: always()` on the upload step so it runs even after test failures, but treat upload failure as a hard error)

## Performance Considerations

N/A — `go test ./...` for this repo's unit test suite completes in under 60 seconds on a standard GitHub-hosted runner. The `-parallel=4` flag matches the existing `GNUmakefile` setting. No caching of test results is needed at this coverage level; Go module cache is handled by `actions/setup-go` with `cache: true`.

---
<!-- Cross-reference: FR descriptions map to Goals (G1=FR-001, G2=FR-002, G3=FR-003+FR-004) in the vision doc.
     FR Acceptance Criteria satisfy Success Criteria (SC1=FR-001 AC-1/AC-3, SC2=FR-002 AC-1/AC-2, SC3=FR-003 AC-1/AC-2, SC4=FR-004 AC-1).
     Quality Rules and Edge Cases address Constraints and Risks from the vision doc.
     FR-IDs defined here are referenced in tasks-template (Spec References) and validation-template (FR Coverage). -->
