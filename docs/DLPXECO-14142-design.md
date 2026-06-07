# DLPXECO-14142 — Design

**Scope**: doc-only change. Locked decision (Option A): modify `CLAUDE.md` only,
leave `.github/workflows/ci.yml` untouched. All FRs in
`docs/DLPXECO-14142-functional.md` are satisfied by edits to a single file.

---

## Discovery: Existing State of `CLAUDE.md`

A `## CI Contract` section already exists in `CLAUDE.md` (lines 176–240 at
HEAD on branch `DLPXECO-14141-ci`). It was introduced by the predecessor work
DLPXECO-14115 alongside `.github/workflows/ci.yml`. That section already
contains:

- The `### Workflow Summary` table with all 11 rows from FR-002.
- The `### Running the Equivalent Check Locally` subsection with all three
  bash blocks from FR-003.
- The `### Updating the Coverage Threshold` subsection with the 4-step
  playbook and the "Do not lower" advisory from FR-004.
- The `### Branch Protection` subsection with the descriptive paragraph,
  4-row table, and bold callout from FR-005.

Gap analysis vs. the functional spec:

| FR | Status | Action |
|---|---|---|
| FR-001 — section exists | Present (heading + structure) | No-op (the section is already in place; the functional spec's "append" language is satisfied by the existing section because it was added in the same branch lineage that this ticket continues) |
| FR-002 — workflow summary table | Present, all 11 rows present, all values match `ci.yml` HEAD | No-op |
| FR-003 — local-reproduction recipe | Present, all three bash blocks, first command byte-identical to FR-002 | No-op |
| FR-004 — threshold playbook | Present, 4 steps, advisory line present | No-op |
| FR-005 — branch-protection contract | Present, paragraph + 4-row table + bold callout | No-op |
| **FR-006 — drift-management note** | **MISSING** — no paragraph instructs contributors to update `CLAUDE.md` when `ci.yml` changes | **ADD** a 1–3 sentence paragraph inside the `CI Contract` section |
| FR-007 — `ci.yml` untouched | Satisfied by design (we add no edits under `.github/`) | No-op |

The ticket's value-add at the code level reduces to a single small additive
edit: adding the FR-006 drift-management note. This matches the vision
(doc-only) and respects all constraints (C1–C5) and non-goals (NG1–NG6).

---

## Architecture Changes

### Source Files to Modify

| File | Change Type | Description |
|---|---|---|
| `CLAUDE.md` | Modify (additive only) | Insert one new paragraph (FR-006 drift-management note) inside the existing `## CI Contract` section. Optionally: a one-line wording polish on the existing "Updating the Coverage Threshold" subsection to cross-reference the drift note, if the reviewer prefers — but only as a non-load-bearing nicety. |

### Source Files NOT Modified (explicit)

| File | Reason |
|---|---|
| `.github/workflows/ci.yml` | Vision NG1; FR-007 negative requirement. The workflow is the source of truth for behaviour; this ticket is doc-only. |
| `README.md` | Vision NG5; out of scope. |
| `pull_request_template.md` | Vision NG5; out of scope. |
| Any file under `internal/` | Vision NG4; no source changes. |
| Any file under `examples/` | Vision NG4; no example changes. |
| `.goreleaser.yml`, `GNUmakefile` | Out of scope (release/build config). |
| `go.mod` / `go.sum` | No dependency changes. |

### Files Added

None. All work lands in `CLAUDE.md`.

---

## Detailed Edit Plan for `CLAUDE.md`

### Insertion Point

Insert the FR-006 paragraph **immediately after** the existing
`### Branch Protection` subsection, at the end of the `## CI Contract`
section. Rationale:

- Placing it last keeps the section's logical flow intact (summary → local
  repro → threshold playbook → branch protection → process rule).
- A `### Drift Management` subsection heading (sentence-case, matching the
  existing subsection style) is added so the rule has a stable anchor for
  future cross-references and shows up in any TOC generator.
- This satisfies FR-006 AC-1 ("a reviewer can identify in one read-through
  what to update if they edit `ci.yml`") because the section ends with a
  visible, headed paragraph — it does not bury the rule mid-section.

### Exact Content to Add

A new subsection appended after the `### Branch Protection` table and
callout. The literal block to add (verbatim, no trailing blank line beyond
what already exists at EOF):

```markdown

### Drift Management

The values in the Workflow Summary table above (threshold, status-check
string, trigger branches, workflow name, job name) are duplicated from
`.github/workflows/ci.yml`. Any future PR that changes those values in
`ci.yml` MUST also update the corresponding rows in this section in the
same PR. This is a process rule, not a tooling-enforced check — reviewers
should call out mismatches during PR review.
```

Notes on wording:

- "MUST" (uppercase) signals a process requirement, matching the convention
  used elsewhere in `CLAUDE.md` (e.g. the "Commits must be signed" line under
  Contribution Notes).
- The paragraph explicitly mentions the five values that are duplicated
  (threshold, status-check string, trigger branches, workflow name, job
  name) so a reviewer scanning the diff can immediately see what is in
  scope.
- The closing sentence states the rule is process-only, not
  tool-enforced — satisfying FR-006 AC-2 ("does NOT promise tooling
  enforcement that does not exist") and matching vision constraint C2.

### Diff Shape

Expected `git diff CLAUDE.md` after the edit:

```
@@ existing tail of CI Contract section @@
   **The exact status-check string to enter in GitHub's branch protection UI is: `ci / unit-tests`**
+
+### Drift Management
+
+The values in the Workflow Summary table above (threshold, status-check
+string, trigger branches, workflow name, job name) are duplicated from
+`.github/workflows/ci.yml`. Any future PR that changes those values in
+`ci.yml` MUST also update the corresponding rows in this section in the
+same PR. This is a process rule, not a tooling-enforced check — reviewers
+should call out mismatches during PR review.
```

Total: 8 lines added, 0 lines removed, 0 lines modified.

---

## Interfaces and Data Models

Not applicable. This is a doc-only change with no code interfaces, schemas,
or data models. No SDK calls, no resource schema fields, no provider
registration changes.

---

## Behavioural Impact

| Surface | Before | After |
|---|---|---|
| CI workflow execution | Runs unit tests + coverage threshold on PRs to `main`/`develop` and pushes to `main`/`develop` | Identical — `ci.yml` is unchanged |
| `go build` / `make build` | Builds the provider | Identical — no source changes |
| `make test` / `make testacc` | Runs unit / acceptance tests | Identical — no test changes |
| `terraform init` / `terraform apply` against the built provider | Provisions resources via DCT | Identical — no resource-schema changes |
| Coverage threshold | `2` (per `ci.yml`) | Identical — `ci.yml` is unchanged |
| Documentation read flow | A contributor reads `## CI Contract` and learns workflow, local repro, threshold playbook, branch protection | A contributor additionally reads the `### Drift Management` paragraph and knows to update `CLAUDE.md` alongside `ci.yml` |

No runtime, build, or test behaviour changes.

---

## Test Plan

See `docs/DLPXECO-14142-test-plan.md` for the per-FR verification matrix.
Summary:

- All verification is structural (does the doc say X?) and diff-based (does
  `git diff` touch only the expected files?). No unit tests, no acceptance
  tests, no new test files.
- The existing CI workflow (`ci / unit-tests`) continues to provide unit-test
  + coverage gating; since this ticket adds no Go code, it does not change
  the coverage total and the threshold (`2`%) remains comfortably satisfied.

---

## Risks and Mitigations

| Risk | Likelihood | Impact | Mitigation |
|---|---|---|---|
| Reviewer asks to remove the new "Drift Management" subsection as "obvious" | Medium | Low | Cite vision constraint C2 in the PR description; the constraint explicitly calls out that the drift rule must be documented, not assumed |
| Reviewer asks to also update `README.md` or `pull_request_template.md` | Medium | Low | Re-affirm vision NG5 in the PR description; offer a follow-up ticket if the team disagrees |
| Reviewer asks to lower the duplication by removing values from the table and pointing to `ci.yml` instead | Low | Medium — would unwind FR-002 | Cite vision G1 + C1: the whole point of the contract doc is copy-paste-able values. The drift note (FR-006) is the explicit cost-management mechanism for the duplication |
| Markdown lint / `mdl` config in the repo flags the new paragraph | Low | Low | None expected — wording follows the existing `CLAUDE.md` style (sentence-case `###` headings, prose paragraphs, no exotic constructs). If `mdl` is wired in, fix in place |
| The existing section already passes all FRs and reviewers ask "why does this ticket need a PR at all?" | Medium | Low | The PR adds the FR-006 paragraph that is genuinely missing from the existing section. Cite the FR-006 row of the gap-analysis table above in the PR description |

---

## Rollback

Trivial: `git revert` the single commit. No data migrations, no schema
versions, no consumer impact.

---

## Out-of-Scope Confirmations (mirror functional spec)

- No edits to `.github/workflows/ci.yml`.
- No edits to `README.md`, `pull_request_template.md`, `.goreleaser.yml`,
  `GNUmakefile`, `main.go`, anything under `internal/`, anything under
  `examples/`, anything under `docs/resources/` or `docs/guides/`.
- No change to the coverage threshold value.
- No new CI jobs, lint configurations, or test files.
- No programmatic configuration of GitHub branch protection.
