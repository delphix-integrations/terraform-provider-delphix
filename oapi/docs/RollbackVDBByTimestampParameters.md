# RollbackVDBByTimestampParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | The point in time from which to execute the operation. Mutually exclusive with timestamp_in_database_timezone. If the timestamp is not set, selects the latest point. | [optional] 
**TimestampInDatabaseTimezone** | Pointer to **string** | The point in time from which to execute the operation, expressed as a date-time in the timezone of the source database. Mutually exclusive with timestamp. | [optional] 

## Methods

### NewRollbackVDBByTimestampParameters

`func NewRollbackVDBByTimestampParameters() *RollbackVDBByTimestampParameters`

NewRollbackVDBByTimestampParameters instantiates a new RollbackVDBByTimestampParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRollbackVDBByTimestampParametersWithDefaults

`func NewRollbackVDBByTimestampParametersWithDefaults() *RollbackVDBByTimestampParameters`

NewRollbackVDBByTimestampParametersWithDefaults instantiates a new RollbackVDBByTimestampParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *RollbackVDBByTimestampParameters) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RollbackVDBByTimestampParameters) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RollbackVDBByTimestampParameters) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *RollbackVDBByTimestampParameters) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTimestampInDatabaseTimezone

`func (o *RollbackVDBByTimestampParameters) GetTimestampInDatabaseTimezone() string`

GetTimestampInDatabaseTimezone returns the TimestampInDatabaseTimezone field if non-nil, zero value otherwise.

### GetTimestampInDatabaseTimezoneOk

`func (o *RollbackVDBByTimestampParameters) GetTimestampInDatabaseTimezoneOk() (*string, bool)`

GetTimestampInDatabaseTimezoneOk returns a tuple with the TimestampInDatabaseTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampInDatabaseTimezone

`func (o *RollbackVDBByTimestampParameters) SetTimestampInDatabaseTimezone(v string)`

SetTimestampInDatabaseTimezone sets TimestampInDatabaseTimezone field to given value.

### HasTimestampInDatabaseTimezone

`func (o *RollbackVDBByTimestampParameters) HasTimestampInDatabaseTimezone() bool`

HasTimestampInDatabaseTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


