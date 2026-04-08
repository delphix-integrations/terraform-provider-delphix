# Logging Rules

## Inside Resource Functions

Use `tflog` with the provider prefix constants from `commons.go`:

```go
tflog.Info(ctx,  DLPX+INFO+"message")
tflog.Warn(ctx,  DLPX+WARN+"message")
tflog.Error(ctx, DLPX+ERROR+"message")
```

Do **not** use `InfoLog`, `WarnLog`, or `ErrorLog` (defined in `logging.go`) inside resource functions — those are for package-level init code only.

## What to Log

- Info: start and end of each CRUD operation, key parameter values (non-sensitive).
- Warn: non-fatal API anomalies, skipped optional steps.
- Error: API failures, unexpected state, job failures. Always include the error message or job ID.

## Sensitive Data

Never log credential fields, API keys, passwords, or any value marked `Sensitive: true` in the schema.
