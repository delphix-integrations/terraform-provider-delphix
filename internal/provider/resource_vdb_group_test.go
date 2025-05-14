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

func TestAccVdbGroup_create_positive(t *testing.T) {
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

func TestAccVdbGroup_tags(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccVdbPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVdbGroupDestroy,
		Steps: []resource.TestStep{
			{
				// Create VDB group with initial tags
				Config: testAccCheckDctVDBGroupConfigWithTags(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDctVdbGroupResourceExists("delphix_vdb.test", "delphix_vdb_group.test"),
					resource.TestCheckResourceAttr("delphix_vdb_group.test", "tags.#", "3"),
					resource.TestCheckTypeSetElemNestedAttrs("delphix_vdb_group.test", "tags.*", map[string]string{
						"key":   "environment",
						"value": "test",
					}),
					resource.TestCheckTypeSetElemNestedAttrs("delphix_vdb_group.test", "tags.*", map[string]string{
						"key":   "environment",
						"value": "dev",
					}),
					resource.TestCheckTypeSetElemNestedAttrs("delphix_vdb_group.test", "tags.*", map[string]string{
						"key":   "purpose",
						"value": "testing",
					}),
				),
			},
			{
				// Update tags - add new tag, modify existing tag, remove tag
				Config: testAccCheckDctVDBGroupConfigWithTagsUpdated(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("delphix_vdb_group.test", "tags.#", "2"),
					resource.TestCheckTypeSetElemNestedAttrs("delphix_vdb_group.test", "tags.*", map[string]string{
						"key":   "environment",
						"value": "prod",
					}),
					resource.TestCheckTypeSetElemNestedAttrs("delphix_vdb_group.test", "tags.*", map[string]string{
						"key":   "owner",
						"value": "team-a",
					}),
				),
			},
			{
				// Remove all tags
				Config: testAccCheckDctVDBGroupConfigBasicWithName("my-vdb-group-name-2"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("delphix_vdb_group.test", "tags.#", "0"),
				),
			},
		},
	})
}

func TestAccVdbGroup_update_name(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccVdbPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVdbGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDctVDBGroupConfigBasicWithName("my-vdb-group-original"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("delphix_vdb_group.test", "name", "my-vdb-group-original"),
				),
			},
			{
				Config: testAccCheckDctVDBGroupConfigBasicWithName("my-vdb-group-renamed"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("delphix_vdb_group.test", "name", "my-vdb-group-renamed"),
				),
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
		name                   = "TESTVDB1"
		provision_type         = "snapshot"
	}
	resource "delphix_vdb_group" "new_group" {
		name                   = "my-vdb-group-name-1"
		vdb_ids                = [delphix_vdb.new.id]
	}
	`, datasource_id)
}

func testAccCheckDctVDBGroupConfigWithTags() string {
	datasource_id := os.Getenv("DATASOURCE_ID")
	return fmt.Sprintf(`
	resource "delphix_vdb" "test" {
		auto_select_repository = true
		source_data_id = "%s"
		name = "TESTVDB2"
		provision_type = "snapshot"
	}

	resource "delphix_vdb_group" "test" {
		name = "my-vdb-group-name-2"
		vdb_ids = [delphix_vdb.test.id]

		tags {
			key = "environment"
			value = "test"
		}
		tags {
			key = "environment"
			value = "dev"
		}
		tags {
			key = "purpose"
			value = "testing"
		}
	}
	`, datasource_id)
}

func testAccCheckDctVDBGroupConfigWithTagsUpdated() string {
	datasource_id := os.Getenv("DATASOURCE_ID")
	return fmt.Sprintf(`
	resource "delphix_vdb" "test" {
		auto_select_repository = true
		source_data_id = "%s"
		name = "TESTVDB2"
		provision_type = "snapshot"
	}

	resource "delphix_vdb_group" "test" {
		name = "my-vdb-group-name-2"
		vdb_ids = [delphix_vdb.test.id]

		tags {
			key = "environment"
			value = "prod"
		}
		tags {
			key = "owner"
			value = "team-a"
		}
	}
	`, datasource_id)
}

func testAccCheckDctVDBGroupConfigBasicWithName(groupName string) string {
	datasource_id := os.Getenv("DATASOURCE_ID")
	return fmt.Sprintf(`
	resource "delphix_vdb" "test" {
		auto_select_repository = true
		source_data_id         = "%s"
		name                   = "TESTVDB2"
		provision_type         = "snapshot"
	}
	resource "delphix_vdb_group" "test" {
		name                   = "%s"
		vdb_ids                = [delphix_vdb.test.id]
	}
	`, datasource_id, groupName)
}

func testAccCheckDctVDBGroupConfigBasicWithVdbId(vdbName string, groupName string) string {
	datasource_id := os.Getenv("DATASOURCE_ID")
	return fmt.Sprintf(`
	resource "delphix_vdb" "%s" {
		auto_select_repository = true
		source_data_id         = "%s"
		name                   = "%s"
		provision_type         = "snapshot"
	}
	resource "delphix_vdb_group" "test" {
		name                   = "%s"
		vdb_ids                = [delphix_vdb.%s.id]
	}
	`, vdbName, datasource_id, vdbName, groupName, vdbName)
}

func testAccCheckDctVDBGroupConfigWithTwoVdbs(groupName, vdbName string) string {
	datasource_id := os.Getenv("DATASOURCE_ID")
	return fmt.Sprintf(`
	resource "delphix_vdb" "vdb3" {
		auto_select_repository = true
		source_data_id         = "%s"
		name                   = "TESTVDB3"
		provision_type         = "snapshot"
	}
	resource "delphix_vdb" "vdb4" {
		auto_select_repository = true
		source_data_id         = "%s"
		name                   = "TESTVDB4"
		provision_type         = "snapshot"
	}
	resource "delphix_vdb_group" "test" {
		name                   = "%s"
		vdb_ids                = [delphix_vdb.%s.id]
	}
	`, datasource_id, datasource_id, groupName, vdbName)
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

func TestAccVdbGroup_update_vdb_ids(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccVdbPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVdbGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDctVDBGroupConfigWithTwoVdbs("my-vdb-group-vdbid", "vdb3"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("delphix_vdb_group.test", "vdb_ids.#", "1"),
				),
			},
			{
				Config: testAccCheckDctVDBGroupConfigWithTwoVdbs("my-vdb-group-vdbid", "vdb4"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("delphix_vdb_group.test", "vdb_ids.#", "1"),
				),
			},
		},
	})
}
