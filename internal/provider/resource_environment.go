package provider

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"

	dctapi "github.com/delphix/dct-sdk-go/v21"
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
			"is_cluster": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"cluster_home": {
				Type:     schema.TypeString,
				Optional: true,
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
			"ase_db_username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ase_db_password": {
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
			"database_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oracle_base": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bits": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"allow_provisioning": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_staging": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_replica": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_windows_target": {
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
			"namespace_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"namespace_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"hosts": {
				Type:     schema.TypeList,
				Optional: true,
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
						"ssh_port": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"toolkit_path": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"oracle_tde_keystores_root_path": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"java_home": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"nfs_addresses": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"os_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"memory_size": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"processor_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"timezone": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"available": {
							Type:     schema.TypeBool,
							Computed: true,
						},
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
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"database_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"oracle_base": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bits": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"allow_provisioning": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_staging": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"user_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceEnvironmentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Function to add an environment in an engine.

	var diags diag.Diagnostics
	client := meta.(*apiClient).client

	var hostname, os_name, toolkit_path, java_home string
	var ssh_port int
	var nfs_addresses interface{}
	// process hosts
	if v, has_v := d.GetOk("hosts"); has_v {
		hosts := v.([]interface{})
		if len(hosts) > 0 {
			host := hosts[0].(map[string]interface{}) // Cast host to a map
			os_name = host["os_name"].(string)
			// oracle_tde_keystores_root_path = host["oracle_tde_keystores_root_path"].(string)
			hostname = host["hostname"].(string)
			toolkit_path = host["toolkit_path"].(string)
			if val, ok := host["ssh_port"]; ok {
				ssh_port = val.(int)
			}
			java_home = host["java_home"].(string)
			nfs_addresses = host["nfs_addresses"]
		}
	}

	createEnvParams := dctapi.NewEnvironmentCreateParameters(
		d.Get("engine_id").(string),
		os_name,
		hostname,
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
	if toolkit_path != "" {
		createEnvParams.SetToolkitPath(toolkit_path)
	}
	if ssh_port != 0 {
		createEnvParams.SetSshPort(int64(ssh_port))
	}
	if java_home != "" {
		createEnvParams.SetJavaHome(java_home)
	}
	if v, has_v := d.GetOk("ase_db_username"); has_v {
		createEnvParams.SetAseDbUsername(v.(string))
	}
	if v, has_v := d.GetOk("ase_db_password"); has_v {
		createEnvParams.SetAseDbPassword(v.(string))
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
	if nfs_addresses != nil {
		createEnvParams.SetNfsAddresses(toStringArray(nfs_addresses))
	}
	if v, has_v := d.GetOk("tags"); has_v {
		createEnvParams.SetTags(toTagArray(v))
	}

	apiReq := client.EnvironmentsAPI.CreateEnvironment(ctx)
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
		return client.EnvironmentsAPI.GetEnvironmentById(ctx, envId).Execute()
	})

	if diags != nil {
		_, diags := PollForObjectDeletion(ctx, func() (interface{}, *http.Response, error) {
			return client.EnvironmentsAPI.GetEnvironmentById(ctx, envId).Execute()
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
	d.Set("name", envRes.GetName())
	d.Set("id", envRes.GetId())
	d.Set("namespace", envRes.GetNamespace())
	d.Set("namespace_name", envRes.GetNamespaceName())
	d.Set("namespace_id", envRes.GetNamespaceId())
	d.Set("is_replica", envRes.GetIsReplica())
	d.Set("engine_id", envRes.GetEngineId())
	d.Set("is_cluster", envRes.GetIsCluster())
	d.Set("enabled", envRes.GetEnabled())
	tflog.Info(ctx, "is WindowsTarget"+strconv.FormatBool(envRes.GetIsWindowsTarget()))
	d.Set("is_windows_target", envRes.GetIsWindowsTarget())
	d.Set("staging_environment", envRes.GetStagingEnvironment())
	d.Set("cluster_home", envRes.GetClusterHome())
	d.Set("hosts", flattenHosts(envRes.GetHosts()))
	d.Set("repositories", flattenHostRepositories(envRes.GetRepositories()))
	d.Set("tags", flattenTags(envRes.Tags))

	if user_ref, has_user_ref := d.GetOk("user_ref"); has_user_ref {
		// this is set from update
		tflog.Info(ctx, "~~~~~~~~Setting username in state(read)")
		resUserList, httpResUserList, errUserList := client.EnvironmentsAPI.ListEnvironmentUsers(ctx, envId).Execute()
		if diags := apiErrorResponseHelper(ctx, resUserList, httpResUserList, errUserList); diags != nil {
			return diag.Errorf("unable to retrieve user list")
		}

		for _, users := range resUserList.GetUsers() {
			if strings.EqualFold(users.GetUserRef(), user_ref.(string)) {
				d.Set("username", users.GetUsername())
			}
		}
	}

	return diags
}

func resourceEnvironmentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	// get the changed keys
	changedKeys := make([]string, 0, len(d.State().Attributes))
	for k := range d.State().Attributes {
		if strings.Contains(k, "tags") { // this is because the changed keys are of the form tag.0.keydi
			k = "tags"
		}
		if strings.Contains(k, "hosts") { // this is because the changed keys are of the form tag.0.keydi
			k = "hosts"
		}
		if d.HasChange(k) {
			tflog.Info(ctx, ">>>>>@@@<<<<<<"+k)
			changedKeys = append(changedKeys, k)
		}
	}
	for _, ck := range changedKeys {
		tflog.Info(ctx, "!!!!!!!!!!!!!!!"+ck)
	}
	client := meta.(*apiClient).client
	environmentId := d.Get("id").(string)
	var updateFailure, destructiveUpdate bool = false, false
	var nonUpdatableField []string
	var dsourceItems []dctapi.DSource
	var vdbs []dctapi.VDB
	var vdbDiags, dsourceDiags diag.Diagnostics
	var disableDsourceFailure bool = false
	// if changedKeys contains non updatable field set a flag
	for _, key := range changedKeys {
		tflog.Info(ctx, "!!!!!!!!!!!!!!!"+key)
		if !updatableEnvKeys[key] {
			updateFailure = true
			tflog.Info(ctx, ">>>>>!!!<<<<<<"+key)
			nonUpdatableField = append(nonUpdatableField, key)
		}
	}

	if updateFailure {
		tflog.Info(ctx, "######updatefailure")
		revertChanges(d, changedKeys)
		return diag.Errorf("cannot update options %v. Please refer to provider documentation for updatable params.", nonUpdatableField)
	}
	// find if destructive update
	for _, key := range changedKeys {
		if isDestructiveEnvUpdate[key] {
			tflog.Info(ctx, "######isDestructiveUpdate"+key)
			destructiveUpdate = true
		}
	}

	if destructiveUpdate {
		// get dsources and vdbs
		vdbs, vdbDiags = filterVDBs(ctx, client, environmentId)
		tflog.Info(ctx, "######vdbs")
		if vdbDiags.HasError() {
			revertChanges(d, changedKeys)
			return vdbDiags
		}

		// get sources to get dsources
		sources, sourceDiag := filterSources(ctx, client, environmentId)
		tflog.Info(ctx, "######sources")
		if sourceDiag.HasError() {
			revertChanges(d, changedKeys)
			return sourceDiag
		}
		var sourceIds []string
		for _, item := range sources {
			sourceIds = append(sourceIds, *item.Id)
		}

		// retrieve dsources from source list

		if len(sourceIds) > 0 {
			tflog.Info(ctx, "######DSource")
			dsourceItems, dsourceDiags = filterdSources(ctx, client, sourceIds)
			if dsourceDiags != nil {
				revertChanges(d, changedKeys)
				return dsourceDiags
			}
		}

		// disable vdb
		for _, item := range vdbs {
			tflog.Info(ctx, "######disableVDB")
			if diags := disableVDB(ctx, client, *item.Id); diags != nil {
				tflog.Error(ctx, "failure in disabling vdbs")
				//disableVdbFailure = true
				revertChanges(d, changedKeys)
				return diags
			}
		}

		// disable dsources
		for _, item := range dsourceItems {
			tflog.Info(ctx, "######disabledSource")
			if diags := disabledSource(ctx, client, *item.Id); diags != nil {
				tflog.Error(ctx, "failure in disabling Dsources")
				disableDsourceFailure = true
			}
		}
		if disableDsourceFailure {
			//enable back vdbs and return
			for _, item := range vdbs {
				tflog.Info(ctx, "######disableDsourceFailure")
				if diags := enableVDB(ctx, client, *item.Id); diags != nil {
					revertChanges(d, changedKeys)
					return diags
				}
			}
		}
	}
	var failureEvents []string
	// if no disable failures, proceed to updateen
	if d.HasChanges(
		"name",
		"cluster_home",
		"description",
	) {
		// env update
		tflog.Info(ctx, "env")
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
		if d.HasChange("description") {
			if v, has_v := d.GetOk("description"); has_v {
				envUpdateParam.SetDescription(v.(string))
			}
		}

		res, httpRes, err := client.EnvironmentsAPI.UpdateEnvironment(ctx, environmentId).EnvironmentUpdateParameters(*envUpdateParam).Execute()
		if diags := apiErrorResponseHelper(ctx, res, httpRes, err); diags != nil {
			revertChanges(d, changedKeys)
			updateFailure = true
			failureEvents = append(failureEvents, httpRes.Body.Close().Error())
		}

		job_res, job_err := PollJobStatus(*res.Job.Id, ctx, client)
		if job_err != "" {
			tflog.Warn(ctx, DLPX+WARN+"Env Host Update Job Polling failed but continuing with update. Error: "+job_err)
		}
		tflog.Info(ctx, DLPX+INFO+"Job result is "+job_res)
		if job_res == Failed || job_res == Canceled || job_res == Abandoned {
			tflog.Error(ctx, DLPX+ERROR+"Job "+job_res+" "+*res.Job.Id+"!")
			revertChanges(d, changedKeys)
			updateFailure = true
			failureEvents = append(failureEvents, job_err)
			// return diag.Errorf("[NOT OK] Job %s %s with error %s", *res.Job.Id, job_res, job_err)
		}
	}
	if d.HasChanges(
		"username",
		"password",
	) {
		tflog.Info(ctx, "envUser")
		// envUser Update
		envUserUpdateParam := dctapi.NewEnvironmentUserParams()
		if d.HasChange("username") || d.HasChange("password") {
			if v, has_v := d.GetOk("username"); has_v {
				envUserUpdateParam.SetUsername(v.(string))
			}
			if v, has_v := d.GetOk("password"); has_v {
				envUserUpdateParam.SetPassword(v.(string))
			}
		}
		// get the user ref
		tflog.Info(ctx, "~~~~~~~~Getting the userlist")
		resUserList, httpResUserList, errUserList := client.EnvironmentsAPI.ListEnvironmentUsers(ctx, environmentId).Execute()
		if diags := apiErrorResponseHelper(ctx, resUserList, httpResUserList, errUserList); diags != nil {
			revertChanges(d, changedKeys)
			return diags
		}

		var user_ref string

		username, _ := d.GetChange("username")
		for _, users := range resUserList.GetUsers() {
			tflog.Info(ctx, "~~~~~~~~Getting the users"+*users.Username)
			if strings.EqualFold(users.GetUsername(), username.(string)) {
				user_ref = users.GetUserRef()
				break
			}
		}
		if user_ref == "" {
			revertChanges(d, changedKeys)
			return diag.Errorf("no matching user found in the environment list to update")
		}

		// this is to propagate the value to read call which is defined at the end.
		// we will use the user_ref to filter from the list of users in the env
		tflog.Info(ctx, "~~~~~~~~Setting the user_ref"+user_ref)
		d.Set("user_ref", user_ref)

		tflog.Info(ctx, "~~~~~~~~Updating the user"+user_ref)
		resEnvUser, httpResEnvUser, errEnvUser := client.EnvironmentsAPI.UpdateEnvironmentUser(ctx, environmentId, user_ref).EnvironmentUserParams(*envUserUpdateParam).Execute()
		if diags := apiErrorResponseHelper(ctx, resEnvUser, httpResEnvUser, errEnvUser); diags != nil {
			revertChanges(d, changedKeys)
			updateFailure = true
			failureEvents = append(failureEvents, httpResEnvUser.Body.Close().Error())
		}

		job_res, job_err := PollJobStatus(*resEnvUser.Job.Id, ctx, client)
		if job_err != "" {
			tflog.Warn(ctx, DLPX+WARN+"Env User Update Job Polling failed but continuing with update. Error: "+job_err)
		}
		tflog.Info(ctx, DLPX+INFO+"Job result is "+job_res)
		if job_res == Failed || job_res == Canceled || job_res == Abandoned {
			tflog.Error(ctx, DLPX+ERROR+"Job "+job_res+" "+*resEnvUser.Job.Id+"!")
			revertChanges(d, changedKeys)
			updateFailure = true
			failureEvents = append(failureEvents, job_err)
			// return diag.Errorf("[NOT OK] Job %s %s with error %s", *resEnvUser.Job.Id, job_res, job_err)
		}
	}
	if d.HasChanges(
		"hosts",
		"connector_port",
		// "java_home",
		// "hostname",
		// "ssh_port",
		// "toolkit_path",
		// "nfs_addresses",
		// "oracle_tde_keystores_root_path",
	) {
		tflog.Info(ctx, "hosts")
		// host update
		var hostId string

		// get changes
		oldHosts, newHosts := d.GetChange("hosts")

		// signifies the hostname that will be updated
		oldHost := oldHosts.([]interface{})
		oldHostName := oldHost[0].(map[string]interface{})["hostname"].(string)

		// retrieving new params for the update
		newHost := newHosts.([]interface{})
		newHostName := newHost[0].(map[string]interface{})["hostname"].(string)
		newSshPort := int64(newHost[0].(map[string]interface{})["ssh_port"].(int))
		newToolkitPath := newHost[0].(map[string]interface{})["toolkit_path"].(string)
		newJavaHome := newHost[0].(map[string]interface{})["java_home"].(string)
		newNfsAddress := newHost[0].(map[string]interface{})["nfs_addresses"]

		// get the hosts list
		hostsList := d.Get("hosts").([]interface{})

		// retrieve the hostId corresponding to the old host name (that will be updated)
		for _, host := range hostsList {
			if oldHostName == host.(map[string]interface{})["hostname"].(string) {
				hostId = host.(map[string]interface{})["id"].(string)
				tflog.Info(ctx, "<>>>>>><<<<<<<>>>>>> hostsId: "+hostId)
				break
			} else {
				// if not found, proceed with enable and finally display the failure events
				updateFailure = true
				failureEvents = append(failureEvents, "No hostname %s found to update", oldHostName)
			}
		}

		if !updateFailure {
			tflog.Info(ctx, DLPX+INFO+" hostID "+hostId)
			tflog.Info(ctx, DLPX+INFO+" environmentId "+environmentId)

			hostUpdateParam := dctapi.NewHostUpdateParameters()
			if d.HasChange("connector_port") {
				if v, has_v := d.GetOk("connector_port"); has_v {
					hostUpdateParam.SetConnectorPort(v.(int32))
				}
			}
			if newJavaHome != "" {
				hostUpdateParam.SetJavaHome(newJavaHome)
			}
			if newHostName != "" {
				hostUpdateParam.SetHostname(newHostName)
			}
			if newSshPort != 0 {
				hostUpdateParam.SetSshPort(newSshPort)
			}
			if newToolkitPath != "" {
				hostUpdateParam.SetToolkitPath(newToolkitPath)
			}
			if newNfsAddress != nil {
				hostUpdateParam.SetNfsAddresses(toStringArray(newNfsAddress))
			}
			// if d.HasChange("oracle_tde_keystores_root_path") {
			// 	if v, has_v := d.GetOk("oracle_tde_keystores_root_path"); has_v {
			// 		hostUpdateParam.SetOracleTdeKeystoresRootPath(v.(string))
			// 	}
			// }

			hostUpdateRes, hostHttpRes, hostUpdateErr := client.EnvironmentsAPI.UpdateHost(ctx, environmentId, hostId).HostUpdateParameters(*hostUpdateParam).Execute()
			if diags := apiErrorResponseHelper(ctx, hostUpdateRes, hostHttpRes, hostUpdateErr); diags != nil {
				revertChanges(d, changedKeys)
				updateFailure = true
				failureEvents = append(failureEvents, hostHttpRes.Body.Close().Error())
			}

			job_res, job_err := PollJobStatus(*hostUpdateRes.Job.Id, ctx, client)
			if job_err != "" {
				tflog.Warn(ctx, DLPX+WARN+"Env Host Update Job Polling failed but continuing with update. Error: "+job_err)
			}
			tflog.Info(ctx, DLPX+INFO+"Job result is "+job_res)
			if job_res == Failed || job_res == Canceled || job_res == Abandoned {
				tflog.Error(ctx, DLPX+ERROR+"Job "+job_res+" "+*hostUpdateRes.Job.Id+"!")
				revertChanges(d, changedKeys)
				updateFailure = true
				failureEvents = append(failureEvents, job_err)
				// return diag.Errorf("[NOT OK] Job %s %s with error %s", *hostUpdateRes.Job.Id, job_res, job_err)
			}
		}

	}
	if d.HasChanges(
		"tags",
	) { // tags update
		tflog.Info(ctx, ">>>>>>>>>>>>tags")
		if d.HasChange("tags") {
			// delete old tag
			tflog.Info(ctx, ">>>>>>>>>>>>delete tags")
			oldTag, newTag := d.GetChange("tags")
			if len(toTagArray(oldTag)) != 0 {
				tflog.Info(ctx, "&&&&&&&&&&&>>>>>>>>>>>>delete tags"+toTagArray(oldTag)[0].GetKey()+" "+toTagArray(oldTag)[0].GetValue())
				deleteTag := *dctapi.NewDeleteTag()
				tagDelResp, tagDelErr := client.EnvironmentsAPI.DeleteEnvironmentTags(ctx, environmentId).DeleteTag(deleteTag).Execute()
				tflog.Info(ctx, ">>DELETE TAG RESP: "+tagDelResp.Status)
				if diags := apiErrorResponseHelper(ctx, nil, tagDelResp, tagDelErr); diags != nil {
					revertChanges(d, changedKeys)
					updateFailure = true
					failureEvents = append(failureEvents, tagDelResp.Body.Close().Error())
				}
			}
			// create tag
			if len(toTagArray(newTag)) != 0 {
				tflog.Info(ctx, ">>>>>>>>>>>>create tags")
				_, httpResp, tagCrtErr := client.EnvironmentsAPI.CreateEnvironmentTags(ctx, environmentId).TagsRequest(*dctapi.NewTagsRequest(toTagArray(newTag))).Execute()
				if diags := apiErrorResponseHelper(ctx, nil, httpResp, tagCrtErr); diags != nil {
					revertChanges(d, changedKeys)
					return diags
				}
			}
		}

	}

	if destructiveUpdate {
		// enable Dsources back
		for _, item := range dsourceItems {
			if diags := enableDsource(ctx, client, *item.Id); diags != nil {
				return diags
			}
		}
		// enable VDB back
		for _, item := range vdbs {
			if diags := enableVDB(ctx, client, *item.Id); diags != nil {
				return diags
			}
		}
	}

	// return the error back
	if updateFailure {
		tflog.Error(ctx, "??????ERPRORORRRRRRRR???")
		return diag.Errorf("[NOT OK] Update failed with error %s", failureEvents)
	}
	// ========================Old=============================
	// if d.HasChanges(
	// 	"engine_id",
	// 	"id",
	// 	"namespace",
	// 	"enabled",
	// 	"os_name",
	// 	"os_version",
	// 	"memory_size",
	// 	"processor_type",
	// 	"timezone",
	// 	"available",
	// 	"is_cluster",
	// 	"connector_port",
	// 	"is_target",
	// 	"tags") {

	// 	// revert and set the old value to the changed keys
	// 	for _, key := range changedKeys {
	// 		old, _ := d.GetChange(key)
	// 		d.Set(key, old)
	// 	}

	// 	return diag.Errorf("cannot update one (or more) of the options changed. Please refer to provider documentation for updatable params.")
	// }
	// client := meta.(*apiClient).client
	// environmentId := d.Get("id").(string)

	// if d.HasChanges(
	// 	"allow_provisioning",
	// 	"is_staging",
	// ) {
	// 	// repo update
	// }

	// if d.HasChanges(
	// 	"username",
	// 	"password",
	// ) {
	// 	// envuser update
	// }

	// if d.HasChanges(
	// 	"name",
	// 	"cluster_home",
	// 	"description") {
	// 	// env update
	// 	envUpdateParam := dctapi.NewEnvironmentUpdateParameters()
	// 	if d.HasChange("name") {
	// 		if v, has_v := d.GetOk("name"); has_v {
	// 			envUpdateParam.SetName(v.(string))
	// 		}
	// 	}
	// 	if d.HasChange("cluster_home") {
	// 		if v, has_v := d.GetOk("cluster_home"); has_v {
	// 			envUpdateParam.SetClusterHome(v.(string))
	// 		}
	// 	}
	// 	if d.HasChange("description") {
	// 		if v, has_v := d.GetOk("description"); has_v {
	// 			envUpdateParam.SetDescription(v.(string))
	// 		}
	// 	}

	// 	res, httpRes, err := client.EnvironmentsAPI.UpdateEnvironment(ctx, environmentId).EnvironmentUpdateParameters(*envUpdateParam).Execute()
	// 	if diags := apiErrorResponseHelper(ctx, res, httpRes, err); diags != nil {
	// 		revertChanges(d, changedKeys)
	// 		return diags
	// 	}

	// 	job_res, job_err := PollJobStatus(*res.Job.Id, ctx, client)
	// 	if job_err != "" {
	// 		tflog.Warn(ctx, DLPX+WARN+"Env Host Update Job Polling failed but continuing with update. Error: "+job_err)
	// 	}
	// 	tflog.Info(ctx, DLPX+INFO+"Job result is "+job_res)
	// 	if job_res == Failed || job_res == Canceled || job_res == Abandoned {
	// 		tflog.Error(ctx, DLPX+ERROR+"Job "+job_res+" "+*res.Job.Id+"!")
	// 		revertChanges(d, changedKeys)
	// 		return diag.Errorf("[NOT OK] Job %s %s with error %s", *res.Job.Id, job_res, job_err)
	// 	}
	// }

	// if d.HasChanges(
	// 	"hostname",
	// 	"java_home",
	// 	"ssh_port",
	// 	"toolkit_path",
	// 	"nfs_addresses",
	// 	"connector_port",
	// ) { // call host update API
	// 	hostsList := d.Get("hosts").([]interface{})
	// 	firstHostMap, ok := hostsList[0].(map[string]interface{})
	// 	if !ok {
	// 		return diag.Errorf("Unexpected data type for first host element")
	// 	}

	// 	hostID, ok := firstHostMap["id"].(string)
	// 	if !ok {
	// 		return diag.Errorf("Error getting 'id' attribute from first host")
	// 	}

	// 	tflog.Info(ctx, DLPX+INFO+" hostID "+hostID)
	// 	tflog.Info(ctx, DLPX+INFO+" environmentId "+environmentId)

	// 	//1
	// 	items, diags := filterVDBs(ctx, client, environmentId)
	// 	if diags != nil {
	// 		revertChanges(d, changedKeys)
	// 		return diags
	// 	}

	// 	// 2
	// 	var disableDiags diag.Diagnostics
	// 	for _, item := range items {
	// 		if diags := disableVDB(ctx, client, *item.Id); diags != nil {
	// 			revertChanges(d, changedKeys)
	// 			disableDiags = diags
	// 		}
	// 	}

	// 	if disableDiags != nil { // if fialure attempt enable
	// 		disableDiags = attemptEnableVDBs(ctx, client, items)
	// 		if disableDiags != nil {
	// 			return disableDiags
	// 		}
	// 	}

	// 	//3
	// 	sourceItems, diags := filterSources(ctx, client, environmentId)
	// 	if diags != nil {
	// 		revertChanges(d, changedKeys)
	// 		return diags
	// 	}

	// 	var ids []string
	// 	for _, item := range sourceItems {
	// 		ids = append(ids, *item.Id)
	// 	}

	// 	//4
	// 	var dsouceItems []dctapi.DSource
	// 	if len(ids) > 0 {
	// 		dsouceItems, diags = filterdSources(ctx, client, ids)
	// 		if diags != nil {
	// 			revertChanges(d, changedKeys)
	// 			return diags
	// 		}
	// 	}

	// 	//5
	// 	for _, item := range dsouceItems {
	// 		if diags := disabledSource(ctx, client, *item.Id); diags != nil {
	// 			revertChanges(d, changedKeys)
	// 			return diags //if failure should we revert
	// 		}
	// 	}

	// 	hostUpdateParam := dctapi.NewHostUpdateParameters()
	// 	if d.HasChange("hostname") {
	// 		if v, has_v := d.GetOk("hostname"); has_v {
	// 			hostUpdateParam.SetHostname(v.(string))
	// 		}
	// 	}
	// 	if d.HasChange("java_home") {
	// 		if v, has_v := d.GetOk("java_home"); has_v {
	// 			hostUpdateParam.SetJavaHome(v.(string))
	// 		}
	// 	}
	// 	if d.HasChange("ssh_port") {
	// 		if v, has_v := d.GetOk("ssh_port"); has_v {
	// 			hostUpdateParam.SetSshPort(v.(int64))
	// 		}
	// 	}
	// 	if d.HasChange("toolkit_path") {
	// 		if v, has_v := d.GetOk("toolkit_path"); has_v {
	// 			hostUpdateParam.SetToolkitPath(v.(string))
	// 		}
	// 	}
	// 	if d.HasChange("nfs_addresses") {
	// 		if v, has_v := d.GetOk("nfs_addresses"); has_v {
	// 			hostUpdateParam.SetNfsAddresses(toStringArray(v))
	// 		}
	// 	}
	// 	if d.HasChange("dsp_keystore_path") {
	// 		if v, has_v := d.GetOk("dsp_keystore_path"); has_v {
	// 			hostUpdateParam.SetDspKeystorePath(v.(string))
	// 		}
	// 	}
	// 	if d.HasChange("dsp_keystore_password") {
	// 		if v, has_v := d.GetOk("dsp_keystore_password"); has_v {
	// 			hostUpdateParam.SetDspKeystorePassword(v.(string))
	// 		}
	// 	}
	// 	if d.HasChange("dsp_keystore_alias") {
	// 		if v, has_v := d.GetOk("dsp_keystore_alias"); has_v {
	// 			hostUpdateParam.SetDspKeystoreAlias(v.(string))
	// 		}
	// 	}
	// 	if d.HasChange("dsp_truststore_path") {
	// 		if v, has_v := d.GetOk("dsp_truststore_path"); has_v {
	// 			hostUpdateParam.SetDspTruststorePath(v.(string))
	// 		}
	// 	}
	// 	if d.HasChange("dsp_truststore_password") {
	// 		if v, has_v := d.GetOk("dsp_truststore_password"); has_v {
	// 			hostUpdateParam.SetDspTruststorePassword(v.(string))
	// 		}
	// 	}
	// 	if d.HasChange("oracle_jdbc_keystore_password") {
	// 		if v, has_v := d.GetOk("oracle_jdbc_keystore_password"); has_v {
	// 			hostUpdateParam.SetOracleJdbcKeystorePassword(v.(string))
	// 		}
	// 	}
	// 	if d.HasChange("oracle_tde_keystores_root_path") {
	// 		if v, has_v := d.GetOk("oracle_tde_keystores_root_path"); has_v {
	// 			hostUpdateParam.SetOracleTdeKeystoresRootPath(v.(string))
	// 		}
	// 	}
	// 	if d.HasChange("connector_port") {
	// 		if v, has_v := d.GetOk("connector_port"); has_v {
	// 			hostUpdateParam.SetConnectorPort(v.(int32))
	// 		}
	// 	}
	// 	if d.HasChange("connector_authentication_key") {
	// 		if v, has_v := d.GetOk("connector_authentication_key"); has_v {
	// 			hostUpdateParam.SetConnectorAuthenticationKey(v.(string))
	// 		}
	// 	}

	// 	hostUpdateRes, hostHttpRes, hostUpdateErr := client.EnvironmentsAPI.UpdateHost(ctx, environmentId, hostID).HostUpdateParameters(*hostUpdateParam).Execute()
	// 	if diags := apiErrorResponseHelper(ctx, hostUpdateRes, hostHttpRes, hostUpdateErr); diags != nil {
	// 		revertChanges(d, changedKeys)
	// 		return diags
	// 	}

	// 	job_res, job_err := PollJobStatus(*hostUpdateRes.Job.Id, ctx, client)
	// 	if job_err != "" {
	// 		tflog.Warn(ctx, DLPX+WARN+"Env Host Update Job Polling failed but continuing with update. Error: "+job_err)
	// 	}
	// 	tflog.Info(ctx, DLPX+INFO+"Job result is "+job_res)
	// 	if job_res == Failed || job_res == Canceled || job_res == Abandoned {
	// 		tflog.Error(ctx, DLPX+ERROR+"Job "+job_res+" "+*hostUpdateRes.Job.Id+"!")
	// 		revertChanges(d, changedKeys)
	// 		return diag.Errorf("[NOT OK] Job %s %s with error %s", *hostUpdateRes.Job.Id, job_res, job_err)
	// 	}

	// 	for _, item := range dsouceItems {
	// 		if diags := enabledSource(ctx, client, *item.Id); diags != nil {
	// 			return diags //if failure should we enable
	// 		}
	// 	}

	// 	for _, item := range items {
	// 		if diags := enableVDB(ctx, client, *item.Id); diags != nil {
	// 			return diags //if failure should we enable
	// 		}
	// 	}
	// }
	// ========================Old=============================
	readDiags := resourceEnvironmentRead(ctx, d, meta)
	if readDiags.HasError() {
		return readDiags
	}
	return diags
}

// func attemptEnableVDBs(ctx context.Context, client *dctapi.APIClient, items []dctapi.VDB) diag.Diagnostics {
// 	var disableDiags diag.Diagnostics
// 	tflog.Info(ctx, DLPX+INFO+" vdb disable failed ...attempting enable for vdbs")
// 	for _, item := range items {
// 		if diags := enableVDB(ctx, client, *item.Id); diags != nil {
// 			disableDiags = append(disableDiags, diags...)
// 			return disableDiags
// 		}
// 	}
// 	return disableDiags
// }

func revertChanges(d *schema.ResourceData, changedKeys []string) {
	for _, key := range changedKeys {
		old, _ := d.GetChange(key)
		d.Set(key, old)
	}
}

func resourceEnvironmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*apiClient).client
	envId := d.Id()

	apiRes, httpRes, err := client.EnvironmentsAPI.DeleteEnvironment(ctx, envId).Execute()

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
		return client.EnvironmentsAPI.GetEnvironmentById(ctx, envId).Execute()
	})

	return diags
}
