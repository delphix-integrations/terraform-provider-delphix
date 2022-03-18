package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func new_provider() *schema.Provider {
	return Provider("acc-test-version")()
}

func init() {
	testAccProvider = new_provider()
	testAccProviders = map[string]*schema.Provider{
		"delphix": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := new_provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = new_provider()
}

func testAccPreCheck(t *testing.T) {
	if err := os.Getenv("DCT_KEY"); err == "" {
		t.Fatal("DCT_KEY must be set for acceptance tests")
	}
	if err := os.Getenv("DCT_HOST"); err == "" {
		t.Fatal("DCT_HOST must be set for acceptance tests")
	}
}
