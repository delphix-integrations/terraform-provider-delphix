# ProvisionVDBResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **string** | The ID of the provisioning job. | [optional] 
**Vdb** | Pointer to [**VDB**](VDB.md) |  | [optional] 

## Methods

### NewProvisionVDBResponse

`func NewProvisionVDBResponse() *ProvisionVDBResponse`

NewProvisionVDBResponse instantiates a new ProvisionVDBResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionVDBResponseWithDefaults

`func NewProvisionVDBResponseWithDefaults() *ProvisionVDBResponse`

NewProvisionVDBResponseWithDefaults instantiates a new ProvisionVDBResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *ProvisionVDBResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ProvisionVDBResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ProvisionVDBResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *ProvisionVDBResponse) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetVdb

`func (o *ProvisionVDBResponse) GetVdb() VDB`

GetVdb returns the Vdb field if non-nil, zero value otherwise.

### GetVdbOk

`func (o *ProvisionVDBResponse) GetVdbOk() (*VDB, bool)`

GetVdbOk returns a tuple with the Vdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdb

`func (o *ProvisionVDBResponse) SetVdb(v VDB)`

SetVdb sets Vdb field to given value.

### HasVdb

`func (o *ProvisionVDBResponse) HasVdb() bool`

HasVdb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


