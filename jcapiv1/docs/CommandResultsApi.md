# \CommandResultsApi

All URIs are relative to *https://console.jumpcloud.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETECommandresultsId**](CommandResultsApi.md#DELETECommandresultsId) | **Delete** /commandresults/{id} | Delete a Command result
[**GETCommandresults**](CommandResultsApi.md#GETCommandresults) | **Get** /commandresults | List all Command Results
[**GETCommandresultsId**](CommandResultsApi.md#GETCommandresultsId) | **Get** /commandresults/{id} | List an individual Command result


# **DELETECommandresultsId**
> InlineResponse2003 DELETECommandresultsId($id, $contentType, $accept)

Delete a Command result

Deletes a specific command result.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **contentType** | **string**|  | [optional] [default to application/json]
 **accept** | **string**|  | [optional] [default to application/json]

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETCommandresults**
> InlineResponse2002 GETCommandresults($contentType, $accept, $fields, $limit, $skip, $sort)

List all Command Results

Returns all command results.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string**|  | [optional] [default to application/json]
 **accept** | **string**|  | [optional] [default to application/json]
 **fields** | **string**| The comma separated fileds included in the returned records. If omitted the default list of fields will be returned.  | [optional] [default to ]
 **limit** | **int32**| The number of records to return at once. | [optional] [default to 10]
 **skip** | **int32**| The offset into the records to return. | [optional] [default to 0]
 **sort** | **string**| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [optional] [default to ]

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETCommandresultsId**
> Commandresults GETCommandresultsId($id, $contentType, $accept, $fields, $limit, $skip, $sort)

List an individual Command result

Returns a specific command result.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **contentType** | **string**|  | [optional] [default to application/json]
 **accept** | **string**|  | [optional] [default to application/json]
 **fields** | **string**| The comma separated fileds included in the returned records. If omitted the default list of fields will be returned.  | [optional] [default to ]
 **limit** | **int32**| The number of records to return at once. | [optional] [default to 10]
 **skip** | **int32**| The offset into the records to return. | [optional] [default to 0]
 **sort** | **string**| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [optional] [default to ]

### Return type

[**Commandresults**](commandresults.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

