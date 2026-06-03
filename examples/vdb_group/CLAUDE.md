# examples/vdb_group/

Example for the `delphix_vdb_group` resource — grouping VDBs together for coordinated operations and tag management.

## Files

| File | Purpose |
|---|---|
| [main.tf](main.tf) | Creates a VDB group containing one or more VDB IDs |

## Usage Pattern

```hcl
resource "delphix_vdb_group" "my_group" {
  name    = "My Group Name"
  vdb_ids = [
    delphix_vdb.vdb1.id,
    delphix_vdb.vdb2.id,
  ]
}
```

## Notes

- `vdb_ids` references must be VDB resource IDs managed in the same Terraform state, or literal DCT VDB IDs.
- VDB groups support import via `terraform import delphix_vdb_group.<name> <group-id>` (available since provider v4.1.0).
- Tags can be added to the group resource for organizing and filtering in DCT.
- `terraform.tfstate` is checked in as a reference artifact — do not treat it as canonical state.
