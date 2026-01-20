# Resource: <resource name> delphix_database_plugin
This resource helps to upload database plugin json to Engine.
## Example Usage

Upload database plugin to delphix engine
```hcl
resource "delphix_database_plugin" "upload" {
  engine_host = "https://engine.dlpx.co"
  file_path = "PATH-TO-PLUGIN-JSON"
}
```
## Argument Reference
The following arguments apply to all configurations.
* `engine_host` - (Required) Hostname of the engine to upload plugin on. [NOTE] Engine should be already registered into DCT.
* `file_path` - (Required) Path to plugin json file.


## Limitations
Please contact Delphix Support to make a feature request and/or create a GitHub Issue.