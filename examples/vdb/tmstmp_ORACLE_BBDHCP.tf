terraform {
  required_providers {
    delphix = {
      version = "0.0-dev"
      source  = "delphix.com/local/dct"
    }
  }
}

provider "delphix" {
  tls_insecure_skip = true
  key               = "xxx"
  host              = "xxx"
}

resource "delphix_vdb" "oracle_tmstmp_bbdhcp" {
  provision_type  = "timestamp"
  source_data_id  = "3-ORACLE_DB_CONTAINER-1"
  engine_id       = 3
  #target_group_id = "foo"
  vdb_name        = "vdb1"
  database_name   = "vdb1"
  #truncate_log_on_checkpoint = true #ONLY ON ASE
  username = "bogus_user"
  password = "bogus_pass"
  environment_id = "UNIX_HOST_ENVIRONMENT-2"
  environment_user_id = "ora11107"
  #repository_id
  auto_select_repository = true

  pre_refresh {
    name    = "1"
    command = "touch pre_refresh.1"
  }

  pre_refresh {
    name    = "2"
    command = "touch pre_refresh.2"
  }

  post_refresh {
    name    = "1"
    command = "touch post_refresh.1"
  }
  post_refresh {
    name    = "2"
    command = "touch post_refresh.2"
  }

  pre_rollback {
    name    = "1"
    command = "touch pre_rollback.1"
  }
  pre_rollback {
    name    = "2"
    command = "touch pre_rollback.2"
  }

  post_rollback {
    name    = "1"
    command = "touch post_rollback.1"
  }
  post_rollback {
    name    = "2"
    command = "touch post_rollback.2"
  }

  configure_clone {
    name    = "1"
    command = "touch configure_clone.1"
  }
  configure_clone {
    name    = "2"
    command = "touch configure_clone.2"
  }

  pre_snapshot {
    name    = "1"
    command = "touch pre_snapshot.1"
  }
  pre_snapshot {
    name    = "2"
    command = "touch pre_snapshot.2"
  }

  post_snapshot {
    name    = "1"
    command = "touch post_snapshot.1"
  }
  post_snapshot {
    name    = "2"
    command = "touch post_snapshot.2"
  }

  pre_start {
    name    = "1"
    command = "touch pre_start.1"
  }
  pre_start {
    name    = "2"
    command = "touch pre_start.2"
  }

  post_start {
    name    = "1"
    command = "touch post_start.1"
  }
  post_start {
    name    = "2"
    command = "touch post_start.2"
  }

  pre_stop {
    name    = "1"
    command = "touch pre_stop.1"
  }
  pre_stop {
    name    = "2"
    command = "touch pre_stop.2"
  }

  post_stop {
    name    = "1"
    command = "touch post_stop.1"
  }
  post_stop {
    name    = "2"
    command = "touch post_stop.2"
  }

  vdb_restart = true
  #template_id = "DATABASE_TEMPLATE-1"
  #file_mapping_rules
  #oracle_instance_name
  unique_name     = "unique"
  #mount_point     = "/var/mnt"
  open_reset_logs = true
  #snapshot_policy_id
  #retention_policy_id
  # recovery_model = "FULL" only on mssql
  # pre_script = "/var/prs.sh"
  # post_script = "/var/pos.sh"
  # cdc_on_provision = true only  on mssql
  online_log_size   = 4
  online_log_groups = 2
  archive_log       = true
  new_dbid          = false
  listener_ids = ["ORACLE_NODE_LISTENER-2"]
  custom_env_vars = {
    MY_ENV_VAR1 = "$ORACLE_HOME"
  }
  timestamp = "2022-03-23T13:51:17Z"
}
