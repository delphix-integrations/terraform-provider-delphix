/**
* Summary: This template showcases the properties available when uploading a database plugin.
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

resource "delphix_database_plugin" "plugin_upload" {
  engine_host = "ENGINE-HOST"
  file_path = "PATH-TO-PLUGIN-JSON"
}
