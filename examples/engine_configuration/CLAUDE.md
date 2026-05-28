# examples/engine_configuration/

Examples for the `delphix_engine_configuration` resource — initial setup of a Delphix Engine: storage, NTP, SMTP, DNS, SSO, and other system-level settings.

## Files

| File | Purpose |
|---|---|
| [main.tf](main.tf) | Comprehensive multi-scenario reference — all storage types and optional system configs |
| [make_and_clean.sh](make_and_clean.sh) | Shell helper to run `terraform apply` and `terraform destroy` in sequence |
| [results/](results/) | Log output from QA infrastructure runs (reference artifacts, not for editing) |

## Storage Scenarios in main.tf

| Scenario | `device_type` | `cloud_provider` | `auth_type` |
|---|---|---|---|
| Block storage (default) | `BLOCK` | — | — |
| AWS Object Storage (role-based) | `OBJECT` | `AWS` | `ROLE` |
| AWS Object Storage (access key) | `OBJECT` | `AWS` | `ACCESS_KEY` |
| Azure Object Storage | `OBJECT` | `AZURE` | `MANAGED_IDENTITIES` |
| GCP Object Storage | `OBJECT` | `GCP` | —  |

## Optional System Config Blocks

```hcl
# NTP
ntp_timezone = "America/New_York"
ntp_servers  = ["pool.ntp.org"]

# SMTP
smtp_config {
  server             = "smtp.example.com"
  port               = 25
  from_email_address = "noreply@example.com"
  tls_authentication = true
}

# DNS
dns_config {
  servers = ["8.8.8.8"]
  domains = ["example.com"]
}

# Web Proxy
web_proxy_config {
  host = "proxy.example.com"
  port = 8080
}

# SSO (SAML)
sso_config {
  enabled                 = true
  response_skew_time      = 120
  max_authentication_age  = 86400
  saml_metadata           = "<YOUR-XML>"
}
```

## Notes

- `engine_host` must be the HTTP/HTTPS address of the engine directly (not DCT).
- `sys_password` + `sys_new_password` are used for the first-time engine setup — after initial config, only `password` matters.
- GCP Object Storage does not require `auth_type` — it uses the engine's service account.
- Available since provider v4.2.0; GCP support added in v4.3.0.
