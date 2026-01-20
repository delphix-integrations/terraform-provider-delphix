# Resource: <resource name> delphix_engine_dct_registration
This resource helps to register Delphix Engine to Delphix Data Control Tower(DCT)
## Example Usage

Register an Compliance Engine.
```hcl
resource "delphix_engine_dct_registration" "register4" {
  hostname     = "ccblo.dlpxdc.co"
  name         = "compliance-engine"
  engine_type = "CC"  
  username     = "admin"
  password     = "xxxx"
  compliance_user = "admin"
  compliance_password = "xxxx"
  insecure_ssl = true
}
```
Register an Virtualisation Engine.
```hcl
resource "delphix_engine_dct_registration" "register4" {
  hostname     = "cdblo.dlpxdc.co"
  name         = "virtual-engine"
  engine_type = "CD"  
  username     = "admin"
  password     = "xxxx"
  insecure_ssl = true
}
```
## Argument Reference
The following arguments apply to all configurations.
* `hostname` - (Required) Hostname of the engine to register on DCT
* `name` - (Required) Name of engine to display in DCT.
* `engine_type ` - Type of Engine  `CC ` or  `CD `.
* `username ` - (Required) Username for engine admin.
* `password ` - (Required) Password of engine admin.
* `compliance_user ` - (Required if  `engine_type ` is  ` CC`) Username for compliance engine admin.
* `compliance_password ` - (Required if  `engine_type ` is  ` CC`) Password of engine admin.
* `insecure_ssl ` - Boolean flag for secure SSL.



## Limitations
Please contact Delphix Support to make a feature request and/or create a GitHub Issue.