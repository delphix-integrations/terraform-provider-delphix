# DLPXECO-14142 — Vision

## Summary

Document the CI contract for the `terraform-provider-delphix` repository so that
contributors, reviewers, and repo maintainers have a single authoritative reference
for **what runs in CI, when it runs, how to reproduce it locally, and what status
check must be required by branch protection**.

This is a **doc-only enhancement**. The CI workflow file (`.github/workflows/ci.yml`)
is already in place and is **not modified** by this work — the source of truth for
behaviour stays in `ci.yml`. This ticket exists to make that behaviour discoverable
and reviewable from `CLAUDE.md` (and downstream from the README / contributor docs).

---

## Problem

`ci.yml` was added by DLPXECO-14115 (predecessor ticket on branch `DLPXECO-14142-ci`)
to run unit tests with coverage enforcement on every PR. The workflow works, but:

1. **Discoverability gap** — a new contributor reading `CLAUDE.md` sees build and
   test commands but no statement of what CI actually runs, on which branches, or
   what status check name a maintainer must wire into branch protection.
2. **Branch-protection ambiguity** — the exact status-check string GitHub expects
   (`ci / unit-tests`, i.e. `<workflow-name> / <job-name>`) is non-obvious and
   easy to mistype. Without it documented, a maintainer configuring branch
   protection can paste the wrong string, silently disabling the gate.
3. **Reproducibility gap** — the test command and coverage-threshold logic live
   only in YAML. A contributor wanting to reproduce CI locally before pushing has
   to read the workflow and reconstruct the command by hand.
4. **Threshold-change ambiguity** — the process for raising/lowering
   `COVERAGE_THRESHOLD` is not written down. Anyone editing it currently has to
   infer the convention (measure → edit → document in PR).

The cost of leaving this undocumented is small per-incident but compounding:
every new contributor pays the discovery tax, and every branch-protection setup
or threshold change risks a quiet mistake.

---

## Goals

1. **G1 — Authoritative CI contract in `CLAUDE.md`.** Add a "CI Contract" section
   to `CLAUDE.md` that enumerates: workflow file path, workflow name, job name,
   status-check string, trigger events, runner, Go version source, test command,
   coverage artifact name + retention, threshold env var location, and the
   baseline coverage measured at the time the threshold was set.
2. **G2 — Local-reproduction recipe.** Document the exact local commands a
   contributor can run to reproduce CI's pass/fail decision (`go test ...
   -coverprofile=...` followed by `go tool cover -func=...`), including the
   per-function and HTML report variants.
3. **G3 — Threshold-update playbook.** Document the four-step process for
   updating `COVERAGE_THRESHOLD` (measure → edit `ci.yml` → document in PR →
   takes effect next run) and the team-agreement rule for lowering it.
4. **G4 — Branch-protection contract.** Document the required branch-protection
   settings for `main` (and call out that a repo maintainer must apply them in
   the GitHub UI — the doc describes the contract, it does not configure
   GitHub). Surface the exact status-check string verbatim so it can be copy-
   pasted into the GitHub UI.
5. **G5 — Dual-branch coverage.** The documented contract must accurately
   describe that CI runs on both `main` **and** `develop` (matching `ci.yml`
   triggers `pull_request: {branches: [main, develop]}` and
   `push: {branches: [main, develop]}`), so the doc does not become stale the
   moment someone opens a PR targeting `develop`.

---

## Non-Goals (NG)

1. **NG1 — Do not modify `.github/workflows/ci.yml`.** The workflow is the source
   of truth for CI behaviour. This ticket only documents what is already there.
   Any behavioural change (new jobs, threshold change, new triggers) is a
   separate ticket.
2. **NG2 — Do not configure GitHub branch protection programmatically.** Branch
   protection is applied by a repo maintainer in the GitHub UI. The doc states
   the required contract; it does not (and cannot, from this repo) enforce it.
3. **NG3 — Do not change the coverage threshold value.** The current value (`2`,
   measured baseline `2.3%`) stays as set by DLPXECO-14115. Raising the floor is
   a separate decision.
4. **NG4 — Do not add new tests, lint configuration, or new CI jobs.** Out of
   scope for a doc-only ticket.
5. **NG5 — Do not modify the README or `pull_request_template.md`.** Those are
   contributor-facing surfaces with their own review owners; updating them is a
   follow-up if/when the team decides `CLAUDE.md` should be cross-referenced.
6. **NG6 — Do not document acceptance tests (`TF_ACC=1`) as part of CI.** They
   are intentionally excluded from CI (no live DCT credentials in runners); the
   doc must state this exclusion explicitly so future contributors don't try to
   "fix" it.

---

## Constraints

1. **C1 — Single source of truth for behaviour.** When the doc describes a value
   (threshold, trigger branches, job name), it must phrase it as "current value
   in `ci.yml`" rather than restating the YAML, OR it must commit to keeping the
   two in sync. We choose the latter for the small set of values that matter
   (threshold, status-check string, trigger branches) because copy-paste-able
   exact values are the whole point of the doc.
2. **C2 — Drift risk acknowledgement.** Because we duplicate a small number of
   values from `ci.yml` into `CLAUDE.md`, any future PR that changes those
   values in `ci.yml` MUST also update `CLAUDE.md` in the same PR. This is a
   process rule, not a tooling enforcement; we accept the residual drift risk
   as the cost of having a copy-paste-able contract doc.
3. **C3 — No external links to mutable infrastructure.** The doc must not link
   to internal Delphix Jenkins/Buildkite/etc. dashboards or to ephemeral GitHub
   Actions run URLs. All references stay inside the repo (`ci.yml`, `go.mod`,
   `GNUmakefile`) or to stable public docs (Terraform Registry, DCT SDK,
   `actions/setup-go`).
4. **C4 — Markdown style consistency with existing `CLAUDE.md`.** Use the same
   heading depth, table style, and code-block fencing as the existing sections
   (Build & Test Commands, Provider Resources, Key Architectural Patterns) so
   the new section reads as a natural continuation, not a bolt-on.
5. **C5 — Dual-branch wording must be unambiguous.** The trigger description
   must say "PRs to `main` or `develop`" and "pushes to `main` or `develop`"
   explicitly — not "PRs to the default branch" or "PRs to the protected
   branch", because the project uses both `main` and `develop` per
   `CONTRIBUTION NOTES` ("branch from `develop` for features, `main` for
   bugfixes").

---

## Risks

| ID | Risk | Likelihood | Impact | Mitigation |
|---|---|---|---|---|
| R1 | Doc and `ci.yml` drift after a future threshold change | Medium | Low — CI still gates correctly; only the doc misleads | Add an explicit note in the threshold-update playbook (G3) instructing the editor to update both files in the same PR. |
| R2 | A maintainer pastes the wrong status-check string into branch protection (e.g. `unit-tests` instead of `ci / unit-tests`) | Low | High — gate silently disabled | Surface the exact string in a dedicated bold callout: "The exact status-check string to enter in GitHub's branch protection UI is: `ci / unit-tests`". |
| R3 | Contributor reads the doc, runs the local command, sees lower coverage than CI (because CI excludes `TF_ACC`) and assumes regression | Low | Low — confusion only | Document explicitly that the baseline (`2.3%`) was measured without `TF_ACC`, matching CI's behaviour. Add a side-note that running with `TF_ACC=1` locally will show a higher number that is not comparable to CI. |
| R4 | `ci.yml` is later renamed or its job is later renamed, invalidating the status-check string | Low | High — same as R2 | Same mitigation as C2: process rule that any rename must update `CLAUDE.md` in the same PR. Cannot be enforced from this ticket. |
| R5 | Scope creep — reviewer asks for README + PR-template updates during review | Medium | Low — delays merge | Reaffirm NG5 in the PR description; offer to file follow-up ticket(s) if the team agrees those surfaces also need the cross-reference. |

---

## Acceptance Criteria (overview — full FR-level criteria in functional.md)

A reviewer should be able to verify the following from the resulting PR diff alone:

- `CLAUDE.md` gains a `## CI Contract` section (or equivalent heading) containing
  all eight items listed in G1.
- The section includes a copy-paste-able local-reproduction command block (G2).
- The section includes the four-step threshold-update playbook (G3).
- The section includes a branch-protection subsection with the exact required
  status-check string surfaced in a bold/callout style (G4, R2).
- The trigger description explicitly names both `main` and `develop` for both
  `pull_request` and `push` (G5, C5).
- `.github/workflows/ci.yml` is **not** modified (NG1) — `git diff` shows
  changes only under `CLAUDE.md` (and any companion spec files under `docs/`).
- No new files are added under `.github/`, `internal/`, or `examples/` (NG4).

---

## Open Questions

None. The decision to keep this doc-only (Option A), to trigger CI on both
`main` and `develop` (matching the already-merged `ci.yml`), and to leave
`ci.yml` untouched is locked per the user's explicit confirmation. Any
deviation requires a new ticket.
