# Resource: <resource name> delphix_appdata_dsource

In Delphix terminology, a dSource is a database that the Delphix Continuous Data Engine uses to create and update virtual copies of your database. 
A dSource is created and managed by the Delphix Continuous Data Engine.

The Appdata dSource resource allows Terraform to create and delete AppData dSources. This specifically enables the apply and destroy Terraform commands. Modification of existing dSource resources via the apply command is not supported. All supported parameters are listed below.

## System Requirements

* Data Control Tower v10.0.1+ is required for dSource management. Lower versions are not supported.
* This Appdata dSource Resource only supports Appdata based datasource's , such as POSTGRES,SAP HANA, IBM Db2, etc.The below examples are shown from the PostgreSQL context. See the Oracle dSource Resource for the support of Oracle. The Delphix Provider does not support Oracle, SQL Server, or SAP ASE.

## Upgrade Guide
* Any new dSource created post Version>=3.2.1 can set `wait_time` to wait for snapshot creation , dSources created prior to this version will not support this capability 

## Note
* `status` and `enabled` are subject to change in the tfstate file based on the dSource state.

## Example Usage

The linking of a dSource can be configured through various ingestion approaches. Each configuration is customized to the connector and its supported options. The three PostgreSQL parameter sets below show working examples.

```hcl
# Link dSource using external backup. 

resource "delphix_appdata_dsource" "dsource_name" {
  source_value                  = SOURCE_VALUE
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
  source_value                  = SOURCE_VALUE
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
  source_value                  = SOURCE_VALUE
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

* `source_value` - (Required) Id or Name of the source to link.

* `group_id` - (Required)  Id of the dataset group where this dSource should belong to.

* `log_sync_enabled` - (Required) True if LogSync should run for this database.

* `make_current_account_owner` - (Required) Whether the account creating this reporting schedule must be configured as owner of the reporting schedule.

* `description` - The notes/description for the dSource.

* `link_type` - (Required) The type of link to create. Default is AppDataDirect.
    * `AppDataDirect` - Represents the AppData specific parameters of a link request for a source directly replicated into the Delphix Engine.
    * `AppDataStaged` - Represents the AppData specific parameters of a link request for a source with a staging source.

* `name` - The unique name of the dSource. If unset, a name is randomly generated.

* `staging_mount_base` - The base mount point for the NFS mount on the staging environment [AppDataStaged only].

* `environment_user` - (Required) The OS user to use for linking.

* `staging_environment` - (Required) The environment used as an intermediate stage to pull data into Delphix [AppDataStaged only].

* `staging_environment_user` - The environment user used to access the staging environment [AppDataStaged only].

* `tags` - The tags to be created for dSource. This is a map of 2 parameters:
    * `key` - (Required) Key of the tag
    * `value` - (Required) Value of the tag

* `ops_pre_sync` - Operations to perform before syncing the created dSource. These operations can quiesce any data prior to syncing
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
    
* `ops_post_sync` - Operations to perform after syncing a created dSource.
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

* `excludes` - List of subdirectories in the source to exclude when syncing data.These paths are relative to the root of the source directory. [AppDataDirect only]

* `follow_symlinks` - List of symlinks in the source to follow when syncing data.These paths are relative to the root of the source directory. All other symlinks are preserved. [AppDataDirect only]

* `parameters` - The JSON payload is based on the type of dSource being created. Different data sources require different parameters.

* `sync_parameters` - The JSON payload conforming to the snapshot parameters definition in a LUA toolkit or platform plugin.

* `skip_wait_for_snapshot_creation` - By default this resource will wait for a snapshot to be created post-dSource creation. This ensure a snapshot is available during the VDB provisioning. This behavior can be skipped by setting this parameter to `true`.

* `wait_time` - By default this resource waits 0 minutes for a snapshot to be created. Increase the integer value as needed for larger dSource snapshots. This parameter can be ignored if 'skip_wait_for_snapshot_creation' is set to `true`.
