# \VDBsApi

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVdb**](VDBsApi.md#DeleteVdb) | **Post** /vdbs/{vdbId}/delete | Delete a VDB.
[**DisableVdb**](VDBsApi.md#DisableVdb) | **Post** /vdbs/{vdbId}/disable | Disable a VDB.
[**EnableVdb**](VDBsApi.md#EnableVdb) | **Post** /vdbs/{vdbId}/enable | Enable a VDB.
[**GetVdbById**](VDBsApi.md#GetVdbById) | **Get** /vdbs/{vdbId} | Get a VDB by ID.
[**GetVdbs**](VDBsApi.md#GetVdbs) | **Get** /vdbs | List all vdbs.
[**ProvisionVdbBySnapshot**](VDBsApi.md#ProvisionVdbBySnapshot) | **Post** /vdbs/provision_by_snapshot | Provision a new VDB by snapshot.
[**ProvisionVdbByTimestamp**](VDBsApi.md#ProvisionVdbByTimestamp) | **Post** /vdbs/provision_by_timestamp | Provision a new VDB by timestamp.
[**RefreshVdbBySnapshot**](VDBsApi.md#RefreshVdbBySnapshot) | **Post** /vdbs/{vdbId}/refresh_by_snapshot | Refresh a VDB by snapshot.
[**RefreshVdbByTimestamp**](VDBsApi.md#RefreshVdbByTimestamp) | **Post** /vdbs/{vdbId}/refresh_by_timestamp | Refresh a VDB by timestamp.
[**RollbackVdbBySnapshot**](VDBsApi.md#RollbackVdbBySnapshot) | **Post** /vdbs/{vdbId}/rollback_by_snapshot | Rollback a VDB by snapshot.
[**RollbackVdbByTimestamp**](VDBsApi.md#RollbackVdbByTimestamp) | **Post** /vdbs/{vdbId}/rollback_by_timestamp | Rollback a VDB by timestamp.
[**StartVdb**](VDBsApi.md#StartVdb) | **Post** /vdbs/{vdbId}/start | Start a VDB.
[**StopVdb**](VDBsApi.md#StopVdb) | **Post** /vdbs/{vdbId}/stop | Stop a VDB.



## DeleteVdb

> DeleteVDBResponse DeleteVdb(ctx, vdbId).DeleteVDBParameters(deleteVDBParameters).Execute()

Delete a VDB.

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
    vdbId := "vdbId_example" // string | The ID of the VDB.
    deleteVDBParameters := *openapiclient.NewDeleteVDBParameters() // DeleteVDBParameters | The parameters to delete a VDB. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBsApi.DeleteVdb(context.Background(), vdbId).DeleteVDBParameters(deleteVDBParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBsApi.DeleteVdb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVdb`: DeleteVDBResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBsApi.DeleteVdb`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbId** | **string** | The ID of the VDB. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVdbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteVDBParameters** | [**DeleteVDBParameters**](DeleteVDBParameters.md) | The parameters to delete a VDB. | 

### Return type

[**DeleteVDBResponse**](DeleteVDBResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableVdb

> DisableVDBResponse DisableVdb(ctx, vdbId).DisableVDBParameters(disableVDBParameters).Execute()

Disable a VDB.

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
    vdbId := "vdbId_example" // string | The ID of the VDB.
    disableVDBParameters := *openapiclient.NewDisableVDBParameters() // DisableVDBParameters | The parameters to disable a VDB. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBsApi.DisableVdb(context.Background(), vdbId).DisableVDBParameters(disableVDBParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBsApi.DisableVdb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableVdb`: DisableVDBResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBsApi.DisableVdb`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbId** | **string** | The ID of the VDB. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableVdbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disableVDBParameters** | [**DisableVDBParameters**](DisableVDBParameters.md) | The parameters to disable a VDB. | 

### Return type

[**DisableVDBResponse**](DisableVDBResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableVdb

> EnableVDBResponse EnableVdb(ctx, vdbId).EnableVDBParameters(enableVDBParameters).Execute()

Enable a VDB.

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
    vdbId := "vdbId_example" // string | The ID of the VDB.
    enableVDBParameters := *openapiclient.NewEnableVDBParameters() // EnableVDBParameters | The parameters to enable a VDB. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBsApi.EnableVdb(context.Background(), vdbId).EnableVDBParameters(enableVDBParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBsApi.EnableVdb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableVdb`: EnableVDBResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBsApi.EnableVdb`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbId** | **string** | The ID of the VDB. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableVdbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableVDBParameters** | [**EnableVDBParameters**](EnableVDBParameters.md) | The parameters to enable a VDB. | 

### Return type

[**EnableVDBResponse**](EnableVDBResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVdbById

> VDB GetVdbById(ctx, vdbId).Execute()

Get a VDB by ID.

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
    vdbId := "vdbId_example" // string | The ID of the VDB.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBsApi.GetVdbById(context.Background(), vdbId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBsApi.GetVdbById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVdbById`: VDB
    fmt.Fprintf(os.Stdout, "Response from `VDBsApi.GetVdbById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbId** | **string** | The ID of the VDB. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVdbByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VDB**](VDB.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVdbs

> ListVDBsResponse GetVdbs(ctx).Limit(limit).Cursor(cursor).Execute()

List all vdbs.

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
    limit := int32(50) // int32 | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. (optional) (default to 100)
    cursor := "RXlhbCBpcyBncmVhdAo=" // string | Cursor to fetch the next or previous page of results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBsApi.GetVdbs(context.Background()).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBsApi.GetVdbs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVdbs`: ListVDBsResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBsApi.GetVdbs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVdbsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. | [default to 100]
 **cursor** | **string** | Cursor to fetch the next or previous page of results. | 

### Return type

[**ListVDBsResponse**](ListVDBsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisionVdbBySnapshot

> ProvisionVDBResponse ProvisionVdbBySnapshot(ctx).ProvisionVDBBySnapshotParameters(provisionVDBBySnapshotParameters).Execute()

Provision a new VDB by snapshot.

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
    provisionVDBBySnapshotParameters := *openapiclient.NewProvisionVDBBySnapshotParameters() // ProvisionVDBBySnapshotParameters | The parameters to provision a VDB.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBsApi.ProvisionVdbBySnapshot(context.Background()).ProvisionVDBBySnapshotParameters(provisionVDBBySnapshotParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBsApi.ProvisionVdbBySnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisionVdbBySnapshot`: ProvisionVDBResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBsApi.ProvisionVdbBySnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvisionVdbBySnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provisionVDBBySnapshotParameters** | [**ProvisionVDBBySnapshotParameters**](ProvisionVDBBySnapshotParameters.md) | The parameters to provision a VDB. | 

### Return type

[**ProvisionVDBResponse**](ProvisionVDBResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisionVdbByTimestamp

> ProvisionVDBResponse ProvisionVdbByTimestamp(ctx).ProvisionVDBByTimestampParameters(provisionVDBByTimestampParameters).Execute()

Provision a new VDB by timestamp.

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
    provisionVDBByTimestampParameters := *openapiclient.NewProvisionVDBByTimestampParameters("source-123") // ProvisionVDBByTimestampParameters | The parameters to provision a VDB.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBsApi.ProvisionVdbByTimestamp(context.Background()).ProvisionVDBByTimestampParameters(provisionVDBByTimestampParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBsApi.ProvisionVdbByTimestamp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisionVdbByTimestamp`: ProvisionVDBResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBsApi.ProvisionVdbByTimestamp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvisionVdbByTimestampRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provisionVDBByTimestampParameters** | [**ProvisionVDBByTimestampParameters**](ProvisionVDBByTimestampParameters.md) | The parameters to provision a VDB. | 

### Return type

[**ProvisionVDBResponse**](ProvisionVDBResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshVdbBySnapshot

> RefreshVDBBySnapshotResponse RefreshVdbBySnapshot(ctx, vdbId).RefreshVDBBySnapshotParameters(refreshVDBBySnapshotParameters).Execute()

Refresh a VDB by snapshot.

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
    vdbId := "vdbId_example" // string | The ID of the VDB.
    refreshVDBBySnapshotParameters := *openapiclient.NewRefreshVDBBySnapshotParameters() // RefreshVDBBySnapshotParameters | The parameters to refresh a VDB. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBsApi.RefreshVdbBySnapshot(context.Background(), vdbId).RefreshVDBBySnapshotParameters(refreshVDBBySnapshotParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBsApi.RefreshVdbBySnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshVdbBySnapshot`: RefreshVDBBySnapshotResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBsApi.RefreshVdbBySnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbId** | **string** | The ID of the VDB. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshVdbBySnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refreshVDBBySnapshotParameters** | [**RefreshVDBBySnapshotParameters**](RefreshVDBBySnapshotParameters.md) | The parameters to refresh a VDB. | 

### Return type

[**RefreshVDBBySnapshotResponse**](RefreshVDBBySnapshotResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshVdbByTimestamp

> RefreshVDBByTimestampResponse RefreshVdbByTimestamp(ctx, vdbId).RefreshVDBByTimestampParameters(refreshVDBByTimestampParameters).Execute()

Refresh a VDB by timestamp.

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
    vdbId := "vdbId_example" // string | The ID of the VDB.
    refreshVDBByTimestampParameters := *openapiclient.NewRefreshVDBByTimestampParameters() // RefreshVDBByTimestampParameters | The parameters to refresh a VDB. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBsApi.RefreshVdbByTimestamp(context.Background(), vdbId).RefreshVDBByTimestampParameters(refreshVDBByTimestampParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBsApi.RefreshVdbByTimestamp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshVdbByTimestamp`: RefreshVDBByTimestampResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBsApi.RefreshVdbByTimestamp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbId** | **string** | The ID of the VDB. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshVdbByTimestampRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refreshVDBByTimestampParameters** | [**RefreshVDBByTimestampParameters**](RefreshVDBByTimestampParameters.md) | The parameters to refresh a VDB. | 

### Return type

[**RefreshVDBByTimestampResponse**](RefreshVDBByTimestampResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbackVdbBySnapshot

> RollbackVDBBySnapshotResponse RollbackVdbBySnapshot(ctx, vdbId).RollbackVDBBySnapshotParameters(rollbackVDBBySnapshotParameters).Execute()

Rollback a VDB by snapshot.

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
    vdbId := "vdbId_example" // string | The ID of the VDB.
    rollbackVDBBySnapshotParameters := *openapiclient.NewRollbackVDBBySnapshotParameters() // RollbackVDBBySnapshotParameters | The parameters to rollback a VDB. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBsApi.RollbackVdbBySnapshot(context.Background(), vdbId).RollbackVDBBySnapshotParameters(rollbackVDBBySnapshotParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBsApi.RollbackVdbBySnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RollbackVdbBySnapshot`: RollbackVDBBySnapshotResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBsApi.RollbackVdbBySnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbId** | **string** | The ID of the VDB. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbackVdbBySnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rollbackVDBBySnapshotParameters** | [**RollbackVDBBySnapshotParameters**](RollbackVDBBySnapshotParameters.md) | The parameters to rollback a VDB. | 

### Return type

[**RollbackVDBBySnapshotResponse**](RollbackVDBBySnapshotResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbackVdbByTimestamp

> RollbackVDBByTimestampResponse RollbackVdbByTimestamp(ctx, vdbId).RollbackVDBByTimestampParameters(rollbackVDBByTimestampParameters).Execute()

Rollback a VDB by timestamp.

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
    vdbId := "vdbId_example" // string | The ID of the VDB.
    rollbackVDBByTimestampParameters := *openapiclient.NewRollbackVDBByTimestampParameters() // RollbackVDBByTimestampParameters | The parameters to rollback a VDB. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBsApi.RollbackVdbByTimestamp(context.Background(), vdbId).RollbackVDBByTimestampParameters(rollbackVDBByTimestampParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBsApi.RollbackVdbByTimestamp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RollbackVdbByTimestamp`: RollbackVDBByTimestampResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBsApi.RollbackVdbByTimestamp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbId** | **string** | The ID of the VDB. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbackVdbByTimestampRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rollbackVDBByTimestampParameters** | [**RollbackVDBByTimestampParameters**](RollbackVDBByTimestampParameters.md) | The parameters to rollback a VDB. | 

### Return type

[**RollbackVDBByTimestampResponse**](RollbackVDBByTimestampResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartVdb

> StartVDBResponse StartVdb(ctx, vdbId).Execute()

Start a VDB.

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
    vdbId := "vdbId_example" // string | The ID of the VDB.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBsApi.StartVdb(context.Background(), vdbId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBsApi.StartVdb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartVdb`: StartVDBResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBsApi.StartVdb`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbId** | **string** | The ID of the VDB. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartVdbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StartVDBResponse**](StartVDBResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopVdb

> StopVDBResponse StopVdb(ctx, vdbId).Execute()

Stop a VDB.

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
    vdbId := "vdbId_example" // string | The ID of the VDB.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBsApi.StopVdb(context.Background(), vdbId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBsApi.StopVdb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopVdb`: StopVDBResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBsApi.StopVdb`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbId** | **string** | The ID of the VDB. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopVdbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StopVDBResponse**](StopVDBResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

