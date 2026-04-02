# internal/provider

All provider logic lives here. This is a **Terraform Plugin SDK v2** package that implements every resource and the provider itself.

## File Map

| File | Purpose |
|------|---------|
| `provider.go` | Provider schema, DCT client setup, `configure()` |
| `commons.go` | Job-status constants, updatable/destructive field maps per resource |
| `utility.go` | `PollJobStatus()`, `CustomizeDiffTags()`, shared helpers |
| `models.go` | Shared structs (e.g. hook models) |
| `logging.go` | `InfoLog`, `WarnLog`, `ErrorLog` — `log.Logger` wrappers with `[DELPHIX]` prefix |
| `security.go` | `SecureClearByteSlice()`, `SecureClearString()`, `SecureString` type |
| `engine_api.go` | Direct engine REST calls (session login, SSO) |
| `engine_api_utility.go` | Helpers for engine API: `validateStorageSize()`, request/response processing |
| `resource_vdb.go` | VDB provisioning (most complex resource) |
| `resource_vdb_group.go` | VDB group management |
| `resource_environment.go` | Environment (host) management |
| `resource_appdata_dsource.go` | AppData dSource ingestion |
| `resource_oracle_dsource.go` | Oracle dSource ingestion |
| `resource_database_postgresql.go` | PostgreSQL source registration |
| `resource_engine_configuration.go` | Engine configuration |
| `resource_engine_registration.go` | Register engines with DCT |
| `resource_engine_plugin_upload.go` | Database plugin upload |

## Adding a New Resource

1. Create `resource_<name>.go` — define `Resource<Name>() *schema.Resource` with Create/Read/Update/Delete funcs.
2. Add updatable-field maps to `commons.go`:
   - `updatable<Name>Keys map[string]bool`
   - `isDestructive<Name>Update map[string]bool` (`true` = force-recreate on change)
3. Register in `provider.go` under `ResourcesMap`: `"delphix_<name>": resource<Name>()`.
4. Create `resource_<name>_test.go` with at least one acceptance test skeleton.

## Key Patterns

### Async Operations
All DCT operations return a job ID. Poll it with:
```go
status, errMsg := PollJobStatus(jobId, ctx, client)
if errMsg != "" {
    return diag.Errorf("job failed: %s", errMsg)
}
```
Terminal states: `COMPLETED`, `FAILED`, `TIMEDOUT`, `CANCELED`, `ABANDONED` (constants in `commons.go`).

### Update Strategy
In each resource's Update func:
```go
// Check if the changed field allows in-place update
if !updatable<Name>Keys[key] {
    d.ForceNew(key)
}
// Check if in-place update is destructive (requires recreate)
if isDestructive<Name>Update[key] {
    d.ForceNew(key)
}
```

### Tag Handling
Apply `CustomizeDiffTags` to every resource via `CustomizeDiff`. Resources support `ignore_tag_changes` (bool) to suppress tag drift detection.

### Logging
Use `tflog` with the constants from `commons.go`:
```go
tflog.Info(ctx, DLPX+INFO+"message")
tflog.Warn(ctx, DLPX+WARN+"message")
tflog.Error(ctx, DLPX+ERROR+"message")
```
Do **not** use `InfoLog`/`WarnLog`/`ErrorLog` (from `logging.go`) inside resource funcs — those are for package-level init only.

### Credentials
Wrap any sensitive string with `SecureString` and call `.Clear(ctx)` in a `defer`. Use `SecureClearByteSlice` for byte slices.

## Test Conventions

- Unit tests: no `TF_ACC`, no live DCT needed. Run with `make test`.
- Acceptance tests: set `TF_ACC=1`, require `DCT_HOST` and `DCT_KEY`, and often resource-specific vars (`DATASOURCE_ID`, `ENVIRONMENT_ID`, etc.). Run with `make testacc` or a targeted `go test -run TestAcc<Name>`.
- Test function naming: `TestAcc<ResourceName>_<scenario>` for acceptance tests, `Test<FunctionName>` for unit tests.
