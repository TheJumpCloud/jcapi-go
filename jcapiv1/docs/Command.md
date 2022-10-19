# Command

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** | The command to execute on the server. | [default to null]
**CommandRunners** | **[]string** | An array of IDs of the Command Runner Users that can execute this command. | [optional] [default to null]
**CommandType** | **string** | The Command OS | [default to linux]
**Files** | **[]string** | An array of file IDs to include with the command. | [optional] [default to null]
**LaunchType** | **string** | How the command will execute. | [optional] [default to null]
**ListensTo** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**Organization** | **string** | The ID of the organization. | [optional] [default to null]
**Schedule** | **string** | A crontab that consists of: [ (seconds) (minutes) (hours) (days of month) (months) (weekdays) ] or [ immediate ]. If you send this as an empty string, it will run immediately.  | [optional] [default to null]
**ScheduleRepeatType** | **string** | When the command will repeat. | [optional] [default to null]
**ScheduleYear** | **int32** | The year that a scheduled command will launch in. | [optional] [default to null]
**Shell** | **string** | The shell used to run the command. | [optional] [default to null]
**Sudo** | **bool** |  | [optional] [default to null]
**Systems** | **[]string** | An array of system IDs to run the command on. Not available if you are using Groups. | [optional] [default to null]
**Template** | **string** | The template this command was created from | [optional] [default to null]
**TimeToLiveSeconds** | **int32** | Time in seconds a command can wait in the queue to be run before timing out | [optional] [default to null]
**Timeout** | **string** | The time in seconds to allow the command to run for. | [optional] [default to null]
**Trigger** | **string** | The name of the command trigger. | [optional] [default to null]
**User** | **string** | The ID of the system user to run the command as. This field is required when creating a command with a commandType of \&quot;mac\&quot; or \&quot;linux\&quot;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

