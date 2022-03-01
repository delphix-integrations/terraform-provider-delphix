package provider

import (
	"context"
	"log"
	"os"

	openapi "github.com/Uddipaan-Hazarika/demo-go-sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVdb() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider scaffolding.",

		CreateContext: resourceVdbCreate,
		ReadContext:   resourceVdbRead,
		UpdateContext: resourceVdbUpdate,
		DeleteContext: resourceVdbDelete,

		Schema: map[string]*schema.Schema{
			"vdb": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_select_repository": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"source_data_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"job_id": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"database_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
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
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"environment_id": {
							Type:     schema.TypeString,
							Computed: true,
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
						"database_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"truncate_log_on_checkpoint": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"username": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"password": {
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
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"shell": {
										Type:     schema.TypeString,
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
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"shell": {
										Type:     schema.TypeString,
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
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"shell": {
										Type:     schema.TypeString,
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
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"shell": {
										Type:     schema.TypeString,
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
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"shell": {
										Type:     schema.TypeString,
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
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"shell": {
										Type:     schema.TypeString,
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
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"shell": {
										Type:     schema.TypeString,
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
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"shell": {
										Type:     schema.TypeString,
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
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"shell": {
										Type:     schema.TypeString,
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
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"shell": {
										Type:     schema.TypeString,
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
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"shell": {
										Type:     schema.TypeString,
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
							Type:     schema.TypeString,
							Optional: true,
						},
						"custom_env_vars": {
							Type:     schema.TypeString,
							Optional: true,
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
					},
				},
			},
		},
	}
}

func resourceVdbCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	client := meta.(*apiClient).client

	vdbResource := d.Get("vdb").([]interface{})
	vdb := vdbResource[0].(map[string]interface{})

	provisionVDBBySnapshotParameters := openapi.NewProvisionVDBBySnapshotParameters()
	provisionVDBBySnapshotParameters.SetSourceDataId(vdb["source_data_id"].(string))
	provisionVDBBySnapshotParameters.SetAutoSelectRepository(vdb["auto_select_repository"].(bool))

	req := client.VDBsApi.ProvisionVdbBySnapshot(context.WithValue(context.Background(), openapi.ContextAPIKeys, meta.(*apiClient).apiKeyMap))

	res, httpRes, err := req.ProvisionVDBBySnapshotParameters(*provisionVDBBySnapshotParameters).Execute()

	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}

	log.Print(&res)
	log.Print(httpRes)

	//d.Set("id", res.Vdb.Id)
	d.SetId(*res.Vdb.Id)
	job_res := PollJobStatus(*res.JobId, context.WithValue(context.Background(), openapi.ContextAPIKeys, meta.(*apiClient).apiKeyMap), client)
	log.Print(job_res)
	resourceVdbRead(ctx, d, meta)
	return diags

}

func resourceVdbRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*apiClient).client
	vdbResource := d.Get("vdb").([]interface{})
	vdb := vdbResource[0].(map[string]interface{})

	var diags diag.Diagnostics

	vdbId := d.Id()

	res, httpRes, err := client.VDBsApi.GetVdbById(context.WithValue(context.Background(), openapi.ContextAPIKeys, meta.(*apiClient).apiKeyMap), vdbId).Execute()

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	log.Print(res.GetName())
	log.Print(httpRes)

	flatRes := make(map[string]interface{})
	flatRes["database_type"] = res.GetDatabaseType()
	flatRes["name"] = res.GetName()
	flatRes["database_version"] = res.GetDatabaseVersion()
	flatRes["engine_id"] = res.GetEngineId()
	flatRes["status"] = res.GetStatus()
	flatRes["environment_id"] = res.GetEnvironmentId()
	flatRes["ip_address"] = res.GetIpAddress()
	flatRes["fqdn"] = res.GetFqdn()
	flatRes["parent_id"] = res.GetParentId()
	flatRes["group_name"] = res.GetGroupName()
	flatRes["creation_date"] = res.GetCreationDate().String()
	flatRes["id"] = res.GetId()
	flatRes["source_data_id"] = vdb["source_data_id"].(string)
	flatRes["auto_select_repository"] = vdb["auto_select_repository"].(bool)

	d.Set("vdb", []interface{}{flatRes})

	d.SetId(*res.Id)
	return diags
}

func resourceVdbUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return diag.Errorf("not implemented")
}

func resourceVdbDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*apiClient).client

	var diags diag.Diagnostics

	vdbId := d.Id()

	deleteVdbParams := openapi.NewDeleteVDBParametersWithDefaults()
	deleteVdbParams.SetForce(false)

	res, httpRes, err := client.VDBsApi.DeleteVdb(context.WithValue(context.Background(), openapi.ContextAPIKeys, meta.(*apiClient).apiKeyMap), vdbId).DeleteVDBParameters(*deleteVdbParams).Execute()

	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}

	log.Print(&res)
	log.Print(httpRes)

	job_res := PollJobStatus(*res.JobId, context.WithValue(context.Background(), openapi.ContextAPIKeys, meta.(*apiClient).apiKeyMap), client)
	log.Print(job_res)

	return diags
}
