# Security Rules

## Credential Handling

- Wrap any sensitive string value in `SecureString` and defer `.Clear(ctx)` to wipe it from memory when the function returns.
- For byte slices holding credentials, call `SecureClearByteSlice()` in a deferred call.
- Never log, print, or include sensitive values in error messages.

```go
password := SecureString(d.Get("password").(string))
defer password.Clear(ctx)
```

## Schema

- Mark all password and key fields with `Sensitive: true` in the schema so Terraform redacts them in plan output.

## API Communication

- Default `host_scheme` is `https`. Never default to `http`.
- Honor `tls_insecure_skip` only when explicitly set by the user; do not set it to `true` in code.

## Input Validation

- Validate user-supplied values that are passed to shell commands or external APIs to prevent injection.
- Do not embed user-controlled strings directly in system calls.
