# {{classname}}

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

# **AddRegisteredEngineUser**
> RegisteredEngineUser AddRegisteredEngineUser(ctx, engineId, optional)
Associate a new engine user to a registered engine.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **engineId** | **int64**| Numeric ID of the registered engine. | 
 **optional** | ***ManagementApiAddRegisteredEngineUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementApiAddRegisteredEngineUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RegisteredEngineUser**](RegisteredEngineUser.md)|  | 

### Return type

[**RegisteredEngineUser**](RegisteredEngineUser.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateHashicorpVault**
> HashicorpVault CreateHashicorpVault(ctx, body)
Configure a new Hashicorp Vault

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HashicorpVault**](HashicorpVault.md)|  | 

### Return type

[**HashicorpVault**](HashicorpVault.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHashicorpVault**
> DeleteHashicorpVault(ctx, vaultId)
Delete a Hashicorp vault by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **int64**| Numeric ID of the Hashicorp vault | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRegisteredEngineUser**
> DeleteRegisteredEngineUser(ctx, engineId, userId)
Remove a user from the list of users associated to a registered engine.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **engineId** | **int64**| Numeric ID of the registered engine. | 
  **userId** | **int64**| The ID of the registered engine user. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHashicorpVault**
> HashicorpVault GetHashicorpVault(ctx, vaultId)
Get a Hashicorp vault by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vaultId** | **int64**| Numeric ID of the Hashicorp vault | 

### Return type

[**HashicorpVault**](HashicorpVault.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHashicorpVaults**
> []HashicorpVault GetHashicorpVaults(ctx, )
Returns a list of configured Hashicorp vaults.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]HashicorpVault**](HashicorpVault.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRegisteredEngine**
> RegisteredEngine GetRegisteredEngine(ctx, engineId)
Returns a registered engine by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **engineId** | **int64**| Numeric ID of the registered engine. | 

### Return type

[**RegisteredEngine**](RegisteredEngine.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRegisteredEngineUsers**
> []RegisteredEngineUser GetRegisteredEngineUsers(ctx, engineId)
Returns the list of users associated to an registered engine.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **engineId** | **int64**| Numeric ID of the registered engine. | 

### Return type

[**[]RegisteredEngineUser**](RegisteredEngineUser.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRegisteredEngines**
> []RegisteredEngine GetRegisteredEngines(ctx, )
Returns a list of registered engines.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]RegisteredEngine**](RegisteredEngine.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterEngine**
> RegisteredEngine RegisterEngine(ctx, body)
Register an engine.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EngineRegistrationParameter**](EngineRegistrationParameter.md)| The parameters to register an engine. | 

### Return type

[**RegisteredEngine**](RegisteredEngine.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnregisterEngine**
> UnregisterEngine(ctx, engineId)
Unregister an engine.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **engineId** | **int64**| Numeric ID of the registered engine. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRegisteredEngine**
> RegisteredEngine UpdateRegisteredEngine(ctx, body, engineId)
Update a registered engine.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RegisteredEngine**](RegisteredEngine.md)| The updated registration engine information. | 
  **engineId** | **int64**| Numeric ID of the registered engine. | 

### Return type

[**RegisteredEngine**](RegisteredEngine.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

