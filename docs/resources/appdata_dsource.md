# Resource: <resource name> delphix_appdata_dsource

In Delphix terminology, a dSource is a database that the Delphix Continuous Data Engine uses to create and update virtual copies of your database. 
A dSource is created and managed by the Delphix Continuous Data Engine.


The appdata dSource resource allows Terraform to CREATE,READ and DELETE dSources. 
Modification of existing appdata dSource resources is not supported. All supported parameters are listed below

## Example Usage
Appdata dSource linking can be done in 3 methods,the parameters for these methods wary based on the linking mechanism to be used. 

```hcl
# Link dSource using external backup. 

resource "delphix_appdata_dsource" "dsource_name" {
  source_id                  = SOURCE_ID
  group_id                   = GROUP_ID
  log_sync_enabled           = false
  make_current_account_owner = true
  link_type                  = LINK_TYPE
  name                       = DSOURCE_NAME
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
}

# Link dSource using Delphix Initiated Backup.

resource "delphix_appdata_dsource" "dsource_name" {
  source_id                  = SOURCE_ID
  group_id                   = GROUP_ID
  log_sync_enabled           = false
  make_current_account_owner = true
  link_type                  = LINK_TYPE
  name                       = DSOURCE_NAME
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
}

# Link dSource using Single Database Ingestion.

resource "delphix_appdata_dsource" "dsource_name" {
  source_id                  = SOURCE_ID
  group_id                   = GROUP_ID
  log_sync_enabled           = false
  make_current_account_owner = true
  link_type                  = LINK_TYPE
  name                       = DSOURCE_NAME
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
}
```

## Argument Reference

* `source_id` - (Required) Id of the source to link.

* `group_id` - (Required)  Id of the dataset group where this dSource should belong to.

* `log_sync_enabled` - (Required) True if LogSync should run for this database.

* `make_current_account_owner` - (Required) Whether the account creating this reporting schedule must be configured as owner of the reporting schedule.

* `description` - (Optional) The notes/description for the dSource.

* `link_type` - (Required) The type of link to create. Default is AppDataDirect.
    * `AppDataDirect` - Represents the AppData specific parameters of a link request for a source directly replicated into the Delphix Engine.
    * `AppDataStaged` - Represents the AppData specific parameters of a link request for a source with a staging source.

* `name` - (Optional) The unique name of the dSource. If unset, a name is randomly generated.

* `staging_mount_base` - (Optional) The base mount point for the NFS mount on the staging environment [AppDataStaged only].

* `environment_user` - (Required) The OS user to use for linking.

* `staging_environment` - (Required) The environment used as an intermediate stage to pull data into Delphix [AppDataStaged only].

* `staging_environment_user` - (Optional) The environment user used to access the staging environment [AppDataStaged only].

* `tags` - (Optional) The tags to be created for dSource. This is a map of 2 parameters:
    * `key` - (Required) Key of the tag
    * `value` - (Required) Value of the tag

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

* `excludes` - (Optional) List of subdirectories in the source to exclude when syncing data.These paths are relative to the root of the source directory. [AppDataDirect only]

* `follow_symlinks` - (Optional) List of symlinks in the source to follow when syncing data.These paths are relative to the root of the source directory. All other symlinks are preserved. [AppDataDirect only]

* `parameters` - (Optional) The JSON payload conforming to the DraftV4 schema based on the type of application data being manipulated.

* `sync_parameters` - (Optional) The JSON payload conforming to the snapshot parameters definition in a LUA toolkit or platform plugin.