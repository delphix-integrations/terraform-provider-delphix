/**
* Summary: This template showcases the properties available when creating an source.
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

resource "delphix_database_postgresql" "source" {
  name              = "test"
  repository_value  = "REPO-1"
  engine_value      = "2"
  environment_value = "ENV-1"
}

