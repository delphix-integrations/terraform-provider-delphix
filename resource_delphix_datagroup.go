package main

import (
	"encoding/json"
	"fmt"
	"log"

	delphix "github.com/delphix/delphix-go-sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Group struct {
	name         string
	description  string
}

func resourceDelphixGroup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        resourceDelphixGroupCreate,
		Read:          resourceDelphixGroupRead,
		Update:        resourceDelphixGroupUpdate,
		Delete:        resourceDelphixGroupDelete,
		Exists:        resourceDelphixGroupExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{ // List of supported configuration fields for your resource
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func resourceDelphixGroupExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	log.Println("Running Exists")
	client := meta.(*delphix.Client)
	reference := d.Id()
	present, err := client.FindGroupByRef(reference)
	if err != nil || present == nil {
		return false, err
	}
	return true, nil
}
func resourceDelphixGroupCreate(d *schema.ResourceData, meta interface{}) error {
	grp := Group{
		name:         d.Get("name").(string),
		description:  d.Get("description").(string),
	}
	var reference interface{}
	client := meta.(*delphix.Client)
	groupCreateParams := delphix.GroupStruct{
		Type: "Group",
		Name: grp.name,
		Description: grp.description,
	}
	fmt.Println(json.Marshal(groupCreateParams))

	bits, err := json.Marshal(groupCreateParams)
	fmt.Println(string(bits))
	groupRef, err := client.FindGroupByName(grp.name)
	if err != nil {
		log.Println("Error is not nil ------------")
		return err
	} else if groupRef != nil {
		log.Println("GroupRef is not nil ------------")
		return fmt.Errorf("Group \"%s\" already exists", grp.name)
	}
	reference, err = client.CreateGroup(&groupCreateParams)
	if err != nil {
		return err
	} else if reference == nil {
		return fmt.Errorf("Group \"%s\" was not created", grp.name)
	}
	d.SetId(reference.(string))
	return nil
}
func resourceDelphixGroupRead(d *schema.ResourceData, meta interface{}) error {
	log.Println("Running Read")
	client := meta.(*delphix.Client)
	reference := d.Id()
	grpObj, err := client.FindGroupByRef(reference)
	if err != nil {
		return err
	} else if grpObj == nil {
		return fmt.Errorf("Unable find group \"%s\"", reference)
	}
	d.Set("name", grpObj.(map[string]interface{})["name"].(string))
	d.Set("description", grpObj.(map[string]interface{})["description"].(string))
	return nil
}
func resourceDelphixGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*delphix.Client)
	uGrp := Group{
		name:         d.Get("name").(string),
		description:  d.Get("description").(string),
	}
	groupUpdateParams := delphix.GroupStruct{
		Type:        "Group",
		Name:        uGrp.name,
		Description: uGrp.description,
	}
	bits, _ := json.Marshal(groupUpdateParams)
	fmt.Println(string(bits))
	if err := client.UpdateGroup(d.Id(), &groupUpdateParams); err != nil {
		return fmt.Errorf("error updating Group: %s", err.Error())
	}
	return resourceDelphixGroupRead(d, meta)
}
func resourceDelphixGroupDelete(d *schema.ResourceData, meta interface{}) error {
	log.Println("Running Delete")
	client := meta.(*delphix.Client)
	reference := d.Id()
	grpObj, err := client.FindGroupByRef(reference)
	if err != nil {
		return err
	} else if grpObj == nil {
		return fmt.Errorf("unable find group \"%s\"", reference)
	}
	err = client.DeleteGroup(reference)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
