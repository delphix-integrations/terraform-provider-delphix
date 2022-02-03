# {{classname}}

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSnapshotById**](SnapshotsApi.md#GetSnapshotById) | **Get** /snapshots/{snapshotId} | Get a Snapshot by ID.
[**GetSnapshots**](SnapshotsApi.md#GetSnapshots) | **Get** /snapshots | List Snapshots for a dSource or VDB.

# **GetSnapshotById**
> Snapshot GetSnapshotById(ctx, snapshotId)
Get a Snapshot by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotId** | **string**| The ID of the snaphost. | 

### Return type

[**Snapshot**](Snapshot.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSnapshots**
> ListSnaphotsResponse GetSnapshots(ctx, optional)
List Snapshots for a dSource or VDB.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SnapshotsApiGetSnapshotsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotsApiGetSnapshotsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasetId** | **optional.String**| The ID of the dSource or VDB for which to fetch Snapshots. | 
 **limit** | **optional.Int32**| Maximum number of objects to return per query. The value must be between 1 and 1000. Default is 100. | [default to 100]
 **cursor** | **optional.String**| Cursor to fetch the next or previous page of results. | 

### Return type

[**ListSnaphotsResponse**](ListSnaphotsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

