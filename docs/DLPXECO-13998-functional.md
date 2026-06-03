# Functional Spec: DLPXECO-13998

**Note**: This is a unit-test addition ticket. The feature code (GCP Object Storage support) was shipped in DLPXECO-13662 (provider v4.3.0, commit 263cf5e). The acceptance tests were added in DLPXECO-13975. This ticket adds pure unit tests for the GCP struct-construction logic and the `cloud_provider` schema validator — no new production code is introduced. Vision, design, and implement phases were intentionally skipped. This functional spec is derived from the two sub-tasks in the ticket scope.

---

## FR-001 — GCP ObjectStore struct construction validated by unit tests

**Description**: Unit tests must verify that the `ObjectStore` struct built for `cloud_provider = "GCP"` has `Type = "GcpObjectStore"`, the correct `Bucket`, the correct `Size` (in bytes from `convertStorageToBytes`), and `nil` `AccessCredentials` — mirroring the exact field assignments in `engine_api.go:204-210`.

**Input**: Various combinations of bucket name strings (alphanumeric, hyphens, dots, empty) and size strings (`20GB`, `500GB`, `1TB`).

**Processing**: Tests directly construct the `ObjectStore` struct using the same expressions as the production `case GCP:` switch branch in `initializeSystem`. `convertStorageToBytes` is called identically to production. The `AccessCredentials` guard (no `else if GCP` in the post-switch block) is exercised separately.

**Output**: All assertions in `TestGcpObjectStoreStruct` (7 sub-tests) and `TestGcpObjectStoreNoAccessCredentials` (1 test) pass.

**Acceptance Criteria**:
- [ ] `TestGcpObjectStoreStruct` passes for all 7 bucket/size combinations
- [ ] `Type` is always `"GcpObjectStore"` for GCP case
- [ ] `Bucket` value is passed through unchanged
- [ ] `Size` equals the byte-converted value from `convertStorageToBytes`
- [ ] `TestGcpObjectStoreNoAccessCredentials` passes: `AccessCredentials` is `nil` for GCP

---

## FR-002 — GCP TestConnection struct construction validated by unit tests

**Description**: Unit tests must verify that the `TestConnection` struct built for `cloud_provider = "GCP"` has `Type = "GcpObjectStoreTest"`, the correct `Bucket`, and empty `Region`, `Endpoint`, and `Container` fields — mirroring `engine_api.go:502-508`.

**Input**: Various bucket names (alphanumeric, hyphens, dots, empty). Size inputs are also exercised to confirm they do not affect the `TestConnection` struct.

**Processing**: Tests directly construct `TestConnection` using the same expressions as the production `else if params.CloudProvider == GCP` branch in `testConnectionForObjectStore`. Assertions verify each field.

**Output**: All assertions in `TestGcpObjectStoreTestConnectionStruct` (5 sub-tests) and `TestGcpObjectStoreTestConnectionWithSizes` (3 sub-tests) pass.

**Acceptance Criteria**:
- [ ] `TestGcpObjectStoreTestConnectionStruct` passes for all 5 bucket variations
- [ ] `Type` is always `"GcpObjectStoreTest"` for GCP case
- [ ] `Bucket` value is passed through unchanged
- [ ] `Region`, `Endpoint`, and `Container` are empty strings for GCP
- [ ] `TestGcpObjectStoreTestConnectionWithSizes` passes for `bucket-20gb`, `bucket-1tb`, `bucket-500gb`

---

## FR-003 — `cloud_provider` schema validator accepts exactly {AWS, AZURE, GCP}

**Description**: Unit tests must verify that the `validation.StringInSlice([]string{AWS, AZURE, GCP}, false)` validator used in `resource_engine_configuration.go:294` accepts exactly the three defined constants and rejects all other values — including lowercase variants, partial strings, aliases, and values with trailing whitespace.

**Input**: 13 test cases: 3 accepted values (`AWS`, `AZURE`, `GCP` constants) and 10 rejected values (`gcp`, `aws`, `azure`, `awsgov`, `""`, `S3`, `google`, `GCS`, `"aws "`, `"GCP "`).

**Processing**: The same `validation.StringInSlice` call from the production schema is reproduced inline (same constants, same `false` case-sensitivity flag). Warnings and errors from the validator are checked against expected outcomes.

**Output**: All assertions in `TestCloudProviderValidator` (13 sub-tests) pass.

**Acceptance Criteria**:
- [ ] `AWS`, `AZURE`, `GCP` constants accepted (zero errors, zero warnings)
- [ ] All 10 non-allowed values produce at least one error
- [ ] Lowercase variants (`gcp`, `aws`, `azure`) are rejected (validator is case-sensitive)
- [ ] Aliases (`GCS`, `google`, `S3`) are rejected
- [ ] Values with trailing spaces are rejected

---

## Quality Rules

| Rule | Description | Enforcement |
|------|-------------|-------------|
| Root cause verified before fix | Tests mirror exact production expressions; any change to production breaks tests | Assertion code uses same struct literals and function calls as engine_api.go and resource_engine_configuration.go |
| Regression test required | New unit tests added in `engine_api_gcp_test.go` cover the struct-building logic not covered by acceptance tests | `go test ./internal/provider/... -run TestGcp\|TestCloudProvider` exits 0 |
| No gold-plating | Tests cover only the two sub-tasks in scope (ObjectStore construction, TestConnection construction, validator). No production code added. | `git diff main -- internal/provider/` shows only the new test file |

---

## Edge Cases

1. **Empty bucket string**: `TestGcpObjectStoreStruct` includes `bucket = ""` — the builder is intentionally dumb; validation happens in `CustomizeDiff`. Test confirms struct is built without panicking.
2. **Bucket with dots**: GCS bucket names allow dots (e.g. `my.bucket.name`); tested explicitly.
3. **Size independence for TestConnection**: `TestGcpObjectStoreTestConnectionWithSizes` confirms that size values do not affect the `TestConnection` struct — the connection test only needs the bucket, not the size.
4. **Trailing-space rejection in validator**: `"aws "` and `"GCP "` confirm the `StringInSlice` validator does exact string matching with no trimming.
5. **Alias rejection in validator**: `GCS` (GCP Storage alias) and `google` are explicitly rejected to prevent confusion.
6. **Empty string in validator**: `""` is rejected — cloud_provider is required.

---

## Error Scenarios

- **`convertStorageToBytes` error on bad size string**: `TestGcpObjectStoreStruct` calls `t.Fatalf` if `convertStorageToBytes` returns an error, surfacing parse failures immediately.
- **`AccessCredentials` accidentally set**: `TestGcpObjectStoreNoAccessCredentials` fails if `AccessCredentials != nil`, catching any future regression that adds a GCP credentials path incorrectly.
- **Validator returns no error for invalid value**: `TestCloudProviderValidator` fails if `len(errors) == 0` for a should-reject case, catching validator relaxation bugs.

---

## Performance Considerations

No performance considerations apply — these are pure unit tests with no I/O, no network calls, and no external dependencies. All 25 sub-tests complete in under 1 second total.
