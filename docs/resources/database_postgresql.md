# Resource: <resource name> delphix_database_postgresql

In the Delphix Platform, a Database (or Source Config) identifies the environment's location from which a specific source dataset can be ingested from. dSources are then created from these Databases.

## System Requirements

* Data Control Tower v14.0.1+ is required for database management. Lower versions are not supported.
* The Database PostgreSQL resource only supports the Delphix PostgreSQL database type (POSTGRES). This resource does not support Oracle, SQL Server, or SAP ASE.

## Example Usage

```hcl
# Create a postgres database/source. 

resource "delphix_database_postgresql" "source" {
  name             = "test"
  repository_value = "REPO-1"
  engine_value = "2"
  environment_value = "ENV-1"
}

```

## Argument Reference

* `name` - (Required) The name of the new source.

* `repository_value` - (Required)  The Id or Name of the Repository onto which the source will be created..

* `environment_value` - The Id or Name of the environment to create the source on.

* `engine_value` - The Id or Name of the engine to create the source on.

* `id` - The Source object entity ID.

* `database_type` - The type of this source database.

* `namespace_id` - The namespace id of this source database.

* `namespace_name` - The namespace name of this source database.

* `is_replica` - Is this a replicated object.

* `database_version` - The version of this source database.

* `data_uuid` - A universal ID that uniquely identifies this source database.

* `ip_address` - The IP address of the source's host.

* `fqdn` - The FQDN of the source's host.

* `size` - The total size of this source database, in bytes.

* `jdbc_connection_string` - The JDBC connection URL for this source database.

* `plugin_version` - The version of the plugin associated with this source database.

* `toolkit_id` - The ID of the toolkit associated with this source database(AppData only).

* `is_dsource` - Is this associated with dSource.

* `repository` - The repository id for this source.

* `appdata_source_type` - The type of this appdata source database (Appdata Only).

* `tags` -  The tags to be created for database. This is a map of 2 parameters:
    * `key` - Key of the tag
    * `value` - Value of the tag

## Import (Beta)

Use the [`import` block](https://developer.hashicorp.com/terraform/language/import) to add source configs created directly in Data Control Tower into a Terraform state file. 

For example:
```terraform
import { 
    to = delphix_database_postgresql.source_config_import
    id = "source_config_id"
}
```

*This is a beta feature. Delphix offers no guarantees of support or compatibility.*

