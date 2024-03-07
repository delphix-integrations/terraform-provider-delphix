# <provider> Delphix Provider

The Terraform Provider for Delphix enables customers to natively manage data-as-code along with their infrastructure.
With Terraform and Delphix, customers can now automatically provision, manage, and teardown any number of ephemeral data environments to drive enterprise DevOps workflows, such as test data management.

This provider communicates directly with Data Control Tower (DCT) to generated virtual database and other objects. Therefore, DCT must be registered with one or more Delphix Continuous Data Engines.

To learn more about Delphix and DCT APIs, refer to [Delphix Documentation](https://documentation.delphix.com/docs/) and [DCT Documentation](https://dct.delphix.com/docs/latest/) respectively. Please [Contact us](ask-integrations@delphix.com) (ask-integrations@delphix.com) with any questions. 

Customers who are entitled to Data Control Tower may also send support issues through the [Delphix Support Portal](https://support.delphix.com/).

## System Requirements

| Product                        | Version  |
|--------------------------------|----------|
| Data Control Tower (DCT)       | v14+ |
| Delphix Continuous Data Engine | v6.0.0.1+ |

Note: The DCT version above guarantees full provider support. However, each  resource might support older versions. Refer to the specific resource documentation page for more information.

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
      version = "3.1.0"
    }
  }
}

# Configure the DXI Provider
provider "delphix" {
  host = "dct_hostname"
  key = "dct_api_key"
  tls_insecure_skip = true
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
* __tls_insecure_skip__: (Optional) A boolean value which determines whether to skip the SSL/TLS check. The dfault value is `false`. Skipping any SSL/TLS check is not recommended for production environments. 
* __host_scheme__: (Optional) Determines the configured host URL's scheme. The default value is `https`.

Consult the documentation's Resources section for details on individual resources, such as VDB, dSource, and Environment.
