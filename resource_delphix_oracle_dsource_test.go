package main

import (
	"fmt"
	"log"
	"testing"

	delphix "github.com/delphix/delphix-go-sdk"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccDsourceDoesExistBasicCheck(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testDsourceCheckDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckDsourceConfigImported(&testAccDelphixAdminConfig, &testAccDsource),
				Check: resource.ComposeTestCheckFunc(
					testDsourceCheckDoesExist(testAccDsource.name),
					resource.TestCheckResourceAttr(
						"delphix_data_source_oracle.testaccdsource", "name", testAccDsource.name),
					resource.TestCheckResourceAttr(
						"delphix_data_source_oracle.testaccdsource", "instance", testAccDsource.instance),
				),
			},
		},
	})
}
func TestAccDsourceDoesExist_UpdatedName(t *testing.T) {
	uName := "Updated"
	uDescription := "UPDATED"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testDsourceCheckDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckDsourceConfigImported(&testAccDelphixAdminConfig, &testAccDsource),
				Check: resource.ComposeTestCheckFunc(
					testDsourceCheckDoesExist(testAccDsource.name),
					resource.TestCheckResourceAttr(
						"delphix_data_source_oracle.testaccdsource", "name", testAccDsource.name),
					resource.TestCheckResourceAttr(
						"delphix_data_source_oracle.testaccdsource", "instance", testAccDsource.instance),
				),
			},
			resource.TestStep{
				Config: testAccCheckDsourceConfigUpdated(&testAccDelphixAdminConfig, &testAccDsource, uName, uDescription),
				Check: resource.ComposeTestCheckFunc(
					testDsourceCheckDoesExist(uName),
					resource.TestCheckResourceAttr(
						"delphix_data_source_oracle.testaccdsource", "name", uName),
					resource.TestCheckResourceAttr(
						"delphix_data_source_oracle.testaccdsource", "description", uDescription),
				),
			},
		},
	})
}

func TestAccDsourceDoesExist_UpdatedMulti(t *testing.T) {
	uName := "Updated"
	uDescription := "UPDATED"
	uGroup := "Dev Copies"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testDsourceCheckDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckDsourceConfigImported(&testAccDelphixAdminConfig, &testAccDsource),
				Check: resource.ComposeTestCheckFunc(
					testDsourceCheckDoesExist(testAccDsource.name),
					resource.TestCheckResourceAttr(
						"delphix_data_source_oracle.testaccdsource", "name", testAccDsource.name),
					resource.TestCheckResourceAttr(
						"delphix_data_source_oracle.testaccdsource", "instance", testAccDsource.instance),
					resource.TestCheckResourceAttr(
						"delphix_data_source_oracle.testaccdsource", "group_name", testAccDsource.groupName),
				),
			},
			resource.TestStep{
				Config: testAccCheckDsourceConfigMulti(&testAccDelphixAdminConfig, &testAccDsource, uName, uDescription, uGroup),
				Check: resource.ComposeTestCheckFunc(
					testDsourceCheckDoesExist(uName),
					resource.TestCheckResourceAttr(
						"delphix_data_source_oracle.testaccdsource", "name", uName),
					resource.TestCheckResourceAttr(
						"delphix_data_source_oracle.testaccdsource", "description", uDescription),
				),
			},
		},
	})
}

func testAccCheckDsourceConfigImported(c *Config, d *DSource) string {
	config := fmt.Sprintf(`
		provider "delphix" {
			url = "${var.url}"
			delphix_admin_username = "%s"
			delphix_admin_password = "%s"
		}

		resource "delphix_data_source_oracle" "testaccdsource" {
			name = "%s"
			description = "%s"
			user_name = "%s"
			password = "%s"
			group_name = "%s"
			environment = "UNIX_HOST_ENVIRONMENT-51"
			environment_user = "%s"
			link_now = true
			instance = "%s"
			oracle_home = "%s"
		}
		variable "url" {
			default = "%s"
		}`,
		c.username, c.password, d.name, d.description, d.userName, d.password, d.groupName, d.environmentUser, d.instance, d.oracleHome, c.url,
	)
	fmt.Printf("Config:\n%s\n", config)
	return config
}

func testAccCheckDsourceConfigUpdated(c *Config, d *DSource, n string, desc string) string {
	config := fmt.Sprintf(`
		provider "delphix" {
			url = "${var.url}"
			delphix_admin_username = "%s"
			delphix_admin_password = "%s"
		}

		resource "delphix_data_source_oracle" "testaccdsource" {
			name = "%s"
			description = "%s"
			user_name = "%s"
			password = "%s"
			group_name = "%s"
			environment = "${delphix_environment.my-source-env.id}"
			environment_user = "%s"
			link_now = %v
			instance = "%s"
			oracle_home = "%s"
		}
		variable "url" {
			default = "%s"
		}`,
		c.username, c.password, n, desc, d.userName, d.password, d.groupName, d.environmentUser, d.linkNow, d.instance, d.oracleHome, c.url,
	)
	fmt.Printf("Config:\n%s\n", config)
	return config
}

func testAccCheckDsourceConfigMulti(c *Config, d *DSource, n string, desc string, g string) string {
	config := fmt.Sprintf(`
		provider "delphix" {
			url = "${var.url}"
			delphix_admin_username = "%s"
			delphix_admin_password = "%s"
		}

		resource "delphix_data_source_oracle" "testaccdsource" {
			name = "%s"
			description = "%s"
			user_name = "%s"
			password = "%s"
			group_name = "%s"
			environment = "${delphix_environment.my-source-env.id}"
			environment_user = "%s"
			link_now = %v
			instance = "%s"
			oracle_home = "%s"
		}
		variable "url" {
			default = "%s"
		}`,
		c.username, c.password, n, desc, d.userName, d.password, g, d.environmentUser, d.linkNow, d.instance, d.oracleHome, c.url,
	)
	fmt.Printf("Config:\n%s\n", config)
	return config
}

func testDsourceCheckDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*delphix.Client)
	for _, r := range s.RootModule().Resources {
		id := r.Primary.ID
		if dSourceObj, err := client.FindDatabaseByReference(id); err != nil {
			return fmt.Errorf("testDsourceDestroy failed on FindDatabaseByReference(%s)", id)
		} else if dSourceObj != nil {
			return fmt.Errorf("Dsource %s still exits", dSourceObj.(map[string]interface{})["name"].(string))
		}
	}
	return nil
}

func testDsourceCheckDoesExist(name string) resource.TestCheckFunc {
	log.Println("Beginning test")
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*delphix.Client)
		dSourceObj, err := client.FindDatabaseByName(name)
		if err != nil {
			return fmt.Errorf("FindDatabaseByName failed: %s", err)
		}
		if dSourceObj == nil {
			return fmt.Errorf("Database %s does not exist", dSourceObj.(map[string]interface{})["name"].(string))
		}
		return nil
	}
}
