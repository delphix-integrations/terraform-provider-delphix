# DisableVDBParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptCleanup** | Pointer to **bool** | Whether to attempt a cleanup of the VDB before the disable. | [optional] [default to true]

## Methods

### NewDisableVDBParameters

`func NewDisableVDBParameters() *DisableVDBParameters`

NewDisableVDBParameters instantiates a new DisableVDBParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisableVDBParametersWithDefaults

`func NewDisableVDBParametersWithDefaults() *DisableVDBParameters`

NewDisableVDBParametersWithDefaults instantiates a new DisableVDBParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptCleanup

`func (o *DisableVDBParameters) GetAttemptCleanup() bool`

GetAttemptCleanup returns the AttemptCleanup field if non-nil, zero value otherwise.

### GetAttemptCleanupOk

`func (o *DisableVDBParameters) GetAttemptCleanupOk() (*bool, bool)`

GetAttemptCleanupOk returns a tuple with the AttemptCleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptCleanup

`func (o *DisableVDBParameters) SetAttemptCleanup(v bool)`

SetAttemptCleanup sets AttemptCleanup field to given value.

### HasAttemptCleanup

`func (o *DisableVDBParameters) HasAttemptCleanup() bool`

HasAttemptCleanup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


