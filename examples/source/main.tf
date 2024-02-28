terraform {
  required_providers {
    delphix = {
      version = "2.0.4-beta"
      source  = "delphix.com/local/delphix"
    }
  }
}

provider "delphix" {
  tls_insecure_skip = true
  key               = "1.jTElhpXIao7pTNzVCYdkj1HpGXriTBlYbPha1Di8HjvMF6nESA1crkGlljowDs7y"
  host              = "ubuntu-2-uv49-qar-125346-27a4593a.dlpxdc.co"
}


/* Unix Standalone */
resource "delphix_environment" "unixtgt" {
     engine_id = 13
     os_name = "UNIX"
     username = "postgres"
     password = "postgres"
     hostname = "rhel-86-jpv3-qar-125346-27a4593a.dlpxdc.co"
     toolkit_path = "/tmp"
     name = "unixsrc"
     description = "This is a unix src."     
 } 


locals {
  repository_id = [for r in delphix_environment.unixtgt.repositories : r.id if r.name == "Postgres vFiles (14.1)"]
}

resource "delphix_source" "name" {
  name             = "abc5"
  repository_value = local.repository_id[0]
}

locals {
  split_user_id = split("-", delphix_environment.unixtgt.id)
  last_element = element(local.split_user_id, length(local.split_user_id) - 1)
  env_user = "HOST_USER-${local.last_element}"
  group = "${delphix_environment.unixtgt.engine_id}-GROUP-1"
}


resource "delphix_appdata_dsource" "test_app_data_dsource_second" {
  source_value               = delphix_source.name.id
  group_id                   = local.group
  log_sync_enabled           = false
  make_current_account_owner = true
  link_type                  = "AppDataStaged"
  name                       = "appdata_dsource_second_new"
  staging_mount_base         = ""
  environment_user           = local.env_user
  staging_environment        = delphix_environment.unixtgt.id
  parameters = jsonencode({
    delphixInitiatedBackupFlag : true,
    delphixInitiatedBackup : [
      {
        userName : "delphix",
        postgresSourcePort : 5432,
        userPass : "postgres",
        sourceHostAddress : delphix_source.name.fqdn
      }
    ],
    postgresPort : 5433,
    mountLocation : "/datadrive1/provision/ds-assetmanagement-neur-tntnpk-dev-rocsexecution-1"
  })
  sync_parameters = jsonencode({
    resync = true
  })
  ops_pre_sync {
    name    = "key-1"
    command = "echo \"hello world\""
    shell   = "shell"
    credentials_env_vars {
      base_var_name = "delphix"
      password      = "delphix"
    }
  }
}

# resource "time_sleep" "wait_30_seconds" {
#   create_duration = "30s"
#   depends_on = [delphix_appdata_dsource.test_app_data_dsource_second]
# }


resource "delphix_vdb" "example" {
  #depends_on = [time_sleep.wait_30_seconds]
  name                   = "vdb_to_be_created"
  source_data_id         = delphix_appdata_dsource.test_app_data_dsource_second.id
  vdb_restart            = true
  auto_select_repository = true
  appdata_source_params = jsonencode({
    mountLocation     = "/mnt/GAT"
    postgresPort      = 5434
    configSettingsStg = []
  })
  make_current_account_owner = true
  #time_sleep = [time_sleep.wait_30_seconds]
}