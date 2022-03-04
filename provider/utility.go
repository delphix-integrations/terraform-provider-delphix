package provider

import (
	"context"
	"io"
	"log"
	"time"

	openapi "github.com/Uddipaan-Hazarika/demo-go-sdk"
)

var SLEEP_TIME = 5

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
	for res.GetStatus() == "RUNNING" {
		time.Sleep(time.Duration(SLEEP_TIME) * time.Second)
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
		log.Printf("__________JOB-STATUS_________Iteration %d", i)
		log.Print(res.GetStatus())
	}

	return *res.Status, ""
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
