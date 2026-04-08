# New Resource Workflow

When the user asks to create, add, build, scaffold, implement, write, or make a new Terraform resource — or says things like "I need a resource for X", "can you add support for X", "let's implement delphix_X" — automatically invoke the `/new-resource` skill with the resource concept name as the argument. Do NOT wait for the user to type `/new-resource` explicitly.

## Required Pipeline Order

Never write resource code before completing the SDK research step. The pipeline is:

1. **`sdk-researcher` agent** — find all SDK types, methods, and fields for the resource concept.
2. **Scaffold** — use the SDK research output to write real, correctly-typed Go code.
3. **`resource-reviewer` agent** — validate the scaffolded resource before presenting it to the user.

## Passing Context Between Steps

The `sdk-researcher` output (API methods, model fields, schema mapping, availability assessment) must be held in context and used directly in the scaffold step. Do not ask the user to re-supply SDK information that was already discovered in step 1.

## When SDK Support is NONE

If `sdk-researcher` reports NONE for a resource, stop and tell the user before scaffolding. Show what the closest available SDK API is and ask whether to proceed with TODO stubs or target a different resource.
