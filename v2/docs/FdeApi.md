# \FdeApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemsGetFDEKey**](FdeApi.md#SystemsGetFDEKey) | **Get** /systems/{system_id}/fdekey | Get System FDE Key


# **SystemsGetFDEKey**
> InlineResponse200 SystemsGetFDEKey(ctx, systemId)
Get System FDE Key

This endpoint will return the current (latest) fde key saved for a system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **systemId** | **string**|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

