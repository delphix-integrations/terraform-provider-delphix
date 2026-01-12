package provider

import (
	"context"
	"net/http"
	"strings"
	"time"

	dctapi "github.com/delphix/dct-sdk-go/v25"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSource() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Resource for Source creation.",

		CreateContext: resourceDatabasePostgressqlCreate,
		ReadContext:   resourceDatabasePostgressqlRead,
		UpdateContext: resourceDatabasePostgressqlUpdate,
		DeleteContext: resourceDatabasePostgressqlDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"repository_value": {
				Type:     schema.TypeString,
				Required: true,
			},
			"environment_value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"engine_value": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// Output
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"environment_id": {
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
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"jdbc_connection_string": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"plugin_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"toolkit_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_dsource": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"repository": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"appdata_source_type": {
				Type:     schema.TypeString,
				Computed: true,
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
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceDatabasePostgressqlCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*apiClient).client

	sourceCreateParameters := dctapi.NewPostgresSourceCreateParametersWithDefaults()

	if v, has_v := d.GetOk("name"); has_v {
		sourceCreateParameters.SetName(v.(string))
	}
	if v, has_v := d.GetOk("repository_value"); has_v {
		sourceCreateParameters.SetRepositoryId(v.(string))
	}
	if v, has_v := d.GetOk("environment_value"); has_v {
		sourceCreateParameters.SetEnvironmentId(v.(string))
	}
	if v, has_v := d.GetOk("engine_value"); has_v {
		sourceCreateParameters.SetEngineId(v.(string))
	}

	// respect resource create timeout
	createCtx, createCancel := context.WithTimeout(ctx, d.Timeout(schema.TimeoutCreate))
	defer createCancel()

	req := client.SourcesAPI.CreatePostgresSource(createCtx)

	apiRes, httpRes, err := req.PostgresSourceCreateParameters(*sourceCreateParameters).Execute()
	
	// Check if the API call itself timed out
	if err != nil && createCtx.Err() == context.DeadlineExceeded {
		return diag.Errorf("PostgreSQL source creation API call timed out after %s. The request may still be processing on the DCT server. "+
			"Check the Delphix DCT UI or API to verify if a creation job was created. "+
			"If a job exists, wait for it to complete, then import the source. "+
			"To avoid timeouts, increase the timeout: timeouts { create = \"60m\" }", 
			d.Timeout(schema.TimeoutCreate))
	}
	
	if diags := apiErrorResponseHelper(ctx, apiRes, httpRes, err); diags != nil {
		return diags
	}

	d.SetId(apiRes.GetSourceId())

	job_res, job_err := PollJobStatus(apiRes.Job.GetId(), createCtx, client)
	if job_err != "" {
		tflog.Error(ctx, DLPX+ERROR+"Job Polling failed but continuing with Source creation. Error: "+job_err)
	}

	tflog.Info(ctx, DLPX+INFO+"Job result is "+job_res)

	if job_res == Failed || job_res == Canceled || job_res == Abandoned {
		d.SetId("")
		tflog.Error(ctx, DLPX+ERROR+"Job "+job_res+" "+apiRes.Job.GetId())
		return diag.Errorf("[NOT OK] Job %s %s with error %s", apiRes.Job.GetId(), job_res, job_err)
	}

	readDiags := resourceDatabasePostgressqlRead(createCtx, d, meta)

	if readDiags.HasError() {
		return readDiags
	}

	return diags
}

func resourceDatabasePostgressqlRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient).client

	source_id := d.Id()

	// use read timeout for polling during reads
	readCtx, readCancel := context.WithTimeout(ctx, d.Timeout(schema.TimeoutRead))
	defer readCancel()

	res, diags := PollForObjectExistence(readCtx, func() (interface{}, *http.Response, error) {
		return client.SourcesAPI.GetSourceById(readCtx, source_id).Execute()
	})

	if res == nil {
		tflog.Error(ctx, DLPX+ERROR+"PostgreSQL source not found: "+source_id+", removing from state. ")
		d.SetId("")
		return nil
	}

	if diags != nil {
		_, diags := PollForObjectDeletion(readCtx, func() (interface{}, *http.Response, error) {
			return client.SourcesAPI.GetSourceById(readCtx, source_id).Execute()
		})
		// This would imply error in poll for deletion so we just log and exit.
		if diags != nil {
			tflog.Error(ctx, DLPX+ERROR+"Error in polling of source for deletion.")
		} else {
			// diags will be nill in case of successful poll for deletion logic aka 404
			tflog.Error(ctx, DLPX+ERROR+"Error reading the source "+source_id+", removing from state.")
			d.SetId("")
		}

		return nil
	}

	result, ok := res.(*dctapi.Source)
	if !ok {
		return diag.Errorf("Error occured in type casting.")
	}

	repository_value := d.Get("repository_value").(string)

	if repository_value == "" {
		resEnv, httpRes, err := client.EnvironmentsAPI.GetEnvironmentById(ctx, result.GetEnvironmentId()).Execute()

		if diags := apiErrorResponseHelper(ctx, resEnv, httpRes, err); diags != nil {
			return diags
		}
		if result.GetRepository() != "" {
			for _, repo := range resEnv.Repositories {
				if strings.EqualFold(repo.GetId(), result.GetRepository()) {
					repository_value = repo.GetName()
				}
			}
		}
	}

	d.Set("id", result.GetId())
	d.Set("repository_value", repository_value)
	d.Set("environment_id", result.GetEnvironmentId())
	d.Set("database_type", result.GetDatabaseType())
	d.Set("name", result.GetName())
	d.Set("is_replica", result.GetIsReplica())
	d.Set("namespace_id", result.GetNamespaceId())
	d.Set("namespace_name", result.GetNamespaceName())
	d.Set("database_version", result.GetDatabaseVersion())
	d.Set("ip_address", result.GetIpAddress())
	d.Set("data_uuid", result.GetDataUuid())
	d.Set("fqdn", result.GetFqdn())
	d.Set("size", result.GetSize())
	d.Set("jdbc_connection_string", result.GetJdbcConnectionString())
	d.Set("plugin_version", result.GetPluginVersion())
	d.Set("toolkit_id", result.GetToolkitId())
	d.Set("is_dsource", result.GetIsDsource())
	d.Set("repository", result.GetRepository())
	d.Set("appdata_source_type", result.GetAppdataSourceType())

	return diags
}

func resourceDatabasePostgressqlUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	client := meta.(*apiClient).client
	updateSourceParam := dctapi.NewPostgresSourceUpdateParameters()

	// get the changed keys
	changedKeys := make([]string, 0, len(d.State().Attributes))
	for k := range d.State().Attributes {
		if d.HasChange(k) {
			changedKeys = append(changedKeys, k)
		}
	}

	if d.HasChanges(
		"repository_value",
		"environment_value",
		"engine_value") {

		// revert and set the old value to the changed keys
		for _, key := range changedKeys {
			old, _ := d.GetChange(key)
			d.Set(key, old)
		}

		return diag.Errorf("cannot update one (or more) of the options changed. Please refer to provider documentation for updatable params.")
	}

	if d.HasChange("name") {
		updateSourceParam.SetName(d.Get("name").(string))
	}

	// respect resource update timeout
	updateCtx, updateCancel := context.WithTimeout(ctx, d.Timeout(schema.TimeoutUpdate))
	defer updateCancel()

	res, httpRes, err := client.SourcesAPI.UpdatePostgresSourceById(updateCtx, d.Get("id").(string)).PostgresSourceUpdateParameters(*updateSourceParam).Execute()

	if diags := apiErrorResponseHelper(ctx, nil, httpRes, err); diags != nil {
		// revert and set the old value to the changed keys
		for _, key := range changedKeys {
			old, _ := d.GetChange(key)
			d.Set(key, old)
		}
		return diags
	}

	// Check if context timed out before polling
	if updateCtx.Err() == context.DeadlineExceeded {
		return diag.Errorf("PostgreSQL source update timed out after %s. The operation is still running on the DCT (Job ID: %s). "+
			"To resolve:\n"+
			"1. Wait for the job to complete (check Delphix DCT UI or API)\n"+
			"2. Run 'terraform refresh' to update the state with the actual resource details\n"+
			"3. If needed, revert changes manually or reapply the configuration\n"+
			"To avoid timeouts, increase the timeout: timeouts { update = \"60m\" }",
			d.Timeout(schema.TimeoutUpdate), res.Job.GetId())
	}

	job_status, job_err := PollJobStatus(res.Job.GetId(), updateCtx, client)
	if job_err != "" {
		// Check if the error is due to timeout
		if updateCtx.Err() == context.DeadlineExceeded {
			return diag.Errorf("PostgreSQL source update timed out after %s while polling job status. The operation is still running on the DCT (Job ID: %s). "+
				"To resolve:\n"+
				"1. Wait for the job to complete (check Delphix DCT UI or API)\n"+
				"2. Run 'terraform refresh' to update the state with the actual resource details\n"+
				"3. If needed, revert changes manually or reapply the configuration\n"+
				"To avoid timeouts, increase the timeout: timeouts { update = \"60m\" }",
				d.Timeout(schema.TimeoutUpdate), res.Job.GetId())
		}
		tflog.Warn(ctx, DLPX+WARN+"Source Update Job Polling failed but continuing with update. Error :"+job_err)
	}
	tflog.Info(ctx, DLPX+INFO+"Job result is "+job_status)
	if isJobTerminalFailure(job_status) {
		return diag.Errorf("[NOT OK] Source-Update %s. JobId: %s / Error: %s", job_status, res.Job.GetId(), job_err)
	}

	return diags
}

func resourceDatabasePostgressqlDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient).client

	source_id := d.Id()

	// respect resource delete timeout
	deleteCtx, deleteCancel := context.WithTimeout(ctx, d.Timeout(schema.TimeoutDelete))
	defer deleteCancel()

	res, httpRes, err := client.SourcesAPI.DeleteSource(deleteCtx, source_id).Execute()

	if diags := apiErrorResponseHelper(ctx, res, httpRes, err); diags != nil {
		return diags
	}

	// Check if context timed out before polling
	if deleteCtx.Err() == context.DeadlineExceeded {
		return diag.Errorf("PostgreSQL source deletion timed out after %s. The operation is still running on the DCT (Job ID: %s). "+
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
			return diag.Errorf("PostgreSQL source deletion timed out after %s while polling job status. The operation is still running on the DCT (Job ID: %s). "+
				"To resolve:\n"+
				"1. Wait for the job to complete (check Delphix DCT UI or API)\n"+
				"2. Run 'terraform refresh' to check if the resource was deleted\n"+
				"3. If still in state, retry 'terraform destroy'\n"+
				"To avoid timeouts, increase the timeout: timeouts { delete = \"60m\" }",
				d.Timeout(schema.TimeoutDelete), res.Job.GetId())
		}
		tflog.Warn(ctx, DLPX+WARN+"Job Polling failed but continuing with deletion. Error :"+job_err)
	}
	tflog.Info(ctx, DLPX+INFO+" Job result is "+job_status)
	if isJobTerminalFailure(job_status) {
		return diag.Errorf("[NOT OK] Source-Delete %s. JobId: %s / Error: %s", job_status, res.Job.GetId(), job_err)
	}

	_, diags := PollForObjectDeletion(deleteCtx, func() (interface{}, *http.Response, error) {
		return client.SourcesAPI.GetSourceById(deleteCtx, source_id).Execute()
	})

	return diags
}
