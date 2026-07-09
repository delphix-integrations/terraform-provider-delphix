package provider

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccEngineConfiguration_blockDevice(t *testing.T) {
	resourceName := "delphix_engine_configuration.test"
	engineHost := os.Getenv("DELPHIX_ENGINE_HOST")

	if engineHost == "" {
		t.Skip("DELPHIX_ENGINE_HOST environment variable not set")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccEngineConfigurationBlockDevice(engineHost),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEngineConfigurationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "engine_host", engineHost),
					resource.TestCheckResourceAttr(resourceName, "sys_user", "sysadmin"),
					resource.TestCheckResourceAttr(resourceName, "user", "admin"),
					resource.TestCheckResourceAttr(resourceName, "engine_type", "CD"),
					resource.TestCheckResourceAttr(resourceName, "device_type", "BLOCK"),
					resource.TestCheckResourceAttrSet(resourceName, "configured"),
					resource.TestCheckResourceAttrSet(resourceName, "hostname"),
					resource.TestCheckResourceAttrSet(resourceName, "product_type"),
				),
			},
		},
	})
}

func TestAccEngineConfiguration_objectStorageWithRole(t *testing.T) {
	resourceName := "delphix_engine_configuration.test"
	engineHost := os.Getenv("DELPHIX_ENGINE_HOST")
	bucketName := os.Getenv("S3_BUCKET_NAME")

	if engineHost == "" || bucketName == "" {
		t.Skip("DELPHIX_ENGINE_HOST or S3_BUCKET_NAME environment variable not set")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccEngineConfigurationObjectStorageRole(engineHost, bucketName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEngineConfigurationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "device_type", "OBJECT"),
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.region", "us-west-2"),
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.bucket", bucketName),
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.endpoint", "s3.us-west-2.amazonaws.com"),
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.size", "20GB"),
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.auth_type", "ROLE"),
					resource.TestCheckResourceAttr(resourceName, "ntp_servers.0", "pool.ntp.org"),
					resource.TestCheckResourceAttr(resourceName, "ntp_servers.1", "time.nist.gov"),
					resource.TestCheckResourceAttr(resourceName, "ntp_timezone", "America/New_York"),
				),
			},
		},
	})
}

func TestAccEngineConfiguration_objectStorageWithAccessKey(t *testing.T) {
	resourceName := "delphix_engine_configuration.test"
	engineHost := os.Getenv("DELPHIX_ENGINE_HOST")
	bucketName := os.Getenv("S3_BUCKET_NAME")
	accessId := os.Getenv("AWS_ACCESS_KEY_ID")
	accessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	if engineHost == "" || bucketName == "" || accessId == "" || accessKey == "" {
		t.Skip("Required environment variables not set: DELPHIX_ENGINE_HOST, S3_BUCKET_NAME, AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccEngineConfigurationObjectStorageAccessKey(engineHost, bucketName, accessId, accessKey),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEngineConfigurationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "device_type", "OBJECT"),
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.auth_type", "ACCESS_KEY"),
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.access_id", accessId),
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.access_key", accessKey),
					resource.TestCheckResourceAttr(resourceName, "ntp_servers.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "ntp_timezone", "UTC"),
				),
			},
		},
	})
}

func TestAccEngineConfiguration_gcpObjectStorage(t *testing.T) {
	resourceName := "delphix_engine_configuration.test"
	engineHost := os.Getenv("DELPHIX_ENGINE_HOST")
	bucketName := os.Getenv("GCP_BUCKET_NAME")

	if engineHost == "" || bucketName == "" {
		t.Skip("DELPHIX_ENGINE_HOST or GCP_BUCKET_NAME environment variable not set")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccEngineConfigurationGCPObjectStorage(engineHost, bucketName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEngineConfigurationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "device_type", "OBJECT"),
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.cloud_provider", "GCP"),
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.bucket", bucketName),
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.size", "20GB"),
					resource.TestCheckResourceAttr(resourceName, "ntp_servers.0", "pool.ntp.org"),
					resource.TestCheckResourceAttr(resourceName, "ntp_servers.1", "time.nist.gov"),
					resource.TestCheckResourceAttr(resourceName, "ntp_timezone", "America/New_York"),
				),
			},
		},
	})
}

func TestAccEngineConfiguration_gcpObjectStorage_CC(t *testing.T) {
	resourceName := "delphix_engine_configuration.test"
	engineHost := os.Getenv("DELPHIX_ENGINE_HOST")
	bucketName := os.Getenv("GCP_BUCKET_NAME")

	if engineHost == "" || bucketName == "" {
		t.Skip("DELPHIX_ENGINE_HOST or GCP_BUCKET_NAME environment variable not set (requires a CC engine)")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccEngineConfigurationGCPObjectStorageCC(engineHost, bucketName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEngineConfigurationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "engine_type", "CC"),
					resource.TestCheckResourceAttr(resourceName, "device_type", "OBJECT"),
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.cloud_provider", "GCP"),
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.bucket", bucketName),
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.size", "20GB"),
					resource.TestCheckResourceAttr(resourceName, "ntp_servers.0", "pool.ntp.org"),
					resource.TestCheckResourceAttr(resourceName, "ntp_servers.1", "time.nist.gov"),
					resource.TestCheckResourceAttr(resourceName, "ntp_timezone", "America/New_York"),
				),
			},
		},
	})
}

func TestAccEngineConfiguration_azureObjectStorageWithManagedIdentities(t *testing.T) {
	resourceName := "delphix_engine_configuration.test"
	engineHost := os.Getenv("DELPHIX_ENGINE_HOST")
	containerName := os.Getenv("AZURE_CONTAINER_NAME")
	azureAccount := os.Getenv("AZURE_ACCOUNT_NAME")

	if engineHost == "" || containerName == "" || azureAccount == "" {
		t.Skip("DELPHIX_ENGINE_HOST, AZURE_CONTAINER_NAME or AZURE_ACCOUNT_NAME environment variable not set")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccEngineConfigurationAzureObjectStorageManagedIdentities(engineHost, containerName, azureAccount),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEngineConfigurationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "device_type", "OBJECT"),
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.cloud_provider", "AZURE"),
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.azure_container", containerName),
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.azure_account", azureAccount),
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.size", "20GB"),
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.auth_type", "MANAGED_IDENTITIES"),
					resource.TestCheckResourceAttr(resourceName, "ntp_servers.0", "pool.ntp.org"),
					resource.TestCheckResourceAttr(resourceName, "ntp_servers.1", "time.nist.gov"),
					resource.TestCheckResourceAttr(resourceName, "ntp_timezone", "Europe/London"),
				),
			},
		},
	})
}

func TestAccEngineConfiguration_azureObjectStorageWithAccessKey(t *testing.T) {
	resourceName := "delphix_engine_configuration.test"
	engineHost := os.Getenv("DELPHIX_ENGINE_HOST")
	containerName := os.Getenv("AZURE_CONTAINER_NAME")
	azureAccount := os.Getenv("AZURE_ACCOUNT_NAME")
	azureKey := os.Getenv("AZURE_ACCESS_KEY")

	if engineHost == "" || containerName == "" || azureAccount == "" || azureKey == "" {
		t.Skip("Required environment variables not set: DELPHIX_ENGINE_HOST, AZURE_CONTAINER_NAME, AZURE_ACCOUNT_NAME, AZURE_ACCESS_KEY")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccEngineConfigurationAzureObjectStorageAccessKey(engineHost, containerName, azureAccount, azureKey),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEngineConfigurationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "device_type", "OBJECT"),
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.cloud_provider", "AZURE"),
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.auth_type", "ACCESS_KEY"),
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.azure_container", containerName),
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.azure_account", azureAccount),
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.azure_key", azureKey),
					resource.TestCheckResourceAttr(resourceName, "ntp_servers.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "ntp_timezone", "UTC"),
				),
			},
		},
	})
}

func TestAccEngineConfiguration_validationErrors(t *testing.T) {
	engineHost := "http://test-engine.example.com"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccEngineConfigurationObjectStorageMissingParams(engineHost),
				ExpectError: regexp.MustCompile("object_storage_params must be provided when device_type is OBJECT"),
			},
			{
				Config:      testAccEngineConfigurationObjectStorageMissingAccessKey(engineHost),
				ExpectError: regexp.MustCompile("access_id and access_key must be provided when auth_type is ACCESS_KEY"),
			},
			{
				Config:      testAccEngineConfigurationObjectStorageMissingNTP(engineHost),
				ExpectError: regexp.MustCompile("ntp_servers and ntp_timezone must be provided when device_type is OBJECT"),
			},
			{
				Config:      testAccEngineConfigurationInvalidStorageSize(engineHost),
				ExpectError: regexp.MustCompile("must be a valid storage size with units"),
			},
			{
				Config:      testAccEngineConfigurationAzureMissingKey(engineHost),
				ExpectError: regexp.MustCompile("azure_key must be provided when auth_type is ACCESS_KEY for AZURE cloud_provider"),
			},
			{
				Config:      testAccEngineConfigurationGCPMissingBucket(engineHost),
				ExpectError: regexp.MustCompile("bucket must be a non-empty string in object_storage_params for GCP cloud_provider"),
			},
			{
				Config:      testAccEngineConfigurationAzureMissingContainer(engineHost),
				ExpectError: regexp.MustCompile("azure_container must be provided in object_storage_params for AZURE cloud_provider"),
			},
			{
				Config:      testAccEngineConfigurationAzureMissingAccount(engineHost),
				ExpectError: regexp.MustCompile("azure_account must be provided in object_storage_params for AZURE cloud_provider"),
			},
		},
	})
}

func TestValidateStorageSize(t *testing.T) {
	validSizes := []string{
		"100GB",
		"1.5TB",
		"20TB",
		"0.5PB",
		"1000GB",
		"2.5TB",
		"100 GB",
		"20 TB",
	}

	invalidSizes := []string{
		"100",
		"100MB",
		"100KB",
		"abc",
		"100gb",
		"1.5.5TB",
		"TB",
		"",
	}

	for _, size := range validSizes {
		warnings, errors := validateStorageSize(size, "size")
		if len(errors) > 0 {
			t.Errorf("Expected %s to be valid, got errors: %v", size, errors)
		}
		if len(warnings) > 0 {
			t.Errorf("Expected %s to have no warnings, got: %v", size, warnings)
		}
	}

	for _, size := range invalidSizes {
		_, errors := validateStorageSize(size, "size")
		if len(errors) == 0 {
			t.Errorf("Expected %s to be invalid, but got no errors", size)
		}
	}
}

func TestNormalizeStorageSize(t *testing.T) {
	cases := map[string]string{
		"100GB":     "100GB",
		"100 GB":    "100GB",
		"  20 TB ":  "20TB",
		"1.5 PB":    "1.5PB",
		"1000   GB": "1000GB",
		"20\tGB":    "20GB",
		"20 \t GB":  "20GB",
	}
	for in, want := range cases {
		got := normalizeStorageSize(in)
		if got != want {
			t.Errorf("normalizeStorageSize(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestAccEngineConfiguration_comprehensive(t *testing.T) {
	resourceName := "delphix_engine_configuration.test"
	engineHost := os.Getenv("DELPHIX_ENGINE_HOST")
	bucketName := os.Getenv("S3_BUCKET_NAME")

	if engineHost == "" || bucketName == "" {
		t.Skip("DELPHIX_ENGINE_HOST or S3_BUCKET_NAME environment variable not set")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccEngineConfigurationComprehensive(engineHost, bucketName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEngineConfigurationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "device_type", "OBJECT"),
					// DNS Config
					resource.TestCheckResourceAttr(resourceName, "dns_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "dns_config.0.servers.#", "3"),
					// NTP Config
					resource.TestCheckResourceAttr(resourceName, "ntp_servers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "ntp_timezone", "America/New_York"),
					// SMTP Config
					resource.TestCheckResourceAttr(resourceName, "smtp_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "smtp_config.0.server", "smtp.example.com"),
					// Web Proxy Config
					resource.TestCheckResourceAttr(resourceName, "web_proxy_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "web_proxy_config.0.host", "proxy.internal.com"),
					// Analytics and Phone Home
					resource.TestCheckResourceAttr(resourceName, "phone_home_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "user_analytics_enabled", "true"),
					// Object Storage
					resource.TestCheckResourceAttr(resourceName, "object_storage_params.0.auth_type", "ROLE"),
				),
			},
		},
	})
}

func testAccEngineConfigurationComprehensive(engineHost, bucketName string) string {
	return fmt.Sprintf(`
resource "delphix_engine_configuration" "test" {
  engine_host               = "%s"
  sys_user                  = "sysadmin"
  sys_password              = "sysadmin"
  sys_new_password          = "delphix"
  user                      = "admin"
  password                  = "delphix"
  email                     = "test@example.com"
  engine_type               = "CD"
  device_type               = "OBJECT"
  
  # Object Storage Configuration
  object_storage_params {
    cloud_provider = "AWS"
    region    = "us-west-2"
    bucket    = "%s"
    endpoint  = "s3.us-west-2.amazonaws.com"
    size      = "30GB"
    auth_type = "ROLE"
  }
  
  # NTP Configuration
  ntp_servers  = ["pool.ntp.org", "time.nist.gov"]
  ntp_timezone = "America/New_York"
  
  # DNS Configuration
  dns_config {
    servers = ["172.16.105.22", "172.16.105.23", "8.8.8.8"]
    domains = ["example.com", "internal.local", "test.local"]
  }
  
  # SMTP Configuration
  smtp_config {
    server               = "smtp.example.com"
    port                 = 587
    from_email_address   = "noreply@example.com"
    tls_authentication   = true
    send_timeout         = 120
    
    smtp_authentication {
      user     = "smtp_user@example.com"
      password = "smtp_password"
    }
  }
  
  # Web Proxy Configuration
  web_proxy_config {
    host     = "proxy.internal.com"
    port     = 3128
    username = "proxy_admin"
    password = "proxy_secret"
  }
  
  # Analytics and Phone Home
  phone_home_enabled     = true
  user_analytics_enabled = true
}
`, engineHost, bucketName)
}

func testAccCheckEngineConfigurationExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		// In a real scenario, you might make an API call here to verify the resource exists
		return nil
	}
}

// Test configuration templates
func testAccEngineConfigurationBlockDevice(engineHost string) string {
	return fmt.Sprintf(`
resource "delphix_engine_configuration" "test" {
  engine_host      = "%s"
  sys_user         = "sysadmin"
  sys_password     = "sysadmin"
  sys_new_password = "delphix"
  user             = "admin"
  password         = "delphix"
  email        = "test@example.com"
  engine_type  = "CD"
  device_type  = "BLOCK"
}
`, engineHost)
}

func testAccEngineConfigurationObjectStorageRole(engineHost, bucketName string) string {
	return fmt.Sprintf(`
resource "delphix_engine_configuration" "test" {
  engine_host      = "%s"
  sys_user         = "sysadmin"
  sys_password     = "sysadmin"
  sys_new_password = "delphix"
  user             = "admin"
  password         = "delphix"
  email        = "test@example.com"
  engine_type  = "CD"
  device_type  = "OBJECT"
  
  ntp_servers  = ["pool.ntp.org", "time.nist.gov"]
  ntp_timezone = "America/New_York"
  
  object_storage_params {
	cloud_provider = "AWS"
    region    = "us-west-2"
    bucket    = "%s"
    endpoint  = "s3.us-west-2.amazonaws.com"
    size      = "20GB"
    auth_type = "ROLE"
  }
}
`, engineHost, bucketName)
}

func testAccEngineConfigurationObjectStorageAccessKey(engineHost, bucketName, accessId, accessKey string) string {
	return fmt.Sprintf(`
resource "delphix_engine_configuration" "test" {
  engine_host      = "%s"
  sys_user         = "sysadmin"
  sys_password     = "sysadmin"
  sys_new_password = "delphix"
  user             = "admin"
  password         = "delphix"
  email        = "test@example.com"
  engine_type  = "CD"
  device_type  = "OBJECT"
  
  ntp_servers  = ["pool.ntp.org", "time.nist.gov", "1.ubuntu.pool.ntp.org"]
  ntp_timezone = "UTC"
  
  object_storage_params {
	cloud_provider = "AWS"
    region     = "us-west-2"
    bucket     = "%s"
    endpoint   = "s3.us-west-2.amazonaws.com"
    size       = "20GB"
    auth_type  = "ACCESS_KEY"
    access_id  = "%s"
    access_key = "%s"
  }
}
`, engineHost, bucketName, accessId, accessKey)
}

func testAccEngineConfigurationObjectStorageMissingParams(engineHost string) string {
	return fmt.Sprintf(`
resource "delphix_engine_configuration" "test" {
  engine_host      = "%s"
  sys_user         = "sysadmin"
  sys_password     = "sysadmin"
  sys_new_password = "delphix"
  user             = "admin"
  password         = "delphix"
  email        = "test@example.com"
  engine_type  = "CD"
  device_type  = "OBJECT"
  # Missing object_storage_params
}
`, engineHost)
}

func testAccEngineConfigurationObjectStorageMissingAccessKey(engineHost string) string {
	return fmt.Sprintf(`
resource "delphix_engine_configuration" "test" {
  engine_host      = "%s"
  sys_user         = "sysadmin"
  sys_password     = "sysadmin"
  sys_new_password = "delphix"
  user             = "admin"
  password         = "delphix"
  email        = "test@example.com"
  engine_type  = "CD"
  device_type  = "OBJECT"
  
  ntp_servers  = ["pool.ntp.org"]
  ntp_timezone = "UTC"
  
  object_storage_params {
	cloud_provider = "AWS"
    region     = "us-west-2"
    bucket     = "test-bucket"
    endpoint   = "s3.us-west-2.amazonaws.com"
    size       = "20GB"
    auth_type  = "ACCESS_KEY"
    # Missing access_id and access_key
  }
}
`, engineHost)
}

func testAccEngineConfigurationObjectStorageMissingNTP(engineHost string) string {
	return fmt.Sprintf(`
resource "delphix_engine_configuration" "test" {
  engine_host      = "%s"
  sys_user         = "sysadmin"
  sys_password     = "sysadmin"
  sys_new_password = "delphix"
  user             = "admin"
  password         = "delphix"
  email        = "test@example.com"
  engine_type  = "CD"
  device_type  = "OBJECT"
  
  # Missing ntp_servers and ntp_timezone
  
  object_storage_params {
	cloud_provider = "AWS"
    region    = "us-west-2"
    bucket    = "test-bucket"
    endpoint  = "s3.us-west-2.amazonaws.com"
    size      = "20GB"
    auth_type = "ROLE"
  }
}
`, engineHost)
}

func testAccEngineConfigurationInvalidStorageSize(engineHost string) string {
	return fmt.Sprintf(`
resource "delphix_engine_configuration" "test" {
  engine_host      = "%s"
  sys_user         = "sysadmin"
  sys_password     = "sysadmin"
  sys_new_password = "delphix"
  user             = "admin"
  password         = "delphix"
  email        = "test@example.com"
  engine_type  = "CD"
  device_type  = "OBJECT"
  
  ntp_servers  = ["pool.ntp.org"]
  ntp_timezone = "UTC"
  
  object_storage_params {
	cloud_provider = "AWS"
    region    = "us-west-2"
    bucket    = "test-bucket"
    endpoint  = "s3.us-west-2.amazonaws.com"
    size      = "20MB"  # Invalid size unit
    auth_type = "ROLE"
  }
}
`, engineHost)
}

func testAccEngineConfigurationGCPObjectStorage(engineHost, bucketName string) string {
	return fmt.Sprintf(`
resource "delphix_engine_configuration" "test" {
  engine_host      = "%s"
  sys_user         = "sysadmin"
  sys_password     = "sysadmin"
  sys_new_password = "delphix"
  user             = "admin"
  password         = "delphix"
  email            = "test@example.com"
  engine_type      = "CD"
  device_type      = "OBJECT"
  
  ntp_servers  = ["pool.ntp.org", "time.nist.gov"]
  ntp_timezone = "America/New_York"
  
  object_storage_params {
    cloud_provider = "GCP"
    bucket = "%s"
    size   = "20GB"
  }
}
`, engineHost, bucketName)
}

func testAccEngineConfigurationGCPObjectStorageCC(engineHost, bucketName string) string {
	return fmt.Sprintf(`
resource "delphix_engine_configuration" "test" {
  engine_host             = "%s"
  sys_user                = "sysadmin"
  sys_password            = "sysadmin"
  sys_new_password        = "delphix"
  user                    = "admin"
  password                = "delphix"
  email                   = "test@example.com"
  engine_type             = "CC"
  compliance_user         = "admin"
  compliance_password     = "Admin-12"
  compliance_new_password = "Admin@45"
  compliance_email        = "compliance@example.com"
  device_type             = "OBJECT"

  ntp_servers  = ["pool.ntp.org", "time.nist.gov"]
  ntp_timezone = "America/New_York"

  object_storage_params {
    cloud_provider = "GCP"
    bucket = "%s"
    size   = "20GB"
  }
}
`, engineHost, bucketName)
}

func testAccEngineConfigurationGCPMissingBucket(engineHost string) string {
	return fmt.Sprintf(`
resource "delphix_engine_configuration" "test" {
  engine_host      = "%s"
  sys_user         = "sysadmin"
  sys_password     = "sysadmin"
  sys_new_password = "delphix"
  user             = "admin"
  password         = "delphix"
  email        = "test@example.com"
  engine_type  = "CD"
  device_type  = "OBJECT"
  
  ntp_servers  = ["pool.ntp.org"]
  ntp_timezone = "UTC"
  
  object_storage_params {
    cloud_provider = "GCP"
    # Missing bucket parameter
    size = "20GB"
  }
}
`, engineHost)
}

func testAccEngineConfigurationAzureObjectStorageManagedIdentities(engineHost, containerName, azureAccount string) string {
	return fmt.Sprintf(`
resource "delphix_engine_configuration" "test" {
  engine_host      = "%s"
  sys_user         = "sysadmin"
  sys_password     = "sysadmin"
  sys_new_password = "delphix"
  user             = "admin"
  password         = "delphix"
  email        = "test@example.com"
  engine_type  = "CD"
  device_type  = "OBJECT"
  
  ntp_servers  = ["pool.ntp.org", "time.nist.gov"]
  ntp_timezone = "Europe/London"
  
  object_storage_params {
    cloud_provider   = "AZURE"
    azure_container  = "%s"
    azure_account    = "%s"
    size            = "20GB"
    auth_type       = "MANAGED_IDENTITIES"
  }
}
`, engineHost, containerName, azureAccount)
}

func testAccEngineConfigurationAzureObjectStorageAccessKey(engineHost, containerName, azureAccount, azureKey string) string {
	return fmt.Sprintf(`
resource "delphix_engine_configuration" "test" {
  engine_host      = "%s"
  sys_user         = "sysadmin"
  sys_password     = "sysadmin"
  sys_new_password = "delphix"
  user             = "admin"
  password         = "delphix"
  email        = "test@example.com"
  engine_type  = "CD"
  device_type  = "OBJECT"
  
  ntp_servers  = ["pool.ntp.org", "time.nist.gov", "1.ubuntu.pool.ntp.org"]
  ntp_timezone = "UTC"
  
  object_storage_params {
    cloud_provider   = "AZURE"
    azure_container  = "%s"
    azure_account    = "%s"
    azure_key       = "%s"
    size            = "20GB"
    auth_type       = "ACCESS_KEY"
  }
}
`, engineHost, containerName, azureAccount, azureKey)
}

func testAccEngineConfigurationAzureMissingContainer(engineHost string) string {
	return fmt.Sprintf(`
resource "delphix_engine_configuration" "test" {
  engine_host      = "%s"
  sys_user         = "sysadmin"
  sys_password     = "sysadmin"
  sys_new_password = "delphix"
  user             = "admin"
  password         = "delphix"
  email        = "test@example.com"
  engine_type  = "CD"
  device_type  = "OBJECT"
  
  ntp_servers  = ["pool.ntp.org"]
  ntp_timezone = "UTC"
  
  object_storage_params {
    cloud_provider = "AZURE"
    # Missing azure_container parameter
    azure_account = "test-account"
    size = "20GB"
    auth_type = "MANAGED_IDENTITIES"
  }
}
`, engineHost)
}

func testAccEngineConfigurationAzureMissingAccount(engineHost string) string {
	return fmt.Sprintf(`
resource "delphix_engine_configuration" "test" {
  engine_host      = "%s"
  sys_user         = "sysadmin"
  sys_password     = "sysadmin"
  sys_new_password = "delphix"
  user             = "admin"
  password         = "delphix"
  email        = "test@example.com"
  engine_type  = "CD"
  device_type  = "OBJECT"
  
  ntp_servers  = ["pool.ntp.org"]
  ntp_timezone = "UTC"
  
  object_storage_params {
    cloud_provider   = "AZURE"
    azure_container  = "test-container"
    # Missing azure_account parameter
    size = "20GB"
    auth_type = "MANAGED_IDENTITIES"
  }
}
`, engineHost)
}

func testAccEngineConfigurationAzureMissingKey(engineHost string) string {
	return fmt.Sprintf(`
resource "delphix_engine_configuration" "test" {
  engine_host      = "%s"
  sys_user         = "sysadmin"
  sys_password     = "sysadmin"
  sys_new_password = "delphix"
  user             = "admin"
  password         = "delphix"
  email        = "test@example.com"
  engine_type  = "CD"
  device_type  = "OBJECT"

  ntp_servers  = ["pool.ntp.org"]
  ntp_timezone = "UTC"

  object_storage_params {
    cloud_provider   = "AZURE"
    azure_container  = "test-container"
    azure_account    = "test-account"
    size            = "20GB"
    auth_type       = "ACCESS_KEY"
    # Missing azure_key parameter
  }
}
`, engineHost)
}

// =============================================================================
// Unit tests (no live engine required)
//
// The acceptance tests above only run with TF_ACC=1 against a real engine, so
// they contribute nothing to unit-test coverage of resource_engine_configuration.go.
// The tests below exercise the schema, the CustomizeDiff validation closure,
// and the Read/Update/Delete handlers directly so the file is covered by
// `make test`.
// =============================================================================

// TestEngineConfigurationSchemaInternalValidate runs the SDK's schema
// validation over the full resource definition, exercising every schema entry
// and guarding against malformed schema (bad defaults, conflicting
// Required/Computed, etc.).
func TestEngineConfigurationSchemaInternalValidate(t *testing.T) {
	res := resourceEngineConfiguration()
	if err := res.InternalValidate(nil, true); err != nil {
		t.Fatalf("resourceEngineConfiguration schema InternalValidate failed: %v", err)
	}

	if res.CreateContext == nil || res.ReadContext == nil ||
		res.UpdateContext == nil || res.DeleteContext == nil {
		t.Error("expected all CRUD context handlers to be set")
	}
	if res.CustomizeDiff == nil {
		t.Error("expected CustomizeDiff to be set")
	}
	if _, ok := res.Schema["object_storage_params"]; !ok {
		t.Error("expected object_storage_params in schema")
	}
}

// runCustomizeDiff drives the real CustomizeDiff closure through Resource.Diff
// with a raw config and returns any error it produces.
func runCustomizeDiff(t *testing.T, raw map[string]interface{}) error {
	t.Helper()
	res := resourceEngineConfiguration()
	cfg := terraform.NewResourceConfigRaw(raw)
	_, err := res.Diff(context.Background(), nil, cfg, nil)
	return err
}

// baseObjectConfig returns a minimal-but-valid OBJECT/AWS/ROLE config that
// passes CustomizeDiff; individual test cases mutate it to trigger failures.
func baseObjectConfig() map[string]interface{} {
	return map[string]interface{}{
		"engine_host":      "http://engine.example.com",
		"sys_user":         "sysadmin",
		"sys_password":     "sysadmin",
		"sys_new_password": "delphix",
		"user":             "admin",
		"password":         "delphix",
		"email":            "test@example.com",
		"engine_type":      "CD",
		"device_type":      "OBJECT",
		"ntp_servers":      []interface{}{"pool.ntp.org"},
		"ntp_timezone":     "UTC",
		"object_storage_params": []interface{}{
			map[string]interface{}{
				"cloud_provider": "AWS",
				"region":         "us-west-2",
				"bucket":         "my-bucket",
				"endpoint":       "s3.us-west-2.amazonaws.com",
				"size":           "20GB",
				"auth_type":      "ROLE",
			},
		},
	}
}

func TestEngineConfigCustomizeDiff(t *testing.T) {
	tests := []struct {
		name      string
		mutate    func(m map[string]interface{})
		wantErr   bool
		errSubstr string
	}{
		{
			name:    "valid object/aws/role passes",
			mutate:  func(m map[string]interface{}) {},
			wantErr: false,
		},
		{
			name: "block device skips object validation",
			mutate: func(m map[string]interface{}) {
				m["device_type"] = "BLOCK"
				delete(m, "object_storage_params")
				delete(m, "ntp_servers")
				delete(m, "ntp_timezone")
			},
			wantErr: false,
		},
		{
			name: "object without object_storage_params errors",
			mutate: func(m map[string]interface{}) {
				delete(m, "object_storage_params")
			},
			wantErr:   true,
			errSubstr: "object_storage_params must be provided when device_type is OBJECT",
		},
		{
			name: "object without ntp errors",
			mutate: func(m map[string]interface{}) {
				delete(m, "ntp_servers")
				delete(m, "ntp_timezone")
			},
			wantErr:   true,
			errSubstr: "ntp_servers and ntp_timezone must be provided when device_type is OBJECT",
		},
		{
			name: "unknown cloud provider errors",
			mutate: func(m map[string]interface{}) {
				osp := m["object_storage_params"].([]interface{})[0].(map[string]interface{})
				osp["cloud_provider"] = "ORACLE"
			},
			wantErr: true,
			// Schema ValidateFunc does not run during Diff, so the unknown
			// provider surfaces from cloudProviderFor inside CustomizeDiff.
			errSubstr: `unsupported cloud_provider "ORACLE"`,
		},
		{
			name: "aws access_key without credentials errors",
			mutate: func(m map[string]interface{}) {
				osp := m["object_storage_params"].([]interface{})[0].(map[string]interface{})
				osp["auth_type"] = "ACCESS_KEY"
			},
			wantErr:   true,
			errSubstr: "access_id and access_key must be provided when auth_type is ACCESS_KEY",
		},
		{
			name: "azure missing container errors",
			mutate: func(m map[string]interface{}) {
				m["object_storage_params"] = []interface{}{
					map[string]interface{}{
						"cloud_provider": "AZURE",
						"azure_account":  "acct",
						"size":           "20GB",
						"auth_type":      "MANAGED_IDENTITIES",
					},
				}
			},
			wantErr:   true,
			errSubstr: "azure_container must be provided in object_storage_params for AZURE cloud_provider",
		},
		{
			name: "gcp missing bucket errors",
			mutate: func(m map[string]interface{}) {
				m["object_storage_params"] = []interface{}{
					map[string]interface{}{
						"cloud_provider": "GCP",
						"size":           "20GB",
					},
				}
			},
			wantErr:   true,
			errSubstr: "bucket must be a non-empty string in object_storage_params for GCP cloud_provider",
		},
		{
			name: "smtp config with authentication passes",
			mutate: func(m map[string]interface{}) {
				m["smtp_config"] = []interface{}{
					map[string]interface{}{
						"server":             "smtp.example.com",
						"port":               587,
						"from_email_address": "noreply@example.com",
						"smtp_authentication": []interface{}{
							map[string]interface{}{
								"user":     "smtp_user",
								"password": "smtp_pass",
							},
						},
					},
				}
			},
			wantErr: false,
		},
		{
			name: "continuous compliance without compliance fields errors",
			mutate: func(m map[string]interface{}) {
				m["engine_type"] = "CC"
			},
			wantErr:   true,
			errSubstr: "must be provided when engine_type is CONTINUOUS_COMPLIANCE",
		},
		{
			name: "continuous compliance with all compliance fields passes",
			mutate: func(m map[string]interface{}) {
				m["engine_type"] = "CC"
				m["compliance_user"] = "admin"
				m["compliance_password"] = "Admin-12"
				m["compliance_new_password"] = "Admin@45"
				m["compliance_email"] = "compliance@example.com"
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raw := baseObjectConfig()
			tt.mutate(raw)
			err := runCustomizeDiff(t, raw)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tt.errSubstr)
				}
				if tt.errSubstr != "" && !strings.Contains(err.Error(), tt.errSubstr) {
					t.Fatalf("expected error containing %q, got: %v", tt.errSubstr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
		})
	}
}

// engineStub is an in-memory stand-in for an engine's CDB API used to drive
// engineConfigRead without a live engine.
type engineStub struct {
	sessionStatus int
	loginStatus   int
	systemStatus  int
	systemBody    string
}

func (s engineStub) server() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc(ENGINE_APIS["SESSION"], func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(orDefaultStatus(s.sessionStatus))
	})
	mux.HandleFunc(ENGINE_APIS["LOGIN"], func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(orDefaultStatus(s.loginStatus))
	})
	mux.HandleFunc(ENGINE_APIS["SYSTEM_INFO"], func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(orDefaultStatus(s.systemStatus))
		body := s.systemBody
		if body == "" {
			body = `{"type":"OKResult","status":"OK","result":{}}`
		}
		_, _ = w.Write([]byte(body))
	})
	return httptest.NewServer(mux)
}

func orDefaultStatus(v int) int {
	if v == 0 {
		return http.StatusOK
	}
	return v
}

func newEngineConfigData(t *testing.T, id string, extra map[string]interface{}) *schema.ResourceData {
	t.Helper()
	res := resourceEngineConfiguration()
	raw := map[string]interface{}{
		"engine_host":      id,
		"sys_user":         "sysadmin",
		"sys_password":     "sysadmin",
		"sys_new_password": "delphix",
		"user":             "admin",
		"password":         "delphix",
		"email":            "test@example.com",
		"engine_type":      "CD",
		"device_type":      "BLOCK",
	}
	for k, v := range extra {
		raw[k] = v
	}
	d := schema.TestResourceDataRaw(t, res.Schema, raw)
	d.SetId(id)
	return d
}

func TestEngineConfigReadSuccess(t *testing.T) {
	stub := engineStub{
		systemBody: `{"type":"OKResult","status":"OK","result":{
			"configured":true,
			"hostname":"engine-host",
			"productType":"standard",
			"ssoEnabled":false,
			"vendorName":"Delphix",
			"productName":"Delphix Engine",
			"platform":"linux",
			"kernelName":"Linux"
		}}`,
	}
	srv := stub.server()
	defer srv.Close()

	d := newEngineConfigData(t, srv.URL, nil)
	diags := engineConfigRead(context.Background(), d, nil)
	if diags.HasError() {
		t.Fatalf("engineConfigRead returned diagnostics: %+v", diags)
	}

	if got := d.Get("hostname").(string); got != "engine-host" {
		t.Errorf("hostname = %q, want engine-host", got)
	}
	if got := d.Get("configured").(bool); !got {
		t.Errorf("configured = %v, want true", got)
	}
	if got := d.Get("product_name").(string); got != "Delphix Engine" {
		t.Errorf("product_name = %q, want Delphix Engine", got)
	}
	if got := d.Get("platform").(string); got != "linux" {
		t.Errorf("platform = %q, want linux", got)
	}
}

func TestEngineConfigReadCustomAPIVersion(t *testing.T) {
	stub := engineStub{}
	srv := stub.server()
	defer srv.Close()

	// Provide an explicit api_version to exercise the non-default branch.
	d := newEngineConfigData(t, srv.URL, map[string]interface{}{"api_version": "1.11.40"})
	diags := engineConfigRead(context.Background(), d, nil)
	if diags.HasError() {
		t.Fatalf("engineConfigRead returned diagnostics: %+v", diags)
	}
}

func TestEngineConfigReadErrors(t *testing.T) {
	tests := []struct {
		name      string
		stub      engineStub
		badHost   bool
		errSubstr string
	}{
		{
			name:      "session failure",
			stub:      engineStub{sessionStatus: http.StatusInternalServerError},
			errSubstr: "Error starting session",
		},
		{
			name:      "login failure",
			stub:      engineStub{loginStatus: http.StatusUnauthorized},
			errSubstr: "Error logging in",
		},
		{
			name:      "system info unmarshal failure",
			stub:      engineStub{systemBody: "not-json"},
			errSubstr: "Error unmarshalling system info response",
		},
		{
			name:      "unreachable host",
			badHost:   true,
			errSubstr: "Error starting session",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var host string
			if tt.badHost {
				host = "http://127.0.0.1:1" // nothing listening
			} else {
				srv := tt.stub.server()
				defer srv.Close()
				host = srv.URL
			}

			d := newEngineConfigData(t, host, nil)
			diags := engineConfigRead(context.Background(), d, nil)
			if !diags.HasError() {
				t.Fatalf("expected diagnostics error, got none")
			}
			if !strings.Contains(diags[0].Summary, tt.errSubstr) {
				t.Fatalf("expected error containing %q, got: %s", tt.errSubstr, diags[0].Summary)
			}
		})
	}
}

func TestEngineConfigUpdateNotSupported(t *testing.T) {
	d := newEngineConfigData(t, "http://engine.example.com", nil)
	diags := engineConfigUpdate(context.Background(), d, nil)
	if !diags.HasError() {
		t.Fatal("expected engineConfigUpdate to return an error")
	}
	if !strings.Contains(diags[0].Summary, "Action update not available for engine config") {
		t.Errorf("unexpected error: %s", diags[0].Summary)
	}
}

func TestEngineConfigDeleteNotSupported(t *testing.T) {
	t.Setenv("TF_ACC", "")
	d := newEngineConfigData(t, "http://engine.example.com", nil)
	diags := engineConfigDelete(context.Background(), d, nil)
	if !diags.HasError() {
		t.Fatal("expected engineConfigDelete to return an error when TF_ACC is unset")
	}
	if !strings.Contains(diags[0].Summary, "Action delete not available for engine config") {
		t.Errorf("unexpected error: %s", diags[0].Summary)
	}
}

func TestEngineConfigDeleteUnderAcceptance(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	d := newEngineConfigData(t, "http://engine.example.com", nil)
	diags := engineConfigDelete(context.Background(), d, nil)
	if diags.HasError() {
		t.Fatalf("expected engineConfigDelete to succeed under TF_ACC, got: %+v", diags)
	}
	if d.Id() != "" {
		t.Errorf("expected ID to be cleared under TF_ACC, got %q", d.Id())
	}
}

// =============================================================================
// engineConfigCreate
//
// Create drives the full engine-initialization flow over ~17 distinct HTTP
// calls.  We mock all of them with an in-memory engine so the real handler
// runs end-to-end without a live engine.
//
// NOTE on timing: engineConfigCreate calls time.Sleep(10s) before
// initialization and time.Sleep(60s) afterwards. Those sleeps live in the
// production handler and cannot be stubbed from a test. The full-flow test
// therefore takes ~80s of real wall-clock time, which is why `make test` must
// be run with an explicit longer -timeout. The tests are marked t.Parallel()
// so their sleeps overlap.
// =============================================================================

// mockEngineOptions controls how the in-memory engine responds, so a single
// handler can serve the happy path and the various failure paths.
type mockEngineOptions struct {
	complianceUpdateFails bool // compliance user update returns an errorMessage
	initFails             bool // SYSTEM_INITIALIZATION returns an ERROR status
}

func newMockEngine(opts mockEngineOptions) *httptest.Server {
	const okResult = `{"type":"OKResult","status":"OK","action":"","job":"","result":""}`
	writeOK := func(w http.ResponseWriter) { _, _ = w.Write([]byte(okResult)) }

	mux := http.NewServeMux()

	// Session / auth (used repeatedly).
	mux.HandleFunc(ENGINE_APIS["SESSION"], func(w http.ResponseWriter, r *http.Request) { _, _ = w.Write([]byte(`{}`)) })
	mux.HandleFunc(ENGINE_APIS["LOGIN"], func(w http.ResponseWriter, r *http.Request) { _, _ = w.Write([]byte(`{}`)) })

	// Continuous Compliance setup.
	mux.HandleFunc(ENGINE_APIS["START_MASKING"], func(w http.ResponseWriter, r *http.Request) { writeOK(w) })
	mux.HandleFunc(ENGINE_APIS["COMPLIANCE_LOGIN"], func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"Authorization":"masking-token-123"}`))
	})
	complianceUsers := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			_, _ = w.Write([]byte(`{"responseList":[{"userName":"admin","userId":5,"userStatus":"ACTIVE"}]}`))
			return
		}
		if opts.complianceUpdateFails { // PUT update
			_, _ = w.Write([]byte(`{"errorMessage":"compliance update rejected"}`))
			return
		}
		_, _ = w.Write([]byte(`{"userName":"admin","userStatus":"ACTIVE"}`)) // PUT update
	}
	mux.HandleFunc(ENGINE_APIS["COMPLIANCE_USER"], complianceUsers)
	mux.HandleFunc(ENGINE_APIS["COMPLIANCE_USER"]+"/", complianceUsers)

	// Optional service configuration tasks.
	mux.HandleFunc(ENGINE_APIS["SMTP_CONFIG"], func(w http.ResponseWriter, r *http.Request) { writeOK(w) })
	mux.HandleFunc(ENGINE_APIS["PHONE_HOME_CONFIG"], func(w http.ResponseWriter, r *http.Request) { writeOK(w) })
	mux.HandleFunc(ENGINE_APIS["USER_ANALYTICS_CONFIG"], func(w http.ResponseWriter, r *http.Request) { writeOK(w) })
	mux.HandleFunc(ENGINE_APIS["WEB_PROXY_CONFIG"], func(w http.ResponseWriter, r *http.Request) { writeOK(w) })
	mux.HandleFunc(ENGINE_APIS["DNS_CONFIG"], func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			_, _ = w.Write([]byte(`{"status":"OK","result":{"type":"DNSConfig","servers":[],"domain":[]}}`))
			return
		}
		writeOK(w)
	})
	mux.HandleFunc(ENGINE_APIS["NTP_CONFIG"], func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			_, _ = w.Write([]byte(`{"status":"OK","result":{"type":"TimeConfig","systemTimeZone":"UTC","ntpConfig":{"type":"NTPConfig","servers":[]}}}`))
			return
		}
		writeOK(w)
	})

	// Storage / initialization.
	mux.HandleFunc(ENGINE_APIS["STORAGE_DEVICE"], func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"status":"OK","result":[{"type":"StorageDevice","reference":"STORAGE_DEVICE-1","configured":false,"size":1000}]}`))
	})
	mux.HandleFunc(ENGINE_APIS["OBJECT_STORE_TEST_CONNECTION"], func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"status":"OK","result":{"type":"ObjectStoreTestResult","result":true}}`))
	})
	mux.HandleFunc(ENGINE_APIS["SYSTEM_INITIALIZATION"], func(w http.ResponseWriter, r *http.Request) {
		if opts.initFails {
			_, _ = w.Write([]byte(`{"type":"ErrorResult","status":"ERROR","error":{"details":"init blew up"}}`))
			return
		}
		_, _ = w.Write([]byte(`{"type":"OKResult","status":"OK","action":"ACTION-1","job":"","result":""}`))
	})
	// Poll: report COMPLETED so pollActionStatus succeeds and the handler
	// proceeds through the post-init tail (set engine type, SSO, password
	// updates, read).
	mux.HandleFunc(ENGINE_APIS["ACTION"], func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"status":"OK","result":{"type":"ActionResult","state":"COMPLETED","reference":"ACTION-1"}}`))
	})

	// Post-init endpoints.
	mux.HandleFunc(ENGINE_APIS["SYSTEM_INFO"], func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			_, _ = w.Write([]byte(`{"type":"OKResult","status":"OK","result":{"configured":true,"hostname":"h"}}`))
			return
		}
		writeOK(w)
	})
	mux.HandleFunc(ENGINE_APIS["SSO_CONFIG"], func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			_, _ = w.Write([]byte(`{"status":"OK","result":{"entityId":"entity-1"}}`))
			return
		}
		writeOK(w)
	})
	mux.HandleFunc(ENGINE_APIS["USER"], func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/current") {
			_, _ = w.Write([]byte(`{"status":"OK","result":{"type":"User","reference":"USER-2"}}`))
			return
		}
		writeOK(w)
	})

	return httptest.NewServer(mux)
}

// newCreateData builds the ResourceData passed to engineConfigCreate. It uses a
// fully-populated Continuous Compliance + OBJECT(AWS) config with every
// optional service-configuration block enabled, so the create handler executes
// the CC branch and every config task closure.
func newCreateData(t *testing.T, engineHost, engineType string) *schema.ResourceData {
	t.Helper()
	res := resourceEngineConfiguration()
	raw := map[string]interface{}{
		"engine_host":            engineHost,
		"sys_user":               "sysadmin",
		"sys_password":           "sysadmin",
		"sys_new_password":       "delphix",
		"user":                   "admin",
		"password":               "delphix",
		"email":                  "test@example.com",
		"engine_type":            engineType,
		"device_type":            "OBJECT",
		"phone_home_enabled":     true,
		"user_analytics_enabled": true,
		"ntp_servers":            []interface{}{"pool.ntp.org", "time.nist.gov"},
		"ntp_timezone":           "UTC",
		"object_storage_params": []interface{}{
			map[string]interface{}{
				"cloud_provider": "AWS",
				"region":         "us-west-2",
				"bucket":         "my-bucket",
				"endpoint":       "s3.us-west-2.amazonaws.com",
				"size":           "20GB",
				"auth_type":      "ACCESS_KEY",
				"access_id":      "AKIA-TEST",
				"access_key":     "secret-test",
			},
		},
		"smtp_config": []interface{}{
			map[string]interface{}{
				"server":             "smtp.example.com",
				"port":               587,
				"from_email_address": "noreply@example.com",
				"tls_authentication": true,
				"send_timeout":       120,
				"smtp_authentication": []interface{}{
					map[string]interface{}{"user": "smtp_user", "password": "smtp_pass"},
				},
			},
		},
		"dns_config": []interface{}{
			map[string]interface{}{
				"servers": []interface{}{"8.8.8.8", "8.8.4.4"},
				"domains": []interface{}{"example.com"},
			},
		},
		"web_proxy_config": []interface{}{
			map[string]interface{}{
				"host":     "proxy.internal.com",
				"port":     3128,
				"username": "proxy_admin",
				"password": "proxy_secret",
			},
		},
		"sso_config": []interface{}{
			map[string]interface{}{
				"enabled":       true,
				"saml_metadata": "<EntityDescriptor/>",
			},
		},
	}
	if engineType == "CC" {
		raw["compliance_user"] = "admin"
		raw["compliance_password"] = "Admin-12"
		raw["compliance_new_password"] = "Admin@45"
		raw["compliance_email"] = "compliance@example.com"
	}
	return schema.TestResourceDataRaw(t, res.Schema, raw)
}

// TestEngineConfigCreateFullFlow drives the entire create handler end-to-end
// for a Continuous Compliance + OBJECT(AWS) engine with every optional service
// block enabled: CC setup, all configuration tasks, OBJECT storage parameter
// extraction, system initialization, action polling, engine-type selection,
// SSO, both password updates, and the final read. The mock engine returns
// success for every endpoint, so the handler runs to completion.
//
// This is the slow test (~80s: the 10s masking sleep, the 10s pre-init sleep,
// and the 60s post-init sleep all run for real); see the package note above.
func TestEngineConfigCreateFullFlow(t *testing.T) {
	t.Parallel()
	srv := newMockEngine(mockEngineOptions{})
	defer srv.Close()

	d := newCreateData(t, srv.URL, "CC")
	diags := engineConfigCreate(context.Background(), d, nil)
	if diags.HasError() {
		t.Fatalf("expected engineConfigCreate to succeed, got: %+v", diags)
	}
	if d.Id() != srv.URL {
		t.Errorf("expected resource ID %q, got %q", srv.URL, d.Id())
	}
	// engineConfigRead runs at the end of create and populates computed fields.
	if got := d.Get("hostname").(string); got == "" {
		t.Error("expected hostname to be populated by the trailing read")
	}
}

// TestEngineConfigCreateInitFailureRollback drives the Continuous Compliance
// rollback path: when system initialization fails, the handler restores the
// compliance user's password before returning the initialization error.
func TestEngineConfigCreateInitFailureRollback(t *testing.T) {
	t.Parallel()
	srv := newMockEngine(mockEngineOptions{initFails: true})
	defer srv.Close()

	d := newCreateData(t, srv.URL, "CC")
	diags := engineConfigCreate(context.Background(), d, nil)
	if !diags.HasError() {
		t.Fatal("expected engineConfigCreate to fail when initialization errors")
	}
	if !strings.Contains(diags[0].Summary, "Error initializing system") {
		t.Fatalf("unexpected error: %s", diags[0].Summary)
	}
}

// TestEngineConfigCreateComplianceSetup covers the Continuous Compliance branch
// (start masking, compliance-user login, compliance-user update). The mock
// rejects the compliance-user update so the handler returns right after the CC
// setup, before reaching the pre-init sleep. Only the 10s masking-startup sleep
// is incurred.
func TestEngineConfigCreateComplianceSetup(t *testing.T) {
	t.Parallel()
	srv := newMockEngine(mockEngineOptions{complianceUpdateFails: true})
	defer srv.Close()

	d := newCreateData(t, srv.URL, "CC")
	diags := engineConfigCreate(context.Background(), d, nil)
	if !diags.HasError() {
		t.Fatal("expected engineConfigCreate to fail at the compliance-user update step")
	}
	if !strings.Contains(diags[0].Summary, "Error in updating compliance user details") {
		t.Fatalf("unexpected error: %s", diags[0].Summary)
	}
}

// TestEngineConfigCreateSessionFailure covers the earliest error path: a failed
// session start returns immediately, before any sleeps.
func TestEngineConfigCreateSessionFailure(t *testing.T) {
	t.Parallel()
	d := newCreateData(t, "http://127.0.0.1:1", "CD") // nothing listening
	diags := engineConfigCreate(context.Background(), d, nil)
	if !diags.HasError() {
		t.Fatal("expected engineConfigCreate to fail when the session cannot start")
	}
	if !strings.Contains(diags[0].Summary, "Error starting session") {
		t.Fatalf("unexpected error: %s", diags[0].Summary)
	}
}
