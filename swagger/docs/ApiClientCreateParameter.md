# ApiClientCreateParameter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | An optional logical name for the API client. | [optional] [default to null]
**GenerateApiKey** | **bool** | Whether an API key must be generated for this API client. This must be set if the API client will be used for API key based authentication, and unset otherwise. | [optional] [default to true]
**ApiClientId** | **string** | The unique ID which is used to identity the identity of an API request. The web server (nginx) configuration must be configured so as to include the external ID as the value of the X_CLIENT_ID HTTP request header when requests are proxied. If this value isn&#x27;t set, the application will automatically generate one. For OAuth2/JWT based authentication, this typically corresponds to a value extracted from the JWT, uniquely identifying the API client. | [optional] [default to null]
**IsAdmin** | **bool** |  | [optional] [default to false]
**EngineUsersMapping** | [**[]EngineUserMapping**](EngineUserMapping.md) | Mapping of engine ID to the engine User ID. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

