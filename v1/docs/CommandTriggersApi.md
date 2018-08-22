# \CommandTriggersApi

All URIs are relative to *https://console.jumpcloud.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommandTriggerWebhookPost**](CommandTriggersApi.md#CommandTriggerWebhookPost) | **Post** /command/trigger/{triggername} | Launch a command via a Trigger


# **CommandTriggerWebhookPost**
> CommandTriggerWebhookPost(ctx, triggername, contentType, accept)
Launch a command via a Trigger

This endpoint allows you to launch a command based on a defined trigger.  #### Sample Requests  **Launch a Command via a Trigger**  ``` curl --silent \\      -X 'POST' \\      -H \"x-api-key: {API_KEY}\" \\      \"https://console.jumpcloud.com/api/command/trigger/{TriggerName}\" ``` **Launch a Command via a Trigger passing a JSON object to the command** ``` curl --silent \\      -X 'POST' \\      -H \"x-api-key: {API_KEY}\" \\      -H 'Accept: application/json' \\      -H 'Content-Type: application/json' \\      -d '{ \"srcip\":\"192.168.2.32\", \"attack\":\"Cross Site Scripting Attempt\" }' \\      \"https://console.jumpcloud.com/api/command/trigger/{TriggerName}\" ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **triggername** | **string**|  | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

