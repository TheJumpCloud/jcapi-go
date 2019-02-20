# \CommandResultsApi

All URIs are relative to *https://console.jumpcloud.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommandResultsDelete**](CommandResultsApi.md#CommandResultsDelete) | **Delete** /commandresults/{id} | Delete a Command result
[**CommandResultsGet**](CommandResultsApi.md#CommandResultsGet) | **Get** /commandresults/{id} | List an individual Command result
[**CommandResultsList**](CommandResultsApi.md#CommandResultsList) | **Get** /commandresults | List all Command Results


# **CommandResultsDelete**
> Commandresult CommandResultsDelete(ctx, id, contentType, accept, optional)
Delete a Command result

This endpoint deletes a specific command result.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/commandresults/{CommandID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'   ````

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **xOrgId** | **string**|  | [default to ]

### Return type

[**Commandresult**](commandresult.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommandResultsGet**
> Commandresult CommandResultsGet(ctx, id, contentType, accept, optional)
List an individual Command result

This endpoint returns a specific command result.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/commandresults/{CommandResultID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **fields** | **string**| Use a space seperated string of field parameters to include the data in the response. If omitted the default list of fields will be returned.  | [default to ]
 **filter** | **string**| A filter to apply to the query. | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**Commandresult**](commandresult.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommandResultsList**
> Commandresultslist CommandResultsList(ctx, contentType, accept, optional)
List all Command Results

This endpoint returns all command results.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/commandresults \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key:{API_KEY}'   ```

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
 **fields** | **string**| Use a space seperated string of field parameters to include the data in the response. If omitted the default list of fields will be returned.  | [default to ]
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | **string**| Use space separated sort parameters to sort the collection. Default sort is ascending. Prefix with &#x60;-&#x60; to sort descending.  | [default to ]
 **filter** | **string**| A filter to apply to the query. | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**Commandresultslist**](commandresultslist.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

