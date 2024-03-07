# Resource: <resource name> delphix_oracle_dsource

In Delphix terminology, a dSource is a database that the Delphix Continuous Data Engine uses to create and update virtual copies of your database. 
A dSource is created and managed by the Delphix Continuous Data Engine.

The Oracle dSource resource allows Terraform to create and delete Oracle dSources. This specifically enables the apply and destroy Terraform commands. Modification of existing dSource resources via the apply command is not supported. All supported parameters are listed below.

## System Requirements

* Data Control Tower v10.0.1+ is required for dSource management. Lower versions are not supported.
* This Oracle dSource Resource only supports Oracle. See the AppData dSource Resource for the support of other connectors (i.e. AppData), such as PostgreSQL and SAP HANA. The Delphix Provider does not support SQL Server or SAP ASE.

## Example Usage

* The linking of a dSource can be performed via direct ingestion as shown in the example below

```hcl
# Link Oracle dSource

resource "delphix_oracle_dsource" "test_oracle_dsource" {
  name                       = "test2"
  source_value               = "DBOMSRB331B3"
  group_id                   = "3-GROUP-1"
  log_sync_enabled           = false
  make_current_account_owner = true
  environment_user_id        = "HOST_USER-1"
  rman_channels              = 2
  files_per_set              = 5
  check_logical              = false
  encrypted_linking_enabled  = false
  compressed_linking_enabled = true
  bandwidth_limit            = 0
  number_of_connections      = 1
  diagnose_no_logging_faults = true
  pre_provisioning_enabled   = false
  link_now                   = true
  force_full_backup          = false
  double_sync                = false
  skip_space_check           = false
  do_not_resume              = false
  files_for_full_backup      = []
  log_sync_mode              = "UNDEFINED"
  log_sync_interval          = 5
}
```

## Argument Reference

* `source_value` - (Required) Id or Name of the source to link.

* `group_id` - (Required)  Id of the dataset group where this dSource should belong to.

* `log_sync_enabled` - (Required) True if LogSync should run for this database.

* `make_current_account_owner` - (Required) Whether the account creating this reporting schedule must be configured as owner of the reporting schedule.

* `description` - (Optional) The notes/description for the dSource.

* `external_file_path` - (Optional) External file path.

* `environment_user_id` - (Optional) Id of the environment user to use for linking.

* `backup_level_enabled` - (Optional) Boolean value indicates whether LEVEL-based incremental backups can be used on the source database.

* `rman_channels` - (Optional)  Number of parallel channels to use.

* `files_per_set` - (Optional) Number of data files to include in each RMAN backup set.

* `check_logical` - (Optional) True if extended block checking should be used for this linked database.

* `encrypted_linking_enabled` - (Optional) True if SnapSync data from the source should be retrieved through an encrypted connection. Enabling this feature can decrease the performance of SnapSync from the source but has no impact on the performance of VDBs created from the retrieved data.

* `compressed_linking_enabled` - (Optional) True if SnapSync data from the source should be compressed over the network. Enabling this feature will reduce network bandwidth consumption and may significantly improve throughput, especially over slow network.

* `bandwidth_limit` - (Optional) Bandwidth limit (MB/s) for SnapSync and LogSync network traffic. A value of 0 means no limit.

* `number_of_connections` - (Optional) Total number of transport connections to use during SnapSync.

* `diagnose_no_logging_faults` - (Optional)  If true, NOLOGGING operations on this container are treated as faults and cannot be resolved manually.

* `pre_provisioning_enabled` - (Optional) If true, pre-provisioning will be performed after every sync.

* `link_now` - (Optional) True if initial load should be done immediately.

* `force_full_backup` - (Optional) Whether or not to take another full backup of the source database.

* `double_sync` - (Optional) True if two SnapSyncs should be performed in immediate succession to reduce the number of logs required to provision the snapshot. This may significantly reduce the time necessary to provision from a snapshot.

* `skip_space_check` - (Optional) Skip check that tests if there is enough space available to store the database in the Delphix Engine. The Delphix Engine estimates how much space a database will occupy after compression and prevents SnapSync if insufficient space is available. This safeguard can be overridden using this option. This may be useful when linking highly compressible databases.

* `do_not_resume` - (Optional) Indicates whether a fresh SnapSync must be started regardless if it was possible to resume the current SnapSync. If true, we will not resume but instead ignore previous progress and backup all datafiles even if already completed from previous failed SnapSync. This does not force a full backup, if an incremental was in progress this will start a new incremental snapshot.

* `files_for_full_backup` - (Optional) List of datafiles to take a full backup of. This would be useful in situations where certain datafiles could not be backed up during previous SnapSync due to corruption or because they went offline.

* `log_sync_mode` - (Optional) LogSync operation mode for this database [ ARCHIVE_ONLY_MODE, ARCHIVE_REDO_MODE, UNDEFINED ]. 

* `log_sync_interval` - (Optional) Interval between LogSync requests, in seconds.

* `non_sys_password` - (Optional) Password for non sys user authentication (Single tenant only).

* `non_sys_username` - (Optional) Non-SYS database user to access this database. Only required for username-password auth (Single tenant only).

* `non_sys_vault` - (Optional) The name or reference of the vault from which to read the database credentials (Single tenant only).

* `non_sys_hashicorp_vault_engine` - (Optional) Vault engine name where the credential is stored (Single tenant only).

* `non_sys_hashicorp_vault_secret_path` - (Optional)  Path in the vault engine where the credential is stored (Single tenant only).

* `non_sys_hashicorp_vault_username_key` - (Optional) Hashicorp vault key for the username in the key-value store (Single tenant only).

* `non_sys_hashicorp_vault_secret_key` - (Optional) Hashicorp vault key for the password in the key-value store (Single tenant only).

* `non_sys_azure_vault_name` - (Optional) Azure key vault name (Single tenant only).

* `non_sys_azure_vault_username_key` - (Optional) Azure vault key for the username in the key-value store (Single tenant only).

* `non_sys_azure_vault_secret_key` - (Optional) Azure vault key for the password in the key-value store (Single tenant only).

* `non_sys_cyberark_vault_query_string` - (Optional) Query to find a credential in the CyberArk vault (Single tenant only).

* `fallback_username` - (Optional) The database fallback username. Optional if bequeath connections are enabled (to be used in case of bequeath connection failures). Only required for username-password auth..

* `fallback_password` - (Optional) Password for fallback username.

* `fallback_vault` - (Optional) The name or reference of the vault from which to read the database credentials.

* `fallback_hashicorp_vault_engine` - (Optional) Vault engine name where the credential is stored.

* `fallback_hashicorp_vault_secret_path` - (Optional) Path in the vault engine where the credential is stored.

* `fallback_hashicorp_vault_username_key` - (Optional) Hashicorp vault key for the username in the key-value store.

* `fallback_hashicorp_vault_secret_key` - (Optional) Hashicorp vault key for the password in the key-value store.

* `fallback_azure_vault_name` - (Optional) Azure key vault name.

* `fallback_azure_vault_username_key` - (Optional) Azure vault key for the username in the key-value store.

* `fallback_azure_vault_secret_key` - (Optional) Azure vault key for the password in the key-value store.

* `fallback_cyberark_vault_query_string` - (Optional) Query to find a credential in the CyberArk vault.

* `tags` - (Optional) The tags to be created for dSource. This is a map of 2 parameters:
    * `key` - (Required) Key of the tag
    * `value` - (Required) Value of the tag

* `ops_pre_log_sync` - (Optional) Operations to perform after syncing a created dSource and before running the LogSync.
    * `name` - Name of the hook
    * `command` - Command to be executed
    * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]` 
    * `credentials_env_vars` - List of environment variables that will contain credentials for this operation
        * `base_var_name` - Base name of the environment variables. Variables are named by appending '_USER', '_PASSWORD', '_PUBKEY' and '_PRIVKEY' to this base name, respectively. Variables whose values are not entered or are not present in the type of credential or vault selected, will not be set.
        * `password` - Password to assign to the environment variables.
        * `vault` - The name or reference of the vault to assign to the environment variables.
        * `hashicorp_vault_engine` - Vault engine name where the credential is stored.
        * `hashicorp_vault_secret_path` -  Path in the vault engine where the credential is stored.
        * `hashicorp_vault_username_key` - Hashicorp vault key for the username in the key-value store.
        * `hashicorp_vault_secret_key` - Hashicorp vault key for the password in the key-value store.
        * `azure_vault_name` - Azure key vault name.
        * `azure_vault_username_key` - Azure vault key in the key-value store.
        * `azure_vault_secret_key` - Azure vault key in the key-value store.
        * `cyberark_vault_query_string` - Query to find a credential in the CyberArk vault.

* `ops_pre_sync` - (Optional) Operations to perform before syncing the created dSource. These operations can quiesce any data prior to syncing
    * `name` - Name of the hook
    * `command` - Command to be executed
    * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]` 
    * `credentials_env_vars` - List of environment variables that will contain credentials for this operation
        * `base_var_name` - Base name of the environment variables. Variables are named by appending '_USER', '_PASSWORD', '_PUBKEY' and '_PRIVKEY' to this base name, respectively. Variables whose values are not entered or are not present in the type of credential or vault selected, will not be set.
        * `password` - Password to assign to the environment variables.
        * `vault` - The name or reference of the vault to assign to the environment variables.
        * `hashicorp_vault_engine` - Vault engine name where the credential is stored.
        * `hashicorp_vault_secret_path` -  Path in the vault engine where the credential is stored.
        * `hashicorp_vault_username_key` - Hashicorp vault key for the username in the key-value store.
        * `hashicorp_vault_secret_key` - Hashicorp vault key for the password in the key-value store.
        * `azure_vault_name` - Azure key vault name.
        * `azure_vault_username_key` - Azure vault key in the key-value store.
        * `azure_vault_secret_key` - Azure vault key in the key-value store.
        * `cyberark_vault_query_string` - Query to find a credential in the CyberArk vault.
    
* `ops_post_sync` - (Optional) Operations to perform after syncing a created dSource.
    * `name` - Name of the hook
    * `command` - Command to be executed
    * `shell` - Type of shell. Valid values are `[bash, shell, expect, ps, psd]` 
    * `credentials_env_vars` - List of environment variables that will contain credentials for this operation
        * `base_var_name` - Base name of the environment variables. Variables are named by appending '_USER', '_PASSWORD', '_PUBKEY' and '_PRIVKEY' to this base name, respectively. Variables whose values are not entered or are not present in the type of credential or vault selected, will not be set.
        * `password` - Password to assign to the environment variables.
        * `vault` - The name or reference of the vault to assign to the environment variables.
        * `hashicorp_vault_engine` - Vault engine name where the credential is stored.
        * `hashicorp_vault_secret_path` -  Path in the vault engine where the credential is stored.
        * `hashicorp_vault_username_key` - Hashicorp vault key for the username in the key-value store.
        * `hashicorp_vault_secret_key` - Hashicorp vault key for the password in the key-value store.
        * `azure_vault_name` - Azure key vault name.
        * `azure_vault_username_key` - Azure vault key in the key-value store.
        * `azure_vault_secret_key` - Azure vault key in the key-value store.
        * `cyberark_vault_query_string` - Query to find a credential in the CyberArk vault.

* `skip_wait_for_snapshot_creation` - (Optional) By default this resource waits for snapshot to be created post dsource creation , default behaviour can be skipped by setting this parameter to true.

* `wait_time` - (Optional) By default this resource waits for snapshot to be created for 3 minuites , this pameter can be used to set the wait_time in minutes as per need.