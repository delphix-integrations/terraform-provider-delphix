package provider

import (
	"context"
	"log"
	"os"
	"time"

	openapi "github.com/Uddipaan-Hazarika/demo-go-sdk"
)

var SLEEP_TIME = 5

// Job Polling function that makes call to the job status API and checks for status of the JOB
// Input is job status, context with the APIKEY and the client
// Returns the status of the given JOB-ID
func PollJobStatus(job_id string, ctx context.Context, client *openapi.APIClient) string {

	res, httpRes, err := client.JobsApi.GetJobById(ctx, job_id).Execute()

	if err != nil {
		log.Print("\n______JOB-STATUS____ERR__________\n")
		log.Print(err)
		os.Exit(1)
	}

	log.Print("\n_______JOB-STATUS___RES__________\n")
	log.Print(&res)
	log.Print("\n_______JOB-STATUS___HTTP-RESP__________\n")
	log.Print(httpRes)
	var i = 0
	for res.GetStatus() == "RUNNING" {
		time.Sleep(time.Duration(SLEEP_TIME) * time.Second)
		res, _, _ = client.JobsApi.GetJobById(ctx, job_id).Execute()
		i++
		log.Printf("__________JOB-STATUS_________Iteration %d", i)
		log.Print(res)
	}

	return *res.Status
}
