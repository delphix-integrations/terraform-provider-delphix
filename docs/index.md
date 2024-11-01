# <provider> Delphix Provider

The Terraform Provider for Delphix enables you to natively manage data-as-code along with your infrastructure.

With Terraform and Delphix, you can now automatically provision, manage, and teardown any number of ephemeral data environments to drive enterprise DevOps workflows, such as test data management.

This provider communicates directly with Data Control Tower (DCT) to generated virtual database and other objects. Therefore, DCT must be registered with one or more Delphix Continuous Data Engines.

To learn more about Delphix and DCT APIs, refer to [Delphix Documentation](https://documentation.delphix.com/docs/) and [DCT Documentation](https://dct.delphix.com/docs/latest/) respectively. Please [Contact us](ask-integrations@delphix.com) (ask-integrations@delphix.com) with any questions. 

If you are entitled to Data Control Tower then you may also send support issues through the [Delphix Support Portal](https://support.delphix.com/).

## System Requirements

| Product                        | Version  |
|--------------------------------|----------|
| Data Control Tower (DCT)       | v22+ |
| Delphix Continuous Data Engine | v16.0.0.0+ |

Note: The DCT version above guarantees full provider support. However, each  resource might support older versions. Refer to the specific resource documentation page for more information.

## Release Notes

The Delphix Provider's complete release notes can be found in the [Delphix Ecosystem Documentation](https://ecosystem.delphix.com/docs/main/release-notes-terraform).

## Connectivity and Authentication

All communication is performed through HTTPS. The Delphix Provider uses Data Control Tower (DCT) APIs to communicate with Delphix Continuous Data Engines. 

Authentication with DCT APIs are managed using API Keys. For generation of an API key, please refer to [DCT API Keys](https://dct.delphix.com/docs/latest/api-keys).

## Example Usage

The following script demonstrates how to configure the Delphix Provider to connect with Data Control Tower and then provision a VDB. Additional resource guides and documentation can be found on the left hand side. 

```hcl
terraform {
  required_providers {
    delphix = {
      source = "delphix-integrations/delphix"
      version = "3.3.0"
    }
  }
}

# Configure the DXI Provider
provider "delphix" {
  host = "dct_hostname"
  key = "dct_api_key"
  tls_insecure_skip = false
}

# Provision a VDB
resource "delphix_vdb" "vdb_name" {
  auto_select_repository = true
  source_data_id         = "DATASOURCE_ID"
}
```

### Example Global Parameter Reference

* __host__: The hostname for DCT.
* __key__ : The API Key which is used to authenticate with DCT. (Example `apk 2.abc123...`).
* __tls_insecure_skip__: (Optional) A boolean value which determines whether to skip the SSL/TLS check. The default value is `false`. Skipping any SSL/TLS check is not recommended for production environments. 
* __host_scheme__: (Optional) Determines the configured host URL's scheme. The default value is `https`.
  
Consult the documentation's Resources section for details on individual resources, such as VDB, dSource, and Environment.

## Support Matrix

Feature/Product | Provider Version | DCT version 
--- |------------------| --- | 
delphix_vdb | v 1.0.0          | v 2.0.0 
delphix_vdb_group | v 1.0.0          | v 2.0.0 
delphix_environment | v 1.0.0          | v 2.0.0 
delphix_appdata_dsource | v 2.1.0          | v 10.0.0 
delphix_oracle_dsource | v 3.1.0          | v 10.0.0 
delphix_database_postgresql | v 3.2.0          | v 14.0.0 
delphix_vdb update<br>delphix_database_postgresql import | v 3.3.0  | v 22.0.0

