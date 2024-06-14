# Resource: <resource name> delphix_vdb

In Delphix terminology, a VDB is a database provisioned from either a dSource or another VDB which is a full read/write copy of the source data. 
A VDB is created and managed by the Delphix Continuous Data Engine.

The VDB resource allows Terraform to create, update, and delete Delphix VDBs. This specifically enables the apply and destroy Terraform commands. Update operation does not support all VDB parameters. All supported parameters are listed below.

## Example Usage
Provisioning can be done in 2 methods, provision by snapshot and provision by timestamp.

```hcl
# Provision a VDB using latest snapshot.

resource "delphix_vdb" "vdb_name" {
  auto_select_repository = true
  source_data_id         = "DATASOURCE_ID"
}

# Provision a VDB using timestamp and post refresh hooks

resource "delphix_vdb" "vdb_name2" {
  provision_type         = "timestamp"
  auto_select_repository = true
  source_data_id         = "DATASOURCE_ID"
  timestamp              = "2021-05-01T08:51:34.148000+00:00"

  post_refresh {
    name    = "HOOK_NAME"
    command = "COMMAND"
  }
}

# Provision a VDB from a bookmark with a single VDB

resource "delphix_vdb" "test_vdb" {
  provision_type         = "bookmark"
  auto_select_repository = true
  bookmark_id            = "BOOKMARK_ID"
  environment_id         = "ENV_ID"
}

# Provision a VDB using snapshot and pre refresh hooks

resource "delphix_vdb" "vdb_name" {
  provision_type         = "snapshot"
  auto_select_repository = true
  source_data_id         = "DATASOURCE_ID"

  pre_refresh {
    name    = "HOOK_NAME"
    command = "COMMAND"
  }
}
```

## Argument Reference

* `source_data_id` - (Optional) The ID or name of the source object (dSource or VDB) to provision from. All other objects referenced by the parameters must live on the same engine as the source.

* `engine_id` - (Optional) The ID or name of the Engine onto which to provision. If the source ID unambiguously identifies a source object, this parameter is unnecessary and ignored.

* `target_group_id` - (Optional) The ID of the group into which the VDB will be provisioned. If unset, a group is selected randomly on the Engine.

* `name` - (Optional) The unique name of the provisioned VDB within a group. If unset, a name is randomly generated.

* `database_name` - (Optional) The name of the database on the target environment. Defaults to name.

* `cdb_id` - (Optional) The ID of the container database (CDB) to provision an Oracle Multitenant database into. When this is not set, a new vCDB will be provisioned.

* `cluster_node_ids` - (Optional) The cluster node ids, name or addresses for this provision operation (Oracle RAC Only).

* `truncate_log_on_checkpoint` - (Optional) Whether to truncate log on checkpoint (ASE only).

* `os_username` - (Optional) The name of the privileged user to run the provision operation (Oracle Only).

* `os_password` - (Optional) The password of the privileged user to run the provision operation (Oracle Only).

* `db_username` - (Optional) [Updatable] The username of the database user (Oracle, ASE Only). Only for update.

* `db_password` - (Optional) [Updatable] The password of the database user (Oracle, ASE Only). Only for update.

* `environment_id` - (Optional) The ID or name of the target environment where to provision the VDB. If repository_id unambigously identifies a repository, this is unnecessary and ignored. Otherwise, a compatible repository is randomly selected on the environment.

* `environment_user_id` - (Optional)[Updatable] The environment user ID to use to connect to the target environment.

* `repository_id` - (Optional) The ID of the target repository where to provision the VDB. A repository typically corresponds to a database installation (Oracle home, database instance, ...). Setting this attribute implicitly determines the environment where to provision the VDB.

* `auto_select_repository` - (Optional) Option to automatically select a compatible environment and repository. Mutually exclusive with repository_id.

* `pre_refresh` - (Optional) The commands to execute on the target environment before refreshing the VDB. This is a map of 5 parameters:
  * `name` - Name of the hook
  * `command` - (Required)Command to be executed
  * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]` 
  * `element_id` - Element ID for the hook
  * `has_credentials` - Flag to indicate if it has credentials

* `post_refresh` - (Optional) The commands to execute on the target environment after refreshing the VDB. This is a map of 5 parameters:
  * `name` - Name of the hook
  * `command` - (Required)Command to be executed
  * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`
  * `element_id` - Element ID for the hook
  * `has_credentials` - Flag to indicate if it has credentials

* `pre_rollback` - (Optional) The commands to execute on the target environment before rewinding the VDB. This is a map of 5 parameters:
  * `name` - Name of the hook
  * `command` - (Required)Command to be executed
  * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`
  * `element_id` - Element ID for the hook
  * `has_credentials` - Flag to indicate if it has credentials

* `post_rollback` - (Optional) The commands to execute on the target environment after rewinding the VDB. This is a map of 5 parameters:
  * `name` - Name of the hook
  * `command` - (Required)Command to be executed
  * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`
  * `element_id` - Element ID for the hook
  * `has_credentials` - Flag to indicate if it has credentials

* `configure_clone` - (Optional) The commands to execute on the target environment when the VDB is created or refreshed. This is a map of 5 parameters:
  * `name` - Name of the hook
  * `command` - (Required)Command to be executed
  * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`
  * `element_id` - Element ID for the hook
  * `has_credentials` - Flag to indicate if it has credentials

* `pre_snapshot` - (Optional) The commands to execute on the target environment before snapshotting a virtual source. These commands can quiesce any data prior to snapshotting. This is a map of 5 parameters:
  * `name` - Name of the hook
  * `command` - (Required)Command to be executed
  * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`
  * `element_id` - Element ID for the hook
  * `has_credentials` - Flag to indicate if it has credentials

* `post_snapshot` - (Optional) The commands to execute on the target environment after snapshotting a virtual source. This is a map of 5 parameters:
  * `name` - Name of the hook
  * `command` - (Required)Command to be executed
  * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`
  * `element_id` - Element ID for the hook
  * `has_credentials` - Flag to indicate if it has credentials

* `pre_start` - (Optional) The commands to execute on the target environment before starting a virtual source. This is a map of 5 parameters:
  * `name` - Name of the hook
  * `command` - (Required)Command to be executed
  * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`

* `post_start` - (Optional) The commands to execute on the target environment after starting a virtual source. This is a map of 5 parameters:
  * `name` - Name of the hook
  * `command` - (Required)Command to be executed
  * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`
  * `element_id` - Element ID for the hook
  * `has_credentials` - Flag to indicate if it has credentials

* `pre_stop` - (Optional) The commands to execute on the target environment before stopping a virtual source. This is a map of 5 parameters:
  * `name` - Name of the hook
  * `command` - (Required)Command to be executed
  * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`
  * `element_id` - Element ID for the hook
  * `has_credentials` - Flag to indicate if it has credentials

* `post_stop` - (Optional) The commands to execute on the target environment after stopping a virtual source. This is a map of 5 parameters:
  * `name` - Name of the hook
  * `command` - (Required)Command to be executed
  * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`
  * `element_id` - Element ID for the hook
  * `has_credentials` - Flag to indicate if it has credentials

* `vdb_restart` - (Optional) [Updatable] Indicates whether the Engine should automatically restart this virtual source when target host reboot is detected.

* `auxiliary_template_id` - (Optional) The ID of the configuration template to apply to the auxiliary container database. This is only relevant when provisioning a Multitenant pluggable database into an existing CDB, i.e when the cdb_id property is set. (Oracle Only)

* `template_id` - (Optional) [Updatable] The ID of the target VDB Template (Oracle Only).

* `file_mapping_rules` - (Optional) Target VDB file mapping rules (Oracle Only). Rules must be line separated (\n or \r) and each line must have the format "pattern:replacement". Lines are applied in order.

* `oracle_instance_name` - (Optional) Target VDB SID name (Oracle Only).

* `unique_name` - (Optional) Target VDB db_unique_name (Oracle Only).

* `vcdb_name` - (Optional) When provisioning an Oracle Multitenant vCDB (when the cdb_id property is not set), the name of the provisioned vCDB (Oracle Multitenant Only).

* `vcdb_database_name` - (Optional) When provisioning an Oracle Multitenant vCDB (when the cdb_id property is not set), the database name of the provisioned vCDB. Defaults to the value of the vcdb_name property. (Oracle Multitenant Only).

* `mount_point` - (Optional) Mount point for the VDB (Oracle, ASE Only).

* `open_reset_logs` - (Optional) Whether to open the database after provision (Oracle Only).

* `snapshot_policy_id` - (Optional) The ID of the snapshot policy for the VDB.

* `retention_policy_id` - (Optional) The ID of the retention policy for the VDB.

* `recovery_model` - (Optional) Recovery model of the source database (MSSql Only). Valid values are `[ FULL, SIMPLE, BULK_LOGGED ]`

* `pre_script` - (Optional) [Updatable] PowerShell script or executable to run prior to provisioning (MSSql Only).

* `post_script` - (Optional) [Updatable] PowerShell script or executable to run after provisioning (MSSql Only).

* `cdc_on_provision` - (Optional) [Updatable] Option to enable change data capture (CDC) on both the provisioned VDB and subsequent snapshot-related operations (e.g. refresh, rewind) (MSSql Only).

* `online_log_size` - (Optional) Online log size in MB (Oracle Only).

* `online_log_groups` - (Optional) Number of online log groups (Oracle Only).

* `archive_log` - (Optional) Option to create a VDB in archivelog mode (Oracle Only).

* `new_dbid` - (Optional) [Updatable] Option to generate a new DB ID for the created VDB (Oracle Only).

* `masked` - (Optional) Option to create a Masked VDB. Note: You should define a `configure_clone` script in the Hooks step to mask the dataset. The selection of the "Mask this VDB" option will cause the data to be marked as masked, whether you have defined a script to do so or not.
If you do not define a script to mask the dataset, the data will not be masked unless there is a masking job associated with the source dataset.

* `listener_ids` - (Optional) [Updatable] The listener IDs for this provision operation (Oracle Only). This is a list of listener ids. For eg: [ "listener-123", "listener-456" ]

* `custom_env_vars` - (Optional) 
Environment variable to be set when the engine creates a VDB. See the Engine documentation for the list of allowed/denied environment variables and rules about substitution. This is an ordered map of key-value pairs. For eg: { "MY_ENV_VAR1": "$ORACLE_HOME", "MY_ENV_VAR2": "$CRS_HOME/after" }

* `custom_env_files` - (Optional) Environment files to be sourced when the Engine creates a VDB. This path can be followed by parameters. Paths and parameters are separated by spaces. Valid values are a list of env_files. For eg: [ "/export/home/env_file_1", "/export/home/env_file_2" ]

* `timestamp` - (Optional) The point in time from which to execute the operation. Mutually exclusive with timestamp_in_database_timezone. If the timestamp is not set, selects the latest point.

* `timestamp_in_database_timezone` - (Optional) The point in time from which to execute the operation, expressed as a date-time in the timezone of the source database. Mutually exclusive with timestamp.

* `snapshot_id` - (Optional) The ID or name of the snapshot from which to execute the operation. If the snapshot_id is not, selects the latest snapshot.

* `bookmark_id` - (Optional) The ID or name of the bookmark from which to execute the operation. The bookmark must contain only one VDB.

* `tags` - (Optional) The tags to be created for VDB. This is a map of 2 parameters:
  * `key` - (Required) Key of the tag
  * `value` - (Required) Value of the tag

* `make_current_account_owner` - (Optional) Whether the account provisioning this VDB must be configured as owner of the VDB. 

* `config_params` - (Optional) Database configuration parameter overrides

* `appdata_source_params` - The JSON payload conforming to the DraftV4 schema based on the type of application data being manipulated.

* `appdata_config_params` - (Optional) The list of parameters specified by the source config schema in the toolkit

* `additional_mount_points` - (Optional) Specifies additional locations on which to mount a subdirectory of an AppData container
  * `shared_path` - (Required) Relative path within the container of the directory that should be mounted.
  * `mount_path` - (Required) Absolute path on the target environment were the filesystem should be mounted
  * `environment_id` - (Required) The entity ID of the environment on which the file system will be mounted.

* `vcdb_tde_key_identifier` - (Optional) ID of the key created by Delphix. (Oracle Multitenant Only)

* `cdb_tde_keystore_password` - (Optional) The password for the Transparent Data Encryption keystore associated with the CDB. (Oracle Multitenant Only)

* `target_vcdb_tde_keystore_path` - (Optional) Path to the keystore of the target vCDB. (Oracle Multitenant Only)

* `tde_key_identifier` - (Optional) ID of the key created by Delphix. (Oracle Multitenant Only)

* `tde_exported_key_file_secret` - (Optional) Secret to be used while exporting and importing vPDB encryption keys if Transparent Data Encryption is enabled on the vPDB. (Oracle Multitenant Only)

* `parent_tde_keystore_password` - (Optional) The password of the keystore specified in parentTdeKeystorePath. (Oracle Multitenant Only)

* `parent_tde_keystore_path` - (Optional) Path to a copy of the parent's Oracle transparent data encryption keystore on the target host. Required to provision from snapshots containing encrypted database files. (Oracle Multitenant Only)

* `oracle_rac_custom_env_vars` - (Optional) Environment variable to be set when the engine creates an Oracle RAC VDB. See the Engine documentation for the list of allowed/denied environment variables and rules about substitution.
  * `node_id` - (Required) The node id of the cluster.
  * `name` - (Required) Name of the environment variable
  * `value` - (Required) Value of the environment variable.

* `oracle_rac_custom_env_files` - (Optional) Environment files to be sourced when the Engine creates an Oracle RAC VDB. This path can be followed by parameters. Paths and parameters are separated by spaces.
  * `node_id` - (Required) The node id of the cluster.
  * `path_parameters` - (Required) This references a file from which certain parameters will be loaded.


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
