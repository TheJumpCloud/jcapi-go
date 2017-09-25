# \SystemusersApi

All URIs are relative to *https://console.jumpcloud.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETESystemusersId**](SystemusersApi.md#DELETESystemusersId) | **Delete** /systemusers/{id} | Delete a system user
[**GETSystemusers**](SystemusersApi.md#GETSystemusers) | **Get** /systemusers | List all system users
[**GETSystemusersId**](SystemusersApi.md#GETSystemusersId) | **Get** /systemusers/{id} | List a system user
[**GETSystemusersSystemUserIDSystems**](SystemusersApi.md#GETSystemusersSystemUserIDSystems) | **Get** /systemusers/{systemUserID}/systems | List system user binding
[**POSTSystemusers**](SystemusersApi.md#POSTSystemusers) | **Post** /systemusers | Create a system user
[**PUTSystemusersId**](SystemusersApi.md#PUTSystemusersId) | **Put** /systemusers/{id} | Update a system user
[**PUTSystemusersSystemUserIDSystems**](SystemusersApi.md#PUTSystemusersSystemUserIDSystems) | **Put** /systemusers/{systemUserID}/systems | Update a system user binding


# **DELETESystemusersId**
> DELETESystemusersId($id, $contentType, $accept)

Delete a system user

Delete a particular system user.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
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

# **GETSystemusers**
> InlineResponse2007 GETSystemusers($limit, $skip, $sort, $fields, $filter, $contentType, $accept)

List all system users

Returns all systemusers.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**| The number of records to return at once. | [optional] [default to 10]
 **skip** | **int32**| The offset into the records to return. | [optional] [default to 0]
 **sort** | **string**| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [optional] [default to ]
 **fields** | **string**| The comma separated fileds included in the returned records. If omitted the default list of fields will be returned.  | [optional] [default to ]
 **filter** | **string**|  | [optional] [default to ]
 **contentType** | **string**|  | [optional] [default to application/json]
 **accept** | **string**|  | [optional] [default to application/json]

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETSystemusersId**
> InlineResponse2008 GETSystemusersId($id, $contentType, $accept, $fields, $limit, $skip, $sort)

List a system user

Get a particular System User.


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

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETSystemusersSystemUserIDSystems**
> InlineResponse2009 GETSystemusersSystemUserIDSystems($systemUserID, $contentType, $accept, $fields, $limit, $skip, $sort)

List system user binding

List system bindings for a specific system user in a system and user binding format.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemUserID** | **string**|  | 
 **contentType** | **string**|  | [optional] [default to application/json]
 **accept** | **string**|  | [optional] [default to application/json]
 **fields** | **string**| The comma separated fileds included in the returned records. If omitted the default list of fields will be returned.  | [optional] [default to ]
 **limit** | **int32**| The number of records to return at once. | [optional] [default to 10]
 **skip** | **int32**| The offset into the records to return. | [optional] [default to 0]
 **sort** | **string**| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [optional] [default to ]

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTSystemusers**
> SystemUser POSTSystemusers($body, $contentType, $accept)

Create a system user

Add new System Users.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SystemUser**](SystemUser.md)|  | [optional] 
 **contentType** | **string**|  | [optional] [default to application/json]
 **accept** | **string**|  | [optional] [default to application/json]

### Return type

[**SystemUser**](SystemUser.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PUTSystemusersId**
> PUTSystemusersId($id, $body, $contentType, $accept)

Update a system user

Update a system user record and return the modified record.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **body** | [**Body1**](Body1.md)|  | [optional] 
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

# **PUTSystemusersSystemUserIDSystems**
> PUTSystemusersSystemUserIDSystems($systemUserID, $body, $contentType, $accept)

Update a system user binding

Adds or removes a system binding for a user.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemUserID** | **string**|  | 
 **body** | [**Systemuserbindings**](Systemuserbindings.md)|  | [optional] 
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

