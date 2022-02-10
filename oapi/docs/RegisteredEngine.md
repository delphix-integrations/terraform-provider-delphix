# RegisteredEngine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Hostname** | **string** |  | 
**PrimaryUser** | Pointer to **int64** | Id of the primary user for this engine. The primary user must be an engine admin. | [optional] 
**InsecureSsl** | Pointer to **bool** | Allow connections to the engine over HTTPs without validating the TLS certificate. Even though the connection to the engine might be performed over HTTPs, setting this property eliminates the protection against a man-in-the-middle attach for connections to this engine. Instead, consider creating a truststore with a Certificate Authority to validate the engine&#39;s certificate, and set the truststore_path propery.  | [optional] [default to false]
**UnsafeSslHostnameCheck** | Pointer to **bool** | Ignore validation of the name associated to the TLS certificate when connecting to the engine over HTTPs. Setting this value must only be done if the TLS certificate of the engine does not match the hostname, and the TLS configuration of the engine cannot be fixed. Setting this property reduces the protection against a man-in-the-middle attack for connections to this engine. This is ignored if insecure_ssl is set.  | [optional] [default to false]
**TruststoreFilename** | Pointer to **NullableString** | File name of a truststore which can be used to validate the TLS certificate of the engine. The truststore must be available at /etc/config/certs/&lt;truststore_filename&gt;  | [optional] 
**TruststorePassword** | Pointer to **NullableString** | Password to read the truststore.  | [optional] 
**Status** | Pointer to **NullableString** | the status of the engine  | [optional] [readonly] 

## Methods

### NewRegisteredEngine

`func NewRegisteredEngine(name string, hostname string, ) *RegisteredEngine`

NewRegisteredEngine instantiates a new RegisteredEngine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisteredEngineWithDefaults

`func NewRegisteredEngineWithDefaults() *RegisteredEngine`

NewRegisteredEngineWithDefaults instantiates a new RegisteredEngine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegisteredEngine) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegisteredEngine) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegisteredEngine) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RegisteredEngine) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RegisteredEngine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegisteredEngine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegisteredEngine) SetName(v string)`

SetName sets Name field to given value.


### GetHostname

`func (o *RegisteredEngine) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *RegisteredEngine) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *RegisteredEngine) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetPrimaryUser

`func (o *RegisteredEngine) GetPrimaryUser() int64`

GetPrimaryUser returns the PrimaryUser field if non-nil, zero value otherwise.

### GetPrimaryUserOk

`func (o *RegisteredEngine) GetPrimaryUserOk() (*int64, bool)`

GetPrimaryUserOk returns a tuple with the PrimaryUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryUser

`func (o *RegisteredEngine) SetPrimaryUser(v int64)`

SetPrimaryUser sets PrimaryUser field to given value.

### HasPrimaryUser

`func (o *RegisteredEngine) HasPrimaryUser() bool`

HasPrimaryUser returns a boolean if a field has been set.

### GetInsecureSsl

`func (o *RegisteredEngine) GetInsecureSsl() bool`

GetInsecureSsl returns the InsecureSsl field if non-nil, zero value otherwise.

### GetInsecureSslOk

`func (o *RegisteredEngine) GetInsecureSslOk() (*bool, bool)`

GetInsecureSslOk returns a tuple with the InsecureSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSsl

`func (o *RegisteredEngine) SetInsecureSsl(v bool)`

SetInsecureSsl sets InsecureSsl field to given value.

### HasInsecureSsl

`func (o *RegisteredEngine) HasInsecureSsl() bool`

HasInsecureSsl returns a boolean if a field has been set.

### GetUnsafeSslHostnameCheck

`func (o *RegisteredEngine) GetUnsafeSslHostnameCheck() bool`

GetUnsafeSslHostnameCheck returns the UnsafeSslHostnameCheck field if non-nil, zero value otherwise.

### GetUnsafeSslHostnameCheckOk

`func (o *RegisteredEngine) GetUnsafeSslHostnameCheckOk() (*bool, bool)`

GetUnsafeSslHostnameCheckOk returns a tuple with the UnsafeSslHostnameCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsafeSslHostnameCheck

`func (o *RegisteredEngine) SetUnsafeSslHostnameCheck(v bool)`

SetUnsafeSslHostnameCheck sets UnsafeSslHostnameCheck field to given value.

### HasUnsafeSslHostnameCheck

`func (o *RegisteredEngine) HasUnsafeSslHostnameCheck() bool`

HasUnsafeSslHostnameCheck returns a boolean if a field has been set.

### GetTruststoreFilename

`func (o *RegisteredEngine) GetTruststoreFilename() string`

GetTruststoreFilename returns the TruststoreFilename field if non-nil, zero value otherwise.

### GetTruststoreFilenameOk

`func (o *RegisteredEngine) GetTruststoreFilenameOk() (*string, bool)`

GetTruststoreFilenameOk returns a tuple with the TruststoreFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruststoreFilename

`func (o *RegisteredEngine) SetTruststoreFilename(v string)`

SetTruststoreFilename sets TruststoreFilename field to given value.

### HasTruststoreFilename

`func (o *RegisteredEngine) HasTruststoreFilename() bool`

HasTruststoreFilename returns a boolean if a field has been set.

### SetTruststoreFilenameNil

`func (o *RegisteredEngine) SetTruststoreFilenameNil(b bool)`

 SetTruststoreFilenameNil sets the value for TruststoreFilename to be an explicit nil

### UnsetTruststoreFilename
`func (o *RegisteredEngine) UnsetTruststoreFilename()`

UnsetTruststoreFilename ensures that no value is present for TruststoreFilename, not even an explicit nil
### GetTruststorePassword

`func (o *RegisteredEngine) GetTruststorePassword() string`

GetTruststorePassword returns the TruststorePassword field if non-nil, zero value otherwise.

### GetTruststorePasswordOk

`func (o *RegisteredEngine) GetTruststorePasswordOk() (*string, bool)`

GetTruststorePasswordOk returns a tuple with the TruststorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruststorePassword

`func (o *RegisteredEngine) SetTruststorePassword(v string)`

SetTruststorePassword sets TruststorePassword field to given value.

### HasTruststorePassword

`func (o *RegisteredEngine) HasTruststorePassword() bool`

HasTruststorePassword returns a boolean if a field has been set.

### SetTruststorePasswordNil

`func (o *RegisteredEngine) SetTruststorePasswordNil(b bool)`

 SetTruststorePasswordNil sets the value for TruststorePassword to be an explicit nil

### UnsetTruststorePassword
`func (o *RegisteredEngine) UnsetTruststorePassword()`

UnsetTruststorePassword ensures that no value is present for TruststorePassword, not even an explicit nil
### GetStatus

`func (o *RegisteredEngine) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegisteredEngine) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegisteredEngine) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RegisteredEngine) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *RegisteredEngine) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *RegisteredEngine) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


