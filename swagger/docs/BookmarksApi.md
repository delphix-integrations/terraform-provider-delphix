# {{classname}}

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBookmark**](BookmarksApi.md#CreateBookmark) | **Post** /bookmarks | Create a bookmark at the current time.
[**DeleteBookmark**](BookmarksApi.md#DeleteBookmark) | **Delete** /bookmarks/{bookmarkId} | Delete a bookmark.
[**GetBookmarkById**](BookmarksApi.md#GetBookmarkById) | **Get** /bookmarks/{bookmarkId} | Get a bookmark by ID.
[**GetBookmarks**](BookmarksApi.md#GetBookmarks) | **Get** /bookmarks | List all bookmarks.
[**RestoreBookmark**](BookmarksApi.md#RestoreBookmark) | **Post** /bookmarks/{bookmarkId}/restore | Restore VDBs to the bookmark creation time.

# **CreateBookmark**
> CreateBookmarkResponse CreateBookmark(ctx, body)
Create a bookmark at the current time.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Bookmark**](Bookmark.md)| The parameters to create a bookmark. | 

### Return type

[**CreateBookmarkResponse**](CreateBookmarkResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBookmark**
> DeleteBookmark(ctx, bookmarkId)
Delete a bookmark.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bookmarkId** | **string**| The ID of the bookmark. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBookmarkById**
> Bookmark GetBookmarkById(ctx, bookmarkId)
Get a bookmark by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bookmarkId** | **string**| The ID of the bookmark. | 

### Return type

[**Bookmark**](Bookmark.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBookmarks**
> ListBookmarksResponse GetBookmarks(ctx, )
List all bookmarks.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListBookmarksResponse**](ListBookmarksResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreBookmark**
> RestoreBookmarkResponse RestoreBookmark(ctx, bookmarkId)
Restore VDBs to the bookmark creation time.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bookmarkId** | **string**| The ID of the bookmark. | 

### Return type

[**RestoreBookmarkResponse**](RestoreBookmarkResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

