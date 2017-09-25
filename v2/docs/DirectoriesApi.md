# \DirectoriesApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoriesList**](DirectoriesApi.md#DirectoriesList) | **Get** /directories | List All Directories


# **DirectoriesList**
> []Directory DirectoriesList($contentType, $accept, $fields, $limit, $skip, $sort)

List All Directories

This endpoint returns all active directories (LDAP, O365 Suite, G-Suite).  #### Sample Request ``` https://console.jumpcloud.com/api/v2/directories ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **fields** | **string**| The comma separated fields included in the returned records. If omitted the default list of fields will be returned.  | [optional] [default to ]
 **limit** | **int32**| The number of records to return at once. | [optional] [default to 10]
 **skip** | **int32**| The offset into the records to return. | [optional] [default to 0]
 **sort** | **string**| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [optional] [default to ]

### Return type

[**[]Directory**](Directory.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

