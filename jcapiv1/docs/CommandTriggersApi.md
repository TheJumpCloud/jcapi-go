# \CommandTriggersApi

All URIs are relative to *https://console.jumpcloud.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**POSTCommandTriggerWebhook**](CommandTriggersApi.md#POSTCommandTriggerWebhook) | **Post** /command/trigger/{triggername} | Run a Command assigned to a webhook


# **POSTCommandTriggerWebhook**
> POSTCommandTriggerWebhook($triggername, $contentType, $accept)

Run a Command assigned to a webhook


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **triggername** | **string**|  | 
 **contentType** | **string**|  | [optional] [default to application/json]
 **accept** | **string**|  | [optional] [default to application/json]

### Return type

void (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

