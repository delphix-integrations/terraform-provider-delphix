# Vdb

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The VDB object entity ID. | [optional] [default to null]
**DatabaseType** | **string** | The database type of this VDB. | [optional] [default to null]
**Name** | **string** | The container name of this VDB. | [optional] [default to null]
**DatabaseVersion** | **string** | The database version of this VDB. | [optional] [default to null]
**Size** | **int64** | The total size of this VDB, in bytes. | [optional] [default to null]
**EngineId** | **string** | A reference to the Engine that this VDB belongs to. | [optional] [default to null]
**Status** | **string** | The runtime status of the VDB. &#x27;Unknown&#x27; if all attempts to connect to the dataset failed. | [optional] [default to null]
**EnvironmentId** | **string** | A reference to the Environment that hosts this VDB. | [optional] [default to null]
**IpAddress** | **string** | The IP address of the VDB&#x27;s host. | [optional] [default to null]
**Fqdn** | **string** | The FQDN of the VDB&#x27;s host. | [optional] [default to null]
**ParentId** | **string** | A reference to the parent dataset of this VDB. | [optional] [default to null]
**GroupName** | **string** | The name of the group containing this VDB. | [optional] [default to null]
**CreationDate** | [**time.Time**](time.Time.md) | The date this VDB was created. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

