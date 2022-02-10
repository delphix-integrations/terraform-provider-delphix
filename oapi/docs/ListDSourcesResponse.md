# ListDSourcesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]DSource**](DSource.md) |  | [optional] 
**Errors** | Pointer to [**[]Error**](Error.md) | Sadly, sometimes requests to the API are not successful. Failures can occur for a wide range of reasons. The Errors object contains information about full or partial failures which might have occurred during the request. | [optional] 
**ResponseMetadata** | Pointer to [**PaginatedResponseMetadata**](PaginatedResponseMetadata.md) |  | [optional] 

## Methods

### NewListDSourcesResponse

`func NewListDSourcesResponse() *ListDSourcesResponse`

NewListDSourcesResponse instantiates a new ListDSourcesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDSourcesResponseWithDefaults

`func NewListDSourcesResponseWithDefaults() *ListDSourcesResponse`

NewListDSourcesResponseWithDefaults instantiates a new ListDSourcesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListDSourcesResponse) GetItems() []DSource`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListDSourcesResponse) GetItemsOk() (*[]DSource, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListDSourcesResponse) SetItems(v []DSource)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListDSourcesResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetErrors

`func (o *ListDSourcesResponse) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ListDSourcesResponse) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ListDSourcesResponse) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ListDSourcesResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetResponseMetadata

`func (o *ListDSourcesResponse) GetResponseMetadata() PaginatedResponseMetadata`

GetResponseMetadata returns the ResponseMetadata field if non-nil, zero value otherwise.

### GetResponseMetadataOk

`func (o *ListDSourcesResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool)`

GetResponseMetadataOk returns a tuple with the ResponseMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMetadata

`func (o *ListDSourcesResponse) SetResponseMetadata(v PaginatedResponseMetadata)`

SetResponseMetadata sets ResponseMetadata field to given value.

### HasResponseMetadata

`func (o *ListDSourcesResponse) HasResponseMetadata() bool`

HasResponseMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


