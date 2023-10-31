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



resource "delphix_oracle_dsource" "test_oracle_dsource" {
  name                       = "test2"
  source_value               = "DBOMSRB331B3"
  group_id                   = "3-GROUP-1"
  log_sync_enabled           = false
  make_current_account_owner = true
  environment_user_id        = "HOST_USER-1"
  rman_channels              = 2
  files_per_set              = 5
  check_logical              = false
  encrypted_linking_enabled  = false
  compressed_linking_enabled = true
  bandwidth_limit            = 0
  number_of_connections      = 1
  diagnose_no_logging_faults = true
  pre_provisioning_enabled   = false
  link_now                   = true
  force_full_backup          = false
  double_sync                = false
  skip_space_check           = false
  do_not_resume              = false
  files_for_full_backup      = []
  log_sync_mode              = "UNDEFINED"
  log_sync_interval          = 5
}


