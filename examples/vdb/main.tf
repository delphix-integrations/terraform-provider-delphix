terraform {
required_providers {
    delphix = {
      version = "0.0-dev"
      source  = "delphix.com/local/delphix"
    }
  }
}

provider "delphix" {
  key = "API_KEY"
  host = "HOST_NAME"
}

resource "delphix_resource_vdb" "test_vdb" {
  provider = delphix
    vdb {
        auto_select_repository = true
        source_data_id = "DSOURCE_ID"
    }
}