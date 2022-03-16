package provider

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	openapi "github.com/Uddipaan-Hazarika/demo-go-sdk"
)

var SLEEP_TIME = 10

// Job Polling function that makes call to the job status API and checks for status of the JOB
// Input is job status, context and the client
// Returns the status of the given JOB-ID and Error body as a string
func PollJobStatus(job_id string, ctx context.Context, client *openapi.APIClient) (string, string) {

	res, httpRes, err := client.JobsApi.GetJobById(ctx, job_id).Execute()
	if err != nil {
		resBody, err := ResponseBodyToString(httpRes.Body)
		if err != nil {
			return "", err.Error()
		}
		log.Print(err.Error())
		return "", resBody
	}

	var i = 0
	for res.GetStatus() == Running {
		time.Sleep(time.Duration(JOB_STATUS_SLEEP_TIME) * time.Second)
		res, httpRes, err = client.JobsApi.GetJobById(ctx, job_id).Execute()
		if err != nil {
			resBody, err := ResponseBodyToString(httpRes.Body)
			if err != nil {
				return "", err.Error()
			}
			log.Print(err.Error())
			return "", resBody
		}
		i++
		log.Printf("DCT JobId:%s / Status:%s / Error:%s / PollIteration:%d", job_id, res.GetStatus(), res.GetErrorDetails(), i)
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
		log.Print("Error occured in reading body of the response.")
		return "", err
	}
	return string(bytes), nil
}

func PollForObjectCreation(apiCall func() (interface{}, *http.Response, error)) bool {
	return PollForStatusCode(apiCall, http.StatusOK, 20)
}

func PollForObjectDeletion(apiCall func() (interface{}, *http.Response, error)) bool {
	return PollForStatusCode(apiCall, http.StatusNotFound, 0)
}

// poll counter is the retry counter for which an api call should be retried.
func PollForStatusCode(apiCall func() (interface{}, *http.Response, error), statusCode int, pollCounter int) bool {
	counter := 0
	for {
		if pollCounter != 0 && counter > pollCounter { //retry is exhausted then break
			log.Print("Breaking poll for status as retry exhausted")
			return false
		} else {
			time.Sleep(time.Duration(STATUS_POLL_SLEEP_TIME) * time.Second) //otherwise sleep and make the call
		}
		_, httpRes, _ := apiCall()
		if httpRes.StatusCode == statusCode {
			log.Print("Breaking poll for status as status reached")
			break
		}
		counter = counter + 1
	}
	return true
}

func toStringArray(array interface{}) []string {
	items := []string{}
	for _, item := range array.([]interface{}) {
		items = append(items, item.(string))
	}
	return items
}

func flattenHosts(hosts []openapi.Host) []interface{} {
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
