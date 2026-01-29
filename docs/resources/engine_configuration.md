# Resource: <resource name> delphix_engine_configuration
This resource helps to configure Delphix Engine.
This resource allow you to configure Delphix Engine with several parameters like DNS, SMTP, NTP, etc .
Engine can be either of type Virtualization or Compliance.

## Example Usage

### Configure an Engine with Object store.
```hcl
resource "delphix_engine_configuration" "config" {
  engine_host  = "<full-url-of-engine-host>"
  sys_user     = "sysadmin"
  sys_password = "xxxx"
  sys_new_password = "xxxx"
  compliance_user = "admin"
  compliance_password = "<compliance-current-password>"
  compliance_email = "<email>"
  compliance_new_password = "<compliance-new-password>"
  user         = "admin"
  password     = "<admin-password>"
  email        = "<admin-email>"
  engine_type  = "<CC or CD>"

  device_type = "OBJECT"
  object_storage_params {
    cloud_provider = "AWS"
    auth_type = "ROLE"
    region = "<aws-region>"
    bucket = "<aws-bucket-name>"
    endpoint = "<endpoint>"
    size = "<size>"
  }

  smtp_config {
    server = "<smtp-server>"
    port = <smtp-port>
    from_email_address = "<email>"
    send_timeout = 80
    tls_authentication = true
    smtp_authentication {
        user = "<user>"
        password = "xxxx"
    }
  }

  dns_config {
    domains = ["<domain1>", "<domain2>"]
    servers = ["<server1>", "<sever2>"]
  }

  phone_home_enabled = true

  web_proxy_config {
    host = "<host>"
    port = 8081
    username = "<web-proxy-user>"
    password = "xxxx"
  }

  user_analytics_enabled = true

  sso_config {
     enabled=true
     response_skew_time=120
     max_authentication_age=86400
     saml_metadata = <<EOF <sso-metadata> EOF
   }
}
```

### Configure an Engine with Block store.
```hcl
resource "delphix_engine_configuration" "config" {
  engine_host  = "<full-url-of-engine-host>"
  sys_user     = "sysadmin"
  sys_password = "xxxx"
  sys_new_password = "xxxx"
  user         = "admin"
  password     = "<admin-password>"
  email        = "<admin-email>"
  engine_type  = "<CC or CD>"
  device_type = "BLOCK"
}
```

### Object storage params for AZURE based storage.
```hcl
object_storage_params {
    cloud_provider = "AZURE"
    azure_account = "<azure-account>"
    auth_type = "MANAGED_IDENTITIES"
    size = "<size>"
    azure_container = "<azure-container>"
  }
```

### Object storage params for AWS based storage.
```hcl
object_storage_params {
    cloud_provider = "AWS"
    auth_type = "ROLE"
    region = "<aws-region>"
    bucket = "<aws-bucket-name>"
    endpoint = "<endpoint>"
    size = "<size>"
  }
```

## Argument Reference
The following arguments apply to all configurations.
* `engine_host` - (Required) Full URL of Engine to configure. example - `https://example-engine.dlpxdc.co`.
* `api_version` - Engine API version which is used to create the API session.
The oldest supported engine version is `1.11.40` and is the default. For a specific API version, please refer to [API version mapping for engine](https://help.delphix.com/cd/current/content/api_version_information.htm) [NOTE: Due to a known issue, the API version has to be set to `1.11.47` specifically for engine version `2026.1.0.0`]
* `engine_type ` - (Required) Type of Engine to configure.This can be either `CD` (for Virtualisation Engine ) or `CC` (for Masking Engine).
* `sys_user` - (Required) Name of system administrator user.
* `sys_password ` - (Required) Current password of system administrator.
* `sys_new_password ` - (Required) New password of system administrator.
* `email ` - (Required) Email to be set for system administrator user.
* `compliance_user ` - Name of compliance engine user. Required if `engine_type` is set to `CC`.
* `compliance_password ` - Current password of compliance user. 
* `compliance_new_password ` - New password of compliance user. 
[NOTE] Password must be between 6-12 characters and contain 1 digit, 1 uppercase alphabet character, and 1 special character
* `compliance_email ` - Email to set for compliance user.
* `user ` - (Required) Name of administrator user.
* `password ` - (Required) New password of administrator user.
* `ntp_timezone ` - Timezone for NTP server (Required for Object Storage)
* `ntp_servers ` - List of NTP servers (Required for Object Storage)
* `device_type ` - This is to configure storage. This can be either `BLOCK`(for Block Storage) or `OBJECT`(AWS S3 Object store).
* `object_storage_params ` - Configuration parameters for `OBJECT` storage.
     * `cloud_provider ` - Cloud provider for Object Store. This can be either `AWS` or `AZURE`
     * `auth_type ` - Authentication type for Object Store. This can be either `ROLE` or `ACCESS_KEY` if `cloud_provider` is `AWS` and `MANAGED_IDENTITIES` and `ACCESS_KEY` if `cloud_provider` is `AZURE`.
     * `region` - (Required for `AWS `) Region of the bucket for Object store.
     * `bucket` - (Required for `AWS `) Name of the bucket for Object store.
     * `endpoint ` - (Required for `AWS `) Endpoint for Object store.
     * `size ` - Size of the Object store to configure.
     * `access_id ` - (Required for `AWS `) access id if using `auth_type ` as `ACCESS_KEY`.
     * `access_key ` - (Required for `AWS `) access key if using `auth_type ` as `ACCESS_KEY`.
     * `azure_container ` - (Required for `AZURE `) container name in case of Azure based Object storage.
     * `azure_account ` - (Required for `AZURE `) Azure account in case of Azure based Object storage.
     * `azure_key` - (Required for `AZURE `) access key for Azure if `auth_type` is `ACCESS_KEY`.
   
* `smtp_config ` - Configuration parameters for SMTP. Configure the Delphix Engine's SMTP sending service to enable email notifications. Your sysadmin email will be used for receiving system reports, events, and fault notifications.
       * `server ` - (Required) SMTP server
       * `port ` - (Required) Port number for SMTP.
       * `from_email_address ` - (Required) Email address.
       * `send_timeout ` - Maximum timeout to wait, in seconds, when sending mail. Default is  `60 seconds`.
       * `tls_authentication ` - Boolean flag to enable/disable TLS authentication.
       * `smtp_authentication ` - Authentication for SMTP.
           * `user ` - username for SMTP authentication.
           * `password ` - Password for SMTP authentication.
* `dns_config ` - Configuration parameters for DNS.
     * `servers ` - List of DNS server to add.
     * `domains ` - List of DNS domains to add.
* `phone_home_enabled ` - Boolean flag for Delphix Phone Home. This service will automatically send a minimal support bundle once a day to the Delphix support site over HTTPS. This will help with future support and troubleshooting.
* `web_proxy_config ` - Configuration parameters for Web Proxy. The Web Proxy Server will be used to communicate with Delphix Corp. for support, troubleshooting, upgrades, updates, and patches.
     * `host ` - (Required) Host for Web proxy.
     * `port ` - (Required) Port for Web proxy, default is `8081`.
     * `username ` - Username for  Web proxy.
     * `password` - Password for Web proxy.
* `user_analytics_enabled ` - Boolean flag for setting User click analytics. This service will automatically send a stream of anonymous, non-personal metadata describing user interaction with the product's user interface.
* `sso_config ` - Configuration for SSO based authentication.
     * `enabled ` - (Required) Boolean flag for sso enable/disable.
     * `saml_metadata` - (Required) SSO/SAML metadata.
     * `response_skew_time ` - Maximum time difference allowed between a SAML response and the engine's current time, in seconds. If not set, it defaults to `120 seconds `.
     * `max_authentication_age ` - This defines how far in the past to accept authentications to the identity provider, in seconds. If not set, it defaults to `86,400 seconds `(one day).
    
## Limitations
This resource is not updatable, deletable and importable. Please contact Delphix Support to make a feature request and/or create a GitHub Issue.