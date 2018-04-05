package main

import (
	"fmt"
	"log"
	"testing"

	delphix "github.com/delphix/delphix-go-sdk"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccVDBDoesExistBasicCheck(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testVDBCheckDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckVDBConfigImported(&testAccDelphixAdminConfig, &testAccVDB),
				Check: resource.ComposeTestCheckFunc(
					testVDBCheckDoesExist(testAccVDB.name),
					resource.TestCheckResourceAttr(
						"delphix_vdb.testaccvdb", "name", testAccVDB.name),
					resource.TestCheckResourceAttr(
						"delphix_vdb.testaccvdb", "db_name", testAccVDB.dbName),
				),
			},
		},
	})
}

func TestAccVDBDoesExist_UpdatedName(t *testing.T) {
	updatedVDBName := "UpdatedName"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testVDBCheckDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckVDBConfigImported(&testAccDelphixAdminConfig, &testAccVDB),
				Check: resource.ComposeTestCheckFunc(
					testVDBCheckDoesExist(testAccVDB.name),
					resource.TestCheckResourceAttr(
						"delphix_vdb.testaccvdb", "name", testAccVDB.name),
					resource.TestCheckResourceAttr(
						"delphix_vdb.testaccvdb", "db_name", testAccVDB.dbName),
				),
			},
			resource.TestStep{
				Config: testAccCheckVDBConfigUpdatedName(&testAccDelphixAdminConfig, &testAccVDB, updatedVDBName),
				Check: resource.ComposeTestCheckFunc(
					testVDBCheckDoesExist(updatedVDBName),
					resource.TestCheckResourceAttr(
						"delphix_vdb.testaccvdb", "name", updatedVDBName),
					resource.TestCheckResourceAttr(
						"delphix_vdb.testaccvdb", "db_name", testAccVDB.dbName),
				),
			},
		},
	})
}
func TestAccVDBDoesExist_UpdatedMult(t *testing.T) {
	updatedGroupName := "Dev Copies"
	updatedVDBName := "UpdatedName"
	updatedDBName := "update"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testVDBCheckDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckVDBConfigImported(&testAccDelphixAdminConfig, &testAccVDB),
				Check: resource.ComposeTestCheckFunc(
					testVDBCheckDoesExist(testAccVDB.name),
					resource.TestCheckResourceAttr(
						"delphix_vdb.testaccvdb", "name", testAccVDB.name),
					resource.TestCheckResourceAttr(
						"delphix_vdb.testaccvdb", "db_name", testAccVDB.dbName),
				),
			},
			resource.TestStep{
				Config: testAccCheckVDBConfigUpdatedMult(&testAccDelphixAdminConfig, &testAccVDB, updatedGroupName, updatedVDBName, updatedDBName),
				Check: resource.ComposeTestCheckFunc(
					testVDBCheckDoesExist(updatedVDBName),
					resource.TestCheckResourceAttr(
						"delphix_vdb.testaccvdb", "group_name", updatedGroupName),
					resource.TestCheckResourceAttr(
						"delphix_vdb.testaccvdb", "name", updatedVDBName),
					resource.TestCheckResourceAttr(
						"delphix_vdb.testaccvdb", "db_name", updatedDBName),
				),
			},
		},
	})
}

func testAccCheckVDBConfigImported(c *Config, v *VDB) string {
	return fmt.Sprintf(`
		provider "delphix" {
		url = "${var.url}"
		delphix_admin_username = "%s"
		delphix_admin_password = "%s"
		}

		resource "delphix_vdb" "testaccvdb" {
		group_name = "%s"
		name = "%s"
		db_name = "%s"
		source = "ORACLE_DB_CONTAINER-96"
		environment = "UNIX_HOST_ENVIRONMENT-53"
		oracle_home = "%s"
		}

		variable "url" {
		default = "%s"
		}`,
		c.username, c.password, v.groupName, v.name, v.dbName, v.oracleHome, c.url,
	)
}

func testAccCheckVDBConfigUpdatedName(c *Config, v *VDB, n string) string {
	return fmt.Sprintf(`
		provider "delphix" {
		url = "${var.url}"
		delphix_admin_username = "%s"
		delphix_admin_password = "%s"
		}

		resource "delphix_vdb" "testaccvdb" {
		group_name = "%s"
		name = "%s"
		db_name = "%s"
		source = "ORACLE_DB_CONTAINER-96"
		environment = "UNIX_HOST_ENVIRONMENT-53"
		oracle_home = "%s"
		}

		variable "url" {
		default = "%s"
		}`,
		c.username, c.password, v.groupName, n, v.dbName, v.oracleHome, c.url,
	)
}

func testAccCheckVDBConfigUpdatedMult(c *Config, v *VDB, g string, n string, d string) string {
	return fmt.Sprintf(`
		provider "delphix" {
		url = "${var.url}"
		delphix_admin_username = "%s"
		delphix_admin_password = "%s"
		}

		resource "delphix_vdb" "testaccvdb" {
		group_name = "%s"
		name = "%s"
		db_name = "%s"
		source = "ORACLE_DB_CONTAINER-96"
		environment = "UNIX_HOST_ENVIRONMENT-53"
		oracle_home = "%s"
		}

		variable "url" {
		default = "%s"
		}`,
		c.username, c.password, g, n, d, v.oracleHome, c.url,
	)
}

func testVDBCheckDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*delphix.Client)
	for _, r := range s.RootModule().Resources {
		id := r.Primary.ID
		if vdbObj, err := client.FindDatabaseByReference(id); err != nil {
			return fmt.Errorf("testVDBDestroy failed on FindDatabaseByReference(%s)", id)
		} else if vdbObj != nil {
			return fmt.Errorf("VDB %s still exits", vdbObj.(map[string]interface{})["name"].(string))
		}
	}
	return nil
}

func testVDBCheckDoesExist(name string) resource.TestCheckFunc {
	log.Println("Beginning test")
	return func(s *terraform.State) error {
		// Leaving this here to remind me how I got info from state
		// for _, r := range s.RootModule().Resources {
		// 	id := r.Primary.ID
		// 	attributes := r.Primary.Attributes
		// 	log.Printf("id: %s\n", id)
		// 	log.Printf("attributes: %s\n", attributes["name"])
		// }
		client := testAccProvider.Meta().(*delphix.Client)
		vdbObj, err := client.FindDatabaseByName(name)
		if err != nil {
			return fmt.Errorf("FindDatabaseByName failed: %s", err)
		}
		if vdbObj == nil {
			return fmt.Errorf("Database %s does not exist", vdbObj.(map[string]interface{})["name"].(string))
		}
		return nil
	}
}
