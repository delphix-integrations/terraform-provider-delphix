# \BookmarksApi

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBookmark**](BookmarksApi.md#CreateBookmark) | **Post** /bookmarks | Create a bookmark at the current time.
[**DeleteBookmark**](BookmarksApi.md#DeleteBookmark) | **Delete** /bookmarks/{bookmarkId} | Delete a bookmark.
[**GetBookmarkById**](BookmarksApi.md#GetBookmarkById) | **Get** /bookmarks/{bookmarkId} | Get a bookmark by ID.
[**GetBookmarks**](BookmarksApi.md#GetBookmarks) | **Get** /bookmarks | List all bookmarks.
[**RestoreBookmark**](BookmarksApi.md#RestoreBookmark) | **Post** /bookmarks/{bookmarkId}/restore | Restore VDBs to the bookmark creation time.



## CreateBookmark

> CreateBookmarkResponse CreateBookmark(ctx).Bookmark(bookmark).Execute()

Create a bookmark at the current time.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bookmark := *openapiclient.NewBookmark("my-bookmark-123", []string{"VdbIds_example"}) // Bookmark | The parameters to create a bookmark.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BookmarksApi.CreateBookmark(context.Background()).Bookmark(bookmark).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookmarksApi.CreateBookmark``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBookmark`: CreateBookmarkResponse
    fmt.Fprintf(os.Stdout, "Response from `BookmarksApi.CreateBookmark`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBookmarkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bookmark** | [**Bookmark**](Bookmark.md) | The parameters to create a bookmark. | 

### Return type

[**CreateBookmarkResponse**](CreateBookmarkResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBookmark

> DeleteBookmark(ctx, bookmarkId).Execute()

Delete a bookmark.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bookmarkId := "bookmarkId_example" // string | The ID of the bookmark.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BookmarksApi.DeleteBookmark(context.Background(), bookmarkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookmarksApi.DeleteBookmark``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookmarkId** | **string** | The ID of the bookmark. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBookmarkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBookmarkById

> Bookmark GetBookmarkById(ctx, bookmarkId).Execute()

Get a bookmark by ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bookmarkId := "bookmarkId_example" // string | The ID of the bookmark.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BookmarksApi.GetBookmarkById(context.Background(), bookmarkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookmarksApi.GetBookmarkById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBookmarkById`: Bookmark
    fmt.Fprintf(os.Stdout, "Response from `BookmarksApi.GetBookmarkById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookmarkId** | **string** | The ID of the bookmark. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBookmarkByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Bookmark**](Bookmark.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBookmarks

> ListBookmarksResponse GetBookmarks(ctx).Execute()

List all bookmarks.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BookmarksApi.GetBookmarks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookmarksApi.GetBookmarks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBookmarks`: ListBookmarksResponse
    fmt.Fprintf(os.Stdout, "Response from `BookmarksApi.GetBookmarks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBookmarksRequest struct via the builder pattern


### Return type

[**ListBookmarksResponse**](ListBookmarksResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreBookmark

> RestoreBookmarkResponse RestoreBookmark(ctx, bookmarkId).Execute()

Restore VDBs to the bookmark creation time.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bookmarkId := "bookmarkId_example" // string | The ID of the bookmark.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BookmarksApi.RestoreBookmark(context.Background(), bookmarkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookmarksApi.RestoreBookmark``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestoreBookmark`: RestoreBookmarkResponse
    fmt.Fprintf(os.Stdout, "Response from `BookmarksApi.RestoreBookmark`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookmarkId** | **string** | The ID of the bookmark. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreBookmarkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestoreBookmarkResponse**](RestoreBookmarkResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

