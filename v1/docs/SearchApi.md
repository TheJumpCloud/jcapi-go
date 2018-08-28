# \SearchApi

All URIs are relative to *https://console.jumpcloud.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchOrginzationsPost**](SearchApi.md#SearchOrginzationsPost) | **Post** /search/organizations | Search Organizations
[**SearchSystemsPost**](SearchApi.md#SearchSystemsPost) | **Post** /search/systems | Search Systems
[**SearchSystemusersPost**](SearchApi.md#SearchSystemusersPost) | **Post** /search/systemusers | Search System Users


# **SearchOrginzationsPost**
> Systemslist SearchOrginzationsPost(ctx, contentType, accept, optional)
Search Organizations

This endpoint will return Organization data based on your search parameters. This endpoint WILL NOT allow you to add a new Orginization.  You can use the supported parameters and pass those in the body of request.   The parameters must be passed as Content-Type application/json.   #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/search/organizations \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   \"search\":{     \"fields\" : [\"settings.name\"],     \"searchTerm\": \"Second\"     },   \"fields\": [\"_id\", \"displayName\", \"logoUrl\"],   \"limit\" : 0,   \"skip\" : 0 }' ```

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
 **body** | [**Search**](Search.md)|  | 
 **fields** | **string**| Use a space seperated string of field parameters to include the data in the response. If omitted the default list of fields will be returned.  | [default to ]
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **string**|  | [default to &lt;&lt;your org id&gt;&gt;]
 **skip** | **int32**| The offset into the records to return. | [default to 0]

### Return type

[**Systemslist**](systemslist.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSystemsPost**
> Systemslist SearchSystemsPost(ctx, contentType, accept, optional)
Search Systems

Return Systems in multi-record format allowing for the passing of the 'filter' parameter. This WILL NOT allow you to add a new system.  To support advanced filtering you can use the `filter` parameter that can only be passed in the body of POST /api/search/_* routes. The `filter` parameter must be passed as Content-Type application/json supports advanced filtering using the mongodb JSON query syntax.   The `filter` parameter is an object with a single property, either and or or with the value of the property being an array of query expressions.   This allows you to filter records using the logic of matching ALL or ANY records in the array of query expressions. If the and or or are not included the default behavior is to match ALL query expressions.   #### Sample Request  ``` curl -X POST https://console.jumpcloud.com/api/search/systems \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{ \"filter\" :     {         \"or\" :             [                 {\"hostname\" : { \"$regex\" : \"^www\" }},                 {\"hostname\" : {\"$regex\" : \"^db\"}}             ]     }, \"fields\" : \"os hostname displayName\" }' ```

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
 **body** | [**Search**](Search.md)|  | 
 **fields** | **string**| Use a space seperated string of field parameters to include the data in the response. If omitted the default list of fields will be returned.  | [default to ]
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **string**|  | [default to &lt;&lt;your org id&gt;&gt;]
 **skip** | **int32**| The offset into the records to return. | [default to 0]

### Return type

[**Systemslist**](systemslist.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSystemusersPost**
> Systemuserslist SearchSystemusersPost(ctx, contentType, accept, optional)
Search System Users

Return System Users in multi-record format allowing for the passing of the 'filter' parameter. This WILL NOT allow you to add a new system user.  To support advanced filtering you can use the `filter` parameter that can only be passed in the body of POST /api/search/_* routes. The `filter` parameter must be passed as Content-Type application/json supports advanced filtering using the mongodb JSON query syntax.   The `filter` parameter is an object with a single property, either and or or with the value of the property being an array of query expressions.   This allows you to filter records using the logic of matching ALL or ANY records in the array of query expressions. If the and or or are not included the default behavior is to match ALL query expressions.  #### Sample Request  ``` curl -X POST https://console.jumpcloud.com/api/search/systemsusers \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{ \"filter\" : [{\"email\" : { \"$regex\" : \"gmail.com$\"}}], \"fields\" : \"email username sudo\" }' ```

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
 **body** | [**Search**](Search.md)|  | 
 **fields** | **string**| Use a space seperated string of field parameters to include the data in the response. If omitted the default list of fields will be returned.  | [default to ]
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **xOrgId** | **string**|  | [default to &lt;&lt;your org id&gt;&gt;]

### Return type

[**Systemuserslist**](systemuserslist.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

