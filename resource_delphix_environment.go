package main

import (
	"encoding/json"
	"fmt"
	"log"

	delphix "github.com/ajaytho/delphix-go-sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Environment struct {
	name         string
	description  string
	userName     string
	userPassword string
	address      string
	toolkitPath  string
	serverID     string
}

func resourceDelphixEnvironment() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        resourceDelphixEnvironmentCreate,
		Read:          resourceDelphixEnvironmentRead,
		Update:        resourceDelphixEnvironmentUpdate,
		Delete:        resourceDelphixEnvironmentDelete,
		Exists:        resourceDelphixEnvironmentExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{ // List of supported configuration fields for your resource
			"user_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"user_password": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"toolkit_path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"server_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}
func resourceDelphixEnvironmentExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	log.Println("Running Exists")
	client := meta.(*delphix.Client)
	reference := d.Id()
	present, err := client.FindEnvironmentByReference(reference)
	if err != nil || present == nil {
		return false, err
	}
	return true, nil
}
func resourceDelphixEnvironmentCreate(d *schema.ResourceData, meta interface{}) error {
	env := Environment{
		name:         d.Get("name").(string),
		description:  d.Get("description").(string),
		userName:     d.Get("user_name").(string),
		userPassword: d.Get("user_password").(string),
		address:      d.Get("address").(string),
		toolkitPath:  d.Get("toolkit_path").(string),
	}
	var reference interface{}
	client := meta.(*delphix.Client)
	environmentCreateParams := delphix.HostEnvironmentCreateParametersStruct{
		Type: "HostEnvironmentCreateParameters",
		PrimaryUser: &delphix.EnvironmentUserStruct{
			Type: "EnvironmentUser",
			Name: env.userName,
			Credential: &delphix.PasswordCredentialStruct{
				Type:     "PasswordCredential",
				Password: env.userPassword,
			},
		},
		HostEnvironment: &delphix.UnixHostEnvironmentStruct{
			Type:        "UnixHostEnvironment",
			Name:        env.name,
			Description: env.description,
		},
		HostParameters: &delphix.UnixHostCreateParametersStruct{
			Type: "UnixHostCreateParameters",
			Host: &delphix.UnixHostStruct{
				Type:        "UnixHost",
				Address:     env.address,
				ToolkitPath: env.toolkitPath,
			},
		},
	}

	fmt.Println(json.Marshal(environmentCreateParams))

	bits, err := json.Marshal(environmentCreateParams)
	fmt.Println(string(bits))
	environmentRef, err := client.FindEnvironmentByName(env.name)
	if err != nil {
		return err
	} else if environmentRef != nil {
		return fmt.Errorf("Environment \"%s\" already exists", env.name)
	}
	reference, err = client.CreateEnvironment(&environmentCreateParams)
	if err != nil {
		return err
	} else if reference == nil {
		return fmt.Errorf("Environment \"%s\" was not created", env.name)
	}
	d.SetId(reference.(string))
	return nil
}
func resourceDelphixEnvironmentRead(d *schema.ResourceData, meta interface{}) error {
	log.Println("Running Read")
	client := meta.(*delphix.Client)
	reference := d.Id()
	envObj, err := client.FindEnvironmentByReference(reference)
	if err != nil {
		return err
	} else if envObj == nil {
		return fmt.Errorf("Unable find environment \"%s\"", reference)
	}
	d.Set("name", envObj.(map[string]interface{})["name"].(string))
	d.Set("description", envObj.(map[string]interface{})["description"].(string))
	return nil
}
func resourceDelphixEnvironmentUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*delphix.Client)
	uEnv := Environment{
		name:         d.Get("name").(string),
		description:  d.Get("description").(string),
		userName:     d.Get("user_name").(string),
		userPassword: d.Get("user_password").(string),
		address:      d.Get("address").(string),
		toolkitPath:  d.Get("toolkit_path").(string),
	}
	environmentUpdateParams := delphix.UnixHostEnvironmentStruct{
		Type:        "UnixHostEnvironment",
		Name:        uEnv.name,
		Description: uEnv.description,
	}
	bits, _ := json.Marshal(environmentUpdateParams)
	fmt.Println(string(bits))
	if err := client.UpdateEnvironment(d.Id(), &environmentUpdateParams); err != nil {
		return fmt.Errorf("error updating Environment: %s", err.Error())
	}
	return resourceDelphixEnvironmentRead(d, meta)
}
func resourceDelphixEnvironmentDelete(d *schema.ResourceData, meta interface{}) error {
	log.Println("Running Delete")
	client := meta.(*delphix.Client)
	reference := d.Id()
	envObj, err := client.FindEnvironmentByReference(reference)
	if err != nil {
		return err
	} else if envObj == nil {
		return fmt.Errorf("Unable find environment \"%s\"", reference)
	}
	err = client.DeleteEnvironment(reference)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
