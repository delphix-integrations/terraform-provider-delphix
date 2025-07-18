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

// TestAccVdbGroup_create_positive tests the basic creation of a VDB group with a single VDB
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

// TestAccVdbGroup_tags tests various tag management scenarios:
// 1. Initial tag creation with multiple tags (including multiple values for same key)
// 2. Tag updates including:
//   - Modifying an existing tag's value (environment: test -> prod)
//   - Adding a new tag (owner: team-a)
//   - Removing a tag (purpose: testing)
//
// 3. Complete tag removal (setting tags to empty)
func TestAccVdbGroup_tags(t *testing.T) {
	var vdbGroupId string
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
					// Store the VDB group ID for later verification
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources["delphix_vdb_group.test"]
						if !ok {
							return fmt.Errorf("Not found: delphix_vdb_group.test")
						}
						vdbGroupId = rs.Primary.ID
						fmt.Printf("VDB Group ID: %s\n", vdbGroupId)
						return nil
					},
				),
			},
			{
				// Update tags - add new tag, modify existing tag, remove tag
				Config:       testAccCheckDctVDBGroupConfigWithTagsUpdated(),
				ResourceName: "delphix_vdb_group.test",
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
				Config:       testAccCheckDctVDBGroupConfigBasicWithName("my-vdb-group-name-2"),
				ResourceName: "delphix_vdb_group.test",
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("delphix_vdb_group.test", "tags.#", "0"),
				),
			},
		},
	})
}

// TestAccVdbGroup_update_name tests the ability to rename a VDB group
func TestAccVdbGroup_update_name(t *testing.T) {
	var vdbGroupId string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccVdbPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVdbGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDctVDBGroupConfigBasicWithName("my-vdb-group-original"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("delphix_vdb_group.test", "name", "my-vdb-group-original"),
					// Store the VDB group ID for later verification
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources["delphix_vdb_group.test"]
						if !ok {
							return fmt.Errorf("Not found: delphix_vdb_group.test")
						}
						vdbGroupId = rs.Primary.ID
						fmt.Printf("VDB Group ID: %s\n", vdbGroupId)
						return nil
					},
				),
			},
			{
				Config:       testAccCheckDctVDBGroupConfigBasicWithName("my-vdb-group-renamed"),
				ResourceName: "delphix_vdb_group.test",
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("delphix_vdb_group.test", "name", "my-vdb-group-renamed"),
				),
			},
		},
	})
}

// TestAccVdbGroup_update_vdb_ids tests VDB group membership changes:
// 1. Initial creation with one VDB
// 2. Update to use a different VDB
// 3. Update to use multiple VDBs
func TestAccVdbGroup_update_vdb_ids(t *testing.T) {
	var vdbGroupId string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccVdbGroupPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVdbGroupDestroy,
		Steps: []resource.TestStep{
			{
				// Create VDB group with first VDB
				Config: testAccCheckDctVDBGroupConfigWithTwoVdbs("my-vdb-group-vdbid", "vdb3"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("delphix_vdb_group.test", "vdb_ids.#", "1"),
					resource.TestCheckResourceAttrPair("delphix_vdb_group.test", "vdb_ids.0", "delphix_vdb.vdb3", "id"),
					// Store the VDB group ID for later verification
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources["delphix_vdb_group.test"]
						if !ok {
							return fmt.Errorf("Not found: delphix_vdb_group.test")
						}
						vdbGroupId = rs.Primary.ID
						fmt.Printf("VDB Group ID: %s\n", vdbGroupId)
						return nil
					},
				),
			},
			{
				// Update to use second VDB
				Config:       testAccCheckDctVDBGroupConfigWithTwoVdbs("my-vdb-group-vdbid", "vdb4"),
				ResourceName: "delphix_vdb_group.test",
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("delphix_vdb_group.test", "vdb_ids.#", "1"),
					resource.TestCheckResourceAttrPair("delphix_vdb_group.test", "vdb_ids.0", "delphix_vdb.vdb4", "id"),
				),
			},
			{
				// Update to use both VDBs
				Config:       testAccCheckDctVDBGroupConfigWithMultipleVdbs("my-vdb-group-vdbid"),
				ResourceName: "delphix_vdb_group.test",
				ImportState:  true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("delphix_vdb_group.test", "vdb_ids.#", "2"),
					resource.TestCheckResourceAttrSet("delphix_vdb.vdb3", "id"),
					resource.TestCheckResourceAttrSet("delphix_vdb.vdb4", "id"),
					resource.TestCheckResourceAttr("delphix_vdb.vdb3", "name", "TESTVDB3"),
					resource.TestCheckResourceAttr("delphix_vdb.vdb4", "name", "TESTVDB4"),
				),
			},
		},
	})
}

// Helper Functions

// testAccCheckDctVDBGroupConfigBasic creates a basic VDB group with a single VDB
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

// testAccCheckDctVDBGroupConfigWithTags creates a VDB group with initial tags:
// - environment: test
// - environment: dev (multiple values for same key)
// - purpose: testing
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
		ignore_tag_changes = false
	}
	`, datasource_id)
}

// testAccCheckDctVDBGroupConfigWithTagsUpdated updates the tags to:
// - environment: prod (modified value)
// - owner: team-a (new tag)
// Note: purpose: testing is removed
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
		ignore_tag_changes = false
	}
	`, datasource_id)
}

// testAccCheckDctVDBGroupConfigBasicWithName creates a VDB group with no tags
// Used for testing tag removal and name updates
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
		ignore_tag_changes     = false
	}
	`, datasource_id, groupName)
}

// testAccCheckDctVDBGroupConfigBasicWithVdbId creates a VDB group with a specific VDB
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

// testAccCheckDctVDBGroupConfigWithTwoVdbs creates a VDB group with one of two possible VDBs
func testAccCheckDctVDBGroupConfigWithTwoVdbs(groupName, vdbName string) string {
	datasource_id := os.Getenv("DATASOURCE_ID")
	datasource_id_2 := os.Getenv("DATASOURCE_ID_2")
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
	`, datasource_id, datasource_id_2, groupName, vdbName)
}

// testAccCheckDctVDBGroupConfigWithMultipleVdbs creates a VDB group with multiple VDBs
func testAccCheckDctVDBGroupConfigWithMultipleVdbs(groupName string) string {
	datasource_id := os.Getenv("DATASOURCE_ID")
	datasource_id_2 := os.Getenv("DATASOURCE_ID_2")
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
		vdb_ids                = [delphix_vdb.vdb3.id, delphix_vdb.vdb4.id]
	}
	`, datasource_id, datasource_id_2, groupName)
}

func testAccVdbGroupPreCheck(t *testing.T) {
	testAccPreCheck(t)
	if err := os.Getenv("DATASOURCE_ID"); err == "" {
		t.Fatal("DATASOURCE_ID must be set for vdb acceptance tests")
	}
	if err := os.Getenv("DATASOURCE_ID_2"); err == "" {
		t.Fatal("DATASOURCE_ID_2 must be set for vdb acceptance tests")
	}
}

// testAccCheckDctVdbGroupResourceExists verifies that a VDB group exists and contains the expected VDB
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

// testAccCheckVdbGroupDestroy verifies that a VDB group has been properly destroyed
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
