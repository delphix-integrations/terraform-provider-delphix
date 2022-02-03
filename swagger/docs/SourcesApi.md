# {{classname}}

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSourceById**](SourcesApi.md#GetSourceById) | **Get** /sources/{sourceId} | Get a source by ID.
[**GetSources**](SourcesApi.md#GetSources) | **Get** /sources | List all sources.

# **GetSourceById**
> Source GetSourceById(ctx, sourceId)
Get a source by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sourceId** | **string**| The ID of the source. | 

### Return type

[**Source**](Source.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSources**
> ListSourcesResponse GetSources(ctx, )
List all sources.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListSourcesResponse**](ListSourcesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

