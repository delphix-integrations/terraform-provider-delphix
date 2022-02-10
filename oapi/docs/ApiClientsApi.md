# \ApiClientsApi

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiClient**](ApiClientsApi.md#CreateApiClient) | **Post** /management/api-clients | Create a new API Client. 
[**DeleteApiClient**](ApiClientsApi.md#DeleteApiClient) | **Delete** /management/api-clients/{id} | Delete an API client
[**GetApiClient**](ApiClientsApi.md#GetApiClient) | **Get** /management/api-clients/{id} | Get an API client by id
[**GetApiClients**](ApiClientsApi.md#GetApiClients) | **Get** /management/api-clients | Returns a list of API clients.
[**UpdateApiClient**](ApiClientsApi.md#UpdateApiClient) | **Put** /management/api-clients/{id} | Update an Api client. 



## CreateApiClient

> ApiClientCreateResponse CreateApiClient(ctx).ApiClientCreateParameter(apiClientCreateParameter).Execute()

Create a new API Client. 

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
    apiClientCreateParameter := *openapiclient.NewApiClientCreateParameter() // ApiClientCreateParameter | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiClientsApi.CreateApiClient(context.Background()).ApiClientCreateParameter(apiClientCreateParameter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiClientsApi.CreateApiClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiClient`: ApiClientCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `ApiClientsApi.CreateApiClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiClientCreateParameter** | [**ApiClientCreateParameter**](ApiClientCreateParameter.md) |  | 

### Return type

[**ApiClientCreateResponse**](ApiClientCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiClient

> DeleteApiClient(ctx, id).Execute()

Delete an API client

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
    id := int64(789) // int64 | Numeric ID of the Api client

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiClientsApi.DeleteApiClient(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiClientsApi.DeleteApiClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Numeric ID of the Api client | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiClient

> ApiClient GetApiClient(ctx, id).Execute()

Get an API client by id

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
    id := int64(789) // int64 | Numeric ID of the Api client

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiClientsApi.GetApiClient(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiClientsApi.GetApiClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiClient`: ApiClient
    fmt.Fprintf(os.Stdout, "Response from `ApiClientsApi.GetApiClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Numeric ID of the Api client | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiClient**](ApiClient.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiClients

> []ApiClient GetApiClients(ctx).Execute()

Returns a list of API clients.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiClientsApi.GetApiClients(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiClientsApi.GetApiClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiClients`: []ApiClient
    fmt.Fprintf(os.Stdout, "Response from `ApiClientsApi.GetApiClients`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiClientsRequest struct via the builder pattern


### Return type

[**[]ApiClient**](ApiClient.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiClient

> ApiClient UpdateApiClient(ctx, id).ApiClient(apiClient).Execute()

Update an Api client. 

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
    id := int64(789) // int64 | Numeric ID of the Api client
    apiClient := *openapiclient.NewApiClient(false) // ApiClient | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiClientsApi.UpdateApiClient(context.Background(), id).ApiClient(apiClient).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiClientsApi.UpdateApiClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiClient`: ApiClient
    fmt.Fprintf(os.Stdout, "Response from `ApiClientsApi.UpdateApiClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Numeric ID of the Api client | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiClient** | [**ApiClient**](ApiClient.md) |  | 

### Return type

[**ApiClient**](ApiClient.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

