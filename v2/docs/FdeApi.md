# \FdeApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemsGetFDEKey**](FdeApi.md#SystemsGetFDEKey) | **Get** /systems/{system_id}/fdekey | Get System FDE Key


# **SystemsGetFDEKey**
> Systemfdekey SystemsGetFDEKey(ctx, systemId, optional)
Get System FDE Key

This endpoint will return the current (latest) fde key saved for a system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **systemId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemId** | **string**|  | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**Systemfdekey**](systemfdekey.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

