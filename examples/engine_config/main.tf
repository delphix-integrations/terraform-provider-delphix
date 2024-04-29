/**
* Summary: This template showcases the properties available when creating an source.
*/

terraform {
  required_providers {
    delphix = {
      version = "3.3.0-beta"
      source  = "delphix.com/dct/delphix"
    }
  }
}

provider "delphix" {
  tls_insecure_skip = true
  key               = "1.jTElhpXIao7pTNzVCYdkj1HpGXriTBlYbPha1Di8HjvMF6nESA1crkGlljowDs7y"
  host              = "ubuntu-2-uv49-qar-125346-27a4593a.dlpxdc.co"
}

resource "delphix_engine_configuration" "config" {
  engine_host  = "http://eg22.dlpxdc.co"
  api_version  = "1.11.31"
  sys_user     = "xxx"
  sys_password = "xxx"
  user         = "xxx"
  password     = "xxx"
  email        = "noreply@delphix.com"
  engine_type  = "CD"
}

