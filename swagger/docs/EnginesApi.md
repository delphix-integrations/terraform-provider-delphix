# {{classname}}

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEngineById**](EnginesApi.md#GetEngineById) | **Get** /engines/{engineId} | Returns an engine by ID.
[**GetEngines**](EnginesApi.md#GetEngines) | **Get** /engines | List all engines.

# **GetEngineById**
> Engine GetEngineById(ctx, engineId)
Returns an engine by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **engineId** | **string**| The ID of the engine. | 

### Return type

[**Engine**](Engine.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEngines**
> ListEnginesResponse GetEngines(ctx, )
List all engines.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListEnginesResponse**](ListEnginesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

