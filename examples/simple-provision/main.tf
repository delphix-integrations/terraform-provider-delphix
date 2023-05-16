# Configure the connection to Data Control Tower
provider "delphix" {
  host = var.dct_hostname
  key = var.dct_api_key
  tls_insecure_skip = true
}

# Provision a VDB 1
resource "delphix_vdb" "provision_vdb_1" {
  name                   = "tfmtest1"
  source_data_id         = var.source_data_id_1
  auto_select_repository = true
}

# Create a VDB Group with VDB 1
resource "delphix_vdb_group" "create_vdb_group" {
  name  = "Terraform Demo Group"
  vdb_ids = [
    delphix_vdb.provision_vdb_1.id
  ]
}