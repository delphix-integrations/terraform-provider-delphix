package provider

import (
	"context"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"

	dctapi "github.com/delphix/dct-sdk-go/v23"
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
			"source_value": {
				Type:     schema.TypeString,
				Required: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return true
				},
			},
			"group_id": {
				Type:     schema.TypeString,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return true
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_sync_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"make_current_account_owner": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return true
				},
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
						},
						"element_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"credentials_env_vars": {
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
						},
						"element_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"credentials_env_vars": {
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
						},
						"element_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"credentials_env_vars": {
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
			"database_type": {
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
			"is_detached": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"engine_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_id": {
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
			"is_appdata": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"sync_policy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"retention_policy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exported_data_directory": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"wait_time": {
				Type:     schema.TypeInt,
				Default:  0,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return true
				},
			},
			"skip_wait_for_snapshot_creation": {
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return true
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
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
	req := client.DSourcesAPI.LinkOracleDatabase(ctx)

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

	readDiags := resourceOracleDsourceRead(ctx, d, meta)

	if readDiags.HasError() {
		return readDiags
	}

	return diags
}

func resourceOracleDsourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

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

	ops_pre_sync_Raw, _ := d.Get("ops_pre_sync").([]interface{})
	oldOpsPreSync := toSourceOperationArray(ops_pre_sync_Raw)

	ops_post_sync_Raw, _ := d.Get("ops_post_sync").([]interface{})
	oldOpsPostSync := toSourceOperationArray(ops_post_sync_Raw)

	ops_pre_log_sync_Raw, _ := d.Get("ops_pre_log_sync").([]interface{})
	oldOpsPreLogSync := toSourceOperationArray(ops_pre_log_sync_Raw)

	d.Set("id", result.GetId())
	d.Set("database_type", result.GetDatabaseType())
	d.Set("name", result.GetName())
	d.Set("is_replica", result.GetIsReplica())
	d.Set("database_version", result.GetDatabaseVersion())
	d.Set("content_type", result.GetContentType())
	d.Set("data_uuid", result.GetDataUuid())
	d.Set("creation_date", result.GetCreationDate().String())
	d.Set("group_name", result.GetGroupName())
	d.Set("enabled", result.GetEnabled())
	d.Set("is_detached", result.GetIsDetached())
	d.Set("engine_id", result.GetEngineId())
	d.Set("source_id", result.GetSourceId())
	d.Set("status", result.GetStatus())
	d.Set("engine_name", result.GetEngineName())
	d.Set("current_timeflow_id", result.GetCurrentTimeflowId())
	d.Set("is_appdata", result.GetIsAppdata())
	d.Set("hooks", result.GetHooks())
	d.Set("sync_policy_id", result.GetSyncPolicyId())
	d.Set("retention_policy_id", result.GetReplicaRetentionPolicyId())
	d.Set("log_sync_enabled", result.GetLogsyncEnabled())
	d.Set("exported_data_directory", result.GetExportedDataDirectory())
	d.Set("ops_pre_sync", flattenDSourceHooks(result.GetHooks().OpsPreSync, oldOpsPreSync))
	d.Set("ops_post_sync", flattenDSourceHooks(result.GetHooks().OpsPostSync, oldOpsPostSync))
	d.Set("ops_pre_log_sync", flattenDSourceHooks(result.GetHooks().OpsPreLogSync, oldOpsPreLogSync))
	return diags
}

func resourceOracleDsourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	client := meta.(*apiClient).client
	updateOracleDsource := dctapi.NewUpdateOracleDsourceParameters()

	dsourceId := d.Get("id").(string)

	// get the changed keys
	changedKeys := make([]string, 0, len(d.State().Attributes))
	for k := range d.State().Attributes {
		if strings.Contains(k, "tags") { // this is because the changed keys are of the form tag.0.keydi
			k = "tags"
		}
		if strings.Contains(k, "ops_pre_sync") {
			k = "ops_pre_sync"
		}
		if strings.Contains(k, "ops_pre_log_sync") {
			k = "ops_pre_log_sync"
		}
		if strings.Contains(k, "ops_post_sync") {
			k = "ops_post_sync"
		}
		if d.HasChange(k) {
			tflog.Debug(ctx, "changed keys"+k)
			changedKeys = append(changedKeys, k)
		}
	}

	var updateFailure bool = false
	var nonUpdatableField []string

	// check if the changed keys are updatable
	for _, key := range changedKeys {
		if !updatableOracleDsourceKeys[key] {
			updateFailure = true
			tflog.Debug(ctx, "non updatable field: "+key)
			nonUpdatableField = append(nonUpdatableField, key)
		}
	}

	// if not updatable keys are provided, error out
	if updateFailure {
		revertChanges(d, changedKeys)
		return diag.Errorf("cannot update options %v. Please refer to provider documentation for updatable params.", nonUpdatableField)
	}

	// set changed params in the updateOracleDsource
	if d.HasChange("name") {
		updateOracleDsource.SetName(d.Get("name").(string))
	}
	if d.HasChange("environment_user_id") {
		updateOracleDsource.SetEnvironmentUserId(d.Get("environment_user_id").(string))
	}
	if d.HasChange("backup_level_enabled") {
		updateOracleDsource.SetBackupLevelEnabled(d.Get("backup_level_enabled").(bool))
	}
	if d.HasChange("rman_channels") {
		updateOracleDsource.SetRmanChannels(int32(d.Get("rman_channels").(int)))
	}
	if d.HasChange("files_per_set") {
		updateOracleDsource.SetFilesPerSet(int32(d.Get("files_per_set").(int)))
	}
	if d.HasChange("check_logical") {
		updateOracleDsource.SetCheckLogical(d.Get("check_logical").(bool))
	}
	if d.HasChange("encrypted_linking_enabled") {
		updateOracleDsource.SetEncryptedLinkingEnabled(d.Get("encrypted_linking_enabled").(bool))
	}
	if d.HasChange("compressed_linking_enabled") {
		updateOracleDsource.SetCompressedLinkingEnabled(d.Get("compressed_linking_enabled").(bool))
	}
	if d.HasChange("bandwidth_limit") {
		updateOracleDsource.SetBandwidthLimit(int32(d.Get("bandwidth_limit").(int)))
	}
	if d.HasChange("number_of_connections") {
		updateOracleDsource.SetNumberOfConnections(int32(d.Get("number_of_connections").(int)))
	}
	if d.HasChange("pre_provisioning_enabled") {
		updateOracleDsource.SetPreProvisioningEnabled(d.Get("pre_provisioning_enabled").(bool))
	}
	if d.HasChange("diagnose_no_logging_faults") {
		updateOracleDsource.SetDiagnoseNoLoggingFaults(d.Get("diagnose_no_logging_faults").(bool))
	}
	if d.HasChange("external_file_path") {
		updateOracleDsource.SetExternalFilePath(d.Get("external_file_path").(string))
	}

	// update hooks
	ndsh := dctapi.NewDSourceHooks()

	if d.HasChange("ops_pre_sync") {
		if v, has_v := d.GetOk("ops_pre_sync"); has_v {
			ndsh.SetOpsPreSync(toHookArray(v))
		} else {
			ndsh.SetOpsPreSync([]dctapi.Hook{})
		}
	}

	if d.HasChange("ops_pre_log_sync") {
		if v, has_v := d.GetOk("ops_pre_log_sync"); has_v {
			ndsh.SetOpsPreLogSync(toHookArray(v))
		} else {
			ndsh.SetOpsPreLogSync([]dctapi.Hook{})
		}
	}

	if d.HasChange("ops_post_sync") {
		if v, has_v := d.GetOk("ops_post_sync"); has_v {
			ndsh.SetOpsPostSync(toHookArray(v))
		} else {
			ndsh.SetOpsPostSync([]dctapi.Hook{})
		}
	}

	if ndsh != nil {
		updateOracleDsource.SetHooks(*ndsh)
	}

	res, httpRes, err := client.DSourcesAPI.UpdateOracleDsourceById(ctx, dsourceId).UpdateOracleDsourceParameters(*updateOracleDsource).Execute()

	if diags := apiErrorResponseHelper(ctx, nil, httpRes, err); diags != nil {
		// revert and set the old value to the changed keys
		revertChanges(d, changedKeys)
		return diags
	}

	job_status, job_err := PollJobStatus(*res.Job.Id, ctx, client)
	if job_err != "" {
		tflog.Warn(ctx, DLPX+WARN+"Dsource Update Job Polling failed but continuing with update. Error: "+job_err)
	}
	tflog.Info(ctx, DLPX+INFO+"Job result is "+job_status)
	if isJobTerminalFailure(job_status) {
		return diag.Errorf("[NOT OK] Dsource-Update %s. JobId: %s / Error: %s", job_status, *res.Job.Id, job_err)
	}

	if d.HasChanges(
		"tags",
	) { // tags update
		tflog.Debug(ctx, "updating tags")
		if d.HasChange("tags") {
			// delete old tag
			tflog.Debug(ctx, "deleting old tags")
			oldTag, newTag := d.GetChange("tags")
			if len(toTagArray(oldTag)) != 0 {
				tflog.Debug(ctx, "tag to be deleted: "+toTagArray(oldTag)[0].GetKey()+" "+toTagArray(oldTag)[0].GetValue())
				deleteTag := *dctapi.NewDeleteTag()
				tagDelResp, tagDelErr := client.DSourcesAPI.DeleteTagsDsource(ctx, dsourceId).DeleteTag(deleteTag).Execute()
				if diags := apiErrorResponseHelper(ctx, nil, tagDelResp, tagDelErr); diags != nil {
					revertChanges(d, changedKeys)
					updateFailure = true
				}
			}
			// create tag
			if len(toTagArray(newTag)) != 0 {
				tflog.Info(ctx, "creating new tags")
				_, httpResp, tagCrtErr := client.DSourcesAPI.CreateTagsDsource(ctx, dsourceId).TagsRequest(*dctapi.NewTagsRequest(toTagArray(newTag))).Execute()
				if diags := apiErrorResponseHelper(ctx, nil, httpResp, tagCrtErr); diags != nil {
					revertChanges(d, changedKeys)
					return diags
				}
			}
		}
	}

	return diags
}

func resourceOracleDsourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
