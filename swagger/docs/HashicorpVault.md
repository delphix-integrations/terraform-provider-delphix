# HashicorpVault

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [optional] [default to null]
**EnvVariables** | **map[string]string** | Environment variables to set when invoking the Vault CLI tool. The environment variables will be used both to login to the vault (if this step is required) and to retrieve engine username and passwords.  | [optional] [default to null]
**LoginCommandArgs** | **[]string** | Arguments to the \&quot;vault\&quot; CLI tool to be used to fetch a client token (or \&quot;login\&quot;). If supporting files, such as TLS certificates, must be used to authenticate, they can be mounted to the /etc/config directory. This property must not be set when using the TOKEN authentication method as login is not required.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

