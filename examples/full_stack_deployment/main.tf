/**
* Summary: This template showcases a mock example to 
* 1) Provision an Azure VM.
* 2) Create a Target environment from that VM.
* 3) Link and Sync a dSource. Create a new snapshot.
* 4) Provision a new VDB from that Oracle dSource's snapshot.
* *** Warning: This is only an example. It will not work out of the box.***
*/

terraform {
  required_providers {
    delphix = {
      version = ">=3.3.2"
      source  = "delphix-integrations/delphix"
    }
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~>3.0"
    }
  }
}

// *** Requirement***: Various variables used throughout the example.
locals {
  dct-key           = "<1.XXXX>"
  dct-host          = "<DCT HOSTNAME>"
  vm-hostname       = "oracle-linux-host"
  vm-username       = "<USERNAME>"
  vm-password       = "<PASSWORD>"
  source-db-name    = "<SOURCE DATABASE NAME>" // Name of Database dynamically identified on the source environment
  dsource-name      = "full_deploy_dsource"
  vdb-name          = "full_deploy_vdb"
}

provider "delphix" {
  tls_insecure_skip = true
  key               = local.dct-key
  host              = loca.dct-host
}


// *** Requirement ***: This is an example only and will not work without significant modification and additional files.
// See the official documentation here for a full VM deployment: https://learn.microsoft.com/en-us/azure/virtual-machines/linux/quick-create-terraform
// The VM creation terraform resource can be replaced with an equivalent resource GCP, AWS, VMWare, etc that's compatible with Delphix Continuous Data.
// Consult your organization's DevOps expert for guidance on how to provision a VM that's approved for your company.
resource "azurerm_linux_virtual_machine" "azure_vm" {
  name                  = "Delphix Oracle Target"
  location              = azurerm_resource_group.rg.location // Not provided
  resource_group_name   = azurerm_resource_group.rg.name // Not provided
  network_interface_ids = [azurerm_network_interface.my_terraform_nic.id] // Not provided
  size                  = "Standard_DS1_v2"

  os_disk {
    name                 = "myOsDisk"
    caching              = "ReadWrite"
    storage_account_type = "Premium_LRS"
  }

  source_image_reference {
    publisher = "Canonical"
    offer     = "0001-com-ubuntu-server-jammy"
    sku       = "22_04-lts-gen2"
    version   = "latest"
  }

  computer_name  = local.vm-hostname
  admin_username = local.vm-username
  admin_password = local.vm-password // Not secure. Example only

  boot_diagnostics {
    storage_account_uri = azurerm_storage_account.my_storage_account.primary_blob_endpoint // Not provided
  }
}

// Add the Azure VM as a Delphix environment.
resource "delphix_environment" "linux-oracle-target" {
     name = local.vm-hostname
     os_name = "UNIX"
     hostname = local.vm-hostname
     username = local.vm-username
     password = local.vm-password
     engine_id = 1
     toolkit_path = "/home/delphix_os/toolkit"
     description = "This is a unix target for the Oracle VDB."     
 } 

// Link and Sync the dSource and take a new snapshot
// *** Requirement *** This is an Oracle dSource. Updates are likely required.
resource "delphix_oracle_dsource" "full_oracle_dsource" {
  name                       = local.dsource-name
  source_value               = local.source-db-name
  group_id                   = "4-GROUP-1"
  environment_user_id        = "HOST_USER-1"
  log_sync_enabled           = false
  rman_channels              = 2
  files_per_set              = 5
  check_logical              = false
  encrypted_linking_enabled  = false
  compressed_linking_enabled = true
  bandwidth_limit            = 0
  number_of_connections      = 1
  diagnose_no_logging_faults = true
  pre_provisioning_enabled   = false
  link_now                   = true
  force_full_backup          = false
  double_sync                = false
  skip_space_check           = false
  do_not_resume              = false
  files_for_full_backup      = []
  log_sync_mode              = "UNDEFINED"
  log_sync_interval          = 5
  make_current_account_owner = true
}


// Provision by Snapshot the 1 Oracle VDB on the newly created environment
resource "delphix_vdb" "vdb_provision_loop" {
  name                   = local.vdb-name
  source_data_id         = local.dsource-name
  environment_id         = local.vm-hostname
  auto_select_repository = true
}
