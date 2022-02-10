# HashicorpVault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] [readonly] 
**EnvVariables** | Pointer to **map[string]string** | Environment variables to set when invoking the Vault CLI tool. The environment variables will be used both to login to the vault (if this step is required) and to retrieve engine username and passwords.  | [optional] 
**LoginCommandArgs** | Pointer to **[]string** | Arguments to the \&quot;vault\&quot; CLI tool to be used to fetch a client token (or \&quot;login\&quot;). If supporting files, such as TLS certificates, must be used to authenticate, they can be mounted to the /etc/config directory. This property must not be set when using the TOKEN authentication method as login is not required.  | [optional] 

## Methods

### NewHashicorpVault

`func NewHashicorpVault() *HashicorpVault`

NewHashicorpVault instantiates a new HashicorpVault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHashicorpVaultWithDefaults

`func NewHashicorpVaultWithDefaults() *HashicorpVault`

NewHashicorpVaultWithDefaults instantiates a new HashicorpVault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HashicorpVault) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HashicorpVault) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HashicorpVault) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *HashicorpVault) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvVariables

`func (o *HashicorpVault) GetEnvVariables() map[string]string`

GetEnvVariables returns the EnvVariables field if non-nil, zero value otherwise.

### GetEnvVariablesOk

`func (o *HashicorpVault) GetEnvVariablesOk() (*map[string]string, bool)`

GetEnvVariablesOk returns a tuple with the EnvVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVariables

`func (o *HashicorpVault) SetEnvVariables(v map[string]string)`

SetEnvVariables sets EnvVariables field to given value.

### HasEnvVariables

`func (o *HashicorpVault) HasEnvVariables() bool`

HasEnvVariables returns a boolean if a field has been set.

### GetLoginCommandArgs

`func (o *HashicorpVault) GetLoginCommandArgs() []string`

GetLoginCommandArgs returns the LoginCommandArgs field if non-nil, zero value otherwise.

### GetLoginCommandArgsOk

`func (o *HashicorpVault) GetLoginCommandArgsOk() (*[]string, bool)`

GetLoginCommandArgsOk returns a tuple with the LoginCommandArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCommandArgs

`func (o *HashicorpVault) SetLoginCommandArgs(v []string)`

SetLoginCommandArgs sets LoginCommandArgs field to given value.

### HasLoginCommandArgs

`func (o *HashicorpVault) HasLoginCommandArgs() bool`

HasLoginCommandArgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


