# <resource name> delphix_environment

An environment is a a grouping of a single host or a cluster of hosts. environment allows creating hosts or a cluster of hosts.

## Example Usage

### Creating a UNIX standalone environment
```hcl
resource "delphix_environment" "unix_env_name" {
     engine_id = 2
     os_name = "UNIX"
     username = "xxx"
     password = "xxx"
     hostname = "db.host.com"
     toolkit_path = "/home/delphix"
     name = "my-env"
     is_cluster = false
     cluster_home = "/home/ghrid"
     staging_environment = "stage"
     connector_port = 5312
     ssh_port = 22
     ase_db_password = "test"
     ase_db_username = "user-123"
     java_home = "/java/home"
     dsp_keystore_alias = "alias"
     dsp_keystore_password = "pass"
     dsp_keystore_path = "path"
     dsp_truststore_password = "pass"
     dsp_truststore_path = "/work"
     description = "desc"
     is_target = false
 }
```
### Creating a UNIX cluster
```hcl
resource "delphix_environment" "unixcluster" {
     engine_id = 2
     os_name = "UNIX"
     username = "xxx"
     password = "xxx"
     hostname = "db.host.com"
     toolkit_path = "/home/delphix"
     name = "unixcluster"
     description = "This is a unix target." 
     is_cluster = true    
     cluster_home = "/u01/app/19.0.0.0/grid"
 }
```
### Creating a WINDOWS standalone target environment
```hcl
resource "delphix_environment" "wintgt" {
     engine_id = 2
     os_name = "WINDOWS"
     username = "xxx"
     password = "xxx"
     hostname = "xxx"
     name = "wintgt"
     connector_port = 9100
     ssh_port = 22
     description = "This is a windows target."
 }
```
### Creating a WINDOWS standalone source environment
```hcl
resource "delphix_environment" "WindowsSrc" {
     engine_id = 2
     os_name = "WINDOWS"
     username = "xxx"
     password = "xxx"
     hostname = "db.host.com"
     name = "WindowsSrc"
     staging_environment = delphix_environment.wintgt.id
 }
```
### Creating a WINDOWS cluster source environment
```hcl
resource "delphix_environment" "winsrc-cluster" {
     engine_id = 2
     is_target = false
     os_name = "WINDOWS"
     username = "xxx"
     password = "xxx"
     hostname = "xxx"
     name = "winsrc-cluster"
     staging_environment = delphix_environment.wintgt.id
     is_cluster = true
 }
```
### Creating a WINDOWS failover cluster that can be used as target
```hcl
resource "delphix_environment" "fc-cluster-0" {
     engine_id = 2
     os_name = "WINDOWS"
     username = "xxx"
     password = "xxx"
     hostname = "xxx"
     name = "fc-cluster-0"
     connector_port = 9100
     description = "This is an FC cluster"
 }
 resource "delphix_environment" "fc-cluster-1" {
     engine_id = 2
     os_name = "WINDOWS"
     username = "xxx"
     password = "xxx"
     hostname = "xxx"
     name = "fc-cluster-1"
     connector_port = 9100
     description = "This is an FC cluster."
 }
resource "delphix_environment" "fc-tgt-cluster" {
     engine_id = 2
     is_target = true
     os_name = "WINDOWS"
     username = "xxx"
     password = "xxx"
     hostname = "db.host.com"
     name = "fc-tgt-cluster"
     staging_environment = delphix_environment.fc-cluster-1.id
     is_cluster = true
 }

```

## Argument Reference

* `name` - (Optional) The name of the environment.
* `engine_id` - (Required) The ID of the Engine onto which to create the environment.
* `os_name` - (Required) Operating system type of the environment. Valid values are `[UNIX, WINDOWS]`
* `is_cluster` - (Optional) Whether the environment to be created is a cluster.
* `cluster_home` - (Optional) Absolute path to cluster home drectory. This parameter is mandatory for UNIX cluster environments.
* `hostname` - (Required) host address of the machine.
* `staging_environment` - (Optional) Id of the connector environment which is used to connect to this source environment. This is mandatory parameter when creating Windows source environments.
* `connector_port` - (Optional) Specify port on which Delphix connector will run. This is mandatory parameter when creating Windows target environments.
* `is_target` - (Optional) Whether the environment to be created is a target cluster environment. This property is used only when creating Windows cluster environments.
* `ssh_port` - (Optional) ssh port of the host.
* `toolkit_path` - (Optional) The path for the toolkit that resides on the host.
* `username` - (Optional) Username of the OS.
* `password` - (Optional) Password of the OS.
* `vault` - (Optional) The name or reference of the vault from which to read the host credentials.
* `hashicorp_vault_engine` - (Optional) Vault engine name where the credential is stored.
* `hashicorp_vault_secret_path` - (Optional) Path in the vault engine where the credential is stored.
* `hashicorp_vault_username_key` - (Optional) Key for the username in the key-value store.
* `hashicorp_vault_secret_key` - (Optional) Key for the password in the key-value store.
* `cyberark_vault_query_string` - (Optional) Query to find a credential in the CyberArk vault.
* `nfs_addresses` - (Optional) array of ip address or hostnames. Valid values are a list of addresses. For eg: `["192.168.10.2"]`
* `ase_db_username` - (Optional) username of the SAP ASE database.
* `ase_db_password` - (Optional) password of the SAP ASE database.
* `ase_db_vault` - (Optional) The name or reference of the vault from which to read the ASE database credentials.
* `ase_db_hashicorp_vault_engine` - (Optional) Vault engine name where the credential is stored.
* `ase_db_hashicorp_vault_secret_path` - (Optional) Path in the vault engine where the credential is stored.
* `ase_db_hashicorp_vault_username_key` - (Optional) Key for the username in the key-value store.
* `ase_db_hashicorp_vault_secret_key` - (Optional) Key for the password in the key-value store.
* `ase_db_cyberark_vault_query_string` - (Optional) Query to find a credential in the CyberArk vault.
* `java_home` - (Optional) The path to the user managed Java Development Kit (JDK). If not specified, then the OpenJDK will be used.
* `dsp_keystore_path` - (Optional) DSP keystore path.
* `dsp_keystore_password` - (Optional) DSP keystore password.
* `dsp_keystore_alias` - (Optional) DSP keystore alias.
* `dsp_truststore_path` - (Optional) DSP truststore path.
* `dsp_truststore_password` - (Optional) DSP truststore password.
* `description` - (Optional) The environment description.

## Attribute Reference

* `namespace` - The namespace of this environment for replicated and restored objects.
* `engine_id` - A reference to the Engine that this Environment connection is associated with.
* `enabled` - True if this environment is enabled.
* `hosts` - The hosts that are part of this environment.