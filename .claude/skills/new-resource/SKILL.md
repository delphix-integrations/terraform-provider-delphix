---
name: new-resource
description: Scaffold a new Terraform resource for the Delphix provider. Use when the user wants to add a new resource type to the provider.
disable-model-invocation: true
argument-hint: "resource-name"
---

Scaffold a new Terraform resource for the Delphix provider.

Resource name: $ARGUMENTS

Follow these steps:

1. **Research the DCT SDK** — use the `sdk-researcher` agent to search for all types, API methods, and model fields related to `$ARGUMENTS` in the DCT SDK module cache. This determines whether the resource can be fully implemented or needs TODO stubs.

2. **Read an existing resource** for reference — use `internal/provider/resource_environment.go` as a simple CRUD pattern reference, or `internal/provider/resource_vdb.go` for a complex one.

3. **Create the resource file** at `internal/provider/resource_$ARGUMENTS.go` using the SDK types found in step 1:
   - Define `resource$ARGUMENTS()` returning `*schema.Resource` (use CamelCase)
   - Implement `resourceCreate`, `resourceRead`, `resourceUpdate`, `resourceDelete` functions
   - Wire real SDK calls if SDK support is FULL/PARTIAL; add TODO stubs only where SDK support is NONE
   - Use `PollJobStatus()` from `utility.go` for any operation returning a job ID
   - Use `InfoLog.Println(...)` / `WarnLog.Println(...)` / `ErrorLog.Println(...)` for logging
   - Add updatable fields to `commons.go` maps (`updatable$ARGUMENTSKeys`, `isDestructive$ARGUMENTSUpdate`)

4. **Register the resource** in `internal/provider/provider.go` under `ResourcesMap` as `delphix_$ARGUMENTS`

5. **Create a test file** at `internal/provider/resource_$ARGUMENTS_test.go`:
   - Add `TF_ACC` guard before `resource.Test()` if `Config` function makes live API calls
   - Add `PreCheck` validating required env vars
   - Include at least one positive test step (create + read) and a `CheckDestroy`

6. **Create example Terraform config** at `examples/$ARGUMENTS/main.tf`

7. **Run build and tests** — `make build && make test` — fix any compile errors before finishing.

8. **Run the `resource-reviewer` agent** on the newly created resource file to catch any pattern violations.

Remind the user to update docs in `docs/resources/` when the resource is production-ready.
