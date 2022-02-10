# ApiClientCreateParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | An optional logical name for the API client. | [optional] 
**GenerateApiKey** | Pointer to **bool** | Whether an API key must be generated for this API client. This must be set if the API client will be used for API key based authentication, and unset otherwise. | [optional] [default to true]
**ApiClientId** | Pointer to **string** | The unique ID which is used to identity the identity of an API request. The web server (nginx) configuration must be configured so as to include the external ID as the value of the X_CLIENT_ID HTTP request header when requests are proxied. If this value isn&#39;t set, the application will automatically generate one. For OAuth2/JWT based authentication, this typically corresponds to a value extracted from the JWT, uniquely identifying the API client. | [optional] 
**IsAdmin** | Pointer to **bool** |  | [optional] [default to false]
**EngineUsersMapping** | Pointer to [**[]EngineUserMapping**](EngineUserMapping.md) | Mapping of engine ID to the engine User ID. | [optional] 

## Methods

### NewApiClientCreateParameter

`func NewApiClientCreateParameter() *ApiClientCreateParameter`

NewApiClientCreateParameter instantiates a new ApiClientCreateParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiClientCreateParameterWithDefaults

`func NewApiClientCreateParameterWithDefaults() *ApiClientCreateParameter`

NewApiClientCreateParameterWithDefaults instantiates a new ApiClientCreateParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiClientCreateParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiClientCreateParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiClientCreateParameter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiClientCreateParameter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGenerateApiKey

`func (o *ApiClientCreateParameter) GetGenerateApiKey() bool`

GetGenerateApiKey returns the GenerateApiKey field if non-nil, zero value otherwise.

### GetGenerateApiKeyOk

`func (o *ApiClientCreateParameter) GetGenerateApiKeyOk() (*bool, bool)`

GetGenerateApiKeyOk returns a tuple with the GenerateApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateApiKey

`func (o *ApiClientCreateParameter) SetGenerateApiKey(v bool)`

SetGenerateApiKey sets GenerateApiKey field to given value.

### HasGenerateApiKey

`func (o *ApiClientCreateParameter) HasGenerateApiKey() bool`

HasGenerateApiKey returns a boolean if a field has been set.

### GetApiClientId

`func (o *ApiClientCreateParameter) GetApiClientId() string`

GetApiClientId returns the ApiClientId field if non-nil, zero value otherwise.

### GetApiClientIdOk

`func (o *ApiClientCreateParameter) GetApiClientIdOk() (*string, bool)`

GetApiClientIdOk returns a tuple with the ApiClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiClientId

`func (o *ApiClientCreateParameter) SetApiClientId(v string)`

SetApiClientId sets ApiClientId field to given value.

### HasApiClientId

`func (o *ApiClientCreateParameter) HasApiClientId() bool`

HasApiClientId returns a boolean if a field has been set.

### GetIsAdmin

`func (o *ApiClientCreateParameter) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *ApiClientCreateParameter) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *ApiClientCreateParameter) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *ApiClientCreateParameter) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetEngineUsersMapping

`func (o *ApiClientCreateParameter) GetEngineUsersMapping() []EngineUserMapping`

GetEngineUsersMapping returns the EngineUsersMapping field if non-nil, zero value otherwise.

### GetEngineUsersMappingOk

`func (o *ApiClientCreateParameter) GetEngineUsersMappingOk() (*[]EngineUserMapping, bool)`

GetEngineUsersMappingOk returns a tuple with the EngineUsersMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineUsersMapping

`func (o *ApiClientCreateParameter) SetEngineUsersMapping(v []EngineUserMapping)`

SetEngineUsersMapping sets EngineUsersMapping field to given value.

### HasEngineUsersMapping

`func (o *ApiClientCreateParameter) HasEngineUsersMapping() bool`

HasEngineUsersMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


