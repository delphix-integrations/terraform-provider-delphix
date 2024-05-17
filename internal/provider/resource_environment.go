package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-log/tflog"

	dctapi "github.com/delphix/dct-sdk-go/v14"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
				Type:     schema.TypeString,
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
				Optional: true,
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
			"cyberark_vault_query_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"use_kerberos_authentication": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_engine_public_key": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ase_db_vault": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ase_db_hashicorp_vault_engine": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ase_db_hashicorp_vault_secret_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ase_db_hashicorp_vault_username_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ase_db_hashicorp_vault_secret_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ase_db_cyberark_vault_query_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ase_db_use_kerberos_authentication": {
				Type:     schema.TypeBool,
				Optional: true,
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
			"id": {
				Type:     schema.TypeString,
				Computed: true,
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
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						// "ssh_port": {
						// 	Type:     schema.TypeInt,
						// 	Optional: true,
						// },
						// "toolkit_path": {
						// 	Type:     schema.TypeString,
						// 	Optional: true,
						// },
						// "processor_type": {
						// 	Type:     schema.TypeString,
						// 	Optional: true,
						// },
						// "timezone": {
						// 	Type:     schema.TypeString,
						// 	Optional: true,
						// },
						// "available": {
						// 	Type:     schema.TypeBool,
						// 	Optional: true,
						// },
						// "nfs_addresses": {
						// 	Type:     schema.TypeList,
						// 	Optional: true,
						// 	Elem: &schema.Schema{
						// 		Type: schema.TypeString,
						// 	},
						// },
					},
				},
			},
			"repositories": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"database_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"allow_provisioning": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"is_staging": {
							Type:     schema.TypeBool,
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

	createEnvParams := dctapi.NewEnvironmentCreateParameters(
		d.Get("engine_id").(string),
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
	if v, has_v := d.GetOk("vault"); has_v {
		createEnvParams.SetVault(v.(string))
	}
	if v, has_v := d.GetOk("hashicorp_vault_engine"); has_v {
		createEnvParams.SetHashicorpVaultEngine(v.(string))
	}
	if v, has_v := d.GetOk("hashicorp_vault_secret_path"); has_v {
		createEnvParams.SetHashicorpVaultSecretPath(v.(string))
	}
	if v, has_v := d.GetOk("hashicorp_vault_username_key"); has_v {
		createEnvParams.SetHashicorpVaultUsernameKey(v.(string))
	}
	if v, has_v := d.GetOk("hashicorp_vault_secret_key"); has_v {
		createEnvParams.SetHashicorpVaultSecretKey(v.(string))
	}
	if v, has_v := d.GetOk("cyberark_vault_query_string"); has_v {
		createEnvParams.SetCyberarkVaultQueryString(v.(string))
	}
	if v, has_v := d.GetOk("use_kerberos_authentication"); has_v {
		createEnvParams.SetUseKerberosAuthentication(v.(bool))
	}
	if v, has_v := d.GetOk("use_engine_public_key"); has_v {
		createEnvParams.SetUseEnginePublicKey(v.(bool))
	}
	if v, has_v := d.GetOk("ase_db_vault"); has_v {
		createEnvParams.SetAseDbVault(v.(string))
	}
	if v, has_v := d.GetOk("ase_db_hashicorp_vault_engine"); has_v {
		createEnvParams.SetAseDbHashicorpVaultEngine(v.(string))
	}
	if v, has_v := d.GetOk("ase_db_hashicorp_vault_secret_path"); has_v {
		createEnvParams.SetAseDbHashicorpVaultSecretPath(v.(string))
	}
	if v, has_v := d.GetOk("ase_db_hashicorp_vault_username_key"); has_v {
		createEnvParams.SetAseDbHashicorpVaultUsernameKey(v.(string))
	}
	if v, has_v := d.GetOk("ase_db_hashicorp_vault_secret_key"); has_v {
		createEnvParams.SetAseDbHashicorpVaultSecretKey(v.(string))
	}
	if v, has_v := d.GetOk("ase_db_cyberark_vault_query_string"); has_v {
		createEnvParams.SetAseDbCyberarkVaultQueryString(v.(string))
	}
	if v, has_v := d.GetOk("ase_db_use_kerberos_authentication"); has_v {
		createEnvParams.SetAseDbUseKerberosAuthentication(v.(bool))
	}

	// Clusters
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
	if v, has_v := d.GetOk("tags"); has_v {
		createEnvParams.SetTags(toTagArray(v))
	}

	apiReq := client.EnvironmentsApi.CreateEnvironment(ctx)
	apiRes, httpRes, err := apiReq.EnvironmentCreateParameters(*createEnvParams).Execute()

	if diags := apiErrorResponseHelper(ctx, apiRes, httpRes, err); diags != nil {
		return diags
	}

	d.SetId(apiRes.GetEnvironmentId())
	job_status, job_err := PollJobStatus(*apiRes.Job.Id, ctx, client)

	if job_err != "" {
		tflog.Error(ctx, DLPX+ERROR+"Job Polling failed but continuing with env creation. Error: "+job_err)
	}

	if isJobTerminalFailure(job_status) {
		d.SetId("")
		return diag.Errorf("[NOT OK] Env-Create %s. JobId: %s / Error: %s", job_status, *apiRes.Job.Id, job_err)
	}
	// Get environment info and store state.
	readDiags := resourceEnvironmentRead(ctx, d, meta)
	if readDiags.HasError() {
		return readDiags
	}
	return diags
}

func resourceEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient).client
	envId := d.Id()

	apiRes, diags := PollForObjectExistence(ctx, func() (interface{}, *http.Response, error) {
		return client.EnvironmentsApi.GetEnvironmentById(ctx, envId).Execute()
	})

	if diags != nil {
		_, diags := PollForObjectDeletion(ctx, func() (interface{}, *http.Response, error) {
			return client.EnvironmentsApi.GetEnvironmentById(ctx, envId).Execute()
		})
		if diags != nil {
			tflog.Error(ctx, DLPX+ERROR+"Error in polling of environment for deletion.")
		} else {
			tflog.Error(ctx, DLPX+ERROR+"Error Env-Read failed for EnvId: "+envId+". Removing from state file.")
			d.SetId("")
		}
		return nil
	}

	envRes, _ := apiRes.(*dctapi.Environment)
	//d.SetId(envRes.GetId())
	d.Set("id", envRes.GetId())
	d.Set("namespace", envRes.GetNamespace())
	d.Set("enabled", envRes.GetEnabled())
	d.Set("hosts", flattenHosts(envRes.GetHosts()))
	d.Set("repositories", flattenHostRepositories(envRes.GetRepositories()))
	return diags
}

func resourceEnvironmentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	// get the changed keys
	changedKeys := make([]string, 0, len(d.State().Attributes))
	for k := range d.State().Attributes {
		if d.HasChange(k) {
			changedKeys = append(changedKeys, k)
		}
	}

	if d.HasChanges(
		"engine_id",
		"os_name",
		"is_cluster",
		"connector_port",
		"is_target",
		"username",
		"password",
		"vault",
		"hashicorp_vault_engine",
		"hashicorp_vault_secret_path",
		"hashicorp_vault_username_key",
		"hashicorp_vault_secret_key",
		"cyberark_vault_query_string",
		"use_kerberos_authentication",
		"use_engine_public_key",
		"dsp_keystore_path",
		"dsp_keystore_password",
		"dsp_keystore_alias",
		"dsp_truststore_path",
		"dsp_truststore_password",
		"tags") {

		// revert and set the old value to the changed keys
		for _, key := range changedKeys {
			old, _ := d.GetChange(key)
			d.Set(key, old)
		}

		return diag.Errorf("cannot update one (or more) of the options changed. Please refer to provider documentation for updatable params.")
	}
	client := meta.(*apiClient).client
	environmentId := d.Get("id").(string)

	if d.HasChanges(
		"name",
		"cluster_home",
		"staging_environment",
		"ase_db_vault",
		"ase_db_hashicorp_vault_engine",
		"ase_db_hashicorp_vault_secret_path",
		"ase_db_hashicorp_vault_username_key",
		"ase_db_hashicorp_vault_secret_key",
		"ase_db_cyberark_vault_query_string",
		"ase_db_use_kerberos_authentication",
		"ase_db_username",
		"ase_db_password",
		"description") {
		envUpdateParam := dctapi.NewEnvironmentUpdateParameters()
		if d.HasChange("name") {
			if v, has_v := d.GetOk("name"); has_v {
				envUpdateParam.SetName(v.(string))
			}
		}
		if d.HasChange("cluster_home") {
			if v, has_v := d.GetOk("cluster_home"); has_v {
				envUpdateParam.SetClusterHome(v.(string))
			}
		}
		if d.HasChange("staging_environment") {
			if v, has_v := d.GetOk("staging_environment"); has_v {
				envUpdateParam.SetStagingEnvironment(v.(string))
			}
		}
		if d.HasChange("ase_db_vault") {
			if v, has_v := d.GetOk("ase_db_vault"); has_v {
				envUpdateParam.SetAseDbVault(v.(string))
			}
		}
		if d.HasChange("ase_db_hashicorp_vault_engine") {
			if v, has_v := d.GetOk("ase_db_hashicorp_vault_engine"); has_v {
				envUpdateParam.SetAseDbHashicorpVaultEngine(v.(string))
			}
		}
		if d.HasChange("ase_db_hashicorp_vault_secret_path") {
			if v, has_v := d.GetOk("ase_db_hashicorp_vault_secret_path"); has_v {
				envUpdateParam.SetAseDbHashicorpVaultSecretPath(v.(string))
			}
		}
		if d.HasChange("ase_db_hashicorp_vault_username_key") {
			if v, has_v := d.GetOk("ase_db_hashicorp_vault_username_key"); has_v {
				envUpdateParam.SetAseDbHashicorpVaultUsernameKey(v.(string))
			}
		}
		if d.HasChange("ase_db_hashicorp_vault_secret_key") {
			if v, has_v := d.GetOk("ase_db_hashicorp_vault_secret_key"); has_v {
				envUpdateParam.SetAseDbHashicorpVaultSecretKey(v.(string))
			}
		}
		if d.HasChange("ase_db_cyberark_vault_query_string") {
			if v, has_v := d.GetOk("ase_db_cyberark_vault_query_string"); has_v {
				envUpdateParam.SetAseDbCyberarkVaultQueryString(v.(string))
			}
		}
		if d.HasChange("ase_db_use_kerberos_authentication") {
			if v, has_v := d.GetOk("ase_db_use_kerberos_authentication"); has_v {
				envUpdateParam.SetAseDbUseKerberosAuthentication(v.(bool))
			}
		}
		if d.HasChange("ase_db_username") {
			if v, has_v := d.GetOk("ase_db_username"); has_v {
				envUpdateParam.SetAseDbUsername(v.(string))
			}
		}
		if d.HasChange("ase_db_password") {
			if v, has_v := d.GetOk("ase_db_password"); has_v {
				envUpdateParam.SetAseDbPassword(v.(string))
			}
		}
		if d.HasChange("description") {
			if v, has_v := d.GetOk("description"); has_v {
				envUpdateParam.SetDescription(v.(string))
			}
		}

		res, httpRes, err := client.EnvironmentsApi.UpdateEnvironment(ctx, environmentId).EnvironmentUpdateParameters(*envUpdateParam).Execute()
		if diags := apiErrorResponseHelper(ctx, res, httpRes, err); diags != nil {
			revertChanges(d, changedKeys)
			return diags
		}

		job_res, job_err := PollJobStatus(*res.Job.Id, ctx, client)
		if job_err != "" {
			tflog.Warn(ctx, DLPX+WARN+"Env Host Update Job Polling failed but continuing with update. Error: "+job_err)
		}
		tflog.Info(ctx, DLPX+INFO+"Job result is "+job_res)
		if job_res == Failed || job_res == Canceled || job_res == Abandoned {
			tflog.Error(ctx, DLPX+ERROR+"Job "+job_res+" "+*res.Job.Id+"!")
			revertChanges(d, changedKeys)
			return diag.Errorf("[NOT OK] Job %s %s with error %s", *res.Job.Id, job_res, job_err)
		}
	}

	if d.HasChanges("hostname", "java_home", "ssh_port", "toolkit_path", "nfs_addresses") { // call host update API
		hostsList := d.Get("hosts").([]interface{})
		firstHostMap, ok := hostsList[0].(map[string]interface{})
		if !ok {
			return diag.Errorf("Unexpected data type for first host element")
		}

		hostID, ok := firstHostMap["id"].(string)
		if !ok {
			return diag.Errorf("Error getting 'id' attribute from first host")
		}

		tflog.Info(ctx, DLPX+INFO+" hostID "+hostID)
		tflog.Info(ctx, DLPX+INFO+" environmentId "+environmentId)

		//1
		items, diags := filterVDBs(ctx, client, environmentId)
		if diags != nil {
			revertChanges(d, changedKeys)
			return diags
		}

		// 2
		var disableDiags diag.Diagnostics
		for _, item := range items {
			if diags := disableVDB(ctx, client, *item.Id); diags != nil {
				revertChanges(d, changedKeys)
				disableDiags = diags
			}
		}

		if disableDiags != nil { // if fialure attempt enable
			disableDiags = attemptEnableVDBs(ctx, client, items)
			if disableDiags != nil {
				return disableDiags
			}
		}

		//3
		sourceItems, diags := filterSources(ctx, client, environmentId)
		if diags != nil {
			revertChanges(d, changedKeys)
			return diags
		}

		var ids []string
		for _, item := range sourceItems {
			ids = append(ids, *item.Id)
		}

		//4
		var dsouceItems []dctapi.DSource
		if len(ids) > 0 {
			dsouceItems, diags = filterdSources(ctx, client, ids)
			if diags != nil {
				revertChanges(d, changedKeys)
				return diags
			}
		}

		//5
		for _, item := range dsouceItems {
			if diags := disabledSource(ctx, client, *item.Id); diags != nil {
				revertChanges(d, changedKeys)
				return diags //if failure should we revert
			}
		}

		hostUpdateParam := dctapi.NewHostUpdateParameters()
		if d.HasChange("hostname") {
			if v, has_v := d.GetOk("hostname"); has_v {
				hostUpdateParam.SetHostname(v.(string))
			}
		}
		if d.HasChange("java_home") {
			if v, has_v := d.GetOk("java_home"); has_v {
				hostUpdateParam.SetJavaHome(v.(string))
			}
		}
		if d.HasChange("ssh_port") {
			if v, has_v := d.GetOk("ssh_port"); has_v {
				hostUpdateParam.SetSshPort(v.(int64))
			}
		}
		if d.HasChange("toolkit_path") {
			if v, has_v := d.GetOk("toolkit_path"); has_v {
				hostUpdateParam.SetToolkitPath(v.(string))
			}
		}
		if d.HasChange("nfs_addresses") {
			if v, has_v := d.GetOk("nfs_addresses"); has_v {
				hostUpdateParam.SetNfsAddresses(toStringArray(v))
			}
		}

		hostUpdateRes, hostHttpRes, hostUpdateErr := client.EnvironmentsApi.UpdateHost(ctx, environmentId, hostID).HostUpdateParameters(*hostUpdateParam).Execute()
		if diags := apiErrorResponseHelper(ctx, hostUpdateRes, hostHttpRes, hostUpdateErr); diags != nil {
			revertChanges(d, changedKeys)
			return diags
		}

		job_res, job_err := PollJobStatus(*hostUpdateRes.Job.Id, ctx, client)
		if job_err != "" {
			tflog.Warn(ctx, DLPX+WARN+"Env Host Update Job Polling failed but continuing with update. Error: "+job_err)
		}
		tflog.Info(ctx, DLPX+INFO+"Job result is "+job_res)
		if job_res == Failed || job_res == Canceled || job_res == Abandoned {
			tflog.Error(ctx, DLPX+ERROR+"Job "+job_res+" "+*hostUpdateRes.Job.Id+"!")
			revertChanges(d, changedKeys)
			return diag.Errorf("[NOT OK] Job %s %s with error %s", *hostUpdateRes.Job.Id, job_res, job_err)
		}

		for _, item := range dsouceItems {
			if diags := enabledSource(ctx, client, *item.Id); diags != nil {
				return diags //if failure should we enable
			}
		}

		for _, item := range items {
			if diags := enableVDB(ctx, client, *item.Id); diags != nil {
				return diags //if failure should we enable
			}
		}
	}
	return diags
}

func attemptEnableVDBs(ctx context.Context, client *dctapi.APIClient, items []dctapi.VDB) diag.Diagnostics {
	var disableDiags diag.Diagnostics
	tflog.Info(ctx, DLPX+INFO+" vdb disable failed ...attempting enable for vdbs")
	for _, item := range items {
		if diags := enableVDB(ctx, client, *item.Id); diags != nil {
			disableDiags = append(disableDiags, diags...)
			return disableDiags
		}
	}
	return disableDiags
}

func revertChanges(d *schema.ResourceData, changedKeys []string) {
	for _, key := range changedKeys {
		old, _ := d.GetChange(key)
		d.Set(key, old)
	}
}

func resourceEnvironmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*apiClient).client
	envId := d.Id()

	apiRes, httpRes, err := client.EnvironmentsApi.DeleteEnvironment(ctx, envId).Execute()

	if diags := apiErrorResponseHelper(ctx, apiRes, httpRes, err); diags != nil {
		return diags
	}

	job_status, job_err := PollJobStatus(*apiRes.Job.Id, ctx, client)
	if job_err != "" {
		tflog.Error(ctx, DLPX+ERROR+"Job Polling failed but continuing with env deletion. Error: "+job_err)
	}
	if isJobTerminalFailure(job_status) {
		return diag.Errorf("[NOT OK] Env-Delete %s. JobId: %s / Error: %s", job_status, *apiRes.Job.Id, job_err)
	}
	_, diags := PollForObjectDeletion(ctx, func() (interface{}, *http.Response, error) {
		return client.EnvironmentsApi.GetEnvironmentById(ctx, envId).Execute()
	})

	return diags
}
