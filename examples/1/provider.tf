terraform {
required_providers {
    delphix = {
      version = "0.0-dev"
      source  = "delphix.com/local/delphix"
    }
  }
}

provider "delphix" {
  # example configuration here
}

resource "delphix_resource" "example" {
  sample_attribute = "foo"
}