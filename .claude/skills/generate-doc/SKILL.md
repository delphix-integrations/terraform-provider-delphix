---
name: generate-doc
description: Generate or update documentation for a Delphix Terraform provider resource. Use when the user wants to create or refresh docs/resources/<name>.md.
disable-model-invocation: true
argument-hint: "resource-name (e.g. vdb, environment, database_plugin)"
---

Generate documentation for the Delphix Terraform resource: $ARGUMENTS

Steps:

1. **Read the resource schema** from `internal/provider/resource_$ARGUMENTS.go`:
   - Collect every schema field: its type, whether Required/Optional/Computed, description, and whether it is marked `[Updatable]` (i.e. present in `updatable$ARGUMENTSKeys` in `commons.go`)
   - Note any `ValidateFunc`, `Default`, `Sensitive`, or nested blocks

2. **Read an existing doc for reference** — use `docs/resources/vdb.md` or the closest equivalent to match tone, section order, and formatting conventions.

3. **Write (or overwrite) `docs/resources/$ARGUMENTS.md`** with the following structure:

   ```
   # Resource: delphix_$ARGUMENTS

   <One-paragraph description of what this resource manages.>

   ## Example Usage

   <Minimal working HCL example using placeholder values in UPPER_SNAKE_CASE.>

   ## Argument Reference

   <Group fields logically (required first, then optional). Mark updatable fields with [Updatable].
    Document nested blocks as sub-sections.>

   ## Timeout Configuration

   <Standard timeouts block explanation — create/update/delete default 20m — only if the resource uses timeouts.>

   ## Import

   <Import block example using the resource ID — only if the resource supports import.>

   ## Limitations

   <Any known constraints: non-updatable fields, unsupported operations, version restrictions.>
   ```

4. Remind the user to:
   - Review generated content for accuracy against the DCT API docs
   - Add version-specific notes if any fields are restricted to certain DCT versions
