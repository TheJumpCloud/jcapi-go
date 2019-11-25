# \SystemsApi

All URIs are relative to *https://console.jumpcloud.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemsDelete**](SystemsApi.md#SystemsDelete) | **Delete** /systems/{id} | Delete a System
[**SystemsGet**](SystemsApi.md#SystemsGet) | **Get** /systems/{id} | List an individual system
[**SystemsList**](SystemsApi.md#SystemsList) | **Get** /systems | List All Systems
[**SystemsPut**](SystemsApi.md#SystemsPut) | **Put** /systems/{id} | Update a system
[**SystemsSystemusersBindingList**](SystemsApi.md#SystemsSystemusersBindingList) | **Get** /systems/{id}/systemusers | List system user bindings
[**SystemsSystemusersBindingPut**](SystemsApi.md#SystemsSystemusersBindingPut) | **Put** /systems/{id}/systemusers | Update a system&#39;s or user&#39;s binding


# **SystemsDelete**
> System SystemsDelete(ctx, id, contentType, accept, optional)
Delete a System

This endpoint allows you to delete a system. This command will cause the system to uninstall the JumpCloud agent from its self which can can take about a minute. If the system is not connected to JumpCloud the system record will simply be removed.  #### Sample Request ``` curl -X DELETE https://console.jumpcloud.com/api/systems/{SystemID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'   ```

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
 **date** | **string**| Current date header for the System Context API | 
 **authorization** | **string**| Authorization header for the System Context API | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**System**](system.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemsGet**
> System SystemsGet(ctx, id, contentType, accept, optional)
List an individual system

This endpoint returns an individual system.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/systems/{SystemID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'   ```

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
 **fields** | **string**| Use a space seperated string of field parameters to include the data in the response. If omitted, the default list of fields will be returned.  | [default to ]
 **filter** | **string**| A filter to apply to the query. | 
 **date** | **string**| Current date header for the System Context API | 
 **authorization** | **string**| Authorization header for the System Context API | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**System**](system.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemsList**
> Systemslist SystemsList(ctx, contentType, accept, optional)
List All Systems

This endpoint returns all Systems.  #### Sample Requests ``` curl -X GET https://console.jumpcloud.com/api/systems \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'  ```

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
 **fields** | **string**| Use a space seperated string of field parameters to include the data in the response. If omitted, the default list of fields will be returned.  | [default to ]
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **string**|  | [default to ]
 **search** | **string**| A nested object containing a string &#x60;searchTerm&#x60; and a list of &#x60;fields&#x60; to search on. | 
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | **string**| Use space separated sort parameters to sort the collection. Default sort is ascending. Prefix with &#x60;-&#x60; to sort descending.  | [default to ]
 **filter** | **string**| A filter to apply to the query. | 

### Return type

[**Systemslist**](systemslist.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemsPut**
> System SystemsPut(ctx, id, contentType, accept, optional)
Update a system

This endpoint allows you to update a system.  #### Sample Request  ``` curl -X PUT https://console.jumpcloud.com/api/systems/{SystemID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{  \"displayName\":\"Name_Update\",  \"allowSshPasswordAuthentication\":\"true\",  \"allowSshRootLogin\":\"true\",  \"allowMultiFactorAuthentication\":\"true\",  \"allowPublicKeyAuthentication\":\"false\" }' ```

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
 **body** | [**Systemput**](Systemput.md)|  | 
 **date** | **string**| Current date header for the System Context API | 
 **authorization** | **string**| Authorization header for the System Context API | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**System**](system.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemsSystemusersBindingList**
> Systemuserbinding SystemsSystemusersBindingList(ctx, id, contentType, accept, optional)
List system user bindings

Hidden as Tags is deprecated  List system user bindings for a specific system in a system and user binding format.  This endpoint is only used for users still using JumpCloud Tags. If you are using JumpCloud Groups please refer to the documentation found [here](https://docs.jumpcloud.com/2.0/systems/manage-associations-of-a-system).  #### Sample Request  *List system user bindings for specific system*  ``` curl -X https://console.jumpcloud.com/api/systems/{SystemID}/systemusers\\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   \" ```

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
 **fields** | **string**| Use a space seperated string of field parameters to include the data in the response. If omitted, the default list of fields will be returned.  | [default to ]
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | **string**| Use space separated sort parameters to sort the collection. Default sort is ascending. Prefix with &#x60;-&#x60; to sort descending.  | [default to ]
 **filter** | **string**| A filter to apply to the query. | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**Systemuserbinding**](systemuserbinding.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemsSystemusersBindingPut**
> SystemsSystemusersBindingPut(ctx, id, contentType, accept, optional)
Update a system's or user's binding

Hidden as Tags is deprecated  Adds or removes a user binding for a system.  This endpoint is only used for users still using JumpCloud Tags. If you are using JumpCloud Groups please refer to the documentation found [here](https://docs.jumpcloud.com/2.0/systems/manage-associations-of-a-system).  #### Sample Request  *Add (or remove) a system user to (from) a system*  ``` curl \\   -d '{ \"add\": [\"[SYSTEM_USER_ID_TO_ADD_HERE]\"], \"remove\": [\"[SYSTEM_USER_ID_TO_REMOVE_HERE]\"] }' \\   -X PUT \\   -H 'Content-Type: application/json' \\   -H 'Accept: application/json' \\   -H \"x-api-key: [YOUR_API_KEY_HERE]\" \\   \"https://console.jumpcloud.com/api/systems/[SYSTEM_ID_HERE]/systemusers

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
 **body** | [**Systemuserbindingsput**](Systemuserbindingsput.md)|  | 
 **xOrgId** | **string**|  | [default to ]

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

