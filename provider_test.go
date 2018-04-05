package main

import (
	"os"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var (
	//Acceptance Test Vars
	testAccProviders          map[string]terraform.ResourceProvider
	testAccProvider           *schema.Provider
	testAccDelphixAdminConfig Config
	testAccVDB                VDB
	testAccEnvironment        Environment
	testAccDsource            DSource
	//This Client has a slightly different URL format for the non-ACC tests
	testDelphixAdminConfig Config
	testSysadminConfig     Config
)

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"delphix": testAccProvider,
	}
	testAccDelphixAdminConfig.url = os.Getenv("TEST_DELPHIX_URL")
	testAccDelphixAdminConfig.username = os.Getenv("TEST_DELPHIX_USERNAME")
	testAccDelphixAdminConfig.password = os.Getenv("TEST_DELPHIX_PASSWORD")
	//VDB
	testAccVDB.dbName = os.Getenv("TEST_DBNAME")
	testAccVDB.name = os.Getenv("TEST_VDBNAME")
	testAccVDB.groupName = os.Getenv("TEST_GROUP_NAME")
	testAccVDB.oracleHome = os.Getenv("TEST_ORACLE_HOME")
	//dSources
	testAccDsource.name = os.Getenv("TEST_DELPHIX_DSOURCE_NAME")
	testAccDsource.description = os.Getenv("TEST_DELPHIX_DSOURCE_DESCRIPTION")
	testAccDsource.userName = os.Getenv("TEST_DELPHIX_DSOURCE_USERNAME")
	testAccDsource.password = os.Getenv("TEST_DELPHIX_DSOURCE_PASSWORD")
	testAccDsource.groupName = os.Getenv("TEST_DSOURCE_GROUP_NAME")
	testAccDsource.environmentUser = os.Getenv("TEST_DSOURCE_ENVIRONMENT_USER")
	testAccDsource.linkNow, _ = strconv.ParseBool(os.Getenv("TEST_DSOURCE_ENVIRONMENT_LINKNOW"))
	testAccDsource.instance = os.Getenv("TEST_DSOURCE_ENVIRONMENT_INSTANCE")
	testAccDsource.oracleHome = os.Getenv("TEST_DSOURCE_ENVIRONMENT_OH")
	//This url needs the /resources/json/delphix append
	testDelphixAdminConfig.url = testAccDelphixAdminConfig.url + "/resources/json/delphix"
	testDelphixAdminConfig.username = testAccDelphixAdminConfig.username
	testDelphixAdminConfig.password = testAccDelphixAdminConfig.password
	testSysadminConfig.url = testDelphixAdminConfig.url
	testSysadminConfig.username = "sysadmin"
	testSysadminConfig.password = os.Getenv("TEST_SYSADMIN_PASSWORD")
	//Environment
	testAccEnvironment.name = os.Getenv("TEST_ENV_NAME")
	testAccEnvironment.description = os.Getenv("TEST_ENV_DESC")
	testAccEnvironment.address = os.Getenv("TEST_ENV_ADD")
	testAccEnvironment.userName = os.Getenv("TEST_ENV_USER")
	testAccEnvironment.userPassword = os.Getenv("TEST_ENV_PASS")
	testAccEnvironment.toolkitPath = os.Getenv("TEST_ENV_TK")

}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	// We will use this function later on to make sure our test environment is valid.
	// For example, you can make sure here that some environment variables are set.
}
