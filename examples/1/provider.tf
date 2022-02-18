terraform {
required_providers {
    delphix = {
      version = "0.0-dev"
      source  = "delphix.com/local/delphix"
    }
  }
}

provider "delphix" {
  key = "KEY_GOES_HERE"
  host = "IP_GOES_HERE"
}

resource "delphix_resource" "example" {
  sample_attribute = "foo"
}