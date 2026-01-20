package provider

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccEnginePlugin_basic(t *testing.T) {
	resourceName := "delphix_database_plugin.test"
	engineID := os.Getenv("TEST_ENGINE_HOST")
	pluginFilePath := os.Getenv("TEST_PLUGIN_FILE_PATH")

	if engineID == "" || pluginFilePath == "" {
		t.Skip("Required environment variables TEST_ENGINE_HOST and TEST_PLUGIN_FILE_PATH not set")
	}

	// Verify plugin file exists
	if _, err := os.Stat(pluginFilePath); os.IsNotExist(err) {
		t.Skipf("Plugin file does not exist: %s", pluginFilePath)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckEnginePluginDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccEnginePluginBasic(pluginFilePath, engineID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEnginePluginExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "file_path", pluginFilePath),
					resource.TestCheckResourceAttr(resourceName, "engine_host", engineID),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "toolkit_name"),
				),
			},
		},
	})
}

func TestAccEnginePlugin_validationErrors(t *testing.T) {
	engineID := os.Getenv("TEST_ENGINE_HOST")
	if engineID == "" {
		engineID = "test-engine-id"
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccEnginePluginMissingFilePath(engineID),
				ExpectError: regexp.MustCompile("The argument \"file_path\" is required"),
			},
			{
				Config:      testAccEnginePluginMissingEngineID(),
				ExpectError: regexp.MustCompile("The argument \"engine_host\" is required"),
			},
			{
				Config:      testAccEnginePluginNonExistentFile(engineID),
				ExpectError: regexp.MustCompile("file does not exist"),
			},
			{
				Config:      testAccEnginePluginNonJSONFile(engineID),
				ExpectError: regexp.MustCompile("must be a JSON file"),
			},
			{
				Config:      testAccEnginePluginInvalidJSON(engineID),
				ExpectError: regexp.MustCompile("not a valid JSON file"),
			},
		},
	})
}

func testAccCheckEnginePluginExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		client := testAccProvider.Meta().(*apiClient).client
		toolkitReq := client.ToolkitsAPI.GetToolkitById(context.Background(), rs.Primary.ID)
		_, _, err := toolkitReq.Execute()
		if err != nil {
			return fmt.Errorf("Error getting engine: %s", err)
		}

		return nil
	}
}

func testAccCheckEnginePluginDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*apiClient).client
	time.Sleep(time.Duration(10) * time.Second)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "delphix_database_plugin" {
			continue
		}

		// Try to get the plugin - if it still exists, the destroy failed
		if rs.Primary.Attributes["toolkit_name"] != "" {
			// Check if plugin still exists on the engine
			toolkitReq := client.ToolkitsAPI.GetToolkitById(context.Background(), rs.Primary.ID)
			_, httpRes, err := toolkitReq.Execute()
			if err != nil && httpRes != nil && httpRes.StatusCode != 404 {
				return fmt.Errorf("Unexpected error checking for deleted engine plugin: %s", err)
			}
		}
	}

	return nil
}

// Test configuration templates

func testAccEnginePluginBasic(filePath, engineID string) string {
	return fmt.Sprintf(`
resource "delphix_database_plugin" "test" {
  file_path = "%s"
  engine_host = "%s"
}
`, filePath, engineID)
}

// Validation error test configurations

func testAccEnginePluginMissingFilePath(engineID string) string {
	return fmt.Sprintf(`
resource "delphix_database_plugin" "test" {
  engine_host = "%s"
}
`, engineID)
}

func testAccEnginePluginMissingEngineID() string {
	return `
resource "delphix_database_plugin" "test" {
  file_path = "/tmp/test-plugin.json"
}
`
}

func testAccEnginePluginNonExistentFile(engineID string) string {
	return fmt.Sprintf(`
resource "delphix_database_plugin" "test" {
  file_path = "/nonexistent/path/plugin.json"
  engine_host = "%s"
}
`, engineID)
}

func testAccEnginePluginNonJSONFile(engineID string) string {
	textFile := createTempTextFile()
	return fmt.Sprintf(`
resource "delphix_database_plugin" "test" {
  file_path = "%s"
  engine_host = "%s"
}
`, textFile, engineID)
}

func testAccEnginePluginInvalidJSON(engineID string) string {
	invalidJSONFile := createTempInvalidJSONFile()
	return fmt.Sprintf(`
resource "delphix_database_plugin" "test" {
  file_path = "%s"
  engine_host = "%s"
}
`, invalidJSONFile, engineID)
}

// Helper functions
func createTempTextFile() string {
	file, err := os.CreateTemp("", "test-plugin-*.txt")
	if err != nil {
		return "/tmp/test.txt"
	}

	file.WriteString("This is not a JSON file")
	file.Close()

	return file.Name()
}

func createTempInvalidJSONFile() string {
	file, err := os.CreateTemp("", "test-plugin-*.json")
	if err != nil {
		return "/tmp/invalid.json"
	}

	file.WriteString("{ invalid json content without closing brace")
	file.Close()

	return file.Name()
}
