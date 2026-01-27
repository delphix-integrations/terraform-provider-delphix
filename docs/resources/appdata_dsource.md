# Resource: <resource name> delphix_appdata_dsource

In Delphix terminology, a dSource is an internal, read-only database copy that the Delphix Continuous Data Engine uses to create and update virtual copies of your database.   
A dSource is created and managed by the Delphix Continuous Data Engine and syncs with your chosen source database. The AppData dSource resource in Terraform allows you to create, update, delete and import AppData dSources. Updating existing dSource resource parameters via the apply command is supported for the parameters listed below.    
For Oracle, refer to the Oracle dSource resource. The Delphix Provider does not currently support SQL Server or SAP ASE. 

## Note 
* `status` and `enabled` are computed values and are subject to change in the tfstate file based on the dSource state. 
* Parameters `credentials_env_vars` within `ops_pre_sync` and `ops_post_sync` object blocks are not updatable. Therefore, any changes made on a Terraform state file do not reflect the actual value of the actual infrastructure 
* Sensitive values in `credentials_env_vars` are stored as plain text in the state file. We recommend following Terraform’s sensitive input variables documentation. 
* `source_value` and `group_id` parameters cannot be updated after the initial resource creation. However, any differences detected in these parameters are suppressed from the Terraform plan to prevent unnecessary drift detection.
* Only valid for DCT versions 2025.1 and earlier: 
  * `Make_current_account_owner`,`wait_time` and `skip_wait_for_snapshot_creation` are applicable only during the creation of dsource. Note, these parameters are single-use and not applicable to updates. 
  * Any new dSource created post Version>=3.2.1 can set wait_time to wait for snapshot creation, dSources created prior to this version will not support this capability. 

## Example Usage

The linking of a dSource can be configured through various ingestion approaches. Each configuration is customized to the connector and its supported options. The three PostgreSQL parameter sets below show different ingestion configuration examples. 

# Link dSource using external backup  
 
```hcl
resource "delphix_appdata_dsource" "pg_using_external_backup" { 
  name                       = DSOURCE_NAME  
  source_value               = SOURCE_VALUE 
  group_id                   = DATASET_GROUP_ID 
  log_sync_enabled           = false 
  link_type                  = "AppDataStaged"
  staging_mount_base         = MOUNT_PATH 
  environment_user           = ENV_USER 
  staging_environment        = STAGING_ENV 
  parameters = jsonencode({ 
    externalBackup: [  
            { 
                keepStagingInSync: false, 
                backupPath: BKP_PATH, 
                walLogPath: LOG_PATH 
            } 
    ], 
    postgresPort : PORT, 
    mountLocation : MOUNT_PATH 
  }) 
  sync_parameters = jsonencode({ 
    resync = true 
  }) 
  make_current_account_owner = true 
  
  timeouts {
    create = "20m"
    update = "20m"
    delete = "20m"
  }
} 
```

# Link dSource using Delphix Initiated Backup 
 
```hcl
resource "delphix_appdata_dsource" "pg_using_delphix_initiated_backup" { 
  name                       = DSOURCE_NAME 
  source_value               = SOURCE_VALUE 
  group_id                   = DATASET_GROUP_ID 
  log_sync_enabled           = false 
  link_type                  = "AppDataStaged" 
  staging_mount_base         = MOUNT_PATH 
  environment_user           = ENV_USER 
  staging_environment        = STAGING_ENV 
  parameters = jsonencode({ 
    delphixInitiatedBackupFlag : true, 
    delphixInitiatedBackup : [ 
      { 
        userName : USERNAME, 
        postgresSourcePort : SOURCE_PORT, 
        userPass : PASSWORD, 
        sourceHostAddress : SOURCE_ADDRESS 
      } 
    ], 
    postgresPort : PORT, 
    mountLocation : MOUNT_PATH 
  }) 
  sync_parameters = jsonencode({ 
    resync = true 
  }) 
  make_current_account_owner = true 
  
  timeouts {
    create = "20m"
    update = "20m"
    delete = "20m"
  }
} 
```

# Link dSource using Single Database Ingestion
 
```hcl
resource "delphix_appdata_dsource" "pg_using_single_db_ingestion" { 
  name                       = DSOURCE_NAME 
  source_value               = SOURCE_VALUE 
  group_id                   = DATASET_GROUP_ID 
  log_sync_enabled           = false 
  link_type                  = "AppDataStaged" 
  staging_mount_base         = MOUNT_PATH 
  environment_user           = ENV_USER 
  staging_environment        = STAGING_ENV 
  parameters = jsonencode({ 
    singleDatabaseIngestionFlag : true, 
    singleDatabaseIngestion : [ 
        { 
            databaseUserName: DBUSER_NAME, 
            sourcePort: SOURCE_PORT, 
            dumpJobs: 2, 
            restoreJobs: 2, 
            databaseName: DB_NAME, 
            databaseUserPassword: DB_PASS, 
            dumpDir: DIR, 
            sourceHost: SOURCE_HOST 
            postgresqlFile: FILE 
        } 
    ], 
    postgresPort : PORT, 
    mountLocation : MOUNT_PATH 
  }) 
  sync_parameters = jsonencode({ 
    resync = true 
  })
  make_current_account_owner = true 
  
  timeouts {
    create = "20m"
    update = "20m"
    delete = "20m"
  }
} 
```  

## Argument Reference

### General Linking Requirements 
* `name` - The unique name of the dSource. If unset, a name is randomly generated. [Updatable] 
* `source_value` - (Required) ID or Name of the Source to link. 
* `description` - The notes/description for the dSource. [Updatable] 
* `group_id` - (Required) ID of the dataset group where this dSource should belong to. 
* `rollback_on_failure` -  When a dSource linking operation fails during SnapSync, it results in a tainted dsource on the engine. By setting this flag to true, the tainted dSource will be removed from both the Terraform state and the engine. By default, the flag is to false, meaning the tainted dSource is maintained on the Terraform state. 

### Full Backup and Transactional Log requirements 
* `log_sync_enabled` - (Required) True if LogSync should run for this database.  
* `link_type` - (Required) The type of link to create. Default is AppDataDirect.  
  * `AppDataDirect` - Represents the AppData specific parameters of a link request for a source directly replicated into the Delphix Continuous Data Engine.  
  * `AppDataStaged` - Represents the AppData specific parameters of a link request for a source with a staging source.  

#### AppDataDirect properties 
* `excludes` - List of subdirectories in the source to exclude when syncing data.These paths are relative to the root of the source directory.  
* `follow_symlinks` - List of symlinks in the source to follow when syncing data. These paths are relative to the root of the source directory. All other symlinks are preserved. 

#### AppDataStaged properties 
* `staging_mount_base` - The base mount point for the NFS mount on the staging environment.  
* `environment_user` - (Required) The OS user to use for linking. [Updatable] 
* `staging_environment` - (Required) The environment used as an intermediate stage to pull data into Delphix. [Updatable] 
* `staging_environment_user` - Specifies the environment user that accesses the staging environment. [Updatable] 
* `parameters` - The JSON payload is based on the type of dSource being created. Different data sources require different parameters. Available parameters can be found within the data connector’s schema.json. [Updatable] 
* `sync_parameters` - The JSON payload conforming to the snapshot parameters definition in a Continuous Data plugin. 
* `sync_policy_id` - The ID of the SnapSync policy for the dSource. [Updatable] 
* `retention_policy_id` - The ID of the Retention policy for the dSource. [Updatable] 

### Hooks 
Any combination of the following hooks can be provided on the AppData dSource resource. The available arguments are identical for each hook and are consolidated in a single list to save space.  

#### Names 
* `ops_pre_sync`: Operations to perform before syncing the created dSource. These operations can quiesce any data prior to syncing. See argument list below.  
* `ops_post_sync`: Operations to perform after syncing a created dSource. See argument list below. 

#### Arguments 
* `name` - Name of the hook [Updatable]  
* `command` - Command to be executed [Updatable]  
* `shell` - Type of shell. Valid values are [bash, shell, expect, ps, psd] [Updatable]  
* `credentials_env_vars` - List of environment variables that contain credentials for this operation  
  * `base_var_name` - Base name of the environment variables. Variables are named by appending '_USER', '_PASSWORD', '_PUBKEY' and '_PRIVKEY' to this base name, respectively. Variables whose values are not entered or present in the type of credential or vault selected will not be set.  
  * `password` - Password to assign to the environment variables.  
  * `vault` - The name or reference of the vault to assign to the environment variables.  
  * `hashicorp_vault_engine` - Vault engine name where the credential is stored.  
  * `hashicorp_vault_secret_path` - Path in the vault engine where the credential is stored.  
  * `hashicorp_vault_username_key` - Hashicorp vault key for the username in the key-value store.  
  * `hashicorp_vault_secret_key` - Hashicorp vault key for the password in the key-value store.  
  * `azure_vault_name` - Azure key vault name.  
  * `azure_vault_username_key` - Azure vault key in the key-value store.  
  * `azure_vault_secret_key` - Azure vault key in the key-value store.  
  * `cyberark_vault_query_string` - Query to find a credential in the CyberArk vault. 

### Initial Ingestion and Snapshot [Deprecated] 
The following arguments enable the user to control how the first ingestion and snapshot (SnapSync) should be taken.   
* `skip_wait_for_snapshot_creation` - In DCT v2025.1, waiting for Ingestion and Snapshotting (aka SnapSync) to complete is default functionality. Therefore, these the arguments skip_wait_for_snapshot_creation and wait_time are ignored. In future versions of the provider, we will look at re-implementing the skip SnapSync behavior  
* `wait_time` - In DCT v2025.1, waiting for Ingestion and Snapshotting (aka SnapSync) to complete is default functionality. Therefore, these the arguments skip_wait_for_snapshot_creation and wait_time are ignored. In future versions of the provider, we will look at re-implementing the skip SnapSync behavior. 

### Advanced 
The following arguments apply to all dSources but they are not often necessary for simple sources. 
* `make_current_account_owner` - (Optional) Indicates whether the account creating this reporting schedule must be configured as owner of the reporting schedule. Default: `true`. [Updatable] 
* `tags` - The tags to be created for dSource. This is a map of two parameters: 
  * `key` - (Required) Key of the tag 
  * `value` - (Required) Value of the tag 
* `ignore_tag_changes` – (Optional) Whether changes in the tags are identified by Terraform. Default: `true` (changes to tags are ignored).

## Timeout Configuration

The `timeouts` block is a Terraform meta-argument that's handled specially by Terraform itself and should not be treated as a regular resource attribute. It's used to configure operation timeouts but doesn't represent actual infrastructure state.

The AppData dSource resource supports customizable timeouts for create, update, and delete operations:

```hcl
resource "delphix_appdata_dsource" "example" {
  # ... other configuration ...
  
  timeouts {
    create = "20m"  # Default: 20 minutes
    update = "20m"  # Default: 20 minutes
    delete = "20m"  # Default: 20 minutes
  }
}
```

If an operation exceeds the configured timeout:
- For CREATE operations: The resource will not be added to Terraform state. Check DCT UI to verify if the dSource was created, then import it if necessary.
- For UPDATE operations: Changes may be partially applied. Verify the dSource state in DCT UI.
- For DELETE operations: The resource may still exist in DCT. Verify and manually delete if necessary.

## Import
Use the import block to add Appdata dSources created directly in DCT into a Terraform state file.  

For example:  
```hcl
import {    
    to = delphix_appdata_dsource.dsrc_import_demo 
    id = "dsource_id"    
}   
```

## Limitations 
Not all properties can be updated using the update command. Attempts to update an unsupported property will return a runtime error message. 