/**
* Summary: This template showcases the properties available when provisioning a SAP Hana database from a DCT bookmark.
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

resource "delphix_vdb" "example" {
  bookmark_id             = ""
  name                    = "vdb_to_be_created"
  source_data_id          = "dsource-name"
  vdb_restart             = true
  environment_id          = "env-name"
  environment_user_id     = "environment_user_name"
  target_group_id         = "group-123"
  snapshot_policy_id      = "test_snapshot_policy"
  database_name           = "dbname_to_be_created"
  mount_point             = "/var/mnt"
  auto_select_repository  = true
  retention_policy_id     = "test_retention_policy"
  custom_env_files = ["/export/home/env_file_1"]
  custom_env_vars = {
    MY_ENV_VAR1 = "$HOME"
    MY_ENV_VAR2 = "$CRS_HOME/after"
  }
  repository_id = ""
  appdata_source_params = jsonencode({
    mountLocation    = "/mnt/bkmrk"
    configServices = [{
        "sTenantServicePort" =  "indexserver:30049"
      },
      {
        "sTenantServicePort" =  "xsengine:30052"
      }]
    tDatabaseName   = "tfbkmrk"
    tSystemUserName = "SYSTEM"
    tSystemUserPassword ="Delphix_123"
  })
  config_params jsonencode({
    processes = 150
  })
  appdata_config_params jsonencode({
    param = "value"
  })
  additional_mount_points = [{
      shared_path = "/",
      mount_path = "/work",
      environment_id = "environment-123"
  }]
  post_snapshot {
    name            = "string"
    command         = "string"
    shell           = "bash"
    element_id      = "string"
    has_credentials = true
  }
  pre_snapshot {
    name            = "string"
    command         = "string"
    shell           = "bash"
    element_id      = "string"
    has_credentials = true
  }
  pre_stop {
    name            = "string"
    command         = "string"
    shell           = "bash"
    element_id      = "string"
    has_credentials = true
  }
  configure_clone {
    name            = "string"
    command         = "string"
    shell           = "bash"
    element_id      = "string"
    has_credentials = true
  }
  post_refresh {
    name            = "string"
    command         = "string"
    shell           = "bash"
    element_id      = "string"
    has_credentials = true
  }
  post_stop {
    name            = "string"
    command         = "string"
    shell           = "bash"
    element_id      = "string"
    has_credentials = true
  }
  post_rollback {
    name            = "string"
    command         = "string"
    shell           = "bash"
    element_id      = "string"
    has_credentials = true
  }
  post_start {
    name            = "string"
    command         = "string"
    shell           = "bash"
    element_id      = "string"
    has_credentials = true
  }
  pre_rollback {
    name            = "string"
    command         = "string"
    shell           = "bash"
    element_id      = "string"
    has_credentials = true
  }
  pre_start {
    name            = "string"
    command         = "string"
    shell           = "bash"
    element_id      = "string"
    has_credentials = true
  }
  pre_refresh {
    name            = "string"
    command         = "string"
    shell           = "bash"
    element_id      = "string"
    has_credentials = true
  }
  make_current_account_owner = true
  tags {
    key   = "key-1"
    value = "value-1"
  }
}