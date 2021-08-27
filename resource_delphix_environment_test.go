package main

import (
	"fmt"
	"log"
	"testing"

	delphix "github.com/ajaytho/delphix-go-sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccEnvironmentDoesExistBasicCheck(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testEnvironmentCheckDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckEnvironmentConfigImported(&testAccDelphixAdminConfig, &testAccEnvironment),
				Check: resource.ComposeTestCheckFunc(
					testEnvironmentCheckDoesExist(testAccEnvironment.name),
					resource.TestCheckResourceAttr(
						"delphix_environment.test_acc_env", "name", testAccEnvironment.name),
					resource.TestCheckResourceAttr(
						"delphix_environment.test_acc_env", "description", testAccEnvironment.description),
				),
			},
		},
	})
}

func TestAccEnvironmentDoesExistBasicCheckPublicKey(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testEnvironmentCheckDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckEnvironmentConfigPublicKey(&testAccDelphixAdminConfig, &testAccEnvironment),
				Check: resource.ComposeTestCheckFunc(
					testEnvironmentCheckDoesExist(testAccEnvironment.name),
					resource.TestCheckResourceAttr(
						"delphix_environment.test_acc_env", "name", testAccEnvironment.name),
					resource.TestCheckResourceAttr(
						"delphix_environment.test_acc_env", "description", testAccEnvironment.description),
				),
			},
		},
	})
}

func TestAccEnvironmentDoesExist_UpdatedName(t *testing.T) {
	updatedEnvironmentName := "UpdatedName"
	updatedDescription := "UpdatedDescription"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testEnvironmentCheckDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckEnvironmentConfigImported(&testAccDelphixAdminConfig, &testAccEnvironment),
				Check: resource.ComposeTestCheckFunc(
					testEnvironmentCheckDoesExist(testAccEnvironment.name),
					resource.TestCheckResourceAttr(
						"delphix_environment.test_acc_env", "name", testAccEnvironment.name),
					resource.TestCheckResourceAttr(
						"delphix_environment.test_acc_env", "description", testAccEnvironment.description),
				),
			},
			resource.TestStep{
				Config: testAccCheckEnvironmentConfigUpdatedName(&testAccDelphixAdminConfig, &testAccEnvironment, updatedEnvironmentName, updatedDescription),
				Check: resource.ComposeTestCheckFunc(
					testEnvironmentCheckDoesExist(updatedEnvironmentName),
					resource.TestCheckResourceAttr(
						"delphix_environment.test_acc_env", "name", updatedEnvironmentName),
					resource.TestCheckResourceAttr(
						"delphix_environment.test_acc_env", "description", updatedDescription),
				),
			},
		},
	})
}

func TestAccEnvironmentDoesExist_UpdatedServerID(t *testing.T) {
	origServerID := "54321"
	updatedServerID := "12345"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testEnvironmentCheckDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckEnvironmentConfigUpdatedServerID(&testAccDelphixAdminConfig, &testAccEnvironment, origServerID),
				Check: resource.ComposeTestCheckFunc(
					testEnvironmentCheckDoesExist(testAccEnvironment.name),
					resource.TestCheckResourceAttr(
						"delphix_environment.test_acc_env", "name", testAccEnvironment.name),
					resource.TestCheckResourceAttr(
						"delphix_environment.test_acc_env", "server_id", origServerID),
				),
			},
			resource.TestStep{
				Config: testAccCheckEnvironmentConfigUpdatedServerID(&testAccDelphixAdminConfig, &testAccEnvironment, updatedServerID),
				Check: resource.ComposeTestCheckFunc(
					testEnvironmentCheckDoesExist(testAccEnvironment.name),
					resource.TestCheckResourceAttr(
						"delphix_environment.test_acc_env", "name", testAccEnvironment.name),
					resource.TestCheckResourceAttr(
						"delphix_environment.test_acc_env", "server_id", updatedServerID),
				),
			},
		},
	})
}

func testAccCheckEnvironmentConfigUpdatedServerID(c *Config, e *Environment, s string) string {
	return fmt.Sprintf(`
		provider "delphix" {
		url = "${var.url}"
		delphix_admin_username = "%s"
		delphix_admin_password = "%s"
		}

		resource "delphix_environment" "test_acc_env" {
		name = "%s"
		description = "%s"
		address = "%s"
		user_name = "%s"
		user_password = "%s"
		toolkit_path = "%s"
		server_id = "%s"
		}

		variable "url" {
		default = "%s"
		}
		`,
		c.username, c.password, e.name, e.address, e.address, e.userName, e.userPassword, e.toolkitPath, s, c.url,
	)
}
func TestAccEnvironmentDoesExist_UpdatedMulti(t *testing.T) {
	updatedEnvironmentName := "UpdatedName"
	updatedDescription := "UpdatedDescription"
	updatedToolkitPath := "/home/delphix/toolkit"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testEnvironmentCheckDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckEnvironmentConfigImported(&testAccDelphixAdminConfig, &testAccEnvironment),
				Check: resource.ComposeTestCheckFunc(
					testEnvironmentCheckDoesExist(testAccEnvironment.name),
					resource.TestCheckResourceAttr(
						"delphix_environment.test_acc_env", "name", testAccEnvironment.name),
					resource.TestCheckResourceAttr(
						"delphix_environment.test_acc_env", "description", testAccEnvironment.description),
					resource.TestCheckResourceAttr(
						"delphix_environment.test_acc_env", "toolkit_path", testAccEnvironment.toolkitPath),
				),
			},
			resource.TestStep{
				Config: testAccCheckEnvironmentConfigUpdatedMulti(&testAccDelphixAdminConfig, &testAccEnvironment, updatedEnvironmentName, updatedDescription, updatedToolkitPath),
				Check: resource.ComposeTestCheckFunc(
					testEnvironmentCheckDoesExist(updatedEnvironmentName),
					resource.TestCheckResourceAttr(
						"delphix_environment.test_acc_env", "name", updatedEnvironmentName),
					resource.TestCheckResourceAttr(
						"delphix_environment.test_acc_env", "description", updatedDescription),
					resource.TestCheckResourceAttr(
						"delphix_environment.test_acc_env", "toolkit_path", updatedToolkitPath),
				),
			},
		},
	})
}

func testAccCheckEnvironmentConfigImported(c *Config, e *Environment) string {
	return fmt.Sprintf(`
		provider "delphix" {
		url = "${var.url}"
		delphix_admin_username = "%s"
		delphix_admin_password = "%s"
		}

		resource "delphix_environment" "test_acc_env" {
		name = "%s"
		description = "%s"
		address = "%s"
		user_name = "%s"
		user_password = "%s"
		toolkit_path = "%s"
		}

		variable "url" {
		default = "%s"
		}
		`,
		c.username, c.password, e.name, e.description, e.address, e.userName, e.userPassword, e.toolkitPath, c.url,
	)
}

func testAccCheckEnvironmentConfigPublicKey(c *Config, e *Environment) string {
	return fmt.Sprintf(`
		provider "delphix" {
		url = "${var.url}"
		delphix_admin_username = "%s"
		delphix_admin_password = "%s"
		}

		resource "delphix_environment" "test_acc_env" {
		name = "%s"
		description = "%s"
		address = "%s"
		user_name = "%s"
		toolkit_path = "%s"
		public_key = true
		}

		variable "url" {
		default = "%s"
		}
		`,
		c.username, c.password, e.name, e.description, e.address, e.userName, e.toolkitPath, c.url,
	)
}

func testAccCheckEnvironmentConfigUpdatedName(c *Config, e *Environment, n string, d string) string {
	return fmt.Sprintf(`
		provider "delphix" {
		url = "${var.url}"
		delphix_admin_username = "%s"
		delphix_admin_password = "%s"
		}

		resource "delphix_environment" "test_acc_env" {
		name = "%s"
		description = "%s"
		address = "%s"
		user_name = "%s"
		user_password = "%s"
		toolkit_path = "%s"
		}

		variable "url" {
		default = "%s"
		}
		`,
		c.username, c.password, n, d, e.address, e.userName, e.userPassword, e.toolkitPath, c.url,
	)
}

func testAccCheckEnvironmentConfigUpdatedMulti(c *Config, e *Environment, n string, d string, p string) string {
	return fmt.Sprintf(`
		provider "delphix" {
		url = "${var.url}"
		delphix_admin_username = "%s"
		delphix_admin_password = "%s"
		}

		resource "delphix_environment" "test_acc_env" {
		name = "%s"
		description = "%s"
		address = "%s"
		user_name = "%s"
		user_password = "%s"
		toolkit_path = "%s"
		}

		variable "url" {
		default = "%s"
		}
		`,
		c.username, c.password, n, d, e.address, e.userName, e.userPassword, p, c.url,
	)
}

func testEnvironmentCheckDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*delphix.Client)
	for _, r := range s.RootModule().Resources {
		id := r.Primary.ID
		if envObj, err := client.FindEnvironmentByReference(id); err != nil {
			return fmt.Errorf("testEnvironmentDestroy failed on FindEnvironmentByReference(%s)", id)
		} else if envObj != nil {
			return fmt.Errorf("Environment %s still exits", envObj.(map[string]interface{})["name"].(string))
		}
	}
	return nil
}

func testEnvironmentCheckDoesExist(name string) resource.TestCheckFunc {
	log.Println("Beginning test")
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*delphix.Client)
		envObj, err := client.FindEnvironmentByName(name)
		if err != nil {
			return fmt.Errorf("FindEnvironmentByName failed: %v", err)
		}
		if envObj == nil {
			return fmt.Errorf("Environment %s does not exist", envObj.(map[string]interface{})["name"].(string))
		}
		return nil
	}
}
