# Resource: <resource name> delphix_vdb_group

A vdb group is a collection of virtual databases and datesets. vdb_group allows creating such collection for some vdbs.

## Example Usage

```hcl
Creating a vdb group and assigning vdb with vdb_id = my_vdb_id

resource "delphix_vdb_group" "vdb_group_name" {
  name  = "my vdb group"
  vdb_ids = ["my_vdb_id"]
  tags {
    key = "environment"
    value = "production"
  }
  tags {
    key = "project"
    value = "terraform"
  }
}
```

## Argument Reference

* `id` - A unique identifier for the entity.

* `name` - A unique name for the entity.

* `vdb_ids` - The list of VDB IDs in this VDBGroup.

* `tags` - The tags to be created for the VDB group. This is a map of 2 parameters:
  * `key` - The key of the tag.
  * `value` - The value of the tag.

## Attribute Reference

This resource exports same attributes as the arguments.
