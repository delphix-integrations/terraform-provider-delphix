terraform {
  required_providers {
    delphix = {
      version = "1.0.0-beta"
      source  = "delphix.com/dct/delphix"
    }
  }
}

provider "delphix" {
  tls_insecure_skip = true
  key = "xxx"
  host = "localhost"
}

/* Unix Standalone */
resource "delphix_environment" "unixtgt" {
     engine_id = 1
     os_name = "UNIX"
     username = "xxx"
     password = "xxx"
     hostname = "xxx"
     toolkit_path = "/home/delphix_os/toolkit"
     name = "unixtgt"
     description = "This is a unix target."     
 } 

/* Unix Standalone using Hashicorp vault 
resource "delphix_environment" "unixtgt" {
  engine_id = 1
  os_name   = "UNIX"
  hostname  = "xxx"

  vault                        = "xxx"
  hashicorp_vault_engine       = "xxx"
  hashicorp_vault_secret_path  = "xxx"
  hashicorp_vault_username_key = "xxx"
  hashicorp_vault_secret_key   = "xxx"

  toolkit_path = "/home/delphix_os/toolkit"
  name         = "unixtgt"
  description  = "This is a unix target."
} */

/* Unix Standalone using CyberArk vault 
resource "delphix_environment" "unixtgt" {
  engine_id = 1
  os_name   = "UNIX"
  hostname  = "xxx"

  vault                        = "xxx"
  cyberark_query_string        = "xxx"

  toolkit_path = "/home/delphix_os/toolkit"
  name         = "unixtgt"
  description  = "This is a unix target."
} */

 /* Win Standalone - Target*/
/* resource "delphix_environment" "wintgt" {
     engine_id = 2
     os_name = "WINDOWS"
     username = "xxx"
     password = "xxx"
     hostname = "xxx"
     name = "wintgt"
     connector_port = 9100
     ssh_port = 22
     description = "This is a windows target."
 }  */

 /* Win Standalone - Source*/
/* resource "delphix_environment" "WindowsSrc" {
     engine_id = 2
     os_name = "WINDOWS"
     username = "xxx"
     password = "xxx"
     hostname = "xxx"
     name = "WindowsSrc"
     staging_environment = delphix_environment.wintgt.id
 } */


/* Unix Standalone - All Params */
 /* resource "delphix_environment" "env_name" {
     engine_id = 2
     os_name = "UNIX"
     username = "xxx"
     password = "xxx"
     hostname = "xxx"
     toolkit_path = "/home/delphix"
     name = "Test"
     is_cluster = false
     cluster_home = "/home/ghrid"
     staging_environment = "stage"
     connector_port = 5312
     ssh_port = 22
     ase_db_password = "pass"
     ase_db_username = "user"
     java_home = "/j/h"
     dsp_keystore_alias = "alias"
     dsp_keystore_password = "pass"
     dsp_keystore_path = "path"
     dsp_truststore_password = "pass"
     dsp_truststore_path = "path"
     description = "desc"
     is_target = false
 } */


/* Win Cluster - Source*/
/* resource "delphix_environment" "winsrc-cluster" {
     engine_id = 2
     is_target = false
     os_name = "WINDOWS"
     username = "xxx"
     password = "xxx"
     hostname = "xxx"
     name = "winsrc-cluster"
     staging_environment = delphix_environment.wintgt.id
     is_cluster = true
 }   */

/* Unix Cluster */
 /* resource "delphix_environment" "unixcluster" {
     engine_id = 2
     os_name = "UNIX"
     username = "xxx"
     password = "xxx"
     hostname = "xxx"
     toolkit_path = "/work"
     name = "unixcluster"
     description = "This is a unix target." 
     is_cluster = true    
     cluster_home = "/u01/app/19.0.0.0/grid"
 } */


 /* Windows Failover Cluster - Used as target */
 /* resource "delphix_environment" "fc-cluster-0" {
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
     hostname = "xxx"
     name = "fc-tgt-cluster"
     staging_environment = "2-WINDOWS_HOST_ENVIRONMENT-35"
     is_cluster = true
 }   */