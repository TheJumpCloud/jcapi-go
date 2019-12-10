# \OrganizationsApi

All URIs are relative to *https://console.jumpcloud.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationList**](OrganizationsApi.md#OrganizationList) | **Get** /organizations | Get Organization Details
[**OrganizationPut**](OrganizationsApi.md#OrganizationPut) | **Put** /organizations/{id} | Update an Organization
[**OrganizationsGet**](OrganizationsApi.md#OrganizationsGet) | **Get** /organizations/{id} | Get an Organization


# **OrganizationList**
> Organizationslist OrganizationList(ctx, contentType, accept, optional)
Get Organization Details

This endpoint returns Organization Details.  #### Sample Request   ``` curl -X GET \\   https://console.jumpcloud.com/api/organizations \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'   ```

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
 **filter** | **string**| A filter to apply to the query. | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | **string**| Use space separated sort parameters to sort the collection. Default sort is ascending. Prefix with &#x60;-&#x60; to sort descending.  | [default to ]
 **search** | **string**| A nested object containing a string &#x60;searchTerm&#x60; and a list of &#x60;fields&#x60; to search on. | 

### Return type

[**Organizationslist**](organizationslist.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationPut**
> Organization OrganizationPut(ctx, id, contentType, accept, optional)
Update an Organization

This endpoint allows you to update an Organization.  Note: `passwordPolicy` settings are only used when `passwordCompliance` is set to \"custom\". We discourage the use of non-custom passwordCompliance values.  `hasStripeCustomerId` is deprecated and will be removed.  #### Sample Request  ``` curl -X PUT https://console.jumpcloud.com/api/organizations/{OrganizationID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   \"settings\": {     \"contactName\": \"Admin Name\",     \"contactEmail\": \"admin@company.com\",     \"systemUsersCanEdit\":true,     \"passwordPolicy\": {       \"enableMaxHistory\": true,       \"maxHistory\": 3     }   } }' ```

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
 **body** | [**Body**](Body.md)|  | 

### Return type

[**Organization**](organization.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsGet**
> OrganizationsGet(ctx, id, contentType, accept, optional)
Get an Organization

This endpoint returns a particular Organization.  #### Sample Request  ``` curl -X GET https://console.jumpcloud.com/api/organizations/{OrganizationID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

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

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

