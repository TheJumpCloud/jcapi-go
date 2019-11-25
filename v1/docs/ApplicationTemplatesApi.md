# \ApplicationTemplatesApi

All URIs are relative to *https://console.jumpcloud.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationTemplatesGet**](ApplicationTemplatesApi.md#ApplicationTemplatesGet) | **Get** /application-templates/{id} | Get an Application Template
[**ApplicationTemplatesList**](ApplicationTemplatesApi.md#ApplicationTemplatesList) | **Get** /application-templates | List Application Templates


# **ApplicationTemplatesGet**
> Applicationtemplate ApplicationTemplatesGet(ctx, id, contentType, accept, optional)
Get an Application Template

The endpoint returns a specific SSO / SAML Application Template.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/application-templates/{id} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'  ```

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
 **fields** | **string**| The comma separated fields included in the returned records. If omitted the default list of fields will be returned. | 
 **limit** | **int32**| The number of records to return at once. | 
 **skip** | **int32**| The offset into the records to return. | 
 **sort** | **string**|  | [default to The comma separated fields used to sort the collection. Default sort is ascending, prefix with - to sort descending.]
 **filter** | **string**| A filter to apply to the query. | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**Applicationtemplate**](applicationtemplate.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationTemplatesList**
> Applicationtemplateslist ApplicationTemplatesList(ctx, contentType, accept, optional)
List Application Templates

The endpoint returns all the SSO / SAML Application Templates.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/application-templates \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'  ```

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
 **fields** | **string**| The comma separated fields included in the returned records. If omitted the default list of fields will be returned. | 
 **limit** | **int32**| The number of records to return at once. | 
 **skip** | **int32**| The offset into the records to return. | 
 **sort** | **string**|  | [default to The comma separated fields used to sort the collection. Default sort is ascending, prefix with - to sort descending.]
 **filter** | **string**| A filter to apply to the query. | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**Applicationtemplateslist**](applicationtemplateslist.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

