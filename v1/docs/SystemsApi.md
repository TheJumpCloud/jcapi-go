# \SystemsApi

All URIs are relative to *https://console.jumpcloud.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETESystemsId**](SystemsApi.md#DELETESystemsId) | **Delete** /systems/{id} | Delete a System
[**GETSystems**](SystemsApi.md#GETSystems) | **Get** /systems | List All Systems
[**GETSystemsId**](SystemsApi.md#GETSystemsId) | **Get** /systems/{id} | List an individual system
[**GETSystemsSystemIDSystemusers**](SystemsApi.md#GETSystemsSystemIDSystemusers) | **Get** /systems/{systemID}/systemusers | List system user bindings
[**PUTSystemsId**](SystemsApi.md#PUTSystemsId) | **Put** /systems/{id} | Update a system
[**PUTSystemsSystemIDSystemusers**](SystemsApi.md#PUTSystemsSystemIDSystemusers) | **Put** /systems/{systemID}/systemusers | Update a systems or user&#39;s binding


# **DELETESystemsId**
> DELETESystemsId($id, $contentType, $accept)

Delete a System

Delete a system record by its id. This command will cause the system to uninstall the JumpCloud agent from its self which can can take about a minute. If the system is not connected to JumpCloud the system record will simply be removed.


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

# **GETSystems**
> InlineResponse2005 GETSystems($contentType, $accept, $fields, $limit, $skip, $sort)

List All Systems

Returns all Systems.


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

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETSystemsId**
> InlineResponse2006 GETSystemsId($id, $contentType, $accept, $fields, $limit, $skip, $sort)

List an individual system

Returns an individual system.


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

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETSystemsSystemIDSystemusers**
> GETSystemsSystemIDSystemusers($systemID, $contentType, $accept, $fields, $limit, $skip, $sort)

List system user bindings

List system user bindings for a specific system in a system and user binding format.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemID** | **string**|  | 
 **contentType** | **string**|  | [optional] [default to application/json]
 **accept** | **string**|  | [optional] [default to application/json]
 **fields** | **string**| The comma separated fileds included in the returned records. If omitted the default list of fields will be returned.  | [optional] [default to ]
 **limit** | **int32**| The number of records to return at once. | [optional] [default to 10]
 **skip** | **int32**| The offset into the records to return. | [optional] [default to 0]
 **sort** | **string**| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [optional] [default to ]

### Return type

void (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PUTSystemsId**
> PUTSystemsId($id, $body, $contentType, $accept)

Update a system

Update a system record by its id and return the modified system record in single record format.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **body** | [**Body**](Body.md)|  | [optional] 
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

# **PUTSystemsSystemIDSystemusers**
> PUTSystemsSystemIDSystemusers($systemID, $body, $contentType, $accept)

Update a systems or user's binding

Adds or removes a user binding for a system.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemID** | **string**|  | 
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

