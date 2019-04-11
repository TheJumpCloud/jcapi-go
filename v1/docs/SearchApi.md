# \SearchApi

All URIs are relative to *https://console.jumpcloud.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchOrganizationsPost**](SearchApi.md#SearchOrganizationsPost) | **Post** /search/organizations | Search Organizations
[**SearchSystemsPost**](SearchApi.md#SearchSystemsPost) | **Post** /search/systems | Search Systems
[**SearchSystemusersPost**](SearchApi.md#SearchSystemusersPost) | **Post** /search/systemusers | Search System Users


# **SearchOrganizationsPost**
> Organizationslist SearchOrganizationsPost(ctx, contentType, accept, optional)
Search Organizations

This endpoint will return Organization data based on your search parameters. This endpoint WILL NOT allow you to add a new Organization.  You can use the supported parameters and pass those in the body of request.   The parameters must be passed as Content-Type application/json.   #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/search/organizations \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   \"search\":{     \"fields\" : [\"settings.name\"],     \"searchTerm\": \"Second\"     },   \"fields\": [\"_id\", \"displayName\", \"logoUrl\"],   \"limit\" : 0,   \"skip\" : 0 }' ```

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
 **fields** | **string**| Use a space seperated string of field parameters to include the data in the response. If omitted, the default list of fields will be returned.  | [default to ]
 **filter** | **string**| A filter to apply to the query. | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]

### Return type

[**Organizationslist**](organizationslist.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSystemsPost**
> Systemslist SearchSystemsPost(ctx, contentType, accept, optional)
Search Systems

Return Systems in multi-record format allowing for the passing of the `filter` and `searchFilter` parameters. This WILL NOT allow you to add a new system.  To support advanced filtering you can use the `filter` and `searchFilter` parameters that can only be passed in the body of POST /api/search/_* routes. The `filter` and `searchFilter` parameters must be passed as Content-Type application/json.  The `filter` parameter is an object with a single property, either `and` or `or` with the value of the property being an array of query expressions.   This allows you to filter records using the logic of matching ALL or ANY records in the array of query expressions. If the `and` or `or` are not included the default behavior is to match ALL query expressions.  The `searchFilter` parameter allows text searching on supported fields by specifying a `searchTerm` and a list of `fields` to query on. If any `field` has a partial text match on the`searchTerm` the record will be returned.   #### Sample Request  Exact search for a list of hostnames ``` curl -X POST https://console.jumpcloud.com/api/search/systems \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   \"filter\": {     \"or\": [       {\"hostname\" : \"my-hostname\"},       {\"hostname\" : \"other-hostname\"}     ]   },   \"fields\" : \"os hostname displayName\" }' ```  Text search for a hostname or display name ``` curl -X POST https://console.jumpcloud.com/api/search/systems \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   \"searchFilter\": {     \"searchTerm\": \"my-host\",     \"fields\": [\"hostname\", \"displayName\"]   },   \"fields\": \"os hostname displayName\" }' ```  Combining `filter` and `searchFilter` to search for names that match a given OS ``` curl -X POST https://console.jumpcloud.com/api/search/systems \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   \"searchFilter\": {     \"searchTerm\": \"my-host\",     \"fields\": [\"hostname\", \"displayName\"]   },   \"filter\": {     \"or\": [       {\"os\" : \"Ubuntu\"},       {\"os\" : \"Mac OS X\"}     ]   },   \"fields\": \"os hostname displayName\" }' ```

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
 **fields** | **string**| Use a space seperated string of field parameters to include the data in the response. If omitted, the default list of fields will be returned.  | [default to ]
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **string**|  | [default to ]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | **string**| A filter to apply to the query. | 

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

Return System Users in multi-record format allowing for the passing of the `filter` and `searchFilter` parameters. This WILL NOT allow you to add a new system user.  To support advanced filtering you can use the `filter` and `searchFilter` parameters that can only be passed in the body of POST /api/search/_* routes. The `filter` and `searchFilter` parameters must be passed as Content-Type application/json.  The `filter` parameter is an object with a single property, either `and` or `or` with the value of the property being an array of query expressions.   This allows you to filter records using the logic of matching ALL or ANY records in the array of query expressions. If the `and` or `or` are not included the default behavior is to match ALL query expressions.  The `searchFilter` parameter allows text searching on supported fields by specifying a `searchTerm` and a list of `fields` to query on. If any `field` has a partial text match on the`searchTerm` the record will be returned.   #### Sample Request  Exact search for a list of system users in a department ``` curl -X POST https://console.jumpcloud.com/api/search/systemusers \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   \"filter\" : [{\"department\" : \"IT\"}],   \"fields\" : \"email username sudo\" }' ```  Text search for system users with and email on a domain ``` curl -X POST https://console.jumpcloud.com/api/search/systemusers \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   \"searchFilter\" : {     \"searchTerm\": \"@jumpcloud.com\",     \"fields\": [\"email\"]   },   \"fields\" : \"email username sudo\" }' ```  Combining `filter` and `searchFilter` to text search for system users with and email on a domain who are in a list of departments ``` curl -X POST https://console.jumpcloud.com/api/search/systemusers \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   \"searchFilter\": {     \"searchTerm\": \"@jumpcloud.com\",     \"fields\": [\"email\"]   },   \"filter\": {     \"or\": [       {\"department\" : \"IT\"},       {\"department\" : \"Sales\"}     ]   },   \"fields\" : \"email username sudo\" }' ```

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
 **fields** | **string**| Use a space seperated string of field parameters to include the data in the response. If omitted, the default list of fields will be returned.  | [default to ]
 **filter** | **string**| A filter to apply to the query. | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **xOrgId** | **string**|  | [default to ]

### Return type

[**Systemuserslist**](systemuserslist.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

