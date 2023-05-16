
### Simple Getting Stated 

1) Update the dct_hostname and dct_api_key variables in the `variables.tf` or in Terrafrom Cloud with the DCT Server and API Key. 
For example:
- DCT: uv123abcfei59h6qajyy.vm.cld.sr
- API Key: 2.ZAAgpjHxljW7A7g...

2) Update `Oracle_QA` with an applicable VDB name and run the following commands. 
```
# Create all resources
terraform apply -var="source_data_id_1=Oracle_QA"

# Destroy resources
terraform destroy"
```


### Troubleshoot: Invalid Resource State

If you find that you've lost the sync between DCT and Terraform, use the `terraform state rm` command to help reconfigure without starting over.
```
terraform state rm delphix_vdb.provision_vdb_1

terraform state rm delphix_vdb_group.create_vdb_group
```

[Documentation](https://developer.hashicorp.com/terraform/cli/commands/state/rm)
