---
name: new-resource
description: Scaffold a new Terraform resource for the Delphix provider. Use when the user wants to add a new resource type to the provider.
disable-model-invocation: true
argument-hint: "resource-name"
---

Scaffold a new Terraform resource for the Delphix provider.

Resource name: $ARGUMENTS

## Pipeline

This skill runs a two-agent pipeline before writing any code. Do not skip either step.

---

### Step 1 — SDK Research (sdk-researcher agent)

Run the `sdk-researcher` agent with `$ARGUMENTS` as the concept name.

Wait for the agent to complete and capture its full output. The output must include:
- All available API methods (Create, Read, Update, Delete, List)
- All model fields for CreateParameters, UpdateParameters, and the response struct
- A Terraform schema mapping suggestion
- An availability assessment (FULL / PARTIAL / NONE)

Store this output as **SDK_REPORT**. Do not proceed to Step 2 until SDK_REPORT is complete.

---

### Step 2 — Scaffold (using SDK_REPORT)

Using SDK_REPORT from Step 1 as the authoritative source of truth for all SDK calls and field names:

**a. Read a reference resource**
- Simple pattern: `internal/provider/resource_environment.go`
- Complex pattern: `internal/provider/resource_vdb.go`

**b. Create the resource file** at `internal/provider/resource_$ARGUMENTS.go`:
- Define `resource$ARGUMENTS()` returning `*schema.Resource` (CamelCase name)
- Implement `resourceCreate`, `resourceRead`, `resourceUpdate`, `resourceDelete`
- Use exact SDK method signatures from SDK_REPORT — do not guess or invent method names
- **Implement real API calls** using SDK_REPORT as the authoritative source:
  - Build the request object using the exact CreateParameters/UpdateParameters fields from SDK_REPORT
  - Set every field from the Terraform schema onto the request object before calling the API
  - Map every field from the API response back to `d.Set(...)` in Read
  - Add `// TODO:` stubs only for fields where SDK_REPORT explicitly reports NONE support
- Use `PollJobStatus()` from `utility.go` for any operation returning a job ID
- Use `tflog.Info/Warn/Error(ctx, DLPX+INFO/WARN/ERROR+"...")` for logging inside resource funcs
- Wrap credential fields in `SecureString` and defer `.Clear(ctx)`
- Add `Timeouts` block with 20-minute defaults for Create/Update/Delete
- Add `CustomizeDiff: CustomizeDiffTags` if the resource has a `tags` field

**c. Update `commons.go`** — add two maps:
- `updatable$ARGUMENTSKeys map[string]bool` — use SDK_REPORT Update parameters to populate
- `isDestructive$ARGUMENTSUpdate map[string]bool` — mark fields that require destroy+recreate

**d. Register** in `internal/provider/provider.go` under `ResourcesMap` as `"delphix_$ARGUMENTS"`

**e. Write tests** — invoke the `/write-test` skill with `$ARGUMENTS` as the argument. Wait for it to complete before proceeding. This produces the full test file including `PreCheck`, `CheckDestroy`, `ResourceExists`, Config functions, and all `TestAcc*` cases.

**f. Create example config** at `examples/$ARGUMENTS/main.tf`

---

### Step 3 — Build & Validate

1. Run `make build && make test` — fix any compile errors.
2. Run the `resource-reviewer` agent on the new resource file to catch pattern violations.
3. Report the reviewer's verdict to the user.

Remind the user to update `docs/resources/$ARGUMENTS.md` when the resource is production-ready.
