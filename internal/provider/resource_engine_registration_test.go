package provider

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccEngineRegistration_withSSLConfig(t *testing.T) {
	resourceName := "delphix_engine_dct_registration.test"
	engineName := "test-engine-ssl-" + randomString(8)
	engineHostname := os.Getenv("TEST_ENGINE_HOSTNAME")
	engineUsername := os.Getenv("TEST_ENGINE_USERNAME")
	enginePassword := os.Getenv("TEST_ENGINE_PASSWORD")

	if engineHostname == "" || engineUsername == "" || enginePassword == "" {
		t.Skip("Required environment variables not set")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckEngineRegistrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccEngineRegistrationWithSSL(
					engineName, engineHostname, engineUsername, enginePassword,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEngineRegistrationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", engineName),
					resource.TestCheckResourceAttr(resourceName, "insecure_ssl", "true"),
					resource.TestCheckResourceAttr(resourceName, "unsafe_ssl_hostname_check", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

func TestAccEngineRegistration_withTags(t *testing.T) {
	resourceName := "delphix_engine_dct_registration.test"
	engineName := "test-engine-tags-" + randomString(8)
	engineHostname := os.Getenv("TEST_ENGINE_HOSTNAME")
	engineUsername := os.Getenv("TEST_ENGINE_USERNAME")
	enginePassword := os.Getenv("TEST_ENGINE_PASSWORD")

	if engineHostname == "" || engineUsername == "" || enginePassword == "" {
		t.Skip("Required environment variables not set")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckEngineRegistrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccEngineRegistrationWithTags(
					engineName, engineHostname, engineUsername, enginePassword,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEngineRegistrationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", engineName),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.0.key", "environment"),
					resource.TestCheckResourceAttr(resourceName, "tags.0.value", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.1.key", "team"),
					resource.TestCheckResourceAttr(resourceName, "tags.1.value", "qa"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

func TestAccEngineRegistration_validationErrors(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccEngineRegistrationMissingName(),
				ExpectError: regexp.MustCompile("The argument \"name\" is required"),
			},
			{
				Config:      testAccEngineRegistrationMissingHostname(),
				ExpectError: regexp.MustCompile("The argument \"hostname\" is required"),
			},
			{
				Config:      testAccEngineRegistrationMissingCredentials(),
				ExpectError: regexp.MustCompile("The argument \"username\" is required"),
			},
		},
	})
}

func testAccCheckEngineRegistrationExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		client := testAccProvider.Meta().(*apiClient).client
		_, _, err := client.ManagementAPI.GetRegisteredEngine(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("Error getting registered engine: %s", err)
		}

		return nil
	}
}

func testAccCheckEngineRegistrationDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*apiClient).client
	time.Sleep(time.Duration(60) * time.Second)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "delphix_engine_dct_registration" {
			continue
		}

		_, httpRes, err := client.ManagementAPI.GetRegisteredEngine(context.Background(), rs.Primary.ID).Execute()
		if err == nil {
			return fmt.Errorf("Engine registration still exists")
		}

		// Check if it's a 404 - which means the resource was successfully deleted
		if httpRes != nil && httpRes.StatusCode == 404 {
			continue
		}

		// If we get any other error, return it
		return fmt.Errorf("Unexpected error checking for deleted engine registration: %s", err)
	}

	return nil
}

// Test configuration templates

func testAccEngineRegistrationWithSSL(name, hostname, username, password string) string {
	return fmt.Sprintf(`
resource "delphix_engine_dct_registration" "test" {
  name                        = "%s"
  hostname                    = "%s"
  username                    = "%s"
  password                    = "%s"
  insecure_ssl                = true
  unsafe_ssl_hostname_check   = true
  engine_type = "CD"
}
`, name, hostname, username, password)
}

// Validation error test configurations

func testAccEngineRegistrationMissingName() string {
	return `
resource "delphix_engine_dct_registration" "test" {
  hostname = "test.example.com"
  username = "admin"
  password = "password"
  engine_type = "CD"
}
`
}

func testAccEngineRegistrationMissingHostname() string {
	return `
resource "delphix_engine_dct_registration" "test" {
  name     = "test-engine"
  username = "admin"
  password = "password"
  engine_type = "CD"
}
`
}

func testAccEngineRegistrationMissingCredentials() string {
	return `
resource "delphix_engine_dct_registration" "test" {
  name     = "test-engine"
  hostname = "test.example.com"
  password = "password"
  engine_type = "CD"
}
`
}

// Utility function to generate random strings for unique resource names
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func testAccEngineRegistrationWithTags(name, hostname, username, password string) string {
	return fmt.Sprintf(`
resource "delphix_engine_dct_registration" "test" {
  name     = "%s"
  hostname = "%s"
  username = "%s"
  password = "%s"
  insecure_ssl = true
  engine_type = "CD"
  tags {
    key   = "environment"
    value = "test"
  }
  
  tags {
    key   = "team"
    value = "qa"
  }
}
`, name, hostname, username, password)
}
