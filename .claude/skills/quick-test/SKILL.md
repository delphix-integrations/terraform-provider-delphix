---
name: quick-test
description: Use when asked to run, check, or review tests for a specific Delphix Terraform provider resource without the full feature-implement pipeline. Triggers on "quick test X", "run tests for X", "test the X resource", "check tests for X".
---

Interactively run tests for a provider resource. Collects the resource name and required env vars from the user before running anything.

## Step 1 — Ask for the resource name (if not already provided)

If the user did not supply a resource name as an argument, ask:

> Which resource would you like to test? (e.g. `engine_configuration`, `vdb`, `environment`)

Wait for the response. Accept any of these formats and normalize to snake_case:
- kebab-case: `engine-configuration` → `engine_configuration`
- PascalCase: `EngineConfiguration` → `engine_configuration`
- full Terraform name: `delphix_engine_configuration` → `engine_configuration`

## Step 2 — Locate and read the test file

Look for `internal/provider/resource_<name>_test.go`. If it does not exist, stop and tell the user:

> No test file found for `delphix_<name>`. Use `/new-resource` to scaffold one.

Read the full file. Identify all test functions and group them:

| Type | Pattern | Needs live DCT? |
|---|---|---|
| Unit | `func Test<Name>...` (no `Acc` in name) | No |
| Acceptance | `func TestAcc<Name>...` | Yes |

## Step 3 — Discover required env vars

For every acceptance test function (`TestAcc*`):
- Scan its body for all `os.Getenv("VAR_NAME")` calls
- Collect the unique set of variable names across all functions
- Always include `DCT_HOST` and `DCT_KEY` (required by `testAccPreCheck`)

Check which vars are already exported in the current shell:
```bash
echo "DCT_HOST=${DCT_HOST:-<not set>}"
echo "DCT_KEY=${DCT_KEY:-<not set>}"
# repeat for each discovered var
```

## Step 4 — Collect all missing values in one prompt

Build the variable table dynamically from Step 3's scan — every `os.Getenv(...)` call found in the test file must appear as a row. Do not hardcode or assume a fixed list; the set of variables is resource-specific.

**Ask once, do not ask again.** Present the complete table of all required vars (set and missing) in a single message, then wait for the user to paste whatever values they have. Do not ask for variables one-by-one or in multiple rounds.

> To run acceptance tests for `delphix_<name>`, I found **N environment variables** required across all test functions. Please paste the values you have below (any format: `KEY=value`). I'll skip tests whose vars are missing.
>
> | Variable | Found in | Status |
> |---|---|---|
> | `DCT_HOST` | all `TestAcc*` (via `testAccPreCheck`) | ✅ set |
> | `DCT_KEY` | all `TestAcc*` (via `testAccPreCheck`) | ❌ missing |
> | `VAR_A` | `TestAccFoo_scenario1`, `TestAccFoo_scenario2` | ❌ missing |
> | `VAR_B` | `TestAccFoo_scenario3` | ✅ set |
> | *(one row per unique var found in the file)* | | |

After the user responds (or if they provide no values), move on — do not ask again. If a required var is still missing, note in the report that the test was skipped rather than asking the user to supply it.

## Step 5 — Run unit tests immediately

Unit tests require no env vars. Run them right away:

```bash
go test ./internal/provider -run Test<PascalCaseName>[^A][^c][^c] -v -count=1 -timeout 30s
```

(The negative lookahead pattern excludes `TestAcc*` from this run. Alternatively filter by listing the exact unit test names found in step 2.)

Also run the full package unit suite to catch regressions:
```bash
make test
```

## Step 6 — Run acceptance tests (after env vars are confirmed)

Only once all required vars are available, construct and show the command:

```bash
TF_ACC=1 \
  DCT_HOST=<value> DCT_KEY=<value> \
  <OTHER_VAR>=<value> ... \
  go test ./internal/provider \
  -run TestAcc<PascalCaseName> \
  -v -timeout 120m
```

Ask the user: "Ready to run acceptance tests now?" Before executing, confirm the user says yes.

## Step 7 — Report results

| Category | Tests found | Result |
|---|---|---|
| Unit tests | N | ✅ All passed / ❌ N failed |
| Acceptance tests | N | ✅ All passed / ❌ N failed / ⚠️ Not run |

Show full failure output for any failing test.

## Rules

- Never run acceptance tests without the user confirming all env vars are set and saying "yes, run them"
- Never create or modify test files — this skill only runs and reports
- Never log, print, or echo sensitive values like `DCT_KEY` or passwords in plain text
- Always run `make test` before reporting clean unit tests
- If the user asks to *write new* tests, stop and invoke `/write-test` instead
