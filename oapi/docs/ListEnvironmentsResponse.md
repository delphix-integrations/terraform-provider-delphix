# ListEnvironmentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Environment**](Environment.md) |  | [optional] 
**Errors** | Pointer to [**[]Error**](Error.md) | Sadly, sometimes requests to the API are not successful. Failures can occur for a wide range of reasons. The Errors object contains information about full or partial failures which might have occurred during the request. | [optional] 
**ResponseMetadata** | Pointer to [**PaginatedResponseMetadata**](PaginatedResponseMetadata.md) |  | [optional] 

## Methods

### NewListEnvironmentsResponse

`func NewListEnvironmentsResponse() *ListEnvironmentsResponse`

NewListEnvironmentsResponse instantiates a new ListEnvironmentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEnvironmentsResponseWithDefaults

`func NewListEnvironmentsResponseWithDefaults() *ListEnvironmentsResponse`

NewListEnvironmentsResponseWithDefaults instantiates a new ListEnvironmentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListEnvironmentsResponse) GetItems() []Environment`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListEnvironmentsResponse) GetItemsOk() (*[]Environment, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListEnvironmentsResponse) SetItems(v []Environment)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListEnvironmentsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetErrors

`func (o *ListEnvironmentsResponse) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ListEnvironmentsResponse) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ListEnvironmentsResponse) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ListEnvironmentsResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetResponseMetadata

`func (o *ListEnvironmentsResponse) GetResponseMetadata() PaginatedResponseMetadata`

GetResponseMetadata returns the ResponseMetadata field if non-nil, zero value otherwise.

### GetResponseMetadataOk

`func (o *ListEnvironmentsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool)`

GetResponseMetadataOk returns a tuple with the ResponseMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMetadata

`func (o *ListEnvironmentsResponse) SetResponseMetadata(v PaginatedResponseMetadata)`

SetResponseMetadata sets ResponseMetadata field to given value.

### HasResponseMetadata

`func (o *ListEnvironmentsResponse) HasResponseMetadata() bool`

HasResponseMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


