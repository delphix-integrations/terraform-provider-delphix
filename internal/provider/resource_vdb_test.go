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

func TestAccVdb_provision_positive(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccVdbPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVdbDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDctVDBConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDctVdbResourceExists("delphix_vdb.new"),
					resource.TestCheckResourceAttr("delphix_vdb.new", "parent_id", os.Getenv("DATASOURCE_ID"))),
				ExpectNonEmptyPlan: true,
			},
			{
				// positive update test case
				Config: testAccUpdatePositive("vdbu", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDctVdbResourceExists("delphix_vdb.new"),
					resource.TestCheckResourceAttr("delphix_vdb.new", "vdb_name", "vdbu"),
					resource.TestCheckResourceAttr("delphix_vdb.new", "vdb_restart", "true")),
				ExpectNonEmptyPlan: true, //due to the delay in GET call, we get some inconsistencies
			},
			{
				// negative update test case
				Config:      testAccUpdateNegative(false),
				ExpectError: regexp.MustCompile("Error running apply: exit status 1"),
			},
		},
	})
}

func testAccVdbPreCheck(t *testing.T) {
	testAccPreCheck(t)
	if err := os.Getenv("DATASOURCE_ID"); err == "" {
		t.Fatal("DATASOURCE_ID must be set for vdb acceptance tests")
	}
}

func testAccCheckDctVDBConfigBasic() string {
	datasource_id := os.Getenv("DATASOURCE_ID")
	return fmt.Sprintf(`
	resource "delphix_vdb" "new" {
		auto_select_repository = true
    	source_data_id         = "%s"
	}
	`, datasource_id)
}

func testAccCheckDctVdbResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vdbId := rs.Primary.ID
		if vdbId == "" {
			return fmt.Errorf("No VdbID set")
		}

		client := testAccProvider.Meta().(*apiClient).client

		res, _, err := client.VDBsApi.GetVdbById(context.Background(), vdbId).Execute()

		if err != nil {
			return err
		}

		parentId := res.GetParentId()
		if parentId != os.Getenv("DATASOURCE_ID") {
			return fmt.Errorf("parentId does not match DATASOURCE_ID")
		}

		return nil
	}
}

func testAccCheckVdbDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*apiClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "delphix_vdb" {
			continue
		}

		vdbId := rs.Primary.ID

		_, httpResp, _ := client.VDBsApi.GetVdbById(context.Background(), vdbId).Execute()

		if httpResp == nil {
			return fmt.Errorf("VDB has not been deleted")
		}

		if httpResp.StatusCode != 404 {
			return fmt.Errorf("Expected a 404 Not Found for a deleted VDB but got %d", httpResp.StatusCode)
		}
	}

	return nil
}
func testAccUpdateNegative(value bool) string {
	datasource_id := os.Getenv("DATASOURCE_ID")
	return fmt.Sprintf(`
	resource "delphix_vdb" "new" {
		auto_select_repository = "%t"
    	source_data_id         = "%s"
	}
	`, value, datasource_id)
}

func testAccUpdatePositive(name string, vdb_restart bool) string {
	datasource_id := os.Getenv("DATASOURCE_ID")
	return fmt.Sprintf(`
	resource "delphix_vdb" "new" {
		auto_select_repository = true
    	source_data_id         = "%s"
		vdb_name = "%s"
		vdb_restart = "%t"
	}
	`, datasource_id, name, vdb_restart)
}
