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
  key               = "1.XXXX"
  host              = "127.0.0.1"
}

resource "delphix_vdb" "vdb_name" {
  provision_type         = "snapshot"
  auto_select_repository = true
  source_data_id         = "dsource"

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
}

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
