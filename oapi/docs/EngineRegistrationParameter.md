# EngineRegistrationParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Hostname** | **string** |  | 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**HashicorpVaultUsernameCommandArgs** | Pointer to **[]string** | Arguments to pass to the Vault CLI tool to retrieve the username for the engine. | [optional] 
**HashicorpVaultPasswordCommandArgs** | Pointer to **[]string** | Arguments to pass to the Vault CLI tool to retrieve the password for the engine. | [optional] 
**HashicorpVaultId** | Pointer to **NullableInt64** | Reference to the Hashicorp vault to use to retrieve engine credentials. | [optional] 
**InsecureSsl** | Pointer to **bool** | Allow connections to the engine over HTTPs without validating the TLS certificate. Even though the connection to the engine might be performed over HTTPs, setting this property eliminates the protection against a man-in-the-middle attach for connections to this engine. Instead, consider creating a truststore with a Certificate Authority to validate the engine&#39;s certificate, and set the truststore_path propery.  | [optional] [default to false]
**UnsafeSslHostnameCheck** | Pointer to **bool** | Ignore validation of the name associated to the TLS certificate when connecting to the engine over HTTPs. Setting this value must only be done if the TLS certificate of the engine does not match the hostname, and the TLS configuration of the engine cannot be fixed. Setting this property reduces the protection against a man-in-the-middle attack for connections to this engine. This is ignored if insecure_ssl is set.  | [optional] [default to false]
**TruststoreFilename** | Pointer to **NullableString** | File name of a truststore which can be used to validate the TLS certificate of the engine. The truststore must be available at /etc/config/certs/&lt;truststore_filename&gt;  | [optional] 
**TruststorePassword** | Pointer to **NullableString** | Password to read the truststore.  | [optional] 

## Methods

### NewEngineRegistrationParameter

`func NewEngineRegistrationParameter(name string, hostname string, ) *EngineRegistrationParameter`

NewEngineRegistrationParameter instantiates a new EngineRegistrationParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineRegistrationParameterWithDefaults

`func NewEngineRegistrationParameterWithDefaults() *EngineRegistrationParameter`

NewEngineRegistrationParameterWithDefaults instantiates a new EngineRegistrationParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EngineRegistrationParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EngineRegistrationParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EngineRegistrationParameter) SetName(v string)`

SetName sets Name field to given value.


### GetHostname

`func (o *EngineRegistrationParameter) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *EngineRegistrationParameter) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *EngineRegistrationParameter) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetUsername

`func (o *EngineRegistrationParameter) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EngineRegistrationParameter) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EngineRegistrationParameter) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *EngineRegistrationParameter) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *EngineRegistrationParameter) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *EngineRegistrationParameter) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *EngineRegistrationParameter) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EngineRegistrationParameter) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EngineRegistrationParameter) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EngineRegistrationParameter) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *EngineRegistrationParameter) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *EngineRegistrationParameter) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetHashicorpVaultUsernameCommandArgs

`func (o *EngineRegistrationParameter) GetHashicorpVaultUsernameCommandArgs() []string`

GetHashicorpVaultUsernameCommandArgs returns the HashicorpVaultUsernameCommandArgs field if non-nil, zero value otherwise.

### GetHashicorpVaultUsernameCommandArgsOk

`func (o *EngineRegistrationParameter) GetHashicorpVaultUsernameCommandArgsOk() (*[]string, bool)`

GetHashicorpVaultUsernameCommandArgsOk returns a tuple with the HashicorpVaultUsernameCommandArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultUsernameCommandArgs

`func (o *EngineRegistrationParameter) SetHashicorpVaultUsernameCommandArgs(v []string)`

SetHashicorpVaultUsernameCommandArgs sets HashicorpVaultUsernameCommandArgs field to given value.

### HasHashicorpVaultUsernameCommandArgs

`func (o *EngineRegistrationParameter) HasHashicorpVaultUsernameCommandArgs() bool`

HasHashicorpVaultUsernameCommandArgs returns a boolean if a field has been set.

### SetHashicorpVaultUsernameCommandArgsNil

`func (o *EngineRegistrationParameter) SetHashicorpVaultUsernameCommandArgsNil(b bool)`

 SetHashicorpVaultUsernameCommandArgsNil sets the value for HashicorpVaultUsernameCommandArgs to be an explicit nil

### UnsetHashicorpVaultUsernameCommandArgs
`func (o *EngineRegistrationParameter) UnsetHashicorpVaultUsernameCommandArgs()`

UnsetHashicorpVaultUsernameCommandArgs ensures that no value is present for HashicorpVaultUsernameCommandArgs, not even an explicit nil
### GetHashicorpVaultPasswordCommandArgs

`func (o *EngineRegistrationParameter) GetHashicorpVaultPasswordCommandArgs() []string`

GetHashicorpVaultPasswordCommandArgs returns the HashicorpVaultPasswordCommandArgs field if non-nil, zero value otherwise.

### GetHashicorpVaultPasswordCommandArgsOk

`func (o *EngineRegistrationParameter) GetHashicorpVaultPasswordCommandArgsOk() (*[]string, bool)`

GetHashicorpVaultPasswordCommandArgsOk returns a tuple with the HashicorpVaultPasswordCommandArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultPasswordCommandArgs

`func (o *EngineRegistrationParameter) SetHashicorpVaultPasswordCommandArgs(v []string)`

SetHashicorpVaultPasswordCommandArgs sets HashicorpVaultPasswordCommandArgs field to given value.

### HasHashicorpVaultPasswordCommandArgs

`func (o *EngineRegistrationParameter) HasHashicorpVaultPasswordCommandArgs() bool`

HasHashicorpVaultPasswordCommandArgs returns a boolean if a field has been set.

### SetHashicorpVaultPasswordCommandArgsNil

`func (o *EngineRegistrationParameter) SetHashicorpVaultPasswordCommandArgsNil(b bool)`

 SetHashicorpVaultPasswordCommandArgsNil sets the value for HashicorpVaultPasswordCommandArgs to be an explicit nil

### UnsetHashicorpVaultPasswordCommandArgs
`func (o *EngineRegistrationParameter) UnsetHashicorpVaultPasswordCommandArgs()`

UnsetHashicorpVaultPasswordCommandArgs ensures that no value is present for HashicorpVaultPasswordCommandArgs, not even an explicit nil
### GetHashicorpVaultId

`func (o *EngineRegistrationParameter) GetHashicorpVaultId() int64`

GetHashicorpVaultId returns the HashicorpVaultId field if non-nil, zero value otherwise.

### GetHashicorpVaultIdOk

`func (o *EngineRegistrationParameter) GetHashicorpVaultIdOk() (*int64, bool)`

GetHashicorpVaultIdOk returns a tuple with the HashicorpVaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultId

`func (o *EngineRegistrationParameter) SetHashicorpVaultId(v int64)`

SetHashicorpVaultId sets HashicorpVaultId field to given value.

### HasHashicorpVaultId

`func (o *EngineRegistrationParameter) HasHashicorpVaultId() bool`

HasHashicorpVaultId returns a boolean if a field has been set.

### SetHashicorpVaultIdNil

`func (o *EngineRegistrationParameter) SetHashicorpVaultIdNil(b bool)`

 SetHashicorpVaultIdNil sets the value for HashicorpVaultId to be an explicit nil

### UnsetHashicorpVaultId
`func (o *EngineRegistrationParameter) UnsetHashicorpVaultId()`

UnsetHashicorpVaultId ensures that no value is present for HashicorpVaultId, not even an explicit nil
### GetInsecureSsl

`func (o *EngineRegistrationParameter) GetInsecureSsl() bool`

GetInsecureSsl returns the InsecureSsl field if non-nil, zero value otherwise.

### GetInsecureSslOk

`func (o *EngineRegistrationParameter) GetInsecureSslOk() (*bool, bool)`

GetInsecureSslOk returns a tuple with the InsecureSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSsl

`func (o *EngineRegistrationParameter) SetInsecureSsl(v bool)`

SetInsecureSsl sets InsecureSsl field to given value.

### HasInsecureSsl

`func (o *EngineRegistrationParameter) HasInsecureSsl() bool`

HasInsecureSsl returns a boolean if a field has been set.

### GetUnsafeSslHostnameCheck

`func (o *EngineRegistrationParameter) GetUnsafeSslHostnameCheck() bool`

GetUnsafeSslHostnameCheck returns the UnsafeSslHostnameCheck field if non-nil, zero value otherwise.

### GetUnsafeSslHostnameCheckOk

`func (o *EngineRegistrationParameter) GetUnsafeSslHostnameCheckOk() (*bool, bool)`

GetUnsafeSslHostnameCheckOk returns a tuple with the UnsafeSslHostnameCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsafeSslHostnameCheck

`func (o *EngineRegistrationParameter) SetUnsafeSslHostnameCheck(v bool)`

SetUnsafeSslHostnameCheck sets UnsafeSslHostnameCheck field to given value.

### HasUnsafeSslHostnameCheck

`func (o *EngineRegistrationParameter) HasUnsafeSslHostnameCheck() bool`

HasUnsafeSslHostnameCheck returns a boolean if a field has been set.

### GetTruststoreFilename

`func (o *EngineRegistrationParameter) GetTruststoreFilename() string`

GetTruststoreFilename returns the TruststoreFilename field if non-nil, zero value otherwise.

### GetTruststoreFilenameOk

`func (o *EngineRegistrationParameter) GetTruststoreFilenameOk() (*string, bool)`

GetTruststoreFilenameOk returns a tuple with the TruststoreFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruststoreFilename

`func (o *EngineRegistrationParameter) SetTruststoreFilename(v string)`

SetTruststoreFilename sets TruststoreFilename field to given value.

### HasTruststoreFilename

`func (o *EngineRegistrationParameter) HasTruststoreFilename() bool`

HasTruststoreFilename returns a boolean if a field has been set.

### SetTruststoreFilenameNil

`func (o *EngineRegistrationParameter) SetTruststoreFilenameNil(b bool)`

 SetTruststoreFilenameNil sets the value for TruststoreFilename to be an explicit nil

### UnsetTruststoreFilename
`func (o *EngineRegistrationParameter) UnsetTruststoreFilename()`

UnsetTruststoreFilename ensures that no value is present for TruststoreFilename, not even an explicit nil
### GetTruststorePassword

`func (o *EngineRegistrationParameter) GetTruststorePassword() string`

GetTruststorePassword returns the TruststorePassword field if non-nil, zero value otherwise.

### GetTruststorePasswordOk

`func (o *EngineRegistrationParameter) GetTruststorePasswordOk() (*string, bool)`

GetTruststorePasswordOk returns a tuple with the TruststorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruststorePassword

`func (o *EngineRegistrationParameter) SetTruststorePassword(v string)`

SetTruststorePassword sets TruststorePassword field to given value.

### HasTruststorePassword

`func (o *EngineRegistrationParameter) HasTruststorePassword() bool`

HasTruststorePassword returns a boolean if a field has been set.

### SetTruststorePasswordNil

`func (o *EngineRegistrationParameter) SetTruststorePasswordNil(b bool)`

 SetTruststorePasswordNil sets the value for TruststorePassword to be an explicit nil

### UnsetTruststorePassword
`func (o *EngineRegistrationParameter) UnsetTruststorePassword()`

UnsetTruststorePassword ensures that no value is present for TruststorePassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


