terraform {
  required_providers {
    delphix = {
      version = "0.0-dev"
      source  = "delphix.com/dct/delphix"
    }
  }
}

provider "delphix" {
  tls_insecure_skip = true
  key               = "xxx"
  host              = "localhost"
}

resource "delphix_vdb" "tfvora1" {
  provision_type         = "snapshot"
  auto_select_repository = true
  source_data_id         = "orasrc"
  environment_id = "2-UNIX_HOST_ENVIRONMENT-5"

  configure_clone {
    name    = "post-provision-hook"
    command = "echo $1"
    shell   = "bash"
  }
}

/* resource "delphix_vdb" "vdb_name2" {
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
} */
