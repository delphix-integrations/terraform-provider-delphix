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
  key = "1.XXXX"
  host = "localhost"
}

resource "delphix_vdb" "vdb_name" {
  auto_select_repository = true
  source_data_id         = "DSOURCE_ID"
}
