# <resource name> delphix_vdb

The VDB resource allows the provisioning, updation and deletion of a Delphix virtual database infrastructure (VDB). However, update is not supported for all parameters, the ones supported is mentioned below.

## Example Usage
Provisioning can be done in 2 methods, provision by snapshot and provision by timestamp.

```hcl
# Provision a VDB with only non optional parameters

resource "delphix_vdb" "vdb_name" {
  auto_select_repository = true
  source_data_id         = "DATASOURCE_ID"
}

# Provision a VDB using timestamp

resource "delphix_vdb" "vdb_name2" {
  provision_type         = "timestamp"
  auto_select_repository = true
  source_data_id         = "dsource2"
  timestamp              = "2021-05-01T08:51:34.148000+00:00"

  post_refresh {
    name    = "n"
    command = "time"
    shell   = "bash"
  }
  post_refresh {
    name    = "n2"
    command = "time"
    shell   = "bash"
  }
}

# Provision a VDB using snapshot

resource "delphix_vdb" "vdb_name" {
  provision_type         = "snapshot"
  auto_select_repository = true
  source_data_id         = "dsource"

  pre_refresh {
    name    = "n"
    command = "time"
    shell   = "bash"
  }
  pre_refresh {
    name    = "n2"
    command = "time"
    shell   = "bash"
  }
}
```

## Argument Reference

* `source_data_id` - (Required) The ID of the source object (dSource or VDB) to provision from. All other objects referenced by the parameters must live on the same engine as the source.

* `engine_id` - (Optional) The ID of the Engine onto which to provision. If the source ID unambiguously identifies a source object, this parameter is unnecessary and ignored.

* `target_group_id` - (Optional) The ID of the group into which the VDB will be provisioned. If unset, a group is selected randomly on the Engine.

* `vdb_name` - (Optional) The unique name of the provisioned VDB within a group. If unset, a name is randomly generated.

* `database_name` - (Optional) The name of the database on the target environment. Defaults to vdb_name.

* `truncate_log_on_checkpoint` - (Required) Whether to truncate log on checkpoint (ASE only).

* `username` - (Required) [Updatable] The name of the privileged user to run the provision operation (Oracle Only).

* `password` - (Required) [Updatable] The password of the privileged user to run the provision operation (Oracle Only).

* `environment_id` - (Optional) The ID of the target environment where to provision the VDB. If repository_id unambigously identifies a repository, this is unnecessary and ignored. Otherwise, a compatible repository is randomly selected on the environment.

* `environment_user_id` - [Updatable] The environment user ID to use to connect to the target environment.

* `repository_id` - (Optional) The ID of the target repository where to provision the VDB. A repository typically corresponds to a database installation (Oracle home, database instance, ...). Setting this attribute implicitly determines the environment where to provision the VDB.

* `auto_select_repository` - (Required) Option to automatically select a compatible environment and repository. Mutually exclusive with repository_id.

* `pre_refresh` - (Optional) The commands to execute on the target environment before refreshing the VDB.

* `post_refresh` - (Optional) The commands to execute on the target environment after refreshing the VDB.

* `pre_rollback` - (Optional) The commands to execute on the target environment before rewinding the VDB.

* `post_rollback` - (Optional) The commands to execute on the target environment after rewinding the VDB.

* `configure_clone` - (Optional) The commands to execute on the target environment when the VDB is created or refreshed.

* `pre_snapshot` - (Optional) The commands to execute on the target environment before snapshotting a virtual source. These commands can quiesce any data prior to snapshotting.

* `post_snapshot` - (Optional) The commands to execute on the target environment after snapshotting a virtual source.

* `pre_start` - (Optional) The commands to execute on the target environment before starting a virtual source.

* `post_start` - (Optional) The commands to execute on the target environment after starting a virtual source.

* `pre_stop` - (Optional) The commands to execute on the target environment before stopping a virtual source.

* `post_stop` - (Optional) The commands to execute on the target environment after stopping a virtual source.

* `vdb_restart` - (Optional) [Updatable] Indicates whether the Engine should automatically restart this virtual source when target host reboot is detected.

* `template_id` - (Required) [Updatable] The ID of the target VDB Template (Oracle Only).

* `file_mapping_rules` - (Required) Target VDB file mapping rules (Oracle Only). Rules must be line separated (\n or \r) and each line must have the format "pattern:replacement". Lines are applied in order.

* `oracle_instance_name` - (Required) Target VDB SID name (Oracle Only).

* `unique_name` - (Required) Target VDB db_unique_name (Oracle Only).

* `mount_point` - (Required) Mount point for the VDB (Oracle, ASE Only).

* `open_reset_logs` - (Required) Whether to open the database after provision (Oracle Only).

* `snapshot_policy_id` - (Optional) The ID of the snapshot policy for the VDB.

* `retention_policy_id` - (Optional) The ID of the retention policy for the VDB.

* `recovery_model` - (Optional) Recovery model of the source database (MSSql Only).

* `pre_script` - (Required) [Updatable] PowerShell script or executable to run prior to provisioning (MSSql Only).

* `post_script` - (Optional) [Updatable] PowerShell script or executable to run after provisioning (MSSql Only).

* `cdc_on_provision` - (Optional) [Updatable] Option to enable change data capture (CDC) on both the provisioned VDB and subsequent snapshot-related operations (e.g. refresh, rewind) (MSSql Only).

* `online_log_size` - (Required) Online log size in MB (Oracle Only).

* `online_log_groups` - (Required) Number of online log groups (Oracle Only).

* `archive_log` - (Required) Option to create a VDB in archivelog mode (Oracle Only).

* `new_dbid` - (Required) [Updatable] Option to generate a new DB ID for the created VDB (Oracle Only).

* `listener_ids` - (Required) [Updatable] The listener IDs for this provision operation (Oracle Only).

* `custom_env_vars` - (Optional) 
Environment variable to be set when the engine creates a VDB. See the Engine documentation for the list of allowed/denied environment variables and rules about substitution.

* `custom_env_files` - (Optional) Environment files to be sourced when the Engine creates a VDB. This path can be followed by parameters. Paths and parameters are separated by spaces.

* `timestamp` - (Optional) The point in time from which to execute the operation. Mutually exclusive with timestamp_in_database_timezone. If the timestamp is not set, selects the latest point.

* `timestamp_in_database_timezone` - (Optional) The point in time from which to execute the operation, expressed as a date-time in the timezone of the source database. Mutually exclusive with timestamp.

* `snapshot_id` - (Optional) The ID of the snapshot from which to execute the operation. If the snapshot_id is not, selects the latest snapshot.


## Attribute Reference

* `id` - The VDB object entity ID.

* `database_type` - The database type of this VDB.

* `name` - The container name of this VDB.

* `database_version` - The database version of this VDB.

* `size` - The total size of this VDB, in bytes.

* `engine_id` - A reference to the Engine that this VDB belongs to. 

* `status` - The runtime status of the VDB. 'Unknown' if all attempts to connect to the dataset failed.

* `environment_id` - A reference to the Environment that hosts this VDB.

* `ip_address` - The IP address of the VDB's host.

* `fqdn` - The FQDN of the VDB's host.

* `parent_id` - A reference to the parent dataset of this VDB.

* `group_name` - The name of the group containing this VDB.

* `tags` - A list of key value pair.

* `creation_date` - The date this VDB was created.
