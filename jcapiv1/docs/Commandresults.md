# Commandresults

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** | The command that was executed on the system. | [optional] [default to null]
**System** | **string** | The ID of the system on which the command was executed. | [optional] [default to null]
**Organization** | **string** | The ID of the organization. | [optional] [default to null]
**User** | **string** | The user which ran the command. | [optional] [default to null]
**Files** | **[]string** | An array of file IDs that were included with the command. | [optional] [default to null]
**RequestTime** | [**time.Time**](time.Time.md) | The time that the command was sent. | [optional] [default to null]
**ResponseTime** | [**time.Time**](time.Time.md) | The time that the command was completed. | [optional] [default to null]
**ResponseId** | **string** | The ID of the parent command. | [optional] [default to null]
**ResponseDataOutput** | **int32** | the stdout from the command that ran. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


