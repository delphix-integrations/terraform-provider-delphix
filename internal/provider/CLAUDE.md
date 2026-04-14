# internal/provider/

Core provider implementation. Everything Terraform calls lives here — resource CRUD handlers, schema definitions, shared utilities, and tests. All files share the `package provider` namespace.

## File Map

### Entry Point
| File | Purpose |
|---|---|
| [provider.go](provider.go) | Registers all resources with Terraform, handles provider-level config (`DCT_KEY`, `DCT_HOST`, etc.) |

### Resources
| File | Terraform Resource |
|---|---|
| [resource_vdb.go](resource_vdb.go) | `delphix_vdb` — VDB provisioning (snapshot/timestamp/bookmark provision types) |
| [resource_vdb_group.go](resource_vdb_group.go) | `delphix_vdb_group` — Group VDBs, manage tags |
| [resource_environment.go](resource_environment.go) | `delphix_environment` — Unix/Windows environments (standalone, cluster, RAC) |
| [resource_appdata_dsource.go](resource_appdata_dsource.go) | `delphix_appdata_dsource` — AppData staged/direct linked sources |
| [resource_oracle_dsource.go](resource_oracle_dsource.go) | `delphix_oracle_dsource` — Oracle RMAN-linked sources |
| [resource_database_postgresql.go](resource_database_postgresql.go) | `delphix_database_postgresql` — PostgreSQL database objects |
| [resource_engine_configuration.go](resource_engine_configuration.go) | `delphix_engine_configuration` — Engine storage (Block/Object: AWS/Azure/GCP), NTP, SMTP, SSO |
| [resource_engine_registration.go](resource_engine_registration.go) | `delphix_engine_dct_registration` — Register CD/CC engines with DCT |
| [resource_engine_plugin_upload.go](resource_engine_plugin_upload.go) | `delphix_database_plugin` — Upload database plugins to engines |

### Shared Infrastructure
| File | Purpose |
|---|---|
| [commons.go](commons.go) | Constants: job states (`PENDING`, `STARTED`, `COMPLETED`, `FAILED`, `CANCELED`), updatable VDB key maps, destructive update flags |
| [models.go](models.go) | Shared Go structs used across multiple resources |
| [utility.go](utility.go) | `PollJobStatus()` async job polling, HTTP response helpers, shared schema helpers |
| [engine_api.go](engine_api.go) | Direct engine API calls (bypasses DCT for engine-level operations) |
| [engine_api_utility.go](engine_api_utility.go) | Utilities specific to engine API calls |
| [logging.go](logging.go) | Logging helpers — always use `tflog` with `[DELPHIX]` prefix |
| [security.go](security.go) | TLS/HTTP client configuration, `DCT_TLS_INSECURE_SKIP` handling |

### Tests
| File | Covers |
|---|---|
| [provider_test.go](provider_test.go) | Provider-level configuration tests |
| [resource_vdb_test.go](resource_vdb_test.go) | VDB acceptance tests |
| [resource_vdb_group_test.go](resource_vdb_group_test.go) | VDB group tests |
| [resource_environment_test.go](resource_environment_test.go) | Environment CRUD tests |
| [resource_appdata_dsource_test.go](resource_appdata_dsource_test.go) | AppData dSource tests |
| [resource_oracle_dsource_test.go](resource_oracle_dsource_test.go) | Oracle dSource tests |
| [resource_database_postgresql_test.go](resource_database_postgresql_test.go) | PostgreSQL database tests |
| [resource_engine_configuration_test.go](resource_engine_configuration_test.go) | Engine configuration tests |
| [resource_engine_registration_test.go](resource_engine_registration_test.go) | Engine registration tests |
| [resource_engine_plugin_upload_test.go](resource_engine_plugin_upload_test.go) | Plugin upload tests |
| [security_test.go](security_test.go) | TLS/security configuration tests |

## Key Patterns

### Async Job Polling
All DCT mutating calls return a job ID. Use `PollJobStatus()` from [utility.go](utility.go) — do not inline your own polling loop.

```go
jobStatus, err := PollJobStatus(jobId, ctx, client)
if jobStatus != COMPLETED {
    return diag.Errorf("Job failed: %s", jobStatus)
}
```

### CRUD Function Signatures
```go
func resourceXxxCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics
func resourceXxxRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics
func resourceXxxUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics
func resourceXxxDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics
```

### Logging
```go
tflog.Info(ctx, "[DELPHIX] Starting VDB creation")
tflog.Warn(ctx, "[DELPHIX] Field X is deprecated")
tflog.Error(ctx, "[DELPHIX] Job failed", map[string]interface{}{"job_id": jobId})
```

### Running Tests
```bash
# Unit tests only
make test

# Acceptance tests (requires live DCT)
TF_ACC=1 DCT_KEY=<key> DCT_HOST=<host> make testacc
```

Acceptance tests run with a 120-minute timeout and 4 parallel workers.
