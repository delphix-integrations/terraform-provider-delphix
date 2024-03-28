package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	dctapi "github.com/delphix/dct-sdk-go/v14"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOracleDsource() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Resource for Oracle dSource creation.",

		CreateContext: resourceOracleDsourceCreate,
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
				Optional: true,
			},
			"environment_user_id": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
			},
			"check_logical": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"encrypted_linking_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"compressed_linking_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"bandwidth_limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"number_of_connections": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"diagnose_no_logging_faults": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"pre_provisioning_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"link_now": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"force_full_backup": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"double_sync": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"skip_space_check": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"do_not_resume": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"files_for_full_backup": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"log_sync_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"log_sync_interval": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"non_sys_username": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"non_sys_password": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"non_sys_vault": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"non_sys_hashicorp_vault_engine": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"non_sys_hashicorp_vault_secret_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"non_sys_hashicorp_vault_username_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"non_sys_hashicorp_vault_secret_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"non_sys_azure_vault_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"non_sys_azure_vault_username_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"non_sys_azure_vault_secret_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"non_sys_cyberark_vault_query_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fallback_username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fallback_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fallback_vault": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fallback_hashicorp_vault_engine": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fallback_hashicorp_vault_secret_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fallback_hashicorp_vault_username_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fallback_hashicorp_vault_secret_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fallback_azure_vault_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fallback_azure_vault_username_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fallback_azure_vault_secret_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fallback_cyberark_vault_query_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ops_pre_log_sync": {
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
				Default:  3,
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

func resourceOracleDsourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*apiClient).client

	oracleDSourceLinkSourceParameters := dctapi.NewOracleDSourceLinkSourceParametersWithDefaults()

	if v, has_v := d.GetOk("name"); has_v {
		oracleDSourceLinkSourceParameters.SetName(v.(string))
	}
	if v, has_v := d.GetOk("source_value"); has_v {
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
		oracleDSourceLinkSourceParameters.SetRmanChannels(int32(v.(int)))
	}
	if v, has_v := d.GetOk("files_per_set"); has_v {
		oracleDSourceLinkSourceParameters.SetFilesPerSet(int32(v.(int)))
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
		oracleDSourceLinkSourceParameters.SetBandwidthLimit(int32(v.(int)))
	}
	if v, has_v := d.GetOk("number_of_connections"); has_v {
		oracleDSourceLinkSourceParameters.SetNumberOfConnections(int32(v.(int)))
	}
	if v, has_v := d.GetOkExists("diagnose_no_logging_faults"); has_v {
		oracleDSourceLinkSourceParameters.SetDiagnoseNoLoggingFaults(v.(bool))
	}
	if v, has_v := d.GetOkExists("pre_provisioning_enabled"); has_v {
		oracleDSourceLinkSourceParameters.SetPreProvisioningEnabled(v.(bool))
	}
	if v, has_v := d.GetOkExists("link_now"); has_v {
		oracleDSourceLinkSourceParameters.SetLinkNow(v.(bool))
	}
	if v, has_v := d.GetOkExists("force_full_backup"); has_v {
		oracleDSourceLinkSourceParameters.SetForceFullBackup(v.(bool))
	}
	if v, has_v := d.GetOkExists("double_sync"); has_v {
		oracleDSourceLinkSourceParameters.SetDoubleSync(v.(bool))
	}
	if v, has_v := d.GetOkExists("skip_space_check"); has_v {
		oracleDSourceLinkSourceParameters.SetSkipSpaceCheck(v.(bool))
	}
	if v, has_v := d.GetOkExists("do_not_resume"); has_v {
		oracleDSourceLinkSourceParameters.SetDoNotResume(v.(bool))
	}
	if v, has_v := d.GetOk("files_for_full_backup"); has_v {
		oracleDSourceLinkSourceParameters.SetFilesForFullBackup(toIntArray(v))
	}
	if v, has_v := d.GetOk("log_sync_mode"); has_v {
		oracleDSourceLinkSourceParameters.SetLogSyncMode(v.(string))
	}
	if v, has_v := d.GetOk("log_sync_interval"); has_v {
		oracleDSourceLinkSourceParameters.SetLogSyncInterval(int32(v.(int)))
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

func toIntArray(array interface{}) []int32 {
	items := []int32{}
	for _, item := range array.([]interface{}) {
		items = append(items, int32(item.(int)))
	}
	return items
}
