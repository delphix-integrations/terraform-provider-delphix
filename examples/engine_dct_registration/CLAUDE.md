# examples/engine_dct_registration/

Example for the `delphix_engine_dct_registration` resource — registering a Delphix Continuous Data (CD) or Continuous Compliance (CC) engine with DCT.

## Files

| File | Purpose |
|---|---|
| [main.tf](main.tf) | Two examples: CD engine registration and CC engine registration |

## Engine Types

```hcl
# Continuous Data engine
resource "delphix_engine_dct_registration" "cd_engine" {
  hostname     = "engine.example.com"
  name         = "prod-cd-engine"
  username     = "admin"
  password     = "secret"
  insecure_ssl = false
  engine_type  = "CD"
}

# Continuous Compliance engine
resource "delphix_engine_dct_registration" "cc_engine" {
  hostname            = "cc-engine.example.com"
  name                = "prod-cc-engine"
  username            = "admin"
  password            = "secret"
  insecure_ssl        = false
  engine_type         = "CC"
  compliance_user     = "admin"
  compliance_password = "secret"
}
```

## Notes

- `insecure_ssl = true` skips TLS certificate verification — only use in dev/test environments.
- CC engines require additional `compliance_user` / `compliance_password` credentials.
- `engine_type` must be exactly `"CD"` or `"CC"`.
- Available since provider v4.2.0.
- After registration, the engine ID returned by DCT can be used as `engine_id` in `delphix_environment` resources.
