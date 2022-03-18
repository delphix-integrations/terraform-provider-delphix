package provider

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccEnvironment_positive(t *testing.T) {
	engineId := os.Getenv("ACC_ENV_ENGINE_ID")
	username := os.Getenv("ACC_ENV_USERNAME")
	password := os.Getenv("ACC_ENV_PASSWORD")
	hostname := os.Getenv("ACC_ENV_HOSTNAME")
	toolkitPath := os.Getenv("ACC_ENV_TOOLKIT_PATH")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccEnvPreCheck(t, engineId, username, password, hostname, toolkitPath) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckEnvDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDctEnvConfigBasic(engineId, username, password, hostname, toolkitPath),
				Check: resource.ComposeTestCheckFunc(
					// TODO: hostname isn't not set yet?
					testAccCheckDctEnvResourceExists("delphix_environment.new_env", hostname),
					resource.TestCheckResourceAttr("delphix_environment.new_env", "hostname", hostname)),
			},
		},
	})
}

func testAccEnvPreCheck(t *testing.T, engineId string, username string, password string, hostname string, toolkitPath string) {
	testAccPreCheck(t)
	if engineId == "" {
		t.Fatal("ACC_ENV_ENGINE_ID must be set for env acceptance tests")
	}
	if username == "" {
		t.Fatal("ACC_ENV_USERNAME must be set for env acceptance tests")
	}
	if password == "" {
		t.Fatal("ACC_ENV_PASSWORD must be set for env acceptance tests")
	}
	if hostname == "" {
		t.Fatal("ACC_ENV_HOSTNAME must be set for env acceptance tests")
	}
	if toolkitPath == "" {
		t.Fatal("ACC_ENV_TOOLKIT_PATH must be set for env acceptance tests")
	}
}

func escape(s string) string {
	// Escape backslash or terraform interepts it as a special character
	return strings.ReplaceAll(s, "\\", "\\\\")
}

func testAccCheckDctEnvConfigBasic(engineId string, username string, password string, hostname string, toolkitPath string) string {
	return fmt.Sprintf(`
	resource "delphix_environment" "new_env" {
		engine_id = %s
		os_name = "UNIX"
		username = "%s"
		password = "%s"
		hostname = "%s"
		toolkit_path = "%s"
		name = "test-acc-name"
	}
	`, engineId, escape(username), escape(password), escape(hostname), escape(toolkitPath))
}

func testAccCheckDctEnvResourceExists(n string, hostname string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		EnvId := rs.Primary.ID
		if EnvId == "" {
			return fmt.Errorf("No EnvID set")
		}

		client := testAccProvider.Meta().(*apiClient).client
		res, _, err := client.EnvironmentsApi.GetEnvironmentById(context.Background(), EnvId).Execute()
		if err != nil {
			return err
		}

		actualHostname := res.GetHosts()[0].GetHostname()
		if actualHostname != hostname {
			return fmt.Errorf("actualHostname %s does not match hostname %s", actualHostname, hostname)
		}

		return nil
	}
}

func testAccCheckEnvDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*apiClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "delphix_environment" {
			continue
		}

		EnvId := rs.Primary.ID

		_, httpResp, _ := client.EnvironmentsApi.GetEnvironmentById(context.Background(), EnvId).Execute()
		if httpResp == nil {
			return fmt.Errorf("Environment has not been deleted")
		}

		if httpResp.StatusCode != 404 {
			return fmt.Errorf("Exepcted a 404 Not Found for a deleted Environment but got %d", httpResp.StatusCode)
		}
	}

	return nil
}
