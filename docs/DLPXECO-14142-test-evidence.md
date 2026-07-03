# DLPXECO-14142 — Test Evidence

**Phase**: test
**Date**: 2026-06-07
**Branch**: DLPXECO-14141-ci
**Working-tree state**: `CLAUDE.md` modified (+9 lines, 0 deletions); no other repo files modified by this ticket.

This ticket is **doc-only** (see `docs/DLPXECO-14142-design.md`). Verification is structural
(grep/regex/diff against `CLAUDE.md`) plus the CI-equivalent `go test` repro from FR-003 of
the functional spec. No new Go tests, fixtures, or test infra were added.

---

## Verification Matrix Results

All checks reference `docs/DLPXECO-14142-test-plan.md`.

| Check | Tied to | Command | Expected | Observed | Result |
|---|---|---|---|---|---|
| V-1 | FR-001 AC-1 | `grep -c '^## CI Contract$' CLAUDE.md` | `1` | `1` | PASS |
| V-2 | FR-001 AC-2 | `git diff CLAUDE.md \| grep -c '^-[^-]'` (deletion count) | `0` | `0` | PASS |
| V-3 | FR-001 AC-3 / FR-007 AC-1, AC-2 | `git diff -- .github/ \| wc -l` | `0` | `0` | PASS |
| V-4 | FR-002 AC-1 | Workflow Summary pipe-row count (header+sep+11 data ≥ 13) | `≥13` | `13` | PASS |
| V-5 | FR-002 AC-2 | Per-row cross-read of table values vs `ci.yml` HEAD | match | match (table from DLPXECO-14115 still mirrors `ci.yml` HEAD; no `ci.yml` change in this PR) | PASS |
| V-6 | FR-002 AC-3 | `TF_ACC` exclusion paragraph follows the table | present | present at CLAUDE.md L197 | PASS |
| V-7 | FR-002 AC-4 | Trigger row names main+develop twice | match | `\| Trigger \| pull_request to main or develop; push to main or develop \|` | PASS |
| V-8 | FR-003 AC-1 | First `go test` under local-repro byte-identical to FR-002 | match | `go test ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s` | PASS |
| V-9 | FR-003 AC-2 | Three ```` ```bash ```` fences in local-repro subsection | `3` | `3` | PASS |
| V-10 | FR-003 AC-3 (informational) | Run CI command locally; see V-25 below | `total: ~2.3%` | `total: 2.3%` | PASS |
| V-11 | FR-004 AC-1 | Threshold playbook has 4 ordered items | `4` | `4` | PASS |
| V-12 | FR-004 AC-2 | Step 2 names `.github/workflows/ci.yml` | match | match (CLAUDE.md L220) | PASS |
| V-13 | FR-004 AC-3 | "Do not lower the threshold without team agreement." present | present | present (CLAUDE.md L224) | PASS |
| V-14 | FR-005 AC-1 | Branch Protection opening paragraph says descriptive, not enforcing | `describes\|configure` match | `This document describes the required contract; it does not programmatically configure GitHub.` | PASS |
| V-15 | FR-005 AC-2 | Branch Protection pipe-row count (header+sep+4 data ≥ 6) | `≥6` | `6` | PASS |
| V-16 | FR-005 AC-3 | Bold callout with `ci / unit-tests` immediately after table | present | `**The exact status-check string to enter in GitHub's branch protection UI is: \`ci / unit-tests\`**` | PASS |
| V-17 | FR-005 AC-4 | Branch Protection subsection mentions `main` | present | 2 occurrences | PASS |
| V-18 | FR-006 AC-1 | `### Drift Management` present; paragraph enumerates duplicated values | present + 5 values named | `### Drift Management` heading present; paragraph names `threshold`, `status-check string`, `trigger branches`, `workflow name`, `job name` | PASS |
| V-19 | FR-006 AC-2 | Note explicitly says "process rule" / "not tooling" | match | `This is a process rule, not a tooling-enforced check` | PASS |
| V-20 | QR-1 | No Jenkins / Buildkite / Actions-run URLs in CI Contract | no matches | rc=1 (grep no match) | PASS |
| V-21 | QR-2 | Markdown style consistent (`##` / `###` / pipe tables / fenced bash) | visual | consistent — no new constructs | PASS |
| V-22 | QR-3 | No source/test diff vs HEAD outside `CLAUDE.md` and `docs/` | `0` lines | `0` | PASS |
| V-23 | QR-4 | All duplicated values reproducible from `ci.yml` HEAD | covered by V-5 | covered | PASS |
| V-24 | QR-5 | No "default branch" / "the protected branch" stand-ins | no matches | rc=1 (grep no match) | PASS |

**Structural verdict: 24 / 24 PASS.**

---

## V-25: CI-Equivalent Unit-Test Run (FR-003 AC-3 / Pre-Flight #5)

The CI workflow (`.github/workflows/ci.yml`, job `unit-tests`) runs:

```bash
go test ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s
```

The same command was executed locally on the working tree containing my CLAUDE.md
change. Environment was stripped of `TF_ACC` and per-resource credential env vars
(`ACC_ENV_ENGINE_ID`, `DATASOURCE_ID`, `ORACLE_DSOURCE_SOURCE_VALUE`) to match
CI conditions exactly.

### Result

```
ok    terraform-provider-delphix                       0.887s    coverage: 0.0% of statements
ok    terraform-provider-delphix/internal/provider    0.884s    coverage: 2.3% of statements
EXIT=0
total: (statements) 2.3%
```

### Baseline Comparison

| Tree state | Total coverage | Exit code |
|---|---|---|
| HEAD without DLPXECO-14142 change | `2.3%` | `0` |
| HEAD with DLPXECO-14142 change | `2.3%` | `0` |

The doc-only edit produced **zero behavioural delta** — coverage is identical to the pre-change baseline, comfortably above the `2%` threshold enforced by `ci.yml`.

### Note on Earlier False Negative

A first invocation of the CI command in this phase produced acceptance-test failures
(`ACC_ENV_ENGINE_ID must be set`, `DATASOURCE_ID must be set`, …). This was traced to
environment-variable leakage from the parent shell, not to my edit. After clearing
those env vars and the test cache (`go clean -testcache`), the command passed on
both the pre-change and post-change trees. The CI runner has no such env leakage,
so this would not reproduce in CI.

---

## Files Inspected During Verification

- `CLAUDE.md` (working tree, post-edit)
- `.github/workflows/ci.yml` (HEAD; not modified — confirmed by `git diff -- .github/`)
- `docs/DLPXECO-14142-design.md`
- `docs/DLPXECO-14142-functional.md`
- `docs/DLPXECO-14142-test-plan.md`
