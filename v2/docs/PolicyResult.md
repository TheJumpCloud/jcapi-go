# PolicyResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyID** | **string** | ObjectId uniquely identifying the parent Policy. | [optional] [default to null]
**SystemID** | **string** | ObjectId uniquely identifying the parent System. | [optional] [default to null]
**Id** | **string** | ObjectId uniquely identifying a Policy Result. | [optional] [default to null]
**StartedAt** | [**time.Time**](time.Time.md) | The start of the policy application. | [optional] [default to null]
**EndedAt** | [**time.Time**](time.Time.md) | The end of the policy application. | [optional] [default to null]
**Success** | **bool** | True if the policy was successfully applied; false otherwise. | [optional] [default to null]
**ExitStatus** | **int32** | The 32-bit unsigned exit status from the applying the policy. | [optional] [default to null]
**StdErr** | **string** | The STDERR output from applying the policy. | [optional] [default to null]
**StdOut** | **string** | The STDOUT output from applying the policy. | [optional] [default to null]
**State** | **string** | Enumeration describing the state of the policy. Success, failed, or pending. | [optional] [default to null]
**Detail** | **string** | Details pertaining to the policy result. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


