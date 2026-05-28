# Functional Spec: DLPXECO-13975

**Note**: This is a testing-only ticket. The feature code was shipped in DLPXECO-13662 (provider v4.3.0, commit 263cf5e). The vision, design, and implement phases were intentionally skipped. This functional spec is derived retrospectively from the shipped implementation to satisfy the validation phase's coverage gating.

---

## FR-001 — GCP `cloud_provider` accepted in schema

**Description**: The `object_storage_params.cloud_provider` schema field must accept the value `"GCP"` in addition to the existing `"AWS"` and `"AZURE"` values.

**Input**: A `delphix_engine_configuration` Terraform resource with `device_type = "OBJECT"` and `object_storage_params { cloud_provider = "GCP" }`.

**Processing**: The schema's `ValidateFunc` (`validation.StringInSlice`) must include `"GCP"` in its allowed values list. The provider must not reject the configuration at plan time.

**Output**: Terraform plan and apply succeed without validation errors for the `cloud_provider` field.

**Acceptance Criteria**:
- [ ] `cloud_provider = "GCP"` passes schema validation
- [ ] `cloud_provider = "GCP"` is listed alongside `"AWS"` and `"AZURE"` in the `ValidateFunc`

---

## FR-002 — GCP validation: `bucket` required when `cloud_provider = "GCP"`

**Description**: When `cloud_provider` is `"GCP"`, the `bucket` field must be required via `CustomizeDiff`. If `bucket` is absent, the plan must fail with a clear error.

**Input**: A `delphix_engine_configuration` resource with `cloud_provider = "GCP"` and no `bucket` value.

**Processing**: The `CustomizeDiff` function checks for the `bucket` key in the `object_storage_params` block when `cloud_provider == GCP`. If absent, it returns an error.

**Output**: Terraform plan fails with error `"bucket must be provided in object_storage_params for GCP cloud_provider"`.

**Acceptance Criteria**:
- [ ] Plan fails with the expected error message when `bucket` is omitted for GCP
- [ ] Plan succeeds when `bucket` is provided

---

## FR-003 — GCP object store payload built with `Type: "GcpObjectStore"`

**Description**: When calling the engine API to configure object storage with `cloud_provider = "GCP"`, the provider must construct the payload with `Type: "GcpObjectStore"` and include the `Bucket` field.

**Input**: A valid GCP object storage configuration with `bucket` set.

**Processing**: In `engine_api.go`, the `case GCP:` branch constructs an `ObjectStore` struct with `Type: "GcpObjectStore"` and the `Bucket` value from configuration.

**Output**: The engine API receives a payload with `"type": "GcpObjectStore"` and the correct bucket name.

**Acceptance Criteria**:
- [ ] Engine API payload has `type = "GcpObjectStore"` for GCP configurations
- [ ] `Bucket` field is populated in the payload

---

## FR-004 — GCP connection test payload built with `Type: "GcpObjectStoreTest"`

**Description**: When the provider tests the GCP object store connection (prior to committing the configuration), it must send a `TestConnection` payload with `Type: "GcpObjectStoreTest"`.

**Input**: GCP configuration with `cloud_provider = "GCP"` and a valid `bucket`.

**Processing**: In `engine_api.go`, the `else if params.CloudProvider == GCP` branch in the connection test function constructs a `TestConnection` with `Type: "GcpObjectStoreTest"` and `Bucket`.

**Output**: The engine API connection test endpoint receives `{"type": "GcpObjectStoreTest", "bucket": "<bucket_name>"}`.

**Acceptance Criteria**:
- [ ] Connection test payload has `type = "GcpObjectStoreTest"`
- [ ] `Bucket` is populated in the test payload

---

## FR-005 — GCP bucket parameter passed through to engine API

**Description**: The `bucket` value from the Terraform configuration must be extracted from the resource data and passed through to the `ObjectStoreParams` struct for use in both the object store setup and connection test payloads.

**Input**: `object_storage_params[0]["bucket"]` string value in resource data.

**Processing**: In `resource_engine_configuration.go`, the `else if params.CloudProvider == GCP` branch assigns `params.Bucket = object_storage_params[0].(map[string]interface{})["bucket"].(string)`.

**Output**: The `params.Bucket` field is populated and used in downstream API calls.

**Acceptance Criteria**:
- [ ] `bucket` is read from Terraform state and stored in `params.Bucket`
- [ ] The bucket name appears correctly in the engine API payload

---

## FR-006 — `GCP` constant defined

**Description**: A `GCP` string constant must be defined in `models.go` for use throughout the provider codebase, consistent with the existing `AWS` and `AZURE` constants.

**Input**: Codebase references to `GCP` constant.

**Processing**: `models.go` defines `GCP = "GCP"` alongside other cloud provider constants.

**Output**: All references to the GCP cloud provider string use the constant, not a hardcoded string literal.

**Acceptance Criteria**:
- [ ] `GCP = "GCP"` constant exists in `models.go`
- [ ] No hardcoded `"GCP"` string literals in resource or engine API files

---

## FR-007 — End-to-end: CD engine configured with GCP Object Storage via Terraform

**Description**: A complete Terraform lifecycle (plan, apply, read, destroy) must succeed when configuring a Delphix CD engine with GCP Object Storage.

**Input**: A live CD engine VM (GCP cloud), a valid GCP bucket, and valid `delphix_engine_configuration` HCL.

**Processing**: `TestAccEngineConfiguration_gcpObjectStorage` runs the full Terraform acceptance test lifecycle against a live engine.

**Output**: Engine is configured with `device_type=OBJECT`, `cloud_provider=GCP`, the specified bucket, NTP servers, and timezone. All Terraform resource attribute checks pass.

**Acceptance Criteria**:
- [ ] Terraform apply completes without errors
- [ ] `device_type = "OBJECT"` verified in state
- [ ] `object_storage_params.0.cloud_provider = "GCP"` verified in state
- [ ] `object_storage_params.0.bucket = <bucket>` verified in state
- [ ] `object_storage_params.0.size = "20GB"` verified in state
- [ ] NTP servers and timezone verified in state

---

## FR-008 — CC engine configured with GCP Object Storage via Terraform

**Description**: A complete Terraform lifecycle must succeed when configuring a Delphix CC engine with GCP Object Storage. The CC path requires `engine_type = "CONTINUOUS_COMPLIANCE"` and additional compliance credentials.

**Input**: A live CC engine VM (GCP cloud), a valid GCP bucket, and valid HCL with `engine_type = "CONTINUOUS_COMPLIANCE"`.

**Processing**: `TestAccEngineConfiguration_gcpObjectStorage_CC` runs the full acceptance test lifecycle against a live CC engine.

**Output**: CC engine is configured with GCP Object Storage; Terraform resource attribute checks pass including CC-specific fields.

**Acceptance Criteria**:
- [ ] Terraform apply completes without errors on a CC engine
- [ ] `engine_type = "CONTINUOUS_COMPLIANCE"` does not block GCP object storage configuration

---

## Quality Rules

| Rule | Description | Domain | Enforcement |
|------|-------------|--------|-------------|
| API backward compatibility preserved | Existing AWS and Azure object storage configurations must continue to work after GCP support is added | feature | Run `TestAccEngineConfiguration_awsObjectStorage` and `TestAccEngineConfiguration_azureObjectStorage` (or confirm no AWS/Azure test failures) |
| Input validated at point of entry | `bucket` requirement for GCP enforced via `CustomizeDiff` before API calls | feature | Check `CustomizeDiff` logic in `resource_engine_configuration.go` |
| No secrets in log output | Credentials and passwords not logged via `tflog` | feature | Grep for `tflog` calls near sensitive fields |
| Regression test for the new feature | At least one acceptance test for GCP Object Storage | feature | `TestAccEngineConfiguration_gcpObjectStorage` in `resource_engine_configuration_test.go` |
