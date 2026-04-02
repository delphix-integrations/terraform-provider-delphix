---
name: new-resource
description: Scaffold a new Terraform resource for the Delphix provider. Use when the user wants to add a new resource type to the provider.
disable-model-invocation: true
argument-hint: "resource-name"
---

Scaffold a new Terraform resource for the Delphix provider.

Resource name: $ARGUMENTS

Follow these steps:

1. **Read existing resources** for reference — look at `internal/provider/resource_vdb.go` or a simpler resource like `internal/provider/resource_environment.go` to understand the pattern.

2. **Create the resource file** at `internal/provider/resource_$ARGUMENTS.go` following the Terraform Plugin SDK v2 pattern:
   - Define `Resource$ARGUMENTS()` returning `*schema.Resource`
   - Implement `resourceCreate`, `resourceRead`, `resourceUpdate`, `resourceDelete` functions
   - Use `PollJobStatus()` from `utility.go` for async DCT operations
   - Use `InfoLog`, `WarnLog`, `ErrorLog` from `logging.go`
   - Add updatable fields to `commons.go` maps (`updatable$ARGUMENTSKeys`, `isDestructive$ARGUMENTSUpdate`)

3. **Register the resource** in `internal/provider/provider.go` under `ResourcesMap`

4. **Create a test file** at `internal/provider/resource_$ARGUMENTS_test.go` with at least one acceptance test skeleton

5. **Create example Terraform config** at `examples/resources/delphix_$ARGUMENTS/resource.tf`

Remind the user to:
- Add the DCT SDK method calls for the new resource
- Update docs in `docs/resources/`
