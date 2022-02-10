# Source

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Source object entity ID. | [optional] 
**DatabaseType** | Pointer to **NullableString** | The type of this source database. | [optional] 
**Name** | Pointer to **NullableString** | The name of this source database. | [optional] 
**DatabaseVersion** | Pointer to **NullableString** | The version of this source database. | [optional] 
**EnvironmentId** | Pointer to **NullableString** | A reference to the Environment that hosts this source database. | [optional] 
**DataUuid** | Pointer to **NullableString** | A universal ID that uniquely identifies this source database. | [optional] 
**IpAddress** | Pointer to **NullableString** | The IP address of the source&#39;s host. | [optional] 
**Fqdn** | Pointer to **NullableString** | The FQDN of the source&#39;s host. | [optional] 
**Size** | Pointer to **NullableInt64** | The total size of this source database, in bytes. | [optional] 
**JdbcConnectionString** | Pointer to **NullableString** | The JDBC connection URL for this source database. | [optional] 
**PluginVersion** | Pointer to **NullableString** | The version of the plugin associated with this source database. | [optional] 

## Methods

### NewSource

`func NewSource() *Source`

NewSource instantiates a new Source object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceWithDefaults

`func NewSourceWithDefaults() *Source`

NewSourceWithDefaults instantiates a new Source object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Source) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Source) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Source) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Source) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDatabaseType

`func (o *Source) GetDatabaseType() string`

GetDatabaseType returns the DatabaseType field if non-nil, zero value otherwise.

### GetDatabaseTypeOk

`func (o *Source) GetDatabaseTypeOk() (*string, bool)`

GetDatabaseTypeOk returns a tuple with the DatabaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseType

`func (o *Source) SetDatabaseType(v string)`

SetDatabaseType sets DatabaseType field to given value.

### HasDatabaseType

`func (o *Source) HasDatabaseType() bool`

HasDatabaseType returns a boolean if a field has been set.

### SetDatabaseTypeNil

`func (o *Source) SetDatabaseTypeNil(b bool)`

 SetDatabaseTypeNil sets the value for DatabaseType to be an explicit nil

### UnsetDatabaseType
`func (o *Source) UnsetDatabaseType()`

UnsetDatabaseType ensures that no value is present for DatabaseType, not even an explicit nil
### GetName

`func (o *Source) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Source) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Source) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Source) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Source) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Source) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDatabaseVersion

`func (o *Source) GetDatabaseVersion() string`

GetDatabaseVersion returns the DatabaseVersion field if non-nil, zero value otherwise.

### GetDatabaseVersionOk

`func (o *Source) GetDatabaseVersionOk() (*string, bool)`

GetDatabaseVersionOk returns a tuple with the DatabaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseVersion

`func (o *Source) SetDatabaseVersion(v string)`

SetDatabaseVersion sets DatabaseVersion field to given value.

### HasDatabaseVersion

`func (o *Source) HasDatabaseVersion() bool`

HasDatabaseVersion returns a boolean if a field has been set.

### SetDatabaseVersionNil

`func (o *Source) SetDatabaseVersionNil(b bool)`

 SetDatabaseVersionNil sets the value for DatabaseVersion to be an explicit nil

### UnsetDatabaseVersion
`func (o *Source) UnsetDatabaseVersion()`

UnsetDatabaseVersion ensures that no value is present for DatabaseVersion, not even an explicit nil
### GetEnvironmentId

`func (o *Source) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *Source) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *Source) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *Source) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### SetEnvironmentIdNil

`func (o *Source) SetEnvironmentIdNil(b bool)`

 SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil

### UnsetEnvironmentId
`func (o *Source) UnsetEnvironmentId()`

UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
### GetDataUuid

`func (o *Source) GetDataUuid() string`

GetDataUuid returns the DataUuid field if non-nil, zero value otherwise.

### GetDataUuidOk

`func (o *Source) GetDataUuidOk() (*string, bool)`

GetDataUuidOk returns a tuple with the DataUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUuid

`func (o *Source) SetDataUuid(v string)`

SetDataUuid sets DataUuid field to given value.

### HasDataUuid

`func (o *Source) HasDataUuid() bool`

HasDataUuid returns a boolean if a field has been set.

### SetDataUuidNil

`func (o *Source) SetDataUuidNil(b bool)`

 SetDataUuidNil sets the value for DataUuid to be an explicit nil

### UnsetDataUuid
`func (o *Source) UnsetDataUuid()`

UnsetDataUuid ensures that no value is present for DataUuid, not even an explicit nil
### GetIpAddress

`func (o *Source) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Source) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Source) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *Source) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *Source) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *Source) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetFqdn

`func (o *Source) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *Source) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *Source) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *Source) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### SetFqdnNil

`func (o *Source) SetFqdnNil(b bool)`

 SetFqdnNil sets the value for Fqdn to be an explicit nil

### UnsetFqdn
`func (o *Source) UnsetFqdn()`

UnsetFqdn ensures that no value is present for Fqdn, not even an explicit nil
### GetSize

`func (o *Source) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Source) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Source) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *Source) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *Source) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *Source) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetJdbcConnectionString

`func (o *Source) GetJdbcConnectionString() string`

GetJdbcConnectionString returns the JdbcConnectionString field if non-nil, zero value otherwise.

### GetJdbcConnectionStringOk

`func (o *Source) GetJdbcConnectionStringOk() (*string, bool)`

GetJdbcConnectionStringOk returns a tuple with the JdbcConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJdbcConnectionString

`func (o *Source) SetJdbcConnectionString(v string)`

SetJdbcConnectionString sets JdbcConnectionString field to given value.

### HasJdbcConnectionString

`func (o *Source) HasJdbcConnectionString() bool`

HasJdbcConnectionString returns a boolean if a field has been set.

### SetJdbcConnectionStringNil

`func (o *Source) SetJdbcConnectionStringNil(b bool)`

 SetJdbcConnectionStringNil sets the value for JdbcConnectionString to be an explicit nil

### UnsetJdbcConnectionString
`func (o *Source) UnsetJdbcConnectionString()`

UnsetJdbcConnectionString ensures that no value is present for JdbcConnectionString, not even an explicit nil
### GetPluginVersion

`func (o *Source) GetPluginVersion() string`

GetPluginVersion returns the PluginVersion field if non-nil, zero value otherwise.

### GetPluginVersionOk

`func (o *Source) GetPluginVersionOk() (*string, bool)`

GetPluginVersionOk returns a tuple with the PluginVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginVersion

`func (o *Source) SetPluginVersion(v string)`

SetPluginVersion sets PluginVersion field to given value.

### HasPluginVersion

`func (o *Source) HasPluginVersion() bool`

HasPluginVersion returns a boolean if a field has been set.

### SetPluginVersionNil

`func (o *Source) SetPluginVersionNil(b bool)`

 SetPluginVersionNil sets the value for PluginVersion to be an explicit nil

### UnsetPluginVersion
`func (o *Source) UnsetPluginVersion()`

UnsetPluginVersion ensures that no value is present for PluginVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


