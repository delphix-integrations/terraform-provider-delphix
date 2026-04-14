# examples/appdata_dsource/

Examples for the `delphix_appdata_dsource` resource — linking AppData (non-Oracle, non-MSSQL) databases as dSources in DCT.

## Subdirectories

| Directory | Database Type | Notes |
|---|---|---|
| [postgresql/](postgresql/) | PostgreSQL | AppData staged source with `delphixInitiatedBackup` parameters |
| [mysql/](mysql/) | MySQL | AppData direct-linked source |
| [import_postg/](import_postg/) | PostgreSQL | `terraform import` workflow for an existing AppData dSource |

## Key Fields

```hcl
resource "delphix_appdata_dsource" "example" {
  source_value               = "<SOURCE-CONFIG-ID>"   # e.g. "1-APPDATA_STAGED_SOURCE_CONFIG-6"
  group_id                   = "<GROUP-ID>"           # e.g. "1-GROUP-1"
  link_type                  = "AppDataStaged"        # AppDataStaged | AppDataDirect
  name                       = "my_dsource"
  log_sync_enabled           = false
  make_current_account_owner = true
  staging_environment        = "<ENV-ID>"
  environment_user           = "<HOST-USER-ID>"
  staging_mount_base         = ""

  # Database-specific JSON parameters
  parameters = jsonencode({
    delphixInitiatedBackupFlag: true,
    delphixInitiatedBackup: [{
      userName: "db_user",
      postgresSourcePort: 5432,
      userPass: "db_pass",
      sourceHostAddress: "source-host"
    }]
  })
}
```

## Notes

- The `parameters` field is a JSON-encoded string — its shape depends entirely on the AppData plugin installed on the engine. Refer to plugin documentation for the correct schema.
- `link_type` determines whether DCT uses a staging server (`AppDataStaged`) or mounts directly from the source (`AppDataDirect`).
- Import is supported since provider v4.0.0.
- The `main.tf` files use commented-out blocks — uncomment and fill in values for your environment.
