package provider

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func Test_source_create_positive(t *testing.T) {
	name := os.Getenv("SOURCE_NAME")
	repository_id := os.Getenv("REPOSITORY_VALUE")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testsourcePreCheck(t, repository_id, name)
		},
		Providers:    testAccProviders,
		CheckDestroy: testSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testsourceBasic("", name),
				ExpectError: regexp.MustCompile(`.*`),
			},
			{
				Config: testsourceBasic(repository_id, name),
				Check: resource.ComposeTestCheckFunc(
					testSourceExists("delphix_database_postgresql.new_dsource", name),
					resource.TestCheckResourceAttr("delphix_database_postgresql.new_dsource", "name", name)),
			},
			{
				Config: testsourceUpdate(repository_id, "update_name"),
				Check: resource.ComposeTestCheckFunc(
					testSourceExists("delphix_database_postgresql.new_dsource", "update_name"),
					resource.TestCheckResourceAttr("delphix_database_postgresql.new_dsource", "name", "update_name")),
			},
		},
	})
}

func testsourcePreCheck(t *testing.T, repo_value string, name string) {
	testAccPreCheck(t)
	if repo_value == "" {
		t.Fatal("REPOSITORY_VALUE must be set for env acceptance tests")
	}
	if name == "" {
		t.Fatal("SOURCE_NAME must be set for env acceptance tests")
	}
}

func testsourceBasic(repo_value string, name string) string {
	return fmt.Sprintf(`
resource "delphix_database_postgresql" "new_dsource" {
  repository_value                  = "%s"
  name                       = "%s"
}
	`, repo_value, name)
}

func testSourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*apiClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "delphix_database_postgresql" {
			continue
		}

		sourceId := rs.Primary.ID

		_, httpResp, _ := client.SourcesAPI.GetSourceById(context.Background(), sourceId).Execute()
		if httpResp == nil {
			return fmt.Errorf("Source has not been deleted")
		}

		if httpResp.StatusCode != 404 {
			return fmt.Errorf("Exepcted a 404 Not Found for a deleted Source but got %d", httpResp.StatusCode)
		}
	}

	return nil
}

func testsourceUpdate(repo_value string, name string) string {
	return fmt.Sprintf(`
resource "delphix_database_postgresql" "new_dsource" {
  repository_value                  = "%s"
  name                       = "%s"
}
	`, repo_value, name)
}

func testSourceExists(n string, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		sourceId := rs.Primary.ID
		if sourceId == "" {
			return fmt.Errorf("No sourceId set")
		}

		client := testAccProvider.Meta().(*apiClient).client
		res, _, err := client.SourcesAPI.GetSourceById(context.Background(), sourceId).Execute()
		if err != nil {
			return err
		}

		resSourceId := res.GetName()
		if resSourceId != name {
			return fmt.Errorf("SourceId mismatch")
		}

		return nil
	}
}
