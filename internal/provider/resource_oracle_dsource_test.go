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

func TestOracleDsource_create_positive(t *testing.T) {
	sourcevalue := os.Getenv("ORACLE_DSOURCE_SOURCE_VALUE")
	groupId := os.Getenv("ORACLE_DSOURCE_GROUP_ID")
	name := os.Getenv("ORACLE_DSOURCE_NAME")
	environmentUser := os.Getenv("ORACLE_DSOURCE_ENV_USER")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testOracleDsourcePreCheck(t, sourcevalue, groupId, name, environmentUser)
		},
		Providers:    testAccProviders,
		CheckDestroy: testDsourceDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testOracleDsourceBasic(name, "", groupId, environmentUser),
				ExpectError: regexp.MustCompile(`.*`),
			},
			{
				Config: testOracleDsourceBasic(name, sourcevalue, groupId, environmentUser),
				Check: resource.ComposeTestCheckFunc(
					testOracleDsourceExists("delphix_oracle_dsource.test_oracle_dsource", sourcevalue),
					resource.TestCheckResourceAttr("delphix_oracle_dsource.test_oracle_dsource", "source_id", sourcevalue)),
			},
			{
				Config: testOracleDsourceBasic("update_name", sourcevalue, groupId, environmentUser),
				Check:  resource.ComposeAggregateTestCheckFunc(
				// irrelevant
				),
				ExpectError: regexp.MustCompile(`.*`),
			},
		},
	})
}

func testOracleDsourcePreCheck(t *testing.T, sourceId string, groupId string, name string, environmentUser string) {
	testAccPreCheck(t)
	if sourceId == "" {
		t.Fatal("ORACLE_DSOURCE_SOURCE_VALUE must be set for env acceptance tests")
	}
	if groupId == "" {
		t.Fatal("ORACLE_DSOURCE_GROUP_ID must be set for env acceptance tests")
	}
	if name == "" {
		t.Fatal("ORACLE_DSOURCE_NAME must be set for env acceptance tests")
	}
	if environmentUser == "" {
		t.Fatal("ORACLE_DSOURCE_ENV_USER must be set for env acceptance tests")
	}
}

func testOracleDsourceBasic(name string, sourceValue string, groupId string, environmentUser string) string {
	return fmt.Sprintf(`
resource "delphix_oracle_dsource" "test_oracle_dsource" {
  name                       = "%s"
  source_value               = "%s"
  group_id                   = "%s"
  log_sync_enabled           = false
  make_current_account_owner = true
  environment_user_id        = "%s"
  rman_channels              = 2
  files_per_set              = 5
  check_logical              = false
  encrypted_linking_enabled  = false
  compressed_linking_enabled = true
  bandwidth_limit            = 0
  number_of_connections      = 1
  diagnose_no_logging_faults = true
  pre_provisioning_enabled   = false
  link_now                   = true
  force_full_backup          = false
  double_sync                = false
  skip_space_check           = false
  do_not_resume              = false
  files_for_full_backup      = []
  log_sync_mode              = "UNDEFINED"
  log_sync_interval          = 5
}
	`, name, sourceValue, groupId, environmentUser)
}

func testOracleDsourceExists(n string, sourceValue string) resource.TestCheckFunc {
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
		if resSourceId != sourceValue {
			return fmt.Errorf("SourceId mismatch")
		}

		return nil
	}
}

func testOracleDsourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*apiClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "delphix_oracle_dsource" {
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
