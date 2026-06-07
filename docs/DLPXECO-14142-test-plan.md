# DLPXECO-14142 — Test Plan

Scope: doc-only change. All verification is structural (grep / regex /
diff) against `CLAUDE.md` and the repo's `.github/` tree. No Go unit
tests, no acceptance tests, no new test files are added.

The existing CI workflow (`ci / unit-tests`) continues to gate unit-test
+ coverage on the PR for free; it does not need a new test for this
ticket because no Go code is modified.

---

## Verification Matrix

| Check ID | Tied to | What to verify | How to verify |
|---|---|---|---|
| V-1 | FR-001 AC-1 | `CLAUDE.md` has exactly one `^## CI Contract$` heading | `grep -c '^## CI Contract$' CLAUDE.md` returns `1` |
| V-2 | FR-001 AC-2 | All existing content outside the additive edit is preserved byte-for-byte | `git diff -- CLAUDE.md` shows only additions inside `## CI Contract`; no deletions outside the new subsection |
| V-3 | FR-001 AC-3 / FR-007 AC-1 / AC-2 | `.github/` is untouched | `git diff origin/main -- .github/` produces empty output; `git diff origin/main -- .github/workflows/ci.yml` produces empty output |
| V-4 | FR-002 AC-1 | The Workflow Summary table has all 11 rows in order | Visual inspection + `awk '/^### Workflow Summary$/,/^### Running/' CLAUDE.md \| grep -c '^\|'` returns at least `13` (1 header + 1 separator + 11 data rows) |
| V-5 | FR-002 AC-2 | Each table value matches `ci.yml` HEAD | Manual cross-read: `Workflow name` row matches `name:` in `ci.yml`; `Job name` row matches `jobs.unit-tests.name`; `Trigger` row matches `on.pull_request.branches` + `on.push.branches`; `Test command` matches the shell command under `Run unit tests with coverage`; `Coverage artifact` matches `actions/upload-artifact@v4 name:`; `Coverage threshold` value matches `env.COVERAGE_THRESHOLD` |
| V-6 | FR-002 AC-3 | The `TF_ACC` exclusion paragraph follows the table | `grep -A1 'no \`TF_ACC\`' CLAUDE.md` (or visual) shows the paragraph immediately after the table |
| V-7 | FR-002 AC-4 | Trigger row names both `main` and `develop` for both `pull_request` and `push` | Grep: `grep -E 'Trigger.*main.*develop.*main.*develop' CLAUDE.md` returns a match |
| V-8 | FR-003 AC-1 | First bash command under `Running the Equivalent Check Locally` is byte-identical to FR-002 test command | `awk '/^### Running the Equivalent/,/^### Updating/' CLAUDE.md \| grep -m1 '^go test '` returns `go test ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s` |
| V-9 | FR-003 AC-2 | All three code blocks use ` ```bash ` fence | `awk '/^### Running the Equivalent/,/^### Updating/' CLAUDE.md \| grep -c '^```bash$'` returns `3` |
| V-10 | FR-003 AC-3 (informational) | Local recipe produces a coverage.out and a total line | `go test ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s && go tool cover -func=coverage.out \| tail -1` prints a `total:` line in the same ballpark as the documented baseline `2.3%` (within ±0.5 percentage points when run without `TF_ACC=1`) |
| V-11 | FR-004 AC-1 | Threshold playbook has exactly four ordered-list items | `awk '/^### Updating the Coverage Threshold$/,/^### Branch Protection$/' CLAUDE.md \| grep -cE '^[1-4]\. '` returns `4` |
| V-12 | FR-004 AC-2 | Step 2 names the file `.github/workflows/ci.yml` | `awk '/^### Updating the Coverage Threshold$/,/^### Branch Protection$/' CLAUDE.md \| grep -F '.github/workflows/ci.yml'` returns a match |
| V-13 | FR-004 AC-3 | "Do not lower the threshold without team agreement" appears immediately after step 4 | Visual inspection; grep for the exact phrase |
| V-14 | FR-005 AC-1 | Branch Protection opening paragraph says the doc is descriptive, not enforcing | `awk '/^### Branch Protection$/,/^### Drift Management$/' CLAUDE.md \| grep -E 'describes|contract|not.*configure'` returns at least one match |
| V-15 | FR-005 AC-2 | Branch Protection table has exactly four rows in the documented order | `awk '/^### Branch Protection$/,/^### Drift Management$/' CLAUDE.md \| grep -c '^\|'` returns at least `6` (header + separator + 4 data rows) |
| V-16 | FR-005 AC-3 | Bold callout `ci / unit-tests` line follows the table | `grep -F '**The exact status-check string to enter in GitHub' CLAUDE.md` returns a match; the line contains `` `ci / unit-tests` `` |
| V-17 | FR-005 AC-4 | Subsection talks about `main` branch | `awk '/^### Branch Protection$/,/^### Drift Management$/' CLAUDE.md \| grep -F 'main'` returns a match |
| V-18 | FR-006 AC-1 | A reviewer can identify in one read-through what to update if `ci.yml` changes | `grep -F 'Drift Management' CLAUDE.md` returns a match; the paragraph explicitly enumerates the duplicated values (threshold, status-check string, trigger branches, workflow name, job name) |
| V-19 | FR-006 AC-2 | The note does NOT promise tooling enforcement | `awk '/^### Drift Management$/,/^$/' CLAUDE.md \| grep -iE 'process rule|not.*tooling'` returns a match (positive assertion of process-only nature) |
| V-20 | QR-1 | No links to Jenkins / Buildkite / internal dashboards / specific Actions run URLs in the new section | `awk '/^## CI Contract$/,EOF' CLAUDE.md \| grep -iE 'jenkins\|buildkite\|actions/runs/'` returns no matches |
| V-21 | QR-2 | Markdown style is consistent | Visual inspection: `##` for section, `###` for subsections, pipe-delimited tables, fenced `` ```bash `` blocks |
| V-22 | QR-3 | No behavioural code changes | `git diff origin/main -- ':!CLAUDE.md' ':!docs/'` returns empty |
| V-23 | QR-4 | All duplicated values are reproducible from `ci.yml` HEAD | Covered by V-5 |
| V-24 | QR-5 | Trigger description names both branches; never uses "default branch" / "protected branch" as a stand-in | `grep -iE 'default branch\|the protected branch' CLAUDE.md` returns no matches in the new section |

---

## Pre-Flight Sanity (run locally before raising PR)

```bash
# 1. Confirm only CLAUDE.md (and docs/) are touched
git diff --stat origin/main -- ':!docs/' | grep -v '^ CLAUDE.md ' && echo "FAIL: extra files touched" || echo "PASS: only CLAUDE.md modified outside docs/"

# 2. Confirm .github/ is empty in the diff
git diff origin/main -- .github/ | wc -l  # expect 0

# 3. Confirm exactly one CI Contract heading
grep -c '^## CI Contract$' CLAUDE.md       # expect 1

# 4. Confirm Drift Management subsection was added
grep -c '^### Drift Management$' CLAUDE.md  # expect 1

# 5. Local CI repro (unit tests + coverage)
go test ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s
go tool cover -func=coverage.out | tail -1  # expect total ≈ 2.3%, > 2% threshold
```

All checks must return the expected values before raising the PR.

---

## CI Behaviour Expectations

| Aspect | Expectation |
|---|---|
| `ci / unit-tests` status check | PASS — no Go code changed; coverage unchanged from current baseline `2.3%`, comfortably above the `2%` threshold |
| `codeql.yml` (if it runs) | PASS — no code changed |
| Any markdown lint hook | PASS — wording matches existing `CLAUDE.md` style |

No new CI jobs are added; no thresholds are raised; no tests are added.
