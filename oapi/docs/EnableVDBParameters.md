# EnableVDBParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptStart** | Pointer to **bool** | Whether to attempt a startup of the VDB after the enable. | [optional] [default to true]

## Methods

### NewEnableVDBParameters

`func NewEnableVDBParameters() *EnableVDBParameters`

NewEnableVDBParameters instantiates a new EnableVDBParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableVDBParametersWithDefaults

`func NewEnableVDBParametersWithDefaults() *EnableVDBParameters`

NewEnableVDBParametersWithDefaults instantiates a new EnableVDBParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptStart

`func (o *EnableVDBParameters) GetAttemptStart() bool`

GetAttemptStart returns the AttemptStart field if non-nil, zero value otherwise.

### GetAttemptStartOk

`func (o *EnableVDBParameters) GetAttemptStartOk() (*bool, bool)`

GetAttemptStartOk returns a tuple with the AttemptStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptStart

`func (o *EnableVDBParameters) SetAttemptStart(v bool)`

SetAttemptStart sets AttemptStart field to given value.

### HasAttemptStart

`func (o *EnableVDBParameters) HasAttemptStart() bool`

HasAttemptStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


