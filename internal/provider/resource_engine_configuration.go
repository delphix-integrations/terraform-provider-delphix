package provider

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEngineConfiguration() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Resource for Engine initialization.",

		CreateContext: engineConfigCreate,
		ReadContext:   engineConfigRead,
		UpdateContext: engineConfigUpdate,
		DeleteContext: engineConfigDelete,

		Schema: map[string]*schema.Schema{
			"engine_host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sys_user": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"sys_password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			// "sys_new_password": {
			// 	Type:      schema.TypeString,
			// 	Required:  true,
			// 	Sensitive: true,
			// },
			"user": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"engine_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"configured": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"product_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sso_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"vendor_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"product_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"engine_qualifier": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"support_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"build_title": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"build_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"build_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled_features": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"toggleable_features": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// "apiVersion": {
			// 	Type:     schema.TypeString,
			// 	Computed: true,
			// },
			"banner": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"current_locale": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"kernel_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"platform": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func engineConfigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	// Create a cookie jar to store session cookies
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}

	var default_email string
	engine_host, _ := d.Get("engine_host").(string)
	version, _ := d.Get("api_version").(string)
	sys_user, _ := d.Get("sys_user").(string)
	sys_curr_pass, _ := d.Get("sys_password").(string)
	//sys_new_pass, _ := d.Get("sys_new_password").(string)
	user, _ := d.Get("user").(string)
	email, has_email := d.GetOk("email")
	if has_email {
		default_email = email.(string)
	}
	password := d.Get("password").(string)
	engine_type := d.Get("engine_type").(string)

	// // Update sys_user password
	// readDiags := UpdateUserPassword(ctx, client, engine_host, version, sys_user, sys_curr_pass, sys_new_pass, "SYSTEM")
	// if readDiags.HasError() {
	// 	return readDiags
	// }

	// Start a session
	tflog.Info(ctx, DLPX+INFO+"start Session for "+engine_host)
	err := startSession(ctx, client, engine_host, version)
	if err != nil {
		return diag.Errorf("Error starting session: %s", err)
	}

	// Authenticate/login
	tflog.Info(ctx, DLPX+INFO+"login as "+sys_user)
	err = login(ctx, client, engine_host, sys_user, sys_curr_pass, "SYSTEM")
	if err != nil {
		return diag.Errorf("Error logging in: %s", err)
	}

	// Initialize Engine
	result, readDiags := initializeSystemAndDevices(ctx, client, engine_host, user, default_email, password)
	if readDiags.HasError() {
		return readDiags
	}

	// Poll for initialization to complete
	readDiags = pollActionStatus(ctx, client, engine_host, result.Action)
	if readDiags.HasError() {
		return readDiags
	}

	// Sleep a minute for the initialization to complete
	time.Sleep(time.Duration(60) * time.Second)

	// Start a session
	tflog.Info(ctx, DLPX+INFO+"start Session for "+engine_host)
	err = startSession(ctx, client, engine_host, version)
	if err != nil {
		return diag.Errorf("Error starting session: %s", err)
	}

	// Authenticate/login
	tflog.Info(ctx, DLPX+INFO+"login as "+sys_user)
	err = login(ctx, client, engine_host, sys_user, sys_curr_pass, "SYSTEM")
	if err != nil {
		return diag.Errorf("Error logging in: %s", err)
	}

	resp, err := setEngieType(ctx, client, engine_host, engine_type)
	if err != nil {
		return diag.Errorf("Error setting engine type: %s", err)
	}

	tflog.Info(ctx, DLPX+INFO+"engine type resp "+string(resp))

	//Update defaultUser password
	readDiags = UpdateUserPassword(ctx, client, engine_host, version, user, password, password, "DOMAIN")
	if readDiags.HasError() {
		return readDiags
	}

	tflog.Info(ctx, DLPX+INFO+"System initialization successful!")

	d.SetId(engine_host)

	readDiags = engineConfigRead(ctx, d, meta)
	if readDiags.HasError() {
		return readDiags
	}

	return diags
}

func engineConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func engineConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	engineId := d.Id()
	version, _ := d.Get("api_version").(string)

	// Create a cookie jar to store session cookies
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}

	// Start a session
	err := startSession(ctx, client, engineId, version)
	if err != nil {
		diag.Errorf("Error starting session: %v", err)
	}

	// Authenticate/login
	err = login(ctx, client, engineId, d.Get("sys_user").(string), d.Get("sys_password").(string), "SYSTEM")
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
	d.Set("product_type", response.Result["productType"])
	//d.Set("engineType", response.Result["engineType"])
	d.Set("sso_enabled", response.Result["ssoEnabled"])
	d.Set("vendor_name", response.Result["vendorName"])
	d.Set("product_name", response.Result["productName"])
	d.Set("engine_qualifier", response.Result["engineQualifier"])
	d.Set("support_url", response.Result["supportURL"])
	d.Set("build_title", response.Result["buildTitle"])
	d.Set("build_timestamp", response.Result["buildTimestamp"])
	d.Set("build_version", response.Result["buildVersion"])
	d.Set("enabled_features", response.Result["enabledFeatures"])
	d.Set("toggleable_features", response.Result["toggleableFeatures"])
	//d.Set("apiVersion", response.Result["apiVersion"])
	d.Set("banner", response.Result["banner"])
	d.Set("locales", response.Result["locales"])
	d.Set("current_locale", response.Result["currentLocale"])
	d.Set("kernel_name", response.Result["kernelName"])
	d.Set("platform", response.Result["platform"])

	return diags
}

func engineConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
