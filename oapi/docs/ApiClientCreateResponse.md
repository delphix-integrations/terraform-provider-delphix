# ApiClientCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | The opaque token to use to authenticate for other API calls. The token must be included in all HTTP API calls in a request header named \&quot;Authorization\&quot;, and prefixed with \&quot;apk \&quot; to denote that it is a proprietary API key format. For instance, if the token is \&quot;abc123\&quot;, request must contain the following HTTP request header: \&quot;Authorization: apk abc123\&quot;.  | [optional] 
**ApiClientEntityId** | Pointer to **int64** |  | [optional] 

## Methods

### NewApiClientCreateResponse

`func NewApiClientCreateResponse() *ApiClientCreateResponse`

NewApiClientCreateResponse instantiates a new ApiClientCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiClientCreateResponseWithDefaults

`func NewApiClientCreateResponseWithDefaults() *ApiClientCreateResponse`

NewApiClientCreateResponseWithDefaults instantiates a new ApiClientCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *ApiClientCreateResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ApiClientCreateResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ApiClientCreateResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ApiClientCreateResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetApiClientEntityId

`func (o *ApiClientCreateResponse) GetApiClientEntityId() int64`

GetApiClientEntityId returns the ApiClientEntityId field if non-nil, zero value otherwise.

### GetApiClientEntityIdOk

`func (o *ApiClientCreateResponse) GetApiClientEntityIdOk() (*int64, bool)`

GetApiClientEntityIdOk returns a tuple with the ApiClientEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiClientEntityId

`func (o *ApiClientCreateResponse) SetApiClientEntityId(v int64)`

SetApiClientEntityId sets ApiClientEntityId field to given value.

### HasApiClientEntityId

`func (o *ApiClientCreateResponse) HasApiClientEntityId() bool`

HasApiClientEntityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


