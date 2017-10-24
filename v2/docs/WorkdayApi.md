# \WorkdayApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkdaysDelete**](WorkdayApi.md#WorkdaysDelete) | **Delete** /workdays/{id} | Delete Workday
[**WorkdaysGet**](WorkdayApi.md#WorkdaysGet) | **Get** /workdays/{id} | Get Workday
[**WorkdaysList**](WorkdayApi.md#WorkdaysList) | **Get** /workdays | List Workdays
[**WorkdaysPost**](WorkdayApi.md#WorkdaysPost) | **Post** /workdays | Create new Workday
[**WorkdaysPut**](WorkdayApi.md#WorkdaysPut) | **Put** /workdays/{id} | Update Workday
[**WorkdaysReport**](WorkdayApi.md#WorkdaysReport) | **Get** /workdays/{id}/report | Get Workday Report Results
[**WorkdaysSettings**](WorkdayApi.md#WorkdaysSettings) | **Get** /workdays/settings | Get Workday Settings


# **WorkdaysDelete**
> WorkdaysDelete($id, $contentType, $accept, $body)

Delete Workday

This endpoint allows you to delete a workday


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **body** | [**WorkdayRequest**](WorkdayRequest.md)|  | [optional] 

### Return type

void (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkdaysGet**
> WorkdayOutput WorkdaysGet($id, $contentType, $accept)

Get Workday


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]

### Return type

[**WorkdayOutput**](workday-output.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkdaysList**
> []WorkdayOutput WorkdaysList($contentType, $accept, $fields, $filter, $limit, $skip, $sort)

List Workdays


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **fields** | **string**| The comma separated fields included in the returned records. If omitted the default list of fields will be returned.  | [optional] [default to ]
 **filter** | **string**| Supported operators are: eq, ne, gt, ge, lt, le, between, search | [optional] [default to ]
 **limit** | **int32**| The number of records to return at once. | [optional] [default to 10]
 **skip** | **int32**| The offset into the records to return. | [optional] [default to 0]
 **sort** | **string**| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [optional] [default to ]

### Return type

[**[]WorkdayOutput**](workday-output.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkdaysPost**
> WorkdayOutput WorkdaysPost($contentType, $accept, $body)

Create new Workday

This endpoint allows you to create a new workday object


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **body** | [**Body**](Body.md)|  | [optional] 

### Return type

[**WorkdayOutput**](workday-output.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkdaysPut**
> WorkdayOutput WorkdaysPut($id, $contentType, $accept, $body)

Update Workday

This endpoint allows you to update the name and report_url for a Workday Authentication Edit


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **body** | [**WorkdayInput**](WorkdayInput.md)|  | [optional] 

### Return type

[**WorkdayOutput**](workday-output.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkdaysReport**
> WorkdayReportResult WorkdaysReport($id, $contentType, $accept, $fields, $filter, $sort, $skip)

Get Workday Report Results


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **fields** | **string**| The comma separated fields included in the returned records. If omitted the default list of fields will be returned.  | [optional] [default to ]
 **filter** | **string**| Supported operators are: eq, ne, gt, ge, lt, le, between, search | [optional] [default to ]
 **sort** | **string**| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [optional] [default to ]
 **skip** | **int32**| The offset into the records to return. | [optional] [default to 0]

### Return type

[**WorkdayReportResult**](workday-report-result.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkdaysSettings**
> InlineResponse200 WorkdaysSettings($contentType, $accept, $state)

Get Workday Settings

This endpoint allows you to obtain all settings needed for creating a workday instance, namely the URL to initiate an OAuth negotiation


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **state** | **string**|  | [optional] 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

