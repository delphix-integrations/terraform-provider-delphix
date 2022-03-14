package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	openapi "github.com/Uddipaan-Hazarika/demo-go-sdk"
	//openapi "github.com/kurian87/dct-go-sdk"
)

func resourceEnvironment() *schema.Resource {
	return &schema.Resource{
		// Description is used by the doc genertor and language server.
		Description: "Provider Resource to add an environment to Delphix.",

		CreateContext: resourceEnvironmentCreate,
		ReadContext:   resourceEnvironmentRead,
		UpdateContext: resourceEnvironmentUpdate,
		DeleteContext: resourceEnvironmentDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"engine_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"os_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_cluster": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"cluster_home": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"staging_environment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"connector_port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"is_target": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ssh_port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"toolkit_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
			},
			"nfs_addresses": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ase_db_username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ase_db_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"java_home": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dsp_keystore_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dsp_keystore_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dsp_keystore_alias": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dsp_truststore_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dsp_truststore_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"hosts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostname": {
							Type:     schema.TypeString,
							Required: true,
						},
						"os_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"os_version": {
							Type:     schema.TypeString,
							Required: true,
						},
						"memory_size": {
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceEnvironmentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Function to add an environment in an engine.

	var diags diag.Diagnostics
	client := meta.(*apiClient).client

	createEnvParams := openapi.NewEnvironmentCreateParameters(
		int64(d.Get("engine_id").(int)),
		d.Get("os_name").(string),
		d.Get("hostname").(string),
	)

	//General
	if v, has_v := d.GetOk("username"); has_v {
		createEnvParams.SetUsername(v.(string))
	}
	if v, has_v := d.GetOk("password"); has_v {
		createEnvParams.SetPassword(v.(string))
	}
	if v, has_v := d.GetOk("name"); has_v {
		createEnvParams.SetName(v.(string))
	}
	if v, has_v := d.GetOk("toolkit_path"); has_v {
		createEnvParams.SetToolkitPath(v.(string))
	}
	if v, has_v := d.GetOk("ssh_port"); has_v {
		createEnvParams.SetSshPort(int64(v.(int)))
	}
	if v, has_v := d.GetOk("ase_db_username"); has_v {
		createEnvParams.SetAseDbUsername(v.(string))
	}
	if v, has_v := d.GetOk("ase_db_password"); has_v {
		createEnvParams.SetAseDbPassword(v.(string))
	}
	if v, has_v := d.GetOk("java_home"); has_v {
		createEnvParams.SetJavaHome(v.(string))
	}
	if v, has_v := d.GetOk("dsp_keystore_path"); has_v {
		createEnvParams.SetDspKeystorePath(v.(string))
	}
	if v, has_v := d.GetOk("dsp_keystore_password"); has_v {
		createEnvParams.SetDspKeystorePassword(v.(string))
	}
	if v, has_v := d.GetOk("dsp_keystore_alias"); has_v {
		createEnvParams.SetDspKeystoreAlias(v.(string))
	}
	if v, has_v := d.GetOk("dsp_truststore_path"); has_v {
		createEnvParams.SetDspTruststorePath(v.(string))
	}
	if v, has_v := d.GetOk("dsp_truststore_password"); has_v {
		createEnvParams.SetDspTruststorePassword(v.(string))
	}
	if v, has_v := d.GetOk("description"); has_v {
		createEnvParams.SetDescription(v.(string))
	}

	// // Clusters
	os_name := d.Get("os_name").(string)
	if v := d.Get("is_cluster"); v.(bool) {
		createEnvParams.SetIsCluster(v.(bool))
		if os_name == "WINDOWS" {
			createEnvParams.SetIsTarget(d.Get("is_target").(bool))
		}
	}
	if v, has_v := d.GetOk("cluster_home"); has_v {
		createEnvParams.SetClusterHome(v.(string))
	}

	// Win Specific
	if v, has_v := d.GetOk("connector_port"); has_v {
		createEnvParams.SetConnectorPort(int32(v.(int)))
	}

	if v, has_v := d.GetOk("staging_environment"); has_v {
		createEnvParams.SetStagingEnvironment(v.(string))
	}
	if v, has_v := d.GetOk("nfs_addresses"); has_v {
		createEnvParams.SetNfsAddresses(toStringArray(v))
	}

	apiReq := client.EnvironmentsApi.CreateEnvironments(ctx)
	apiRes, httpRes, err := apiReq.EnvironmentCreateParameters(*createEnvParams).Execute()

	if err != nil {
		resBody, httpErr := ResponseBodyToString(httpRes.Body)
		if httpErr != nil {
			log.Fatal(httpErr)
			return diag.FromErr(httpErr)
		}
		return diag.Errorf(resBody)
	}

	d.SetId(*apiRes.EnvironmentId)

	job_status, job_err := PollJobStatus(*apiRes.JobId, ctx, client)
	log.Printf("JobType: Env-Create / JobId: %s / Status: %s / Error: %s", *apiRes.JobId, job_status, job_err)

	if job_err != "" || job_status == "FAILED" {
		return diag.Errorf("JobType: Env-Create / JobId: %s / Status: %s / Error: ", *apiRes.JobId, job_status, job_err)
	}
	// Get environment info and store state.
	resourceEnvironmentRead(ctx, d, meta)
	return diags
}

func resourceEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*apiClient).client
	envId := d.Id()
	apiRes, httpRes, err := client.EnvironmentsApi.GetEnvironmentById(ctx, envId).Execute()

	if err != nil {
		resBody, httpErr := ResponseBodyToString(httpRes.Body)
		if httpErr != nil {
			return diag.FromErr(httpErr)
		}
		return diag.Errorf(resBody)
	}

	d.Set("namespace", apiRes.GetNamespace())
	d.Set("engine_id", apiRes.GetEngineId())
	d.Set("enabled", apiRes.GetEnabled())
	d.Set("hosts", apiRes.GetHosts())
	return diags
}

func resourceEnvironmentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("Not Implemented: resourceEnvironmentUpdate")
	var diags diag.Diagnostics
	return diags
}

func resourceEnvironmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	client := meta.(*apiClient).client
	envId := d.Id()
	apiRes, httpRes, err := client.EnvironmentsApi.DeleteEnvironment(ctx, envId).Execute()

	if err != nil {
		resBody, httpErr := ResponseBodyToString(httpRes.Body)
		if httpErr != nil {
			return diag.FromErr(httpErr)
		}
		return diag.Errorf(resBody)
	}

	job_status, job_err := PollJobStatus(*apiRes.JobId, ctx, client)
	if job_err != "" || job_status == "FAILED" {
		return diag.Errorf("JobType: Env-Delete / JobId: %s / Status: %s / Error: %s", *apiRes.JobId, job_status, job_err)
	}
	log.Printf("JobType: Env-Delete / JobId: %s / JobStatus: %s", *apiRes.JobId, job_status)
	return diags
}
