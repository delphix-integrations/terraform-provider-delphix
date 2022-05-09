terraform {
  required_providers {
    delphix = {
      version = "1.0-beta"
      source  = "delphix.com/dct/delphix"
    }
  }
}

provider "delphix" {
  tls_insecure_skip = true
  key               = "1.089N1yoNUoHis8cJqA6BHRaf5OnS1HWmDAlQgxzNmhxNESamd9T4CNuLyvjw8eVF"
  host              = "localhost"
}

resource "delphix_vdb" "vdb_name2" {
  provision_type         = "snapshot"
  auto_select_repository = true
  source_data_id         = "dsource"
  vdb_name = "vdb-from-tf2"
  oracle_instance_name = "vdbtf"
  database_name = "vdbtf"
  unique_name = "vdbtf"
  db_username = "oracle"
  db_password = "oracle"
  template_id = "myTemplate"
  listener_ids = ["LISTENER"]
}