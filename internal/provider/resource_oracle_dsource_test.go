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

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testOracleDsourcePreCheck(t, sourcevalue, groupId, name)
		},
		Providers:    testAccProviders,
		CheckDestroy: testDsourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testOracleDsourceBasic(name, sourcevalue, groupId),
				Check: resource.ComposeTestCheckFunc(
					testOracleDsourceExists("delphix_oracle_dsource.test_oracle_dsource", sourcevalue),
					resource.TestCheckResourceAttr("delphix_oracle_dsource.test_oracle_dsource", "source_id", sourcevalue)),
			},
			{
				// positive update test case
				Config: testOracleDsourceBasic("update_name", sourcevalue, groupId),
				Check: resource.ComposeTestCheckFunc(
					testOracleDsourceExists("delphix_oracle_dsource.test_oracle_dsource", sourcevalue),
					resource.TestCheckResourceAttr("delphix_oracle_dsource.test_oracle_dsource", "name", "update_name"),
					resource.TestCheckResourceAttr("delphix_oracle_dsource.test_oracle_dsource", "group_id", groupId)),
			},
			{
				// negative update test case
				Config:      testOracleDsourceBasic(name, sourcevalue, "non-existent"),
				ExpectError: regexp.MustCompile("Error running apply: exit status 1"),
			},
			{
				Config:      testOracleDsourceBasic(name, "", groupId),
				ExpectError: regexp.MustCompile(`.*`),
			},
		},
	})
}

func testOracleDsourcePreCheck(t *testing.T, sourceId string, groupId string, name string) {
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
}

func testOracleDsourceBasic(name string, sourceValue string, groupId string) string {
	return fmt.Sprintf(`
resource "delphix_oracle_dsource" "test_oracle_dsource" {
  name                       = "%s"
  source_value               = "%s"
  group_id                   = "%s"
}
	`, name, sourceValue, groupId)
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
		res, _, err := client.DSourcesAPI.GetDsourceById(context.Background(), dsourceId).Execute()
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

		_, httpResp, _ := client.DSourcesAPI.GetDsourceById(context.Background(), dsourceId).Execute()
		if httpResp == nil {
			return fmt.Errorf("Dsource has not been deleted")
		}

		if httpResp.StatusCode != 404 {
			return fmt.Errorf("Exepcted a 404 Not Found for a deleted Dsource but got %d", httpResp.StatusCode)
		}
	}

	return nil
}
