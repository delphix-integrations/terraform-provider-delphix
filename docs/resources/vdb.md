# Resource: <resource name> delphix_vdb

In Delphix terminology, a virtual database (VDB) is a full read/write copy of the source data. It is created (provisioned) from either a dSource or another VDB's data snapshot.
A VDB is created and managed by the Delphix Continuous Data Engine.

The VDB (delphix_vdb) resource allows Terraform to create, update, and delete Delphix VDBs. This specifically enables the `plan`, `apply`, `update`, and `destroy` Terraform commands. All supported parameters are listed below.

## Example Usage
Provisioning can be done using one of three methods: provision by snapshot, timestamp, and bookmark.

```hcl
# Provision a VDB using latest snapshot.

resource "delphix_vdb" "vdb_name_provision_by_snapshot" {
  auto_select_repository = true
  source_data_id         = "<DATASOURCE_ID_OR_NAME>"
  snapshot_id            = "<SNAPSHOT_ID>" # Leave empty to select the latest snapshot
}

# Provision a VDB from a bookmark and on a Target environment

resource "delphix_vdb" "vdb_name_provision_by_bookmark_on_target_environment" {
  provision_type         = "bookmark"
  auto_select_repository = true
  bookmark_id            = "<BOOKMARK_ID_OR_NAME>"
  environment_id         = "<ENV_ID>"
}

# Provision a VDB using timestamp and configure post refresh hook

resource "delphix_vdb" "vdb_name_provion_by_timestamp_with_hook" {
  provision_type         = "timestamp"
  auto_select_repository = true
  source_data_id         = "<DSOURCE_OR_VDB_ID>"
  timestamp              = "2021-05-01T08:51:34.148000+00:00" # Timestamp must be available on the source dataset.

  post_refresh {
    command         = "echo \"Hello World\""
    name            = "Sample Hook"
    shell           = "SHELL"
  }
}

```

## Argument Reference

* `provision_type` - The type of provisioning to be carried out. Defaults to snapshot. Valid values are `[snapshot, bookmark, timestamp]` 

* `timestamp` - The point in time from which to execute the provision operation. Mutually exclusive with timestamp_in_database_timezone. If the timestamp is not set, selects the latest point.

* `timestamp_in_database_timezone` - The point in time from which to execute the provision operation, expressed as a date-time in the timezone of the source database. Mutually exclusive with timestamp.

* `snapshot_id` - The ID or name of the Snapshot from which to execute the provision operation. If the `snapshot_id` is empty or the paramter is not specified, the latest snapshot is automatically selected.

* `bookmark_id` - The ID or name of the Bookmark from which to execute the provision operation. The Bookmark must contain only one VDB.

* `source_data_id` - The ID or name of the source dataset (dSource, VDB, or Snapshot) to provision from. All other objects referenced by the following parameters must live on the same Continuous Data Engine as the chosen source.

* `engine_id` - The ID or name of the Continuous Data Engine onto which to provision. If the source ID unambiguously identifies a source object, this parameter is unnecessary and ignored.

* `target_group_id` - The ID of the Continuous Data Engine's Dataset Group into which the VDB will be provisioned. If empty, the "Unassigned" Dataset Group is used.

* `name` - [Updatable] The unique name of the VDB. If empty, a name is randomly generated.

* `environment_id` - The ID or name of the Target environment where to provision the VDB. If "repository_id" unambigously identifies a repository, then this value is ignored.

* `environment_user_id` - [Updatable] The environment user ID to use to connect to the Target environment.

* `repository_id` - The ID of the Target environment's repository where to provision the VDB. A repository typically corresponds to a database installation (Oracle home, database instance, etc). Setting this parameter may implicitly determines the environment where to provision the VDB.

* `auto_select_repository` - TRUE or FALSE value to automatically select a compatible environment and repository. Mutually exclusive with "repository_id".

* `pre_refresh` - [Updatable] The commands to execute on the Target environment before refreshing the VDB. This is a map of three parameters:
  * `name` - Name of the hook.
  * `command` - (Required, if hook is specified) Command to be executed.
  * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`.

* `post_refresh` - [Updatable] The commands to execute on the Target environment after refreshing the VDB. This is a map of three parameters:
  * `name` - Name of the hook.
  * `command` - (Required, if hook is specified) Command to be executed.
  * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`.

* `pre_self_refresh` - [Updatable] The commands to execute on the Target environment before a self refresh on the VDB. This is a map of three parameters:
  * `name` - Name of the hook.
  * `command` - (Required, if hook is specified) Command to be executed.
  * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`.

* `post_self_refresh` - [Updatable] The commands to execute on the Target environment after a self refresh on the VDB. This is a map of three parameters:
  * `name` - Name of the hook.
  * `command` - (Required, if hook is specified) Command to be executed.
  * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`.

* `pre_rollback` - (Deprecated) [Updatable] The commands to execute on the Target environment before a rollback on the VDB. This is a map of three parameters:
  * `name` - Name of the hook.
  * `command` - (Required, if hook is specified) Command to be executed.
  * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`.

* `post_rollback` - (Deprecated) [Updatable] The commands to execute on the Target environment after a rollback on the VDB. This is a map of three parameters:
  * `name` - Name of the hook.
  * `command` - (Required, if hook is specified) Command to be executed.
  * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`.

* `configure_clone` - [Updatable] The commands to execute on the Target environment when the VDB is created or refreshed. This is a map of three parameters:
  * `name` - Name of the hook.
  * `command` - (Required, if hook is specified) Command to be executed.
  * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`.

* `pre_snapshot` - [Updatable] The commands to execute on the Target environment before snapshotting a virtual database. These commands can quiesce any data prior to snapshotting. This is a map of five parameters:
  * `name` - Name of the hook.
  * `command` - (Required, if hook is specified) Command to be executed.
  * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`.

* `post_snapshot` - [Updatable] The commands to execute on the Target environment after snapshotting a virtual database. This is a map of three parameters:
  * `name` - Name of the hook.
  * `command` - (Required, if hook is specified) Command to be executed.
  * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`.

* `pre_start` - [Updatable] The commands to execute on the Target environment before starting a virtual database. This is a map of three parameters:
  * `name` - Name of the hook.
  * `command` - (Required, if hook is specified) Command to be executed.
  * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`.

* `post_start` - [Updatable] The commands to execute on the Target environment after starting a virtual database. This is a map of three parameters:
  * `name` - Name of the hook.
  * `command` - (Required, if hook is specified) Command to be executed.
  * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`.

* `pre_stop` - [Updatable] The commands to execute on the Target environment before stopping a virtual database. This is a map of three parameters:
  * `name` - Name of the hook.
  * `command` - (Required, if hook is specified) Command to be executed.
  * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`.

* `post_stop` - [Updatable] The commands to execute on the Target environment after stopping a virtual database. This is a map of three parameters:
  * `name` - Name of the hook.
  * `command` - (Required, if hook is specified) Command to be executed.
  * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`.

* `vdb_restart` - [Updatable] Indicates whether the Continuous Data Engine should automatically restart this virtual database when Target environment reboot is detected.

* `snapshot_policy_id` - The ID of the Snapshot Policy for the VDB.

* `retention_policy_id` - The ID of the Snapshot Retention Policy for the VDB.

* `masked` - TRUE or FALSE boolean to set a VDB as "Masked". Note: You should define a `configure_clone` script in the Hooks step to mask the dataset. The selection of this option will cause the data to be marked as masked, regardless of whether you have defined a script to do so or not.
If you do not define a script to mask the dataset, the data will not be masked unless there is a masking job associated with the dataset.

* `custom_env_vars` - 
Environment variable to be set when a VDB is created. See the Continuous Data ENgine documentation for the list of allowed/denied environment variables and rules about substitution. This is an ordered map of key-value pairs. For eg: { "MY_ENV_VAR1": "$ORACLE_HOME", "MY_ENV_VAR2": "$CRS_HOME/after" }

* `custom_env_files` - Environment files to be sourced when a VDB is created. This path can be followed by parameters. Paths and parameters are separated by spaces. Valid values are a list of env_files. For eg: [ "/export/home/env_file_1", "/export/home/env_file_2" ]

* `tags` - [Updatable] The tags to be created for the VDB. This is a map of two parameters:
  * `key` - (Required) Key of the tag
  * `value` - (Required) Value of the tag

* `make_current_account_owner` - Boolean to determine if the account provisioning this VDB will be the "Owner" of the VDB. 

* `config_params` - [Updatable] The database configuration override parameters.

* `appdata_source_params` - [Updatable] The JSON payload conforming to the DraftV4 schema based on the type of application data being manipulated. These values are unique to each AppData (PostgreSQL, MySQL, etc) connector. Consult the connector documentation for more details.

* `additional_mount_points` - [Updatable] Specifies additional locations on which to mount a subdirectory of an AppData container
  * `shared_path` - (Required) Relative path within the container of the directory that should be mounted.
  * `mount_path` - Absolute path on the target environment were the filesystem should be mounted
  * `environment_id` - The entity ID of the environment on which the file system will be mounted.

* `instance_name` - The VDB's SID name (Oracle Only).

* `open_reset_logs` - TRUE or FALSE value which determines whether to open the database after provision (Oracle Only).

* `online_log_size` - The online log size in MB (Oracle Only).

* `online_log_groups` - The number of online log groups (Oracle Only).

* `archive_log` - TRUE or FALSE boolean to create a VDB in `archivelog` mode (Oracle Only).

* `new_dbid` - [Updatable] TRUE or FALSE boolean to generate a new DB ID for the created VDB (Oracle Only).

* `listener_ids` - [Updatable] The listener IDs for this provision operation. This is a list of listener ids. For eg: [ "listener-123", "listener-456" ]  (Oracle Only).

* `file_mapping_rules` - The VDB file mapping rules (Oracle Only). Rules must be line separated (\n or \r) and each line must have the format "pattern:replacement". Lines are applied in order.

* `unique_name` - The VDB's db_unique_name (Oracle Only).

* `auxiliary_template_id` - The ID of the configuration template to apply to the auxiliary container database (CDB). This is only relevant when provisioning a Multitenant pluggable database into an existing CDB, i.e when the cdb_id property is set. (Oracle Only)

* `cdb_id` - The ID of the container database (CDB) to provision an Oracle Multitenant database into. If empty, a new vCDB will be provisioned. (Oracle only)

* `os_username` - The name of the privileged user to run the provision operation (Oracle only).

* `os_password` - The password of the privileged user to run the provision operation (Oracle only).

* `vcdb_tde_key_identifier` - ID of the key created by the Continuous Data Engine. (Oracle Multitenant Only)

* `cdb_tde_keystore_password` - [Updatable] The password for the Transparent Data Encryption keystore associated with the CDB. (Oracle Multitenant Only)

* `target_vcdb_tde_keystore_path` - [Updatable] Path to the keystore of the vCDB. (Oracle Multitenant Only)

* `tde_key_identifier` - [Updatable] ID of the key created by the Continuous Data Engine. (Oracle Multitenant Only)

* `tde_exported_key_file_secret` - Secret to be used while exporting and importing vPDB encryption keys if Transparent Data Encryption is enabled on the vPDB. (Oracle Multitenant Only)

* `parent_tde_keystore_password` - [Updatable] The password of the keystore specified in parentTdeKeystorePath. (Oracle Multitenant Only)

* `parent_tde_keystore_path` - [Updatable] Path to a copy of the parent's Oracle transparent data encryption keystore on the target host. Required to provision from snapshots containing encrypted database files. (Oracle Multitenant Only)

* `vcdb_name` - The name of the provisioned vCDB when the cdb_id property is not set  (Oracle Multitenant Only).

* `vcdb_database_name` - The database name of the provisioned vCDB wwhen the cdb_id property is not set. Defaults to the value of the vcdb_name property (Oracle Multitenant Only).

* `cluster_node_ids` - The cluster node ids, name, or addresses for this provision operation (Oracle RAC Only).

* `oracle_rac_custom_env_vars` - Environment variable to be set when the engine creates an Oracle RAC VDB. See the Engine documentation for the list of allowed/denied environment variables and rules about substitution.
  * `node_id` - (Required) The node id of the cluster.
  * `name` - (Required) Name of the environment variable
  * `value` - (Required) Value of the environment variable.

* `oracle_rac_custom_env_files` - Environment files to be sourced when the Engine creates an Oracle RAC VDB. This path can be followed by parameters. Paths and parameters are separated by spaces.
  * `node_id` - (Required) The node id of the cluster.
  * `path_parameters` - (Required) This references a file from which certain parameters will be loaded.

* `db_username` - [Updatable] The username of the database (Oracle, SAP ASE only). Only for update.

* `db_password` - [Updatable] The password of the database (Oracle, SAP ASE only). Only for update.

* `template_id` - [Updatable] The ID of the VDB Configuration Template (Oracle, SQL Server Only).

* `database_name` - The name of the database on the Target environment. Defaults to "name" (Oracle, MSSQL, SAP ASE).

* `mount_point` - [Updatable] The mount point for the VDB (Oracle, ASE Only).

* `truncate_log_on_checkpoint` - TRUE or FALSE value to truncate the logs on checkpoints (SAP ASE only).

* `recovery_model` - Recovery model of the source database. Valid values are `[ FULL, SIMPLE, BULK_LOGGED ]`  (SQL Server Only).

* `pre_script` - [Updatable] PowerShell script or executable to run prior to provisioning (SQL Server Only).

* `post_script` - [Updatable] PowerShell script or executable to run after provisioning (SQL Server Only).

* `cdc_on_provision` - [Updatable] Option to enable change data capture (CDC) on both the provisioned VDB and subsequent snapshot-related operations (e.g. refresh, rewind) (SQL Server Only).


## Attribute Reference

* `id` - The VDB object entity ID.

* `database_type` - The database type of this VDB.

* `name` - The container name of this VDB.

* `database_version` - The database version of this VDB.

* `engine_id` - A reference to the Engine that this VDB belongs to. 

* `environment_id` - A reference to the Environment that hosts this VDB.

* `ip_address` - The IP address of the VDB's host.

* `fqdn` - The FQDN of the VDB's host.

* `parent_id` - A reference to the parent dataset of this VDB.

* `group_name` - The name of the group containing this VDB.

* `tags` - A list of key value pair.

* `creation_date` - The date this VDB was created.

## Import (Beta)

Use the [`import` block](https://developer.hashicorp.com/terraform/language/import) to add VDBs created directly in Data Control Tower into a Terraform state file. 

For example:
```terraform
import { 
    to = delphix_vdb.vdb_import_demo
    id = "vdb_id" 
}
```

*This is a beta feature. Delphix offers no guarantees of support or compatibility.*

## Limitations

Not all properties are supported through the `update` command. Properties that are not supported by the `update` command are presented via an error message at runtime.
