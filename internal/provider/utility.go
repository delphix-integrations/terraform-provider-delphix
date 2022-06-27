package provider

import (
	"context"
	"io"
	"net/http"
	"time"

	dctapi "github.com/delphix/dct-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

var SLEEP_TIME = 10

// Job Polling function that makes call to the job status API and checks for status of the JOB
// Input is job status, context and the client
// Returns the status of the given JOB-ID and Error body as a string
func PollJobStatus(job_id string, ctx context.Context, client *dctapi.APIClient) (string, string) {

	res, httpRes, err := client.JobsApi.GetJobById(ctx, job_id).Execute()
	if err != nil {
		resBody, err := ResponseBodyToString(httpRes.Body)
		if err != nil {
			return "", err.Error()
		}
		ErrorLog.Print(err.Error())
		return "", resBody
	}

	var i = 0
	for res.GetStatus() == Pending || res.GetStatus() == Started {
		time.Sleep(time.Duration(JOB_STATUS_SLEEP_TIME) * time.Second)
		res, httpRes, err = client.JobsApi.GetJobById(ctx, job_id).Execute()
		if err != nil {
			if httpRes == nil {
				return "", "Received nil response for Job ID " + job_id
			}
			resBody, err := ResponseBodyToString(httpRes.Body)
			if err != nil {
				return "", err.Error()
			}
			ErrorLog.Print(err.Error())
			return "", resBody
		}
		i++
		InfoLog.Printf("DCT-JobId:%s has Status:%s", job_id, res.GetStatus())
	}

	return res.GetStatus(), res.GetErrorDetails()
}

// ResponseBodyToString parses the response body from io.readCloser() to string for
// displaying to user in case of any error.
// INPUT: body of any http response.
// OUTPUT: Body of the response in string format and Error object that may occur during the conversion.
func ResponseBodyToString(body io.ReadCloser) (string, error) {
	bytes, err := io.ReadAll(body)
	if err != nil {
		ErrorLog.Print("Error occured in reading body of the response.")
		return "", err
	}
	return string(bytes), nil
}

func PollForObjectExistence(apiCall func() (interface{}, *http.Response, error)) (interface{}, diag.Diagnostics) {
	// Function to check if an object exists in the Delphix estate.
	return PollForStatusCode(apiCall, http.StatusOK, 10)
}

func PollForObjectDeletion(apiCall func() (interface{}, *http.Response, error)) (interface{}, diag.Diagnostics) {
	// Function to check if an object does not exist in the Delphix estate.
	return PollForStatusCode(apiCall, http.StatusNotFound, 10)
}

// poll counter is the retry counter for which an api call should be retried.
func PollForStatusCode(apiCall func() (interface{}, *http.Response, error), statusCode int, maxRetry int) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	var res interface{}
	var httpRes *http.Response
	var err error
	for i := 0; maxRetry == 0 || i < maxRetry; i++ {
		if res, httpRes, err = apiCall(); httpRes.StatusCode == statusCode {
			InfoLog.Printf("[OK] Breaking poll - Status %d reached.", statusCode)
			return res, nil
		}
		time.Sleep(time.Duration(STATUS_POLL_SLEEP_TIME) * time.Second)
	}
	diags = apiErrorResponseHelper(res, httpRes, err)
	InfoLog.Printf("[NOT OK] Breaking poll - Retry exhausted for status %d", statusCode)
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

func apiErrorResponseHelper(res interface{}, httpRes *http.Response, err error) diag.Diagnostics {
	// Helper function to return Diagnostics object if there is
	// a failure during API call.
	var diags diag.Diagnostics
	if err != nil {
		resBody, nerr := ResponseBodyToString(httpRes.Body)
		if nerr != nil {
			ErrorLog.Printf("An error occured: %v", err)
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
