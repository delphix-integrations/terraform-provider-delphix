package provider

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
	"time"

	dctapi "github.com/delphix/dct-sdk-go/v25"
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
		CustomizeDiff: CustomizeDiffTags,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rollback_on_failure": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"source_value": {
				Type:     schema.TypeString,
				Optional: true,
				DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
					if oldValue != newValue {
						tflog.Info(context.Background(), "updating source_value is not allowed. plan changes are suppressed")
					}
					return d.Id() != ""
				},
			},
			"group_id": {
				Type:     schema.TypeString,
				Optional: true,
				DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
					if oldValue != newValue {
						tflog.Info(context.Background(), "updating group_id is not allowed. plan changes are suppressed")
					}
					return d.Id() != ""
				},
			},
			"sync_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"retention_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Default:  true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// Suppress diff ONLY when upgrading from null/empty to default true (silent upgrade)
					// Do NOT suppress when user explicitly changes from false to true
					if (old == "" || old == "<null>") && new == "true" {
						rawConfig := d.GetRawConfig()
						if rawConfig.IsKnown() && !rawConfig.IsNull() {
							attr := rawConfig.GetAttr("make_current_account_owner")
							if attr.IsNull() || !attr.IsKnown() {
								return true
							}
						}
					}
					return false
				},
			},
			"link_type": {
				Type:     schema.TypeString,
				Optional: true,
				DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
					tflog.Info(context.Background(), "In DiffSuppressFunc of link_type")
					if oldValue != newValue {
						tflog.Info(context.Background(), "updating link_type is not allowed. plan changes are suppressed")
					}
					return d.Id() != ""
				},
			},
			"staging_mount_base": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"staging_environment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"staging_environment_user": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"environment_user": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ignore_tag_changes": {
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// Suppress diff ONLY when upgrading from null/empty to default true (silent upgrade)
					// Do NOT suppress when user explicitly changes from false to true
					if (old == "" || old == "<null>") && new == "true" {
						rawConfig := d.GetRawConfig()
						if rawConfig.IsKnown() && !rawConfig.IsNull() {
							attr := rawConfig.GetAttr("ignore_tag_changes")
							if attr.IsNull() || !attr.IsKnown() {
								return true
							}
						}
					}
					return false
				},
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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
				DiffSuppressFunc: func(_, old, new string, d *schema.ResourceData) bool {
					if ignore, ok := d.GetOk("ignore_tag_changes"); ok && ignore.(bool) {
						return true
					}
					return false
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
				Optional: true,
			},
			"sync_parameters": {
				Type:     schema.TypeString,
				Optional: true,
				DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
					if oldValue != newValue {
						tflog.Info(context.Background(), "updating sync_parameters is not allowed. plan changes are suppressed")
					}
					return d.Id() != ""
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
				Default:  0,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if old != new {
						tflog.Info(context.Background(), "updating wait_time is not allowed. plan changes are suppressed")
					}
					return d.Id() != ""
				},
			},
			"skip_wait_for_snapshot_creation": {
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if old != new {
						tflog.Info(context.Background(), "updating skip_wait_for_snapshot_creation is not allowed. plan changes are suppressed")
					}
					return d.Id() != ""
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAppdataDsourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*apiClient).client

	// respect resource create timeout
	createCtx, createCancel := context.WithTimeout(ctx, d.Timeout(schema.TimeoutCreate))
	defer createCancel()

	appDataDSourceLinkSourceParameters := dctapi.NewAppDataDSourceLinkSourceParametersWithDefaults()

	if v, has_v := d.GetOk("name"); has_v {
		appDataDSourceLinkSourceParameters.SetName(v.(string))
	}
	if v, has_v := d.GetOk("source_value"); has_v {
		appDataDSourceLinkSourceParameters.SetSourceId(v.(string))
	}
	if v, has_v := d.GetOk("sync_policy_id"); has_v {
		appDataDSourceLinkSourceParameters.SetSyncPolicyId(v.(string))
	}
	if v, has_v := d.GetOk("retention_policy_id"); has_v {
		appDataDSourceLinkSourceParameters.SetRetentionPolicyId(v.(string))
	}
	if v, has_v := d.GetOk("group_id"); has_v {
		appDataDSourceLinkSourceParameters.SetGroupId(v.(string))
	}
	if v, has_v := d.GetOk("description"); has_v {
		appDataDSourceLinkSourceParameters.SetDescription(v.(string))
	}
	if v, has_v := d.GetOk("log_sync_enabled"); has_v {
		appDataDSourceLinkSourceParameters.SetLogSyncEnabled(v.(bool))
	}
	if v, has_v := d.GetOk("make_current_account_owner"); has_v {
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
	if v, has_v := d.GetOk("excludes"); has_v {
		appDataDSourceLinkSourceParameters.SetExcludes(toStringArray(v))
	}
	if v, has_v := d.GetOk("follow_symlinks"); has_v {
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

	req := client.DSourcesAPI.LinkAppdataDatabase(createCtx)

	apiRes, httpRes, err := req.AppDataDSourceLinkSourceParameters(*appDataDSourceLinkSourceParameters).Execute()
	
	// Check if the API call itself timed out
	if err != nil && createCtx.Err() == context.DeadlineExceeded {
		return diag.Errorf("dSource creation API call timed out after %s. "+
			"Check DCT UI for job status. If created, find the dSource ID and import it.",
			d.Timeout(schema.TimeoutCreate))
	}
	
	if diags := apiErrorResponseHelper(ctx, apiRes, httpRes, err); diags != nil {
		return diags
	}

	// Check for nil apiRes or Job to prevent crashes
	if apiRes == nil || apiRes.Job == nil {
		return diag.Errorf("dSource creation failed: received nil response or job from API")
	}

	// Store dSource ID temporarily - don't set in state until job completes
	dsourceId := apiRes.GetDsourceId()

	job_res, job_err := PollJobStatus(apiRes.Job.GetId(), createCtx, client)
	if job_err != "" {
		tflog.Error(ctx, DLPX+ERROR+"Job Polling failed but continuing with dSource creation. Error: "+job_err)
	}

	// Check if context was cancelled due to timeout
	if createCtx.Err() != nil {
		// Don't set ID in state - let user verify and import
		if createCtx.Err() == context.DeadlineExceeded {
			return diag.Errorf("dSource creation timed out after %s (Job ID: %s, dSource ID: %s). "+
				"Check DCT UI to verify job completion, then import it.",
				d.Timeout(schema.TimeoutCreate), apiRes.Job.GetId(), dsourceId)
		}
		return diag.Errorf("dSource creation was cancelled (Job ID: %s): %v", apiRes.Job.GetId(), createCtx.Err())
	}

	tflog.Info(ctx, DLPX+INFO+"Job result is "+job_res)

	rollback_on_failure := d.Get("rollback_on_failure").(bool)

	if job_res == Failed || job_res == Canceled || job_res == Abandoned {
		tflog.Error(ctx, DLPX+ERROR+"Job "+job_res+" "+apiRes.Job.GetId()+"!")
		if rollback_on_failure {
			if job_res == Failed {
				res := isSnapSyncFailure(apiRes.Job.GetId(), ctx, client)
				if res {
					deleteDiags := resourceDsourceDelete(ctx, d, meta)
					if deleteDiags.HasError() {
						return deleteDiags
					}
					d.SetId("")
				}
			}
		} else {
			readDiags := resourceDsourceRead(ctx, d, meta)

			if readDiags.HasError() {
				return readDiags
			}
		}
		return diag.Errorf("[NOT OK] Job %s %s with error %s", apiRes.Job.GetId(), job_res, job_err)
	}

	// Check context again before proceeding to snapshot polling
	if createCtx.Err() != nil {
		if createCtx.Err() == context.DeadlineExceeded {
			return diag.Errorf("dSource creation timed out after %s during snapshot polling (Job ID: %s). "+
				"The dSource may have been created. To resolve:\n"+
				"1. Check the Delphix DCT UI or API to verify the dSource exists\n"+
				"2. If created successfully, import it.",
				d.Timeout(schema.TimeoutCreate), apiRes.Job.GetId())
		}
		return diag.Errorf("dSource creation was cancelled after job completion (Job ID: %s): %v", apiRes.Job.GetId(), createCtx.Err())
	}

	PollSnapshotStatus(d, createCtx, client)

	// Check context one more time before reading state
	if createCtx.Err() != nil {
		if createCtx.Err() == context.DeadlineExceeded {
			return diag.Errorf("dSource creation timed out after %s during final state read (Job ID: %s). "+
				"The dSource may have been created. To resolve:\n"+
				"1. Check the Delphix DCT UI or API to verify the dSource exists\n"+
				"2. If created successfully, import it.",
				d.Timeout(schema.TimeoutCreate), apiRes.Job.GetId())
		}
		return diag.Errorf("dSource creation was cancelled during final state read (Job ID: %s): %v", apiRes.Job.GetId(), createCtx.Err())
	}
	
	// Only set ID in state after successful completion
	d.SetId(dsourceId)

	readDiags := resourceDsourceRead(createCtx, d, meta)

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

	ops_pre_sync_Raw, _ := d.Get("ops_pre_sync").([]interface{})
	oldOpsPreSync := toSourceOperationArray(ops_pre_sync_Raw)

	ops_post_sync_Raw, _ := d.Get("ops_post_sync").([]interface{})
	oldOpsPostSync := toSourceOperationArray(ops_post_sync_Raw)

	// _, rollback_on_failure_exists := d.GetOk("rollback_on_failure")
	// if !rollback_on_failure_exists {
	// 	// its an import or upgrade, set to default value
	// 	d.Set("rollback_on_failure", false)
	// }

	// Set make_current_account_owner to default true if not explicitly set
	if _, has_make_current := d.GetOk("make_current_account_owner"); !has_make_current {
		d.Set("make_current_account_owner", true)
	}

	// Set ignore_tag_changes to default true if not explicitly set
	if _, has_ignore_tags := d.GetOk("ignore_tag_changes"); !has_ignore_tags {
		d.Set("ignore_tag_changes", true)
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
	d.Set("sync_policy_id", result.GetSyncPolicyId())
	d.Set("retention_policy_id", result.GetRetentionPolicyId())
	d.Set("ops_pre_sync", flattenDSourceHooks(result.GetHooks().OpsPreSync, oldOpsPreSync))
	d.Set("ops_post_sync", flattenDSourceHooks(result.GetHooks().OpsPostSync, oldOpsPostSync))

	// get the tags and set it
	resTagsDsrc, httpRes, err := client.DSourcesAPI.GetTagsDsource(ctx, dsource_id).Execute()
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"Failed to fetch tags for dSource: "+dsource_id+". Error: "+err.Error())
	} else if httpRes != nil && httpRes.StatusCode >= 400 {
		tflog.Error(ctx, DLPX+ERROR+"Failed to fetch tags for dSource: "+dsource_id+". HTTP Status: "+httpRes.Status)
	} else {
		// check if tags are returned and set them to the state
		HandleRawConfigReadContext(ctx, d, resTagsDsrc)
	}
	return diags
}

func resourceDsourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var updateFailure bool = false
	var nonUpdatableField []string
	client := meta.(*apiClient).client

	// respect resource update timeout
	updateCtx, updateCancel := context.WithTimeout(ctx, d.Timeout(schema.TimeoutUpdate))
	defer updateCancel()
	updateAppdataDsource := dctapi.NewUpdateAppDataDSourceParameters()

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
		if strings.Contains(k, "ops_post_sync") {
			k = "ops_post_sync"
		}
		if d.HasChange(k) {
			tflog.Debug(ctx, "changed keys"+k)
			changedKeys = append(changedKeys, k)
		}
	}

	// check if the changed keys are updatable
	for _, key := range changedKeys {
		if !updatableAppdataDsourceKeys[key] {
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
		updateAppdataDsource.SetName(d.Get("name").(string))
	}
	if d.HasChange("description") {
		updateAppdataDsource.SetDescription(d.Get("description").(string))
	}
	if d.HasChange("staging_environment") {
		updateAppdataDsource.SetStagingEnvironment(d.Get("staging_environment").(string))
	}
	if d.HasChange("staging_environment_user") {
		updateAppdataDsource.SetStagingEnvironmentUser(d.Get("staging_environment_user").(string))
	}
	if d.HasChange("environment_user") {
		updateAppdataDsource.SetEnvironmentUser(d.Get("environment_user").(string))
	}
	if d.HasChange("parameters") {
		if v, has_v := d.GetOk("parameters"); has_v {
			params := make(map[string]interface{})
			json.Unmarshal([]byte(v.(string)), &params)
			updateAppdataDsource.SetParameters(params)
		}
	}
	if d.HasChange("sync_policy_id") {
		updateAppdataDsource.SetSyncPolicyId(d.Get("sync_policy_id").(string))
	}
	if d.HasChange("retention_policy_id") {
		updateAppdataDsource.SetRetentionPolicyId(d.Get("retention_policy_id").(string))
	}
	if d.HasChange("ops_pre_sync") {
		if v, has_v := d.GetOk("ops_pre_sync"); has_v {
			updateAppdataDsource.SetOpsPreSync(toSourceOperationArray(v))
		} else {
			updateAppdataDsource.SetOpsPreSync([]dctapi.SourceOperation{})
		}
	}
	if d.HasChange("ops_pre_sync") {
		if v, has_v := d.GetOk("ops_post_sync"); has_v {
			updateAppdataDsource.SetOpsPostSync(toSourceOperationArray(v))
		} else {
			updateAppdataDsource.SetOpsPostSync([]dctapi.SourceOperation{})
		}
	}
	// check if the updateAppdataDsource is not empty
	if !isStructEmpty(updateAppdataDsource) {
		tflog.Debug(ctx, "updating appdata dsource")
		res, httpRes, err := client.DSourcesAPI.UpdateAppdataDsourceById(updateCtx, dsourceId).UpdateAppDataDSourceParameters(*updateAppdataDsource).Execute()

		if diags := apiErrorResponseHelper(ctx, nil, httpRes, err); diags != nil {
			// revert and set the old value to the changed keys
			revertChanges(d, changedKeys)
			return diags
		}

		if res != nil {
			job_status, job_err := PollJobStatus(res.Job.GetId(), updateCtx, client)
			if job_err != "" {
				tflog.Warn(ctx, DLPX+WARN+"Appdata Dsource Update Job Polling failed but continuing with update. Error: "+job_err)
			}
			tflog.Info(ctx, DLPX+INFO+"Job result is "+job_status)
			if isJobTerminalFailure(job_status) {
				return diag.Errorf("[NOT OK] Appdata Dsource Update %s. JobId: %s / Error: %s", job_status, res.Job.GetId(), job_err)
			}
		}
	}

	// update tags
	if !d.Get("ignore_tag_changes").(bool) {
		apiRes, httpRes, err := client.DSourcesAPI.GetDsourceById(updateCtx, dsourceId).Execute()
		if diags := apiErrorResponseHelper(ctx, apiRes, httpRes, err); diags != nil {
			d.SetId("")
			return diags
		}
		tags := flattenTags(apiRes.GetTags())
		tflog.Debug(ctx, "Existing tags", map[string]interface{}{
			"tags": tags,
		})
		newRaw := d.GetRawConfig()
		if newRaw.IsKnown() || !newRaw.IsNull() {
			attr := newRaw.GetAttr("tags")
			tflog.Debug(ctx, "New tags raw config value", map[string]interface{}{
				"tags": newRaw,
			})
			d.Set("tags", flattenTags(apiRes.GetTags()))
			if attr.IsNull() || !attr.IsKnown() || attr.LengthInt() == 0 {
				// This now correctly gives [] if the user set tags = []
				if len(tags) != 0 {
					tflog.Info(ctx, DLPX+INFO+"Tags field is not set, deleting all existing tags")
					httpRes, err := client.DSourcesAPI.DeleteTagsDsource(ctx, dsourceId).Execute()
					if diags := apiErrorResponseHelper(ctx, nil, httpRes, err); diags != nil {
						return diags
					}
				}
				return resourceOracleDsourceRead(ctx, d, meta)
			}
		}
		oldTags, newTags := d.GetChange("tags")
		if !reflect.DeepEqual(oldTags, newTags) {
			tflog.Debug(ctx, "updating tags")
			// delete old tag
			tflog.Debug(ctx, "deleting old tags")
			if len(toTagArray(oldTags)) != 0 {
				tflog.Debug(ctx, "tag to be deleted: "+toTagArray(oldTags)[0].GetKey()+" "+toTagArray(oldTags)[0].GetValue())
				deleteTag := *dctapi.NewDeleteTag()
				tagDelResp, tagDelErr := client.DSourcesAPI.DeleteTagsDsource(updateCtx, dsourceId).DeleteTag(deleteTag).Execute()
				if diags := apiErrorResponseHelper(ctx, nil, tagDelResp, tagDelErr); diags != nil {
					revertChanges(d, changedKeys)
					updateFailure = true
				}
			}
			// create tag
			if len(toTagArray(newTags)) != 0 {
				tflog.Info(ctx, "creating new tags")
				_, httpResp, tagCrtErr := client.DSourcesAPI.CreateTagsDsource(updateCtx, dsourceId).TagsRequest(*dctapi.NewTagsRequest(toTagArray(newTags))).Execute()
				if diags := apiErrorResponseHelper(ctx, nil, httpResp, tagCrtErr); diags != nil {
					revertChanges(d, changedKeys)
					return diags
				}
			}
		}
	}

	return resourceOracleDsourceRead(updateCtx, d, meta)
}

func resourceDsourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient).client

	dsourceId := d.Id()

	deleteDsourceParams := dctapi.NewDeleteDSourceRequest(dsourceId)
	deleteDsourceParams.SetForce(false)

	// respect resource delete timeout
	deleteCtx, deleteCancel := context.WithTimeout(ctx, d.Timeout(schema.TimeoutDelete))
	defer deleteCancel()

	res, httpRes, err := client.DSourcesAPI.DeleteDsource(deleteCtx).DeleteDSourceRequest(*deleteDsourceParams).Execute()

	// Check if the API call itself timed out
	if err != nil && deleteCtx.Err() == context.DeadlineExceeded {
		return diag.Errorf("AppData dSource deletion API call timed out after %s. The request may still be processing on the DCT server. "+
			"Check the Delphix DCT UI or API to verify if a deletion job was created (dSource ID: %s). "+
			"If a job exists, wait for it to complete, then run 'terraform refresh' to verify the resource was deleted. "+
			"If the resource still exists in state, retry 'terraform destroy'. "+
			"To avoid timeouts, increase the timeout: timeouts { delete = \"60m\" }",
			d.Timeout(schema.TimeoutDelete), dsourceId)
	}

	if diags := apiErrorResponseHelper(ctx, res, httpRes, err); diags != nil {
		return diags
	}

	// Check for nil res or Job to prevent crashes
	if res == nil || res.Job == nil {
		return diag.Errorf("dSource deletion failed: received nil response or job from API")
	}

	// Check if context timed out before polling
	if deleteCtx.Err() == context.DeadlineExceeded {
		return diag.Errorf("AppData dSource deletion timed out after %s. The operation is still running on the DCT (Job ID: %s). "+
			"To resolve:\n"+
			"1. Wait for the job to complete (check Delphix DCT UI or API)\n"+
			"2. Run 'terraform refresh' to check if the resource was deleted\n"+
			"3. If still in state, retry 'terraform destroy'\n"+
			"To avoid timeouts, increase the timeout: timeouts { delete = \"60m\" }",
			d.Timeout(schema.TimeoutDelete), res.GetId())
	}

	job_status, job_err := PollJobStatus(res.GetId(), deleteCtx, client)
	if job_err != "" {
		// Check if the error is due to timeout
		if deleteCtx.Err() == context.DeadlineExceeded {
			return diag.Errorf("AppData dSource deletion timed out after %s while polling job status. The operation is still running on the DCT (Job ID: %s). "+
				"To resolve:\n"+
				"1. Wait for the job to complete (check Delphix DCT UI or API)\n"+
				"2. Run 'terraform refresh' to check if the resource was deleted\n"+
				"3. If still in state, retry 'terraform destroy'\n"+
				"To avoid timeouts, increase the timeout: timeouts { delete = \"60m\" }",
				d.Timeout(schema.TimeoutDelete), res.GetId())
		}
		tflog.Warn(ctx, DLPX+WARN+"Job Polling failed but continuing with deletion. Error :"+job_err)
	}
	tflog.Info(ctx, DLPX+INFO+"Job result is "+job_status)
	if isJobTerminalFailure(job_status) {
		return diag.Errorf("[NOT OK] dSource-Delete %s. JobId: %s / Error: %s", job_status, res.GetId(), job_err)
	}
	_, diags := PollForObjectDeletion(deleteCtx, func() (interface{}, *http.Response, error) {
		return client.DSourcesAPI.GetDsourceById(deleteCtx, dsourceId).Execute()
	})

	return diags
}
