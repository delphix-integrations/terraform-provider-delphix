package provider

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"testing"

	dctapi "github.com/delphix/dct-sdk-go/v10"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
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
			},
			{
				// negative update test case
				Config:      testAccUpdateNegative(false),
				ExpectError: regexp.MustCompile("Error running apply: exit status 1"),
			},
		},
	})
}

var bookmark_id string
var vdb_id string

func TestAccVdb_bookmark_provision(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccVdbPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVdbDestroyBookmark,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDctVDBBookmarkConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDctVdbBookmarkResourceExists()),
			},
		},
	})
}

func TestAccVdb_appdata_provision(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccVdbAppDataPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVdbDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDctVDBConfigAppDataBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDctAppDataVdbResourceExists("delphix_vdb.new_appdata"),
					resource.TestCheckResourceAttr("delphix_vdb.new_appdata", "parent_id", os.Getenv("APPDATA_DATASOURCE_ID"))),
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

func testAccVdbAppDataPreCheck(t *testing.T) {
	testAccPreCheck(t)
	if err := os.Getenv("APPDATA_DATASOURCE_ID"); err == "" {
		t.Fatal("DATASOURCE_ID must be set for vdb acceptance tests")
	}
	if err := os.Getenv("APPDATA_SOURCE_PARAMS"); err == "" {
		t.Fatal("APPDATA_SOURCE_PARAMS must be set for vdb acceptance tests")
	}
	if err := os.Getenv("APPDATA_CONFIG_PARAMS"); err == "" {
		t.Fatal("APPDATA_CONFIG_PARAMS must be set for vdb acceptance tests")
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

func testAccCheckDctVDBConfigAppDataBasic() string {
	appdata_datasource_id := os.Getenv("APPDATA_DATASOURCE_ID")
	appdata_source_params := os.Getenv("APPDATA_SOURCE_PARAMS")
	appdata_config_params := os.Getenv("APPDATA_CONFIG_PARAMS")
	return fmt.Sprintf(`
	resource "delphix_vdb" "new_appdata" {
		auto_select_repository = true
    	source_data_id         = "%s"
		appdata_source_params  = jsonencode(%s)
		appdata_config_params  = jsonencode(%s)
	}
	`, appdata_datasource_id, appdata_source_params, appdata_config_params)
}

func testAccCheckDctVDBBookmarkConfigBasic() string {
	// init client
	cfg := dctapi.NewConfiguration()
	cfg.Host = os.Getenv("DCT_HOST")
	cfg.Scheme = "https"
	cfg.HTTPClient = &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}
	cfg.AddDefaultHeader("Authorization", "apk "+os.Getenv("DCT_KEY"))
	client := dctapi.NewAPIClient(cfg)

	// create vdb
	provisionVDBBySnapshotParameters := dctapi.NewProvisionVDBBySnapshotParameters()
	provisionVDBBySnapshotParameters.SetAutoSelectRepository(true)
	provisionVDBBySnapshotParameters.SetSourceDataId(os.Getenv("DATASOURCE_ID"))

	vdb_req := client.VDBsApi.ProvisionVdbBySnapshot(context.Background())

	vdb_res, vdb_http_res, vdb_err := vdb_req.ProvisionVDBBySnapshotParameters(*provisionVDBBySnapshotParameters).Execute()
	if diags := apiErrorResponseHelper(vdb_res, vdb_http_res, vdb_err); diags != nil {
		println("An error occured during vdb creation: " + vdb_err.Error())
		return "" // return empty config to indicate config error
	}
	vdb_id = *vdb_res.VdbId

	// poll for vdb
	vdb_job_res, vdb_job_err := PollJobStatus(*vdb_res.Job.Id, context.Background(), client)

	if vdb_job_res == Failed || vdb_job_res == Canceled || vdb_job_res == Abandoned {
		println("An error occured during vdb job polling " + vdb_job_err)
		return ""
	}

	//create bookmark
	bookmark := dctapi.NewBookmarkCreateParametersWithDefaults()
	bookmark.SetVdbIds([]string{vdb_id})
	bookmark.SetName(acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

	bookmark_req := client.BookmarksApi.CreateBookmark(context.Background()).BookmarkCreateParameters(*bookmark)
	bk_res, bk_http_res, bk_err := bookmark_req.Execute()

	if diags := apiErrorResponseHelper(bk_res, bk_http_res, bk_err); diags != nil {
		println("An error occured during bookmark creation: " + bk_err.Error())
		return ""
	}
	bookmark_id = *bk_res.Bookmark.Id

	// poll for bookmark
	bk_job_res, bk_job_err := PollJobStatus(*bk_res.Job.Id, context.Background(), client)

	if bk_job_res == Failed || bk_job_res == Canceled || bk_job_res == Abandoned {
		println("An error occured during bookmark job polling: " + bk_job_err)
		return "" // return empty config to indicate config error
	}

	resource := fmt.Sprintf(`
	resource "delphix_vdb" "vdb_bookmark" {
	provision_type         = "bookmark"
	auto_select_repository = true
	bookmark_id            = "%s"
	}
	`, bookmark_id)

	print(resource)

	return resource

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

func testAccCheckDctAppDataVdbResourceExists(n string) resource.TestCheckFunc {
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
		if parentId != os.Getenv("APPDATA_DATASOURCE_ID") {
			return fmt.Errorf("parentId does not match APPDATA_DATASOURCE_ID")
		}

		return nil
	}
}

func testAccCheckDctVdbBookmarkResourceExists() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources["delphix_vdb.vdb_bookmark"]

		if !ok {
			return fmt.Errorf("Not found: delphix_vdb.vdb_bookmark")
		}

		vdbId := rs.Primary.ID
		if vdbId == "" {
			return fmt.Errorf("No VdbID set")
		}

		client := testAccProvider.Meta().(*apiClient).client

		get_vdb_response, _, get_vdb_error := client.VDBsApi.GetVdbById(context.Background(), vdbId).Execute()

		if get_vdb_error != nil {
			return get_vdb_error
		}

		get_bookmark_response, _, get_bookmark_error := client.BookmarksApi.GetBookmarkById(context.Background(), bookmark_id).Execute()

		if get_bookmark_error != nil {
			return get_bookmark_error
		}

		sourceId := get_bookmark_response.GetVdbIds()[0]
		parentId := get_vdb_response.GetParentId()
		if parentId != sourceId {
			return fmt.Errorf("Single-VDB Bookmark's parentId does not match newly created VDB's sourceId")
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

func testAccCheckVdbDestroyBookmark(s *terraform.State) error {
	client := testAccProvider.Meta().(*apiClient).client

	print("Deleting parent vdb " + vdb_id)
	deleteVdbParams := dctapi.NewDeleteVDBParametersWithDefaults()
	deleteVdbParams.SetForce(false)
	client.VDBsApi.DeleteVdb(context.Background(), vdb_id).DeleteVDBParameters(*deleteVdbParams).Execute()

	return testAccCheckVdbDestroy(s)
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
