# ListEnginesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Engine**](Engine.md) |  | [optional] 
**Errors** | Pointer to [**[]Error**](Error.md) | Sadly, sometimes requests to the API are not successful. Failures can occur for a wide range of reasons. The Errors object contains information about full or partial failures which might have occurred during the request. | [optional] 

## Methods

### NewListEnginesResponse

`func NewListEnginesResponse() *ListEnginesResponse`

NewListEnginesResponse instantiates a new ListEnginesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEnginesResponseWithDefaults

`func NewListEnginesResponseWithDefaults() *ListEnginesResponse`

NewListEnginesResponseWithDefaults instantiates a new ListEnginesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListEnginesResponse) GetItems() []Engine`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListEnginesResponse) GetItemsOk() (*[]Engine, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListEnginesResponse) SetItems(v []Engine)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListEnginesResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetErrors

`func (o *ListEnginesResponse) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ListEnginesResponse) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ListEnginesResponse) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ListEnginesResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


