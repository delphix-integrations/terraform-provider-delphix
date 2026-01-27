terraform {
  required_providers {
    delphix = {
      version = "4.2.0"
      source  = "delphix.com/dct/delphix"
    }
  }
}

provider "delphix" {
  tls_insecure_skip = true
  key               = "1.XXXX"
  host              = "HOSTNAME"
}

/* Data/Virtualization Engine Registeration with DCT */
resource "delphix_engine_dct_registration" "register" {
  hostname     = "eg21.dlpxdc.co"
  name         = "test_tf"
  username     = "xxx"
  password     = "xxx"
  insecure_ssl = true
  engine_type = "CD"
}

/*Compliance Engine Registeration with DCT*/
resource "delphix_engine_dct_registration" "register" {
  hostname     = "eg22.dlpxdc.co"
  name         = "test_tf"
  username     = "xxx"
  password     = "xxx"
  insecure_ssl = true
  engine_type = "CC"
  compliance_user = "admin"
  compliance_password = "xxx"
}