terraform {
  required_version = ">=0.13.0"
required_providers {
    dct-goapi = {
      version = "0.1"
      source  = "delphix.com/local/dct-goapi"
    }
  }
}

provider "dct-goapi" {
  # example configuration here
}

resource "dct-goapi_resource" "example" {
  sample_attribute = "foo"
}