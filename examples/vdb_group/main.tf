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


locals {
  vdbs = {
    "vdb4" = { snapshot_id = "6-ORACLE_DB_CONTAINER-21", name = "us4" },
    "vdb5" = { snapshot_id = "6-ORACLE_DB_CONTAINER-23", name = "us5" },
    "vdb1" = { snapshot_id = "6-ORACLE_DB_CONTAINER-7", name = "us1" },
    "vdb2" = { snapshot_id = "6-ORACLE_DB_CONTAINER-1", name = "us2" },
    "vdb3" = { snapshot_id = "6-ORACLE_DB_CONTAINER-5", name = "us3" }
  }
}

resource "delphix_vdb" "example" {
  for_each               = try(local.vdbs, {})
  name                   = each.value.name
  source_data_id         = each.value.snapshot_id
  auto_select_repository = true

}

#sort helps to maintain thr order of the vdbs to avoid erroneous drift
resource "delphix_vdb_group" "this" {
  name    = "random"
  vdb_ids = sort(flatten([for vdb in delphix_vdb.example : vdb.id]))
}
