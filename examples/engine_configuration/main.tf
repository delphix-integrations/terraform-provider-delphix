/**
* Summary: This template showcases the properties available when creating an source.
*/

terraform {
  required_providers {
    delphix = {
      version = "4.2.0"
      source  = "delphix-integrations/delphix"
    }
  }
}

provider "delphix" {
  tls_insecure_skip = true
  key               = "1.XXXX"
  host              = "HOSTNAME"
}

/* BLOCK STORAGE */
resource "delphix_engine_configuration" "config" {
  engine_host  = "http://eg22.dlpxdc.co"
  api_version  = "1.11.X"
  sys_user     = "xxx"
  sys_password = "xxx"
  sys_new_password = "xxx"
  user         = "xxx"
  password     = "xxx"
  email        = "noreply@delphix.com"
  engine_type  = "CD"
  device_type = "BLOCK"
}

/* BLOCK STORAGE With NTP configuration */
resource "delphix_engine_configuration" "config" {
  engine_host  = "http://eg22.dlpxdc.co"
  sys_user     = "XXXX"
  sys_password = "XXXX"
  sys_new_password = "xxx"
  user         = "XXXX"
  password     = "XXXX"
  email        = "noreply@delphix.com"
  engine_type  = "CD"
  device_type = "BLOCK"
  ntp_timezone = "America/Anchorage"
  ntp_servers = ["Europe.pool.ntp.org"]
}

/* OBJECT STORAGE with ROLE based configurations (AWS)*/
resource "delphix_engine_configuration" "config2" {
  engine_host  = "http://object.dlpxdc.co"
  sys_user     = "XXXX"
  sys_password = "XXXX"
  sys_new_password = "xxx"
  user         = "XXXX"
  password     = "XXXX"
  email        = "no-reply@delphix.com"
  engine_type  = "CD"
  device_type = "OBJECT"
  object_storage_params {
    cloud_provider = "AWS"
    auth_type = "ROLE"
    region = "<region>"
    bucket = "<bucket_name>"
    endpoint = "<endpoint>"
    size = "<size>"
  }
  ntp_timezone = "Africa/Asmera"
  ntp_servers = ["Europe.pool.ntp.org"]
}

/* OBJECT STORAGE with ACCESS_KEY based configuration (AWS)*/
resource "delphix_engine_configuration" "config2" {
  engine_host  = "http://object.dlpxdc.co"
  sys_user     = "XXXX"
  sys_password = "XXXX"
  sys_new_password = "xxx"
  user         = "XXXX"
  password     = "XXXX"
  email        = "no-reply@delphix.com"
  engine_type  = "CD"
  device_type = "OBJECT"
  object_storage_params {
    cloud_provider = "<AWS>"
    auth_type = "ACCESS_KEY"
    region = "us-west-2"
    bucket = "<bucket_name>"
    endpoint = "<endpoint>"
    size = "<size>"
    access_id = "XXXX"
    access_key = "XXXX" 
  }
  ntp_timezone = "Africa/Asmera"
  ntp_servers = ["Europe.pool.ntp.org"]
}

/* OBJECT STORAGE with MANAGED_IDENTITIED based configuration (AZURE)*/
resource "delphix_engine_configuration" "config2" {
  engine_host  = "http://object.dlpxdc.co"
  sys_user     = "XXXX"
  sys_password = "XXXX"
  sys_new_password = "xxx"
  user         = "XXXX"
  password     = "XXXX"
  email        = "no-reply@delphix.com"
  engine_type  = "CD"
  device_type = "OBJECT"
  object_storage_params {
    cloud_provider = "AZURE"
    auth_type = "MANAGED_IDENTITIES"
    azure_container = "<azure_container>"
    azure_account = "<azure_account>"
    size = "<size>"
  }
  ntp_timezone = "Africa/Asmera"
  ntp_servers = ["Europe.pool.ntp.org"]
}

/*SMTP, NTP, DNS, WEB PROXY, USER ANALYTICS, PHONEHOME CONFIGS*/
resource "delphix_engine_configuration" "config2" {
  engine_host  = "http://object.dlpxdc.co"
  sys_user     = "XXXX"
  sys_password = "XXXX"
  sys_new_password = "xxx"
  user         = "XXXX"
  password     = "XXXX"
  email        = "no-reply@delphix.com"
  engine_type  = "CD"
  device_type = "OBJECT"
  object_storage_params {
    cloud_provider = "AWS"
    auth_type = "ROLE"
    region = "<region>"
    bucket = "<bucket_name>"
    endpoint = "<endpoint>"
    size = "<size>"
  }
  ntp_timezone = "Africa/Asmera"
  ntp_servers = ["Europe.pool.ntp.org"]
  smtp_config {
    server = "<smtp_server>"
    port = 25
    from_email_address = "noreply@perforce.com"
    send_timeout = 80
    tls_authentication = true
  }
  dns_config {
    servers = ["<dns-server1>","<dns-server2>"]
    domains = ["<domain1>","<domain2>"]
  }
  phone_home_enabled = true
  
  web_proxy_config {
    host = "<web_proxy_host>"
    port = 8081
  }
  user_analytics_enabled = true
}

/* SSO Config */
resource "delphix_engine_configuration" "config2" {
  engine_host  = "http://object.dlpxdc.co"
  sys_user     = "XXXX"
  sys_password = "XXXX"
  sys_new_password = "xxx"
  user         = "XXXX"
  password     = "XXXX"
  email        = "no-reply@delphix.com"
  engine_type  = "CD"
  device_type = "OBJECT"
  object_storage_params {
    cloud_provider = "AWS"
    auth_type = "ROLE"
    region = "<region>"
    bucket = "<bucket_name>"
    endpoint = "<endpoint>"
    size = "<size>"
  }

  sso_config {
    enabled=true
    response_skew_time=120
    max_authentication_age=86400
    saml_metadata = <<EOF
    <YOUR-SSO-XML-METADATA>
    EOF
   }
}