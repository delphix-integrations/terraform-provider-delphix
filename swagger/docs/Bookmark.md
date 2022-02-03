# Bookmark

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Bookmark object entity ID. | [optional] [default to null]
**Name** | **string** | The user-defined name of this bookmark. | [default to null]
**CreationDate** | [**time.Time**](time.Time.md) | The date and time that this bookmark was created. | [optional] [default to null]
**VdbIds** | **[]string** | The list of VDB IDs associated with this bookmark. | [default to null]
**Retention** | **int64** | The retention policy for this bookmark, in days. A value of -1 indicates the bookmark should be kept forever. | [optional] [default to null]
**Status** | **string** | A message with details about operation progress or state of this bookmark. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

