# Test Plan: DLPXECO-14109

**Jira**: [DLPXECO-14109](https://perforce.atlassian.net/browse/DLPXECO-14109)
**Derived from**: `docs/DLPXECO-14109-design.md` `## Affected Components` and `## Version Compatibility`

<!-- This is the authoritative scenario list for the test-generation phase. Every row in `## Scenarios` becomes a Go test function (`TestXxx`) under `internal/provider/` or a helper invocation in this repo's `make test` target. Where a scenario is a pure dependency-graph / module-resolution check, the test is a shell-level assertion in a *_test.go using `os/exec` is acceptable but the project convention is to do these as scripted CI checks; we mark them clearly so the test-generation phase does not try to invent a Go unit test where a `go list`/`grep` assertion is the only sensible option. -->

---

## Test Approach

Go unit tests via the Terraform Plugin SDK v2 testing framework, executed by `make test` (Go 1.25, parallel=4, timeout=30s as configured in `GNUmakefile`). The merge gate is `make test` only — acceptance tests (`TF_ACC=1 make testacc`) require a live DCT instance and run on-demand pre-release per `.claude/test/testing.md`. Module-resolution and CVE-clearance scenarios are verified by deterministic shell-level checks (`go list -m all`, `grep`, `govulncheck`) captured in the validation doc rather than as new Go test functions, since adding new tests for what is a dependency-only bump would violate vision NG3 ("no new test scenarios beyond those needed to prove still works after bump"). The existing unit-test suite is the regression gate for FR-002.

## Environment / Landscape

- **Landscape**: developer machine (`darwin_arm64` baseline) plus the cross-compile matrix exercised by `make release` dry-run.
- **Service under test**: the compiled provider plugin binary; for the merge gate, no live DCT or Delphix engine is required.
- **Go toolchain**: 1.25 (pinned in `go.mod`); the same toolchain CI uses.
- **VMs to provision**: none. `.claude/test/test-infra.md` is consulted by the test-infra phase, but no `## VMs` entries are needed for a dependency-bump regression — the merge gate is all-local. The acceptance run that happens pre-release (out of merge gate per EC-6) reuses the existing DCT staging fixture and is not part of this plan.
- **Post-bump tooling needed**: `go` 1.25+, optional `govulncheck` (Go vulnerability scanner), plus access to the internal Security Check pipeline for SC5 / FR-003 verification.

## Versions to Cover

<!-- Pulled from docs/DLPXECO-14109-design.md ## Version Compatibility — only "Supported = Yes" rows appear here. The provider has a single Go toolchain target so this table is short by design. -->

| Version | Why | Required? |
|---------|-----|-----------|
| Go 1.25 (toolchain pinned in go.mod) | The build/test gate runs against this toolchain in CI and on developer machines | Yes |
| `terraform-plugin-sdk/v2 v2.33.0` | Direct dep, pinned per NG1; carries the original CVE-tagged transitive imports | Yes |
| `terraform-plugin-log v0.9.0` | Direct dep, pinned per NG1; carries CVE-2026-39824 path referenced by issue #157 | Yes |
| `dct-sdk-go v25.6.0` | Direct dep, pinned; independent dependency closure but exercised by every resource-level test | Yes |
| Cross-compile matrix (goreleaser): freebsd/linux/darwin/windows × amd64/386/arm/arm64 | Confirms no per-OS regression introduced by the bumps (EC-7) | Yes |
| `golang.org/x/crypto v0.52.0`, `x/net v0.55.0`, `grpc v1.79.3`, `x/sys v0.44.0` | The exact set this ticket installs — module-resolution scenarios assert these versions | Yes |

## Scenarios

<!-- Each row maps to at least one FR-* in docs/DLPXECO-14109-functional.md. Negative / failure scenarios are explicit. Scenarios marked [resolution-check] are deterministic shell assertions captured in the validation doc, not new Go test functions — this is consistent with NG3 (no new tests beyond regression). Scenarios marked [regression] re-use the existing test suite; the "test function" column references existing TestXxx names in internal/provider/. -->

| # | Scenario | Maps to FR | Versions | Expected outcome |
|---|----------|-----------|----------|------------------|
| S1 | `go.mod` lists the four target modules at exactly the required versions after `go get` [resolution-check] | FR-001 (AC-1) | Go 1.25 | `grep -E "golang.org/x/(crypto\|net\|sys)\|google.golang.org/grpc" go.mod` returns four lines: `golang.org/x/crypto v0.52.0 // indirect`, `golang.org/x/net v0.55.0 // indirect`, `golang.org/x/sys v0.44.0 // indirect`, `google.golang.org/grpc v1.79.3 // indirect`. |
| S2 | `go mod tidy` is idempotent on the bumped tree [resolution-check] | FR-001 (AC-2) | Go 1.25 | Running `go mod tidy` twice in sequence produces no diff on `go.mod` or `go.sum` after the second run (`git diff go.mod go.sum` is empty). |
| S3 | `go list -m all` resolves the four modules without `replace` indirection [resolution-check] | FR-001 (AC-3), Quality Rule "Indirect-only / no replace" | Go 1.25 | `go list -m all \| grep -E "(golang.org/x/(crypto\|net\|sys)\|google.golang.org/grpc)"` shows the four required versions; no `=>` token appears; `grep "^replace " go.mod` returns no new lines vs `main`. |
| S4 | `git diff main -- go.mod` shows only intentional bumps; `go.sum` is regenerated [resolution-check] | FR-001 (AC-4), Quality Rule "No production source change" | Go 1.25 | Diff of `go.mod` contains exactly the four bumped lines (plus any transitively-required floor bumps reviewed under EC-1); `git diff main -- go.sum` is non-empty and consistent. `git diff main -- internal/provider/` is empty (or only the EC-3 syntactic compile-fix). |
| S5 | `make build` exits 0 on the post-bump tree [regression] | FR-002 (AC-1), Quality Rule "Build + unit tests pass" | Go 1.25, darwin_arm64 baseline | `make build` exits 0; produces `terraform-provider-delphix` binary under `bin/` (or current install dir). Build output is captured in `docs/DLPXECO-14109-build-output.md`. |
| S6 | Cross-compile matrix builds clean via `make release` dry-run [regression / EC-7] | FR-002 (AC-1) extended | Go 1.25 × {freebsd, linux, darwin, windows} × {amd64, 386, arm, arm64} | `make release` dry-run (or `goreleaser release --snapshot --clean`) completes with exit 0; no per-OS link failure attributable to a bumped module. Surfaced in validation doc Section 7. |
| S7 | `make test` exits 0 with passed-count >= pre-bump baseline [regression] | FR-002 (AC-2), Quality Rule "Build + unit tests pass", Quality Rule "Backward compatibility" | Go 1.25 | `make test` exits 0; passed-count == pre-bump baseline (no silently-dropped tests); failed-count == 0; skipped-count unchanged. Captured in `docs/DLPXECO-14109-test-evidence.md` `## Functional (primary)`. |
| S8 | `go vet ./...` reports no new issues vs baseline [regression] | FR-002 (AC-3) | Go 1.25 | `go vet ./...` exits 0; any pre-existing warnings are unchanged in count and content. New warnings traced to a bumped module trigger Open Questions / Risks `R: go vet ...` mitigation. |
| S9 | Provider schema and `ResourcesMap` are byte-stable [regression] | Quality Rule "Backward compatibility" | Go 1.25 | `git diff main -- internal/provider/provider.go` is empty for the `ResourcesMap` block; `tfplugindocs generate` (if regenerated) yields no diff under `docs/resources/` and `docs/index.md`. |
| S10 | TestProvider / provider-level unit tests pass post-bump [regression] | FR-002 (AC-2) | Go 1.25 | Existing `TestProvider` and `TestProvider_impl` in `provider_test.go` continue to pass; provider schema validation walks the same fields with the same `Required`/`Computed`/`Sensitive` flags. |
| S11 | Each resource's unit-test functions pass post-bump [regression] | FR-002 (AC-2), Quality Rule "Backward compatibility" | Go 1.25 | Every `TestXxx` (non-acceptance) function across `resource_*_test.go` keeps its result vs pre-bump baseline. Per `.claude/test/testing.md`, this is `make test` end-to-end with `parallel=4 timeout=30s`. |
| S12 | Negative — `make build` halts with a clear compile error if a renamed grpc symbol surfaces (EC-3) [failure path] | FR-002 (AC-4), EC-3 | Go 1.25 | If a test file imports a removed symbol (e.g. `grpc.WithInsecure`), `make build` exits non-zero with a compile error naming the file and line. The implementer applies the minimum syntactic rename (e.g. to `grpc.WithTransportCredentials(insecure.NewCredentials())`), re-runs `make build`, and gets exit 0. The single-line edit is the only change under `internal/provider/`. |
| S13 | Negative — escalate if the compile error is in non-test provider source (EC-4 / Risk #5) [failure path] | FR-002 (AC-4), EC-4 | Go 1.25 | If `make build` fails in any non-`_test.go` file under `internal/provider/`, the implementer STOPs, files a follow-up "bump terraform-plugin-sdk/v2" ticket, and re-scopes this ticket to only the bumps that compile cleanly. No `replace` directive is added; no silent rollback of the security fix. The validate phase records the escalation in Section 9. |
| S14 | Negative — `go get <m>@<v>` returns 410 / not found (ERR-1) [failure path] | FR-001 | Go 1.25 + proxy.golang.org availability | The implementer confirms the version exists upstream (`v0.52.0`, `v0.55.0`, `v1.79.3`, `v0.44.0` are all released as of ticket creation); if `go get` still fails, the cause is GOPROXY / VPN. The implementer does not silently downgrade; surfaces and re-tries. |
| S15 | Post-bump Security Check reports 0 of the original 23 CVEs as still applicable [validation] | FR-003 (AC-1) | Internal Security Check pipeline | Each of the 23 originally-flagged CVE IDs (9 C / 3 H / 9 M / 2 L) is absent from the post-bump report. Pre/post counts are captured in `docs/DLPXECO-14109-validation.md` Section 1 / Section 9 and the PR description. |
| S16 | `govulncheck ./...` (local preflight) shows no Critical/High attributable to the four bumped modules [validation] | FR-003 (AC-2) | Go 1.25 + govulncheck | `govulncheck ./...` exits 0, or its remaining findings are independent of the four target modules (recorded in validation doc Section 9 as out-of-scope per NG5). |
| S17 | PR description records before-state `9C/3H/9M/2L` and after-state `0/0/0/0` for the original batch [validation] | FR-003 (AC-3) | n/a (doc check) | PR body contains the before/after CVE breakdown table, with links to the original Security Check report and the post-bump scan. |
| S18 | Newly-introduced unrelated CVE finding (if any) is logged in validation doc Section 9 [validation / EC-5] | FR-003 (AC-4) | Internal Security Check pipeline | If `go mod tidy` cascade raised a different module to a version with its own CVE, the finding is documented in Section 9 with a follow-up ticket reference. Out of scope per NG5 — does not block FR-003. |
| S19 | Signed-commit check passes on every commit on `vuln-fix` [process / Quality Rule] | Quality Rule "Signed commits" | n/a | `git log --show-signature main..vuln-fix` shows a valid signature on every commit. The PR is blocked otherwise per `CLAUDE.md` § Contribution Notes. |

## Out of Scope

<!-- Pulled directly from vision Non-Goals + a couple of explicit "we are NOT writing this test" lines to head off reviewer questions. -->

- **NG3 — no new functional tests beyond regression**: we do not add new Go `TestXxx` functions for grpc / x/net / x/crypto / x/sys behaviour. The existing test suite is the regression gate. If a test had to be touched, it is an EC-3 single-line syntactic rename, not a new test.
- **NG1 — no direct-dependency bumps**: we do not test against a bumped `terraform-plugin-sdk/v2`, `terraform-plugin-log`, or `dct-sdk-go`. Those are explicitly pinned for this ticket.
- **NG5 — out-of-scope CVEs**: CVEs in test-only deps, in `dct-sdk-go`, or in `tfplugindocs` / `goreleaser` tooling are not exercised here. If surfaced by Section 9 of the validation doc, they are documented but do not gate FR-003.
- **NG2 — no behavioural test rewrites**: no `resource_*_test.go` file is rewritten to assert "the new module behaves differently." If behaviour drift is observed (e.g. EC-2 HTTP/2 client behaviour), it is documented in release notes and exercised at the next acceptance run, not absorbed into a new unit test here.
- **EC-6 — full acceptance suite (`TF_ACC=1 make testacc`) is out of merge gate**: covered by the on-demand pre-release run against staging. Not blocking for this PR.

## Test Data Requirements

- **No fixture seeding**: this is a dependency-only change. Unit tests run with their existing inline HCL fixtures (the `resource.TestStep.Config` strings in each `*_test.go`). No mock server, no recorded HTTP fixtures, no JSON corpora need to be updated.
- **Pre-bump baseline capture (one-time)**: before running `go get`, the implementer captures the pre-bump `make test` output (test count, pass/fail/skip breakdown) so S7 can compare against it. This baseline is recorded in `docs/DLPXECO-14109-test-evidence.md` `## Versions` / `## Landscape / Environment`.
- **`go.mod` / `go.sum` pre-state**: the validation doc captures `git show main:go.mod` and `git show main:go.sum` (or a `git diff main`) so the dependency delta is fully reproducible.

## Exit Criteria

- All `[regression]` scenarios (S5, S6, S7, S8, S9, S10, S11) PASS on the Go 1.25 toolchain (S6 across the goreleaser cross-compile matrix).
- All `[resolution-check]` scenarios (S1, S2, S3, S4) PASS — module resolution matches the spec exactly.
- All `[validation]` scenarios (S15, S16, S17, S18) are addressed by the validate phase, with verdicts captured in `docs/DLPXECO-14109-validation.md`.
- Negative-path scenarios (S12, S13, S14) are documented as expected escalation paths. If S13 fires, the ticket is re-scoped per EC-4 and the PR does not merge until the follow-up SDK-bump ticket is opened.
- `[process]` scenarios (S19) PASS — every commit signed.
- Smoke suite (existing tests in `internal/provider/*_test.go` excluding any file touched by EC-3) PASSes via `make test`.
- No scenario marked SKIPPED without a documented reason in `docs/DLPXECO-14109-test-evidence.md` Notes column.

---
<!-- Cross-references:
     - Each Scenario row → either drives a deterministic shell check captured in docs/DLPXECO-14109-validation.md (resolution-check / validation rows),
       or re-runs an existing TestXxx in internal/provider/*_test.go (regression rows), or documents an escalation path (failure-path rows).
     - Each FR in docs/DLPXECO-14109-functional.md → at least one scenario here:
         FR-001 → S1, S2, S3, S4, S14
         FR-002 → S5, S6, S7, S8, S10, S11, S12, S13
         FR-003 → S15, S16, S17, S18
         Quality Rules / NGs → S4, S9, S19 + Out of Scope section
     - Versions column → subset of docs/DLPXECO-14109-design.md ## Version Compatibility "Supported = Yes" rows.
     - The test-generation phase will skip generating new Go TestXxx functions for this ticket per NG3 — the existing suite IS the regression gate. Validation: feature-executor.md Phase: test-generation Step 2 reads this file. -->
