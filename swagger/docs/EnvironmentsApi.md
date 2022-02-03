# {{classname}}

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEnvironmentById**](EnvironmentsApi.md#GetEnvironmentById) | **Get** /environments/{environmentId} | Returns an environment by ID.
[**GetEnvironments**](EnvironmentsApi.md#GetEnvironments) | **Get** /environments | List all environments.

# **GetEnvironmentById**
> Environment GetEnvironmentById(ctx, environmentId)
Returns an environment by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environmentId** | **string**| The ID of the environment. | 

### Return type

[**Environment**](Environment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvironments**
> ListEnvironmentsResponse GetEnvironments(ctx, )
List all environments.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListEnvironmentsResponse**](ListEnvironmentsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

