# examples/vdb/

Examples for the `delphix_vdb` resource — Virtual Database provisioning.

## Top-Level File

[main.tf](main.tf) — Two VDB examples showing `snapshot` and `timestamp` provision types with pre/post refresh hooks.

## Subdirectories by Database Type

Each database type has up to three provision method subdirectories:

| DB Type | snapshot | timestamp | bookmark |
|---|---|---|---|
| [oracle/](oracle/) | yes | yes | yes |
| [postgresql/](postgresql/) | yes | yes | yes |
| [mssql/](mssql/) | yes | yes | yes |
| [mysql/](mysql/) | yes | — | — |
| [hana/](hana/) | yes | yes | yes |
| [sybase/](sybase/) | yes | yes | yes |

## Provision Types

| Type | Required Field | Description |
|---|---|---|
| `snapshot` | `source_data_id` | Provision from the latest available snapshot |
| `timestamp` | `source_data_id`, `timestamp` | Provision to a specific point in time (ISO 8601) |
| `bookmark` | `source_data_id`, `bookmark_id` | Provision from a named bookmark |

## Key Fields

```hcl
resource "delphix_vdb" "example" {
  provision_type         = "snapshot"       # snapshot | timestamp | bookmark
  auto_select_repository = true
  source_data_id         = "<dsource-id>"

  # Optional hooks (can repeat multiple times)
  pre_refresh {
    name    = "my-hook"
    command = "echo hello"
    shell   = "bash"
  }
}
```

## Notes

- `auto_select_repository = true` lets DCT pick the target repository automatically — use this unless you need to target a specific mount point.
- Hooks (`pre_refresh`, `post_refresh`, `pre_rollback`, `post_rollback`, etc.) can be repeated as multiple blocks.
- The `timestamp` field must be ISO 8601 with timezone: `"2021-05-01T08:51:34.148000+00:00"`.
