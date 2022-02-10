# \SnapshotsApi

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSnapshotById**](SnapshotsApi.md#GetSnapshotById) | **Get** /snapshots/{snapshotId} | Get a Snapshot by ID.
[**GetSnapshots**](SnapshotsApi.md#GetSnapshots) | **Get** /snapshots | List Snapshots for a dSource or VDB.



## GetSnapshotById

> Snapshot GetSnapshotById(ctx, snapshotId).Execute()

Get a Snapshot by ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    snapshotId := "snapshotId_example" // string | The ID of the snaphost.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsApi.GetSnapshotById(context.Background(), snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.GetSnapshotById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSnapshotById`: Snapshot
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.GetSnapshotById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | The ID of the snaphost. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Snapshot**](Snapshot.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnapshots

> ListSnaphotsResponse GetSnapshots(ctx).DatasetId(datasetId).Limit(limit).Cursor(cursor).Execute()

List Snapshots for a dSource or VDB.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    datasetId := "datasetId_example" // string | The ID of the dSource or VDB for which to fetch Snapshots. (optional)
    limit := int32(50) // int32 | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. (optional) (default to 100)
    cursor := "RXlhbCBpcyBncmVhdAo=" // string | Cursor to fetch the next or previous page of results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsApi.GetSnapshots(context.Background()).DatasetId(datasetId).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.GetSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSnapshots`: ListSnaphotsResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.GetSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasetId** | **string** | The ID of the dSource or VDB for which to fetch Snapshots. | 
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. | [default to 100]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. | 

### Return type

[**ListSnaphotsResponse**](ListSnaphotsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

