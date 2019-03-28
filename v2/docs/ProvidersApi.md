# \ProvidersApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProvidersListAdministrators**](ProvidersApi.md#ProvidersListAdministrators) | **Get** /providers/{provider_id}/administrators | List Provider Administrators
[**ProvidersPostAdmins**](ProvidersApi.md#ProvidersPostAdmins) | **Post** /providers/{provider_id}/administrators | Create a new Provider Administrator


# **ProvidersListAdministrators**
> InlineResponse200 ProvidersListAdministrators(ctx, providerId, contentType, accept, optional)
List Provider Administrators

This endpoint returns a list of the Administrators associated with the Provider.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/providers/{ProviderID}/administrators \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string**|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **fields** | [**[]string**](string.md)| The comma separated fields included in the returned records. If omitted the default list of fields will be returned.  | 
 **filter** | [**[]string**](string.md)| Supported operators are: eq, ne, gt, ge, lt, le, between, search, in | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProvidersPostAdmins**
> Administrator ProvidersPostAdmins(ctx, providerId, contentType, accept, optional)
Create a new Provider Administrator

This endpoint allows you to create a provider administrator. You must be associated to given provider to use this route.  **Sample Request**  ``` curl -X POST https://console.jumpcloud.com/api/v2/providers/{ProviderID}/administrators \\     -H 'Accept: application/json' \\     -H 'Context-Type: application/json' \\     -H 'x-api-key: {API_KEY}' \\     -d '{       \"email\":\"{ADMIN_EMAIL}\"     }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providerId** | **string**|  | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string**|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **body** | [**ProviderAdminReq**](ProviderAdminReq.md)|  | 

### Return type

[**Administrator**](Administrator.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

