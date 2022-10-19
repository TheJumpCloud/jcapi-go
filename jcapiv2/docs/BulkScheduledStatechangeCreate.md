# BulkScheduledStatechangeCreate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationEmailOverride** | **string** | Send the activation or welcome email to the specified email address upon activation. Can only be used with a single user_id and scheduled activation. This field will be ignored if &#x60;send_activation_emails&#x60; is explicitly set to false. | [optional] [default to null]
**SendActivationEmails** | **bool** | Set to true to send activation or welcome email(s) to each user_id upon activation. Set to false to suppress emails. Can only be used with scheduled activation(s). | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) | Date and time that scheduled action should occur | [default to null]
**State** | **string** | The state to move the user(s) to | [default to null]
**UserIds** | **[]string** | Array of system user ids to schedule for a state change | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

