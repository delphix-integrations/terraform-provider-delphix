# Functional Specification: DLPXECO-14142

**Jira**: DLPXECO-14142 — Document the CI contract in `CLAUDE.md`
**Generated from**: Vision doc (`docs/DLPXECO-14142-vision.md`) — Option A (doc-only),
dual-branch trigger (`main` + `develop`), `.github/workflows/ci.yml` untouched (NG1).

---

## Scope at a Glance

| Item | Value |
|---|---|
| Files modified | `CLAUDE.md` only |
| Files added | none under `.github/`, `internal/`, `examples/`, `docs/resources/`, `docs/guides/` |
| Files NOT modified | `.github/workflows/ci.yml`, `README.md`, `pull_request_template.md`, any source under `internal/` |
| Quality bar | A reviewer can re-derive every documented value from `.github/workflows/ci.yml` HEAD without running anything |

---

## FR-001: Add a "CI Contract" section to `CLAUDE.md`

### Description

Append a new top-level section (heading depth `##`) titled `CI Contract` to
`CLAUDE.md`, placed after the existing `Useful Links` section and before any
future trailing sections. The section is the single in-repo authoritative
reference for what CI does and how to reproduce it locally.

### Input

- Current `CLAUDE.md` content (must be preserved verbatim outside the new section)
- Current `.github/workflows/ci.yml` content (used as the source of truth for
  every documented value)

### Processing

1. Read existing `CLAUDE.md` and confirm no section titled `CI Contract` exists.
2. Append the new section at the end of the file (after `## Useful Links`).
3. Use the exact heading style (`##` for section, `###` for subsections) and
   table style (pipe-delimited with `---` separator row) already used elsewhere
   in `CLAUDE.md`.
4. Do not reflow or edit any other content in `CLAUDE.md`.

### Output

- `CLAUDE.md` gains exactly one new `## CI Contract` section.
- `git diff CLAUDE.md` shows only additions (no deletions, no modifications
  to pre-existing lines).

### Acceptance Criteria

- [ ] AC-1: `CLAUDE.md` contains exactly one heading matching `^## CI Contract$`.
- [ ] AC-2: All existing content in `CLAUDE.md` is preserved byte-for-byte
      outside the new section (verified via `git diff` review).
- [ ] AC-3: No file under `.github/` is modified (verified via
      `git diff .github/` showing empty output).

---

## FR-002: Document the workflow summary table

### Description

Inside the `CI Contract` section, add a `### Workflow Summary` subsection
containing a table with one row per documented attribute of the CI workflow.

### Input

- `.github/workflows/ci.yml` HEAD content (workflow name, job name, triggers,
  runner, Go-version source, test command, artifact name + retention, threshold
  value).

### Processing

The table MUST include these eight rows, in this order:

| Item | Value |
|---|---|
| Workflow file | `` `.github/workflows/ci.yml` `` |
| Workflow name | `` `ci` `` (matches `name:` in `ci.yml`) |
| Job name | `` `unit-tests` `` (matches `jobs.unit-tests.name`) |
| Status-check string | `` `ci / unit-tests` `` (the `<workflow> / <job>` form GitHub requires) |
| Trigger | `pull_request` to `main` or `develop`; `push` to `main` or `develop` |
| Runner | `` `ubuntu-latest` `` |
| Go version | Auto-detected from `go.mod` via `actions/setup-go@v5` |
| Test command | `` `go test ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s` `` |
| Coverage artifact | `` `coverage-report` `` (7-day retention) |
| Coverage threshold | `COVERAGE_THRESHOLD` in `ci.yml` `env:` block (current value: `2`%) |
| Baseline at threshold set | `2.3%` (measured 2026-06-06; unit tests only, no `TF_ACC`) |

Each value must be reproducible by reading `ci.yml` at the current HEAD.

### Output

- A markdown table rendered cleanly in `CLAUDE.md`.
- Below the table, one paragraph explicitly stating that acceptance tests
  (`TF_ACC=1`) are NOT run in CI and noting WHY (no live DCT credentials in
  CI runners; auto-excluded because the workflow does not export `TF_ACC`).

### Acceptance Criteria

- [ ] AC-1: All eleven rows listed above appear in the table in the order shown.
- [ ] AC-2: Every value in the table matches the corresponding value in
      `.github/workflows/ci.yml` at HEAD. (If `ci.yml` changes in the same PR,
      this AC compares against the post-PR `ci.yml`.)
- [ ] AC-3: The `TF_ACC` exclusion paragraph immediately follows the table.
- [ ] AC-4: The trigger row mentions BOTH `main` and `develop` for both
      `pull_request` and `push` (vision G5, C5).

---

## FR-003: Document the local-reproduction recipe

### Description

Add a `### Running the Equivalent Check Locally` subsection with three
copy-paste-able shell blocks: the primary recipe (matches CI), the
per-function breakdown, and the HTML report.

### Input

- The same `go test` command documented in FR-002 (must match exactly).

### Processing

Add three fenced `bash` code blocks, in this order:

1. **Primary recipe** — runs the exact CI command and prints the total coverage line:
   ```bash
   go test ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s
   go tool cover -func=coverage.out | tail -1
   ```
2. **Per-function breakdown** — sorts functions by coverage ascending:
   ```bash
   go tool cover -func=coverage.out | sort -t$'\t' -k3 -n
   ```
3. **HTML report** — opens an interactive line-by-line coverage view:
   ```bash
   go tool cover -html=coverage.out
   ```

### Output

Three syntactically valid `bash` blocks. The first command must be
byte-identical to the test command documented in FR-002.

### Acceptance Criteria

- [ ] AC-1: The first command line under the `Running the Equivalent Check
      Locally` heading is byte-identical to the FR-002 test command.
- [ ] AC-2: All three code blocks use the fence style `` ```bash `` (matching
      existing `CLAUDE.md` convention).
- [ ] AC-3: Running the first block locally in a clean checkout produces a
      `coverage.out` file and a `total:` line that is comparable to (within
      noise of) the documented baseline `2.3%` when run without `TF_ACC=1`.
      (This AC is informational; reviewer may spot-check.)

---

## FR-004: Document the threshold-update playbook

### Description

Add a `### Updating the Coverage Threshold` subsection containing a four-step
ordered list and a one-sentence team-agreement guard.

### Input

- The location of `COVERAGE_THRESHOLD` (in the `env:` block of `ci.yml`).

### Processing

The subsection MUST contain, in order:

1. An ordered list of exactly four steps:
   1. Measure current coverage locally using the commands above.
   2. Edit `COVERAGE_THRESHOLD` in `.github/workflows/ci.yml`.
   3. Document the old value, new value, and reason in the PR description.
   4. The change takes effect on the next CI run.
2. A one-line guard immediately after the list:
   > Do not lower the threshold without team agreement.

### Output

A 4-step ordered list and a single bold/plain advisory line.

### Acceptance Criteria

- [ ] AC-1: The ordered list contains exactly four items.
- [ ] AC-2: Step 2 names the file `.github/workflows/ci.yml` (full path).
- [ ] AC-3: The advisory line appears immediately after step 4.
- [ ] AC-4: The list does NOT instruct the editor to update `CLAUDE.md`'s
      threshold value separately — that is covered as a global rule in FR-005
      (drift management).

---

## FR-005: Document the branch-protection contract

### Description

Add a `### Branch Protection` subsection containing: a one-paragraph
explanation that branch protection is applied by a repo maintainer in the
GitHub UI (not by this doc); a table of required settings; and a bold callout
giving the exact status-check string.

### Input

- The status-check string from FR-002 (`ci / unit-tests`).

### Processing

1. Add an opening paragraph stating that the doc DESCRIBES the contract and
   does not (cannot) configure GitHub programmatically.
2. Add a table with the following four rows (in this order):

   | Setting | Required Value |
   |---|---|
   | Require status checks to pass before merging | Enabled |
   | Status check to require | `` `ci / unit-tests` `` |
   | Require branches to be up to date before merging | Enabled |
   | Who can bypass | No one (recommended) |

3. Add a bold callout line immediately after the table:
   > **The exact status-check string to enter in GitHub's branch protection UI is: `ci / unit-tests`**

### Output

Paragraph + 4-row table + bold callout line.

### Acceptance Criteria

- [ ] AC-1: The opening paragraph explicitly states the doc is descriptive,
      not programmatically enforcing.
- [ ] AC-2: The table has exactly four rows in the order shown.
- [ ] AC-3: The bold callout line appears immediately after the table and
      contains the exact backtick-wrapped string `` `ci / unit-tests` ``.
- [ ] AC-4: The subsection scope is limited to the `main` branch's contract
      (matches vision G4 and existing repo convention).

---

## FR-006: Drift-management note (process rule)

### Description

Inside the `CI Contract` section (anywhere reviewer-visible), add a short note
that any future PR which changes the workflow file, the threshold value, the
job name, the workflow name, or the trigger branches MUST also update the
corresponding rows in `CLAUDE.md` in the same PR.

### Input

- Vision constraint C2 (drift risk acknowledgement).

### Processing

Add one short paragraph (1–3 sentences) stating the process rule. Phrase it as
a contributor instruction, not as a tool-enforced check.

### Output

One paragraph inside the `CI Contract` section.

### Acceptance Criteria

- [ ] AC-1: A reviewer reading the section can identify, in one read-through,
      what to update in `CLAUDE.md` if they edit `ci.yml`.
- [ ] AC-2: The note does NOT promise tooling enforcement that does not exist.

---

## FR-007: Confirm `ci.yml` is untouched (negative requirement / NG1)

### Description

This is a negative requirement that holds across the PR.

### Input

- The PR's `git diff` output.

### Processing

CI / reviewer verifies that `git diff --stat` on the merge candidate shows
ZERO bytes changed under `.github/workflows/`.

### Output

A clean diff scoped to `CLAUDE.md` (and any companion spec files this
pipeline produces under `docs/`).

### Acceptance Criteria

- [ ] AC-1: `git diff origin/main -- .github/workflows/ci.yml` is empty.
- [ ] AC-2: `git diff origin/main -- .github/` is empty.

---

## Quality Rules (apply to all FRs)

- **QR-1 — No external links to ephemeral infrastructure.** No links to
  Jenkins, Buildkite, internal dashboards, or specific GitHub Actions run
  URLs. (Vision C3.)
- **QR-2 — Markdown style.** Match existing `CLAUDE.md`: `##` for sections,
  `###` for subsections, pipe-delimited tables, fenced code blocks with
  language tag. (Vision C4.)
- **QR-3 — No behavioural changes.** This ticket adds documentation only. If
  during review someone discovers a real defect in `ci.yml`, file a new
  ticket — do not roll a fix into this PR. (Vision NG1.)
- **QR-4 — Single source of truth for behaviour stays in `ci.yml`.** Where
  the doc duplicates a value, the value must be reproducible from `ci.yml`
  HEAD. (Vision C1.)
- **QR-5 — Trigger description must name both branches explicitly.** Never
  use "the default branch" or "the protected branch" as a stand-in for
  `main` / `develop`. (Vision C5, G5.)

---

## Out of Scope (mirrors vision non-goals)

- Editing `.github/workflows/ci.yml` (NG1)
- Configuring GitHub branch protection programmatically (NG2)
- Changing the coverage threshold value (NG3)
- Adding tests, lint, or new CI jobs (NG4)
- Updating `README.md` or `pull_request_template.md` (NG5)
- Documenting acceptance tests as part of CI (NG6)
