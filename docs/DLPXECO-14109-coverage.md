# Spec-Code Coverage: DLPXECO-14109

**Source spec**: `docs/DLPXECO-14109-functional.md`
**Evidence collection method**: `grep -rn` from feature-implement test-phase coverage step. Every PASS row carries a concrete `file:line` citation from grep output. No reasoning-only PASSes.

| FR-ID | Description | Status | Evidence (file:line or "none") |
|-------|-------------|--------|-------------------------------|
| FR-001 | Bump four transitive Go modules (`golang.org/x/crypto`, `golang.org/x/net`, `golang.org/x/sys`, `google.golang.org/grpc`) to patched versions in `go.mod` / `go.sum`. | PASS | `go.mod:48` — `golang.org/x/crypto v0.52.0`; `go.mod:50` — `golang.org/x/sys v0.45.0 // indirect`; `go.mod:53` — `google.golang.org/grpc v1.79.3 // indirect`; `go.mod:60` — `golang.org/x/net v0.55.0 // indirect`. Additional resolution-check evidence in `docs/DLPXECO-14109-test-evidence.md:52` (S1 row). |
| FR-002 | Provider must compile and pass unit tests on the bumped tree without source-level changes to production code (test-only EC-3 rename allowed). | PASS | `internal/provider/provider_test.go:24` — `func TestProvider(t *testing.T)` (PASSED post-bump per evidence `## Functional (primary)` S10); `internal/provider/provider_test.go:30` — `func TestProvider_impl(t *testing.T)` (PASSED post-bump). Full 14-test unit suite PASS recorded at `docs/DLPXECO-14109-test-evidence.md:54` (S7). `go vet ./...` clean (`docs/DLPXECO-14109-test-evidence.md:56` S8). `git diff main -- internal/provider/` empty (S4 / S9). |
| FR-003 | Security Check reports zero of the original 23 CVEs (9 C / 3 H / 9 M / 2 L) on the post-bump tree. | DEFERRED | Validate-phase scope per `docs/DLPXECO-14109-test-plan.md` Scenarios S15–S18. Local preflight `govulncheck ./...` not run in test phase (not installed; not a merge-gate requirement). Acceptance criteria AC-1, AC-2, AC-3, AC-4 will be exercised by the validate phase via internal Security Check pipeline. Test-phase evidence chain (FR-001 module versions confirmed; FR-002 regression clean) is the prerequisite that makes FR-003 verifiable in validate. |

## Notes

- FR-003 is intentionally `DEFERRED` rather than `FAIL` — the test phase per `docs/DLPXECO-14109-test-plan.md` Test Approach does not run the Security Check pipeline; that is reserved for validate. The orchestrator's "after test → before validate" gate requires `docs/DLPXECO-14109-test-evidence.md` to exist (which it now does) and accepts DEFERRED rows on FRs whose primary evidence path is post-test.
- Per NG3, no new `TestXxx` functions were added in this iteration. Coverage for FR-002 relies on the existing unit suite as the regression gate. The `TestProvider` / `TestProvider_impl` citations above are the structural backbone — every additional non-acceptance unit test in `internal/provider/*_test.go` re-walks the schema and would surface a regression if the bump introduced one.
- `golang.org/x/sys` resolved to `v0.45.0` (one minor above the `v0.44.0` test-plan floor). FR-001 acceptance criteria reference patched-version floors, not exact pins; `v0.45.0 ≥ v0.44.0` therefore satisfies AC-1. See `docs/DLPXECO-14109-test-evidence.md` OQ-1.
