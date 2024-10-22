/**
* Summary: This template showcases the properties available when provisioning an Oracle database from a DCT bookmark.
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
  os_username             = "os-user-x"
  os_password             = "os-password-x"
  vdb_restart             = true
  environment_id          = "oracle-env-name"
  environment_user_id     = "environment_user_name"
  target_group_id         = "group-123"
  open_reset_logs         = true
  archive_log             = true
  online_log_groups       = 2
  snapshot_policy_id      = "test_snapshot_policy"
  unique_name             = "dbdhcp2"
  online_log_size         = 4
  database_name           = "dbname_to_be_created"
  mount_point             = "/var/mnt"
  auto_select_repository  = true
  custom_env_files = ["/export/home/env_file_1"]
  custom_env_vars = {
    MY_ENV_VAR1 = "$ORACLE_HOME"
    MY_ENV_VAR2 = "$CRS_HOME/after"
  }
  file_mapping_rules      = "/datafile/dbdhcp3/oradata/dbdhcp3:/data\n/u03/app/ora11202/product/11.2.0/dbhome_1/dbs/dbv_R2V4.dbf:/data/dbv_R2V4.dbf"
  new_dbid                = true
  cluster_node_ids        = ["ORACLE_CLUSTER_NODE-ID"]
  auxiliary_template_id   = "aux-template-1"
  instance_name    = "dbdhcp2"
  retention_policy_id     = "test_retention_policy"
  template_id             = "template-1"
  repository_id = ""
  listener_ids  = ["id1","id2"]
  cdb_id = ""
  vcdb_name = "" //(MT)
  vcdb_database_name = "" //(MT)
  target_vcdb_tde_keystore_path  = "" //(MT)
  parent_tde_keystore_password = "" //(MT)
  cdb_tde_keystore_password = "" //(MT)
  tde_key_identifier = "" //(MT)
  vcdb_tde_key_identifier = "" //(MT)
  parent_tde_keystore_path = "" //(MT)
  tde_exported_key_file_secret = "" //(MT)
  oracle_rac_custom_env_vars = [{
      node_id = "ORACLE_CLUSTER_NODE-1",
      name = "MY_ENV_VAR1",
      value = "$CRS_HOME/after"
    }] //(RAC)
  oracle_rac_custom_env_files = [
    {
      node_id = "ORACLE_CLUSTER_NODE-1",
      path_parameters = "/export/home/env_file_1"
    }] //(RAC)
  config_params jsonencode({
    processes = 150
  })
  pre_start {
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
  post_start {
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
  post_snapshot {
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
  pre_snapshot {
    name            = "string"
    command         = "string"
    shell           = "bash"
    element_id      = "string"
    has_credentials = true
  }
  tags {
    key   = "key-1"
    value = "value-1"
  }
  make_current_account_owner = true
}