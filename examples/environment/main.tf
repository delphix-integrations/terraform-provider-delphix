terraform {
  required_providers {
    delphix = {
      version = "0.0-dev"
      source  = "delphix.com/local/dct"
    }
  }
}

provider "delphix" {
  tls_insecure_skip = true
  key = "1.yZHqmkhAfAIbM2h4FhJykPscvL7OiGItB5KZbH76rjhIYWHWOOiV1tcJdZZEhjFk"
  host = "localhost"
}

/* Unix Standalone */
 /* resource "delphix_environment" "unixtgt" {
     engine_id = 3
     os_name = "UNIX"
     username = "root"
     password = "sailboat"
     hostname = "unixsingletgt.dlpxdc.co"
     toolkit_path = "/opt/toolkit"
     name = "unixtgt"
     description = "This is a unix target."     
 } */

 /* Win Standalone - Target*/
/* resource "delphix_environment" "wintgt" {
     engine_id = 3
     os_name = "WINDOWS"
     username = "administrator"
     password = "virtual4uNme"
     hostname = "winsingletgt.dlpxdc.co"
     name = "wintgt"
     connector_port = 9100
     ssh_port = 22
     description = "This is a windows target."
 }  */

 /* Win Standalone - Source*/
 /* resource "delphix_environment" "env_name" {
     engine_id = 2
     os_name = "WINDOWS"
     username = "delphix\\delphix_src"
     password = "delphix"
     hostname = "10.0.1.50"
     name = "WindowsSrc"
     staging_environment = "WINDOWS_HOST_ENVIRONMENT-17"
 } */


/* Unix Standalone - All Params */
 /* resource "delphix_environment" "env_name" {
     engine_id = 2
     os_name = "UNIX"
     username = "delphix"
     password = "delphix"
     hostname = "10.0.1.30"
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
     engine_id = 3
     is_target = false
     os_name = "WINDOWS"
     username = "administrator"
     password = "virtual4uNme"
     hostname = "tftest-node-0.dlpxdc.co"
     name = "winsrc-cluster"
     staging_environment = "3-WINDOWS_HOST_ENVIRONMENT-5"
     is_cluster = true
 } */

/* Unix Cluster */
 resource "delphix_environment" "unixcluster" {
     engine_id = 3
     os_name = "UNIX"
     username = "oracle"
     password = "oracle"
     hostname = "raczqy16ab61.dcol1.delphix.com"
     toolkit_path = "/work"
     name = "unixcluster"
     description = "This is a unix target." 
     is_cluster = true    
     cluster_home = "/u01/app/19.0.0.0/grid"
 }