# Snapshot

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Snapshot ID. | [optional] [default to null]
**Timestamp** | [**time.Time**](time.Time.md) | The logical time of the data contained in this Snapshot. | [optional] [default to null]
**Location** | **string** | Database specific identifier for the data contained in this Snapshot, such as the Log Sequence Number (LSN) for MSsql databases, System Change Number (SCN) for Oracle databases. | [optional] [default to null]
**DatasetId** | **string** | The ID of the Snapshot&#x27;s dSource or VDB. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

