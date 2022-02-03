# RefreshVdbByTimestampParameters

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | [**time.Time**](time.Time.md) | The point in time from which to execute the operation. Mutually exclusive with timestamp_in_database_timezone. If the timestamp is not set, selects the latest point. | [optional] [default to null]
**TimestampInDatabaseTimezone** | **string** | The point in time from which to execute the operation, expressed as a date-time in the timezone of the source database. Mutually exclusive with timestamp. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

