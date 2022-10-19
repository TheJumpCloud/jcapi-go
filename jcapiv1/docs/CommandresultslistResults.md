# CommandresultslistResults

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the command result. | [optional] [default to null]
**Command** | **string** | The command that was executed on the system. | [optional] [default to null]
**ExitCode** | **int32** | The stderr output from the command that ran. | [optional] [default to null]
**Name** | **string** | The name of the command. | [optional] [default to null]
**RequestTime** | [**time.Time**](time.Time.md) | The time (UTC) that the command was sent. | [optional] [default to null]
**ResponseTime** | [**time.Time**](time.Time.md) | The time (UTC) that the command was completed. | [optional] [default to null]
**Sudo** | **bool** | If the user had sudo rights. | [optional] [default to null]
**System** | **string** | The display name of the system the command was executed on. | [optional] [default to null]
**SystemId** | **string** | The id of the system the command was executed on. | [optional] [default to null]
**User** | **string** | The user the command ran as. | [optional] [default to null]
**WorkflowId** | **string** | The id for the command that ran on the system. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

