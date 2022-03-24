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

resource "delphix_vdb" "vdb_name" {
  provision_type  = "snapshot"
  source_data_id  = "DBOMSRDC200B-84322"
  engine_id       = 1
  #target_group_id = "foo"
  vdb_name        = "vdbor"
  database_name   = "vdbor"
  #username
  #password
  #environment_id
  #environment_user_id
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
  #file_mapping_rules
  #oracle_instance_name
  unique_name     = "udiunique"
  #mount_point     = "/var/mnt"
  open_reset_logs = true
  #snapshot_policy_id
  #retention_policy_id
  online_log_size   = 4
  online_log_groups = 2
  archive_log       = true
  new_dbid          = true
  custom_env_vars = {
    MY_ENV_VAR1 = "$ORACLE_HOME"
  }
  snapshot_id = "1-ORACLE_SNAPSHOT-1"





  /*
  pre_refresh {
    name    = "n"
    command = "time"
    shell   = "bash"
  }
  pre_refresh {
    name    = "n2"
    command = "time"
    shell   = "bash"
  }
  */
}
/*
resource "delphix_vdb" "vdb_name2" {
  provision_type         = "timestamp"
  auto_select_repository = true
  source_data_id         = "dsource2"
  timestamp              = "2021-05-01T08:51:34.148000+00:00"

  post_refresh {
    name    = "n"
    command = "time"
    shell   = "bash"
  }
  post_refresh {
    name    = "n2"
    command = "time"
    shell   = "bash"
  }
}
*/
