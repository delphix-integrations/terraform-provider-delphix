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

resource "delphix_engine_registration" "register" {
  hostname     = "eg21.dlpxdc.co"
  name         = "test_tf"
  username     = "xxx"
  password     = "xxx"
  insecure_ssl = true
}
