# Environment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Environment object entity ID. | [optional] [default to null]
**Name** | **string** | The name of this environment. | [optional] [default to null]
**Namespace** | **string** | The namespace of this environment for replicated and restored objects. | [optional] [default to null]
**EngineId** | **int64** | A reference to the Engine that this Environment connection is associated with. | [optional] [default to null]
**Enabled** | **bool** | True if this environment is enabled. | [optional] [default to null]
**IsCluster** | **bool** | True if this environment is a cluster of hosts. | [optional] [default to null]
**Hosts** | [**[]Host**](Host.md) | The hosts that are part of this environment. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

