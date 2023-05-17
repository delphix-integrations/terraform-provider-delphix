###
# VDB Group Information
###
output "vdb_group_id" {
  value = delphix_vdb_group.create_vdb_group.id
}

output "vdb_group_name" {
  value = delphix_vdb_group.create_vdb_group.name
}

###
# VDB 1 Information
###
output "vdb_id_1" {
  value = delphix_vdb.provision_vdb_1.id
}

output "vdb_name_1" {
  value = delphix_vdb.provision_vdb_1.name
}

output "vdb_ip_address_1" {
  value = delphix_vdb.provision_vdb_1.ip_address
}

output "vdb_database_type_1" {
  value = delphix_vdb.provision_vdb_1.database_type
}

output "vdb_database_version_1" {
  value = delphix_vdb.provision_vdb_1.database_version
}
