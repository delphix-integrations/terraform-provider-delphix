# Resource: <resource name> delphix_environment

In Delphix, an environment is either a single instance host or cluster of hosts that run database software. 
Environments can either be a source (where data comes from), staging (where data are prepared/masked) or target (where data are delivered and used by developers and testers). 
Each environment has its own properties and information depending on the operating system, installation, purpose, etc. 
The Delphix Environment resource allows Terraform to create, update, and delete Environments. This specifically enables the `apply`, `import`, and `destroy` Terraform commands. 
Updating existing Delphix Environment resource parameters via the `apply` command is supported for the parameters specified below.   
Note: In DCT, these are called Infrastructure Connections. 

## Example Usage

### Create UNIX standalone environment
```hcl
resource "delphix_environment" "unix_env_name" { 
     engine_id = 2 
     os_type = "UNIX"  
     name = "my-env" 
     username = "xxx" 
     password = "xxx" 
     hosts { 
        hostname = "db.host.com"        
        toolkit_path = "/home/delphix" 
        ssh_port = 22                   
        java_home = "/java/home"       
     } 
     is_cluster = false 
     cluster_home = "/home/ghrid" 
     staging_environment = "stage" 
     connector_port = 5312 
     ase_db_password = "test" 
     ase_db_username = "user-123" 
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
resource "delphix_environment" "unix_cluster" { 
     engine_id = 2 
     os_type = "UNIX" 
     name = "unixcluster" 
     description = "This is a unix target." 
     username = "xxx" 
     password = "xxx" 
     hosts { 
        hostname = "db.host.com" 
        toolkit_path = "/home/delphix" 
     } 
     is_cluster = true     
     cluster_home = "/u01/app/19.0.0.0/grid" 
 } 
```

### Creating UNIX standalone target environment using HashiCorp Vault 
```hcl
resource "delphix_environment" "unix_with_hashi_vault" { 
     engine_id = 2 
     os_type = "UNIX" 
     name = "unixtgt" 
     hosts { 
        hostname = "xxx" 
        toolkit_path = "/home/delphix" 
     } 
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
resource "delphix_environment" "unix_with_ca_vault" { 
     engine_id = 2 
     os_type = "UNIX" 
     name = "unixtgt" 
     description = "This is unix target." 
     hosts { 
        hostname = "xxx" 
        toolkit_path = "/home/delphix" 
     } 
     vault = "vault-name" 
     cyberark_query_string = "xxx" 
 } 
``` 

### Creating a WINDOWS standalone target environment 
```hcl
resource "delphix_environment" "win_tgt" { 
     engine_id = 2 
     os_type = "WINDOWS" 
     name = "wintgt" 
     description = "This is a windows target." 
     username = "xxx" 
     password = "xxx" 
     hosts { 
        hostname = "xxx" 
        ssh_port = 22 
     } 
     connector_port = 9100  
 } 
```

### Creating a WINDOWS standalone source environment 
```hcl
resource "delphix_environment" "win_standalone" { 
     engine_id = 2 
     os_type = "WINDOWS" 
     name = "WindowsSrc" 
     username = "xxx" 
     password = "xxx" 
     hosts { 
        hostname = "db.host.com" 
     }  
     staging_environment = delphix_environment.wintgt.id 
 } 
```

### Creating a WINDOWS cluster source environment
```hcl
resource "delphix_environment" "winsrc_cluster" { 
     engine_id = 2 
     is_target = false 
     os_type = "WINDOWS" 
     name = "winsrc-cluster" 
     username = "xxx" 
     password = "xxx" 
     hosts { 
        hostname = "xxx" 
     } 
     staging_environment = delphix_environment.wintgt.id 
     is_cluster = true 
 } 
``` 

### Creating a WINDOWS failover cluster that can be used as target 
```hcl
resource "delphix_environment" "win_fc_cluster_0" { 
     engine_id = 2 
     os_type = "WINDOWS" 
     name = "fc-cluster-0" 
     description = "This is an FC cluster" 
     username = "xxx" 
     password = "xxx" 
     hosts { 
        hostname = "xxx" 
     } 
     connector_port = 9100 
 } 
 resource "delphix_environment" "win_fc_cluster_1" { 
     engine_id = 2 
     os_type = "WINDOWS" 
     name = "fc-cluster-1" 
     description = "This is an FC cluster." 
     username = "xxx" 
     password = "xxx" 
     hosts { 
        hostname = "xxx" 
     } 
     connector_port = 9100 
 } 
resource "delphix_environment" "win_fc_tgt_cluster" { 
     engine_id = 2 
     is_target = true 
     os_type = "WINDOWS" 
     name = "fc-tgt-cluster" 
     username = "xxx" 
     password = "xxx" 
     hosts { 
        hostname = "db.host.com" 
     } 
     staging_environment = delphix_environment.fc-cluster-1.id 
     is_cluster = true 
 } 
```

## Argument Reference

### General Linking Requirements 
* `name` - The name of the environment. [Updatable] 
* `description` - The environment description. [Updatable] 
* `os_type` - (Required) Operating system type of the environment. Valid values are [UNIX, WINDOWS] 
* `engine_id` - (Required) The DCT ID of the Engine on which to create the environment. This ID can be obtained by querying the DCT engines API. A Delphix Engine must be registered with DCT first for it to create an Engine ID. 
* `is_cluster` - Whether the environment to be created is a cluster. 
* `cluster_home` - Absolute path to cluster home directory. This parameter is (Required) for UNIX cluster environments. [Updatable] 
* `staging_environment` - Id of the environment where Delphix (Windows) Connector is installed. This is a required parameter when creating Windows source environments. 
* `connector_port` - Specify port on which Delphix connector will run. This is a (Required) parameter when creating Windows target environments. [Updatable] 
* `is_target` - Whether the environment to be created is a target cluster environment. This property is used only when creating Windows cluster environments. 

### Host Arguments 
* `hostname` - (Required) Host Name or IP Address of the host that being added to Delphix. [Updatable] 
* `ssh_port` - ssh port of the environment. [Updatable] 
* `toolkit_path` - The path where Delphix toolkit can be pushed. [Updatable] 
* `java_home` - The path to the user managed Java Development Kit (JDK). If not specified, then the OpenJDK will be used. [Updatable] 
* `nfs_addresses` - Array of ip address or hostnames. Valid values are a list of addresses. For eg: ["192.168.10.2"] [Updatable] 

### General Authentication Arguments 
* `dsp_keystore_path` - DSP keystore path. 
* `dsp_keystore_password` - DSP keystore password. 
* `dsp_keystore_alias` - DSP keystore alias. 
* `dsp_truststore_path` - DSP truststore path. 
* `dsp_truststore_password` - DSP truststore password. 
* `use_engine_public_key` - Whether to use public key authentication. 

### SQL Server Authentication Arguments 
* `username` - OS username to enable a connection from the engine. [Updatable] 
* `password` - OS user's password. [Updatable] 
* `vault` - The name or reference of the vault from which to read the host credentials. 
* `hashicorp_vault_engine` – The Hashicorp Vault engine name where the credential is stored. 
* `hashicorp_vault_secret_path` - Path in the Hashicorp Vault engine where the credential is stored. 
* `hashicorp_vault_username_key` - Key for the username in the key-value store. 
* `hashicorp_vault_secret_key` - Key for the password in the key-value store. 
* `cyberark_vault_query_string` - Query to find a credential in the CyberArk vault. 
* `use_kerberos_authentication` - Whether to use Kerberos authentication. 

### SAP ASE (Sybase) Authentication Arguments 
* `ase_db_username` - Username for the SAP ASE database. 
* `ase_db_password` - Password for the SAP ASE database. 
* `ase_db_vault` - The name or reference of the vault from which to read the SAP ASE database credentials. 
* `ase_db_hashicorp_vault_engine` – The Hashicorp Vault engine name where the credential is stored. 
* `ase_db_hashicorp_vault_secret_path` - Path in the Hashicorp Vault engine where the credential is stored. 
* `ase_db_hashicorp_vault_username_key` - Key for the username in the key-value store. 
* `ase_db_hashicorp_vault_secret_key` - Key for the password in the key-value store. 
* `ase_db_cyberark_vault_query_string` - Query to find a credential in the CyberArk vault. 
* `ase_db_use_kerberos_authentication` - Whether to use Kerberos authentication for SAP ASE DB discovery. 

### Advanced Arguments 
* `tags` - The tags to be created for this environment. This is a map of 2 parameters: [Updatable] 
   * `key` - (Required) Key of the tag 
   * `value` - (Required) Value of the tag 
* `ignore_tag_changes` – This flag enables whether changes in the tags are identified by terraform. By default, it is true, i.e, changes in tags of the resource are ignored. 

## Import (Beta) 
Use the import block to add Appdata dSources created directly in DCT into a Terraform state file.  

For example:  
```hcl
import {    
    to = delphix_environment.env_import_demo
    id = "env_id"    
}   
```

This is a beta feature. Delphix offers no guarantees of support or compatibility.  

## Limitations 
Not all properties are supported through the update command. Properties that are not supported by the update command are presented via an error message at runtime.

