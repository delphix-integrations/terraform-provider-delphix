# Resource: <resource name> delphix_vdb 

In Delphix, a virtual database (VDB) is a full read/write copy of the source data. A VDB is created and managed by Data Control Tower (DCT) and Delphix Continuous Data Engine and it is provisioned (created) from a dSource or a VDB's data snapshot.  

The VDB (delphix_vdb) resource allows Terraform to create, update, and delete Delphix VDBs. The resource enables the `plan`, `apply`, `update`, and `destroy` Terraform commands.  


## Example Usage 
Provisioning can be done using one of three methods: provision by snapshot, bookmark, or timestamp.  

### Provisioning a VDB using latest snapshot.
```terraform 

resource "delphix_vdb" "vdb_name_provision_by_snapshot" {  
  auto_select_repository = true  
  source_data_id         = "<DATASOURCE_ID_OR_NAME>"  
  snapshot_id            = "<SNAPSHOT_ID>" # Leave empty to select the latest snapshot  
}  

``` 
### Provisioning a VDB from a bookmark and on a Target environment 
```terraform 
resource "delphix_vdb" "vdb_name_provision_by_bookmark_on_target_environment" {  
  provision_type         = "bookmark"  
  auto_select_repository = true  
  bookmark_id            = "<BOOKMARK_ID_OR_NAME>"  
  environment_id         = "<ENV_ID>"  
}  
``` 
### Provisioning a VDB using timestamp and configure post refresh hook  
```terraform 
resource "delphix_vdb" "vdb_name_provion_by_timestamp_with_hook" {  
  provision_type         = "timestamp"  
  auto_select_repository = true  
  source_data_id         = "<DSOURCE_OR_VDB_ID>"  
  timestamp              = "2021-05-01T08:51:34.148000+00:00"   # Timestamp must be available on the source dataset.  
  post_refresh {  
    command         = "echo \"Hello World\""  
    name            = "Sample Hook"  
    shell           = "SHELL"  
  }  
}  
``` 

## Argument References 

### General Provisioning Requirements  
The following arguments apply to all database platform types.  

* `name` - (Required) The unique name of the VDB. If empty, a name is randomly generated. [Updatable]  
* `source_data_id` - (Required) The ID or name of the source dataset (dSource, VDB, Snapshot, or timestamp point) to provision from.   
    * All other objects referenced by the following parameters must live on the same Continuous Data Engine as the chosen source.  
* `provision_type` - The type of provisioning to be carried out. Defaults to snapshot.   
    * Valid values are `[snapshot, bookmark, timestamp]`. This value determines which matching `snapshot_id`, `bookmark_id`, and `timestamp` argument is required.  
* `snapshot_id` - The ID or name of the Snapshot from which to execute the provision operation. If the `snapshot_id` is empty or the parameter is not specified, the latest snapshot is automatically selected.  
* `bookmark_id` - The ID or name of the Bookmark from which to execute the provision operation. The Bookmark must contain only one VDB.  
* `timestamp` or `timestamp_in_database_timezone` - The point in time from which to execute the provision operation.   
    * If the `provision_type` is set to `timestamp`, but a `timestamp` value is not provided, then the latest available point is selected.   

### Oracle 
The following arguments apply to the Oracle database type. 

__General Oracle__  
<br /> The following arguments apply to all Oracle deployment configurations.  

* `database_name` - The name of the database on the Target environment.   
    * Defaults to "name".  
* `unique_name` - The VDB's unique name (aka db_unique_name).  
* `db_username` - The username of the database. If blank, this VDB will use its parent’s credentials [Updatable]  
* `db_password` - The password of the database.  If blank, this VDB will use its parent’s credentials [Updatable]  
* `template_id` - The ID of the VDB Configuration Template. [Updatable]  
* `mount_point` - The mount point for the VDB. [Updatable]  
* `os_username` - The name of the Target’s operating system user to run the provision operation.  
* `os_password` - The password of the Target’s operating system user to run the provision operation.  
* `instance_name` - The VDB's SID name.  
    * This parameter is available starting from Data Control Tower (DCT) v22+.  
* `open_reset_logs` - TRUE or FALSE value which determines whether to open the database after provision.  
* `online_log_size` - The online log size in MB.  
* `online_log_groups` - The number of online log groups.  
* `archive_log` - TRUE or FALSE boolean to create a VDB in `archivelog` mode.  
* `new_dbid` - TRUE or FALSE boolean to generate a new DB ID for the created VDB. [Updatable]  
* `listener_ids` - The listener IDs for this provision operation. This is a list of listener IDs. For eg: `[ "listener-123", "listener-456" ]`. [Updatable]  
* `file_mapping_rules` - The VDB file mapping rules. Rules must be line separated (\n or \r) and each line must have the format "pattern:replacement". Lines are applied in order.  

__Oracle Multitenant__  
<br /> In addition to the General Oracle arguments, the following list applies to the Oracle Multitenant deployment configuration.  

* `cdb_id` - The ID of the container database (CDB) to provision an Oracle Multitenant database into.   
    * If empty, a new vCDB will be provisioned.   
* `vcdb_name` - The name of the provisioned vCDB when the `cdb_id` property is not set.  
* `vcdb_database_name` - The database name of the provisioned vCDB when the `cdb_id` property is not set.   
    * Defaults to the value of the `vcdb_name` property.  
* `auxiliary_template_id` - The ID of the configuration template to apply to the auxiliary container database (CDB). This is only relevant when provisioning a Multitenant pluggable database into an existing CDB, i.e when the `cdb_id` property is set.  
* `vcdb_tde_key_identifier` - ID of the key created by the Delphix Continuous Data Engine.   
* `cdb_tde_keystore_password` - The password for the Transparent Data Encryption keystore associated with the CDB. [Updatable]  
* `target_vcdb_tde_keystore_path` - [Updatable] Path to the keystore of the vCDB.   
* `tde_key_identifier` - ID of the key created by the Continuous Data Engine. [Updatable]  
* `tde_exported_key_file_secret` - Secret to be used when exporting and importing vPDB encryption keys if Transparent Data Encryption is enabled on the vPDB.   
* `parent_tde_keystore_password` - The password of the keystore specified in parentTdeKeystorePath. [Updatable]  
* `parent_tde_keystore_path` - Path to a copy of the parent's Oracle Transparent Data Encryption keystore on the target host. Required to provision from snapshots containing encrypted database files. [Updatable]  

__Oracle Real Applications Clusters (RAC)__  
<br /> In addition to the General Oracle arguments, the following list applies to the Oracle RAC deployment configuration. All properties marked as required are necessary for Oracle RAC provisions.  

* `cluster_node_ids` - The cluster node IDs, name, or addresses for this provision operation.  
* `oracle_rac_custom_env_vars` - Environment variable to be set when the engine creates an Oracle RAC VDB. See the Delphix Continuous Data Engine documentation for the list of allowed/denied environment variables and rules about substitution.  
    * `node_id` - (Required) The node ID of the cluster.  
    * `name` - (Required) Name of the environment variable  
    * `value` - (Required) Value of the environment variable.  
* `oracle_rac_custom_env_files` - Environment files to be sourced when the Delphix Continuous Data Engine creates an Oracle RAC VDB. This path can be followed by parameters. Paths and parameters are separated by spaces.  
    * `node_id` - (Required) The node ID of the cluster.  
    * `path_parameters` - (Required) This references a file from which certain parameters will be loaded.  

### SQL Server 
The following arguments apply to the Microsoft SQL Server database type.  

* `database_name` - The name of the database on the Target environment. Defaults to "name".   
* `template_id` - The ID of the VDB Configuration Template. [Updatable]  
* `recovery_model` - Recovery model of the source database. Valid values are `[ FULL, SIMPLE, BULK_LOGGED ]`.  
* `cdc_on_provision` - Option to enable change data capture (CDC) on the provisioned VDB and subsequent snapshot-related operations (e.g. refresh, rewind). [Updatable]  
* `pre_script` - PowerShell script or executable to run prior to provisioning. [Updatable]  
* `post_script` - PowerShell script or executable to run after provisioning. [Updatable] 

### SAP ASE 
The following arguments apply to the SAP ASE database type.  

* `database_name` - The name of the database on the Target environment. Defaults to "name".  
* `db_username` - The username of the database. If blank, this VDB will use its parent’s credentials. [Updatable]  
* `db_password` - The password of the database.  If blank, this VDB will use its parent’s credentials. [Updatable]  
* `mount_point` - The mount point for the VDB. [Updatable]  
* `truncate_log_on_checkpoint` - TRUE or FALSE value to truncate the logs on checkpoints.  

### Other Databases 
The following arguments apply to all other databases supported by the AppData (vSDK) framework and include database types such as PostgreSQL, MySQL, MongoDB, etc. 

* `appdata_source_params` - (Required) The JSON payload conforming to the DraftV4 schema based on the type of application data being manipulated. [Updatable]  
    * Consult the connector documentation for more details.   
* `additional_mount_points` - Specifies additional locations on which to mount a subdirectory of an AppData container. [Updatable]  
    * `shared_path` - (Required) Relative path within the container of the directory that should be mounted.  
    * `mount_path` - Absolute path on the Target environment where the filesystem should be mounted.  
    * `environment_id` - The entity ID of the environment on which the file system will be mounted.  
 
### Target Engine, Environment and Repository  
The following arguments allow users to determine where the VDB will be provisioned.  
If left blank, DCT will auto-select a location based on the source.  

* `auto_select_repository` - (Required, unless `repository_id`, `environment_id`, or `cdb_id` are specified). TRUE or FALSE value to automatically select a compatible environment and repository.   
* `engine_id` - The ID or name of the Continuous Data Engine onto which to provision. If the source ID unambiguously identifies a source object, this parameter is unnecessary and ignored.  
* `environment_id` - The ID or name of the Target environment where to provision the VDB.   
    * If `repository_id` unambiguously identifies a repository, then this value is ignored.  
* `environment_user_id` - The environment user ID to use to connect to the Target environment. [Updatable]  
* `repository_id` - The ID of the Target environment's repository where to provision the VDB. A repository typically corresponds to a database installation (Oracle home, database instance, etc). Setting this parameter may implicitly determine the environment where to provision the VDB.  
* `target_group_id` - The ID of the Continuous Data Engine's Dataset Group into which the VDB will be provisioned. If empty, the "Unassigned" Dataset Group is used.  
    * We encourage all users to avoid this argument as it will be deprecated and removed in the future.  

### Advanced 
The following arguments are applicable to all database types, but they are not often necessary for simple provisions.  

* `custom_env_vars` - Environment variable to be set when a VDB is created. See the Delphix Continuous Data Engine documentation for the list of allowed/denied environment variables and rules about substitution. This is an ordered map of key-value pairs. For eg: { "MY_ENV_VAR1": "$ORACLE_HOME", "MY_ENV_VAR2": "$CRS_HOME/path/here" }  
* `custom_env_files` - Environment files to be sourced when a VDB is created. This path can be followed by parameters. Paths and parameters are separated by spaces. Valid values are a list of env_files. For eg: [ "/export/home/env_file_1", "/export/home/env_file_2" ]  
* `config_params` - The database configuration override parameters. [Updatable]   
* `pre_refresh` - The commands to execute on the Target environment before refreshing the VDB. [Updatable]  
This is a map of three parameters:  
    * name - Name of the hook.  
    * command - (Required, if hook is specified). Command to be executed.  
    * shell - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`.  
* `post_refresh` - The commands to execute on the Target environment after refreshing the VDB. [Updatable]  
This is a map of three parameters:  
    * name - Name of the hook.  
    * command - (Required, if hook is specified). Command to be executed.  
    * shell - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`.  
* `pre_rollback` - The commands to execute on the Target environment before a rollback on the VDB. [Updatable]  
This is a map of three parameters:  
    * `name` - Name of the hook.  
    * `command` - (Required, if hook is specified). Command to be executed.  
    * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`.  
* `post_rollback` - The commands to execute on the Target environment after a rollback on the VDB. [Updatable]  
This is a map of three parameters:  
    * `name` - Name of the hook.  
    * `command` - (Required, if hook is specified). Command to be executed.  
    * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`.  
* `configure_clone` - The commands to execute on the Target environment when the VDB is created or refreshed. [Updatable]  
This is a map of three parameters:  
    * `name` - Name of the hook.  
    * `command` - (Required, if hook is specified). Command to be executed.  
    * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`.  
* `pre_snapshot` - The commands to execute on the Target environment before snapshotting a virtual database. These commands can quiesce any data prior to snapshotting. [Updatable]  
This is a map of five parameters:  
    * `name` - Name of the hook.  
    * `command` - (Required, if hook is specified). Command to be executed.  
    * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`.  
* `post_snapshot` - The commands to execute on the Target environment after snapshotting a virtual database. [Updatable]  
This is a map of three parameters:  
    * `name` - Name of the hook.  
    * `command` - (Required, if hook is specified). Command to be executed.  
    * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`.  
* `pre_start` - The commands to execute on the Target environment before starting a virtual database. [Updatable]  
This is a map of three parameters:  
    * `name` - Name of the hook.  
    * `command` - (Required, if hook is specified). Command to be executed.  
    * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`.  
* `post_start` - The commands to execute on the Target environment after starting a virtual database. [Updatable]  
This is a map of three parameters:  
    * `name` - Name of the hook.  
    * `command` - (Required, if hook is specified). Command to be executed.  
    * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`.  
* `pre_stop` - The commands to execute on the Target environment before stopping a virtual database. [Updatable]  
This is a map of three parameters:  
    * `name` - Name of the hook.  
    * `command` - (Required, if hook is specified). Command to be executed.  
    * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`.  
* `post_stop` - The commands to execute on the Target environment after stopping a virtual database. [Updatable]  
This is a map of three parameters:  
    * `name` - Name of the hook.  
    * `command` - (Required, if hook is specified) Command to be executed.  
    * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]`.  
* `vdb_restart` - Indicates whether the Delphix Continuous Data Engine should automatically restart this virtual database when Target environment reboot is detected. [Updatable]  
* `snapshot_policy_id` - The ID of the Snapshot Policy for the VDB.  
* `retention_policy_id` - The ID of the Snapshot Retention Policy for the VDB.  
* `masked` - TRUE or FALSE boolean to set a VDB as "Masked".   
    * You should define a `configure_clone` script in the Hooks step to mask the dataset. The selection of this option will cause the data to be marked as masked, regardless of whether you have defined a script to do so or not. If you do not define a script to mask the dataset, the data will not be masked unless there is a masking job associated with the dataset.  
* `tags` - The tags to be created for the VDB. [Updatable]  
This is a map of two required parameters:  
    * `key` - Key of the tag.  
    * `value` - Value of the tag.  
* `make_current_account_owner` - Default True. Boolean to determine if the account provisioning this VDB will be the "Owner" of the VDB.  

## Import (Beta)  
Use the [`import` block](https://developer.hashicorp.com/terraform/language/import) to add VDBs created directly in DCT into a Terraform state file.  

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