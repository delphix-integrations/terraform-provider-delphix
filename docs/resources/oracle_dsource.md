# Resource: <resource name> delphix_oracle_dsource 

In Delphix terminology, a dSource is an internal, read-only database copy that the Delphix Continuous Data Engine uses to create and update virtual copies of your database.  

A dSource is created and managed by the Delphix Continuous Data Engine and syncs with your chosen source database. 

The Oracle dSource resource allows Terraform to create and delete Oracle dSources via Terraform automation. This specifically enables the `apply`, `import`, and `destroy` Terraform commands. 

Updating existing dSource resource parameters via the `apply` command is supported for the parameters listed below.  

This Oracle dSource resource only supports Oracle. 

For other connectors, such as PostgreSQL and SAP HANA, refer to the AppData dSource resource. The Delphix Provider does not currently support SQL Server or SAP ASE. 


## Note 

* `status` and `enabled` are computed values and are subject to change in the tfstate file based on the dSource state. 
* Parameters `credentials_env_vars` within `ops_pre_sync`, `ops_post_sync` and `ops_pre_log_sync` object blocks are not updatable. Any changes reflected on the state file do not reflect the actual value of the actual infrastructure. 
* Sensitive values in `credentials_env_vars` are stored as plain text in the state file. 
* `Make_current_account_owner `,`wait_time` and `skip_wait_for_snapshot_creation` are relevant only during the creation of dsource. Note, they can only be used once and are not applicable to updates.
*  `source_value` and `group_id` parameters cannot be updated after the initial resource creation. However, any differences detected in these parameters are suppressed from the Terraform plan to prevent unnecessary drift detection

  

## Example Usage 

* The linking of a dSource can be performed via direct ingestion as shown in the example below:


```hcl 

# Link Oracle dSource 

resource "delphix_oracle_dsource" "test_oracle_dsource" { 
  name                       = "test2" 
  source_value               = "DBOMSRB331B3"
} 

``` 

## Argument References 

 ### General Linking Requirements 

* `name` - The unique name of the dSource. If empty, a name is randomly generated. [Updatable] 
* `source_value` - (Required) ID or name of the source to link. 
* `description` - The notes (or description) for the dSource. 
* `group_id` - ID of the Delphix Continuous Data dataset group where this dSource should belong to. This value is not reflected in DCT. Tags are recommended. 
* `rollback_on_failure` - Dsource linking operation when fails during snapsync creates a tainted dsource on the engine. Setting this flag to true will remove the tainted dsource from state as well as engine. By default, it is set to false, where the tainted dsource is maintained on the terraform state.

### Full Backup and Transaction Log Requirements 

* `external_file_path` - External file path. [Updatable] 
* `encrypted_linking_enabled` - True if SnapSync data from the source should be retrieved through an encrypted connection. Enabling this feature can decrease the performance of SnapSync from the source but has no impact on the performance of VDBs created from the retrieved data. [Updatable] 
* `compressed_linking_enabled` - True if SnapSync data from the source should be compressed over the network. Enabling this feature will reduce network bandwidth consumption and may significantly improve throughput, especially over slow network. [Updatable] 
* `check_logical` - True if extended block checking should be used for this linked database. [Updatable] 
* `files_for_full_backup` - List of datafiles to take a full backup of. This is useful if certain datafiles could not be backed up during previous SnapSync due to corruption or because they went offline. 
* `files_per_set` - The number of data files to include in each RMAN backup set. [Updatable] 
* `rman_channels` - The number of parallel channels to use. [Updatable] 
* `bandwidth_limit` - Bandwidth limit (MB/s) for SnapSync and LogSync network traffic. A value of 0 means no limit. [Updatable] 
* `number_of_connections` - Total number of transport connections to use during SnapSync. [Updatable] 
* `backup_level_enabled` - Boolean value indicates whether LEVEL-based incremental backups can be used on the source database. [Updatable] 
* `diagnose_no_logging_faults` - If true, NOLOGGING operations on this container are treated as faults and cannot be resolved manually. [Updatable]
* `pre_provisioning_enabled` - If true, pre-provisioning will be performed after every sync. [Updatable] 
* `log_sync_enabled` - (Required) True if LogSync should run for this database. 
* `log_sync_mode` - LogSync operation mode for this database [`ARCHIVE_ONLY_MODE`, `ARCHIVE_REDO_MODE`, `UNDEFINED`].  
* `log_sync_interval` - Interval between LogSync requests, in seconds.
* `link_now` - True if initial load should be done immediately. 


### Snapshot 

The following arguments enable the user to control how the first snapshot should be taken.  

* `force_full_backup` - Whether to take another full backup of the source database. 
* `double_sync` - True if two SnapSyncs should be performed in immediate succession to reduce the number of logs required to provision the snapshot. This may significantly reduce the time necessary to provision from a snapshot. 
* `do_not_resume` - Indicates if a fresh SnapSync must be started regardless of whether it was possible to resume the current SnapSync. If true, we will not resume; instead, we will ignore previous progress and back up all datafiles even if they have already been completed from the last failed SnapSync. This does not force a full backup; if an incremental was in progress this will start a new incremental snapshot. 
* `skip_wait_for_snapshot_creation` - In DCT v2025.1, waiting for Ingestion and Snapshotting (aka SnapSync) to complete is default functionality. Therefore, these the arguments skip_wait_for_snapshot_creation and wait_time are ignored. In future versions of the provider, we will look at re-implementing the skip SnapSync behavior 
* `wait_time` - In DCT v2025.1, waiting for Ingestion and Snapshotting (aka SnapSync) to complete is default functionality. Therefore, these the arguments skip_wait_for_snapshot_creation and wait_time are ignored. In future versions of the provider, we will look at re-implementing the skip SnapSync behavior.  

### Password and Password Vault Management 

The following arguments define how the Delphix Continuous Data will authenticate with the source environment and database. 

* `environment_user_id` - ID of the environment user to use for linking. [Updatable] 
* `non_sys_username` - non-SYS database user to access this database. Only required for username-password auth (single tenant only). 
* `non_sys_password` - Password for non sys user authentication (single tenant only). 
* `fallback_username` - The database fallback username. Optional if bequeath connections are enabled (to be used in case of bequeath connection failures). Only required for username-password auth. 
* `fallback_password` - Password for fallback username. 
* `non_sys_vault` - The name or reference of the vault from which to read the database credentials (single tenant only). 
* `non_sys_hashicorp_vault_engine` - Vault engine name where the credential is stored (single tenant only). 
* `non_sys_hashicorp_vault_secret_path` - Path in the vault engine where the credential is stored (single tenant only). 
* `non_sys_hashicorp_vault_username_key` - Hashicorp vault key for the username in the key-value store (single tenant only). 
* `non_sys_hashicorp_vault_secret_key` - Hashicorp vault key for the password in the key-value store (single tenant only). 
* `non_sys_azure_vault_name` - Azure key vault name (single tenant only). 
* `non_sys_azure_vault_username_key` - Azure vault key for the username in the key-value store (single tenant only). 
* `non_sys_azure_vault_secret_key` - Azure vault key for the password in the key-value store (single tenant only). 
* `non_sys_cyberark_vault_query_string` - Query to find a credential in the CyberArk vault (single tenant only). 
* `fallback_vault` - The name or reference of the vault from which to read the database credentials. 
* `fallback_hashicorp_vault_engine` - Vault engine name where the credential is stored. 
* `fallback_hashicorp_vault_secret_path` - Path in the vault engine where the credential is stored. 
* `fallback_hashicorp_vault_username_key` - Hashicorp vault key for the username in the key-value store. 
* `fallback_hashicorp_vault_secret_key` - Hashicorp vault key for the password in the key-value store. 
* `fallback_azure_vault_name` - Azure key vault name. 
* `fallback_azure_vault_username_key` - Azure vault key for the username in the key-value store. 
* `fallback_azure_vault_secret_key` - Azure vault key for the password in the key-value store. 
* `fallback_cyberark_vault_query_string` - Query to find a credential in the CyberArk vault. 


### Advanced  

The following arguments apply to all dSources but they are not often necessary for simple sources. 

* `make_current_account_owner` - Whether the account creating this reporting schedule must be configured as owner of the reporting schedule. Default: true. 
 * `tags` - The tags to be created for dSource. This is a map of 2 parameters: [Updatable] 
    * `key` - (Required) Key of the tag 
    * `value` - (Required) Value of the tag 

### Hooks
Any combination of the following hooks can be provided on the Oracle dSource resource. The available arguments are identical for each hook and are consolidated in a single list to save space. 

#### Names
* `ops_pre_log_sync`: Operations to perform after syncing a created dSource and before running the LogSync. See argument list below. 
* `ops_pre_sync`: Operations to perform before syncing the created dSource. These operations can quiesce any data prior to syncing. See argument list below. 
* `ops_post_sync`: Operations to perform after syncing a created dSource. See argument list below. 

#### Arguments
* `name` - Name of the hook [Updatable]   
* `command` - Command to be executed [Updatable]  
* `shell` - Type of shell. Valid values are [bash, shell, expect, ps, psd] [Updatable] 
* `credentials_env_vars` - List of environment variables that contain credentials for this operation 
* `base_var_name` - Base name of the environment variables. Variables are named by appending '_USER', '_PASSWORD', '_PUBKEY' and '_PRIVKEY' to this base name, respectively. Variables whose values are not entered or present in the type of credential or vault selected will not be set.  
* `password` - Password to assign to the environment variables. - vault - The name or reference of the vault to assign to the environment variables.  
* `hashicorp_vault_engine` - Vault engine name where the credential is stored.  
* `hashicorp_vault_secret_path` - Path in the vault engine where the credential is stored.  
* `hashicorp_vault_username_key` - Hashicorp vault key for the username in the key-value store.  
* `hashicorp_vault_secret_key` - Hashicorp vault key for the password in the key-value store.  
* `azure_vault_name` - Azure key vault name.  
* `azure_vault_username_key` - Azure vault key in the key-value store. 
* `azure_vault_secret_key` - Azure vault key in the key-value store.  
* `cyberark_vault_query_string` - Query to find a credential in the CyberArk vault. 

## Import (Beta)  
Use the [`import` block](https://developer.hashicorp.com/terraform/language/import) to add Oracle Dsources created directly in DCT into a Terraform state file.  

For example:  
```terraform 
import {   
    to = delphix_oracle_dsource.dsrc_import_demo
    id = "dsource_id"   
}  
``` 
*This is a beta feature. Delphix offers no guarantees of support or compatibility.* 

## Limitations 

Not all properties are supported through the `update` command. Properties that are not supported by the `update` command are presented via an error message at runtime. 