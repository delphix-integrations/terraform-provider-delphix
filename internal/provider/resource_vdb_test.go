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
			},
			{
				// positive update test case
				Config: testAccUpdatePositive("vdbu", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDctVdbResourceExists("delphix_vdb.new"),
					resource.TestCheckResourceAttr("delphix_vdb.new", "name", "vdbu"),
					resource.TestCheckResourceAttr("delphix_vdb.new", "vdb_restart", "true")),
				ExpectNonEmptyPlan: true,
			},
			{
				// negative update test case
				Config:      testAccUpdateNegative(false),
				ExpectError: regexp.MustCompile("Error running apply: exit status 1"),
			},
		},
	})
}

func TestAccVdb_bookmark_provision(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccVdbBookmarkPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVdbDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDctVDBBookmarkConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDctVdbBookmarkResourceExists("delphix_vdb.new2")),
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

func testAccVdbBookmarkPreCheck(t *testing.T) {
	testAccPreCheck(t)
	if err := os.Getenv("DCT_BOOKMARK_ID"); err == "" {
		t.Fatal("DCT_BOOKMARK_ID must be set for vdb bookmark acceptance tests")
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

func testAccCheckDctVDBBookmarkConfigBasic() string {
	bookmark_id := os.Getenv("DCT_BOOKMARK_ID")
	return fmt.Sprintf(`
	resource "delphix_vdb" "new2" {
		provision_type         = "bookmark"
		auto_select_repository = true
    	bookmark_id            = "%s"
	}
	`, bookmark_id)
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

func testAccCheckDctVdbBookmarkResourceExists(n string) resource.TestCheckFunc {
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

		get_vdb_reponse, _, get_vdb_error := client.VDBsApi.GetVdbById(context.Background(), vdbId).Execute()

		if get_vdb_error != nil {
			return get_vdb_error
		}

		get_bookmark_response, _, get_bookmark_error := client.BookmarksApi.GetBookmarkById(context.Background(), os.Getenv("DCT_BOOKMARK_ID")).Execute()

		if get_bookmark_error != nil {
			return get_bookmark_error
		}

		sourceId := get_bookmark_response.GetVdbIds()[0]
		parentId := get_vdb_reponse.GetParentId()
		if parentId != sourceId {
			return fmt.Errorf("parentId does not match sourceId")
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
		name = "%s"
		vdb_restart = "%t"
	}
	`, datasource_id, name, vdb_restart)
}
