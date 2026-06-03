# examples/

Runnable Terraform configuration examples, one directory per provider resource or use case. These are referenced by the Registry docs and the Jenkins integration pipeline.

## Directory Index

| Directory | Resource / Use Case |
|---|---|
| [vdb/](vdb/) | `delphix_vdb` — VDB provisioning across all DB types and provision methods |
| [vdb_group/](vdb_group/) | `delphix_vdb_group` — Group VDBs and manage tags |
| [environment/](environment/) | `delphix_environment` — Unix/Windows environment creation and import |
| [appdata_dsource/](appdata_dsource/) | `delphix_appdata_dsource` — AppData linked sources (PostgreSQL, MySQL) |
| [oracle_dsource/](oracle_dsource/) | `delphix_oracle_dsource` — Oracle RMAN-linked data sources |
| [database/](database/) | `delphix_database_postgresql` — PostgreSQL database objects |
| [database_plugin/](database_plugin/) | `delphix_database_plugin` — Upload plugins to engines |
| [engine_configuration/](engine_configuration/) | `delphix_engine_configuration` — Block/Object storage, NTP, SMTP, SSO |
| [engine_dct_registration/](engine_dct_registration/) | `delphix_engine_dct_registration` — Register CD/CC engines with DCT |
| [simple-provision/](simple-provision/) | End-to-end quick-start: provision a VDB and create a VDB group |
| [full_stack_deployment/](full_stack_deployment/) | Multi-provider example: Azure VM + environment + dSource + VDB |
| [jenkins-integration/](jenkins-integration/) | Jenkins pipeline that provisions a VDB, runs tests, then destroys it |
| [jenkins_integration/](jenkins_integration/) | Alternate copy of the Jenkins integration example |
| [upgrade/](upgrade/) | State files from a provider version upgrade scenario |
| [mongodb/](mongodb/) | Placeholder for MongoDB AppData examples |

## Conventions

- Each directory contains at least a `main.tf`.
- Larger examples may split into `variables.tf`, `outputs.tf`, `versions.tf`.
- Placeholder credentials use `"1.XXXX"` / `"HOSTNAME"` — never commit real keys.
- `.terraform/` and `terraform.tfstate*` files in subdirectories are local run artifacts; they are gitignored in most cases but a few are checked in as reference state.
- The `simple-provision/` example is the canonical starting point referenced by the Jenkins pipeline.

## Running an Example

```bash
cd examples/<name>
terraform init
terraform plan
terraform apply
```

Ensure `DCT_KEY` and `DCT_HOST` are set, or hardcode them in the `provider` block (dev only).
