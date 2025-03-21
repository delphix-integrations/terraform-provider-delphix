/**
* Summary: This template showcases a simple example to 
* 1) Set local VDB variables
* 2) Provision 5 VDBs using a for_each loop
* 3) Place the 5 VDBs into a single VDB Group using a for loop
*/

terraform {
  required_providers {
    delphix = {
      version = ">=3.3.2"
      source  = "delphix-integrations/delphix"
    }
  }
}

// *** Requirement***: Update the key and host with valid credentials.
provider "delphix" {
  tls_insecure_skip = true
  key               = "1.XXXX"
  host              = "HOSTNAME"
}


// Create variables for the VDBs.
// *** Requirement***: Update the Snapshot ID with a valid Snapshot. The same snapshot can be reused.
locals {
  vdbs = {
    "vdb1" = { snapshot_id = "6-ORACLE_DB_CONTAINER-7",  name = "us1" },
    "vdb2" = { snapshot_id = "6-ORACLE_DB_CONTAINER-1",  name = "us2" },
    "vdb3" = { snapshot_id = "6-ORACLE_DB_CONTAINER-5",  name = "us3" },
    "vdb4" = { snapshot_id = "6-ORACLE_DB_CONTAINER-21", name = "us4" },
    "vdb5" = { snapshot_id = "6-ORACLE_DB_CONTAINER-23", name = "us5" }
  }
}

// Provision by Snapshot the 5 VDBs in a loop.
// Instead of a for_each loop, you could optionally copy this resource 4 more times and replace the values directly.
resource "delphix_vdb" "vdb_provision_loop" {
  for_each               = try(local.vdbs, {})
  name                   = each.value.name
  source_data_id         = each.value.snapshot_id
  auto_select_repository = true

}
// Place the 5 VDBs in a single VDB Group.
// The sort() function helps to maintain the order of the vdbs to avoid erroneous drift.
resource "delphix_vdb_group" "this" {
  name    = "random"
  vdb_ids = sort(flatten([for vdb in delphix_vdb.example : vdb.id]))
}
