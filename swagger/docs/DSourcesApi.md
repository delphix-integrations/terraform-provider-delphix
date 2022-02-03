# {{classname}}

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDsourceById**](DSourcesApi.md#GetDsourceById) | **Get** /dsources/{dsourceId} | Get a dSource by ID.
[**GetDsources**](DSourcesApi.md#GetDsources) | **Get** /dsources | List all dSources.

# **GetDsourceById**
> DSource GetDsourceById(ctx, dsourceId)
Get a dSource by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dsourceId** | **string**| The ID of the dSource. | 

### Return type

[**DSource**](DSource.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDsources**
> ListDSourcesResponse GetDsources(ctx, )
List all dSources.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListDSourcesResponse**](ListDSourcesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

