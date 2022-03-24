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

resource "delphix_vdb" "mssql_tmstmp_ac2" {
  provision_type  = "timestamp"
  source_data_id  = "winsrc1"
  engine_id       = 5
  target_group_id = "GROUP-2"
  vdb_name        = "vdbms2"
  database_name   = "vdbms2"
  #truncate_log_on_checkpoint = true #ONLY ON ASE
  #username = "bogus_user"
  #password = ""
  environment_id = "5-WINDOWS_HOST_ENVIRONMENT-1"
  # environment_user_id = "QA-AD\\delphix"
  #repository_id
  auto_select_repository = true

  # pre_refresh {
  #   name    = "1"
  #   command = "touch pre_refresh.1"
  # }

  # pre_refresh {
  #   name    = "2"
  #   command = "touch pre_refresh.2"
  # }

  # post_refresh {
  #   name    = "1"
  #   command = "touch post_refresh.1"
  # }
  # post_refresh {
  #   name    = "2"
  #   command = "touch post_refresh.2"
  # }

  # pre_rollback {
  #   name    = "1"
  #   command = "touch pre_rollback.1"
  # }
  # pre_rollback {
  #   name    = "2"
  #   command = "touch pre_rollback.2"
  # }

  # post_rollback {
  #   name    = "1"
  #   command = "touch post_rollback.1"
  # }
  # post_rollback {
  #   name    = "2"
  #   command = "touch post_rollback.2"
  # }

  # configure_clone {
  #   name    = "1"
  #   command = "touch configure_clone.1"
  # }
  # configure_clone {
  #   name    = "2"
  #   command = "touch configure_clone.2"
  # }

  # pre_snapshot {
  #   name    = "1"
  #   command = "touch pre_snapshot.1"
  # }
  # pre_snapshot {
  #   name    = "2"
  #   command = "touch pre_snapshot.2"
  # }

  # post_snapshot {
  #   name    = "1"
  #   command = "touch post_snapshot.1"
  # }
  # post_snapshot {
  #   name    = "2"
  #   command = "touch post_snapshot.2"
  # }

  # pre_start {
  #   name    = "1"
  #   command = "touch pre_start.1"
  # }
  # pre_start {
  #   name    = "2"
  #   command = "touch pre_start.2"
  # }

  # post_start {
  #   name    = "1"
  #   command = "touch post_start.1"
  # }
  # post_start {
  #   name    = "2"
  #   command = "touch post_start.2"
  # }

  # pre_stop {
  #   name    = "1"
  #   command = "touch pre_stop.1"
  # }
  # pre_stop {
  #   name    = "2"
  #   command = "touch pre_stop.2"
  # }

  # post_stop {
  #   name    = "1"
  #   command = "touch post_stop.1"
  # }
  # post_stop {
  #   name    = "2"
  #   command = "touch post_stop.2"
  # }

  vdb_restart = true
  # template_id = "DATABASE_TEMPLATE-1"
  #file_mapping_rules
  #oracle_instance_name
  # unique_name     = "udiunique"
  # mount_point     = "/var/mnt"
  # open_reset_logs = true
  #snapshot_policy_id = "testsnap"
  #retention_policy_id = "RetPolicy"
  recovery_model = "FULL"
  pre_script = "C:\\Program Files\\Delphix\\scripts\\myscript.ps1"
  post_script = "C:\\Program Files\\Delphix\\scripts\\myscript.ps1"
  cdc_on_provision = true
  # online_log_size   = 4
  # online_log_groups = 2
  # archive_log       = true
  # new_dbid          = false
  #listener_ids = ["ORACLE_NODE_LISTENER-13"]
  #custom_env_vars = {
  #  MY_ENV_VAR1 = "$ORACLE_HOME"
  #}
  #timestamp_in_database_timezone = "2018-05-23T21:48:15.000"
  timestamp = "2021-11-29T23:59:46Z"
}
