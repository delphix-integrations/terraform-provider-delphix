# ApiClient

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The entity ID of this API client. | [optional] [default to null]
**ApiClientId** | **string** | The unique ID which is used to identity the identity of an API request. The web server (nginx) configuration must be configured so as to include the external ID as the value of the X_CLIENT_ID HTTP request header when requests are proxied. For OAuth2/JWT based authentication, this typically corresponds to a value extracted from the JWT, uniquely identifying the API client. | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**IsAdmin** | **bool** |  | [default to null]
**EngineUsersMapping** | [**[]EngineUserMapping**](EngineUserMapping.md) | Mapping of engine ID to the engine User ID. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

