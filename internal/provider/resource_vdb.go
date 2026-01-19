package provider

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"

	dctapi "github.com/delphix/dct-sdk-go/v25"
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
		CustomizeDiff: CustomizeDiffTags,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"provision_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "snapshot",
			},
			"auto_select_repository": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Whether to automatically select repository during provisioning. Defaults to false during import.",
			},
			"source_data_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The ID or name of the source dataset (dSource, VDB, Snapshot, or timestamp point) to provision from. Computed during import as root_parent_id.",
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
				Computed: true,
			},
			"cdb_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cluster_node_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
							Computed: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Computed: true,
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
							Computed: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Computed: true,
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
							Computed: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Computed: true,
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
							Computed: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Computed: true,
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
							Computed: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Computed: true,
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
							Computed: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Computed: true,
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
							Computed: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Computed: true,
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
							Computed: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Computed: true,
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
							Computed: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Computed: true,
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
							Computed: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Computed: true,
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
							Computed: true,
						},
						"has_credentials": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"vdb_restart": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"template_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"jdbc_connection_string": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auxiliary_template_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"file_mapping_rules": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed: true,
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
			"masked": {
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
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The ID or name of the Snapshot from which to execute the provision operation. If empty, uses the latest snapshot. Computed during import and read operations.",
			},
			"bookmark_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_dsource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"root_parent_id": {
				Type:     schema.TypeString,
				Computed: true,
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
			"appdata_source_params": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"appdata_config_params": {
				Type:     schema.TypeString,
				Computed: true,
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
			"config_params": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"additional_mount_points": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"shared_path": {
							Type:     schema.TypeString,
							Required: true,
						},
						"mount_path": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"environment_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"vcdb_tde_key_identifier": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cdb_tde_keystore_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_vcdb_tde_keystore_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tde_key_identifier": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tde_exported_key_file_secret": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_tde_keystore_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_tde_keystore_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oracle_rac_custom_env_vars": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
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
			"oracle_rac_custom_env_files": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"path_parameters": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
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
		element_id := item_map["element_id"].(string)
		if element_id != "" {
			hook_item.SetElementId(element_id)
		}
		has_credentials := item_map["has_credentials"].(bool)
		if has_credentials {
			hook_item.SetHasCredentials(has_credentials)
		}

		// defaults to "bash" as per resource schema spec
		hook_item.SetShell(item_map["shell"].(string))
		items = append(items, *hook_item)
	}
	return items
}

func toAdditionalMountPointsArray(array interface{}) []dctapi.AdditionalMountPoint {
	items := []dctapi.AdditionalMountPoint{}
	for _, item := range array.([]interface{}) {
		item_map := item.(map[string]interface{})
		addMntPts := dctapi.NewAdditionalMountPoint()
		addMntPts.SetEnvironmentId(item_map["environment_id"].(string))
		addMntPts.SetMountPath(item_map["mount_path"].(string))
		addMntPts.SetSharedPath(item_map["shared_path"].(string))

		items = append(items, *addMntPts)
	}
	return items
}

func toOracleRacCustomEnvVars(array interface{}) []dctapi.OracleRacCustomEnvVar {
	items := []dctapi.OracleRacCustomEnvVar{}
	for _, item := range array.([]interface{}) {
		item_map := item.(map[string]interface{})
		oracleRacCustomEnvVars := dctapi.NewOracleRacCustomEnvVar()
		oracleRacCustomEnvVars.SetName(item_map["name"].(string))
		oracleRacCustomEnvVars.SetNodeId(item_map["node_id"].(string))
		oracleRacCustomEnvVars.SetValue(item_map["value"].(string))

		items = append(items, *oracleRacCustomEnvVars)
	}
	return items
}

func toOracleRacCustomEnvFiles(array interface{}) []dctapi.OracleRacCustomEnvFile {
	items := []dctapi.OracleRacCustomEnvFile{}
	for _, item := range array.([]interface{}) {
		item_map := item.(map[string]interface{})
		oracleRacCustomEnvFiles := dctapi.NewOracleRacCustomEnvFile()
		oracleRacCustomEnvFiles.SetNodeId(item_map["node_id"].(string))
		oracleRacCustomEnvFiles.SetPathParameters(item_map["path_parameters"].(string))

		items = append(items, *oracleRacCustomEnvFiles)
	}
	return items
}

func helper_provision_by_snapshot(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*apiClient).client

	// respect resource create timeout
	createCtx, createCancel := context.WithTimeout(ctx, d.Timeout(schema.TimeoutCreate))
	defer createCancel()

	provisionVDBBySnapshotParameters := dctapi.NewProvisionVDBBySnapshotParameters()

	// Setters for provisionVDBBySnapshotParameters
	rawConfig := d.GetRawConfig()
	if attr := rawConfig.GetAttr("auto_select_repository"); !attr.IsNull() {
		provisionVDBBySnapshotParameters.SetAutoSelectRepository(d.Get("auto_select_repository").(bool))
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
	if attr := rawConfig.GetAttr("truncate_log_on_checkpoint"); !attr.IsNull() {
		provisionVDBBySnapshotParameters.SetTruncateLogOnCheckpoint(d.Get("truncate_log_on_checkpoint").(bool))
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
	if attr := rawConfig.GetAttr("auto_select_repository"); !attr.IsNull() {
		provisionVDBBySnapshotParameters.SetAutoSelectRepository(d.Get("auto_select_repository").(bool))
	}
	if attr := rawConfig.GetAttr("vdb_restart"); !attr.IsNull() {
		provisionVDBBySnapshotParameters.SetVdbRestart(d.Get("vdb_restart").(bool))
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
	if v, has_v := d.GetOk("instance_name"); has_v {
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
	if attr := rawConfig.GetAttr("open_reset_logs"); !attr.IsNull() {
		provisionVDBBySnapshotParameters.SetOpenResetLogs(d.Get("open_reset_logs").(bool))
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
	if attr := rawConfig.GetAttr("cdc_on_provision"); !attr.IsNull() {
		provisionVDBBySnapshotParameters.SetCdcOnProvision(d.Get("cdc_on_provision").(bool))
	}
	if attr := rawConfig.GetAttr("masked"); !attr.IsNull() {
		provisionVDBBySnapshotParameters.SetMasked(d.Get("masked").(bool))
	}
	if v, has_v := d.GetOk("online_log_size"); has_v {
		provisionVDBBySnapshotParameters.SetOnlineLogSize(int32(v.(int)))
	}
	if v, has_v := d.GetOk("online_log_groups"); has_v {
		provisionVDBBySnapshotParameters.SetOnlineLogGroups(int32(v.(int)))
	}
	if attr := rawConfig.GetAttr("archive_log"); !attr.IsNull() {
		provisionVDBBySnapshotParameters.SetArchiveLog(d.Get("archive_log").(bool))
	}
	if attr := rawConfig.GetAttr("new_dbid"); !attr.IsNull() {
		provisionVDBBySnapshotParameters.SetNewDbid(d.Get("new_dbid").(bool))
	}
	if v, has_v := d.GetOk("listener_ids"); has_v {
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
	if v, has_v := d.GetOk("appdata_source_params"); has_v {
		appdata_source_params := make(map[string]interface{})
		json.Unmarshal([]byte(v.(string)), &appdata_source_params)
		provisionVDBBySnapshotParameters.SetAppdataSourceParams(appdata_source_params)
	}
	if v, has_v := d.GetOk("appdata_config_params"); has_v {
		appdata_config_params := make(map[string]interface{})
		json.Unmarshal([]byte(v.(string)), &appdata_config_params)
		provisionVDBBySnapshotParameters.SetAppdataConfigParams(appdata_config_params)
	}
	if v, has_v := d.GetOk("config_params"); has_v {
		config_params := make(map[string]interface{})
		json.Unmarshal([]byte(v.(string)), &config_params)
		provisionVDBBySnapshotParameters.SetConfigParams(config_params)
	}
	if v, has_v := d.GetOk("make_current_account_owner"); has_v {
		provisionVDBBySnapshotParameters.SetMakeCurrentAccountOwner(v.(bool))
	}
	if v, has_v := d.GetOk("vcdb_tde_key_identifier"); has_v {
		provisionVDBBySnapshotParameters.SetVcdbTdeKeyIdentifier(v.(string))
	}
	if v, has_v := d.GetOk("cdb_tde_keystore_password"); has_v {
		provisionVDBBySnapshotParameters.SetCdbTdeKeystorePassword(v.(string))
	}
	if v, has_v := d.GetOk("target_vcdb_tde_keystore_path"); has_v {
		provisionVDBBySnapshotParameters.SetTargetVcdbTdeKeystorePath(v.(string))
	}
	if v, has_v := d.GetOk("tde_key_identifier"); has_v {
		provisionVDBBySnapshotParameters.SetTdeKeyIdentifier(v.(string))
	}
	if v, has_v := d.GetOk("tde_exported_key_file_secret"); has_v {
		provisionVDBBySnapshotParameters.SetTdeExportedKeyFileSecret(v.(string))
	}
	if v, has_v := d.GetOk("parent_tde_keystore_password"); has_v {
		provisionVDBBySnapshotParameters.SetParentTdeKeystorePassword(v.(string))
	}
	if v, has_v := d.GetOk("parent_tde_keystore_path"); has_v {
		provisionVDBBySnapshotParameters.SetParentTdeKeystorePath(v.(string))
	}
	if v, has_v := d.GetOk("additional_mount_points"); has_v {
		provisionVDBBySnapshotParameters.SetAdditionalMountPoints(toAdditionalMountPointsArray(v))
	}
	if v, has_v := d.GetOk("oracle_rac_custom_env_files"); has_v {
		provisionVDBBySnapshotParameters.SetOracleRacCustomEnvFiles(toOracleRacCustomEnvFiles(v))
	}
	if v, has_v := d.GetOk("oracle_rac_custom_env_vars"); has_v {
		provisionVDBBySnapshotParameters.SetOracleRacCustomEnvVars(toOracleRacCustomEnvVars(v))
	}

	req := client.VDBsAPI.ProvisionVdbBySnapshot(createCtx)

	apiRes, httpRes, err := req.ProvisionVDBBySnapshotParameters(*provisionVDBBySnapshotParameters).Execute()
	
	// Check if the API call itself timed out
	if err != nil && createCtx.Err() == context.DeadlineExceeded {
		return diag.Errorf("VDB provisioning API call timed out after %s. "+
			"Check DCT UI for job status. If created, find the VDB ID and import it.",
			d.Timeout(schema.TimeoutCreate))
	}
	
	if diags := apiErrorResponseHelper(ctx, apiRes, httpRes, err); diags != nil {
		return diags
	}

	// Check for nil apiRes or Job to prevent crashes
	if apiRes == nil || apiRes.Job == nil {
		return diag.Errorf("VDB provisioning failed: received nil response or job from API")
	}

	// Store VDB ID temporarily - don't set in state until job completes
	vdbId := apiRes.GetVdbId()

	job_res, job_err := PollJobStatus(apiRes.Job.GetId(), createCtx, client)
	if job_err != "" {
		tflog.Error(ctx, DLPX+ERROR+"Job Polling failed but continuing with provisioning. Error: "+job_err)
	}
	
	// Check if context was cancelled due to timeout
	if createCtx.Err() != nil {
		// Don't set ID in state - let user verify and import
		if createCtx.Err() == context.DeadlineExceeded {
			return diag.Errorf("VDB provisioning timed out after %s (Job ID: %s, VDB ID: %s). "+
				"Check DCT UI to verify job completion, then import it.",
				d.Timeout(schema.TimeoutCreate), apiRes.Job.GetId(), vdbId)
		}
		return diag.Errorf("VDB provisioning was cancelled (Job ID: %s): %v", apiRes.Job.GetId(), createCtx.Err())
	}
	
	tflog.Info(ctx, DLPX+INFO+"Job result is "+job_res)
	if job_res == Failed || job_res == Canceled || job_res == Abandoned {
		tflog.Error(ctx, DLPX+ERROR+"Job "+job_res+" "+apiRes.Job.GetId()+"!")
		return diag.Errorf("[NOT OK] Job %s %s with error %s", apiRes.Job.GetId(), job_res, job_err)
	}
	
	// Check context before reading state
	if createCtx.Err() != nil {
		if createCtx.Err() == context.DeadlineExceeded {
			return diag.Errorf("VDB provisioning timed out after %s during final state read (Job ID: %s). "+
				"The VDB may have been created. To resolve:\n"+
				"1. Check the Delphix DCT UI or API to verify the VDB exists\n"+
				"2. If created successfully, import it.",
				d.Timeout(schema.TimeoutCreate), apiRes.Job.GetId())
		}
		return diag.Errorf("VDB provisioning was cancelled during final state read (Job ID: %s): %v", apiRes.Job.GetId(), createCtx.Err())
	}
	
	// Only set ID in state after successful completion
	d.SetId(vdbId)
	
	readDiags := resourceVdbRead(createCtx, d, meta)

	if readDiags.HasError() {
		return readDiags
	}

	return diags
}

func helper_provision_by_timestamp(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*apiClient).client

	// respect resource create timeout
	createCtx, createCancel := context.WithTimeout(ctx, d.Timeout(schema.TimeoutCreate))
	defer createCancel()

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
	rawConfig := d.GetRawConfig()
	if attr := rawConfig.GetAttr("truncate_log_on_checkpoint"); !attr.IsNull() {
		provisionVDBByTimestampParameters.SetTruncateLogOnCheckpoint(d.Get("truncate_log_on_checkpoint").(bool))
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
	if attr := rawConfig.GetAttr("auto_select_repository"); !attr.IsNull() {
		provisionVDBByTimestampParameters.SetAutoSelectRepository(d.Get("auto_select_repository").(bool))
	}
	if attr := rawConfig.GetAttr("vdb_restart"); !attr.IsNull() {
		provisionVDBByTimestampParameters.SetVdbRestart(d.Get("vdb_restart").(bool))
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
	if v, has_v := d.GetOk("instance_name"); has_v {
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
	if attr := rawConfig.GetAttr("open_reset_logs"); !attr.IsNull() {
		provisionVDBByTimestampParameters.SetOpenResetLogs(d.Get("open_reset_logs").(bool))
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
	if attr := rawConfig.GetAttr("cdc_on_provision"); !attr.IsNull() {
		provisionVDBByTimestampParameters.SetCdcOnProvision(d.Get("cdc_on_provision").(bool))
	}
	if attr := rawConfig.GetAttr("masked"); !attr.IsNull() {
		provisionVDBByTimestampParameters.SetMasked(d.Get("masked").(bool))
	}
	if v, has_v := d.GetOk("online_log_size"); has_v {
		provisionVDBByTimestampParameters.SetOnlineLogSize(int32(v.(int)))
	}
	if v, has_v := d.GetOk("online_log_groups"); has_v {
		provisionVDBByTimestampParameters.SetOnlineLogGroups(int32(v.(int)))
	}
	if attr := rawConfig.GetAttr("archive_log"); !attr.IsNull() {
		provisionVDBByTimestampParameters.SetArchiveLog(d.Get("archive_log").(bool))
	}
	if attr := rawConfig.GetAttr("new_dbid"); !attr.IsNull() {
		provisionVDBByTimestampParameters.SetNewDbid(d.Get("new_dbid").(bool))
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
			tflog.Error(ctx, DLPX+ERROR+"An error has occurred: "+err.Error())
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
	if v, has_v := d.GetOk("appdata_source_params"); has_v {
		appdata_source_params := make(map[string]interface{})
		json.Unmarshal([]byte(v.(string)), &appdata_source_params)
		provisionVDBByTimestampParameters.SetAppdataSourceParams(appdata_source_params)
	}
	if v, has_v := d.GetOk("appdata_config_params"); has_v {
		appdata_config_params := make(map[string]interface{})
		json.Unmarshal([]byte(v.(string)), &appdata_config_params)
		provisionVDBByTimestampParameters.SetAppdataConfigParams(appdata_config_params)
	}
	if v, has_v := d.GetOk("config_params"); has_v {
		config_params := make(map[string]interface{})
		json.Unmarshal([]byte(v.(string)), &config_params)
		provisionVDBByTimestampParameters.SetConfigParams(config_params)
	}
	if v, has_v := d.GetOk("make_current_account_owner"); has_v {
		provisionVDBByTimestampParameters.SetMakeCurrentAccountOwner(v.(bool))
	}
	if v, has_v := d.GetOk("vcdb_tde_key_identifier"); has_v {
		provisionVDBByTimestampParameters.SetVcdbTdeKeyIdentifier(v.(string))
	}
	if v, has_v := d.GetOk("cdb_tde_keystore_password"); has_v {
		provisionVDBByTimestampParameters.SetCdbTdeKeystorePassword(v.(string))
	}
	if v, has_v := d.GetOk("target_vcdb_tde_keystore_path"); has_v {
		provisionVDBByTimestampParameters.SetTargetVcdbTdeKeystorePath(v.(string))
	}
	if v, has_v := d.GetOk("tde_key_identifier"); has_v {
		provisionVDBByTimestampParameters.SetTdeKeyIdentifier(v.(string))
	}
	if v, has_v := d.GetOk("tde_exported_key_file_secret"); has_v {
		provisionVDBByTimestampParameters.SetTdeExportedKeyFileSecret(v.(string))
	}
	if v, has_v := d.GetOk("parent_tde_keystore_password"); has_v {
		provisionVDBByTimestampParameters.SetParentTdeKeystorePassword(v.(string))
	}
	if v, has_v := d.GetOk("parent_tde_keystore_path"); has_v {
		provisionVDBByTimestampParameters.SetParentTdeKeystorePath(v.(string))
	}
	if v, has_v := d.GetOk("additional_mount_points"); has_v {
		provisionVDBByTimestampParameters.SetAdditionalMountPoints(toAdditionalMountPointsArray(v))
	}
	if v, has_v := d.GetOk("oracle_rac_custom_env_files"); has_v {
		provisionVDBByTimestampParameters.SetOracleRacCustomEnvFiles(toOracleRacCustomEnvFiles(v))
	}
	if v, has_v := d.GetOk("oracle_rac_custom_env_vars"); has_v {
		provisionVDBByTimestampParameters.SetOracleRacCustomEnvVars(toOracleRacCustomEnvVars(v))
	}

	req := client.VDBsAPI.ProvisionVdbByTimestamp(createCtx)

	apiRes, httpRes, err := req.ProvisionVDBByTimestampParameters(*provisionVDBByTimestampParameters).Execute()
	
	// Check if the API call itself timed out
	if err != nil && createCtx.Err() == context.DeadlineExceeded {
		return diag.Errorf("VDB provisioning API call timed out after %s. "+
			"Check DCT UI for job status. If created, find the VDB ID and import it.",
			d.Timeout(schema.TimeoutCreate))
	}
	
	if diags := apiErrorResponseHelper(ctx, apiRes, httpRes, err); diags != nil {
		return diags
	}

	// Check for nil apiRes or Job to prevent crashes
	if apiRes == nil || apiRes.Job == nil {
		return diag.Errorf("VDB provisioning failed: received nil response or job from API")
	}

	// Store VDB ID temporarily - don't set in state until job completes
	vdbId := apiRes.GetVdbId()

	job_res, job_err := PollJobStatus(apiRes.Job.GetId(), createCtx, client)
	if job_err != "" {
		tflog.Error(ctx, DLPX+ERROR+"Job Polling failed but continuing with provisioning. Error: "+job_err)
	}
	
	// Check if context was cancelled due to timeout
	if createCtx.Err() != nil {
		// Don't set ID in state - let user verify and import
		if createCtx.Err() == context.DeadlineExceeded {
			return diag.Errorf("VDB provisioning timed out after %s (Job ID: %s, VDB ID: %s). "+
				"Check DCT UI to verify job completion, then import it.",
				d.Timeout(schema.TimeoutCreate), apiRes.Job.GetId(), vdbId)
		}
		return diag.Errorf("VDB provisioning was cancelled (Job ID: %s): %v", apiRes.Job.GetId(), createCtx.Err())
	}
	
	tflog.Info(ctx, DLPX+INFO+"Job result is "+job_res)
	if job_res == "FAILED" {
		tflog.Error(ctx, DLPX+ERROR+"Job "+apiRes.Job.GetId()+" Failed!")
		return diag.Errorf("[NOT OK] Job %s Failed with error %s", apiRes.Job.GetId(), job_err)
	}
	
	// Check context before reading state
	if createCtx.Err() != nil {
		if createCtx.Err() == context.DeadlineExceeded {
			return diag.Errorf("VDB provisioning timed out after %s during final state read (Job ID: %s). "+
				"The VDB may have been created. To resolve:\n"+
				"1. Check the Delphix DCT UI or API to verify the VDB exists\n"+
				"2. If created successfully, import it.",
				d.Timeout(schema.TimeoutCreate), apiRes.Job.GetId())
		}
		return diag.Errorf("VDB provisioning was cancelled during final state read (Job ID: %s): %v", apiRes.Job.GetId(), createCtx.Err())
	}
	
	// Only set ID in state after successful completion
	d.SetId(vdbId)
	
	readDiags := resourceVdbRead(createCtx, d, meta)

	if readDiags.HasError() {
		return readDiags
	}

	return diags
}

func helper_provision_by_bookmark(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*apiClient).client

	// respect resource create timeout
	createCtx, createCancel := context.WithTimeout(ctx, d.Timeout(schema.TimeoutCreate))
	defer createCancel()

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
	rawConfig := d.GetRawConfig()
	if attr := rawConfig.GetAttr("truncate_log_on_checkpoint"); !attr.IsNull() {
		provisionVDBFromBookmarkParameters.SetTruncateLogOnCheckpoint(d.Get("truncate_log_on_checkpoint").(bool))
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
	if attr := rawConfig.GetAttr("auto_select_repository"); !attr.IsNull() {
		provisionVDBFromBookmarkParameters.SetAutoSelectRepository(d.Get("auto_select_repository").(bool))
	}
	if attr := rawConfig.GetAttr("vdb_restart"); !attr.IsNull() {
		provisionVDBFromBookmarkParameters.SetVdbRestart(d.Get("vdb_restart").(bool))
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
	if v, has_v := d.GetOk("instance_name"); has_v {
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
	if attr := rawConfig.GetAttr("open_reset_logs"); !attr.IsNull() {
		provisionVDBFromBookmarkParameters.SetOpenResetLogs(d.Get("open_reset_logs").(bool))
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
	if attr := rawConfig.GetAttr("cdc_on_provision"); !attr.IsNull() {
		provisionVDBFromBookmarkParameters.SetCdcOnProvision(d.Get("cdc_on_provision").(bool))
	}
	if attr := rawConfig.GetAttr("masked"); !attr.IsNull() {
		provisionVDBFromBookmarkParameters.SetMasked(d.Get("masked").(bool))
	}
	if v, has_v := d.GetOk("online_log_size"); has_v {
		provisionVDBFromBookmarkParameters.SetOnlineLogSize(int32(v.(int)))
	}
	if v, has_v := d.GetOk("online_log_groups"); has_v {
		provisionVDBFromBookmarkParameters.SetOnlineLogGroups(int32(v.(int)))
	}
	if attr := rawConfig.GetAttr("archive_log"); !attr.IsNull() {
		provisionVDBFromBookmarkParameters.SetArchiveLog(d.Get("archive_log").(bool))
	}
	if attr := rawConfig.GetAttr("new_dbid"); !attr.IsNull() {
		provisionVDBFromBookmarkParameters.SetNewDbid(d.Get("new_dbid").(bool))
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
	if v, has_v := d.GetOk("appdata_source_params"); has_v {
		appdata_source_params := make(map[string]interface{})
		json.Unmarshal([]byte(v.(string)), &appdata_source_params)
		provisionVDBFromBookmarkParameters.SetAppdataSourceParams(appdata_source_params)
	}
	if v, has_v := d.GetOk("appdata_config_params"); has_v {
		appdata_config_params := make(map[string]interface{})
		json.Unmarshal([]byte(v.(string)), &appdata_config_params)
		provisionVDBFromBookmarkParameters.SetAppdataConfigParams(appdata_config_params)
	}
	if v, has_v := d.GetOk("config_params"); has_v {
		config_params := make(map[string]interface{})
		json.Unmarshal([]byte(v.(string)), &config_params)
		provisionVDBFromBookmarkParameters.SetConfigParams(config_params)
	}
	if v, has_v := d.GetOk("make_current_account_owner"); has_v {
		provisionVDBFromBookmarkParameters.SetMakeCurrentAccountOwner(v.(bool))
	}
	if v, has_v := d.GetOk("vcdb_tde_key_identifier"); has_v {
		provisionVDBFromBookmarkParameters.SetVcdbTdeKeyIdentifier(v.(string))
	}
	if v, has_v := d.GetOk("cdb_tde_keystore_password"); has_v {
		provisionVDBFromBookmarkParameters.SetCdbTdeKeystorePassword(v.(string))
	}
	if v, has_v := d.GetOk("target_vcdb_tde_keystore_path"); has_v {
		provisionVDBFromBookmarkParameters.SetTargetVcdbTdeKeystorePath(v.(string))
	}
	if v, has_v := d.GetOk("tde_key_identifier"); has_v {
		provisionVDBFromBookmarkParameters.SetTdeKeyIdentifier(v.(string))
	}
	if v, has_v := d.GetOk("tde_exported_key_file_secret"); has_v {
		provisionVDBFromBookmarkParameters.SetTdeExportedKeyFileSecret(v.(string))
	}
	if v, has_v := d.GetOk("parent_tde_keystore_password"); has_v {
		provisionVDBFromBookmarkParameters.SetParentTdeKeystorePassword(v.(string))
	}
	if v, has_v := d.GetOk("parent_tde_keystore_path"); has_v {
		provisionVDBFromBookmarkParameters.SetParentTdeKeystorePath(v.(string))
	}
	if v, has_v := d.GetOk("additional_mount_points"); has_v {
		provisionVDBFromBookmarkParameters.SetAdditionalMountPoints(toAdditionalMountPointsArray(v))
	}
	if v, has_v := d.GetOk("oracle_rac_custom_env_files"); has_v {
		provisionVDBFromBookmarkParameters.SetOracleRacCustomEnvFiles(toOracleRacCustomEnvFiles(v))
	}
	if v, has_v := d.GetOk("oracle_rac_custom_env_vars"); has_v {
		provisionVDBFromBookmarkParameters.SetOracleRacCustomEnvVars(toOracleRacCustomEnvVars(v))
	}

	req := client.VDBsAPI.ProvisionVdbFromBookmark(createCtx)

	apiRes, httpRes, err := req.ProvisionVDBFromBookmarkParameters(*provisionVDBFromBookmarkParameters).Execute()
	
	// Check if the API call itself timed out
	if err != nil && createCtx.Err() == context.DeadlineExceeded {
		return diag.Errorf("VDB provisioning API call timed out after %s. "+
			"Check DCT UI for job status. If created, find the VDB ID and import it.",
			d.Timeout(schema.TimeoutCreate))
	}
	
	if diags := apiErrorResponseHelper(ctx, apiRes, httpRes, err); diags != nil {
		return diags
	}

	// Check for nil apiRes or Job to prevent crashes
	if apiRes == nil || apiRes.Job == nil {
		return diag.Errorf("VDB provisioning failed: received nil response or job from API")
	}

	// Store VDB ID temporarily - don't set in state until job completes
	vdbId := apiRes.GetVdbId()

	job_res, job_err := PollJobStatus(apiRes.Job.GetId(), createCtx, client)
	if job_err != "" {
		tflog.Error(ctx, DLPX+ERROR+"Job Polling failed but continuing with provisioning. Error: "+job_err)
	}
	
	// Check if context was cancelled due to timeout
	if createCtx.Err() != nil {
		// Don't set ID in state - let user verify and import
		if createCtx.Err() == context.DeadlineExceeded {
			return diag.Errorf("VDB provisioning timed out after %s (Job ID: %s, VDB ID: %s). "+
				"Check DCT UI to verify job completion, then import it.",
				d.Timeout(schema.TimeoutCreate), apiRes.Job.GetId(), vdbId)
		}
		return diag.Errorf("VDB provisioning was cancelled (Job ID: %s): %v", apiRes.Job.GetId(), createCtx.Err())
	}
	
	tflog.Info(ctx, DLPX+INFO+"Job result is "+job_res)
	if job_res == Failed || job_res == Canceled || job_res == Abandoned {
		tflog.Error(ctx, DLPX+ERROR+"Job "+job_res+apiRes.Job.GetId()+"!")
		return diag.Errorf("[NOT OK] Job %s %s with error %s", apiRes.Job.GetId(), job_res, job_err)
	}
	
	// Check context before reading state
	if createCtx.Err() != nil {
		if createCtx.Err() == context.DeadlineExceeded {
			return diag.Errorf("VDB provisioning timed out after %s during final state read (Job ID: %s). "+
				"The VDB may have been created. To resolve:\n"+
				"1. Check the Delphix DCT UI or API to verify the VDB exists\n"+
				"2. If created successfully, import it.",
				d.Timeout(schema.TimeoutCreate), apiRes.Job.GetId())
		}
		return diag.Errorf("VDB provisioning was cancelled during final state read (Job ID: %s): %v", apiRes.Job.GetId(), createCtx.Err())
	}
	
	// Only set ID in state after successful completion
	d.SetId(vdbId)
	
	readDiags := resourceVdbRead(createCtx, d, meta)

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

// getSnapshotIdFromTimeflow retrieves the parent snapshot ID for a VDB by querying its current timeflow
func getSnapshotIdFromTimeflow(ctx context.Context, client *dctapi.APIClient, vdbId string) (string, error) {
	// Step 1: Get the VDB to retrieve its current timeflow ID
	vdb, _, err := client.VDBsAPI.GetVdbById(ctx, vdbId).Execute()
	if err != nil {
		return "", err
	}
	
	// Step 2: Get the current timeflow ID from the VDB
	currentTimeflowId, ok := vdb.GetCurrentTimeflowIdOk()
	if !ok || currentTimeflowId == nil || *currentTimeflowId == "" {
		return "", nil // No current timeflow ID found
	}
	
	// Step 3: Get the timeflow using the timeflow ID
	timeflow, _, err := client.TimeflowsAPI.GetTimeflowById(ctx, *currentTimeflowId).Execute()
	if err != nil {
		return "", err
	}
	
	// Step 4: Get the parent snapshot ID from the timeflow
	if parentSnapshotId, ok := timeflow.GetParentSnapshotIdOk(); ok && parentSnapshotId != nil {
		return *parentSnapshotId, nil
	}
	
	return "", nil // No parent snapshot ID found
}

func resourceVdbRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*apiClient).client

	vdbId := d.Id()

	res, diags := PollForObjectExistence(ctx, func() (interface{}, *http.Response, error) {
		return client.VDBsAPI.GetVdbById(ctx, vdbId).Execute()
	})

	if res == nil {
		tflog.Error(ctx, DLPX+ERROR+"VDB not found: "+vdbId+", removing from state. ")
		d.SetId("")
		return nil
	}

	if diags != nil {
		_, diags := PollForObjectDeletion(ctx, func() (interface{}, *http.Response, error) {
			return client.VDBsAPI.GetVdbById(ctx, vdbId).Execute()
		})
		// This would imply error in poll for deletion so we just log and exit.
		if diags != nil {
			tflog.Error(ctx, DLPX+ERROR+"Error in polling of VDB for deletion.")
		} else {
			// diags will be nill in case of successful poll for deletion logic aka 404
			tflog.Error(ctx, DLPX+ERROR+"Error reading the VDB "+vdbId+", removing from state. ")
			d.SetId("")
		}

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
	d.Set("parent_dsource_id", result.GetParentDsourceId())
	d.Set("root_parent_id", result.GetRootParentId())
	d.Set("group_name", result.GetGroupName())
	d.Set("creation_date", result.GetCreationDate().String())
	d.Set("instance_name", result.GetInstanceName())
	
	// Handle auto_select_repository - preserve from config, leave null during import to avoid diffs
	if autoSelectRepo, hasAutoSelectRepo := d.GetOk("auto_select_repository"); hasAutoSelectRepo {
		// If auto_select_repository is in config, keep it in state
		d.Set("auto_select_repository", autoSelectRepo.(bool))
	}
	// If not in config (e.g., during import), don't set it - this avoids plan diffs
	
	// Handle source_data_id - preserve from config or set to root_parent_id during import
	if sourceDataId, hasSourceDataId := d.GetOk("source_data_id"); hasSourceDataId {
		// If source_data_id is in config, keep it in state
		d.Set("source_data_id", sourceDataId.(string))
	} else {
		// For imports or when source_data_id is not specified, set it to parent_id
		if ParentId := result.GetParentId(); ParentId != "" {
			d.Set("source_data_id", ParentId)
			tflog.Debug(ctx, "Set source_data_id from parent_id: "+ParentId)
		}
	}
	
	// Handle snapshot_id - preserve from config or fetch from API during import
	if snapshotId, hasSnapshotId := d.GetOk("snapshot_id"); hasSnapshotId {
		// If snapshot_id is in config, keep it in state
		d.Set("snapshot_id", snapshotId.(string))
	} else {
		// For imports or when snapshot_id is not specified, try to fetch it from the timeflow
		if parentSnapshotId, err := getSnapshotIdFromTimeflow(ctx, client, vdbId); err == nil && parentSnapshotId != "" {
			d.Set("snapshot_id", parentSnapshotId)
			tflog.Debug(ctx, "Retrieved snapshot_id from timeflow: "+parentSnapshotId)
		} else if err != nil {
			tflog.Warn(ctx, DLPX+WARN+"Failed to retrieve snapshot_id from timeflow: "+err.Error())
		}
	}
	
	d.Set("pre_refresh", flattenVDbHooks(result.GetHooks().PreRefresh))
	d.Set("post_refresh", flattenVDbHooks(result.GetHooks().PostRefresh))
	d.Set("configure_clone", flattenVDbHooks(result.GetHooks().ConfigureClone))
	d.Set("pre_snapshot", flattenVDbHooks(result.GetHooks().PreSnapshot))
	d.Set("post_snapshot", flattenVDbHooks(result.GetHooks().PostSnapshot))
	d.Set("pre_start", flattenVDbHooks(result.GetHooks().PreStart))
	d.Set("post_start", flattenVDbHooks(result.GetHooks().PostStart))
	d.Set("pre_stop", flattenVDbHooks(result.GetHooks().PreStop))
	d.Set("post_stop", flattenVDbHooks(result.GetHooks().PostStop))
	d.Set("pre_rollback", flattenVDbHooks(result.GetHooks().PreRollback))
	d.Set("post_rollback", flattenVDbHooks(result.GetHooks().PostRollback))

	if !result.GetIsAppdata() {
		d.Set("database_name", result.GetDatabaseName())
	}

	HandleRawConfigReadContext(ctx, d, result)
	d.Set("vdb_restart", result.GetVdbRestart())

	_, is_provision := d.GetOk("provision_type")
	if !is_provision {
		// its an import, set to default value
		d.Set("provision_type", "snapshot")
	}

	// Set make_current_account_owner to default true if not explicitly set
	if _, has_make_current := d.GetOk("make_current_account_owner"); !has_make_current {
		d.Set("make_current_account_owner", true)
	}

	// Set ignore_tag_changes to default true if not explicitly set
	if _, has_ignore_tags := d.GetOk("ignore_tag_changes"); !has_ignore_tags {
		d.Set("ignore_tag_changes", true)
	}

	d.Set("jdbc_connection_string", result.GetJdbcConnectionString())
	d.Set("cdb_id", result.GetCdbId())
	d.Set("template_id", result.GetTemplateId())
	d.Set("mount_point", result.GetMountPoint())

	appdata_source_params, _ := json.Marshal(result.GetAppdataSourceParams())
	d.Set("appdata_source_params", string(appdata_source_params))
	appdata_config_params, _ := json.Marshal(result.GetAppdataConfigParams())
	d.Set("appdata_config_params", string(appdata_config_params))
	config_params, _ := json.Marshal(result.GetConfigParams())
	d.Set("config_params", string(config_params))
	d.Set("additional_mount_points", flattenAdditionalMountPoints(result.GetAdditionalMountPoints()))

	d.Set("id", vdbId)

	return diags
}

func resourceVdbUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*apiClient).client
	// respect resource update timeout
	updateCtx, updateCancel := context.WithTimeout(ctx, d.Timeout(schema.TimeoutUpdate))
	defer updateCancel()
	updateVDBParam := dctapi.NewUpdateVDBParameters()

	vdbId := d.Get("id").(string)

	// get the changed keys
	changedKeys := make([]string, 0, len(d.State().Attributes))
	for k := range d.State().Attributes {
		if strings.Contains(k, "tags") { // this is because the changed keys are of the form tag.0.keydi
			k = "tags"
		}
		if strings.Contains(k, "pre_refresh") {
			k = "pre_refresh"
		}
		if strings.Contains(k, "post_refresh") {
			k = "post_refresh"
		}
		if strings.Contains(k, "configure_clone") {
			k = "configure_clone"
		}
		if strings.Contains(k, "pre_snapshot") {
			k = "pre_snapshot"
		}
		if strings.Contains(k, "post_snapshot") {
			k = "post_snapshot"
		}
		if strings.Contains(k, "pre_rollback") {
			k = "pre_rollback"
		}
		if strings.Contains(k, "post_rollback") {
			k = "post_rollback"
		}
		if strings.Contains(k, "pre_start") {
			k = "pre_start"
		}
		if strings.Contains(k, "post_start") {
			k = "post_start"
		}
		if strings.Contains(k, "pre_stop") {
			k = "pre_stop"
		}
		if strings.Contains(k, "post_stop") {
			k = "post_stop"
		}
		if strings.Contains(k, "additional_mount_points") {
			k = "additional_mount_points"
		}
		if strings.Contains(k, "listener_ids") {
			k = "listener_ids"
		}
		// Skip auto_select_repository - it's a provision-time parameter that shouldn't trigger updates
		if k == "auto_select_repository" {
			continue
		}
		// Skip timeouts - it's a Terraform meta-argument that shouldn't trigger resource updates
		if strings.Contains(k, "timeouts") {
			continue
		}
		if d.HasChange(k) {
			tflog.Debug(ctx, "changed keys"+k)
			changedKeys = append(changedKeys, k)
		}
	}

	// If no actual changes (e.g., only auto_select_repository changed), skip update and just read
	if len(changedKeys) == 0 {
		tflog.Debug(ctx, "No updatable fields changed, skipping update operation")
		return resourceVdbRead(ctx, d, meta)
	}

	var updateFailure, destructiveUpdate bool = false, false
	var nonUpdatableField []string

	// if changedKeys contains non updatable field set a flag
	for _, key := range changedKeys {
		if !updatableVdbKeys[key] {
			updateFailure = true
			tflog.Debug(ctx, "non updatable field: "+key)
			nonUpdatableField = append(nonUpdatableField, key)
		}
	}

	if updateFailure {
		revertChanges(d, changedKeys)
		return diag.Errorf("cannot update options %v. Please refer to provider documentation for updatable params.", nonUpdatableField)
	}

	// find if destructive update
	for _, key := range changedKeys {
		if isDestructiveVdbUpdate[key] {
			tflog.Debug(ctx, "destructive updates for: "+key)
			destructiveUpdate = true
		}
	}
	if destructiveUpdate {
		if diags := disableVDB(ctx, client, vdbId); diags != nil {
			tflog.Error(ctx, "failure in disabling vdbs")
			revertChanges(d, changedKeys)
			return diags
		}
	}

	nvdh := dctapi.NewVirtualDatasetHooks()

	if d.HasChange("pre_refresh") {
		if v, has_v := d.GetOk("pre_refresh"); has_v {
			nvdh.SetPreRefresh(toHookArray(v))
		} else {
			nvdh.SetPreRefresh([]dctapi.Hook{})
		}
	}

	if d.HasChange("post_refresh") {
		if v, has_v := d.GetOk("post_refresh"); has_v {
			nvdh.SetPostRefresh(toHookArray(v))
		} else {
			nvdh.SetPostRefresh([]dctapi.Hook{})
		}
	}

	if d.HasChange("pre_rollback") {
		if v, has_v := d.GetOk("pre_rollback"); has_v {
			nvdh.SetPreRollback(toHookArray(v))
		} else {
			nvdh.SetPreRollback([]dctapi.Hook{})
		}
	}

	if d.HasChange("post_rollback") {
		if v, has_v := d.GetOk("post_rollback"); has_v {
			nvdh.SetPostRollback(toHookArray(v))
		} else {
			nvdh.SetPostRollback([]dctapi.Hook{})
		}
	}

	if d.HasChange("configure_clone") {
		if v, has_v := d.GetOk("configure_clone"); has_v {
			nvdh.SetConfigureClone(toHookArray(v))
		} else {
			nvdh.SetConfigureClone([]dctapi.Hook{})
		}
	}

	if d.HasChange("pre_snapshot") {
		if v, has_v := d.GetOk("pre_snapshot"); has_v {
			nvdh.SetPreSnapshot(toHookArray(v))
		} else {
			nvdh.SetPreSnapshot([]dctapi.Hook{})
		}
	}

	if d.HasChange("post_snapshot") {
		if v, has_v := d.GetOk("post_snapshot"); has_v {
			nvdh.SetPostSnapshot(toHookArray(v))
		} else {
			nvdh.SetPostSnapshot([]dctapi.Hook{})
		}
	}

	if d.HasChange("pre_start") {
		if v, has_v := d.GetOk("pre_start"); has_v {
			nvdh.SetPreStart(toHookArray(v))
		} else {
			nvdh.SetPreStart([]dctapi.Hook{})
		}
	}

	if d.HasChange("post_start") {
		if v, has_v := d.GetOk("post_start"); has_v {
			nvdh.SetPostStart(toHookArray(v))
		} else {
			nvdh.SetPostStart([]dctapi.Hook{})
		}
	}

	if d.HasChange("pre_stop") {
		if v, has_v := d.GetOk("pre_stop"); has_v {
			nvdh.SetPreStop(toHookArray(v))
		} else {
			nvdh.SetPreStop([]dctapi.Hook{})
		}
	}

	if d.HasChange("post_stop") {
		if v, has_v := d.GetOk("post_stop"); has_v {
			nvdh.SetPostStop(toHookArray(v))
		} else {
			nvdh.SetPostStop([]dctapi.Hook{})
		}
	}

	if nvdh != nil && !isStructEmpty(nvdh) {
		updateVDBParam.SetHooks(*nvdh)
	}

	if d.HasChange("mount_point") {
		updateVDBParam.SetMountPoint(d.Get("mount_point").(string))
	}

	if d.HasChange("custom_env_files") {
		if v, has_v := d.GetOk("custom_env_files"); has_v {
			updateVDBParam.SetCustomEnvFiles(toStringArray(v))
		} else {
			updateVDBParam.SetCustomEnvFiles([]string{})
		}
	}
	if d.HasChange("custom_env_vars") {
		if v, has_v := d.GetOk("custom_env_vars"); has_v {
			custom_env_vars := make(map[string]string)

			for k, v := range v.(map[string]interface{}) {
				custom_env_vars[k] = v.(string)
			}
			updateVDBParam.SetCustomEnvVars(custom_env_vars)
		} else {
			updateVDBParam.SetCustomEnvVars(map[string]string{})
		}
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
	if d.HasChange("additional_mount_points") {
		updateVDBParam.SetAdditionalMountPoints(toAdditionalMountPointsArray(d.Get("additional_mount_points")))
	}
	if d.HasChange("parent_tde_keystore_path") {
		updateVDBParam.SetParentTdeKeystorePath(d.Get("parent_tde_keystore_path").(string))
	}
	if d.HasChange("parent_tde_keystore_password") {
		updateVDBParam.SetParentTdeKeystorePassword(d.Get("parent_tde_keystore_password").(string))
	}
	if d.HasChange("tde_key_identifier") {
		updateVDBParam.SetTdeKeyIdentifier(d.Get("tde_key_identifier").(string))
	}
	if d.HasChange("target_vcdb_tde_keystore_path") {
		updateVDBParam.SetTargetVcdbTdeKeystorePath(d.Get("target_vcdb_tde_keystore_path").(string))
	}
	if d.HasChange("cdb_tde_keystore_password") {
		updateVDBParam.SetCdbTdeKeystorePassword(d.Get("cdb_tde_keystore_password").(string))
	}
	if d.HasChange("appdata_source_params") {
		appdata_source_params := make(map[string]interface{})
		json.Unmarshal([]byte(d.Get("appdata_source_params").(string)), &appdata_source_params)
		updateVDBParam.SetAppdataSourceParams(appdata_source_params)
	}
	if d.HasChange("appdata_config_params") {
		appdata_config_params := make(map[string]interface{})
		json.Unmarshal([]byte(d.Get("appdata_config_params").(string)), &appdata_config_params)
		updateVDBParam.SetAppdataConfigParams(appdata_config_params)
	}
	if d.HasChange("config_params") {
		config_params := make(map[string]interface{})
		json.Unmarshal([]byte(d.Get("config_params").(string)), &config_params)
		updateVDBParam.SetConfigParams(config_params)
	}
	if !isStructEmpty(updateVDBParam) {
		res, httpRes, err := client.VDBsAPI.UpdateVdbById(updateCtx, d.Get("id").(string)).UpdateVDBParameters(*updateVDBParam).Execute()

		if diags := apiErrorResponseHelper(ctx, nil, httpRes, err); diags != nil {
			// revert and set the old value to the changed keys
			revertChanges(d, changedKeys)
			return diags
		}

		if res != nil {
				job_status, job_err := PollJobStatus(res.Job.GetId(), updateCtx, client)
			if job_err != "" {
				tflog.Warn(ctx, DLPX+WARN+"VDB Update Job Polling failed but continuing with update. Error: "+job_err)
			}
			tflog.Info(ctx, DLPX+INFO+"Job result is "+job_status)
			if isJobTerminalFailure(job_status) {
				return diag.Errorf("[NOT OK] VDB-Update %s. JobId: %s / Error: %s", job_status, res.Job.GetId(), job_err)
			}
		}
	}

	// update tags
	if !d.Get("ignore_tag_changes").(bool) {
		apiRes, httpRes, err := client.VDBsAPI.GetVdbById(updateCtx, vdbId).Execute()
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
					httpRes, err := client.VDBsAPI.DeleteVdbTags(ctx, vdbId).Execute()
					if diags := apiErrorResponseHelper(ctx, nil, httpRes, err); diags != nil {
						return diags
					}
				}
				return resourceVdbRead(ctx, d, meta)
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
					tagDelResp, tagDelErr := client.VDBsAPI.DeleteVdbTags(updateCtx, vdbId).DeleteTag(deleteTag).Execute()
				if diags := apiErrorResponseHelper(ctx, nil, tagDelResp, tagDelErr); diags != nil {
					revertChanges(d, changedKeys)
					updateFailure = true
				}
			}
			// create tag
			if len(toTagArray(newTags)) != 0 {
				tflog.Info(ctx, "creating new tags")
				_, httpResp, tagCrtErr := client.VDBsAPI.CreateVdbTags(updateCtx, vdbId).TagsRequest(*dctapi.NewTagsRequest(toTagArray(newTags))).Execute()
				if diags := apiErrorResponseHelper(ctx, nil, httpResp, tagCrtErr); diags != nil {
					revertChanges(d, changedKeys)
					return diags
				}
			}
		}
	}

	if destructiveUpdate {
		if diags := enableVDB(updateCtx, client, vdbId); diags != nil {
			return diags //if failure should we enable
		}
	}

	return resourceVdbRead(updateCtx, d, meta)
}
func resourceVdbDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient).client

	vdbId := d.Id()

	deleteVdbParams := dctapi.NewDeleteVDBParametersWithDefaults()
	deleteVdbParams.SetForce(false)

	// respect resource delete timeout
	deleteCtx, deleteCancel := context.WithTimeout(ctx, d.Timeout(schema.TimeoutDelete))
	defer deleteCancel()

	res, httpRes, err := client.VDBsAPI.DeleteVdb(deleteCtx, vdbId).DeleteVDBParameters(*deleteVdbParams).Execute()

	// Check if the API call itself timed out
	if err != nil && deleteCtx.Err() == context.DeadlineExceeded {
		return diag.Errorf("VDB deletion API call timed out after %s. The request may still be processing on the DCT server. "+
			"Check the Delphix DCT UI or API to verify if a deletion job was created (VDB ID: %s). "+
			"If a job exists, wait for it to complete, then run 'terraform refresh' to verify the resource was deleted. "+
			"If the resource still exists in state, retry 'terraform destroy'. "+
			"To avoid timeouts, increase the timeout: timeouts { delete = \"60m\" }",
			d.Timeout(schema.TimeoutDelete), vdbId)
	}

	if diags := apiErrorResponseHelper(ctx, res, httpRes, err); diags != nil {
		return diags
	}

	// Check for nil res or Job to prevent crashes
	if res == nil || res.Job == nil {
		return diag.Errorf("VDB deletion failed: received nil response or job from API")
	}

	// Check if context timed out before polling
	if deleteCtx.Err() == context.DeadlineExceeded {
		return diag.Errorf("VDB deletion timed out after %s. The operation is still running on the DCT (Job ID: %s). "+
			"To resolve:\n"+
			"1. Wait for the job to complete (check Delphix DCT UI or API)\n"+
			"2. Run 'terraform refresh' to check if the resource was deleted\n"+
			"3. If still in state, retry 'terraform destroy'\n"+
			"To avoid timeouts, increase the timeout: timeouts { delete = \"60m\" }",
			d.Timeout(schema.TimeoutDelete), res.Job.GetId())
	}

	job_status, job_err := PollJobStatus(res.Job.GetId(), deleteCtx, client)
	if job_err != "" {
		// Check if the error is due to timeout
		if deleteCtx.Err() == context.DeadlineExceeded {
			return diag.Errorf("VDB deletion timed out after %s while polling job status. The operation is still running on the DCT (Job ID: %s). "+
				"To resolve:\n"+
				"1. Wait for the job to complete (check Delphix DCT UI or API)\n"+
				"2. Run 'terraform refresh' to check if the resource was deleted\n"+
				"3. If still in state, retry 'terraform destroy'\n"+
				"To avoid timeouts, increase the timeout: timeouts { delete = \"60m\" }",
				d.Timeout(schema.TimeoutDelete), res.Job.GetId())
		}
		tflog.Warn(ctx, DLPX+WARN+"Job Polling failed but continuing with deletion. Error : "+job_err)
	}
	tflog.Info(ctx, DLPX+INFO+"Job result is "+job_status)
	if isJobTerminalFailure(job_status) {
		return diag.Errorf("[NOT OK] VDB-Delete %s. JobId: %s / Error: %s", job_status, res.Job.GetId(), job_err)
	}

	_, diags := PollForObjectDeletion(deleteCtx, func() (interface{}, *http.Response, error) {
		return client.VDBsAPI.GetVdbById(deleteCtx, vdbId).Execute()
	})

	return diags
}
