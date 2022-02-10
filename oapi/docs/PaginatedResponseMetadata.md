# PaginatedResponseMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrevCursor** | Pointer to **string** | Pointer to the previous page of results. Use this value as a cursor query parameter in a subsequent request, along with limit, to navigate through the collection by virtual page. | [optional] 
**NextCursor** | Pointer to **string** | Pointer to the next page of results. Use this value as a cursor query parameter in a subsequent request, along with limit, to navigate through the collection by virtual page. | [optional] 
**Total** | Pointer to **int32** | The total number of results. This value may not be provided. | [optional] 

## Methods

### NewPaginatedResponseMetadata

`func NewPaginatedResponseMetadata() *PaginatedResponseMetadata`

NewPaginatedResponseMetadata instantiates a new PaginatedResponseMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedResponseMetadataWithDefaults

`func NewPaginatedResponseMetadataWithDefaults() *PaginatedResponseMetadata`

NewPaginatedResponseMetadataWithDefaults instantiates a new PaginatedResponseMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrevCursor

`func (o *PaginatedResponseMetadata) GetPrevCursor() string`

GetPrevCursor returns the PrevCursor field if non-nil, zero value otherwise.

### GetPrevCursorOk

`func (o *PaginatedResponseMetadata) GetPrevCursorOk() (*string, bool)`

GetPrevCursorOk returns a tuple with the PrevCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevCursor

`func (o *PaginatedResponseMetadata) SetPrevCursor(v string)`

SetPrevCursor sets PrevCursor field to given value.

### HasPrevCursor

`func (o *PaginatedResponseMetadata) HasPrevCursor() bool`

HasPrevCursor returns a boolean if a field has been set.

### GetNextCursor

`func (o *PaginatedResponseMetadata) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *PaginatedResponseMetadata) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *PaginatedResponseMetadata) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *PaginatedResponseMetadata) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetTotal

`func (o *PaginatedResponseMetadata) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PaginatedResponseMetadata) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PaginatedResponseMetadata) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PaginatedResponseMetadata) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


