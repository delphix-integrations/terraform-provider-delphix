package provider

import (
	"context"
	"net/http"
	"time"

	dctapi "github.com/delphix/dct-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVdb() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Resource for provisioning VDB.",

		CreateContext: resourceVdbCreate,
		ReadContext:   resourceVdbRead,
		UpdateContext: resourceVdbUpdate,
		DeleteContext: resourceVdbDelete,

		Schema: map[string]*schema.Schema{
			"provision_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "snapshot",
			},
			"auto_select_repository": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"source_data_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"engine_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"environment_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"group_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creation_date": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"target_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"database_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cdb_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster_node_ids": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"truncate_log_on_checkpoint": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"os_username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"environment_user_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"repository_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pre_refresh": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"shell": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"element_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"post_refresh": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"shell": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"element_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"pre_rollback": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"shell": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"element_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"post_rollback": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"shell": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"element_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"configure_clone": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"shell": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"element_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"pre_snapshot": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"shell": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"element_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"post_snapshot": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"shell": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"element_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"pre_start": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"shell": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"element_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"post_start": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"shell": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"element_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"pre_stop": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"shell": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"element_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"post_stop": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"shell": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"element_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"vdb_restart": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"template_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"auxiliary_template_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"file_mapping_rules": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oracle_instance_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"unique_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcdb_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcdb_database_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mount_point": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"open_reset_logs": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"snapshot_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"retention_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"recovery_model": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pre_script": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"post_script": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cdc_on_provision": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"online_log_size": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"online_log_groups": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"archive_log": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"new_dbid": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"listener_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"custom_env_vars": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"custom_env_files": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"timestamp": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"timestamp_in_database_timezone": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"snapshot_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bookmark_id": {
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
		},
	}
}

func toHookArray(array interface{}) []dctapi.Hook {
	items := []dctapi.Hook{}
	for _, item := range array.([]interface{}) {
		item_map := item.(map[string]interface{})
		hook_item := dctapi.NewHook(item_map["command"].(string))

		name := item_map["name"].(string)
		if name != "" {
			hook_item.SetName(item_map["name"].(string))
		}

		// defaults to "bash" as per resource schema spec
		hook_item.SetShell(item_map["shell"].(string))
		items = append(items, *hook_item)
	}
	return items
}

func toTagArray(array interface{}) []dctapi.Tag {
	items := []dctapi.Tag{}
	for _, item := range array.([]interface{}) {
		item_map := item.(map[string]interface{})
		tag_item := dctapi.NewTag(item_map["key"].(string), item_map["value"].(string))

		items = append(items, *tag_item)
	}
	return items
}

func helper_provision_by_snapshot(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*apiClient).client

	provisionVDBBySnapshotParameters := dctapi.NewProvisionVDBBySnapshotParameters()

	// Setters for provisionVDBBySnapshotParameters
	if v, has_v := d.GetOkExists("auto_select_repository"); has_v {
		provisionVDBBySnapshotParameters.SetAutoSelectRepository(v.(bool))
	}
	if v, has_v := d.GetOk("source_data_id"); has_v {
		provisionVDBBySnapshotParameters.SetSourceDataId(v.(string))
	}
	if v, has_v := d.GetOk("engine_id"); has_v {
		provisionVDBBySnapshotParameters.SetEngineId(v.(string))
	}
	if v, has_v := d.GetOk("target_group_id"); has_v {
		provisionVDBBySnapshotParameters.SetTargetGroupId(v.(string))
	}
	if v, has_v := d.GetOk("name"); has_v {
		provisionVDBBySnapshotParameters.SetName(v.(string))
	}
	if v, has_v := d.GetOk("database_name"); has_v {
		provisionVDBBySnapshotParameters.SetDatabaseName(v.(string))
	}
	if v, has_v := d.GetOk("cdb_id"); has_v {
		provisionVDBBySnapshotParameters.SetCdbId(v.(string))
	}
	if v, has_v := d.GetOk("cluster_node_ids"); has_v {
		provisionVDBBySnapshotParameters.SetClusterNodeIds(toStringArray(v))
	}
	if v, has_v := d.GetOkExists("truncate_log_on_checkpoint"); has_v {
		provisionVDBBySnapshotParameters.SetTruncateLogOnCheckpoint(v.(bool))
	}
	if v, has_v := d.GetOk("os_username"); has_v {
		provisionVDBBySnapshotParameters.SetOsUsername(v.(string))
	}
	if v, has_v := d.GetOk("os_password"); has_v {
		provisionVDBBySnapshotParameters.SetOsPassword(v.(string))
	}
	if v, has_v := d.GetOk("environment_id"); has_v {
		provisionVDBBySnapshotParameters.SetEnvironmentId(v.(string))
	}
	if v, has_v := d.GetOk("environment_user_id"); has_v {
		provisionVDBBySnapshotParameters.SetEnvironmentUserId(v.(string))
	}
	if v, has_v := d.GetOk("repository_id"); has_v {
		provisionVDBBySnapshotParameters.SetRepositoryId(v.(string))
	}
	if v, has_v := d.GetOkExists("auto_select_repository"); has_v {
		provisionVDBBySnapshotParameters.SetAutoSelectRepository(v.(bool))
	}
	if v, has_v := d.GetOkExists("vdb_restart"); has_v {
		provisionVDBBySnapshotParameters.SetVdbRestart(v.(bool))
	}
	if v, has_v := d.GetOk("template_id"); has_v {
		provisionVDBBySnapshotParameters.SetTemplateId(v.(string))
	}
	if v, has_v := d.GetOk("auxiliary_template_id"); has_v {
		provisionVDBBySnapshotParameters.SetAuxiliaryTemplateId(v.(string))
	}
	if v, has_v := d.GetOk("file_mapping_rules"); has_v {
		provisionVDBBySnapshotParameters.SetFileMappingRules(v.(string))
	}
	if v, has_v := d.GetOk("oracle_instance_name"); has_v {
		provisionVDBBySnapshotParameters.SetOracleInstanceName(v.(string))
	}
	if v, has_v := d.GetOk("unique_name"); has_v {
		provisionVDBBySnapshotParameters.SetUniqueName(v.(string))
	}
	if v, has_v := d.GetOk("vcdb_name"); has_v {
		provisionVDBBySnapshotParameters.SetVcdbName(v.(string))
	}
	if v, has_v := d.GetOk("vcdb_database_name"); has_v {
		provisionVDBBySnapshotParameters.SetVcdbDatabaseName(v.(string))
	}
	if v, has_v := d.GetOk("mount_point"); has_v {
		provisionVDBBySnapshotParameters.SetMountPoint(v.(string))
	}
	if v, has_v := d.GetOkExists("open_reset_logs"); has_v {
		provisionVDBBySnapshotParameters.SetOpenResetLogs(v.(bool))
	}
	if v, has_v := d.GetOk("snapshot_policy_id"); has_v {
		provisionVDBBySnapshotParameters.SetSnapshotPolicyId(v.(string))
	}
	if v, has_v := d.GetOk("retention_policy_id"); has_v {
		provisionVDBBySnapshotParameters.SetRetentionPolicyId(v.(string))
	}
	if v, has_v := d.GetOk("recovery_model"); has_v {
		provisionVDBBySnapshotParameters.SetRecoveryModel(v.(string))
	}
	if v, has_v := d.GetOk("pre_script"); has_v {
		provisionVDBBySnapshotParameters.SetPreScript(v.(string))
	}
	if v, has_v := d.GetOk("post_script"); has_v {
		provisionVDBBySnapshotParameters.SetPostScript(v.(string))
	}
	if v, has_v := d.GetOkExists("cdc_on_provision"); has_v {
		provisionVDBBySnapshotParameters.SetCdcOnProvision(v.(bool))
	}
	if v, has_v := d.GetOk("online_log_size"); has_v {
		provisionVDBBySnapshotParameters.SetOnlineLogSize(int32(v.(int)))
	}
	if v, has_v := d.GetOk("online_log_groups"); has_v {
		provisionVDBBySnapshotParameters.SetOnlineLogGroups(int32(v.(int)))
	}
	if v, has_v := d.GetOkExists("archive_log"); has_v {
		provisionVDBBySnapshotParameters.SetArchiveLog(v.(bool))
	}
	if v, has_v := d.GetOkExists("new_dbid"); has_v {
		provisionVDBBySnapshotParameters.SetNewDbid(v.(bool))
	}
	if v, has_v := d.GetOkExists("listener_ids"); has_v {
		provisionVDBBySnapshotParameters.SetListenerIds(toStringArray(v))
	}
	if v, has_v := d.GetOk("snapshot_id"); has_v {
		provisionVDBBySnapshotParameters.SetSnapshotId(v.(string))
	}
	if v, has_v := d.GetOk("custom_env_files"); has_v {
		provisionVDBBySnapshotParameters.SetCustomEnvFiles(toStringArray(v))
	}
	if v, has_v := d.GetOk("custom_env_vars"); has_v {
		custom_env_vars := make(map[string]string)

		for k, v := range v.(map[string]interface{}) {
			custom_env_vars[k] = v.(string)
		}
		provisionVDBBySnapshotParameters.SetCustomEnvVars(custom_env_vars)
	}

	if v, has_v := d.GetOk("pre_refresh"); has_v {
		provisionVDBBySnapshotParameters.SetPreRefresh(toHookArray(v))
	}
	if v, has_v := d.GetOk("post_refresh"); has_v {
		provisionVDBBySnapshotParameters.SetPostRefresh(toHookArray(v))
	}
	if v, has_v := d.GetOk("pre_rollback"); has_v {
		provisionVDBBySnapshotParameters.SetPreRollback(toHookArray(v))
	}
	if v, has_v := d.GetOk("post_rollback"); has_v {
		provisionVDBBySnapshotParameters.SetPostRollback(toHookArray(v))
	}
	if v, has_v := d.GetOk("configure_clone"); has_v {
		provisionVDBBySnapshotParameters.SetConfigureClone(toHookArray(v))
	}
	if v, has_v := d.GetOk("pre_snapshot"); has_v {
		provisionVDBBySnapshotParameters.SetPreSnapshot(toHookArray(v))
	}
	if v, has_v := d.GetOk("post_snapshot"); has_v {
		provisionVDBBySnapshotParameters.SetPostSnapshot(toHookArray(v))
	}
	if v, has_v := d.GetOk("pre_start"); has_v {
		provisionVDBBySnapshotParameters.SetPreStart(toHookArray(v))
	}
	if v, has_v := d.GetOk("post_start"); has_v {
		provisionVDBBySnapshotParameters.SetPostStart(toHookArray(v))
	}
	if v, has_v := d.GetOk("pre_stop"); has_v {
		provisionVDBBySnapshotParameters.SetPreStop(toHookArray(v))
	}
	if v, has_v := d.GetOk("post_stop"); has_v {
		provisionVDBBySnapshotParameters.SetPostStop(toHookArray(v))
	}
	if v, has_v := d.GetOk("tags"); has_v {
		provisionVDBBySnapshotParameters.SetTags(toTagArray(v))
	}

	req := client.VDBsApi.ProvisionVdbBySnapshot(ctx)

	apiRes, httpRes, err := req.ProvisionVDBBySnapshotParameters(*provisionVDBBySnapshotParameters).Execute()
	if diags := apiErrorResponseHelper(apiRes, httpRes, err); diags != nil {
		return diags
	}

	d.SetId(*apiRes.VdbId)

	job_res, job_err := PollJobStatus(*apiRes.Job.Id, ctx, client)
	if job_err != "" {
		ErrorLog.Printf("Job Polling failed but continuing with provisioning. Error: %s", job_err)
	}
	InfoLog.Printf("Job result is %s", job_res)
	if job_res == Failed || job_res == Canceled || job_res == Abandoned {
		ErrorLog.Printf("Job %s %s!", job_res, *apiRes.Job.Id)
		return diag.Errorf("[NOT OK] Job %s %s with error %s", *apiRes.Job.Id, job_res, job_err)
	}

	readDiags := resourceVdbRead(ctx, d, meta)

	if readDiags.HasError() {
		return readDiags
	}

	return diags
}

func helper_provision_by_timestamp(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*apiClient).client

	provisionVDBByTimestampParameters := dctapi.NewProvisionVDBByTimestampParameters(d.Get("source_data_id").(string))

	// Setters for provisionVDBByTimestampParameters
	if v, has_v := d.GetOk("engine_id"); has_v {
		// provisionVDBByTimestampParameters.SetEngineId(int64(v.(int)))
		provisionVDBByTimestampParameters.SetEngineId(v.(string))
	}
	if v, has_v := d.GetOk("target_group_id"); has_v {
		provisionVDBByTimestampParameters.SetTargetGroupId(v.(string))
	}
	if v, has_v := d.GetOk("name"); has_v {
		provisionVDBByTimestampParameters.SetName(v.(string))
	}
	if v, has_v := d.GetOk("database_name"); has_v {
		provisionVDBByTimestampParameters.SetDatabaseName(v.(string))
	}
	if v, has_v := d.GetOk("cdb_id"); has_v {
		provisionVDBByTimestampParameters.SetCdbId(v.(string))
	}
	if v, has_v := d.GetOk("cluster_node_ids"); has_v {
		provisionVDBByTimestampParameters.SetClusterNodeIds(toStringArray(v))
	}
	if v, has_v := d.GetOkExists("truncate_log_on_checkpoint"); has_v {
		provisionVDBByTimestampParameters.SetTruncateLogOnCheckpoint(v.(bool))
	}
	if v, has_v := d.GetOk("os_username"); has_v {
		provisionVDBByTimestampParameters.SetOsUsername(v.(string))
	}
	if v, has_v := d.GetOk("os_password"); has_v {
		provisionVDBByTimestampParameters.SetOsPassword(v.(string))
	}
	if v, has_v := d.GetOk("environment_id"); has_v {
		provisionVDBByTimestampParameters.SetEnvironmentId(v.(string))
	}
	if v, has_v := d.GetOk("environment_user_id"); has_v {
		provisionVDBByTimestampParameters.SetEnvironmentUserId(v.(string))
	}
	if v, has_v := d.GetOk("repository_id"); has_v {
		provisionVDBByTimestampParameters.SetRepositoryId(v.(string))
	}
	if v, has_v := d.GetOkExists("auto_select_repository"); has_v {
		provisionVDBByTimestampParameters.SetAutoSelectRepository(v.(bool))
	}
	if v, has_v := d.GetOkExists("vdb_restart"); has_v {
		provisionVDBByTimestampParameters.SetVdbRestart(v.(bool))
	}
	if v, has_v := d.GetOk("template_id"); has_v {
		provisionVDBByTimestampParameters.SetTemplateId(v.(string))
	}
	if v, has_v := d.GetOk("auxiliary_template_id"); has_v {
		provisionVDBByTimestampParameters.SetAuxiliaryTemplateId(v.(string))
	}
	if v, has_v := d.GetOk("file_mapping_rules"); has_v {
		provisionVDBByTimestampParameters.SetFileMappingRules(v.(string))
	}
	if v, has_v := d.GetOk("oracle_instance_name"); has_v {
		provisionVDBByTimestampParameters.SetOracleInstanceName(v.(string))
	}
	if v, has_v := d.GetOk("unique_name"); has_v {
		provisionVDBByTimestampParameters.SetUniqueName(v.(string))
	}
	if v, has_v := d.GetOk("vcdb_name"); has_v {
		provisionVDBByTimestampParameters.SetVcdbName(v.(string))
	}
	if v, has_v := d.GetOk("vcdb_database_name"); has_v {
		provisionVDBByTimestampParameters.SetVcdbDatabaseName(v.(string))
	}
	if v, has_v := d.GetOk("mount_point"); has_v {
		provisionVDBByTimestampParameters.SetMountPoint(v.(string))
	}
	if v, has_v := d.GetOkExists("open_reset_logs"); has_v {
		provisionVDBByTimestampParameters.SetOpenResetLogs(v.(bool))
	}
	if v, has_v := d.GetOk("snapshot_policy_id"); has_v {
		provisionVDBByTimestampParameters.SetSnapshotPolicyId(v.(string))
	}
	if v, has_v := d.GetOk("retention_policy_id"); has_v {
		provisionVDBByTimestampParameters.SetRetentionPolicyId(v.(string))
	}
	if v, has_v := d.GetOk("recovery_model"); has_v {
		provisionVDBByTimestampParameters.SetRecoveryModel(v.(string))
	}
	if v, has_v := d.GetOk("pre_script"); has_v {
		provisionVDBByTimestampParameters.SetPreScript(v.(string))
	}
	if v, has_v := d.GetOk("post_script"); has_v {
		provisionVDBByTimestampParameters.SetPostScript(v.(string))
	}
	if v, has_v := d.GetOkExists("cdc_on_provision"); has_v {
		provisionVDBByTimestampParameters.SetCdcOnProvision(v.(bool))
	}
	if v, has_v := d.GetOk("online_log_size"); has_v {
		provisionVDBByTimestampParameters.SetOnlineLogSize(int32(v.(int)))
	}
	if v, has_v := d.GetOk("online_log_groups"); has_v {
		provisionVDBByTimestampParameters.SetOnlineLogGroups(int32(v.(int)))
	}
	if v, has_v := d.GetOkExists("archive_log"); has_v {
		provisionVDBByTimestampParameters.SetArchiveLog(v.(bool))
	}
	if v, has_v := d.GetOkExists("new_dbid"); has_v {
		provisionVDBByTimestampParameters.SetNewDbid(v.(bool))
	}
	if v, has_v := d.GetOk("listener_ids"); has_v {
		provisionVDBByTimestampParameters.SetListenerIds(toStringArray(v))
	}
	if v, has_v := d.GetOk("custom_env_vars"); has_v {
		custom_env_vars := make(map[string]string)

		for k, v := range v.(map[string]interface{}) {
			custom_env_vars[k] = v.(string)
		}
		provisionVDBByTimestampParameters.SetCustomEnvVars(custom_env_vars)
	}
	if v, has_v := d.GetOk("custom_env_files"); has_v {
		provisionVDBByTimestampParameters.SetCustomEnvFiles(toStringArray(v))
	}
	if v, has_v := d.GetOk("timestamp"); has_v {
		tt, err := time.Parse(time.RFC3339, v.(string))
		if err != nil {
			ErrorLog.Printf("An error has occured: %v", err)
			return diag.Errorf("The timestamp parameter %s is not valid RFC3339 format. Please provide valid value. Example: 2021-05-01T08:51:34.148000+00:00", v.(string))
		}
		provisionVDBByTimestampParameters.SetTimestamp(tt)
	}
	if v, has_v := d.GetOk("timestamp_in_database_timezone"); has_v {
		provisionVDBByTimestampParameters.SetTimestampInDatabaseTimezone(v.(string))
	}
	if v, has_v := d.GetOk("pre_refresh"); has_v {
		provisionVDBByTimestampParameters.SetPreRefresh(toHookArray(v))
	}
	if v, has_v := d.GetOk("post_refresh"); has_v {
		provisionVDBByTimestampParameters.SetPostRefresh(toHookArray(v))
	}
	if v, has_v := d.GetOk("pre_rollback"); has_v {
		provisionVDBByTimestampParameters.SetPreRollback(toHookArray(v))
	}
	if v, has_v := d.GetOk("post_rollback"); has_v {
		provisionVDBByTimestampParameters.SetPostRollback(toHookArray(v))
	}
	if v, has_v := d.GetOk("configure_clone"); has_v {
		provisionVDBByTimestampParameters.SetConfigureClone(toHookArray(v))
	}
	if v, has_v := d.GetOk("pre_snapshot"); has_v {
		provisionVDBByTimestampParameters.SetPreSnapshot(toHookArray(v))
	}
	if v, has_v := d.GetOk("post_snapshot"); has_v {
		provisionVDBByTimestampParameters.SetPostSnapshot(toHookArray(v))
	}
	if v, has_v := d.GetOk("pre_start"); has_v {
		provisionVDBByTimestampParameters.SetPreStart(toHookArray(v))
	}
	if v, has_v := d.GetOk("post_start"); has_v {
		provisionVDBByTimestampParameters.SetPostStart(toHookArray(v))
	}
	if v, has_v := d.GetOk("pre_stop"); has_v {
		provisionVDBByTimestampParameters.SetPreStop(toHookArray(v))
	}
	if v, has_v := d.GetOk("post_stop"); has_v {
		provisionVDBByTimestampParameters.SetPostStop(toHookArray(v))
	}
	if v, has_v := d.GetOk("tags"); has_v {
		provisionVDBByTimestampParameters.SetTags(toTagArray(v))
	}

	req := client.VDBsApi.ProvisionVdbByTimestamp(ctx)

	apiRes, httpRes, err := req.ProvisionVDBByTimestampParameters(*provisionVDBByTimestampParameters).Execute()
	if diags := apiErrorResponseHelper(apiRes, httpRes, err); diags != nil {
		return diags
	}

	d.SetId(*apiRes.VdbId)

	job_res, job_err := PollJobStatus(*apiRes.Job.Id, ctx, client)
	if job_err != "" {
		ErrorLog.Printf("Job Polling failed but continuing with provisioning. Error: %v", job_err)
	}
	InfoLog.Printf("Job result is %s", job_res)
	if job_res == "FAILED" {
		ErrorLog.Printf("Job %s Failed!", *apiRes.Job.Id)
		return diag.Errorf("[NOT OK] Job %s Failed with error %s", *apiRes.Job.Id, job_err)
	}

	readDiags := resourceVdbRead(ctx, d, meta)

	if readDiags.HasError() {
		return readDiags
	}

	return diags
}

func helper_provision_by_bookmark(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*apiClient).client

	provisionVDBFromBookmarkParameters := dctapi.NewProvisionVDBFromBookmarkParameters(d.Get("bookmark_id").(string))

	// Setters for provisionVDBFromBookmarkParameters
	if v, has_v := d.GetOk("target_group_id"); has_v {
		provisionVDBFromBookmarkParameters.SetTargetGroupId(v.(string))
	}
	if v, has_v := d.GetOk("name"); has_v {
		provisionVDBFromBookmarkParameters.SetName(v.(string))
	}
	if v, has_v := d.GetOk("database_name"); has_v {
		provisionVDBFromBookmarkParameters.SetDatabaseName(v.(string))
	}
	if v, has_v := d.GetOk("cdb_id"); has_v {
		provisionVDBFromBookmarkParameters.SetCdbId(v.(string))
	}
	if v, has_v := d.GetOk("cluster_node_ids"); has_v {
		provisionVDBFromBookmarkParameters.SetClusterNodeIds(toStringArray(v))
	}
	if v, has_v := d.GetOkExists("truncate_log_on_checkpoint"); has_v {
		provisionVDBFromBookmarkParameters.SetTruncateLogOnCheckpoint(v.(bool))
	}
	if v, has_v := d.GetOk("os_username"); has_v {
		provisionVDBFromBookmarkParameters.SetOsUsername(v.(string))
	}
	if v, has_v := d.GetOk("os_password"); has_v {
		provisionVDBFromBookmarkParameters.SetOsPassword(v.(string))
	}
	if v, has_v := d.GetOk("environment_id"); has_v {
		provisionVDBFromBookmarkParameters.SetEnvironmentId(v.(string))
	}
	if v, has_v := d.GetOk("environment_user_id"); has_v {
		provisionVDBFromBookmarkParameters.SetEnvironmentUserId(v.(string))
	}
	if v, has_v := d.GetOk("repository_id"); has_v {
		provisionVDBFromBookmarkParameters.SetRepositoryId(v.(string))
	}
	if v, has_v := d.GetOkExists("auto_select_repository"); has_v {
		provisionVDBFromBookmarkParameters.SetAutoSelectRepository(v.(bool))
	}
	if v, has_v := d.GetOkExists("vdb_restart"); has_v {
		provisionVDBFromBookmarkParameters.SetVdbRestart(v.(bool))
	}
	if v, has_v := d.GetOk("template_id"); has_v {
		provisionVDBFromBookmarkParameters.SetTemplateId(v.(string))
	}
	if v, has_v := d.GetOk("auxiliary_template_id"); has_v {
		provisionVDBFromBookmarkParameters.SetAuxiliaryTemplateId(v.(string))
	}
	if v, has_v := d.GetOk("file_mapping_rules"); has_v {
		provisionVDBFromBookmarkParameters.SetFileMappingRules(v.(string))
	}
	if v, has_v := d.GetOk("oracle_instance_name"); has_v {
		provisionVDBFromBookmarkParameters.SetOracleInstanceName(v.(string))
	}
	if v, has_v := d.GetOk("unique_name"); has_v {
		provisionVDBFromBookmarkParameters.SetUniqueName(v.(string))
	}
	if v, has_v := d.GetOk("vcdb_name"); has_v {
		provisionVDBFromBookmarkParameters.SetVcdbName(v.(string))
	}
	if v, has_v := d.GetOk("vcdb_database_name"); has_v {
		provisionVDBFromBookmarkParameters.SetVcdbDatabaseName(v.(string))
	}
	if v, has_v := d.GetOk("mount_point"); has_v {
		provisionVDBFromBookmarkParameters.SetMountPoint(v.(string))
	}
	if v, has_v := d.GetOkExists("open_reset_logs"); has_v {
		provisionVDBFromBookmarkParameters.SetOpenResetLogs(v.(bool))
	}
	if v, has_v := d.GetOk("snapshot_policy_id"); has_v {
		provisionVDBFromBookmarkParameters.SetSnapshotPolicyId(v.(string))
	}
	if v, has_v := d.GetOk("retention_policy_id"); has_v {
		provisionVDBFromBookmarkParameters.SetRetentionPolicyId(v.(string))
	}
	if v, has_v := d.GetOk("recovery_model"); has_v {
		provisionVDBFromBookmarkParameters.SetRecoveryModel(v.(string))
	}
	if v, has_v := d.GetOk("pre_script"); has_v {
		provisionVDBFromBookmarkParameters.SetPreScript(v.(string))
	}
	if v, has_v := d.GetOk("post_script"); has_v {
		provisionVDBFromBookmarkParameters.SetPostScript(v.(string))
	}
	if v, has_v := d.GetOkExists("cdc_on_provision"); has_v {
		provisionVDBFromBookmarkParameters.SetCdcOnProvision(v.(bool))
	}
	if v, has_v := d.GetOk("online_log_size"); has_v {
		provisionVDBFromBookmarkParameters.SetOnlineLogSize(int32(v.(int)))
	}
	if v, has_v := d.GetOk("online_log_groups"); has_v {
		provisionVDBFromBookmarkParameters.SetOnlineLogGroups(int32(v.(int)))
	}
	if v, has_v := d.GetOkExists("archive_log"); has_v {
		provisionVDBFromBookmarkParameters.SetArchiveLog(v.(bool))
	}
	if v, has_v := d.GetOkExists("new_dbid"); has_v {
		provisionVDBFromBookmarkParameters.SetNewDbid(v.(bool))
	}
	if v, has_v := d.GetOk("listener_ids"); has_v {
		provisionVDBFromBookmarkParameters.SetListenerIds(toStringArray(v))
	}
	if v, has_v := d.GetOk("custom_env_vars"); has_v {
		custom_env_vars := make(map[string]string)

		for k, v := range v.(map[string]interface{}) {
			custom_env_vars[k] = v.(string)
		}
		provisionVDBFromBookmarkParameters.SetCustomEnvVars(custom_env_vars)
	}
	if v, has_v := d.GetOk("custom_env_files"); has_v {
		provisionVDBFromBookmarkParameters.SetCustomEnvFiles(toStringArray(v))
	}
	if v, has_v := d.GetOk("pre_refresh"); has_v {
		provisionVDBFromBookmarkParameters.SetPreRefresh(toHookArray(v))
	}
	if v, has_v := d.GetOk("post_refresh"); has_v {
		provisionVDBFromBookmarkParameters.SetPostRefresh(toHookArray(v))
	}
	if v, has_v := d.GetOk("pre_rollback"); has_v {
		provisionVDBFromBookmarkParameters.SetPreRollback(toHookArray(v))
	}
	if v, has_v := d.GetOk("post_rollback"); has_v {
		provisionVDBFromBookmarkParameters.SetPostRollback(toHookArray(v))
	}
	if v, has_v := d.GetOk("configure_clone"); has_v {
		provisionVDBFromBookmarkParameters.SetConfigureClone(toHookArray(v))
	}
	if v, has_v := d.GetOk("pre_snapshot"); has_v {
		provisionVDBFromBookmarkParameters.SetPreSnapshot(toHookArray(v))
	}
	if v, has_v := d.GetOk("post_snapshot"); has_v {
		provisionVDBFromBookmarkParameters.SetPostSnapshot(toHookArray(v))
	}
	if v, has_v := d.GetOk("pre_start"); has_v {
		provisionVDBFromBookmarkParameters.SetPreStart(toHookArray(v))
	}
	if v, has_v := d.GetOk("post_start"); has_v {
		provisionVDBFromBookmarkParameters.SetPostStart(toHookArray(v))
	}
	if v, has_v := d.GetOk("pre_stop"); has_v {
		provisionVDBFromBookmarkParameters.SetPreStop(toHookArray(v))
	}
	if v, has_v := d.GetOk("post_stop"); has_v {
		provisionVDBFromBookmarkParameters.SetPostStop(toHookArray(v))
	}
	if v, has_v := d.GetOk("tags"); has_v {
		provisionVDBFromBookmarkParameters.SetPostStop(toHookArray(v))
	}

	req := client.VDBsApi.ProvisionVdbFromBookmark(ctx)

	apiRes, httpRes, err := req.ProvisionVDBFromBookmarkParameters(*provisionVDBFromBookmarkParameters).Execute()
	if diags := apiErrorResponseHelper(apiRes, httpRes, err); diags != nil {
		return diags
	}

	d.SetId(*apiRes.VdbId)

	job_res, job_err := PollJobStatus(*apiRes.Job.Id, ctx, client)
	if job_err != "" {
		ErrorLog.Printf("Job Polling failed but continuing with provisioning. Error: %s", job_err)
	}
	InfoLog.Printf("Job result is %s", job_res)
	if job_res == Failed || job_res == Canceled || job_res == Abandoned {
		ErrorLog.Printf("Job %s %s!", job_res, *apiRes.Job.Id)
		return diag.Errorf("[NOT OK] Job %s %s with error %s", *apiRes.Job.Id, job_res, job_err)
	}

	readDiags := resourceVdbRead(ctx, d, meta)

	if readDiags.HasError() {
		return readDiags
	}

	return diags
}

func resourceVdbCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if _, has_v := d.GetOk("db_username"); has_v {
		return diag.Errorf("db_username can not be set when creating a VDB.")
	}
	if _, has_v := d.GetOk("db_password"); has_v {
		return diag.Errorf("db_password can not be set when creating a VDB.")
	}

	provision_type := d.Get("provision_type").(string)

	if provision_type == "timestamp" {
		if _, has_v := d.GetOk("snapshot_id"); has_v {
			return diag.Errorf("snapshot_id is not supported for provision_type = 'timestamp'")
		} else {
			return helper_provision_by_timestamp(ctx, d, meta)
		}
	} else if provision_type == "snapshot" {
		if _, has_v := d.GetOk("timestamp"); has_v {
			return diag.Errorf("timestamp is not supported for provision_type = 'snapshot'")
		} else {
			return helper_provision_by_snapshot(ctx, d, meta)
		}
	} else if provision_type == "bookmark" {
		if _, has_v := d.GetOk("bookmark_id"); has_v {
			return helper_provision_by_bookmark(ctx, d, meta)
		} else {
			return diag.Errorf("bookmark_id is required for provision_type = 'bookmark'")
		}
	} else {
		return diag.Errorf("provision_type must be 'timestamp', 'snapshot' or 'bookmark'")
	}
}

func resourceVdbRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*apiClient).client

	vdbId := d.Id()

	res, diags := PollForObjectExistence(func() (interface{}, *http.Response, error) {
		return client.VDBsApi.GetVdbById(ctx, vdbId).Execute()
	})

	if diags != nil {
		_, diags := PollForObjectDeletion(func() (interface{}, *http.Response, error) {
			return client.VDBsApi.GetVdbById(ctx, vdbId).Execute()
		})
		// This would imply error in poll for deletion so we just log and exit.
		if diags != nil {
			ErrorLog.Printf("Error in polling of VDB for deletion.")
		} else {
			// diags will be nill in case of successful poll for deletion logic aka 404
			ErrorLog.Printf("Error reading the VDB %s, removing from state.", vdbId)
			d.SetId("")
		}

		// Todo check with 1225 -> if terraform refresh

		return nil
	}

	result, ok := res.(*dctapi.VDB)
	if !ok {
		return diag.Errorf("Error occured in type casting.")
	}

	d.Set("database_type", result.GetDatabaseType())
	d.Set("name", result.GetName())
	d.Set("database_version", result.GetDatabaseVersion())
	d.Set("engine_id", result.GetEngineId())
	d.Set("environment_id", result.GetEnvironmentId())
	d.Set("ip_address", result.GetIpAddress())
	d.Set("fqdn", result.GetFqdn())
	d.Set("parent_id", result.GetParentId())
	d.Set("group_name", result.GetGroupName())
	d.Set("creation_date", result.GetCreationDate().String())
	d.Set("id", vdbId)

	return diags
}

func resourceVdbUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	client := meta.(*apiClient).client
	updateVDBParam := dctapi.NewUpdateVDBParameters()

	// get the changed keys
	changedKeys := make([]string, 0, len(d.State().Attributes))
	for k := range d.State().Attributes {
		if d.HasChange(k) {
			changedKeys = append(changedKeys, k)
		}
	}

	if d.HasChanges(
		"auto_select_repository",
		"source_data_id",
		"id",
		"database_type",
		"database_version",
		"status",
		"ip_address",
		"fqdn",
		"parent_id",
		"group_name",
		"creation_date",
		"target_group_id",
		"database_name",
		"truncate_log_on_checkpoint",
		"repository_id",
		"pre_refresh",
		"post_refresh",
		"pre_rollback",
		"post_rollback",
		"configure_clone",
		"pre_snapshot",
		"post_snapshot",
		"pre_start",
		"post_start",
		"pre_stop",
		"post_stop",
		"file_mapping_rules",
		"oracle_instance_name",
		"unique_name",
		"mount_point",
		"open_reset_logs",
		"snapshot_policy_id",
		"retention_policy_id",
		"recovery_model",
		"online_log_groups",
		"online_log_size",
		"os_username",
		"os_password",
		"archive_log",
		"custom_env_vars",
		"custom_env_files",
		"timestamp",
		"timestamp_in_database_timezone",
		"snapshot_id") {

		// revert and set the old value to the changed keys
		for _, key := range changedKeys {
			old, _ := d.GetChange(key)
			d.Set(key, old)
		}

		return diag.Errorf("cannot update one (or more) of the options changed. Please refer to provider documentation for updatable params.")
	}

	if d.HasChange("template_id") {
		updateVDBParam.SetTemplateId(d.Get("template_id").(string))
	}
	if d.HasChange("name") {
		updateVDBParam.SetName(d.Get("name").(string))
	}
	if d.HasChange("db_username") {
		updateVDBParam.SetDbUsername(d.Get("db_username").(string))
	}
	if d.HasChange("db_password") {
		updateVDBParam.SetDbPassword(d.Get("db_password").(string))
	}
	if d.HasChange("new_dbid") {
		updateVDBParam.SetNewDbid(d.Get("new_dbid").(bool))
	}
	if d.HasChange("vdb_restart") {
		updateVDBParam.SetAutoRestart(d.Get("vdb_restart").(bool))
	}
	if d.HasChange("listener_ids") {
		updateVDBParam.SetListenerIds(toStringArray(d.Get("listener_ids")))
	}
	if d.HasChange("environment_user_id") {
		updateVDBParam.SetEnvironmentUserId(d.Get("environment_user_id").(string))
	}
	if d.HasChange("pre_script") {
		updateVDBParam.SetPreScript(d.Get("pre_script").(string))
	}
	if d.HasChange("post_script") {
		updateVDBParam.SetPostScript(d.Get("post_script").(string))
	}
	if d.HasChange("cdc_on_provision") {
		updateVDBParam.SetCdcOnProvision(d.Get("cdc_on_provision").(bool))
	}

	res, httpRes, err := client.VDBsApi.UpdateVdbById(ctx, d.Get("id").(string)).UpdateVDBParameters(*updateVDBParam).Execute()

	if diags := apiErrorResponseHelper(nil, httpRes, err); diags != nil {
		// revert and set the old value to the changed keys
		for _, key := range changedKeys {
			old, _ := d.GetChange(key)
			d.Set(key, old)
		}
		return diags
	}

	job_status, job_err := PollJobStatus(*res.Job.Id, ctx, client)
	if job_err != "" {
		WarnLog.Printf("VDB Update Job Polling failed but continuing with update. Error :%v", job_err)
	}
	InfoLog.Printf("Job result is %s", job_status)
	if isJobTerminalFailure(job_status) {
		return diag.Errorf("[NOT OK] VDB-Update %s. JobId: %s / Error: %s", job_status, *res.Job.Id, job_err)
	}

	return diags
}

func resourceVdbDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient).client

	vdbId := d.Id()

	deleteVdbParams := dctapi.NewDeleteVDBParametersWithDefaults()
	deleteVdbParams.SetForce(false)

	res, httpRes, err := client.VDBsApi.DeleteVdb(ctx, vdbId).DeleteVDBParameters(*deleteVdbParams).Execute()

	if diags := apiErrorResponseHelper(res, httpRes, err); diags != nil {
		return diags
	}

	job_status, job_err := PollJobStatus(*res.Job.Id, ctx, client)
	if job_err != "" {
		WarnLog.Printf("Job Polling failed but continuing with deletion. Error :%v", job_err)
	}
	InfoLog.Printf("Job result is %s", job_status)
	if isJobTerminalFailure(job_status) {
		return diag.Errorf("[NOT OK] VDB-Delete %s. JobId: %s / Error: %s", job_status, *res.Job.Id, job_err)
	}

	_, diags := PollForObjectDeletion(func() (interface{}, *http.Response, error) {
		return client.VDBsApi.GetVdbById(ctx, vdbId).Execute()
	})

	return diags
}
