# VDB

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The VDB object entity ID. | [optional] 
**DatabaseType** | Pointer to **NullableString** | The database type of this VDB. | [optional] 
**Name** | Pointer to **NullableString** | The container name of this VDB. | [optional] 
**DatabaseVersion** | Pointer to **NullableString** | The database version of this VDB. | [optional] 
**Size** | Pointer to **NullableInt64** | The total size of this VDB, in bytes. | [optional] 
**EngineId** | Pointer to **string** | A reference to the Engine that this VDB belongs to. | [optional] 
**Status** | Pointer to **NullableString** | The runtime status of the VDB. &#39;Unknown&#39; if all attempts to connect to the dataset failed. | [optional] 
**EnvironmentId** | Pointer to **NullableString** | A reference to the Environment that hosts this VDB. | [optional] 
**IpAddress** | Pointer to **NullableString** | The IP address of the VDB&#39;s host. | [optional] 
**Fqdn** | Pointer to **NullableString** | The FQDN of the VDB&#39;s host. | [optional] 
**ParentId** | Pointer to **NullableString** | A reference to the parent dataset of this VDB. | [optional] 
**GroupName** | Pointer to **NullableString** | The name of the group containing this VDB. | [optional] 
**CreationDate** | Pointer to **NullableTime** | The date this VDB was created. | [optional] 

## Methods

### NewVDB

`func NewVDB() *VDB`

NewVDB instantiates a new VDB object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVDBWithDefaults

`func NewVDBWithDefaults() *VDB`

NewVDBWithDefaults instantiates a new VDB object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VDB) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VDB) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VDB) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VDB) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDatabaseType

`func (o *VDB) GetDatabaseType() string`

GetDatabaseType returns the DatabaseType field if non-nil, zero value otherwise.

### GetDatabaseTypeOk

`func (o *VDB) GetDatabaseTypeOk() (*string, bool)`

GetDatabaseTypeOk returns a tuple with the DatabaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseType

`func (o *VDB) SetDatabaseType(v string)`

SetDatabaseType sets DatabaseType field to given value.

### HasDatabaseType

`func (o *VDB) HasDatabaseType() bool`

HasDatabaseType returns a boolean if a field has been set.

### SetDatabaseTypeNil

`func (o *VDB) SetDatabaseTypeNil(b bool)`

 SetDatabaseTypeNil sets the value for DatabaseType to be an explicit nil

### UnsetDatabaseType
`func (o *VDB) UnsetDatabaseType()`

UnsetDatabaseType ensures that no value is present for DatabaseType, not even an explicit nil
### GetName

`func (o *VDB) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VDB) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VDB) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VDB) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VDB) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VDB) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDatabaseVersion

`func (o *VDB) GetDatabaseVersion() string`

GetDatabaseVersion returns the DatabaseVersion field if non-nil, zero value otherwise.

### GetDatabaseVersionOk

`func (o *VDB) GetDatabaseVersionOk() (*string, bool)`

GetDatabaseVersionOk returns a tuple with the DatabaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseVersion

`func (o *VDB) SetDatabaseVersion(v string)`

SetDatabaseVersion sets DatabaseVersion field to given value.

### HasDatabaseVersion

`func (o *VDB) HasDatabaseVersion() bool`

HasDatabaseVersion returns a boolean if a field has been set.

### SetDatabaseVersionNil

`func (o *VDB) SetDatabaseVersionNil(b bool)`

 SetDatabaseVersionNil sets the value for DatabaseVersion to be an explicit nil

### UnsetDatabaseVersion
`func (o *VDB) UnsetDatabaseVersion()`

UnsetDatabaseVersion ensures that no value is present for DatabaseVersion, not even an explicit nil
### GetSize

`func (o *VDB) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VDB) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VDB) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *VDB) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *VDB) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *VDB) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetEngineId

`func (o *VDB) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *VDB) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *VDB) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *VDB) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetStatus

`func (o *VDB) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VDB) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VDB) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VDB) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *VDB) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *VDB) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetEnvironmentId

`func (o *VDB) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *VDB) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *VDB) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *VDB) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### SetEnvironmentIdNil

`func (o *VDB) SetEnvironmentIdNil(b bool)`

 SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil

### UnsetEnvironmentId
`func (o *VDB) UnsetEnvironmentId()`

UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
### GetIpAddress

`func (o *VDB) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VDB) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VDB) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *VDB) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *VDB) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *VDB) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetFqdn

`func (o *VDB) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *VDB) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *VDB) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *VDB) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### SetFqdnNil

`func (o *VDB) SetFqdnNil(b bool)`

 SetFqdnNil sets the value for Fqdn to be an explicit nil

### UnsetFqdn
`func (o *VDB) UnsetFqdn()`

UnsetFqdn ensures that no value is present for Fqdn, not even an explicit nil
### GetParentId

`func (o *VDB) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *VDB) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *VDB) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *VDB) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *VDB) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *VDB) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetGroupName

`func (o *VDB) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *VDB) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *VDB) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *VDB) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *VDB) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *VDB) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetCreationDate

`func (o *VDB) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *VDB) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *VDB) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *VDB) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### SetCreationDateNil

`func (o *VDB) SetCreationDateNil(b bool)`

 SetCreationDateNil sets the value for CreationDate to be an explicit nil

### UnsetCreationDate
`func (o *VDB) UnsetCreationDate()`

UnsetCreationDate ensures that no value is present for CreationDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


