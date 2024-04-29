package provider

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEngineUserManagement() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Resource for Engine User Management.",

		CreateContext: engineUserCreate,
		ReadContext:   engineUserRead,
		UpdateContext: engineUserUpdate,
		DeleteContext: engineUserDelete,

		Schema: map[string]*schema.Schema{
			"engine_host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"login_user": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"login_password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"user_name": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	} // maye be add enabled and other phone no feilds
}

func engineUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	// Create a cookie jar to store session cookies
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}

	engine_host, _ := d.Get("engine_host").(string)
	version, _ := d.Get("version").(string)
	login_user, _ := d.Get("login_user").(string)
	login_password, _ := d.Get("login_password").(string)
	user_name, _ := d.Get("user_name").(string)
	password, _ := d.Get("password").(string)
	user_type, _ := d.Get("user_type").(string)

	// Start a session
	tflog.Info(ctx, DLPX+INFO+"start Session for "+engine_host)
	err := startSession(ctx, client, engine_host, version)
	if err != nil {
		return diag.Errorf("Error starting session: %s", err)
	}

	// Authenticate/login
	tflog.Info(ctx, DLPX+INFO+"login as "+login_user)
	err = login(ctx, client, engine_host, login_user, login_password, user_type)
	if err != nil {
		return diag.Errorf("Error logging in: %s", err)
	}

	//Add check to see if it is already existing user

	action, err := createOrUpdateUser(ctx, client, engine_host, user_name, password, user_type)
	if err != nil {
		return diag.Errorf("Error logging in: %s", err)
	}

	var result ActionResult
	unmarshalErr := json.Unmarshal(action, &result)
	if unmarshalErr != nil {
		tflog.Error(ctx, DLPX+ERROR+"Error unmarshalling: "+unmarshalErr.Error())
	}

	tflog.Info(ctx, DLPX+INFO+"User Create Successfull!")

	d.SetId(engine_host)

	readDiags := engineUserRead(ctx, d, meta)
	if readDiags.HasError() {
		return readDiags
	}

	return diags
}

func engineUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// get the changed keys
	changedKeys := make([]string, 0, len(d.State().Attributes))
	for k := range d.State().Attributes {
		if d.HasChange(k) {
			changedKeys = append(changedKeys, k)
		}
	}
	// revert and set the old value to the changed keys
	for _, key := range changedKeys {
		old, _ := d.GetChange(key)
		d.Set(key, old)
	}

	return diag.Errorf("Action update not available for engine config : dSource")
}

func engineUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	engineId := d.Id()
	version, _ := d.Get("version").(string)

	// Create a cookie jar to store session cookies
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}

	// Start a session
	err := startSession(ctx, client, engineId, version)
	if err != nil {
		diag.Errorf("Error starting session: %v", err)
	}

	// Authenticate/login
	err = login(ctx, client, engineId, d.Get("sys_user").(string), d.Get("sys_new_password").(string), "SYSTEM")
	if err != nil {
		diag.Errorf("Error logging in: %v", err)
	}

	body, err := getSystem(ctx, client, engineId)
	if err != nil {
		diag.Errorf("Error getting system info: %v", err)
	}

	var response SystemInfoResponse
	sysErr := json.Unmarshal(body, &response)
	if sysErr != nil {
		tflog.Error(ctx, DLPX+ERROR+"Error unmarshalling", map[string]interface{}{"error": sysErr.Error()})
	}

	d.Set("configured", response.Result["configured"])
	d.Set("hostname", response.Result["hostname"])

	return diags
}

func engineUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// get the changed keys
	changedKeys := make([]string, 0, len(d.State().Attributes))
	for k := range d.State().Attributes {
		if d.HasChange(k) {
			changedKeys = append(changedKeys, k)
		}
	}
	// revert and set the old value to the changed keys
	for _, key := range changedKeys {
		old, _ := d.GetChange(key)
		d.Set(key, old)
	}

	return diag.Errorf("Action delete not available for engine config : dSource")
}
