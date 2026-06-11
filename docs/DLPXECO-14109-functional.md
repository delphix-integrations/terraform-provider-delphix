# Functional Specification: DLPXECO-14109

**Jira**: [DLPXECO-14109](https://perforce.atlassian.net/browse/DLPXECO-14109)
**Generated from**: Jira ticket "Problem" table + Goals/Success Criteria from `docs/DLPXECO-14109-vision.md`
**Domain**: feature (security remediation — transitive dependency CVE batch)

---

## FR-001: Bump four transitive Go modules to patched versions in `go.mod` / `go.sum`

### Description
Update `go.mod` and `go.sum` so that `golang.org/x/crypto`, `golang.org/x/net`, `google.golang.org/grpc`, and `golang.org/x/sys` resolve at or above their patched versions (`v0.52.0`, `v0.55.0`, `v1.79.3`, `v0.44.0` respectively), clearing the 23 CVEs flagged through `terraform-plugin-sdk/v2 v2.33.0` and `terraform-plugin-log v0.9.0`.

### Input
- Current `go.mod` (Go 1.25, four affected indirect modules at v0.45.0 / v0.47.0 / v1.61.1 / v0.38.0)
- Current `go.sum`
- Required-version table from the Jira ticket
- Local Go toolchain at v1.25+ with network access to `proxy.golang.org`

### Processing
1. From the working tree on branch `vuln-fix`, run targeted `go get` for each module at the required version:
   - `go get golang.org/x/crypto@v0.52.0`
   - `go get golang.org/x/net@v0.55.0`
   - `go get google.golang.org/grpc@v1.79.3`
   - `go get golang.org/x/sys@v0.44.0`
2. Run `go mod tidy` to reconcile `go.sum` and remove any orphaned hashes.
3. Inspect the diff of `go.mod` — confirm that the four target lines moved to the required versions and that no unrelated `require` block changes occurred (transitive floors raised by the bumps are acceptable; record them in the validation doc).
4. Inspect the diff of `go.sum` — confirm that hash entries for the four modules and their dependencies are present.
5. Verify all four modules remain marked `// indirect` (the provider does not import them directly).
6. Stage `go.mod` and `go.sum` only — do not stage any source file under `internal/provider/` unless FR-002 requires it.

### Output
- **Success**: `go.mod` shows exactly:
  - `golang.org/x/crypto v0.52.0 // indirect`
  - `golang.org/x/net v0.55.0 // indirect`
  - `google.golang.org/grpc v1.79.3 // indirect`
  - `golang.org/x/sys v0.44.0 // indirect`
  - `go.sum` regenerated and consistent with the new `go.mod`.
- **Failure (Go toolchain mismatch)**: `go get` reports an unsupported Go version → halt; pin the toolchain explicitly and retry.
- **Failure (`go mod tidy` regression)**: tidy bumps an unrelated module → review, accept only intended changes, revert unrelated bumps via `go mod edit -require=<m@v>`.

### Acceptance Criteria
- [ ] AC-1: After running the four `go get` commands, `grep -E "golang.org/x/(crypto|net|sys)|google.golang.org/grpc" go.mod` returns exactly the four required versions.
- [ ] AC-2: `go mod tidy` runs cleanly and re-running it a second time produces no further diff (idempotency).
- [ ] AC-3: `go list -m all | grep -E "(golang.org/x/(crypto|net|sys)|google.golang.org/grpc)"` resolves to the four required versions — no `=>` indirection, no fork.
- [ ] AC-4: `git diff main -- go.mod` shows only intentional bumps; `git diff main -- go.sum` is non-empty and consistent.

---

## FR-002: Provider must compile and pass unit tests on the bumped tree

### Description
After the dependency bump (FR-001), `make build` and `make test` must both exit 0 on Go 1.25 without any change to provider source code beyond minimal, compile-only adjustments (if a bumped module renamed or removed a symbol consumed transitively in test code).

### Input
- Post-bump `go.mod` / `go.sum`
- All current provider source under `internal/provider/`
- All current test files (`*_test.go`)
- `GNUmakefile` targets `build`, `test`

### Processing
1. Run `make build` on `darwin_arm64` (developer machine). Capture exit code, stdout, stderr.
2. If `make build` fails:
   a. Read the compile error in full.
   b. If the error is in `internal/provider/` provider source: that is a true breaking API in a bumped module — escalate per Risks table (likely needs SDK bump). Do NOT alter provider source to "hide" the break unless the fix is a one-line syntactic rename that obviously preserves semantics.
   c. If the error is in a test file: apply the minimum syntactic fix — e.g. update a renamed import path, swap a removed helper for its replacement. Do not change assertion logic.
3. Run `make test` (unit tests, parallel=4, timeout=30s). Capture exit code and total/passed/failed count.
4. Compare passed-count against the pre-bump baseline. If the count drops, investigate every drop — a silently-skipped test is treated as a failure for this FR.
5. Run `go vet ./...` for static checks.

### Output
- **Success**: `make build` exit 0, `make test` exit 0, `go vet ./...` exit 0, test count unchanged or higher.
- **Failure**: any non-zero exit code → diagnose per step 2 above.
- **Side effect**: if a test-only file needed a syntactic fix, that file is included in the commit; the change is documented in the PR description with the renamed/removed symbol called out.

### Acceptance Criteria
- [ ] AC-1: `make build` exits 0 on the post-bump tree (captured in `docs/DLPXECO-14109-build-output.md` during the build phase).
- [ ] AC-2: `make test` exits 0; passed-count is ≥ pre-bump baseline; failed-count is 0; skipped-count is unchanged.
- [ ] AC-3: `go vet ./...` reports no new issues vs the pre-bump baseline.
- [ ] AC-4: `git diff main -- internal/provider/` is either empty, or contains only minimal compile-fix edits explained line-by-line in the PR.

---

## FR-003: Security Check reports zero of the original 23 CVEs after the bump

### Description
A re-run of the internal Security Check on branch `vuln-fix` (post-bump) must show that none of the 23 originally-flagged CVEs are still applicable to the provider module. New, unrelated findings (if any) are recorded but are out of scope.

### Input
- Post-bump `go.mod` / `go.sum`
- The original Security Check report referenced by the ticket (23 CVEs across the four modules)
- Access to the Security Check tool/pipeline (internal)

### Processing
1. Trigger a Security Check run targeted at the `vuln-fix` branch (or run the local equivalent — `govulncheck ./...` if available — as a fast preflight).
2. Collect the report. For each of the 23 originally-flagged CVE IDs:
   a. Confirm the CVE is no longer reported on the post-bump tree.
   b. If still reported: dig into the report — likely cause is a transitive consumer at a lower-floor version. Apply a targeted `go get` to that consumer or escalate.
3. Record the new report's "Critical / High / Medium / Low" counts and compare against pre-bump.
4. Document any newly-introduced unrelated findings (none expected, but a bump occasionally surfaces a different finding); these are out of scope for this ticket but must be acknowledged.

### Output
- **Success**: post-bump scan reports 0 of the original 23 CVEs as still applicable.
- **Failure**: any of the 23 CVEs persists → diagnose the persistent consumer; either apply an additional targeted bump or escalate as a new ticket.
- **Side effect**: PR description summarizes "23 → 0 of the original batch" and lists any new findings explicitly.

### Acceptance Criteria
- [ ] AC-1: Each of the 23 originally-flagged CVE IDs is absent from the post-bump Security Check report.
- [ ] AC-2: `govulncheck ./...` (or equivalent local preflight) shows no Critical or High findings attributable to the four target modules at the bumped versions.
- [ ] AC-3: PR description documents the before-state (9 C / 3 H / 9 M / 2 L) and the after-state (target: 0 / 0 / 0 / 0 from the original batch).
- [ ] AC-4: Any newly-introduced unrelated finding is logged in `docs/DLPXECO-14109-validation.md` Section 9 (Out-of-Scope Observations).

---

## Quality Rules

| Rule | Description | Enforcement | Status | Evidence |
|------|-------------|-------------|--------|----------|
| No production source change | This is a dependency-only bump. `git diff main -- internal/provider/*.go` is empty, or contains only minimal compile-fix edits explained line-by-line in the PR. | `git diff --stat main..vuln-fix` reviewed in PR; any non-test source change requires explicit reviewer sign-off | _filled by validate_ | _filled by validate_ |
| No new CVEs introduced | Post-bump Security Check shows zero net-new Critical or High findings vs pre-bump baseline | Compare pre- and post-bump Security Check reports; record delta in validation doc | _filled by validate_ | _filled by validate_ |
| Build + unit tests pass | `make build` and `make test` both exit 0 on Go 1.25, with test passed-count ≥ pre-bump baseline | CI-equivalent local runs captured in `docs/DLPXECO-14109-build-output.md` and `docs/DLPXECO-14109-test-evidence.md` | _filled by validate_ | _filled by validate_ |
| go.mod / go.sum consistency | `go mod tidy` is idempotent on the final tree (running it twice produces no diff) | Run `go mod tidy` twice during the build phase; assert second run leaves the tree unchanged | _filled by validate_ | _filled by validate_ |
| Backward compatibility | Provider schema and resource behavior unchanged. No resource is silently removed or relocated, no schema field deprecated, no acceptance-test rewrite required. | `make build` validates schema compiles; `tfplugindocs` output (if regenerated) shows no diff vs main; resource list in `provider.go` `ResourcesMap` is unchanged | _filled by validate_ | _filled by validate_ |
| Indirect-only status preserved | The four modules remain marked `// indirect` in `go.mod` (the provider does not introduce a direct import on them) | grep `^	golang.org/x/(crypto|net|sys)` and `google.golang.org/grpc` in `go.mod` shows `// indirect` on every line | _filled by validate_ | _filled by validate_ |
| No `replace` directives added | The fix works through normal module resolution; no `replace` line is added to `go.mod` | grep `^replace ` in `go.mod` returns no new lines vs `main` | _filled by validate_ | _filled by validate_ |
| Signed commits | All commits on `vuln-fix` are GPG- or SSH-signed per `CLAUDE.md` § Contribution Notes | `git log --show-signature main..vuln-fix` shows valid signatures on every commit | _filled by validate_ | _filled by validate_ |

---

## Edge Cases

- **EC-1 — `go mod tidy` cascades unrelated bumps**: tidy may raise `golang.org/x/text`, `golang.org/x/tools`, or another module beyond what is strictly required by the four targets. Accept transitively-required bumps but explicitly review for any module that does NOT appear in the dependency closure of the four targets; revert those via `go mod edit -require=<m@v>` and re-tidy.
- **EC-2 — Bumped `grpc v1.79.3` requires a newer `golang.org/x/net` floor than `v0.55.0`**: if so, accept the higher net version (still clears CVE), update FR-001 acceptance criteria to record the actual final version, and document the cascade in the PR.
- **EC-3 — Test file references a renamed symbol** (e.g. `grpc.WithInsecure` → `grpc.WithTransportCredentials(insecure.NewCredentials())`): apply the documented one-line rename; do not change assertion logic; flag in PR.
- **EC-4 — `terraform-plugin-sdk/v2 v2.33.0` itself fails to compile against bumped `grpc`**: this is the escalation scenario — halt, open a follow-up ticket to bump the SDK, and re-scope this ticket to only the bumps that compile cleanly with the current SDK.
- **EC-5 — A bumped version is later yanked from proxy.golang.org**: rare, but if it happens during the build phase, fall back to the next-higher minor version that still clears the CVEs and document the substitution.
- **EC-6 — Acceptance tests (`TF_ACC=1 make testacc`) require a live DCT** and can't run as part of the merge gate: acceptance tests are not in the post-gate for this phase, but must be run on demand against staging before the next release tag. Note this in the doc-updates phase.
- **EC-7 — Cross-compile matrix (`make release`) hits a per-OS toolchain issue (e.g. one of the bumped modules pulled in a `cgo`-only path that breaks Windows builds)**: this is rare for these four pure-Go modules but must be verified during the validate phase (E2E section) by running `make release` dry-run.

## Error Scenarios

- **ERR-1 — `go get <module>@<version>` returns 410 / not found**: confirm the version exists upstream (it does — `v0.52.0`, `v0.55.0`, `v1.79.3`, `v0.44.0` are released as of ticket creation). If still failing, check `GOPROXY` / VPN configuration; do not silently downgrade.
- **ERR-2 — `make build` fails inside an indirect module (e.g. `grpc`)** rather than provider code: this signals an upstream incompatibility between two bumped modules at the chosen versions. Investigate the specific compile error; either (a) pick the lowest version of `grpc` that both clears the CVE AND compiles with the SDK, or (b) escalate per EC-4.
- **ERR-3 — `make test` fails with an `http2` or `grpc` runtime panic**: this is a behavior change in the bumped module. Diagnose via `superpowers:systematic-debugging`. Document the root cause; only update tests if the behavior change is intentional upstream and the test was relying on the old behavior.
- **ERR-4 — Security Check tooling unavailable** (VPN / pipeline down): record this in the validation doc and re-run before tagging. Do not declare FR-003 met without evidence.
- **ERR-5 — A new, unrelated CVE surfaces in a newly-bumped indirect dependency** (e.g. `go mod tidy` raises `golang.org/x/text` to a version with its own finding): out of scope for this ticket; record in validation doc Section 9 and open a follow-up.

## Performance Considerations

No runtime performance considerations apply — this is a dependency-version bump only. No new code paths are added; no I/O patterns change. Build time is expected to be unchanged or marginally faster (newer module versions often improve compile time). The only "performance" surface to watch is `make release` cross-compile wall time on CI: a substantial increase (>20%) would warrant investigation but is not expected for these four pure-Go modules.

---
<!-- Cross-reference:
     FR-001 (bump go.mod/go.sum)              → vision G1 (clear 23 CVEs)
     FR-002 (build + tests pass)              → vision G2 (preserve public surface), SC2, SC3, SC4
     FR-003 (security scan clean)             → vision G1 (final proof), SC5
     Quality Rules implement vision Constraints (no direct-dep bump, no replace, signed commits).
     Edge Cases EC-1..EC-7 implement vision Risks (tidy cascade, grpc breakage, TLS default change, etc.).
     FR-IDs defined here will be referenced in tasks-template (Spec References) and validation-template (FR Coverage). -->
