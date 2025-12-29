package provider

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
					resource.TestCheckResourceAttr(resourceName, "api_version", "1.11.46"),
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
	}

	invalidSizes := []string{
		"100",
		"100MB",
		"100KB",
		"abc",
		"100 GB",
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
  api_version               = "1.11.46"
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
  engine_host  = "%s"
  api_version  = "1.11.46"
  sys_user     = "sysadmin"
  sys_password = "sysadmin"
  user         = "admin"
  password     = "delphix"
  email        = "test@example.com"
  engine_type  = "CD"
  device_type  = "BLOCK"
}
`, engineHost)
}

func testAccEngineConfigurationObjectStorageRole(engineHost, bucketName string) string {
	return fmt.Sprintf(`
resource "delphix_engine_configuration" "test" {
  engine_host  = "%s"
  api_version  = "1.11.46"
  sys_user     = "sysadmin"
  sys_password = "sysadmin"
  user         = "admin"
  password     = "delphix"
  email        = "test@example.com"
  engine_type  = "CD"
  device_type  = "OBJECT"
  
  ntp_servers  = ["pool.ntp.org", "time.nist.gov"]
  ntp_timezone = "America/New_York"
  
  object_storage_params {
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
  engine_host  = "%s"
  api_version  = "1.11.46"
  sys_user     = "sysadmin"
  sys_password = "sysadmin"
  user         = "admin"
  password     = "delphix"
  email        = "test@example.com"
  engine_type  = "CD"
  device_type  = "OBJECT"
  
  ntp_servers  = ["pool.ntp.org", "time.nist.gov", "1.ubuntu.pool.ntp.org"]
  ntp_timezone = "UTC"
  
  object_storage_params {
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
  engine_host  = "%s"
  api_version  = "1.11.46"
  sys_user     = "sysadmin"
  sys_password = "sysadmin"
  user         = "admin"
  password     = "delphix"
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
  engine_host  = "%s"
  api_version  = "1.11.46"
  sys_user     = "sysadmin"
  sys_password = "sysadmin"
  user         = "admin"
  password     = "delphix"
  email        = "test@example.com"
  engine_type  = "CD"
  device_type  = "OBJECT"
  
  ntp_servers  = ["pool.ntp.org"]
  ntp_timezone = "UTC"
  
  object_storage_params {
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
  engine_host  = "%s"
  api_version  = "1.11.46"
  sys_user     = "sysadmin"
  sys_password = "sysadmin"
  user         = "admin"
  password     = "delphix"
  email        = "test@example.com"
  engine_type  = "CD"
  device_type  = "OBJECT"
  
  # Missing ntp_servers and ntp_timezone
  
  object_storage_params {
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
  engine_host  = "%s"
  api_version  = "1.11.46"
  sys_user     = "sysadmin"
  sys_password = "sysadmin"
  user         = "admin"
  password     = "delphix"
  email        = "test@example.com"
  engine_type  = "CD"
  device_type  = "OBJECT"
  
  ntp_servers  = ["pool.ntp.org"]
  ntp_timezone = "UTC"
  
  object_storage_params {
    region    = "us-west-2"
    bucket    = "test-bucket"
    endpoint  = "s3.us-west-2.amazonaws.com"
    size      = "20MB"  # Invalid size unit
    auth_type = "ROLE"
  }
}
`, engineHost)
}
