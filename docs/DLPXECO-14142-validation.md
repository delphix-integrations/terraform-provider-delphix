# DLPXECO-14142 — Validation Report

**Phase**: validate
**Date**: 2026-06-07
**Branch**: DLPXECO-14141-ci
**Diff under review**: working-tree change to `CLAUDE.md` (+9 / -0).

---

## 1. Scope Compliance

| Constraint | Spec source | Observed | Verdict |
|---|---|---|---|
| Only `CLAUDE.md` is modified | design.md §"Architecture Changes" | `git diff --stat` lists `CLAUDE.md` and only `CLAUDE.md` outside `docs/` and `.claude/` | PASS |
| `.github/workflows/ci.yml` untouched | FR-007 / vision NG1 | `git diff -- .github/` empty | PASS |
| `README.md`, `pull_request_template.md` untouched | vision NG5 | unchanged | PASS |
| Nothing under `internal/` modified | vision NG4 | unchanged in this ticket's diff | PASS |
| Nothing under `examples/` modified | vision NG4 | unchanged | PASS |
| Coverage threshold value unchanged | vision NG3 | `ci.yml` unchanged → threshold still `2%` | PASS |
| No new CI jobs / lint configs / test files | vision NG4 | none added | PASS |

---

## 2. FR-by-FR Coverage

| FR | Status | Evidence |
|---|---|---|
| FR-001 — `## CI Contract` section exists | PASS | One heading, structure intact (V-1, V-2) |
| FR-002 — Workflow Summary table | PASS | 11 rows, values mirror `ci.yml` HEAD (V-4, V-5, V-6, V-7) |
| FR-003 — Local-reproduction recipe | PASS | All three bash blocks, first command byte-identical to FR-002 (V-8, V-9, V-10) |
| FR-004 — Threshold-update playbook | PASS | 4 ordered steps + guard line (V-11, V-12, V-13) |
| FR-005 — Branch-protection contract | PASS | Descriptive paragraph + 4-row table + bold callout (V-14, V-15, V-16, V-17) |
| FR-006 — Drift-management note | PASS | `### Drift Management` subsection added; enumerates the 5 duplicated values; explicitly process-only (V-18, V-19) |
| FR-007 — `ci.yml` untouched (negative) | PASS | `git diff -- .github/workflows/ci.yml` empty (V-3) |

All 7 FRs satisfied.

---

## 3. Quality-Rule Compliance

| QR | Verdict | Note |
|---|---|---|
| QR-1 — No links to ephemeral infra | PASS | No Jenkins/Buildkite/Actions-run URLs (V-20) |
| QR-2 — Markdown style matches existing `CLAUDE.md` | PASS | `##` section, `###` subsection, pipe tables, ```` ```bash ```` fences (V-21) |
| QR-3 — No behavioural code changes | PASS | `git diff HEAD -- ':!CLAUDE.md' ':!docs/'` returns 0 lines (V-22) |
| QR-4 — Single source of truth in `ci.yml` | PASS | All duplicated values reproducible from `ci.yml` HEAD (V-23) |
| QR-5 — Trigger names both branches explicitly | PASS | No "default branch"/"protected branch" stand-ins (V-24) |

---

## 4. CI Equivalent Verification

Command (matches `ci.yml`):

```bash
go test ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s
```

- Exit: `0`
- `terraform-provider-delphix/internal/provider`: PASS, `2.3%` coverage
- `terraform-provider-delphix` (root): PASS, `0.0%` (expected — no testable statements in `main.go`)
- Aggregate `total:` line: `2.3%`
- Threshold: `2%` (per `ci.yml` env block)
- Margin above threshold: `+0.3 pp` (consistent with the documented baseline)

CI behaviour expectation: **`ci / unit-tests` will report PASS on the PR.**

---

## 5. Diff Inspection

```
 CLAUDE.md | 9 +++++++++
 1 file changed, 9 insertions(+)
```

Hunk:

```diff
@@ -238,3 +238,12 @@ configure GitHub.
 | Who can bypass | No one (recommended) |
 
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

Additive only. Insertion point matches the design ("immediately after the existing
`### Branch Protection` subsection, at the end of the `## CI Contract` section").

---

## 6. Risk Re-Check (from design.md §"Risks and Mitigations")

| Risk | Status post-implementation |
|---|---|
| Reviewer asks to remove `### Drift Management` as "obvious" | Mitigation language ready in PR body (cite vision C2) |
| Reviewer requests `README.md` / `pull_request_template.md` edits | Mitigation language ready (cite vision NG5; offer follow-up ticket) |
| Reviewer requests removing duplication and pointing to `ci.yml` | Mitigation language ready (cite G1 + C1; drift note is the explicit duplication-cost mechanism) |
| Markdown-lint flags the new paragraph | Not present in repo (no `.mdlrc`, no markdown-lint workflow); risk does not materialize |
| Reviewer questions need for PR | Mitigation: gap-analysis table in design.md shows FR-006 was genuinely missing |

No new risks introduced.

---

## 7. Documentation Coverage (Spec → Code Map)

| Spec artefact | Code/doc location |
|---|---|
| FR-001 | `CLAUDE.md` L176 (`## CI Contract` heading) |
| FR-002 | `CLAUDE.md` L181–195 (`### Workflow Summary` + 11-row table) |
| FR-002 AC-3 | `CLAUDE.md` L197–198 (`TF_ACC` exclusion paragraph) |
| FR-003 | `CLAUDE.md` L200–215 (3 bash blocks) |
| FR-004 | `CLAUDE.md` L217–224 (4-step list + guard) |
| FR-005 | `CLAUDE.md` L226–240 (paragraph + table + callout) |
| FR-006 | `CLAUDE.md` L242–249 (`### Drift Management` subsection — **this ticket's net new content**) |
| FR-007 | `git diff -- .github/` empty (negative requirement) |

100% spec coverage.

---

## 8. Open Issues / Warnings

None. No deferred items, no follow-up tickets, no caveats beyond those already captured in the design's "Risks and Mitigations" section.

---

## 9. Pre-Flight Sanity (from test-plan.md)

| Check | Expected | Observed | Verdict |
|---|---|---|---|
| Only `CLAUDE.md` (and `docs/`) touched | only `CLAUDE.md` outside `docs/` and `.claude/` | confirmed via `git diff --stat` | PASS |
| `.github/` empty in diff | 0 lines | 0 | PASS |
| Exactly one `## CI Contract` heading | 1 | 1 | PASS |
| `### Drift Management` subsection added | 1 | 1 | PASS |
| Local CI repro passes above threshold | `≥ 2%` | `2.3%`, exit 0 | PASS |

---

## Overall Verdict: PASS

All 7 FRs satisfied. All 5 QRs satisfied. 24/24 structural checks PASS. CI-equivalent run PASS with coverage `2.3%` above the `2%` threshold. Diff scope strictly limited to `CLAUDE.md` (+9 lines, 0 deletions) and `docs/DLPXECO-14142-*.md` spec/evidence files. `.github/workflows/ci.yml` and all other constrained files are byte-unchanged.

Cleared to raise PR.
