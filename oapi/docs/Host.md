# Host

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** | The hostname or IP address of this host. | [optional] 
**OsName** | Pointer to **string** | The name of the OS on this host. | [optional] 
**OsVersion** | Pointer to **string** | The version of the OS on this host. | [optional] 
**MemorySize** | Pointer to **int64** | The total amount of memory on this host in bytes. | [optional] 

## Methods

### NewHost

`func NewHost() *Host`

NewHost instantiates a new Host object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostWithDefaults

`func NewHostWithDefaults() *Host`

NewHostWithDefaults instantiates a new Host object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *Host) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Host) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Host) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *Host) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetOsName

`func (o *Host) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *Host) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *Host) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *Host) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### GetOsVersion

`func (o *Host) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *Host) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *Host) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *Host) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetMemorySize

`func (o *Host) GetMemorySize() int64`

GetMemorySize returns the MemorySize field if non-nil, zero value otherwise.

### GetMemorySizeOk

`func (o *Host) GetMemorySizeOk() (*int64, bool)`

GetMemorySizeOk returns a tuple with the MemorySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySize

`func (o *Host) SetMemorySize(v int64)`

SetMemorySize sets MemorySize field to given value.

### HasMemorySize

`func (o *Host) HasMemorySize() bool`

HasMemorySize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


