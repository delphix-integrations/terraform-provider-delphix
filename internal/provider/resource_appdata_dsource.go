package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
				Required: true,
			},
			"group_id": {
				Type:     schema.TypeString,
				Required: true,
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
			"ignore_tag_changes": {
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if old != new {
						tflog.Info(context.Background(), "updating ignore_tag_changes is not allowed. plan changes are suppressed")
					}
					return d.Id() != ""
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
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					ignore_tag_changes, _ := d.GetOk("ignore_tag_changes")
					if ignore_tag_changes.(bool) {
						return true
					} else {
						tflog.Debug(context.Background(), fmt.Sprintf("\n [DEBUG] tag changes suppressed : %v", ignore_tag_changes))
						return false
					}
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
	}
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

	req := client.DSourcesAPI.LinkAppdataDatabase(ctx)

	apiRes, httpRes, err := req.AppDataDSourceLinkSourceParameters(*appDataDSourceLinkSourceParameters).Execute()
	if diags := apiErrorResponseHelper(ctx, apiRes, httpRes, err); diags != nil {
		return diags
	}

	d.SetId(apiRes.GetDsourceId())

	job_res, job_err := PollJobStatus(apiRes.Job.GetId(), ctx, client)
	if job_err != "" {
		tflog.Error(ctx, DLPX+ERROR+"Job Polling failed but continuing with dSource creation. Error: "+job_err)
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

	ops_pre_sync_Raw, _ := d.Get("ops_pre_sync").([]interface{})
	oldOpsPreSync := toSourceOperationArray(ops_pre_sync_Raw)

	ops_post_sync_Raw, _ := d.Get("ops_post_sync").([]interface{})
	oldOpsPostSync := toSourceOperationArray(ops_post_sync_Raw)

	_, rollback_on_failure_exists := d.GetOk("rollback_on_failure")
	if !rollback_on_failure_exists {
		// its an import or upgrade, set to default value
		d.Set("rollback_on_failure", false)
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
	d.Set("retention_policy_id", result.GetReplicaRetentionPolicyId())
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
		if len(resTagsDsrc.GetTags()) != 0 {
			tflog.Debug(ctx, DLPX+"Tags are present")
			d.Set("tags", flattenTags(resTagsDsrc.GetTags()))
		}
	}
	return diags
}

func resourceDsourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	var updateFailure bool = false
	var nonUpdatableField []string
	client := meta.(*apiClient).client
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

	res, httpRes, err := client.DSourcesAPI.UpdateAppdataDsourceById(ctx, dsourceId).UpdateAppDataDSourceParameters(*updateAppdataDsource).Execute()

	if diags := apiErrorResponseHelper(ctx, nil, httpRes, err); diags != nil {
		// revert and set the old value to the changed keys
		revertChanges(d, changedKeys)
		return diags
	}

	job_status, job_err := PollJobStatus(res.Job.GetId(), ctx, client)
	if job_err != "" {
		tflog.Warn(ctx, DLPX+WARN+"Appdata Dsource Update Job Polling failed but continuing with update. Error: "+job_err)
	}
	tflog.Info(ctx, DLPX+INFO+"Job result is "+job_status)
	if isJobTerminalFailure(job_status) {
		return diag.Errorf("[NOT OK] Appdata Dsource Update %s. JobId: %s / Error: %s", job_status, res.Job.GetId(), job_err)
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

func resourceDsourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient).client

	dsourceId := d.Id()

	deleteDsourceParams := dctapi.NewDeleteDSourceRequest(dsourceId)
	deleteDsourceParams.SetForce(false)

	res, httpRes, err := client.DSourcesAPI.DeleteDsource(ctx).DeleteDSourceRequest(*deleteDsourceParams).Execute()

	if diags := apiErrorResponseHelper(ctx, res, httpRes, err); diags != nil {
		return diags
	}

	job_status, job_err := PollJobStatus(res.GetId(), ctx, client)
	if job_err != "" {
		tflog.Warn(ctx, DLPX+WARN+"Job Polling failed but continuing with deletion. Error :"+job_err)
	}
	tflog.Info(ctx, DLPX+INFO+"Job result is "+job_status)
	if isJobTerminalFailure(job_status) {
		return diag.Errorf("[NOT OK] dSource-Delete %s. JobId: %s / Error: %s", job_status, res.GetId(), job_err)
	}

	_, diags := PollForObjectDeletion(ctx, func() (interface{}, *http.Response, error) {
		return client.DSourcesAPI.GetDsourceById(ctx, dsourceId).Execute()
	})

	return diags
}
