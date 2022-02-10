# RegisteredEngineUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] [readonly] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**HashicorpVaultUsernameCommandArgs** | Pointer to **[]string** | Arguments to pass to the Vault CLI tool to retrieve the username for the engine. | [optional] 
**HashicorpVaultPasswordCommandArgs** | Pointer to **[]string** | Arguments to pass to the Vault CLI tool to retrieve the password for the engine. | [optional] 
**HashicorpVaultId** | Pointer to **NullableInt64** | Reference to the Hashicorp vault to use to retrieve engine credentials. | [optional] 

## Methods

### NewRegisteredEngineUser

`func NewRegisteredEngineUser() *RegisteredEngineUser`

NewRegisteredEngineUser instantiates a new RegisteredEngineUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisteredEngineUserWithDefaults

`func NewRegisteredEngineUserWithDefaults() *RegisteredEngineUser`

NewRegisteredEngineUserWithDefaults instantiates a new RegisteredEngineUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegisteredEngineUser) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegisteredEngineUser) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegisteredEngineUser) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RegisteredEngineUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *RegisteredEngineUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RegisteredEngineUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RegisteredEngineUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RegisteredEngineUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *RegisteredEngineUser) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *RegisteredEngineUser) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *RegisteredEngineUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RegisteredEngineUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RegisteredEngineUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RegisteredEngineUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *RegisteredEngineUser) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *RegisteredEngineUser) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetHashicorpVaultUsernameCommandArgs

`func (o *RegisteredEngineUser) GetHashicorpVaultUsernameCommandArgs() []string`

GetHashicorpVaultUsernameCommandArgs returns the HashicorpVaultUsernameCommandArgs field if non-nil, zero value otherwise.

### GetHashicorpVaultUsernameCommandArgsOk

`func (o *RegisteredEngineUser) GetHashicorpVaultUsernameCommandArgsOk() (*[]string, bool)`

GetHashicorpVaultUsernameCommandArgsOk returns a tuple with the HashicorpVaultUsernameCommandArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultUsernameCommandArgs

`func (o *RegisteredEngineUser) SetHashicorpVaultUsernameCommandArgs(v []string)`

SetHashicorpVaultUsernameCommandArgs sets HashicorpVaultUsernameCommandArgs field to given value.

### HasHashicorpVaultUsernameCommandArgs

`func (o *RegisteredEngineUser) HasHashicorpVaultUsernameCommandArgs() bool`

HasHashicorpVaultUsernameCommandArgs returns a boolean if a field has been set.

### SetHashicorpVaultUsernameCommandArgsNil

`func (o *RegisteredEngineUser) SetHashicorpVaultUsernameCommandArgsNil(b bool)`

 SetHashicorpVaultUsernameCommandArgsNil sets the value for HashicorpVaultUsernameCommandArgs to be an explicit nil

### UnsetHashicorpVaultUsernameCommandArgs
`func (o *RegisteredEngineUser) UnsetHashicorpVaultUsernameCommandArgs()`

UnsetHashicorpVaultUsernameCommandArgs ensures that no value is present for HashicorpVaultUsernameCommandArgs, not even an explicit nil
### GetHashicorpVaultPasswordCommandArgs

`func (o *RegisteredEngineUser) GetHashicorpVaultPasswordCommandArgs() []string`

GetHashicorpVaultPasswordCommandArgs returns the HashicorpVaultPasswordCommandArgs field if non-nil, zero value otherwise.

### GetHashicorpVaultPasswordCommandArgsOk

`func (o *RegisteredEngineUser) GetHashicorpVaultPasswordCommandArgsOk() (*[]string, bool)`

GetHashicorpVaultPasswordCommandArgsOk returns a tuple with the HashicorpVaultPasswordCommandArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultPasswordCommandArgs

`func (o *RegisteredEngineUser) SetHashicorpVaultPasswordCommandArgs(v []string)`

SetHashicorpVaultPasswordCommandArgs sets HashicorpVaultPasswordCommandArgs field to given value.

### HasHashicorpVaultPasswordCommandArgs

`func (o *RegisteredEngineUser) HasHashicorpVaultPasswordCommandArgs() bool`

HasHashicorpVaultPasswordCommandArgs returns a boolean if a field has been set.

### SetHashicorpVaultPasswordCommandArgsNil

`func (o *RegisteredEngineUser) SetHashicorpVaultPasswordCommandArgsNil(b bool)`

 SetHashicorpVaultPasswordCommandArgsNil sets the value for HashicorpVaultPasswordCommandArgs to be an explicit nil

### UnsetHashicorpVaultPasswordCommandArgs
`func (o *RegisteredEngineUser) UnsetHashicorpVaultPasswordCommandArgs()`

UnsetHashicorpVaultPasswordCommandArgs ensures that no value is present for HashicorpVaultPasswordCommandArgs, not even an explicit nil
### GetHashicorpVaultId

`func (o *RegisteredEngineUser) GetHashicorpVaultId() int64`

GetHashicorpVaultId returns the HashicorpVaultId field if non-nil, zero value otherwise.

### GetHashicorpVaultIdOk

`func (o *RegisteredEngineUser) GetHashicorpVaultIdOk() (*int64, bool)`

GetHashicorpVaultIdOk returns a tuple with the HashicorpVaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashicorpVaultId

`func (o *RegisteredEngineUser) SetHashicorpVaultId(v int64)`

SetHashicorpVaultId sets HashicorpVaultId field to given value.

### HasHashicorpVaultId

`func (o *RegisteredEngineUser) HasHashicorpVaultId() bool`

HasHashicorpVaultId returns a boolean if a field has been set.

### SetHashicorpVaultIdNil

`func (o *RegisteredEngineUser) SetHashicorpVaultIdNil(b bool)`

 SetHashicorpVaultIdNil sets the value for HashicorpVaultId to be an explicit nil

### UnsetHashicorpVaultId
`func (o *RegisteredEngineUser) UnsetHashicorpVaultId()`

UnsetHashicorpVaultId ensures that no value is present for HashicorpVaultId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


