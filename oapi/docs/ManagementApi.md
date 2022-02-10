# \ManagementApi

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRegisteredEngineUser**](ManagementApi.md#AddRegisteredEngineUser) | **Post** /management/engines/{engineId}/users | Associate a new engine user to a registered engine.
[**CreateHashicorpVault**](ManagementApi.md#CreateHashicorpVault) | **Post** /management/vaults/hashicorp | Configure a new Hashicorp Vault
[**DeleteHashicorpVault**](ManagementApi.md#DeleteHashicorpVault) | **Delete** /management/vaults/hashicorp/{vaultId} | Delete a Hashicorp vault by id
[**DeleteRegisteredEngineUser**](ManagementApi.md#DeleteRegisteredEngineUser) | **Delete** /management/engines/{engineId}/users/{userId} | Remove a user from the list of users associated to a registered engine.
[**GetHashicorpVault**](ManagementApi.md#GetHashicorpVault) | **Get** /management/vaults/hashicorp/{vaultId} | Get a Hashicorp vault by id
[**GetHashicorpVaults**](ManagementApi.md#GetHashicorpVaults) | **Get** /management/vaults/hashicorp | Returns a list of configured Hashicorp vaults.
[**GetRegisteredEngine**](ManagementApi.md#GetRegisteredEngine) | **Get** /management/engines/{engineId} | Returns a registered engine by ID.
[**GetRegisteredEngineUsers**](ManagementApi.md#GetRegisteredEngineUsers) | **Get** /management/engines/{engineId}/users | Returns the list of users associated to an registered engine.
[**GetRegisteredEngines**](ManagementApi.md#GetRegisteredEngines) | **Get** /management/engines | Returns a list of registered engines.
[**RegisterEngine**](ManagementApi.md#RegisterEngine) | **Post** /management/engines | Register an engine.
[**UnregisterEngine**](ManagementApi.md#UnregisterEngine) | **Delete** /management/engines/{engineId} | Unregister an engine.
[**UpdateRegisteredEngine**](ManagementApi.md#UpdateRegisteredEngine) | **Put** /management/engines/{engineId} | Update a registered engine.



## AddRegisteredEngineUser

> RegisteredEngineUser AddRegisteredEngineUser(ctx, engineId).RegisteredEngineUser(registeredEngineUser).Execute()

Associate a new engine user to a registered engine.

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
    engineId := int64(789) // int64 | Numeric ID of the registered engine.
    registeredEngineUser := *openapiclient.NewRegisteredEngineUser() // RegisteredEngineUser |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.AddRegisteredEngineUser(context.Background(), engineId).RegisteredEngineUser(registeredEngineUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.AddRegisteredEngineUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRegisteredEngineUser`: RegisteredEngineUser
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.AddRegisteredEngineUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**engineId** | **int64** | Numeric ID of the registered engine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRegisteredEngineUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registeredEngineUser** | [**RegisteredEngineUser**](RegisteredEngineUser.md) |  | 

### Return type

[**RegisteredEngineUser**](RegisteredEngineUser.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHashicorpVault

> HashicorpVault CreateHashicorpVault(ctx).HashicorpVault(hashicorpVault).Execute()

Configure a new Hashicorp Vault

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
    hashicorpVault := *openapiclient.NewHashicorpVault() // HashicorpVault | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.CreateHashicorpVault(context.Background()).HashicorpVault(hashicorpVault).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.CreateHashicorpVault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHashicorpVault`: HashicorpVault
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.CreateHashicorpVault`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHashicorpVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hashicorpVault** | [**HashicorpVault**](HashicorpVault.md) |  | 

### Return type

[**HashicorpVault**](HashicorpVault.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHashicorpVault

> DeleteHashicorpVault(ctx, vaultId).Execute()

Delete a Hashicorp vault by id

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
    vaultId := int64(789) // int64 | Numeric ID of the Hashicorp vault

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.DeleteHashicorpVault(context.Background(), vaultId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.DeleteHashicorpVault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **int64** | Numeric ID of the Hashicorp vault | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHashicorpVaultRequest struct via the builder pattern


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


## DeleteRegisteredEngineUser

> DeleteRegisteredEngineUser(ctx, engineId, userId).Execute()

Remove a user from the list of users associated to a registered engine.

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
    engineId := int64(789) // int64 | Numeric ID of the registered engine.
    userId := int64(789) // int64 | The ID of the registered engine user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.DeleteRegisteredEngineUser(context.Background(), engineId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.DeleteRegisteredEngineUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**engineId** | **int64** | Numeric ID of the registered engine. | 
**userId** | **int64** | The ID of the registered engine user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRegisteredEngineUserRequest struct via the builder pattern


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


## GetHashicorpVault

> HashicorpVault GetHashicorpVault(ctx, vaultId).Execute()

Get a Hashicorp vault by id

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
    vaultId := int64(789) // int64 | Numeric ID of the Hashicorp vault

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.GetHashicorpVault(context.Background(), vaultId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.GetHashicorpVault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHashicorpVault`: HashicorpVault
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.GetHashicorpVault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **int64** | Numeric ID of the Hashicorp vault | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHashicorpVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HashicorpVault**](HashicorpVault.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHashicorpVaults

> []HashicorpVault GetHashicorpVaults(ctx).Execute()

Returns a list of configured Hashicorp vaults.

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
    resp, r, err := apiClient.ManagementApi.GetHashicorpVaults(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.GetHashicorpVaults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHashicorpVaults`: []HashicorpVault
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.GetHashicorpVaults`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHashicorpVaultsRequest struct via the builder pattern


### Return type

[**[]HashicorpVault**](HashicorpVault.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegisteredEngine

> RegisteredEngine GetRegisteredEngine(ctx, engineId).Execute()

Returns a registered engine by ID.

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
    engineId := int64(789) // int64 | Numeric ID of the registered engine.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.GetRegisteredEngine(context.Background(), engineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.GetRegisteredEngine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegisteredEngine`: RegisteredEngine
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.GetRegisteredEngine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**engineId** | **int64** | Numeric ID of the registered engine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegisteredEngineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RegisteredEngine**](RegisteredEngine.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegisteredEngineUsers

> []RegisteredEngineUser GetRegisteredEngineUsers(ctx, engineId).Execute()

Returns the list of users associated to an registered engine.

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
    engineId := int64(789) // int64 | Numeric ID of the registered engine.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.GetRegisteredEngineUsers(context.Background(), engineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.GetRegisteredEngineUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegisteredEngineUsers`: []RegisteredEngineUser
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.GetRegisteredEngineUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**engineId** | **int64** | Numeric ID of the registered engine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegisteredEngineUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]RegisteredEngineUser**](RegisteredEngineUser.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegisteredEngines

> []RegisteredEngine GetRegisteredEngines(ctx).Execute()

Returns a list of registered engines.

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
    resp, r, err := apiClient.ManagementApi.GetRegisteredEngines(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.GetRegisteredEngines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegisteredEngines`: []RegisteredEngine
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.GetRegisteredEngines`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegisteredEnginesRequest struct via the builder pattern


### Return type

[**[]RegisteredEngine**](RegisteredEngine.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterEngine

> RegisteredEngine RegisterEngine(ctx).EngineRegistrationParameter(engineRegistrationParameter).Execute()

Register an engine.

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
    engineRegistrationParameter := *openapiclient.NewEngineRegistrationParameter("Name_example", "Hostname_example") // EngineRegistrationParameter | The parameters to register an engine.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.RegisterEngine(context.Background()).EngineRegistrationParameter(engineRegistrationParameter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.RegisterEngine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterEngine`: RegisteredEngine
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.RegisterEngine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterEngineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **engineRegistrationParameter** | [**EngineRegistrationParameter**](EngineRegistrationParameter.md) | The parameters to register an engine. | 

### Return type

[**RegisteredEngine**](RegisteredEngine.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnregisterEngine

> DeleteEngineResponse UnregisterEngine(ctx, engineId).Execute()

Unregister an engine.

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
    engineId := int64(789) // int64 | Numeric ID of the registered engine.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.UnregisterEngine(context.Background(), engineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.UnregisterEngine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnregisterEngine`: DeleteEngineResponse
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.UnregisterEngine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**engineId** | **int64** | Numeric ID of the registered engine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterEngineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteEngineResponse**](DeleteEngineResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRegisteredEngine

> RegisteredEngine UpdateRegisteredEngine(ctx, engineId).RegisteredEngine(registeredEngine).Execute()

Update a registered engine.

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
    engineId := int64(789) // int64 | Numeric ID of the registered engine.
    registeredEngine := *openapiclient.NewRegisteredEngine("Name_example", "Hostname_example") // RegisteredEngine | The updated registration engine information.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementApi.UpdateRegisteredEngine(context.Background(), engineId).RegisteredEngine(registeredEngine).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.UpdateRegisteredEngine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRegisteredEngine`: RegisteredEngine
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.UpdateRegisteredEngine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**engineId** | **int64** | Numeric ID of the registered engine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRegisteredEngineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registeredEngine** | [**RegisteredEngine**](RegisteredEngine.md) | The updated registration engine information. | 

### Return type

[**RegisteredEngine**](RegisteredEngine.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

