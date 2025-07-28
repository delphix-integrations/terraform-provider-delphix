# Resource: <resource name> delphix_vdb_group 

 

A VDB Group is a collection of virtual databases. The VDB Group (delphix_vdb_group) resource allows creating such a collection in Data Control Tower (DCT).  

 

VDB Groups are valuable when you need to coordinate the provision, refreshes, and bookmark operations across multiple datasets, such as an application that requires datasets from an Oracle database, PostgreSQL database, and Unstructured Files (vFiles). 

 

VDB Groups must contain VDBs from different root dSources. Therefore, they are not intended for bulk actions, such as refreshing multiple datasets to the same point in time. 

 

## Example Usage 

 

```hcl 

Create a VDB Group, add existing VDBs to the group, and assign tags. 

 

resource "delphix_vdb_group" "vdb_group_name" { 

  name = "my vdb group" 

  vdb_ids = ["vdb_id_1”, “vdb_id_2”, “vdb_id_3"] 

  tags { 

    key = "environment" 

    value = "production" 

  } 

  tags { 

    key = "project" 

    value = "terraform" 

  } 

} 

``` 

 

## Argument Reference 

The following arguments apply to all configurations.   

 

* `id` - A unique identifier for the entity. (Output only) 

* `name` - (Required) The unique name for the VDB Group. [Updatable] 

* `vdb_ids` - (Required) The list of VDB IDs or Names for this VDB Group. [Updatable] 

* `tags` - The tags to be created for the VDB Group. [Updatable] 

This is a map of two required parameters: 

* `key` - Key of the tag. 

* `value` - Value of the tag. 

* `ignore_tag_changes` – This flag enables whether changes in the tags are identified by Terraform. By default, this is set to true, meaning changes to the resource's tags are ignored. 

 

## Import 
Use the [`import` block](https://developer.hashicorp.com/terraform/language/import) to add VDB Groups created directly in DCT into a Terraform state file.   
 
For example:   
```terraform  
import {    
    to = delphix_vdb_group_import_demo   
    id = "vdb_group_id"    
}   
``` 

 

## Limitations 

Please contact Delphix Support to make a feature request and/or create a GitHub Issue. 

 