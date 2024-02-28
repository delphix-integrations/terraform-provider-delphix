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

func Test_create_source_positive(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCReateSourceBasic(),
				Check: resource.ComposeTestCheckFunc(
					testSourceExists("delphix_vdb.new")),
			},
		},
	})
}

func testCReateSourceBasic() string {
	name := os.Getenv("SOURCE_NAME")
	repository_id := os.Getenv("REPOSITORY_ID")
	return fmt.Sprintf(`
	resource "delphix_source" "new" {
		name = "%s"
		repository_value         = "%s"
	}
	`, name, repository_id)
}

func testSourceExists(source string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		source, ok := s.RootModule().Resources[source]
		if !ok {
			return fmt.Errorf("Not found: %s", source)
		}

		sourceID := source.Primary.ID
		if sourceID == "" {
			return fmt.Errorf("No sourceID set")
		}

		client := testAccProvider.Meta().(*apiClient).client

		res, _, err := client.SourcesApi.GetSourceById(context.Background(), sourceID).Execute()
		if err != nil {
			return err
		}

		sourceids := res.GetId()
		if !reflect.DeepEqual(sourceids, []string{sourceID}) {
			return fmt.Errorf("Expected the vdb_id in VDB Group vdb_ids property")
		}

		return nil
	}
}

func testSourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*apiClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "delphix_source" {
			continue
		}

		sourceId := rs.Primary.ID

		_, httpResp, _ := client.SourcesApi.GetSourceById(context.Background(), sourceId).Execute()
		if httpResp == nil {
			return fmt.Errorf("Source has not been deleted")
		}

		if httpResp.StatusCode != 404 {
			return fmt.Errorf("Exepcted a 404 Not Found for a deleted Source but got %d", httpResp.StatusCode)
		}
	}

	return nil
}
