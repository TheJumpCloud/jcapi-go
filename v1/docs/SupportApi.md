# \SupportApi

All URIs are relative to *https://console.jumpcloud.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CasePost**](SupportApi.md#CasePost) | **Post** /api/cases | Create a Case


# **CasePost**
> InlineResponse200 CasePost(ctx, contentType, accept, optional)
Create a Case

This endpoint allows you to open a support case.  #### Sample Request  ``` curl -X POST https://console.jumpcloud.com/api/cases \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{  \"subject\":\"{subject}\",  \"description\":\"{description}\",  \"firstname\":\"{firstname}\",  \"lastname\":\"{lastname}\" }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **body** | [**ModelCase**](ModelCase.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

