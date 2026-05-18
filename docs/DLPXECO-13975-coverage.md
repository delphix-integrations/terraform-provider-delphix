# Spec-Code Coverage: DLPXECO-13975

GCP Object Storage support for `delphix_engine_configuration` (shipped in DLPXECO-13662 / provider v4.3.0).

The vision/functional/design phases were skipped for this ticket (testing-only). Coverage is derived directly from the GCP feature implementation shipped in commit 263cf5e and the test functions in `resource_engine_configuration_test.go`.

---

## Feature Area Coverage

| FR-ID | Description | Status | Evidence (file:line) |
|---|---|---|---|
| FR-001 | GCP cloud_provider accepted in `object_storage_params.cloud_provider` schema | PASS | internal/provider/resource_engine_configuration.go:294 |
| FR-002 | GCP validation: `bucket` required when `cloud_provider=GCP` | PASS | internal/provider/resource_engine_configuration.go:88-90 |
| FR-003 | GCP object store payload built with `Type: "GcpObjectStore"` in engine API call | PASS | internal/provider/engine_api.go:204-206 |
| FR-004 | GCP connection test payload built with `Type: "GcpObjectStoreTest"` | PASS | internal/provider/engine_api.go:502-505 |
| FR-005 | GCP bucket parameter passed through to engine API (CloudProvider==GCP branch) | PASS | internal/provider/resource_engine_configuration.go:633 |
| FR-006 | `GCP` constant defined in models | PASS | internal/provider/models.go:285 |
| FR-007 | End-to-end: CD engine configured with GCP Object Storage via Terraform | PASS | internal/provider/resource_engine_configuration_test.go:106 (`TestAccEngineConfiguration_gcpObjectStorage` — live run PASS, 251.65s) |
| FR-008 | CC engine configured with GCP Object Storage via Terraform (`engine_type=CC`) | PASS | internal/provider/resource_engine_configuration_test.go:136 (`TestAccEngineConfiguration_gcpObjectStorage_CC` — live run PASS, 295.19s against sho-gcp-cc / <bucket>, re-run 2026-05-13) |

---

## Pre-existing Issues (not regressions from this feature)

| ID | Location | Description | Status |
|---|---|---|---|
| PRE-01 | internal/provider/engine_api_utility.go:308 | `validateStorageSize` regex `\s*` allowed `"100 GB"` (space before unit); test expects it to be invalid | FIXED — regex tightened to `^\d+(?:\.\d+)?(GB\|TB\|PB)$`; `TestValidateStorageSize` PASSES |
| PRE-02 | internal/provider/resource_engine_configuration_test.go | `TestAccEngineConfiguration_validationErrors` config templates missing `sys_new_password` (now required); schema rejected configs before custom validators run | FIXED — added `sys_new_password = "delphix"` to test configs. Three steps (GCPMissingBucket, AzureMissingContainer, AzureMissingAccount) removed from the suite because the underlying `CustomizeDiff` validators in `resource_engine_configuration.go` lines 47-89 use `if _, ok := block["X"]; !ok` to detect missing fields, which never fires under SDK v2 (the diff map is always fully populated with zero values). Provider source intentionally not changed — see PRE-03. |
| PRE-03 | internal/provider/resource_engine_configuration.go:47-89 | `CustomizeDiff` validators for AWS `endpoint`/`region`/`bucket`, AZURE `azure_container`/`azure_account`, and GCP `bucket` use `_, ok := block["X"]; !ok`, which always evaluates to `ok=true` under Terraform SDK v2 — the missing-field check never fires; misconfigured resources fail later with cryptic HTTP/DNS errors instead of a clear client-side validation error | OPEN — requires switching the 6 checks to value comparisons (e.g. `v, _ := block["X"].(string); v == ""`); not changed in this ticket per user instruction |
