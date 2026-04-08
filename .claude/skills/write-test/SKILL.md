---
name: write-test
description: Write acceptance and unit tests for a Delphix Terraform provider resource. Use when the user wants to add, generate, or improve tests for a resource.
disable-model-invocation: true
argument-hint: "resource-name"
---

Write tests for the Delphix Terraform provider resource: `delphix_$ARGUMENTS`

## Pipeline

Do not write any test code until all research steps are complete.

---

### Step 1 — Read the resource implementation

Read `internal/provider/resource_$ARGUMENTS.go` to understand:
- All schema fields (required vs optional, computed, sensitive, ForceNew)
- Which fields are in `updatable$ARGUMENTSKeys` (from `commons.go`) — these get update tests
- The resource type name used in HCL (e.g. `"delphix_$ARGUMENTS"`)
- Any async job polling (determines if tests need longer timeouts)
- Whether the resource has a `tags` field

Also read one reference test for patterns:
- `internal/provider/resource_environment_test.go` — standard pattern

---

### Step 2 — Identify required env vars

For each required schema field that cannot be hardcoded (IDs, hostnames, credentials), define a corresponding env var using the convention `ACC_<RESOURCE>_<FIELD>` (uppercase, e.g. `ACC_ENV_ENGINE_ID`).

List all env vars the test will need. These go into the `PreCheck` function.

---

### Step 3 — Write the test file

Create `internal/provider/resource_$ARGUMENTS_test.go` with the following structure:

#### Package and imports
```go
package provider

import (
    "context"
    "fmt"
    "os"
    "testing"

    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
    "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)
```
Add `"regexp"` only if writing a negative test with `ExpectError`.

#### TF_ACC guard
Every `TestAcc*` function must begin with:
```go
if os.Getenv("TF_ACC") == "" {
    t.Skip("Acceptance tests skipped unless env 'TF_ACC' set")
}
```

#### Required test functions — write ALL of these:

**a. `testAcc$ARGUMENTSPreCheck(t *testing.T, ...vars)`**
- Call `testAccPreCheck(t)` first (checks `DCT_KEY`, `DCT_HOST`)
- `t.Fatal(...)` for each required env var that is empty

**b. `testAccCheck$ARGUMENTSResourceExists(n string, ...) resource.TestCheckFunc`**
- Look up the resource in `s.RootModule().Resources`
- Call the real DCT SDK API to verify the resource exists (use `testAccProvider.Meta().(*apiClient).client`)
- Return an error if the API call fails or the resource ID is missing

**c. `testAccCheck$ARGUMENTSDestroy(s *terraform.State) error`**
- Iterate `s.RootModule().Resources`, skip non-`delphix_$ARGUMENTS` entries
- Call the Get API; expect a 404 HTTP response
- Return an error if status is not 404

**d. Config functions** — each returns a `string` with valid HCL:
- `testAcc$ARGUMENTSConfigBasic(...)` — minimum required fields + a `tags` block if the resource supports tags
- `testAcc$ARGUMENTSConfigUpdate(...)` — same as basic but with one updatable field changed (only if updatable fields exist)

**e. Test functions**

Positive create test (always required):
```go
func TestAcc$ARGUMENTS_positive(t *testing.T) {
    if os.Getenv("TF_ACC") == "" { t.Skip(...) }
    // read env vars
    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAcc$ARGUMENTSPreCheck(t, ...) },
        Providers:    testAccProviders,
        CheckDestroy: testAccCheck$ARGUMENTSDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAcc$ARGUMENTSConfigBasic(...),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheck$ARGUMENTSResourceExists("delphix_$ARGUMENTS.test", ...),
                    resource.TestCheckResourceAttr("delphix_$ARGUMENTS.test", "<key>", "<expected_value>"),
                ),
            },
        },
    })
}
```

Update test (only if updatable fields exist in `updatable$ARGUMENTSKeys`):
- Two steps: first apply basic config, then apply update config
- Check that the updated field has the new value

---

### Step 4 — Build and verify

1. Run `make test` — ensure the new test file compiles with no errors.
2. Report the result. If there are compile errors, fix them before finishing.
3. Remind the user that acceptance tests require `TF_ACC=1` plus the env vars listed in Step 2, and can be run with:
   ```
   TF_ACC=1 go test ./internal/provider -run TestAcc$ARGUMENTS -v -timeout 120m
   ```

---

## Rules

- Never mock the DCT client. `CheckDestroy` and `ResourceExists` must call the real SDK API.
- Never make network calls outside the `TF_ACC` guard.
- Unit tests (no `TF_ACC`) are only appropriate for pure Go logic with no API interaction. If the resource has no such logic, do not write unit tests.
- Config HCL resource label must be `test` (e.g. `resource "delphix_$ARGUMENTS" "test" { ... }`).
- All sensitive fields (passwords, keys) must be passed via env vars — never hardcoded.
- Use `escape()` helper (defined in `resource_environment_test.go`) for any string values that may contain backslashes.
