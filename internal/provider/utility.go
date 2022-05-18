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
			if httpRes != nil && httpRes.Body != nil {
				resBody, err := ResponseBodyToString(httpRes.Body)
				if err != nil {
					return "", err.Error()
				}
				ErrorLog.Print(err.Error())
				return "", resBody
			}
			return "", "No body in error response"
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

func apiErrorResponseHelper(httpRes *http.Response, err error) diag.Diagnostics {
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
