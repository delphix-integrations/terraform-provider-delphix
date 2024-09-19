package provider

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccVdb_create_vdb_group_positive(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccVdbPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVdbGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDctVDBGroupConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDctVdbGroupResourceExists("delphix_vdb.new", "delphix_vdb_group.new_group")),
			},
		},
	})
}

func testAccCheckDctVDBGroupConfigBasic() string {
	datasource_id := os.Getenv("DATASOURCE_ID")
	return fmt.Sprintf(`
	resource "delphix_vdb" "new" {
		auto_select_repository = true
		source_data_id         = "%s"
	}
	resource "delphix_vdb_group" "new_group" {
		name                   = "my-vdb-group-name"
		vdb_ids                = [delphix_vdb.new.id]
	}
	`, datasource_id)
}

func testAccCheckDctVdbGroupResourceExists(vdbResourceName string, vdbGroupResourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		vdbGroupResource, ok := s.RootModule().Resources[vdbGroupResourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", vdbGroupResourceName)
		}

		vdbGroupId := vdbGroupResource.Primary.ID
		if vdbGroupId == "" {
			return fmt.Errorf("No VdbGroupID set")
		}

		vdbResource, ok := s.RootModule().Resources[vdbResourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", vdbResourceName)
		}

		vdbId := vdbResource.Primary.ID
		if vdbGroupId == "" {
			return fmt.Errorf("No VDbID set")
		}

		client := testAccProvider.Meta().(*apiClient).client

		res, _, err := client.VDBGroupsAPI.GetVdbGroup(context.Background(), vdbGroupId).Execute()
		if err != nil {
			return err
		}

		vdbIds := res.GetVdbIds()
		if !reflect.DeepEqual(vdbIds, []string{vdbId}) {
			return fmt.Errorf("Expected the vdb_id in VDB Group vdb_ids property")
		}

		return nil
	}
}

func testAccCheckVdbGroupDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*apiClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "delphix_vdb_group" {
			continue
		}

		vdbGroupId := rs.Primary.ID

		_, httpResp, _ := client.VDBGroupsAPI.GetVdbGroup(context.Background(), vdbGroupId).Execute()
		if httpResp == nil {
			return fmt.Errorf("VDB Group has not been deleted")
		}

		if httpResp.StatusCode != 404 {
			return fmt.Errorf("Exepcted a 404 Not Found for a deleted VDB Group but got %d", httpResp.StatusCode)
		}
	}

	return nil
}
