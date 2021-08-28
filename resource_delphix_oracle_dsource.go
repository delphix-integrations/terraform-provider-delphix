package main

import (
	"fmt"
	"log"

	delphix "github.com/delphix/delphix-go-sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type DSource struct {
	name            string
	description     string
	userName        string
	password        string
	groupName       string
	environment     string
	environmentUser string
	instance        string
	oracleHome      string
	linkNow         bool
}

func resourceDelphixOracleDSource() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        resourceDelphixOracleDSourceCreate,
		Read:          resourceDelphixOracleDSourceRead,
		Update:        resourceDelphixOracleDSourceUpdate,
		Delete:        resourceDelphixOracleDSourceDelete,
		Exists:        resourceDelphixOracleDSourceExists,
		Schema: map[string]*schema.Schema{ // List of supported configuration fields for your resource
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"environment": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"environment_user": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"oracle_home": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"link_now": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourceDelphixOracleDSourceExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	log.Println("Running Exists")
	client := meta.(*delphix.Client)
	reference := d.Id()
	present, err := client.FindDatabaseByReference(reference)
	if err != nil || present == nil {
		return false, err
	}
	return true, nil
}

func resourceDelphixOracleDSourceCreate(d *schema.ResourceData, meta interface{}) error {
	var reference interface{}
	boolPointer := new(bool)
	*boolPointer = false

	client := meta.(*delphix.Client)
	dSource := DSource{
		name:            d.Get("name").(string),
		description:     d.Get("description").(string),
		userName:        d.Get("user_name").(string),
		password:        d.Get("password").(string),
		environment:     d.Get("environment").(string),
		environmentUser: d.Get("environment_user").(string),
		groupName:       d.Get("group_name").(string),
		instance:        d.Get("instance").(string),
		oracleHome:      d.Get("oracle_home").(string),
		linkNow:         d.Get("link_now").(bool),
	}
	dSourcesExists, err := client.FindDatabaseByName(dSource.name)
	if err != nil {
		return err
	} else if dSourcesExists != nil {
		return fmt.Errorf("%s already exists. Exiting", dSource.name)
	}

	environmentExists, err := client.FindEnvironmentByReference(dSource.environment)
	if err != nil {
		return err
	} else if environmentExists == nil {
		return fmt.Errorf("Environment %s does not exist. Exiting", dSource.name)
	}

	groupObj, err := client.FindGroupRefByName(dSource.groupName)
	if err != nil {
		return err
	} else if groupObj == nil {
		return fmt.Errorf("Group \"%s\" not found", dSource.groupName)
	}

	groupRef := groupObj.(string)

	repoRef, err := client.FindRepoReferenceByEnvironmentRefAndOracleHome(dSource.environment, dSource.oracleHome)
	if err != nil {
		return err
	} else if repoRef == nil {
		return fmt.Errorf("Repo \"%s\" not found on \"%s\"", dSource.oracleHome, dSource.name)
	}

	sourceConfigObj, err := client.FindSourceConfigReferenceByNameAndRepoReference(dSource.instance, repoRef.(string))
	if err != nil {
		return err
	} else if sourceConfigObj == nil {
		return fmt.Errorf("Oracle Instance \"%s\" not found", dSource.instance)
	}

	userRef, err := client.FindEnvironmentUserRefByNameAndEnvironmentReference(dSource.environmentUser, dSource.environment)
	if err != nil {
		return err
	} else if userRef == nil {
		return fmt.Errorf("User \"%s\" not found", dSource.environmentUser)
	}
	if dSource.linkNow == true {
		*boolPointer = true
	}
	l := &delphix.LinkParametersStruct{
		Type:        "LinkParameters",
		Name:        dSource.name,
		Description: dSource.description,
		Group:       groupRef,
		LinkData: &delphix.OracleLinkFromExternalStruct{
			Type:            "OracleLinkFromExternal",
			Config:          sourceConfigObj.(string),
			EnvironmentUser: userRef.(string),
			OracleFallbackUser:          dSource.userName,
			OracleFallbackCredentials: &delphix.PasswordCredentialStruct{
				Type:     "PasswordCredential",
				Password: dSource.password,
			},
			SyncParameters: &delphix.OracleSyncFromExternalParametersStruct{
				SkipSpaceCheck: boolPointer,
				Type: "OracleSyncFromExternalParameters",
			},
			LinkNow: boolPointer,
		},
	}
	log.Println("Value of l")
	log.Println(l)
	reference, err = client.CreateDSource(l)
	if err != nil {
		return err
	} else if reference == nil {
		return fmt.Errorf("Data Source \"%s\" was not created found", dSource.name)
	}

	d.SetId(reference.(string))
	if err != nil {
		return err
	}

	return nil
}

func resourceDelphixOracleDSourceRead(d *schema.ResourceData, meta interface{}) error {
	log.Println("Running Read")
	client := meta.(*delphix.Client)
	reference := d.Id()
	dSourceObj, err := client.FindDatabaseByReference(reference)
	if err != nil {
		return err
	} else if dSourceObj == nil {
		return fmt.Errorf("Unable find database \"%s\"", reference)
	}
	d.Set("name", dSourceObj.(map[string]interface{})["name"])

	return nil
}

func resourceDelphixOracleDSourceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*delphix.Client)

	uDSource := DSource{
		name:            d.Get("name").(string),
		description:     d.Get("description").(string),
		userName:        d.Get("user_name").(string),
		password:        d.Get("password").(string),
		environment:     d.Get("environment").(string),
		environmentUser: d.Get("environment_user").(string),
		groupName:       d.Get("group_name").(string),
		instance:        d.Get("instance").(string),
		oracleHome:      d.Get("oracle_home").(string),
		linkNow:         d.Get("link_now").(bool),
	}

	oracleDatabaseContainer := delphix.OracleDatabaseContainerStruct{
		Type:        "OracleDatabaseContainer",
		Name:        uDSource.name,
		Description: uDSource.description,
	}
	if err := client.UpdateDatabase(d.Id(), &oracleDatabaseContainer); err != nil {
		return fmt.Errorf("error updating dSource: %s", err.Error())
	}

	return resourceDelphixOracleDSourceRead(d, meta)
}

func resourceDelphixOracleDSourceDelete(d *schema.ResourceData, meta interface{}) error {
	log.Println("Running Delete")
	client := meta.(*delphix.Client)
	databaseID := d.Id()
	dSourceObj, err := client.FindDatabaseByReference(databaseID)
	if err != nil {
		return err
	} else if dSourceObj == nil {
		return fmt.Errorf("Unable find database \"%s\"", databaseID)
	}
	reference := dSourceObj.(map[string]interface{})["reference"].(string)
	err = client.DeleteDatabase(reference)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
