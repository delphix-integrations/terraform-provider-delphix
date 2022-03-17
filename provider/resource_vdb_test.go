package provider

import (
	"context"
	"fmt"
	"os"
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

		res, _, err := client.VDBsApi.GetVdbById(context.TODO(), vdbId).Execute()

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

		_, httpResp, _ := client.VDBsApi.GetVdbById(context.TODO(), vdbId).Execute()

		if httpResp == nil {
			return fmt.Errorf("VDB has not been deleted")
		}

		if httpResp.StatusCode != 404 {
			return fmt.Errorf("Exepcted a 404 Not Found for a deleted VDB but got %d", httpResp.StatusCode)
		}
	}

	return nil
}
