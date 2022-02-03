# EngineRegistrationParameter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [default to null]
**Hostname** | **string** |  | [default to null]
**Username** | **string** |  | [optional] [default to null]
**Password** | **string** |  | [optional] [default to null]
**HashicorpVaultUsernameCommandArgs** | **[]string** | Arguments to pass to the Vault CLI tool to retrieve the username for the engine. | [optional] [default to null]
**HashicorpVaultPasswordCommandArgs** | **[]string** | Arguments to pass to the Vault CLI tool to retrieve the password for the engine. | [optional] [default to null]
**HashicorpVaultId** | **int64** | Reference to the Hashicorp vault to use to retrieve engine credentials. | [optional] [default to null]
**InsecureSsl** | **bool** | Allow connections to the engine over HTTPs without validating the TLS certificate. Even though the connection to the engine might be performed over HTTPs, setting this property eliminates the protection against a man-in-the-middle attach for connections to this engine. Instead, consider creating a truststore with a Certificate Authority to validate the engine&#x27;s certificate, and set the truststore_path propery.  | [optional] [default to false]
**UnsafeSslHostnameCheck** | **bool** | Ignore validation of the name associated to the TLS certificate when connecting to the engine over HTTPs. Setting this value must only be done if the TLS certificate of the engine does not match the hostname, and the TLS configuration of the engine cannot be fixed. Setting this property reduces the protection against a man-in-the-middle attack for connections to this engine. This is ignored if insecure_ssl is set.  | [optional] [default to false]
**TruststoreFilename** | **string** | File name of a truststore which can be used to validate the TLS certificate of the engine. The truststore must be available at /etc/config/certs/&lt;truststore_filename&gt;  | [optional] [default to null]
**TruststorePassword** | **string** | Password to read the truststore.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

