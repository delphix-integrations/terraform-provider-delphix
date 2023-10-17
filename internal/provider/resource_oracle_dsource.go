package provider

import (
	"context"
	"net/http"

	dctapi "github.com/delphix/dct-sdk-go/v10"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOracleDsource() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Resource for Oracle dSource creation.",

		CreateContext: resourceOracleDsourceCreate,
		ReadContext:   resourceOracleDsourceRead,
		UpdateContext: resourceOracleDsourceUpdate,
		DeleteContext: resourceOracleDsourceDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_id": {
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
			"external_file_path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"environment_user_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"backup_level_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"rman_channels": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"files_per_set": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"check_logical": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			////////////////////////////TODO ADD REMAINING INPUTS/////

			// Output
			"id": {
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
			"storage_size": {
				Type:     schema.TypeInt,
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
		},
	}
}

func resourceOracleDsourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*apiClient).client

	oracleDSourceLinkSourceParameters := dctapi.NewOracleDSourceLinkSourceParametersWithDefaults()

	if v, has_v := d.GetOk("name"); has_v {
		oracleDSourceLinkSourceParameters.SetName(v.(string))
	}
	if v, has_v := d.GetOk("source_id"); has_v {
		oracleDSourceLinkSourceParameters.SetSourceId(v.(string))
	}
	if v, has_v := d.GetOk("group_id"); has_v {
		oracleDSourceLinkSourceParameters.SetGroupId(v.(string))
	}
	if v, has_v := d.GetOk("description"); has_v {
		oracleDSourceLinkSourceParameters.SetDescription(v.(string))
	}
	if v, has_v := d.GetOkExists("log_sync_enabled"); has_v {
		oracleDSourceLinkSourceParameters.SetLogSyncEnabled(v.(bool))
	}
	if v, has_v := d.GetOkExists("make_current_account_owner"); has_v {
		oracleDSourceLinkSourceParameters.SetMakeCurrentAccountOwner(v.(bool))
	}
	if v, has_v := d.GetOk("tags"); has_v {
		oracleDSourceLinkSourceParameters.SetTags(toTagArray(v))
	}
	if v, has_v := d.GetOk("ops_pre_sync"); has_v {
		oracleDSourceLinkSourceParameters.SetOpsPreSync(toSourceOperationArray(v))
	}
	if v, has_v := d.GetOk("ops_post_sync"); has_v {
		oracleDSourceLinkSourceParameters.SetOpsPostSync(toSourceOperationArray(v))
	}
	if v, has_v := d.GetOk("external_file_path"); has_v {
		oracleDSourceLinkSourceParameters.SetExternalFilePath(v.(string))
	}
	if v, has_v := d.GetOk("environment_user_id"); has_v {
		oracleDSourceLinkSourceParameters.SetEnvironmentUserId(v.(string))
	}
	if v, has_v := d.GetOkExists("backup_level_enabled"); has_v {
		oracleDSourceLinkSourceParameters.SetBackupLevelEnabled(v.(bool))
	}
	if v, has_v := d.GetOk("rman_channels"); has_v {
		oracleDSourceLinkSourceParameters.SetRmanChannels(v.(int32))
	}
	if v, has_v := d.GetOk("files_per_set"); has_v {
		oracleDSourceLinkSourceParameters.SetFilesPerSet(v.(int32))
	}
	if v, has_v := d.GetOkExists("check_logical"); has_v {
		oracleDSourceLinkSourceParameters.SetCheckLogical(v.(bool))
	}
	if v, has_v := d.GetOkExists("encrypted_linking_enabled"); has_v {
		oracleDSourceLinkSourceParameters.SetEncryptedLinkingEnabled(v.(bool))
	}
	if v, has_v := d.GetOkExists("compressed_linking_enabled"); has_v {
		oracleDSourceLinkSourceParameters.SetCompressedLinkingEnabled(v.(bool))
	}
	if v, has_v := d.GetOk("bandwidth_limit"); has_v {
		oracleDSourceLinkSourceParameters.SetBandwidthLimit(v.(int32))
	}
	if v, has_v := d.GetOk("number_of_connections"); has_v {
		oracleDSourceLinkSourceParameters.SetNumberOfConnections(v.(int32))
	}
	if v, has_v := d.GetOkExists("diagnose_no_logging_faults"); has_v {
		oracleDSourceLinkSourceParameters.SetCheckLogical(v.(bool))
	}
	if v, has_v := d.GetOkExists("pre_provisioning_enabled"); has_v {
		oracleDSourceLinkSourceParameters.SetCheckLogical(v.(bool))
	}
	if v, has_v := d.GetOkExists("link_now"); has_v {
		oracleDSourceLinkSourceParameters.SetCheckLogical(v.(bool))
	}
	if v, has_v := d.GetOkExists("force_full_backup"); has_v {
		oracleDSourceLinkSourceParameters.SetCheckLogical(v.(bool))
	}
	if v, has_v := d.GetOkExists("double_sync"); has_v {
		oracleDSourceLinkSourceParameters.SetCheckLogical(v.(bool))
	}
	if v, has_v := d.GetOkExists("skip_space_check"); has_v {
		oracleDSourceLinkSourceParameters.SetCheckLogical(v.(bool))
	}
	if v, has_v := d.GetOkExists("do_not_resume"); has_v {
		oracleDSourceLinkSourceParameters.SetCheckLogical(v.(bool))
	}
	if v, has_v := d.GetOk("files_for_full_backup"); has_v {
		oracleDSourceLinkSourceParameters.SetFilesForFullBackup(toIntArray(v))
	}
	if v, has_v := d.GetOk("log_sync_mode"); has_v {
		oracleDSourceLinkSourceParameters.SetLogSyncMode(v.(string))
	}
	if v, has_v := d.GetOk("log_sync_interval"); has_v {
		oracleDSourceLinkSourceParameters.SetLogSyncInterval(v.(int32))
	}
	if v, has_v := d.GetOk("non_sys_username"); has_v {
		oracleDSourceLinkSourceParameters.SetNonSysUsername(v.(string))
	}
	if v, has_v := d.GetOk("non_sys_password"); has_v {
		oracleDSourceLinkSourceParameters.SetNonSysPassword(v.(string))
	}

	if v, has_v := d.GetOk("non_sys_vault"); has_v {
		oracleDSourceLinkSourceParameters.SetNonSysVault(v.(string))
	}
	if v, has_v := d.GetOk("non_sys_hashicorp_vault_engine"); has_v {
		oracleDSourceLinkSourceParameters.SetNonSysHashicorpVaultEngine(v.(string))
	}
	if v, has_v := d.GetOk("non_sys_hashicorp_vault_secret_path"); has_v {
		oracleDSourceLinkSourceParameters.SetNonSysHashicorpVaultSecretPath(v.(string))
	}
	if v, has_v := d.GetOk("non_sys_hashicorp_vault_username_key"); has_v {
		oracleDSourceLinkSourceParameters.SetNonSysHashicorpVaultUsernameKey(v.(string))
	}

	if v, has_v := d.GetOk("non_sys_hashicorp_vault_secret_key"); has_v {
		oracleDSourceLinkSourceParameters.SetNonSysHashicorpVaultSecretKey(v.(string))
	}

	if v, has_v := d.GetOk("non_sys_azure_vault_name"); has_v {
		oracleDSourceLinkSourceParameters.SetNonSysAzureVaultName(v.(string))
	}

	if v, has_v := d.GetOk("non_sys_azure_vault_username_key"); has_v {
		oracleDSourceLinkSourceParameters.SetNonSysAzureVaultUsernameKey(v.(string))
	}

	if v, has_v := d.GetOk("non_sys_azure_vault_secret_key"); has_v {
		oracleDSourceLinkSourceParameters.SetNonSysAzureVaultSecretKey(v.(string))
	}

	if v, has_v := d.GetOk("non_sys_cyberark_vault_query_string"); has_v {
		oracleDSourceLinkSourceParameters.SetNonSysCyberarkVaultQueryString(v.(string))
	}
	if v, has_v := d.GetOk("fallback_username"); has_v {
		oracleDSourceLinkSourceParameters.SetFallbackUsername(v.(string))
	}
	if v, has_v := d.GetOk("fallback_password"); has_v {
		oracleDSourceLinkSourceParameters.SetFallbackPassword(v.(string))
	}
	if v, has_v := d.GetOk("fallback_vault"); has_v {
		oracleDSourceLinkSourceParameters.SetFallbackVault(v.(string))
	}
	if v, has_v := d.GetOk("fallback_hashicorp_vault_engine"); has_v {
		oracleDSourceLinkSourceParameters.SetFallbackHashicorpVaultEngine(v.(string))
	}
	if v, has_v := d.GetOk("fallback_hashicorp_vault_secret_path"); has_v {
		oracleDSourceLinkSourceParameters.SetFallbackHashicorpVaultSecretPath(v.(string))
	}
	if v, has_v := d.GetOk("fallback_hashicorp_vault_username_key"); has_v {
		oracleDSourceLinkSourceParameters.SetFallbackHashicorpVaultUsernameKey(v.(string))
	}
	if v, has_v := d.GetOk("fallback_hashicorp_vault_secret_key"); has_v {
		oracleDSourceLinkSourceParameters.SetFallbackHashicorpVaultSecretKey(v.(string))
	}
	if v, has_v := d.GetOk("fallback_azure_vault_name"); has_v {
		oracleDSourceLinkSourceParameters.SetFallbackAzureVaultName(v.(string))
	}
	if v, has_v := d.GetOk("fallback_azure_vault_username_key"); has_v {
		oracleDSourceLinkSourceParameters.SetFallbackAzureVaultUsernameKey(v.(string))
	}
	if v, has_v := d.GetOk("fallback_azure_vault_secret_key"); has_v {
		oracleDSourceLinkSourceParameters.SetFallbackAzureVaultSecretKey(v.(string))
	}
	if v, has_v := d.GetOk("fallback_cyberark_vault_query_string"); has_v {
		oracleDSourceLinkSourceParameters.SetFallbackCyberarkVaultQueryString(v.(string))
	}
	if v, has_v := d.GetOk("ops_pre_log_sync"); has_v {
		oracleDSourceLinkSourceParameters.SetOpsPreLogSync(toSourceOperationArray(v))
	}

	req := client.DSourcesApi.LinkOracleDatabase(ctx)

	apiRes, httpRes, err := req.OracleDSourceLinkSourceParameters(*oracleDSourceLinkSourceParameters).Execute()
	if diags := apiErrorResponseHelper(apiRes, httpRes, err); diags != nil {
		return diags
	}

	d.SetId(*apiRes.DsourceId)

	job_res, job_err := PollJobStatus(*apiRes.Job.Id, ctx, client)
	if job_err != "" {
		ErrorLog.Printf("Job Polling failed but continuing with dSource creation. Error: %s", job_err)
	}

	InfoLog.Printf("Job result is %s", job_res)

	if job_res == Failed || job_res == Canceled || job_res == Abandoned {
		d.SetId("")
		ErrorLog.Printf("Job %s %s!", job_res, *apiRes.Job.Id)
		return diag.Errorf("[NOT OK] Job %s %s with error %s", *apiRes.Job.Id, job_res, job_err)
	}

	readDiags := resourceOracleDsourceRead(ctx, d, meta)

	if readDiags.HasError() {
		return readDiags
	}

	return diags
}

func toIntArray(array interface{}) []int32 {
	items := []int32{}
	for _, item := range array.([]interface{}) {
		items = append(items, item.(int32))
	}
	return items
}

func resourceOracleDsourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*apiClient).client

	dsource_id := d.Id()

	res, diags := PollForObjectExistence(func() (interface{}, *http.Response, error) {
		return client.DSourcesApi.GetDsourceById(ctx, dsource_id).Execute()
	})

	if diags != nil {
		_, diags := PollForObjectDeletion(func() (interface{}, *http.Response, error) {
			return client.DSourcesApi.GetDsourceById(ctx, dsource_id).Execute()
		})
		// This would imply error in poll for deletion so we just log and exit.
		if diags != nil {
			ErrorLog.Printf("Error in polling of appdata dSource for deletion.")
		} else {
			// diags will be nill in case of successful poll for deletion logic aka 404
			ErrorLog.Printf("Error reading the appdata dSource %s, removing from state.", dsource_id)
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
	d.Set("storage_size", result.GetStorageSize())
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

func resourceOracleDsourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	return diag.Errorf("Action update not implemented for resource : appdata dsource")
}

func resourceOracleDsourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient).client

	dsourceId := d.Id()

	deleteDsourceParams := dctapi.NewDeleteDSourceRequest(dsourceId)
	deleteDsourceParams.SetForce(false)

	res, httpRes, err := client.DSourcesApi.DeleteDsource(ctx).DeleteDSourceRequest(*deleteDsourceParams).Execute()

	if diags := apiErrorResponseHelper(res, httpRes, err); diags != nil {
		return diags
	}

	job_status, job_err := PollJobStatus(*res.Id, ctx, client)
	if job_err != "" {
		WarnLog.Printf("Job Polling failed but continuing with deletion. Error :%v", job_err)
	}
	InfoLog.Printf("Job result is %s", job_status)
	if isJobTerminalFailure(job_status) {
		return diag.Errorf("[NOT OK] Appdata dSource-Delete %s. JobId: %s / Error: %s", job_status, *res.Id, job_err)
	}

	_, diags := PollForObjectDeletion(func() (interface{}, *http.Response, error) {
		return client.DSourcesApi.GetDsourceById(ctx, dsourceId).Execute()
	})

	return diags
}
