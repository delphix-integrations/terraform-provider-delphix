# Vision: DLPXECO-14115

## Problem Statement

The `terraform-provider-delphix` repository has no automated CI test gate on pull requests: unit tests run only when a developer remembers to run `make test` locally, and there is no coverage threshold to prevent regressions from landing undetected. This means broken code can be merged, test coverage can silently erode over time, and reviewers have no machine-generated evidence that tests pass before they approve a PR.

## Goals

- G1: Add a GitHub Actions workflow that automatically runs `make test` on every pull request and push to `main`, blocking merge when tests fail
- G2: Produce a Go test-coverage artifact on every CI run and enforce a minimum coverage threshold so coverage cannot regress without an explicit decision to lower the bar
- G3: Document the CI contract (workflow name, threshold, how to run locally, how to update the threshold) in `CLAUDE.md` and `CONTRIBUTING.md` so contributors and reviewers share a single authoritative reference

## Non-Goals

- NG1: Does not add or modify acceptance tests (`TF_ACC=1`) — those require live DCT infrastructure and are outside the scope of a lightweight CI gate
- NG2: Does not set up self-hosted or custom runners — the workflow targets the standard `ubuntu-latest` GitHub-hosted runner
- NG3: Does not integrate external coverage reporting services (Codecov, Coveralls, SonarQube) in this iteration — the artifact upload and threshold check are self-contained within GitHub Actions
- NG4: Does not change the content of any existing test file — test authorship is out of scope; this ticket adds the infrastructure around the tests that already exist

## Success Criteria

- SC1: A GitHub Actions workflow file (`.github/workflows/ci.yml`) exists, triggers on `pull_request` to `main` and `push` to `main`, runs `go test ./... -coverprofile=coverage.out`, and the Actions check is visible on every PR
- SC2: The workflow uploads `coverage.out` as a build artifact and fails the job if the overall coverage percentage falls below the defined threshold
- SC3: `CLAUDE.md` and `CONTRIBUTING.md` both contain a CI section that names the workflow, states the current threshold, and explains how to run the equivalent check locally
- SC4: The branch-protection contract is documented — a `## CI Contract` section in `CLAUDE.md` states which status check must pass before merge is permitted

## Stakeholders

| Stakeholder | Interest |
|-------------|----------|
| Feature contributors | Fast, automated feedback that their PR does not break existing tests before requesting review |
| Code reviewers | Machine-generated evidence of test pass and coverage level, reducing manual verification burden |
| Repo maintainers (Delphix integrations team) | Prevent coverage regression and broken builds from landing on `main` without explicit sign-off |
| Security / compliance | Auditability — every merged commit has an associated CI run result in GitHub Actions history |

## Constraints

- Must use GitHub Actions (the repo already uses `.github/workflows/` for CodeQL and release; adding a new workflow file is the natural extension)
- Go version must match `go.mod` (currently `go 1.25`) — the workflow must use the same version to avoid false failures
- The initial coverage threshold must be set at or below the current measured baseline so the PR that adds CI is not immediately blocked by its own check; the threshold can be tightened in a follow-up
- No new Go module dependencies may be introduced — coverage is generated with `go test -coverprofile` (standard library tooling, no third-party needed)
- The workflow must complete within the 6-hour GitHub Actions timeout for the `ubuntu-latest` runner; in practice `make test` completes in under 5 minutes

## Risks

| Risk | Likelihood | Impact | Mitigation |
|------|------------|--------|------------|
| Current unit-test coverage is below any reasonable threshold, making it impossible to set a non-trivial gate without fixing tests first | Medium | High | Measure baseline coverage before setting the threshold; set threshold at `floor(baseline) - 2%` to give a small buffer, document the measured baseline and intended target in `CLAUDE.md` |
| Acceptance tests (`TF_ACC=1`) are accidentally included in the CI run, causing the job to hang waiting for infrastructure that is not present | Low | High | Scope the test command explicitly to unit tests only (`go test ./... -run "^Test[^A][^c][^c]"` or rely on the existing `t.Skip` guards in acceptance tests); document the exclusion pattern |
| Test parallelism (`-parallel=4`) plus the 30-second timeout in `GNUmakefile` is too tight for the CI runner, causing flaky timeouts | Low | Medium | Run with `-timeout=300s` in CI (consistent with the memory note about the engine_configuration create test sleeping ~80 s); keep the Makefile unchanged |
| Go version mismatch between local developer environment and GitHub Actions causes spurious build failures | Low | Medium | Pin the `go-version` in the workflow to match `go.mod`; use `actions/setup-go@v5` with `go-version-file: go.mod` to auto-detect |
| Branch protection rules are not actually configured on the GitHub repo, making the status check advisory only | Medium | Medium | Document the required branch protection settings in `CLAUDE.md`; the repo maintainer must apply them manually in GitHub settings after the workflow lands |

---
<!-- Cross-reference: Goals (G1, G2, G3) map to FR descriptions in the functional spec.
     Success Criteria (SC1–SC4) map to Acceptance Criteria in FR-* entries.
     Constraints and Risks inform the Quality Rules and Edge Cases sections. -->
