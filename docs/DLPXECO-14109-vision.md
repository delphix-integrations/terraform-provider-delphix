# Vision: DLPXECO-14109

**Jira**: [DLPXECO-14109](https://perforce.atlassian.net/browse/DLPXECO-14109) — Security: Remediate 23 transitive dependency vulnerabilities (CVE batch) in terraform provider
**Domain**: feature (security remediation; tracked as `Story` in Jira)
**Source branch**: `vuln-fix` (no worktree; in-place edit)

## Problem Statement

The Security Check scan on `terraform-provider-delphix` flagged **23 vulnerabilities** — 9 Critical, 3 High, 9 Medium, and 2 Low — across four transitive Go modules pulled in via `github.com/hashicorp/terraform-plugin-sdk/v2 v2.33.0` and `github.com/hashicorp/terraform-plugin-log v0.9.0`. Customers running the published provider artifacts (v4.3.0 and earlier) carry these CVEs into every Terraform environment that consumes the provider, exposing them to known auth-bypass, TLS, HTTP/2, and gRPC issues. Until the four upstream modules are bumped to their patched releases, the provider cannot pass internal Security Check gating and cannot be re-published to the Terraform Registry. References: GitHub issue #114 (CVE batch) and #157 (CVE-2026-39824 via terraform-plugin-log).

## Goals

- **G1**: Raise the four transitive modules to the patched versions exactly as required by the security scan — `golang.org/x/crypto` to `v0.52.0`, `golang.org/x/net` to `v0.55.0`, `google.golang.org/grpc` to `v1.79.3`, `golang.org/x/sys` to `v0.44.0` — and clear all 23 flagged CVEs on the next Security Check run.
- **G2**: Preserve the provider's existing public surface — no schema changes, no resource behavior changes, no breaking changes for users of `delphix_vdb`, `delphix_engine_configuration`, or any other resource — verified by `make build`, `make test`, and the existing unit-test suite all passing on Go 1.25.
- **G3**: Keep the bumped versions consistent with `terraform-plugin-sdk/v2 v2.33.0` and `terraform-plugin-log v0.9.0`'s own minimum requirements so that subsequent SDK upgrades do not silently regress these floors.

## Non-Goals

- **NG1**: Do not upgrade `github.com/hashicorp/terraform-plugin-sdk/v2` itself, `github.com/hashicorp/terraform-plugin-log` itself, or any other direct dependency in `go.mod`'s `require` block. Only the four transitive `// indirect` lines are in scope.
- **NG2**: No refactoring of provider source code (`internal/provider/*.go`) — this ticket is dependency-only. Any code change required because of a breaking API in the bumped modules must be the minimum syntactic adjustment, scoped to compile-fix only.
- **NG3**: No new test scenarios beyond those needed to prove "still works after bump" — the existing `make test` suite is the regression gate. No new CVE-specific exploit tests.
- **NG4**: No changes to the release pipeline (`.goreleaser.yml`, `GNUmakefile` version), and no version bump on the provider itself in this ticket (a separate ticket owns the next release).
- **NG5**: Out of scope: vulnerabilities flagged in any test-only dependency, in the `dct-sdk-go` SDK, or in any tool used by `tfplugindocs`/`goreleaser`. Those, if present, are tracked separately.

## Success Criteria

- **SC1**: `go.mod` shows the four indirect modules at exactly the required versions:
  - `golang.org/x/crypto v0.52.0`
  - `golang.org/x/net v0.55.0`
  - `google.golang.org/grpc v1.79.3`
  - `golang.org/x/sys v0.44.0`
  And `go.sum` is regenerated to match.
- **SC2**: `make build` exits 0 on `darwin_arm64` (developer machine) and on the cross-compile matrix (`make release` dry-run completes without errors).
- **SC3**: `make test` exits 0 with no new failures vs the pre-bump baseline; existing test count is preserved (no test is silently skipped due to an API drift).
- **SC4**: `go vet ./...` and `go mod tidy` are both no-ops after the bump (i.e. `go mod tidy` does not further mutate `go.mod`/`go.sum`).
- **SC5**: A fresh Security Check scan on the post-bump branch reports **0** of the 23 originally-flagged CVEs as still applicable. New unrelated findings, if any, are out of scope and noted in the validation doc.
- **SC6**: `go list -m all | grep -E "(golang.org/x/(crypto|net|sys)|google.golang.org/grpc)"` resolves to exactly the four required versions — no `=>` replace directive, no fork.

## Stakeholders

| Stakeholder | Interest |
|-------------|----------|
| Security & compliance team | All 23 CVEs cleared on next scheduled scan; provider passes Security Check gating |
| Terraform Registry consumers (Delphix customers) | Receive a clean, CVE-free provider build on the next published release |
| Delphix ecosystem maintainers | No new compile errors or test regressions to triage post-bump |
| DCT / SDK team | Reassurance that bumping these floors does not block any in-flight SDK or plugin-log upgrades |
| Release manager | Clear "ready to tag" signal — bump lands cleanly with no follow-up code-fix tickets needed |

## Constraints

- **Go toolchain**: must continue to build with `go 1.25` (declared in `go.mod`). The required versions listed above must be compatible with that toolchain — `golang.org/x/net v0.55.0` and `grpc v1.79.3` both require Go 1.23+, which is satisfied.
- **No direct-dependency bumps**: `terraform-plugin-sdk/v2 v2.33.0` and `terraform-plugin-log v0.9.0` stay pinned. The four transitive bumps must be expressible as indirect-only entries in `go.mod` (Go's MVS will record them automatically once they exceed what the SDK requires).
- **No `replace` directives**: the fix must work through normal module resolution. Adding a `replace` line would make the change invisible to downstream consumers of the provider and would not actually clear the CVE.
- **Backward compatibility**: the provider's exported schema, resource names, and acceptance-test behavior must remain unchanged. This is enforced by `make test` and (where available) the acceptance suite.
- **Internal contribution policy**: this is an internal bugfix branch (`vuln-fix`); per `CLAUDE.md`, signed commits and the PR template are required.

## Risks

| Risk | Likelihood | Impact | Mitigation |
|------|------------|--------|------------|
| `grpc v1.79.3` introduces a breaking API change vs `v1.61.1` (18 minor versions) that breaks `terraform-plugin-go` or `terraform-plugin-sdk/v2 v2.33.0` at compile time | Medium | High | Before merging, run `make build` on the bumped tree. If a compile error surfaces in any of the indirect deps, escalate: either pick a closer `grpc` version that still clears the CVE batch, or open a follow-up ticket to upgrade the SDK itself. Do not paper over with `replace`. |
| `golang.org/x/net v0.55.0` ships an HTTP/2 client behavior change that breaks an acceptance test against a live DCT instance | Low | Medium | Acceptance tests require `TF_ACC=1` and a live DCT — they are not part of the merge gate. Surface any net/http2 behavior differences in the validation doc; run `testacc` on demand against staging before tagging the next release. |
| `golang.org/x/crypto v0.52.0` strengthens TLS defaults (e.g. drops a cipher suite) and a customer's DCT endpoint can no longer negotiate | Low | Medium | TLS-related changes between v0.45 and v0.52 are reviewed during validation; documented in the doc-updates phase so release notes call out any cipher-suite floor change. `DCT_TLS_INSECURE_SKIP` remains a documented escape hatch but is explicitly not recommended for production. |
| `go mod tidy` cascades and bumps additional unintended modules (e.g. `golang.org/x/text`, `golang.org/x/tools`) that were not asked for | Medium | Low | Run `go mod tidy` once, diff `go.mod`/`go.sum`, and accept only the four required bumps plus their transitively-required floors. Reject unrelated bumps via targeted `go get`/`go mod edit` if needed. Validation doc captures the final dependency delta. |
| One of the bumped modules deprecates a symbol that `terraform-plugin-sdk/v2 v2.33.0` consumes, producing a compile error in the SDK (not in our code) | Low | High | This is the worst case: it forces a SDK upgrade we explicitly declared out of scope. If hit, halt, open a follow-up "bump terraform-plugin-sdk/v2" ticket, and re-scope this ticket to only the bumps that compile cleanly. Do not silently revert the security fix. |
| Hidden test reliance on a behavior of the old `grpc` (e.g. specific error string in a unit test) breaks `make test` after the bump | Low | Low | Investigate the failing assertion; update the test to use a stable matcher (substring, error type) rather than pinning the bumped module back. Update is local to the test file; no production code change required. |
| Release manager attempts to publish without re-running Security Check post-bump and declares fix premature | Low | Medium | Doc-updates phase explicitly states "re-run Security Check on `vuln-fix` before tagging." Validation verdict requires SC5 evidence. |

---
<!-- Cross-reference: Goals (G1, G2, G3) map to FR-001..FR-003 in the functional spec.
     Success Criteria (SC1..SC6) map to Acceptance Criteria in FR-* entries.
     Constraints and Risks inform the Quality Rules and Edge Cases sections of the functional spec. -->
