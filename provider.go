package main

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns a terraform.ResourceProvider.
func Provider() *schema.Provider {
	return &schema.Provider{ // Source https://github.com/hashicorp/terraform/blob/master/helper/schema
		Schema: providerSchema(),
		ResourcesMap: map[string]*schema.Resource{
			"delphix_group": resourceDelphixGroup(),
			"delphix_vdb": resourceDelphixOracleSIVDB(),
			"delphix_environment": resourceDelphixEnvironment(),
			"delphix_data_source_oracle": resourceDelphixOracleDSource(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"url": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Delphix URL",
			DefaultFunc: schema.EnvDefaultFunc("DELPHIX_URL", nil),
		},
		"delphix_admin_username": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The username for accessing Delphix as the delphix_admin user.",
			DefaultFunc: schema.EnvDefaultFunc("DELPHIX_ADMIN_USERNAME", nil),
		},
		"delphix_admin_password": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The password for accessing Delphix as the delphix_admin user.",
			DefaultFunc: schema.EnvDefaultFunc("DELPHIX_ADMIN_PASSWORD", nil),
			Sensitive:   true,
		},
		"sysadmin_username": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The username for accessing Delphix as the sysadmin user.",
			DefaultFunc: schema.EnvDefaultFunc("DELPHIX_SYSADMIN_USERNAME", nil),
		},
		"sysadmin_password": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The password for accessing Delphix as the sysadmin user.",
			DefaultFunc: schema.EnvDefaultFunc("DELPHIX_SYSADMIN_PASSWORD", nil),
			Sensitive:   true,
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		url:      d.Get("url").(string) + "/resources/json/delphix",
		username: d.Get("delphix_admin_username").(string),
		password: d.Get("delphix_admin_password").(string),
	}

	log.Println("[INFO] Initializing Delphix client")

	client := *config.Client()

	if err := client.LoadAndValidate(); err != nil {
		return nil, err
	}

	return &client, nil
}
