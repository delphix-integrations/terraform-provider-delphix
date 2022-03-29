
terraform {
  required_providers {
    delphix = {
      version = "0.0-dev"
      source  = "delphix.com/local/delphix"
    }
  }
}


provider "delphix" {
  tls_insecure_skip = true
  key               = "XXXX"
  host              = "HOST"
}

resource "delphix_vdb" "new" {
  
	auto_select_repository = true
  source_data_id         = "DATASOURCE_ID"
  vdb_name = "VDBNAME"
	vdb_restart = true
}

