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

#dsource with replication

 resource "delphix_appdata_dsource" "test_app_data_dsource" {
   source_value                  = "1-APPDATA_STAGED_SOURCE_CONFIG-6"
   group_id                   = "1-GROUP-1"
   log_sync_enabled           = false
   make_current_account_owner = true
   link_type                  = "AppDataStaged"
   name                       = "appdata_dsource"
   staging_mount_base         = ""
   environment_user           = "HOST_USER-2"
   staging_environment        = "1-UNIX_HOST_ENVIRONMENT-2"
   parameters = jsonencode({
    dSourceType : "Replication",
    mountPath : "/delphix/zfs3",
    stagingPort : 3310,
    serverId : 106 ,
    sourceip : "10.110.203.176",
    sourceUse : "XXXX",
    sourcePass : "XXXX",
    logSync: true,
    replicationUser: "XXXX",
    replicationPass: "XXXX",
    databaseList: "ALL",
    backupPath: "",
    stagingBasedir: "/usr",
    stagingPass: "XXXX"
   })
   sync_parameters = jsonencode({
     resync = true
   })
 }