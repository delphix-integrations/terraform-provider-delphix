# Resource: <resource name> delphix_environment

In Delphix, an environment is either a single instance host or cluster of hosts that run database software.

Environments can either be a source (where data comes from), staging (where data are prepared/masked) or target (where data are delivered and used by developers and testers).

Each environment has its own properties and information depending on the type of environment it is

## Example Usage

### Create UNIX standalone environment

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
### Create UNIX cluster
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
### Creating UNIX standalone target environment using HashiCorp Vault
```hcl
resource "delphix_environment" "wintgt" {
     engine_id = 2
     os_name = "UNIX"
     hostname = "xxx"
     toolkit_path = "/home/delphix"
     name = "unixtgt"

     vault = "vault-name"
     hashicorp_vault_engine       = "xxx"
     hashicorp_vault_secret_path  = "xxx"
     hashicorp_vault_username_key = "xxx"
     hashicorp_vault_secret_key   = "xxx"

     description = "This is unix target."
 }
```
### Creating UNIX standalone target environment using CyberArk Vault
```hcl
resource "delphix_environment" "wintgt" {
     engine_id = 2
     os_name = "UNIX"
     hostname = "xxx"
     toolkit_path = "/home/delphix"
     name = "unixtgt"

     vault = "vault-name"
     cyberark_query_string = "xxx"

     description = "This is unix target."
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

* `engine_id` - (Required) The DCT ID of the Engine on which to create the environment. This ID can be obtained by querying the DCT engines API. A Delphix Engine must be registered with DCT first for it to create an Engine ID.
* `os_name` - (Required) Operating system type of the environment. Valid values are `[UNIX, WINDOWS]`
* `hostname` - (Required) Host Name or IP Address of the host that being added to Delphix.
* `name` - The name of the environment.
* `is_cluster` - Whether the environment to be created is a cluster.
* `cluster_home` - Absolute path to cluster home drectory. This parameter is (Required) for UNIX cluster environments.
* `staging_environment` - Id of the environment where Delphix Connector is installed. This is a (Required) parameter when creating Windows source environments.
* `connector_port` - Specify port on which Delphix connector will run. This is a (Required) parameter when creating Windows target environments.
* `is_target` - Whether the environment to be created is a target cluster environment. This property is used only when creating Windows cluster environments.
* `ssh_port` - ssh port of the environment.
* `toolkit_path` - The path where Delphix toolkit can be pushed.
* `username` - OS username for Delphix.
* `password` - OS user's password.
* `vault` - The name or reference of the vault from which to read the host credentials.
* `hashicorp_vault_engine` - Vault engine name where the credential is stored.
* `hashicorp_vault_secret_path` - Path in the vault engine where the credential is stored.
* `hashicorp_vault_username_key` - Key for the username in the key-value store.
* `hashicorp_vault_secret_key` - Key for the password in the key-value store.
* `cyberark_vault_query_string` - Query to find a credential in the CyberArk vault.
* `use_kerberos_authentication` - Whether to use kerberos authentication.
* `use_engine_public_key` - Whether to use public key authentication.
* `nfs_addresses` - Array of ip address or hostnames. Valid values are a list of addresses. For eg: `["192.168.10.2"]`
* `ase_db_username` - Username for the SAP ASE database.
* `ase_db_password` - Password for the SAP ASE database.
* `ase_db_vault` - The name or reference of the vault from which to read the ASE database credentials.
* `ase_db_hashicorp_vault_engine` - Vault engine name where the credential is stored.
* `ase_db_hashicorp_vault_secret_path` - Path in the vault engine where the credential is stored.
* `ase_db_hashicorp_vault_username_key` - Key for the username in the key-value store.
* `ase_db_hashicorp_vault_secret_key` - Key for the password in the key-value store.
* `ase_db_cyberark_vault_query_string` - Query to find a credential in the CyberArk vault.
* `ase_db_use_kerberos_authentication` - Whether to use kerberos authentication for ASE DB discovery.
* `java_home` - The path to the user managed Java Development Kit (JDK). If not specified, then the OpenJDK will be used.
* `dsp_keystore_path` - DSP keystore path.
* `dsp_keystore_password` - DSP keystore password.
* `dsp_keystore_alias` - DSP keystore alias.
* `dsp_truststore_path` - DSP truststore path.
* `dsp_truststore_password` - DSP truststore password.
* `description` - The environment description.
* `tags` - The tags to be created for this environment. This is a map of 2 parameters:
  * `key` - (Required) Key of the tag
  * `value` - (Required) Value of the tag

## Attribute Reference

* `namespace` - The namespace of this environment for replicated and restored objects.
* `engine_id` - A reference to the Engine that this Environment connection is associated with.
* `enabled` - True if this environment is enabled.
* `hosts` - The hosts that are part of this environment.
* `repositories` - The repositories that are part of this environment.
