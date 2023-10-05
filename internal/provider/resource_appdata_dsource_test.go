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

func TestDsource_create_positive(t *testing.T) {
	sourceId := os.Getenv("DSOURCE_SOURCE_ID")
	groupId := os.Getenv("DSOURCE_GROUP_ID")
	name := os.Getenv("DSOURCE_NAME")
	environmentUser := os.Getenv("DSOURCE_ENV_USER")
	stagingEnvironment := os.Getenv("DSOURCE_STAGE_ENV")
	postgresPort := os.Getenv("DSOURCE_POSTGRES_PORT")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testDsourcePreCheck(t, sourceId, groupId, name, environmentUser, stagingEnvironment, postgresPort)
		},
		Providers:    testAccProviders,
		CheckDestroy: testDsourceDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testDsourceBasic(sourceId, groupId, name, environmentUser, stagingEnvironment, ""),
				ExpectError: regexp.MustCompile(`.*`),
			},
			{
				Config: testDsourceBasic(sourceId, groupId, name, environmentUser, stagingEnvironment, postgresPort),
				Check: resource.ComposeTestCheckFunc(
					testDsourceExists("delphix_appdata_dsource.new_data_dsource", sourceId),
					resource.TestCheckResourceAttr("delphix_appdata_dsource.new_data_dsource", "source_id", sourceId)),
			},
			{
				Config: testDsourceUpdate(sourceId, groupId, "update_same_dsource", environmentUser, stagingEnvironment, postgresPort),
				Check:  resource.ComposeAggregateTestCheckFunc(
				// irrelevant
				),
				ExpectError: regexp.MustCompile(`.*`),
			},
		},
	})
}

func testDsourcePreCheck(t *testing.T, sourceId string, groupId string, name string, environmentUser string, stagingEnvironment string, postgresPort string) {
	testAccPreCheck(t)
	if sourceId == "" {
		t.Fatal("DSOURCE_SOURCE_ID must be set for env acceptance tests")
	}
	if groupId == "" {
		t.Fatal("DSOURCE_GROUP_ID must be set for env acceptance tests")
	}
	if name == "" {
		t.Fatal("DSOURCE_NAME must be set for env acceptance tests")
	}
	if environmentUser == "" {
		t.Fatal("DSOURCE_ENV_USER must be set for env acceptance tests")
	}
	if stagingEnvironment == "" {
		t.Fatal("DSOURCE_STAGE_ENV must be set for env acceptance tests")
	}
	if postgresPort == "" {
		t.Fatal("DSOURCE_POSTGRES_PORT must be set for env acceptance tests")
	}
}

func testDsourceBasic(sourceId string, groupId string, name string, environmentUser string, stagingEnvironment string, postgresPort string) string {
	return fmt.Sprintf(`
resource "delphix_appdata_dsource" "new_data_dsource" {
  source_id                  = "%s"
  group_id                   = "%s"
  log_sync_enabled           = false
  make_current_account_owner = true
  link_type                  = "AppDataStaged"
  name                       = "%s"
  staging_mount_base         = ""
  environment_user           = "%s"
  staging_environment        = "%s"
  parameters = jsonencode(%s)
  sync_parameters = jsonencode({
    resync = true
  })
}
	`, sourceId, groupId, name, environmentUser, stagingEnvironment, postgresPort)
}

func testDsourceUpdate(sourceId string, groupId string, name string, environmentUser string, stagingEnvironment string, postgresPort string) string {
	return fmt.Sprintf(`
resource "delphix_appdata_dsource" "new_data_dsource" {
  source_id                  = "%s"
  group_id                   = "%s"
  log_sync_enabled           = false
  make_current_account_owner = true
  link_type                  = "AppDataStaged"
  name                       = "%s"
  staging_mount_base         = ""
  environment_user           = "%s"
  staging_environment        = "%s"
  parameters = jsonencode(%s)
  sync_parameters = jsonencode({
    resync = true
  })
}
	`, sourceId, groupId, name, environmentUser, stagingEnvironment, postgresPort)
}

func testDsourceExists(n string, sourceId string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		dsourceId := rs.Primary.ID
		if dsourceId == "" {
			return fmt.Errorf("No dsourceId set")
		}

		client := testAccProvider.Meta().(*apiClient).client
		res, _, err := client.DSourcesApi.GetDsourceById(context.Background(), dsourceId).Execute()
		if err != nil {
			return err
		}

		resSourceId := res.GetSourceId()
		if resSourceId != sourceId {
			return fmt.Errorf("SourceId mismatch")
		}

		return nil
	}
}

func testDsourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*apiClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "delphix_appdata_dsource" {
			continue
		}

		dsourceId := rs.Primary.ID

		_, httpResp, _ := client.DSourcesApi.GetDsourceById(context.Background(), dsourceId).Execute()
		if httpResp == nil {
			return fmt.Errorf("Dsource has not been deleted")
		}

		if httpResp.StatusCode != 404 {
			return fmt.Errorf("Exepcted a 404 Not Found for a deleted Dsource but got %d", httpResp.StatusCode)
		}
	}

	return nil
}
