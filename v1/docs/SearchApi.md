# \SearchApi

All URIs are relative to *https://console.jumpcloud.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**POSTSearchSystemusers**](SearchApi.md#POSTSearchSystemusers) | **Post** /search/systemusers | List System Users


# **POSTSearchSystemusers**
> InlineResponse2004 POSTSearchSystemusers($body, $contentType, $accept, $limit, $skip, $fields)

List System Users

Return System Users in multi-record format allowing for the passing of the 'filter' parameter. The WILL NOT allow you to add a new system.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Search**](Search.md)|  | [optional] 
 **contentType** | **string**|  | [optional] [default to application/json]
 **accept** | **string**|  | [optional] [default to application/json]
 **limit** | **int32**| The number of records to return at once. | [optional] [default to 10]
 **skip** | **int32**| The offset into the records to return. | [optional] [default to 0]
 **fields** | **string**| The comma separated fileds included in the returned records. If omitted the default list of fields will be returned.  | [optional] [default to ]

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

