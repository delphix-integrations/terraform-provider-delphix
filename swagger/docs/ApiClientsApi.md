# {{classname}}

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiClient**](ApiClientsApi.md#CreateApiClient) | **Post** /management/api-clients | Create a new API Client. 
[**DeleteApiClient**](ApiClientsApi.md#DeleteApiClient) | **Delete** /management/api-clients/{id} | Delete an API client
[**GetApiClient**](ApiClientsApi.md#GetApiClient) | **Get** /management/api-clients/{id} | Get an API client by id
[**GetApiClients**](ApiClientsApi.md#GetApiClients) | **Get** /management/api-clients | Returns a list of API clients.
[**UpdateApiClient**](ApiClientsApi.md#UpdateApiClient) | **Put** /management/api-clients/{id} | Update an Api client. 

# **CreateApiClient**
> ApiClientCreateResponse CreateApiClient(ctx, body)
Create a new API Client. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiClientCreateParameter**](ApiClientCreateParameter.md)|  | 

### Return type

[**ApiClientCreateResponse**](ApiClientCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiClient**
> DeleteApiClient(ctx, id)
Delete an API client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Numeric ID of the Api client | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiClient**
> ApiClient GetApiClient(ctx, id)
Get an API client by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Numeric ID of the Api client | 

### Return type

[**ApiClient**](ApiClient.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiClients**
> []ApiClient GetApiClients(ctx, )
Returns a list of API clients.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ApiClient**](ApiClient.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApiClient**
> ApiClient UpdateApiClient(ctx, body, id)
Update an Api client. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiClient**](ApiClient.md)|  | 
  **id** | **int64**| Numeric ID of the Api client | 

### Return type

[**ApiClient**](ApiClient.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

