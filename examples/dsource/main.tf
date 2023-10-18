/**
* Summary: This template showcases the properties available when creating an app data dsource.
*/

terraform {
  required_providers {
    delphix = {
      version = "VERSION"
      source  = "delphix-integrations/delphix"
    }
  }
}

provider "delphix" {
  tls_insecure_skip = true
  key               = "1.XXXX"
  host              = "HOSTNAME"
}


# resource "delphix_appdata_dsource" "test_app_data_dsource" {
#   source_name                  = "1-APPDATA_STAGED_SOURCE_CONFIG-6"
#   group_id                   = "1-GROUP-1"
#   log_sync_enabled           = false
#   make_current_account_owner = true
#   link_type                  = "AppDataStaged"
#   name                       = "appdata_dsource"
#   staging_mount_base         = ""
#   environment_user           = "HOST_USER-2"
#   staging_environment        = "1-UNIX_HOST_ENVIRONMENT-2"
#   parameters = jsonencode({
#     externalBackup : [],
#     delphixInitiatedBackupFlag : true,
#     delphixInitiatedBackup : [
#       {
#         userName : "XXXX",
#         postgresSourcePort : XXXX,
#         userPass : "XXXX",
#         sourceHostAddress : "HOSTNAME"
#       }
#     ],
#     singleDatabaseIngestionFlag : false,
#     singleDatabaseIngestion : [],
#     stagingPushFlag : false,
#     postgresPort : XXXX,
#     configSettingsStg : [],
#     mountLocation : "/tmp/delphix_mnt"
#   })
#   sync_parameters = jsonencode({
#     resync = true
#   })
# }

resource "delphix_appdata_dsource" "test_app_data_dsource_second" {
  source_name                = "1-APPDATA_STAGED_SOURCE_CONFIG-7"
  group_id                   = "1-GROUP-1"
  log_sync_enabled           = false
  make_current_account_owner = true
  link_type                  = "AppDataStaged"
  name                       = "appdata_dsource_second"
  staging_mount_base         = ""
  environment_user           = "HOST_USER-2"
  staging_environment        = "1-UNIX_HOST_ENVIRONMENT-2"
  parameters = jsonencode({
    delphixInitiatedBackupFlag : true,
    delphixInitiatedBackup : [
      {
        userName : "XXXX",
        postgresSourcePort : XXXX,
        userPass : "XXXX",
        sourceHostAddress : "HOSTNAME"
      }
    ],
    postgresPort : XXX,
    mountLocation : "/tmp/delphix_mnt_second"
  })
  sync_parameters = jsonencode({
    resync = true
  })
  ops_pre_sync {
    name    = "key-1"
    command = "echo \"hello world\""
    shell   = "shell"
    credentials_env_vars {
      base_var_name = "XXXX"
      password      = "XXXX"
    }
  }
  ops_post_sync {
    name    = "key-2"
    command = "echo \"hello world\""
    shell   = "shell"
    credentials_env_vars {
      base_var_name = "XXXX"
      password      = "XXXX"
    }
  }
}


# Below are the 3 ways to link dsource with params , use any one of them
#  externalBackup: [
#             {
#                 keepStagingInSync: false,
#                 backupPath: "/var/tmp/backup",
#                 walLogPath: "/var/tmp/backup"
#             }
# ]

# singleDatabaseIngestion: [
#             {
#                 databaseUserName: "postgres",
#                 sourcePort: 5432,
#                 dumpJobs: 2,
#                 restoreJobs: 2,
#                 databaseName: "abcd",
#                 databaseUserPassword: "xxxx",
#                 dumpDir: "abcd",
#                 sourceHost: "abcd",
#                 postgresqlFile: "abcd"
#             }
#         ]

# delphixInitiatedBackup : [
#   {
#     userName : "XXXX",
#     postgresSourcePort : XXXX,
#     userPass : "XXXX",
#     sourceHostAddress : "HOSTNAME"
#   }
# ]
