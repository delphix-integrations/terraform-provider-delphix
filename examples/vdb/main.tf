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
  key = "36.VUKgUKpqLrLSWUHKHMbLzgEwgnZR69WZUSgXqLk0n1ADHIWiHuHQTunCVfy2KTPu"
  host = "localhost"
}

resource "delphix_vdb" "vdb_name" {
  auto_select_repository = true
  source_data_id         = "38-MSSQL_DB_CONTAINER-32"
}

resource "delphix_vdb_group" "vdb_group_name" {
  name  = "my vdb group"
  vdb_ids = [delphix_vdb.vdb_name.id]
}