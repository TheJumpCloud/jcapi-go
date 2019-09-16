# Command

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** | The command to execute on the server. | [default to null]
**CommandRunners** | **[]string** | An array of IDs of the Command Runner Users that can execute this command. | [optional] [default to null]
**CommandType** | **string** | The Command OS | [optional] [default to null]
**Files** | **[]string** | An array of file IDs to include with the command. | [optional] [default to null]
**LaunchType** | **string** | How the command will execute. | [optional] [default to null]
**ListensTo** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Organization** | **string** | The ID of the organization. | [optional] [default to null]
**Schedule** | **string** | A crontab that consists of: [ (seconds) (minutes) (hours) (days of month) (months) (weekdays) ] or [ immediate ]. If you send this as an empty string, it will run immediately.  | [optional] [default to null]
**ScheduleRepeatType** | **string** | When the command will repeat. | [optional] [default to null]
**Sudo** | **bool** |  | [optional] [default to null]
**Systems** | **[]string** | An array of system IDs to run the command on. Not available if you are using Groups. | [optional] [default to null]
**Timeout** | **string** | The time in seconds to allow the command to run for. | [optional] [default to null]
**User** | **string** | The ID of the system user to run the command as. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


