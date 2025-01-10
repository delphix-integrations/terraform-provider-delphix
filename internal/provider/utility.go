package provider

import (
	"context"
	"io"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"time"

	dctapi "github.com/delphix/dct-sdk-go/v23"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var SLEEP_TIME = 10

// Job Polling function that makes call to the job status API and checks for status of the JOB
// Input is job status, context and the client
// Returns the status of the given JOB-ID and Error body as a string
func PollJobStatus(job_id string, ctx context.Context, client *dctapi.APIClient) (string, string) {

	res, httpRes, err := client.JobsAPI.GetJobById(ctx, job_id).Execute()
	if err != nil {
		resBody, resBodyErr := ResponseBodyToString(ctx, httpRes.Body)
		if resBodyErr != nil {
			tflog.Error(ctx, DLPX+ERROR+resBodyErr.Error())
			return "", resBodyErr.Error()
		}
		tflog.Error(ctx, DLPX+ERROR+err.Error())
		return "", resBody
	}

	var i = 0
	for res.GetStatus() == Pending || res.GetStatus() == Started {
		time.Sleep(time.Duration(JOB_STATUS_SLEEP_TIME) * time.Second)
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
		i++
		tflog.Info(ctx, DLPX+INFO+"DCT-JobId:"+job_id+" has Status:"+res.GetStatus())
	}

	return res.GetStatus(), res.GetErrorDetails()
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
	for i := 0; maxRetry == 0 || i < maxRetry; i++ {
		res, httpRes, err = apiCall()
		if httpRes.StatusCode == statusCode {
			tflog.Info(ctx, DLPX+INFO+"[OK] Breaking poll - Status "+strconv.Itoa(statusCode)+" reached.")
			return res, nil
		} else if httpRes.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, DLPX+INFO+"[404 Not found] Breaking poll - Status "+strconv.Itoa(statusCode)+" reached.")
			break
		}
		time.Sleep(time.Duration(STATUS_POLL_SLEEP_TIME) * time.Second)
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
			returnedHost["hostname"] = host.GetHostname()
			returnedHost["os_name"] = host.GetOsName()
			returnedHost["os_version"] = host.GetOsVersion()
			returnedHost["memory_size"] = host.GetMemorySize()
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

func flattenDSourceHooks(hooks []dctapi.Hook) []interface{} {
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
		resBody, nerr := ResponseBodyToString(ctx, httpRes.Body)
		if nerr != nil {
			tflog.Error(ctx, DLPX+ERROR+"An error occurred: "+nerr.Error())
			diags = diag.FromErr(nerr)
		} else {
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
		for attempt := 1; attempt <= maxAttempts; attempt++ {
			snapshotRes, _, api_err = client.DSourcesAPI.GetDsourceSnapshots(ctx, d.Id()).Execute()
			if api_err != nil {
				tflog.Error(ctx, DLPX+ERROR+"Error fetching dSource snapshots: "+api_err.Error())
				break // Exit the loop on error to avoid unnecessary retries
			}
			if len(snapshotRes.GetItems()) > 0 {
				tflog.Info(ctx, DLPX+INFO+"Snapshots are now available.")
				break // Snapshots found, exit the loop
			}
			tflog.Info(ctx, DLPX+INFO+"Attempt "+strconv.Itoa(attempt)+": Waiting for snapshots to become available...")

			if attempt < maxAttempts {
				time.Sleep(time.Duration(STATUS_POLL_SLEEP_TIME) * time.Second) // Wait before retrying
			}
		}

		// After the loop, check for errors or absence of snapshots
		if api_err != nil {
			tflog.Error(ctx, DLPX+ERROR+"Failed to fetch dSource snapshots due to an error.")
		} else if len(snapshotRes.GetItems()) == 0 {
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
	job_res, job_err := PollJobStatus(*apiRes.Job.Id, ctx, client)
	if job_err != "" {
		tflog.Warn(ctx, DLPX+WARN+"VDB disable Job Polling failed. Error: "+job_err)
		//return here
	}
	tflog.Info(ctx, DLPX+INFO+"Job result is "+job_res)
	if job_res == Failed || job_res == Canceled || job_res == Abandoned {
		tflog.Error(ctx, DLPX+ERROR+"Job "+job_res+" "+*apiRes.Job.Id+"!")
		return diag.Errorf("[NOT OK] Job %s %s with error %s", *apiRes.Job.Id, job_res, job_err)
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
	job_res, job_err := PollJobStatus(*apiRes.Job.Id, ctx, client)
	if job_err != "" {
		tflog.Warn(ctx, DLPX+WARN+"VDB enable Job Polling failed. Error: "+job_err)
	}
	tflog.Info(ctx, DLPX+INFO+"Job result is "+job_res)
	if job_res == Failed || job_res == Canceled || job_res == Abandoned {
		tflog.Error(ctx, DLPX+ERROR+"Job "+job_res+" "+*apiRes.Job.Id+"!")
		return diag.Errorf("[NOT OK] Job %s %s with error %s", *apiRes.Job.Id, job_res, job_err)
	}
	return nil
}

func revertChanges(d *schema.ResourceData, changedKeys []string) {
	for _, key := range changedKeys {
		old, _ := d.GetChange(key)
		if !reflect.ValueOf(old).IsZero() { // so that a previously optional param is not set to blank erroraneously
			d.Set(key, old)
		}
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
