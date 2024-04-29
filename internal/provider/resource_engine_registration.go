package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-log/tflog"

	dctapi "github.com/delphix/dct-sdk-go/v14"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEngineRegistration() *schema.Resource {
	return &schema.Resource{
		// Description is used by the doc genertor and language server.
		Description: "Provider Resource to add an environment to Delphix.",

		CreateContext: resourceEngineRegistrationCreate,
		ReadContext:   resourceEngineRegistrationRead,
		UpdateContext: resourceEngineRegistrationUpdate,
		DeleteContext: resourceEngineRegistrationDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
			},
			"masking_username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"masking_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hashicorp_vault_username_command_args": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"hashicorp_vault_masking_username_command_args": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"hashicorp_vault_password_command_args": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"hashicorp_vault_masking_password_command_args": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"hashicorp_vault_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"masking_hashicorp_vault_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"insecure_ssl": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"unsafe_ssl_hostname_check": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"truststore_filename": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"truststore_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpu_core_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"memory_size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_storage_capacity": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_storage_used": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"connection_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"engine_connection_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"connection_status_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"engine_connection_status_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceEngineRegistrationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Function to add an environment in an engine.

	var diags diag.Diagnostics
	client := meta.(*apiClient).client

	registerEngine := dctapi.NewEngineRegistrationParameter(d.Get("name").(string), d.Get("hostname").(string))

	if v, has_v := d.GetOk("username"); has_v {
		registerEngine.SetUsername(v.(string))
	}
	if v, has_v := d.GetOk("password"); has_v {
		registerEngine.SetPassword(v.(string))
	}
	if v, has_v := d.GetOk("masking_username"); has_v {
		registerEngine.SetMaskingUsername(v.(string))
	}
	if v, has_v := d.GetOk("masking_password"); has_v {
		registerEngine.SetMaskingPassword(v.(string))
	}
	if v, has_v := d.GetOk("hashicorp_vault_username_command_args"); has_v {
		registerEngine.SetHashicorpVaultUsernameCommandArgs(toStringArray(v.(string)))
	}
	if v, has_v := d.GetOk("hashicorp_vault_masking_username_command_args"); has_v {
		registerEngine.SetHashicorpVaultMaskingUsernameCommandArgs(toStringArray(v.(string)))
	}
	if v, has_v := d.GetOk("hashicorp_vault_password_command_args"); has_v {
		registerEngine.SetHashicorpVaultPasswordCommandArgs(toStringArray(v.(string)))
	}
	if v, has_v := d.GetOk("hashicorp_vault_masking_password_command_args"); has_v {
		registerEngine.SetHashicorpVaultMaskingPasswordCommandArgs(toStringArray(v.(string)))
	}
	if v, has_v := d.GetOk("hashicorp_vault_id"); has_v {
		registerEngine.SetHashicorpVaultId(v.(int64))
	}
	if v, has_v := d.GetOk("masking_hashicorp_vault_id"); has_v {
		registerEngine.SetMaskingHashicorpVaultId(v.(int64))
	}
	if v, has_v := d.GetOk("insecure_ssl"); has_v {
		registerEngine.SetInsecureSsl(v.(bool))
	}
	if v, has_v := d.GetOk("unsafe_ssl_hostname_check"); has_v {
		registerEngine.SetUnsafeSslHostnameCheck(v.(bool))
	}
	if v, has_v := d.GetOk("truststore_filename"); has_v {
		registerEngine.SetTruststoreFilename(v.(string))
	}
	if v, has_v := d.GetOk("truststore_password"); has_v {
		registerEngine.SetTruststorePassword(v.(string))
	}
	if v, has_v := d.GetOk("tags"); has_v {
		registerEngine.SetTags(toTagArray(v))
	}

	apiReq := client.ManagementApi.RegisterEngine(ctx)
	apiRes, httpRes, err := apiReq.EngineRegistrationParameter(*registerEngine).Execute()

	if diags := apiErrorResponseHelper(ctx, apiRes, httpRes, err); diags != nil {
		return diags
	}
	d.SetId(apiRes.GetId())

	readDiags := resourceEngineRegistrationRead(ctx, d, meta)
	if readDiags.HasError() {
		return readDiags
	}
	return diags
}

func resourceEngineRegistrationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient).client
	engineID := d.Id()

	// getEngineResp := client.ManagementApi.GetRegisteredEngine(ctx, engineID)
	// aapiRes, httpRes, err := getEngineResp.Execute()

	res, diags := PollForObjectExistence(ctx, func() (interface{}, *http.Response, error) {
		return client.ManagementApi.GetRegisteredEngine(ctx, engineID).Execute()
	})

	if diags != nil {
		_, diags := PollForObjectDeletion(ctx, func() (interface{}, *http.Response, error) {
			return client.ManagementApi.GetRegisteredEngine(ctx, engineID).Execute()
		})
		// This would imply error in poll for deletion so we just log and exit.
		if diags != nil {
			tflog.Error(ctx, DLPX+ERROR+"Error in polling of dSource for deletion.")
		} else {
			// diags will be nil in case of successful poll for deletion logic aka 404
			tflog.Error(ctx, DLPX+ERROR+"Error reading the engine "+engineID+", removing from state.")
			d.SetId("")
		}

		return nil
	}

	result, ok := res.(*dctapi.RegisteredEngine)
	if !ok {
		return diag.Errorf("Error occured in type casting.")
	}

	d.Set("id", result.GetId())
	d.Set("uuid", result.GetUuid())
	d.Set("status", result.GetStatus())
	d.Set("connection_status", result.GetConnectionStatus())

	return diags
}

func resourceEngineRegistrationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Info(ctx, DLPX+INFO+"Not Implemented: resourceEngineRegistrationUpdate")
	var diags diag.Diagnostics
	return diags
}

func resourceEngineRegistrationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*apiClient).client
	engineID := d.Id()

	apiRes, httpRes, err := client.ManagementApi.UnregisterEngine(ctx, engineID).Execute()

	if diags := apiErrorResponseHelper(ctx, apiRes, httpRes, err); diags != nil {
		return diags
	}

	// job_status, job_err := PollJobStatus(*apiRes.Job.Id, ctx, client)
	// if job_err != "" {
	// 	tflog.Error(ctx, DLPX+ERROR+"Job Polling failed but continuing with engine removal. Error: "+job_err)
	// }
	// if isJobTerminalFailure(job_status) {
	// 	return diag.Errorf("[NOT OK] Engine-Delete %s. JobId: %s / Error: %s", job_status, *apiRes.Job.Id, job_err)
	// }
	// _, diags := PollForObjectDeletion(ctx, func() (interface{}, *http.Response, error) {
	// 	return client.ManagementApi.GetRegisteredEngine(ctx, engineID).Execute()
	// })

	_, diags := PollForObjectExistence(ctx, func() (interface{}, *http.Response, error) {
		return client.ManagementApi.GetRegisteredEngine(ctx, engineID).Execute()
	})

	if diags != nil {
		tflog.Error(ctx, DLPX+ERROR+"Error in polling registered engine.")
		return diags
	}

	return diags
}
