# Commands

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [optional] [default to null]
**Command** | **string** | The command to execute on the server. | [default to null]
**User** | **string** | The ID of the JC managed user to run the command as. | [default to null]
**Systems** | **[]string** | An array of system IDs to run the command on. | [optional] [default to null]
**Schedule** | **string** | A crontab that consists of: [ (seconds) (minutes) (hours) (days of month) (months) (weekdays) ] or [ immediate ]. If you send this as an empty string, it will run immediately.  | [optional] [default to null]
**Files** | **[]string** | An array of file IDs to include with the command. | [optional] [default to null]
**Tags** | **[]string** | An array of tag IDs to run the command on. | [optional] [default to null]
**Timeout** | **string** | The time in seconds to allow the command to run for. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


