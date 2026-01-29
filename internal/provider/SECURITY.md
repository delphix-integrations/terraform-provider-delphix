# Security - CWE-244 Mitigation

## Overview
This module provides secure memory handling for sensitive data (passwords, credentials, tokens) to prevent heap inspection vulnerabilities (CWE-244: Improper Clearing of Heap Memory Before Release).

## Implementation

### Secure Clearing Functions

1. **SecureClearString(s *string)**
   - Securely clears sensitive string data from memory
   - Overwrites with random data before zeroing
   - Logs clearing operation with debug information
   - Validates successful clearing

2. **SecureClearByteSlice(b []byte)**
   - Securely clears byte slices containing sensitive data
   - Two-pass clearing: random overwrite + zero
   - Verifies all bytes are zeroed
   - Logs operation with hex preview

3. **SecureClearMap(m map[string]interface{}, sensitiveKeys []string)**
   - Clears specific sensitive keys from maps
   - Useful for clearing credentials from configuration maps
   - Tracks number of keys cleared

4. **SecureString type**
   - Wrapper for sensitive strings
   - Automatic cleanup via Clear() method
   - Prevents accidental logging of sensitive data

## Usage

### In Resource Functions

```go
func engineConfigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    // Get sensitive data
    sys_curr_pass := d.Get("sys_password").(string)
    sys_new_pass := d.Get("sys_new_password").(string)
    password := d.Get("password").(string)
    
    // Ensure cleanup on function exit
    defer func() {
        tflog.Info(ctx, "[SECURITY] Clearing sensitive credentials from memory")
        SecureClearString(&sys_curr_pass)
        SecureClearString(&sys_new_pass)
        SecureClearString(&password)
        tflog.Info(ctx, "[SECURITY] All credentials cleared successfully")
    }()
    
    // ... rest of function
}
```

### In API Functions

```go
func login(ctx context.Context, client *http.Client, engine_host string, user string, password string, target string) error {
    loginJSON, err := json.Marshal(loginData)
    if err != nil {
        return err
    }
    
    // Clear sensitive JSON payload after use
    defer func() {
        tflog.Debug(ctx, "[SECURITY] Clearing login credentials from memory")
        SecureClearByteSlice(loginJSON)
        tflog.Debug(ctx, "[SECURITY] Login credentials cleared")
    }()
    
    // ... rest of function
}
```

## Logging

All security operations are logged with appropriate levels:

- **INFO**: High-level security operations (start/end of clearing)
- **DEBUG**: Detailed information including:
  - Length of data being cleared
  - First few characters (obfuscated) for verification
  - Verification that data is actually zeroed

### Example Debug Output

```
[SECURITY] String cleared: original_len=16, prefix_was=MyS***, cleared=true
[SECURITY] ByteSlice cleared: original_len=128, first_bytes_were=a3f2, all_zero=true
[SECURITY] Map cleared: 3 sensitive keys processed
```

## Testing

Run security tests to validate the implementation:

```bash
go test -v ./internal/provider -run TestSecure
```

### Test Coverage

- `TestSecureClearString`: Validates string clearing
- `TestSecureClearByteSlice`: Validates byte slice clearing
- `TestSecureClearMap`: Validates map key clearing
- `TestSecureString`: Validates SecureString wrapper
- `TestSecureClearNilValues`: Tests edge cases (nil values)
- Benchmarks: Performance testing

## Compliance

This implementation addresses:

- **CWE-244**: Improper Clearing of Heap Memory Before Release ('Heap Inspection')
- **OWASP A02:2021**: Cryptographic Failures (sensitive data exposure)
- **PCI DSS Requirement 3**: Protect stored cardholder data

## Best Practices

1. **Always use defer** for cleanup to ensure it happens even on error
2. **Clear ASAP** - clear data as soon as it's no longer needed
3. **Log security operations** - enable audit trail of sensitive data handling
4. **Use TF_LOG=DEBUG** during development to verify clearing works
5. **Never log sensitive values** - only log metadata (length, prefix, etc.)

## Performance Impact

- Minimal overhead (~microseconds per clear operation)
- Random overwrite adds extra security with negligible cost
- Deferred cleanup ensures no memory leaks

## Future Enhancements

- Support for secure credential storage (encrypted memory)
- Integration with hardware security modules (HSM)
- Additional secure data types (SecureInt, SecureMap)
- Memory locking for extremely sensitive data (mlock)
