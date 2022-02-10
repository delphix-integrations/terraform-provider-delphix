# CreateBookmarkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bookmark** | Pointer to [**Bookmark**](Bookmark.md) |  | [optional] 
**JobId** | Pointer to **string** | The initiated job id. | [optional] 

## Methods

### NewCreateBookmarkResponse

`func NewCreateBookmarkResponse() *CreateBookmarkResponse`

NewCreateBookmarkResponse instantiates a new CreateBookmarkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBookmarkResponseWithDefaults

`func NewCreateBookmarkResponseWithDefaults() *CreateBookmarkResponse`

NewCreateBookmarkResponseWithDefaults instantiates a new CreateBookmarkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBookmark

`func (o *CreateBookmarkResponse) GetBookmark() Bookmark`

GetBookmark returns the Bookmark field if non-nil, zero value otherwise.

### GetBookmarkOk

`func (o *CreateBookmarkResponse) GetBookmarkOk() (*Bookmark, bool)`

GetBookmarkOk returns a tuple with the Bookmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmark

`func (o *CreateBookmarkResponse) SetBookmark(v Bookmark)`

SetBookmark sets Bookmark field to given value.

### HasBookmark

`func (o *CreateBookmarkResponse) HasBookmark() bool`

HasBookmark returns a boolean if a field has been set.

### GetJobId

`func (o *CreateBookmarkResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *CreateBookmarkResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *CreateBookmarkResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *CreateBookmarkResponse) HasJobId() bool`

HasJobId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


