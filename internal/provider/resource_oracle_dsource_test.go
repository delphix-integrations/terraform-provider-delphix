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

func TestOracleDsource_create_positive(t *testing.T) {
	sourcevalue := os.Getenv("ORACLE_DSOURCE_SOURCE_VALUE")
	groupId := os.Getenv("ORACLE_DSOURCE_GROUP_ID")
	name := os.Getenv("ORACLE_DSOURCE_NAME")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testOracleDsourcePreCheck(t, sourcevalue, groupId, name)
		},
		Providers:    testAccProviders,
		CheckDestroy: testOracleDsourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testOracleDsourceBasic(name, sourcevalue, groupId),
				Check: resource.ComposeTestCheckFunc(
					testOracleDsourceExists("delphix_oracle_dsource.test_oracle_dsource", sourcevalue),
					resource.TestCheckResourceAttr("delphix_oracle_dsource.test_oracle_dsource", "source_id", sourcevalue)),
			},
			{
				// positive update test case
				Config: testOracleDsourceBasic("update_name", sourcevalue, groupId), // changing the name to update-name
				Check: resource.ComposeTestCheckFunc(
					testOracleDsourceExists("delphix_oracle_dsource.test_oracle_dsource", sourcevalue),
					resource.TestCheckResourceAttr("delphix_oracle_dsource.test_oracle_dsource", "name", "update_name"), // asserting the updated name
					resource.TestCheckResourceAttr("delphix_oracle_dsource.test_oracle_dsource", "group_id", groupId)),
			},
			{
				// negative update test case
				Config:      testOracleDsourceUpdateNegative(name, sourcevalue, "non-existent"),
				ExpectError: regexp.MustCompile("Error running apply: exit status 1"),
			},
			{
				Config:      testOracleDsourceBasic(name, "", groupId),
				ExpectError: regexp.MustCompile(`.*`),
			},
		},
	})
}

func testOracleDsourcePreCheck(t *testing.T, sourceId string, groupId string, name string) {
	testAccPreCheck(t)
	if sourceId == "" {
		t.Fatal("ORACLE_DSOURCE_SOURCE_VALUE must be set for env acceptance tests")
	}
	if groupId == "" {
		t.Fatal("ORACLE_DSOURCE_GROUP_ID must be set for env acceptance tests")
	}
	if name == "" {
		t.Fatal("ORACLE_DSOURCE_NAME must be set for env acceptance tests")
	}
}

func testOracleDsourceBasic(name string, sourceValue string, groupId string) string {
	return fmt.Sprintf(`
resource "delphix_oracle_dsource" "test_oracle_dsource" {
  name                       = "%s"
  source_value               = "%s"
  group_id                   = "%s"
  tags {
	key = "dlpx"
	value = "acc-test"
    }
  ops_pre_sync {
    name            = "string-change-opspresync22"
    command         = "ls -lr"
    shell           = "bash"
    credentials_env_vars {
      base_var_name = "mypass2t"
      password = "password_test"
    }
    credentials_env_vars {
      base_var_name = "mypass3t"
      password = "password_test"
    }
  }
  
  ops_post_sync {
    name            = "string-change-opspostsync22"
    command         = "ls -lrta"
    shell           = "bash"
    credentials_env_vars {
      base_var_name = "mypassopspostsynct"
      password = "password_test"
    }
  }

  ops_pre_log_sync {
    name            = "string-change-opsprelogsync22"
    command         = "ls -lrt"
    shell           = "shell"
    credentials_env_vars {
      base_var_name = "mypassopsprelogsynct"
      password = "password_test"
    }
  }
}
  
	`, name, sourceValue, groupId)
}

func testOracleDsourceUpdateNegative(name string, sourceValue string, description string) string {
	return fmt.Sprintf(`
resource "delphix_oracle_dsource" "test_oracle_dsource" {
  name                       = "%s"
  source_value               = "%s"
  description                = "%s"
  tags {
	key = "dlpx"
	value = "acc-test"
	}
  ops_pre_sync {
    name            = "string-change-opspresync22"
    command         = "ls -lr"
    shell           = "bash"
    credentials_env_vars {
      base_var_name = "mypass2t"
      password = "password_test"
    }
    credentials_env_vars {
      base_var_name = "mypass3t"
      password = "password_test"
    }
  }
  
  ops_post_sync {
    name            = "string-change-opspostsync22"
    command         = "ls -lrta"
    shell           = "bash"
    credentials_env_vars {
      base_var_name = "mypassopspostsynct"
      password = "password_test"
    }
  }

  ops_pre_log_sync {
    name            = "string-change-opsprelogsync22"
    command         = "ls -lrt"
    shell           = "shell"
    credentials_env_vars {
      base_var_name = "mypassopsprelogsynct"
      password = "password_test"
    }
  }
}
	`, name, sourceValue, description)
}

func testOracleDsourceExists(n string, sourceValue string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		dsourceId := rs.Primary.ID
		if dsourceId == "" {
			return fmt.Errorf("No dsourceId set")
		}

		client := testAccProvider.Meta().(*apiClient).client
		res, _, err := client.DSourcesAPI.GetDsourceById(context.Background(), dsourceId).Execute()
		if err != nil {
			return err
		}

		resSourceId := res.GetSourceId()
		if resSourceId != sourceValue {
			return fmt.Errorf("SourceId mismatch")
		}

		return nil
	}
}

func testOracleDsourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*apiClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "delphix_oracle_dsource" {
			continue
		}

		dsourceId := rs.Primary.ID

		_, httpResp, _ := client.DSourcesAPI.GetDsourceById(context.Background(), dsourceId).Execute()
		if httpResp == nil {
			return fmt.Errorf("Dsource has not been deleted")
		}

		if httpResp.StatusCode != 404 {
			return fmt.Errorf("Exepcted a 404 Not Found for a deleted Dsource but got %d", httpResp.StatusCode)
		}
	}

	return nil
}
