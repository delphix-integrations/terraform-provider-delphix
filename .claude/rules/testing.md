# Testing Rules

## Unit Tests

- No `TF_ACC` guard needed.
- No live DCT instance required.
- Run with `make test` (parallel=4, timeout=30s).
- Naming: `Test<FunctionName>` (e.g., `TestPollJobStatus`).

## Acceptance Tests

- Must be guarded: `if os.Getenv("TF_ACC") == "" { t.Skip(...) }`.
- Require `DCT_HOST` and `DCT_KEY` environment variables plus any resource-specific vars (e.g., `DATASOURCE_ID`, `ENVIRONMENT_ID`).
- Run with `make testacc` or `TF_ACC=1 go test ./internal/provider -run TestAcc<Name> -v -timeout 120m`.
- Naming: `TestAcc<ResourceName>_<scenario>` (e.g., `TestAccVdb_provision_positive`).

## Required Test Coverage for Each Resource

Every resource file must have a corresponding `resource_<name>_test.go` that includes:
- A `PreCheck` function calling `testAccPreCheck(t)` and checking resource-specific env vars.
- A `CheckDestroy` function confirming the resource is gone after `terraform destroy`.
- At least one positive test step with a `Config` template and a `Check` function.

## Do Not Mock the Database / DCT Client

Acceptance tests must hit a real DCT instance. Mocked tests that pass do not guarantee production correctness. Use unit tests for pure logic, acceptance tests for API interactions.
