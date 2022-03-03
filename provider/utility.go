package provider

import (
	"context"
	"log"
	"time"

	openapi "github.com/Uddipaan-Hazarika/demo-go-sdk"
)

var SLEEP_TIME = 5

// Job Polling function that makes call to the job status API and checks for status of the JOB
// Input is job status, context with the APIKEY and the client
// Returns the status of the given JOB-ID
func PollJobStatus(job_id string, ctx context.Context, client *openapi.APIClient) (string, error) {

	res, _, err := client.JobsApi.GetJobById(ctx, job_id).Execute()

	if err != nil {
		log.Print(err.Error())
		return "", err
	}

	var i = 0
	for res.GetStatus() == "RUNNING" {
		time.Sleep(time.Duration(SLEEP_TIME) * time.Second)
		res, _, err = client.JobsApi.GetJobById(ctx, job_id).Execute()
		if err != nil {
			log.Print(err.Error())
			return "", err
		}
		i++
		log.Printf("__________JOB-STATUS_________Iteration %d", i)
		log.Print(res.GetStatus())
	}

	return *res.Status, nil
}
