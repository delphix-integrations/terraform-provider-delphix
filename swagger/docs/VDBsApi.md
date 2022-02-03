# {{classname}}

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVdb**](VDBsApi.md#DeleteVdb) | **Post** /vdbs/{vdbId}/delete | Delete a VDB.
[**DisableVdb**](VDBsApi.md#DisableVdb) | **Post** /vdbs/{vdbId}/disable | Disable a VDB.
[**EnableVdb**](VDBsApi.md#EnableVdb) | **Post** /vdbs/{vdbId}/enable | Enable a VDB.
[**GetVdbById**](VDBsApi.md#GetVdbById) | **Get** /vdbs/{vdbId} | Get a VDB by ID.
[**GetVdbs**](VDBsApi.md#GetVdbs) | **Get** /vdbs | List all vdbs.
[**ProvisionVdbBySnapshot**](VDBsApi.md#ProvisionVdbBySnapshot) | **Post** /vdbs/provision_by_snapshot | Provision a new VDB by snapshot.
[**ProvisionVdbByTimestamp**](VDBsApi.md#ProvisionVdbByTimestamp) | **Post** /vdbs/provision_by_timestamp | Provision a new VDB by timestamp.
[**RefreshVdbBySnapshot**](VDBsApi.md#RefreshVdbBySnapshot) | **Post** /vdbs/{vdbId}/refresh_by_snapshot | Refresh a VDB by snapshot.
[**RefreshVdbByTimestamp**](VDBsApi.md#RefreshVdbByTimestamp) | **Post** /vdbs/{vdbId}/refresh_by_timestamp | Refresh a VDB by timestamp.
[**RollbackVdbBySnapshot**](VDBsApi.md#RollbackVdbBySnapshot) | **Post** /vdbs/{vdbId}/rollback_by_snapshot | Rollback a VDB by snapshot.
[**RollbackVdbByTimestamp**](VDBsApi.md#RollbackVdbByTimestamp) | **Post** /vdbs/{vdbId}/rollback_by_timestamp | Rollback a VDB by timestamp.
[**StartVdb**](VDBsApi.md#StartVdb) | **Post** /vdbs/{vdbId}/start | Start a VDB.
[**StopVdb**](VDBsApi.md#StopVdb) | **Post** /vdbs/{vdbId}/stop | Stop a VDB.

# **DeleteVdb**
> DeleteVdbResponse DeleteVdb(ctx, vdbId, optional)
Delete a VDB.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdbId** | **string**| The ID of the VDB. | 
 **optional** | ***VDBsApiDeleteVdbOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VDBsApiDeleteVdbOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DeleteVdbParameters**](DeleteVdbParameters.md)| The parameters to delete a VDB. | 

### Return type

[**DeleteVdbResponse**](DeleteVDBResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableVdb**
> DisableVdbResponse DisableVdb(ctx, vdbId, optional)
Disable a VDB.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdbId** | **string**| The ID of the VDB. | 
 **optional** | ***VDBsApiDisableVdbOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VDBsApiDisableVdbOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DisableVdbParameters**](DisableVdbParameters.md)| The parameters to disable a VDB. | 

### Return type

[**DisableVdbResponse**](DisableVDBResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableVdb**
> EnableVdbResponse EnableVdb(ctx, vdbId, optional)
Enable a VDB.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdbId** | **string**| The ID of the VDB. | 
 **optional** | ***VDBsApiEnableVdbOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VDBsApiEnableVdbOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EnableVdbParameters**](EnableVdbParameters.md)| The parameters to enable a VDB. | 

### Return type

[**EnableVdbResponse**](EnableVDBResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVdbById**
> Vdb GetVdbById(ctx, vdbId)
Get a VDB by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdbId** | **string**| The ID of the VDB. | 

### Return type

[**Vdb**](VDB.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVdbs**
> ListVdBsResponse GetVdbs(ctx, )
List all vdbs.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListVdBsResponse**](ListVDBsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProvisionVdbBySnapshot**
> ProvisionVdbResponse ProvisionVdbBySnapshot(ctx, body)
Provision a new VDB by snapshot.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProvisionVdbBySnapshotParameters**](ProvisionVdbBySnapshotParameters.md)| The parameters to provision a VDB. | 

### Return type

[**ProvisionVdbResponse**](ProvisionVDBResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProvisionVdbByTimestamp**
> ProvisionVdbResponse ProvisionVdbByTimestamp(ctx, body)
Provision a new VDB by timestamp.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProvisionVdbByTimestampParameters**](ProvisionVdbByTimestampParameters.md)| The parameters to provision a VDB. | 

### Return type

[**ProvisionVdbResponse**](ProvisionVDBResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshVdbBySnapshot**
> RefreshVdbBySnapshotResponse RefreshVdbBySnapshot(ctx, vdbId, optional)
Refresh a VDB by snapshot.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdbId** | **string**| The ID of the VDB. | 
 **optional** | ***VDBsApiRefreshVdbBySnapshotOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VDBsApiRefreshVdbBySnapshotOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RefreshVdbBySnapshotParameters**](RefreshVdbBySnapshotParameters.md)| The parameters to refresh a VDB. | 

### Return type

[**RefreshVdbBySnapshotResponse**](RefreshVDBBySnapshotResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshVdbByTimestamp**
> RefreshVdbByTimestampResponse RefreshVdbByTimestamp(ctx, vdbId, optional)
Refresh a VDB by timestamp.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdbId** | **string**| The ID of the VDB. | 
 **optional** | ***VDBsApiRefreshVdbByTimestampOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VDBsApiRefreshVdbByTimestampOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RefreshVdbByTimestampParameters**](RefreshVdbByTimestampParameters.md)| The parameters to refresh a VDB. | 

### Return type

[**RefreshVdbByTimestampResponse**](RefreshVDBByTimestampResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RollbackVdbBySnapshot**
> RollbackVdbBySnapshotResponse RollbackVdbBySnapshot(ctx, vdbId, optional)
Rollback a VDB by snapshot.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdbId** | **string**| The ID of the VDB. | 
 **optional** | ***VDBsApiRollbackVdbBySnapshotOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VDBsApiRollbackVdbBySnapshotOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RollbackVdbBySnapshotParameters**](RollbackVdbBySnapshotParameters.md)| The parameters to rollback a VDB. | 

### Return type

[**RollbackVdbBySnapshotResponse**](RollbackVDBBySnapshotResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RollbackVdbByTimestamp**
> RollbackVdbByTimestampResponse RollbackVdbByTimestamp(ctx, vdbId, optional)
Rollback a VDB by timestamp.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdbId** | **string**| The ID of the VDB. | 
 **optional** | ***VDBsApiRollbackVdbByTimestampOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VDBsApiRollbackVdbByTimestampOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RollbackVdbByTimestampParameters**](RollbackVdbByTimestampParameters.md)| The parameters to rollback a VDB. | 

### Return type

[**RollbackVdbByTimestampResponse**](RollbackVDBByTimestampResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartVdb**
> StartVdbResponse StartVdb(ctx, vdbId)
Start a VDB.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdbId** | **string**| The ID of the VDB. | 

### Return type

[**StartVdbResponse**](StartVDBResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopVdb**
> StopVdbResponse StopVdb(ctx, vdbId)
Stop a VDB.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdbId** | **string**| The ID of the VDB. | 

### Return type

[**StopVdbResponse**](StopVDBResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

