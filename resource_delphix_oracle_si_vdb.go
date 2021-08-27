package main

import (
	"fmt"
	"log"

	delphix "github.com/ajaytho/delphix-go-sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// VDB provides basic VDB object parameters
type VDB struct {
	source      string
	name        string
	environment string
	groupName   string
	dbName      string
	oracleHome  string
	mountBase   string
	snapSource  bool
}

func resourceDelphixOracleSIVDB() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        resourceDelphixOracleSIVDBCreate,
		Read:          resourceDelphixOracleSIVDBRead,
		Update:        resourceDelphixOracleSIVDBUpdate,
		Delete:        resourceDelphixOracleSIVDBDelete,
		Exists:        resourceDelphixOracleSIVDBExists,
		Schema: map[string]*schema.Schema{ // List of supported configuration fields for your resource
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"db_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"environment": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"oracle_home": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"mount_base": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"snap_source": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourceDelphixOracleSIVDBExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	log.Println("Running Exists")
	client := meta.(*delphix.Client)
	reference := d.Id()
	present, err := client.FindDatabaseByReference(reference)
	if err != nil || present == nil {
		return false, err
	}
	return true, nil
}

func resourceDelphixOracleSIVDBCreate(d *schema.ResourceData, meta interface{}) error {
	var reference interface{}
	f := new(bool)
	*f = false
	var instnum = 1

	client := meta.(*delphix.Client)
	vdb := VDB{
		source:      d.Get("source").(string),
		name:        d.Get("name").(string),
		dbName:      d.Get("db_name").(string),
		environment: d.Get("environment").(string),
		groupName:   d.Get("group_name").(string),
		oracleHome:  d.Get("oracle_home").(string),
		mountBase:   d.Get("mount_base").(string),
		snapSource:  d.Get("snap_source").(bool),
	}

	vdbExists, err := client.FindDatabaseByName(vdb.name)
	if err != nil {
		return err
	} else if vdbExists != nil {
		return fmt.Errorf("%s already exists. Exiting", vdb.name)
	}

	environmentExists, err := client.FindEnvironmentByReference(vdb.environment)
	if err != nil {
		return err
	} else if environmentExists == nil {
		return fmt.Errorf("Environment %s does not exist. Exiting", vdb.environment)
	}

	sourceObj, err := client.FindDatabaseByReference(vdb.source)
	if err != nil {
		return err
	} else if sourceObj == nil {
		return fmt.Errorf("Source Database \"%s\" not found", vdb.source)
	}

	groupObj, err := client.FindGroupRefByName(vdb.groupName)
	if err != nil {
		return err
	} else if groupObj == nil {
		return fmt.Errorf("Group \"%s\" not found", vdb.groupName)
	}

	groupRef := groupObj.(string)

	repoRef, err := client.FindRepoReferenceByEnvironmentRefAndOracleHome(vdb.environment, vdb.oracleHome)
	if err != nil {
		return err
	} else if repoRef == nil {
		return fmt.Errorf("Repo \"%s\" not found on \"%s\"", vdb.oracleHome, vdb.environment)
	}

	log.Println("Milestone 1")
	oracleProvisionParameters := delphix.OracleProvisionParametersStruct{
		Type: "OracleProvisionParameters",
		Container: &delphix.OracleDatabaseContainerStruct{
			Type:  "OracleDatabaseContainer",
			Name:  vdb.name,
			Group: groupRef,
		},
		Source: &delphix.OracleVirtualSourceStruct{
			Type:                            "OracleVirtualSource",
			MountBase:                       vdb.mountBase,
			AllowAutoVDBRestartOnHostReboot: f,
		},
		SourceConfig: &delphix.OracleSIConfigStruct{
			Type:         "OracleSIConfig",
			Repository:   repoRef.(string),
			DatabaseName: vdb.dbName,
			UniqueName:   vdb.dbName,
			Instance: &delphix.OracleInstanceStruct{
				Type:           "OracleInstance",
				InstanceName:   vdb.dbName,
				InstanceNumber: &instnum,
			},
		},
		TimeflowPointParameters: delphix.TimeflowPointSemanticStruct{
			Type:      "TimeflowPointSemantic",
			Container: vdb.source,
			Location:  "LATEST_SNAPSHOT",
		},
	}

	log.Println("Milestone 2")
	log.Println(oracleProvisionParameters)

	if vdb.snapSource == true {
		if err = client.SyncDatabase(vdb.source); err != nil {
			return err
		}
	}

	reference, err = client.CreateDatabase(&oracleProvisionParameters)
	if err != nil {
		return err
	} else if reference == nil {
		return fmt.Errorf("Database \"%s\" was not created found", vdb.name)
	}

	d.SetId(reference.(string))
	if err != nil {
		return err
	}

	return nil
}

func resourceDelphixOracleSIVDBRead(d *schema.ResourceData, meta interface{}) error {
	log.Println("Running Read")
	client := meta.(*delphix.Client)
	reference := d.Id()
	vdbObj, err := client.FindDatabaseByReference(reference)
	if err != nil {
		return err
	} else if vdbObj == nil {
		return fmt.Errorf("Unable find database \"%s\"", reference)
	}
	d.Set("name", vdbObj.(map[string]interface{})["name"])

	return nil
}

func resourceDelphixOracleSIVDBUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*delphix.Client)

	uVDB := VDB{
		source:      d.Get("source").(string),
		name:        d.Get("name").(string),
		dbName:      d.Get("db_name").(string),
		environment: d.Get("environment").(string),
		groupName:   d.Get("group_name").(string),
		oracleHome:  d.Get("oracle_home").(string),
	}

	oracleDatabaseContainer := delphix.OracleDatabaseContainerStruct{
		Type: "OracleDatabaseContainer",
		Name: uVDB.name,
	}
	if err := client.UpdateDatabase(d.Id(), &oracleDatabaseContainer); err != nil {
		return fmt.Errorf("error updating VDB: %s", err.Error())
	}

	return resourceDelphixOracleSIVDBRead(d, meta)
}

func resourceDelphixOracleSIVDBDelete(d *schema.ResourceData, meta interface{}) error {
	log.Println("Running Delete")
	client := meta.(*delphix.Client)
	databaseID := d.Id()
	vdbObj, err := client.FindDatabaseByReference(databaseID)
	if err != nil {
		return err
	} else if vdbObj == nil {
		return fmt.Errorf("Unable find database \"%s\"", databaseID)
	}
	reference := vdbObj.(map[string]interface{})["reference"].(string)
	err = client.DeleteDatabase(reference)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
