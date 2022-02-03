# RegisteredEngine

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**Hostname** | **string** |  | [default to null]
**PrimaryUser** | **int64** | Id of the primary user for this engine. The primary user must be an engine admin. | [optional] [default to null]
**InsecureSsl** | **bool** | Allow connections to the engine over HTTPs without validating the TLS certificate. Even though the connection to the engine might be performed over HTTPs, setting this property eliminates the protection against a man-in-the-middle attach for connections to this engine. Instead, consider creating a truststore with a Certificate Authority to validate the engine&#x27;s certificate, and set the truststore_path propery.  | [optional] [default to false]
**UnsafeSslHostnameCheck** | **bool** | Ignore validation of the name associated to the TLS certificate when connecting to the engine over HTTPs. Setting this value must only be done if the TLS certificate of the engine does not match the hostname, and the TLS configuration of the engine cannot be fixed. Setting this property reduces the protection against a man-in-the-middle attack for connections to this engine. This is ignored if insecure_ssl is set.  | [optional] [default to false]
**TruststoreFilename** | **string** | File name of a truststore which can be used to validate the TLS certificate of the engine. The truststore must be available at /etc/config/certs/&lt;truststore_filename&gt;  | [optional] [default to null]
**TruststorePassword** | **string** | Password to read the truststore.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

