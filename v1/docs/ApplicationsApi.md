# \ApplicationsApi

All URIs are relative to *https://console.jumpcloud.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETApplications**](ApplicationsApi.md#GETApplications) | **Get** /applications | Applications


# **GETApplications**
> InlineResponse200 GETApplications($fields, $limit, $skip, $sort)

Applications

The endpoint is used to return all your SSO Applications.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string**| The comma separated fileds included in the returned records. If omitted the default list of fields will be returned. | [optional] 
 **limit** | **int32**| The number of records to return at once. | [optional] 
 **skip** | **int32**| The offset into the records to return. | [optional] 
 **sort** | **string**|  | [optional] [default to The comma separated fields used to sort the collection. Default sort is ascending, prefix with - to sort descending.]

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

