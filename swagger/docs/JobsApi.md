# {{classname}}

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetJobById**](JobsApi.md#GetJobById) | **Get** /jobs/{jobId} | Returns an job by ID.

# **GetJobById**
> Job GetJobById(ctx, jobId)
Returns an job by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**| The ID of the job. | 

### Return type

[**Job**](Job.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

