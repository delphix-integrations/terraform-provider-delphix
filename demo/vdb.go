package main

import (
	"context"
	openapi "dct-goapi/oapi"
	"fmt"
	"os"
)

func GetVdb() {
	apiKeyMap := make(map[string]openapi.APIKey)
	apiKeyMap["ApiKeyAuth"] = openapi.APIKey{
		Key:    "1.GCAE1EOOIa3JUQZcKE2umJa81xOo6MLL210m3NdbMfIEG8VTOz02Xt0wbXlsxS8J",
		Prefix: "apk",
	}

	ctx := context.WithValue(context.Background(), openapi.ContextAPIKeys, apiKeyMap)

	cfg := openapi.NewConfiguration()
	cfg.Host = "dct-apigw.dlpxdc.co"
	cfg.Scheme = "https"

	client := openapi.NewAPIClient(cfg)

	req := client.VDBsApi.GetVdbs(ctx)

	res, _, err := client.VDBsApi.GetVdbsExecute(req)

	if err != nil {
		fmt.Print("\n__________ERR__________\n")
		fmt.Print(err)
		os.Exit(1)
	}

	// fmt.Print("\n__________RES__________\n")
	// fmt.Print(res.GetItems())
	// fmt.Print("\n__________HTTP-RESP__________\n")
	// fmt.Print(httpRes)
	for _, j := range res.Items {
		fmt.Printf("%s \n", *j.Id)
	}

}

func ProvVdbSnap() {

	apiKeyMap := make(map[string]openapi.APIKey)
	apiKeyMap["ApiKeyAuth"] = openapi.APIKey{
		Key:    "1.GCAE1EOOIa3JUQZcKE2umJa81xOo6MLL210m3NdbMfIEG8VTOz02Xt0wbXlsxS8J",
		Prefix: "apk",
	}

	ctx := context.WithValue(context.Background(), openapi.ContextAPIKeys, apiKeyMap)

	cfg := openapi.NewConfiguration()
	cfg.Host = "dct-apigw.dlpxdc.co"
	cfg.Scheme = "https"

	client := openapi.NewAPIClient(cfg)

	provisionVDBBySnapshotParameters := openapi.NewProvisionVDBBySnapshotParameters()
	provisionVDBBySnapshotParameters.SetAutoSelectRepository(true)
	provisionVDBBySnapshotParameters.SetSourceDataId("2-ORACLE_DB_CONTAINER-2")

	req := client.VDBsApi.ProvisionVdbBySnapshot(ctx).ProvisionVDBBySnapshotParameters(*provisionVDBBySnapshotParameters)

	res, httpRes, err := req.ProvisionVDBBySnapshotParameters(*provisionVDBBySnapshotParameters).Execute()

	//res, httpRes, err := client.VDBsApi.ProvisionVdbBySnapshotExecute(req)

	if err != nil {
		fmt.Print("\n__________ERR__________\n")
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Print("\n__________RES__________\n")
	fmt.Print(&res)
	fmt.Print("\n__________HTTP-RESP__________\n")
	fmt.Print(httpRes)
}

func DeleteVdb() {
	apiKeyMap := make(map[string]openapi.APIKey)
	apiKeyMap["ApiKeyAuth"] = openapi.APIKey{
		Key:    "1.GCAE1EOOIa3JUQZcKE2umJa81xOo6MLL210m3NdbMfIEG8VTOz02Xt0wbXlsxS8J",
		Prefix: "apk",
	}

	ctx := context.WithValue(context.Background(), openapi.ContextAPIKeys, apiKeyMap)

	cfg := openapi.NewConfiguration()
	cfg.Host = "dct-apigw.dlpxdc.co"
	cfg.Scheme = "https"

	client := openapi.NewAPIClient(cfg)
	vdbId := "2-ORACLE_DB_CONTAINER-9"

	deleteVdbParams := openapi.NewDeleteVDBParametersWithDefaults()
	deleteVdbParams.SetForce(false)

	res, _, err := client.VDBsApi.DeleteVdb(ctx, vdbId).DeleteVDBParameters(*deleteVdbParams).Execute()

	if err != nil {
		fmt.Print("\n__________ERR__________\n")
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Print("\n__________RES__________\n")
	fmt.Print(res.GetJobId())
}

func GetJobStatus() {
	apiKeyMap := make(map[string]openapi.APIKey)
	apiKeyMap["ApiKeyAuth"] = openapi.APIKey{
		Key:    "1.GCAE1EOOIa3JUQZcKE2umJa81xOo6MLL210m3NdbMfIEG8VTOz02Xt0wbXlsxS8J",
		Prefix: "apk",
	}

	ctx := context.WithValue(context.Background(), openapi.ContextAPIKeys, apiKeyMap)

	cfg := openapi.NewConfiguration()
	cfg.Host = "dct-apigw.dlpxdc.co"
	cfg.Scheme = "https"

	client := openapi.NewAPIClient(cfg)

	res, _, err := client.JobsApi.GetJobById(ctx, "").Execute()

	if err != nil {
		fmt.Print("\n__________ERR__________\n")
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Print("\n__________RES__________\n")
	fmt.Print(res.GetStatus())
}
