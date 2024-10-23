package provider

import (
	"context"
	"encoding/json"
	"net/http"

	dctapi "github.com/delphix/dct-sdk-go/v22"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAppdataDsource() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Resource for appdata dSource creation.",

		CreateContext: resourceAppdataDsourceCreate,
		ReadContext:   resourceDsourceRead,
		UpdateContext: resourceDsourceUpdate,
		DeleteContext: resourceDsourceDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_value": {
				Type:     schema.TypeString,
				Required: true,
			},
			"group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_sync_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"make_current_account_owner": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"link_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"staging_mount_base": {
				Type:     schema.TypeString,
				Required: true,
			},
			"staging_environment": {
				Type:     schema.TypeString,
				Required: true,
			},
			"staging_environment_user": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"environment_user": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ops_pre_sync": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"command": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"shell": {
							Type:     schema.TypeString,
							Optional: true,
						}, "credentials_env_vars": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"base_var_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"password": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"vault": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"hashicorp_vault_engine": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"hashicorp_vault_secret_path": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"hashicorp_vault_username_key": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"hashicorp_vault_secret_key": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"azure_vault_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"azure_vault_username_key": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"azure_vault_secret_key": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"cyberark_vault_query_string": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"ops_post_sync": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"command": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"shell": {
							Type:     schema.TypeString,
							Optional: true,
						}, "credentials_env_vars": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"base_var_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"password": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"vault": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"hashicorp_vault_engine": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"hashicorp_vault_secret_path": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"hashicorp_vault_username_key": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"hashicorp_vault_secret_key": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"azure_vault_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"azure_vault_username_key": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"azure_vault_secret_key": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"cyberark_vault_query_string": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"excludes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"follow_symlinks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"parameters": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sync_parameters": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Output
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"namespace_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"namespace_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_replica": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"database_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"content_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"plugin_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creation_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"group_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"engine_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"engine_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"current_timeflow_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"previous_timeflow_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_appdata": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"wait_time": {
				Type:     schema.TypeInt,
				Default:  0,
				Optional: true,
			},
			"skip_wait_for_snapshot_creation": {
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},
		},
	}
}

func toSourceOperationArray(array interface{}) []dctapi.SourceOperation {
	items := []dctapi.SourceOperation{}
	for _, item := range array.([]interface{}) {
		item_map := item.(map[string]interface{})
		sourceOperation := dctapi.NewSourceOperation(item_map["name"].(string), item_map["command"].(string))
		if item_map["shell"].(string) != "" {
			sourceOperation.SetShell(item_map["shell"].(string))
		}
		sourceOperation.SetCredentialsEnvVars(toCredentialsEnvVariableArray(item_map["credentials_env_vars"]))
		items = append(items, *sourceOperation)
	}
	return items
}

func toCredentialsEnvVariableArray(array interface{}) []dctapi.CredentialsEnvVariable {
	items := []dctapi.CredentialsEnvVariable{}
	for _, item := range array.([]interface{}) {
		item_map := item.(map[string]interface{})

		credentialsEnvVariable_item := dctapi.NewCredentialsEnvVariable(item_map["base_var_name"].(string))
		if item_map["password"].(string) != "" {
			credentialsEnvVariable_item.SetPassword(item_map["password"].(string))
		}
		if item_map["vault"].(string) != "" {
			credentialsEnvVariable_item.SetVault(item_map["vault"].(string))
		}
		if item_map["hashicorp_vault_engine"].(string) != "" {
			credentialsEnvVariable_item.SetHashicorpVaultEngine(item_map["hashicorp_vault_engine"].(string))
		}
		if item_map["hashicorp_vault_secret_path"].(string) != "" {
			credentialsEnvVariable_item.SetHashicorpVaultSecretPath(item_map["hashicorp_vault_secret_path"].(string))
		}
		if item_map["hashicorp_vault_username_key"].(string) != "" {
			credentialsEnvVariable_item.SetHashicorpVaultUsernameKey(item_map["hashicorp_vault_username_key"].(string))
		}
		if item_map["hashicorp_vault_secret_key"].(string) != "" {
			credentialsEnvVariable_item.SetHashicorpVaultSecretKey(item_map["hashicorp_vault_secret_key"].(string))
		}
		if item_map["azure_vault_name"].(string) != "" {
			credentialsEnvVariable_item.SetAzureVaultName(item_map["azure_vault_name"].(string))
		}
		if item_map["azure_vault_username_key"].(string) != "" {
			credentialsEnvVariable_item.SetAzureVaultUsernameKey(item_map["azure_vault_username_key"].(string))
		}
		if item_map["azure_vault_secret_key"].(string) != "" {
			credentialsEnvVariable_item.SetAzureVaultSecretKey(item_map["azure_vault_secret_key"].(string))
		}
		if item_map["cyberark_vault_query_string"].(string) != "" {
			credentialsEnvVariable_item.SetCyberarkVaultQueryString(item_map["cyberark_vault_query_string"].(string))
		}
		items = append(items, *credentialsEnvVariable_item)
	}
	return items
}

func resourceAppdataDsourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*apiClient).client

	appDataDSourceLinkSourceParameters := dctapi.NewAppDataDSourceLinkSourceParametersWithDefaults()

	if v, has_v := d.GetOk("name"); has_v {
		appDataDSourceLinkSourceParameters.SetName(v.(string))
	}
	if v, has_v := d.GetOk("source_value"); has_v {
		appDataDSourceLinkSourceParameters.SetSourceId(v.(string))
	}
	if v, has_v := d.GetOk("group_id"); has_v {
		appDataDSourceLinkSourceParameters.SetGroupId(v.(string))
	}
	if v, has_v := d.GetOk("description"); has_v {
		appDataDSourceLinkSourceParameters.SetDescription(v.(string))
	}
	if v, has_v := d.GetOkExists("log_sync_enabled"); has_v {
		appDataDSourceLinkSourceParameters.SetLogSyncEnabled(v.(bool))
	}
	if v, has_v := d.GetOkExists("make_current_account_owner"); has_v {
		appDataDSourceLinkSourceParameters.SetMakeCurrentAccountOwner(v.(bool))
	}
	if v, has_v := d.GetOk("link_type"); has_v {
		appDataDSourceLinkSourceParameters.SetLinkType(v.(string))
	}
	if v, has_v := d.GetOk("staging_mount_base"); has_v {
		appDataDSourceLinkSourceParameters.SetStagingMountBase(v.(string))
	}
	if v, has_v := d.GetOk("staging_environment"); has_v {
		appDataDSourceLinkSourceParameters.SetStagingEnvironment(v.(string))
	}
	if v, has_v := d.GetOk("staging_environment_user"); has_v {
		appDataDSourceLinkSourceParameters.SetStagingEnvironmentUser(v.(string))
	}
	if v, has_v := d.GetOk("environment_user"); has_v {
		appDataDSourceLinkSourceParameters.SetEnvironmentUser(v.(string))
	}
	if v, has_v := d.GetOk("tags"); has_v {
		appDataDSourceLinkSourceParameters.SetTags(toTagArray(v))
	}
	if v, has_v := d.GetOk("ops_pre_sync"); has_v {
		appDataDSourceLinkSourceParameters.SetOpsPreSync(toSourceOperationArray(v))
	}
	if v, has_v := d.GetOk("ops_post_sync"); has_v {
		appDataDSourceLinkSourceParameters.SetOpsPostSync(toSourceOperationArray(v))
	}
	if v, has_v := d.GetOkExists("excludes"); has_v {
		appDataDSourceLinkSourceParameters.SetExcludes(toStringArray(v))
	}
	if v, has_v := d.GetOkExists("follow_symlinks"); has_v {
		appDataDSourceLinkSourceParameters.SetFollowSymlinks(toStringArray(v))
	}
	if v, has_v := d.GetOk("parameters"); has_v {
		params := make(map[string]interface{})
		json.Unmarshal([]byte(v.(string)), &params)
		appDataDSourceLinkSourceParameters.SetParameters(params)
	}
	if v, has_v := d.GetOk("sync_parameters"); has_v {
		sync_params := make(map[string]interface{})
		json.Unmarshal([]byte(v.(string)), &sync_params)
		appDataDSourceLinkSourceParameters.SetSyncParameters(sync_params)
	}

	req := client.DSourcesAPI.LinkAppdataDatabase(ctx)

	apiRes, httpRes, err := req.AppDataDSourceLinkSourceParameters(*appDataDSourceLinkSourceParameters).Execute()
	if diags := apiErrorResponseHelper(ctx, apiRes, httpRes, err); diags != nil {
		return diags
	}

	d.SetId(*apiRes.DsourceId)

	job_res, job_err := PollJobStatus(*apiRes.Job.Id, ctx, client)
	if job_err != "" {
		tflog.Error(ctx, DLPX+ERROR+"Job Polling failed but continuing with dSource creation. Error: "+job_err)
	}

	tflog.Info(ctx, DLPX+INFO+"Job result is "+job_res)

	if job_res == Failed || job_res == Canceled || job_res == Abandoned {
		d.SetId("")
		tflog.Error(ctx, DLPX+ERROR+"Job "+job_res+" "+*apiRes.Job.Id+"!")
		return diag.Errorf("[NOT OK] Job %s %s with error %s", *apiRes.Job.Id, job_res, job_err)
	}

	PollSnapshotStatus(d, ctx, client)

	readDiags := resourceDsourceRead(ctx, d, meta)

	if readDiags.HasError() {
		return readDiags
	}

	return diags
}

func resourceDsourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*apiClient).client

	dsource_id := d.Id()

	res, diags := PollForObjectExistence(ctx, func() (interface{}, *http.Response, error) {
		return client.DSourcesAPI.GetDsourceById(ctx, dsource_id).Execute()
	})

	if res == nil {
		tflog.Error(ctx, DLPX+ERROR+"Dsource not found: "+dsource_id+", removing from state. ")
		d.SetId("")
		return nil
	}

	if diags != nil {
		_, diags := PollForObjectDeletion(ctx, func() (interface{}, *http.Response, error) {
			return client.DSourcesAPI.GetDsourceById(ctx, dsource_id).Execute()
		})
		// This would imply error in poll for deletion so we just log and exit.
		if diags != nil {
			tflog.Error(ctx, DLPX+ERROR+"Error in polling of dSource for deletion.")
		} else {
			// diags will be nil in case of successful poll for deletion logic aka 404
			tflog.Error(ctx, DLPX+ERROR+"Error reading the dSource "+dsource_id+", removing from state.")
			d.SetId("")
		}

		return nil
	}

	result, ok := res.(*dctapi.DSource)
	if !ok {
		return diag.Errorf("Error occured in type casting.")
	}

	d.Set("id", result.GetId())
	d.Set("database_type", result.GetDatabaseType())
	d.Set("name", result.GetName())
	d.Set("is_replica", result.GetIsReplica())
	d.Set("plugin_version", result.GetPluginVersion())
	d.Set("creation_date", result.GetCreationDate().String())
	d.Set("group_name", result.GetGroupName())
	d.Set("enabled", result.GetEnabled())
	d.Set("engine_id", result.GetEngineId())
	d.Set("source_id", result.GetSourceId())
	d.Set("status", result.GetStatus())
	d.Set("engine_name", result.GetEngineName())
	d.Set("current_timeflow_id", result.GetCurrentTimeflowId())
	d.Set("is_appdata", result.GetIsAppdata())

	return diags
}

func resourceDsourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	return diag.Errorf("Action update not implemented for resource : dSource")
}

func resourceDsourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient).client

	dsourceId := d.Id()

	deleteDsourceParams := dctapi.NewDeleteDSourceRequest(dsourceId)
	deleteDsourceParams.SetForce(false)

	res, httpRes, err := client.DSourcesAPI.DeleteDsource(ctx).DeleteDSourceRequest(*deleteDsourceParams).Execute()

	if diags := apiErrorResponseHelper(ctx, res, httpRes, err); diags != nil {
		return diags
	}

	job_status, job_err := PollJobStatus(*res.Id, ctx, client)
	if job_err != "" {
		tflog.Warn(ctx, DLPX+WARN+"Job Polling failed but continuing with deletion. Error :"+job_err)
	}
	tflog.Info(ctx, DLPX+INFO+"Job result is "+job_status)
	if isJobTerminalFailure(job_status) {
		return diag.Errorf("[NOT OK] dSource-Delete %s. JobId: %s / Error: %s", job_status, *res.Id, job_err)
	}

	_, diags := PollForObjectDeletion(ctx, func() (interface{}, *http.Response, error) {
		return client.DSourcesAPI.GetDsourceById(ctx, dsourceId).Execute()
	})

	return diags
}
