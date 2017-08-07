# \GroupsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsList**](GroupsApi.md#GroupsList) | **Get** /groups | List All Groups


# **GroupsList**
> []Group GroupsList($contentType, $accept, $fields, $filter, $limit, $skip, $sort)

List All Groups

This endpoint returns all Groups that exist in your organization.   #### Available filter fields:   - `name`   - `disabled`   - `type`   -     #### Sample Request  ```  https://console.jumpcloud.com/api/v2/groups ```


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

[**[]Group**](Group.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

