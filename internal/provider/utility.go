package provider

import (
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	dctapi "github.com/delphix/dct-sdk-go/v25"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var SLEEP_TIME = 10

// Job Polling function that makes call to the job status API and checks for status of the JOB
// Input is job status, context and the client
// Returns the status of the given JOB-ID and Error body as a string
func PollJobStatus(job_id string, ctx context.Context, client *dctapi.APIClient) (string, string) {
	const sleepTime = 10 // seconds

	res, httpRes, err := client.JobsAPI.GetJobById(ctx, job_id).Execute()
	if err != nil {
		// handle possible nil httpRes
		if httpRes == nil {
			tflog.Error(ctx, DLPX+ERROR+err.Error())
			return "", err.Error()
		}
		resBody, resBodyErr := ResponseBodyToString(ctx, httpRes.Body)
		if resBodyErr != nil {
			tflog.Error(ctx, DLPX+ERROR+resBodyErr.Error())
			return "", resBodyErr.Error()
		}
		tflog.Error(ctx, DLPX+ERROR+err.Error())
		return "", resBody
	}

	for {
		// if job reached a terminal state return
		if res.GetStatus() != Pending && res.GetStatus() != Started {
			return res.GetStatus(), res.GetErrorDetails()
		}

		// check if context is done (timeout or cancelled)
		select {
		case <-ctx.Done():
			return "", "Job polling cancelled or timed out"
		default:
			// continue
		}

		time.Sleep(time.Duration(sleepTime) * time.Second)

		res, httpRes, err = client.JobsAPI.GetJobById(ctx, job_id).Execute()
		if err != nil {
			if httpRes == nil {
				return "", "Received nil response for Job ID " + job_id
			}
			resBody, resBodyErr := ResponseBodyToString(ctx, httpRes.Body)
			if resBodyErr != nil {
				tflog.Error(ctx, DLPX+ERROR+resBodyErr.Error())
				return "", resBodyErr.Error()
			}
			tflog.Error(ctx, DLPX+ERROR+err.Error())
			return "", resBody
		}
		tflog.Info(ctx, DLPX+INFO+"DCT-JobId:"+job_id+" has Status:"+res.GetStatus())
	}
}

// ResponseBodyToString parses the response body from io.readCloser() to string for
// displaying to user in case of any error.
// INPUT: body of any http response.
// OUTPUT: Body of the response in string format and Error object that may occur during the conversion.
func ResponseBodyToString(ctx context.Context, body io.ReadCloser) (string, error) {
	bytes, err := io.ReadAll(body)
	if err != nil {
		tflog.Error(ctx, DLPX+ERROR+"Error occurred in reading body of the response.")
		return "", err
	}
	return string(bytes), nil
}

func PollForObjectExistence(ctx context.Context, apiCall func() (interface{}, *http.Response, error)) (interface{}, diag.Diagnostics) {
	// Function to check if an object exists in the Delphix estate.
	return PollForStatusCode(ctx, apiCall, http.StatusOK, 10)
}

func PollForObjectDeletion(ctx context.Context, apiCall func() (interface{}, *http.Response, error)) (interface{}, diag.Diagnostics) {
	// Function to check if an object does not exist in the Delphix estate.
	return PollForStatusCode(ctx, apiCall, http.StatusNotFound, 10)
}

// poll counter is the retry counter for which an api call should be retried.
func PollForStatusCode(ctx context.Context, apiCall func() (interface{}, *http.Response, error), statusCode int, maxRetry int) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	var res interface{}
	var httpRes *http.Response
	var err error
	attempt := 0
	for maxRetry == 0 || attempt < maxRetry {
		// return early if context cancelled
		select {
		case <-ctx.Done():
			return nil, diag.Errorf("polling cancelled or timed out: %v", ctx.Err())
		default:
		}

		res, httpRes, err = apiCall()
		if httpRes != nil {
			if httpRes.StatusCode == statusCode {
				tflog.Info(ctx, DLPX+INFO+"[OK] Breaking poll - Status "+strconv.Itoa(statusCode)+" reached.")
				return res, nil
			} else if httpRes.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, DLPX+INFO+"[404 Not found] Breaking poll - Status "+strconv.Itoa(statusCode)+" reached.")
				break
			}
		}

		attempt++

		// sleep but wake early on context done
		sleep := time.NewTimer(time.Duration(STATUS_POLL_SLEEP_TIME) * time.Second)
		select {
		case <-ctx.Done():
			sleep.Stop()
			return nil, diag.Errorf("polling cancelled or timed out: %v", ctx.Err())
		case <-sleep.C:
		}
	}
	diags = apiErrorResponseHelper(ctx, res, httpRes, err)
	tflog.Info(ctx, DLPX+INFO+"[OK] Breaking poll - Retry exhausted for status "+strconv.Itoa(statusCode))
	return nil, diags
}

func toStringArray(array interface{}) []string {
	items := []string{}
	for _, item := range array.([]interface{}) {
		items = append(items, item.(string))
	}
	return items
}

func flattenHosts(hosts []dctapi.Host) []interface{} {
	if hosts != nil {
		returnedHosts := make([]interface{}, len(hosts))
		for i, host := range hosts {
			returnedHost := make(map[string]interface{})
			returnedHost["id"] = host.GetId()
			returnedHost["hostname"] = host.GetHostname()
			returnedHost["os_name"] = host.GetOsName()
			returnedHost["os_version"] = host.GetOsVersion()
			returnedHost["memory_size"] = host.GetMemorySize()
			returnedHost["ssh_port"] = host.GetSshPort()
			returnedHost["toolkit_path"] = host.GetToolkitPath()
			returnedHost["processor_type"] = host.GetProcessorType()
			returnedHost["timezone"] = host.GetTimezone()
			returnedHost["available"] = host.GetAvailable()
			returnedHost["nfs_addresses"] = host.GetNfsAddresses()
			returnedHost["java_home"] = host.GetJavaHome()
			returnedHost["oracle_tde_keystores_root_path"] = host.GetOracleTdeKeystoresRootPath()

			returnedHosts[i] = returnedHost
		}
		return returnedHosts
	}
	return make([]interface{}, 0)
}

func flattenHostRepositories(repos []dctapi.Repository) []interface{} {
	if repos != nil {
		returnedRepos := make([]interface{}, len(repos))
		for i, host := range repos {
			returnedRepo := make(map[string]interface{})
			returnedRepo["id"] = host.GetId()
			returnedRepo["name"] = host.GetName()
			returnedRepo["database_type"] = host.GetDatabaseType()
			returnedRepo["allow_provisioning"] = host.GetAllowProvisioning()
			returnedRepo["is_staging"] = host.GetIsStaging()
			returnedRepo["oracle_base"] = host.GetOracleBase()
			returnedRepo["bits"] = host.GetBits()
			returnedRepos[i] = returnedRepo
		}
		return returnedRepos
	}
	return make([]interface{}, 0)
}

func flattenAdditionalMountPoints(additional_mount_points []dctapi.AdditionalMountPoint) []interface{} {
	if additional_mount_points != nil {
		returned_additional_mount_points := make([]interface{}, len(additional_mount_points))
		for i, additional_mount_point := range additional_mount_points {
			returned_additional_mount_point := make(map[string]interface{})
			returned_additional_mount_point["shared_path"] = additional_mount_point.GetSharedPath()
			returned_additional_mount_point["mount_path"] = additional_mount_point.GetMountPath()
			returned_additional_mount_point["environment_id"] = additional_mount_point.GetEnvironmentId()
			returned_additional_mount_points[i] = returned_additional_mount_point
		}
		return returned_additional_mount_points
	}
	return make([]interface{}, 0)
}

func flattenVDbHooks(hooks []dctapi.Hook) []interface{} {
	if hooks != nil {
		returnedHooks := make([]interface{}, len(hooks))
		for i, hook := range hooks {
			returnedHook := make(map[string]interface{})
			returnedHook["name"] = hook.GetName()
			returnedHook["command"] = hook.GetCommand()
			returnedHook["shell"] = hook.GetShell()
			returnedHook["element_id"] = hook.GetElementId()
			returnedHook["has_credentials"] = hook.GetHasCredentials()
			returnedHooks[i] = returnedHook
		}
		return returnedHooks
	}
	return make([]interface{}, 0)
}

func flattenDSourceHooks(hooks []dctapi.Hook, oldList []dctapi.SourceOperation) []interface{} {
	if hooks != nil {
		returnedHooks := make([]interface{}, len(hooks))
		for i, hook := range hooks {
			returnedHook := make(map[string]interface{})
			returnedHook["name"] = hook.GetName()
			returnedHook["command"] = hook.GetCommand()
			returnedHook["shell"] = hook.GetShell()
			returnedHook["element_id"] = hook.GetElementId()
			returnedHook["has_credentials"] = hook.GetHasCredentials()
			credsEnvVars := []map[string]interface{}{}
			if len(oldList) != 0 {
				for _, cred := range oldList[i].GetCredentialsEnvVars() {
					credsEnvVars = append(credsEnvVars, map[string]interface{}{
						"base_var_name":                cred.BaseVarName,
						"password":                     cred.Password,
						"vault":                        cred.Vault,
						"azure_vault_name":             cred.AzureVaultName,
						"azure_vault_secret_key":       cred.AzureVaultSecretKey,
						"azure_vault_username_key":     cred.AzureVaultUsernameKey,
						"cyberark_vault_query_string":  cred.CyberarkVaultQueryString,
						"hashicorp_vault_engine":       cred.HashicorpVaultEngine,
						"hashicorp_vault_secret_key":   cred.HashicorpVaultSecretKey,
						"hashicorp_vault_secret_path":  cred.HashicorpVaultSecretPath,
						"hashicorp_vault_username_key": cred.HashicorpVaultUsernameKey,
					})
				}
			}
			returnedHook["credentials_env_vars"] = credsEnvVars
			returnedHooks[i] = returnedHook
		}
		return returnedHooks
	}
	return make([]interface{}, 0)
}

func flattenTags(tags []dctapi.Tag) []interface{} {
	if tags != nil {
		returnedTags := make([]interface{}, len(tags))
		for i, tag := range tags {
			returnedTag := make(map[string]interface{})
			returnedTag["key"] = tag.GetKey()
			returnedTag["value"] = tag.GetValue()
			returnedTags[i] = returnedTag
		}
		return returnedTags
	}
	return make([]interface{}, 0)
}

func apiErrorResponseHelper(ctx context.Context, res interface{}, httpRes *http.Response, err error) diag.Diagnostics {
	// Helper function to return Diagnostics object if there is
	// a failure during API call.
	var diags diag.Diagnostics
	if err != nil {
		if httpRes == nil || httpRes.Body == nil {
			tflog.Error(ctx, DLPX+ERROR+"An error occurred: "+err.Error())
			return diag.FromErr(err)
		}
		resBody, nerr := ResponseBodyToString(ctx, httpRes.Body)
		if nerr != nil {
			tflog.Error(ctx, DLPX+ERROR+"An error occurred: "+nerr.Error())
			diags = diag.FromErr(nerr)
		} else {
			tflog.Info(ctx, DLPX+INFO+"Error: "+resBody)
			diags = diag.Errorf(resBody)
		}
		return diags
	}
	return nil
}

func isJobTerminalFailure(job_status string) bool {
	return job_status == Failed || job_status == Canceled || job_status == Abandoned
}

// Poll the /dsources/{dsourceId}/snapshots API till atleast one snapshot is created
func PollSnapshotStatus(d *schema.ResourceData, ctx context.Context, client *dctapi.APIClient) {
	skip := d.Get("skip_wait_for_snapshot_creation") // default false
	wait_time := d.Get("wait_time")                  // default 3 mins

	if !skip.(bool) {
		var snapshotRes *dctapi.ListSnapshotsResponse
		var api_err error
		maxAttempts := int(math.Round(float64(wait_time.(int)*60) / float64(STATUS_POLL_SLEEP_TIME)))
		attempt := 0
		for maxAttempts == 0 || attempt < maxAttempts {
			select {
			case <-ctx.Done():
				tflog.Info(ctx, DLPX+INFO+"PollSnapshotStatus cancelled or timed out")
				return
			default:
			}

			snapshotRes, _, api_err = client.DSourcesAPI.GetDsourceSnapshots(ctx, d.Id()).Execute()
			if api_err != nil {
				tflog.Error(ctx, DLPX+ERROR+"Error fetching dSource snapshots: "+api_err.Error())
				break // Exit the loop on error to avoid unnecessary retries
			}
			if snapshotRes != nil && len(snapshotRes.GetItems()) > 0 {
				tflog.Info(ctx, DLPX+INFO+"Snapshots are now available.")
				break // Snapshots found, exit the loop
			}
			tflog.Info(ctx, DLPX+INFO+"Attempt "+strconv.Itoa(attempt+1)+": Waiting for snapshots to become available...")

			attempt++

			sleep := time.NewTimer(time.Duration(STATUS_POLL_SLEEP_TIME) * time.Second)
			select {
			case <-ctx.Done():
				sleep.Stop()
				tflog.Info(ctx, DLPX+INFO+"PollSnapshotStatus cancelled or timed out during sleep")
				return
			case <-sleep.C:
			}
		}

		// After the loop, check for errors or absence of snapshots
		if api_err != nil {
			tflog.Error(ctx, DLPX+ERROR+"Failed to fetch dSource snapshots due to an error.")
		} else if snapshotRes == nil || len(snapshotRes.GetItems()) == 0 {
			tflog.Info(ctx, DLPX+INFO+"Maximum attempts reached. Snapshots are not available.")
		}
	}
}

func disableVDB(ctx context.Context, client *dctapi.APIClient, vdbId string) diag.Diagnostics {
	tflog.Info(ctx, DLPX+INFO+"Disable VDB "+vdbId)
	disableVDBParam := dctapi.NewDisableVDBParameters()
	apiRes, httpRes, err := client.VDBsAPI.DisableVdb(ctx, vdbId).DisableVDBParameters(*disableVDBParam).Execute()
	if diags := apiErrorResponseHelper(ctx, apiRes, httpRes, err); diags != nil {
		return diags
	}
	job_res, job_err := PollJobStatus(apiRes.Job.GetId(), ctx, client)
	if job_err != "" {
		tflog.Warn(ctx, DLPX+WARN+"VDB disable Job Polling failed. Error: "+job_err)
		//return here
	}
	tflog.Info(ctx, DLPX+INFO+"Job result is "+job_res)
	if job_res == Failed || job_res == Canceled || job_res == Abandoned {
		tflog.Error(ctx, DLPX+ERROR+"Job "+job_res+" "+apiRes.Job.GetId()+"!")
		return diag.Errorf("[NOT OK] Job %s %s with error %s", apiRes.Job.GetId(), job_res, job_err)
	}
	return nil
}

func enableVDB(ctx context.Context, client *dctapi.APIClient, vdbId string) diag.Diagnostics {
	tflog.Info(ctx, DLPX+INFO+"Enable VDB "+vdbId)
	enableVDBParam := dctapi.NewEnableVDBParameters()
	apiRes, httpRes, err := client.VDBsAPI.EnableVdb(ctx, vdbId).EnableVDBParameters(*enableVDBParam).Execute()
	if diags := apiErrorResponseHelper(ctx, apiRes, httpRes, err); diags != nil {
		return diags
	}
	job_res, job_err := PollJobStatus(apiRes.Job.GetId(), ctx, client)
	if job_err != "" {
		tflog.Warn(ctx, DLPX+WARN+"VDB enable Job Polling failed. Error: "+job_err)
	}
	tflog.Info(ctx, DLPX+INFO+"Job result is "+job_res)
	if job_res == Failed || job_res == Canceled || job_res == Abandoned {
		tflog.Error(ctx, DLPX+ERROR+"Job "+job_res+" "+apiRes.Job.GetId()+"!")
		return diag.Errorf("[NOT OK] Job %s %s with error %s", apiRes.Job.GetId(), job_res, job_err)
	}
	return nil
}

func revertChanges(d *schema.ResourceData, changedKeys []string) {
	for _, key := range changedKeys {
		old, _ := d.GetChange(key)
		if !isEmpty(old) { // so that a previously optional param is not set to blank erroraneously
			d.Set(key, old)
		}
	}
}

func isEmpty(value interface{}) bool {
	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Bool:
		return false
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		return v.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	default:
		return v.IsZero()
	}
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

func toIntArray(array interface{}) []int32 {
	items := []int32{}
	for _, item := range array.([]interface{}) {
		items = append(items, int32(item.(int)))
	}
	return items
}

func isSnapSyncFailure(job_id string, ctx context.Context, client *dctapi.APIClient) bool {
	res, httpRes, _ := client.JobsAPI.GetJobById(ctx, job_id).Execute()
	if httpRes.StatusCode == 200 && len(res.GetTasks()) != 0 {
		tflog.Info(ctx, "Status of task 1 is "+res.GetTasks()[0].GetStatus())
		if res.GetTasks()[0].GetStatus() == "COMPLETED" {
			tflog.Info(ctx, "rolling back Dsource")
			return true
		}
	}
	return false
}

func filterVDBs(ctx context.Context, client *dctapi.APIClient, envId string) ([]dctapi.VDB, diag.Diagnostics) {
	tflog.Info(ctx, DLPX+INFO+"Filter VBDs by envId "+envId)
	vdbSearchExpr := dctapi.NewSearchBody()
	vdbSearchExpr.SetFilterExpression(fmt.Sprintf("environment_id eq '%s'", envId))

	apiReq := client.VDBsAPI.SearchVdbs(ctx)
	apiRes, httpRes, err := apiReq.SearchBody(*vdbSearchExpr).Execute()
	if diags := apiErrorResponseHelper(ctx, apiRes, httpRes, err); diags != nil {
		return nil, diags
	}
	return apiRes.Items, nil
}

func filterSources(ctx context.Context, client *dctapi.APIClient, envId string) ([]dctapi.Source, diag.Diagnostics) {
	tflog.Info(ctx, DLPX+INFO+"Filter Sources by envId "+envId)
	sourceSearchExpr := dctapi.NewSearchBody()
	sourceSearchExpr.SetFilterExpression(fmt.Sprintf("environment_id eq '%s'", envId))
	apiReq := client.SourcesAPI.SearchSources(ctx)
	apiRes, httpRes, err := apiReq.SearchBody(*sourceSearchExpr).Execute()
	if diags := apiErrorResponseHelper(ctx, apiRes, httpRes, err); diags != nil {
		return nil, diags
	}
	return apiRes.Items, nil
}

func filterdSources(ctx context.Context, client *dctapi.APIClient, sourceIds []string) ([]dctapi.DSource, diag.Diagnostics) {
	tflog.Info(ctx, DLPX+INFO+"Filter dSources by SourceIds "+strings.Join(sourceIds, ", "))
	dsourceSearchExpr := dctapi.NewSearchBody()
	dsourceSearchExpr.SetFilterExpression(fmt.Sprintf("source_id in ['%s']", strings.Join(sourceIds, "', '")))
	tflog.Info(ctx, DLPX+INFO+"Filter dSources by SourceIds "+dsourceSearchExpr.GetFilterExpression())
	apiReq := client.DSourcesAPI.SearchDsources(ctx)
	apiRes, httpRes, err := apiReq.SearchBody(*dsourceSearchExpr).Execute()
	if diags := apiErrorResponseHelper(ctx, apiRes, httpRes, err); diags != nil {
		return nil, diags
	}
	return apiRes.Items, nil
}

func disabledSource(ctx context.Context, client *dctapi.APIClient, dsourceId string) diag.Diagnostics {
	tflog.Info(ctx, DLPX+INFO+"Disable dSource "+dsourceId)
	disableDsourceParam := dctapi.NewDisableDsourceParameters()
	apiRes, httpRes, err := client.DSourcesAPI.DisableDsource(ctx, dsourceId).DisableDsourceParameters(*disableDsourceParam).Execute()
	if diags := apiErrorResponseHelper(ctx, apiRes, httpRes, err); diags != nil {
		return diags
	}
	job_res, job_err := PollJobStatus(*apiRes.Job.Id, ctx, client)
	if job_err != "" {
		tflog.Warn(ctx, DLPX+WARN+"dSource disable Job Polling failed. Error: "+job_err)
	}
	tflog.Info(ctx, DLPX+INFO+"Job result is "+job_res)
	if job_res == Failed || job_res == Canceled || job_res == Abandoned {
		tflog.Error(ctx, DLPX+ERROR+"Job "+job_res+" "+*apiRes.Job.Id+"!")
		return diag.Errorf("[NOT OK] Job %s %s with error %s", *apiRes.Job.Id, job_res, job_err)
	}
	return nil
} //decide if continue or exit

func enableDsource(ctx context.Context, client *dctapi.APIClient, dsourceId string) diag.Diagnostics {
	tflog.Info(ctx, DLPX+INFO+"Enable dSource "+dsourceId)
	enableDsourceParam := dctapi.NewEnableDsourceParameters()
	apiRes, httpRes, err := client.DSourcesAPI.EnableDsource(ctx, dsourceId).EnableDsourceParameters(*enableDsourceParam).Execute()
	if diags := apiErrorResponseHelper(ctx, apiRes, httpRes, err); diags != nil {
		return diags
	}
	job_res, job_err := PollJobStatus(*apiRes.Job.Id, ctx, client)
	if job_err != "" {
		tflog.Warn(ctx, DLPX+WARN+"dSource enable Job Polling failed. Error: "+job_res)
	}
	tflog.Info(ctx, DLPX+INFO+"Job result is "+job_res)
	if job_res == Failed || job_res == Canceled || job_res == Abandoned {
		tflog.Error(ctx, DLPX+ERROR+"Job "+job_res+" "+*apiRes.Job.Id+"!")
		return diag.Errorf("[NOT OK] Job %s %s with error %s", *apiRes.Job.Id, job_res, job_err)
	}
	return nil
} //decide if continue or exit

func toSourceOperationArray(array interface{}) []dctapi.SourceOperation {
	items := []dctapi.SourceOperation{}
	for _, item := range array.([]interface{}) {
		item_map := item.(map[string]interface{})
		sourceOperation := dctapi.NewSourceOperation(item_map["name"].(string), item_map["command"].(string))
		if item_map["shell"].(string) != "" {
			sourceOperation.SetShell(item_map["shell"].(string))
		}
		sourceOperation.SetCredentialsEnvVars(toCredentialsEnvVariableArray(item_map["credentials_env_vars"]))
		items = append(items, *sourceOperation)
	}
	return items
}

func toCredentialsEnvVariableArray(array interface{}) []dctapi.CredentialsEnvVariable {
	items := []dctapi.CredentialsEnvVariable{}
	for _, item := range array.([]interface{}) {
		item_map := item.(map[string]interface{})

		credentialsEnvVariable_item := dctapi.NewCredentialsEnvVariable(item_map["base_var_name"].(string))
		if item_map["password"].(string) != "" {
			credentialsEnvVariable_item.SetPassword(item_map["password"].(string))
		}
		if item_map["vault"].(string) != "" {
			credentialsEnvVariable_item.SetVault(item_map["vault"].(string))
		}
		if item_map["hashicorp_vault_engine"].(string) != "" {
			credentialsEnvVariable_item.SetHashicorpVaultEngine(item_map["hashicorp_vault_engine"].(string))
		}
		if item_map["hashicorp_vault_secret_path"].(string) != "" {
			credentialsEnvVariable_item.SetHashicorpVaultSecretPath(item_map["hashicorp_vault_secret_path"].(string))
		}
		if item_map["hashicorp_vault_username_key"].(string) != "" {
			credentialsEnvVariable_item.SetHashicorpVaultUsernameKey(item_map["hashicorp_vault_username_key"].(string))
		}
		if item_map["hashicorp_vault_secret_key"].(string) != "" {
			credentialsEnvVariable_item.SetHashicorpVaultSecretKey(item_map["hashicorp_vault_secret_key"].(string))
		}
		if item_map["azure_vault_name"].(string) != "" {
			credentialsEnvVariable_item.SetAzureVaultName(item_map["azure_vault_name"].(string))
		}
		if item_map["azure_vault_username_key"].(string) != "" {
			credentialsEnvVariable_item.SetAzureVaultUsernameKey(item_map["azure_vault_username_key"].(string))
		}
		if item_map["azure_vault_secret_key"].(string) != "" {
			credentialsEnvVariable_item.SetAzureVaultSecretKey(item_map["azure_vault_secret_key"].(string))
		}
		if item_map["cyberark_vault_query_string"].(string) != "" {
			credentialsEnvVariable_item.SetCyberarkVaultQueryString(item_map["cyberark_vault_query_string"].(string))
		}
		items = append(items, *credentialsEnvVariable_item)
	}
	return items
}

// isStructEmpty checks if all fields in a struct are at their zero values
func isStructEmpty(v interface{}) bool {
	val := reflect.ValueOf(v).Elem() // Get the underlying value of the pointer
	for i := 0; i < val.NumField(); i++ {
		if !val.Field(i).IsZero() {
			return false // If any field is not zero, the struct is not empty
		}
	}
	return true
}

// This is CustomizeDiff function that is used to specifically for tags where we want to delete all the tags and due to
// Computed true, the value is computed on its own and all the tags cannot be deleted. To handle that we need to get user config
// ie. Raw Config and override the computed value.
func CustomizeDiffTags(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	tflog.Debug(ctx, "CustomizeDiffTags: Start")
	config := d.GetRawConfig()

	if !config.IsKnown() || config.IsNull() {
		return nil // Config is missing or unknown
	}
	ignore_tag_changes := config.GetAttr("ignore_tag_changes")
	if (!ignore_tag_changes.IsKnown() && ignore_tag_changes.IsNull()) || ignore_tag_changes.True() {
		return nil
	}
	attr := config.GetAttr("tags")
	tflog.Debug(ctx, "CustomizeDiffTags: tags raw config value", map[string]interface{}{
		"tags": attr.GoString(),
	})
	if !attr.IsKnown() || attr.IsNull() {
		// Field is omitted in config
		tflog.Info(ctx, "CustomizeDiffTags: CUSTOM Tags field is not set, ignoring from attr block")
		return nil
	}

	if attr.LengthInt() == 0 {
		tflog.Info(ctx, "CustomizeDiffTags: CUSTOM Tags field is empty, ignoring changes from length block")
		// You can now trigger a diff or set a flag
		err := d.SetNew("tags", []interface{}{})
		if err != nil {
			tflog.Info(ctx, "CustomizeDiffTags: CUSTOM Error setting new tags value", map[string]interface{}{
				"error": err,
			})
			return err
		}
	}

	return nil
}

// searchVDBByName attempts to find a VDB by name and return its ID
// Retries for up to 5 minutes if the resource is not found (may still be creating)
// Creates fresh contexts for each attempt to avoid using expired contexts
func searchVDBByName(ctx context.Context, client *dctapi.APIClient, name string) (string, error) {
	const maxWaitTime = 5 * 60 // 5 minutes in seconds
	const retryInterval = 10   // 10 seconds between retries
	const searchTimeout = 60   // 60 seconds timeout per search attempt
	maxAttempts := maxWaitTime / retryInterval
	
	searchBody := dctapi.NewSearchBody()
	searchBody.SetFilterExpression(fmt.Sprintf("name eq '%s'", name))
	
	for attempt := 0; attempt < maxAttempts; attempt++ {
		// Create a fresh context with timeout for each search attempt
		searchCtx, cancel := context.WithTimeout(context.Background(), time.Duration(searchTimeout)*time.Second)
		apiRes, _, err := client.VDBsAPI.SearchVdbs(searchCtx).SearchBody(*searchBody).Execute()
		cancel() // Always cancel to release resources
		
		if err != nil {
			tflog.Warn(ctx, fmt.Sprintf("Error searching for VDB '%s': %v", name, err))
		} else if len(apiRes.Items) > 0 {
			tflog.Info(ctx, fmt.Sprintf("Found VDB '%s' after %d attempts", name, attempt+1))
			return apiRes.Items[0].GetId(), nil
		}
		
		if attempt < maxAttempts-1 {
			tflog.Info(ctx, fmt.Sprintf("VDB '%s' not found, waiting %d seconds before retry (attempt %d/%d)", name, retryInterval, attempt+1, maxAttempts))
			time.Sleep(time.Duration(retryInterval) * time.Second)
		}
	}
	
	return "", fmt.Errorf("VDB with name '%s' not found after %d minutes", name, maxWaitTime/60)
}

// searchDSourceByName attempts to find a dSource by name and return its ID
// Retries for up to 5 minutes if the resource is not found (may still be creating)
// Creates fresh contexts for each attempt to avoid using expired contexts
func searchDSourceByName(ctx context.Context, client *dctapi.APIClient, name string) (string, error) {
	const maxWaitTime = 5 * 60 // 5 minutes in seconds
	const retryInterval = 10   // 10 seconds between retries
	const searchTimeout = 60   // 60 seconds timeout per search attempt
	maxAttempts := maxWaitTime / retryInterval
	
	searchBody := dctapi.NewSearchBody()
	searchBody.SetFilterExpression(fmt.Sprintf("name eq '%s'", name))
	
	for attempt := 0; attempt < maxAttempts; attempt++ {
		// Create a fresh context with timeout for each search attempt
		searchCtx, cancel := context.WithTimeout(context.Background(), time.Duration(searchTimeout)*time.Second)
		apiRes, _, err := client.DSourcesAPI.SearchDsources(searchCtx).SearchBody(*searchBody).Execute()
		cancel() // Always cancel to release resources
		
		if err != nil {
			tflog.Warn(ctx, fmt.Sprintf("Error searching for dSource '%s': %v", name, err))
		} else if len(apiRes.Items) > 0 {
			tflog.Info(ctx, fmt.Sprintf("Found dSource '%s' after %d attempts", name, attempt+1))
			return apiRes.Items[0].GetId(), nil
		}
		
		if attempt < maxAttempts-1 {
			tflog.Info(ctx, fmt.Sprintf("dSource '%s' not found, waiting %d seconds before retry (attempt %d/%d)", name, retryInterval, attempt+1, maxAttempts))
			time.Sleep(time.Duration(retryInterval) * time.Second)
		}
	}
	
	return "", fmt.Errorf("dSource with name '%s' not found after %d minutes", name, maxWaitTime/60)
}

// GenerateImportBlock generates a Terraform import block for a resource
// and appends it to a session-specific file for easy recovery from timeouts.
// All timeouts in the same terraform apply session write to the same file.
// If resourceID is a placeholder, attempts to search for the actual ID by name.
func GenerateImportBlock(ctx context.Context, client *dctapi.APIClient, resourceType string, resourceName string, resourceID string) string {
	finalID := resourceID
	
	// If we have a placeholder ID, try to find the actual ID by searching
	if strings.HasPrefix(resourceID, "<REPLACE_WITH_") {
		tflog.Info(ctx, fmt.Sprintf("Attempting to find actual ID for %s '%s'", resourceType, resourceName))
		
		var foundID string
		var err error
		
		if resourceType == "delphix_vdb" {
			foundID, err = searchVDBByName(ctx, client, resourceName)
		} else if resourceType == "delphix_appdata_dsource" || resourceType == "delphix_oracle_dsource" {
			foundID, err = searchDSourceByName(ctx, client, resourceName)
		}
		
		if err == nil && foundID != "" {
			tflog.Info(ctx, fmt.Sprintf("Found %s with ID: %s", resourceType, foundID))
			finalID = foundID
		} else {
			tflog.Warn(ctx, fmt.Sprintf("Could not find %s '%s': %v - using placeholder", resourceType, resourceName, err))
		}
	}
	
	importBlock := fmt.Sprintf("import {\n  to = %s.%s\n  id = \"%s\"\n}\n", resourceType, resourceName, finalID)
	
	// Use process ID to group all timeouts from the same terraform apply session
	// Format: terraform_import_blocks_pid12345.txt
	pid := os.Getpid()
	filename := fmt.Sprintf("terraform_import_blocks_pid%d.txt", pid)
	
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Failed to write import block to %s: %v", filename, err))
	} else {
		defer file.Close()
		if _, err := file.WriteString(importBlock); err != nil {
			tflog.Warn(ctx, fmt.Sprintf("Failed to write import block to %s: %v", filename, err))
		} else {
			tflog.Info(ctx, fmt.Sprintf("Import block written to %s for %s.%s (ID: %s)", filename, resourceType, resourceName, finalID))
		}
	}
	
	return importBlock
}
func HandleRawConfigReadContext(ctx context.Context, d *schema.ResourceData, apiRes interface{}) error {
	tflog.Debug(ctx, "HandleRawConfigReadContext:Handling Raw Config Read Context")
	config := d.GetRawConfig()
	tflog.Debug(ctx, "HandleRawConfigReadContext:GOT Raw config")
	if !config.IsKnown() || config.IsNull() {
		tflog.Debug(ctx, "HandleRawConfigReadContext:Config is missing or unknown")
		switch v := apiRes.(type) {
		case *dctapi.VDBGroup:
			d.Set("tags", flattenTags(v.GetTags()))
		case *dctapi.VDB:
			d.Set("tags", flattenTags(v.GetTags()))
		case *dctapi.TagsResponse:
			d.Set("tags", flattenTags(v.GetTags()))
		default:
			fmt.Println("Unknown type")
		}

	} else {
		tflog.Debug(ctx, "Config is known and not null")
		attr := config.GetAttr("tags")
		var check bool = false
		if !attr.IsKnown() || attr.IsNull() {
			// Field is omitted in config
			check = true
			tflog.Info(ctx, "READ:Tags field is not set, ignoring from attr block")
			//d.Set("tags", []interface{}{})
			return nil
		}

		if attr.LengthInt() == 0 && !d.Get("ignore_tag_changes").(bool) {
			tflog.Info(ctx, "READ:Tags field is empty, ignoring changes from length block")
			check = true
			// You can now trigger a diff or set a flag
			d.Set("tags", []interface{}{})
		}

		if !check {
			switch v := apiRes.(type) {
			case *dctapi.VDBGroup:
				d.Set("tags", flattenTags(v.GetTags()))
			case *dctapi.VDB:
				d.Set("tags", flattenTags(v.GetTags()))
			case *dctapi.TagsResponse:
				d.Set("tags", flattenTags(v.GetTags()))
			default:
				fmt.Println("Unknown type")
			}
		}
	}

	return nil
}
