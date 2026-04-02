---
name: validate-resource
description: Validate an existing Delphix Terraform provider resource against provider patterns. Checks CRUD correctness, commons.go entries, error handling, job polling, and test coverage.
argument-hint: "resource-name"
---

Validate the resource `delphix_$ARGUMENTS` in this Terraform provider.

Use the `resource-reviewer` agent to perform a thorough review of:

1. `internal/provider/resource_$ARGUMENTS.go` — CRUD implementation
2. `internal/provider/commons.go` — `updatable$ARGUMENTSKeys` and `isDestructive$ARGUMENTSUpdate` maps (use CamelCase of the resource name)
3. `internal/provider/provider.go` — ResourcesMap registration
4. `internal/provider/resource_$ARGUMENTS_test.go` — test coverage

Report all issues grouped by severity (CRITICAL / WARNING / INFO) and give a final verdict.
