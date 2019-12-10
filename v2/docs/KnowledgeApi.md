# \KnowledgeApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KnowledgeArticlesList**](KnowledgeApi.md#KnowledgeArticlesList) | **Get** /knowledge/articles | List Knowledge Articles
[**KnowledgeSalesforceList**](KnowledgeApi.md#KnowledgeSalesforceList) | **Get** /knowledge/salesforce | List Knowledge Articles


# **KnowledgeArticlesList**
> SalesforceKnowledgeListOutput KnowledgeArticlesList(ctx, optional)
List Knowledge Articles

This endpoint returns a list of knowledge articles hosted in salesforce.  ``` Sample Request curl -X GET https://console.jumpcloud.com/api/v2/knowledge/articles \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | [**[]string**](string.md)| Supported operators are: eq, ne, gt, ge, lt, le, between, search, in | 
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 
 **limit** | **int32**|  | [default to 10]

### Return type

[**SalesforceKnowledgeListOutput**](salesforce-knowledge-list-output.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KnowledgeSalesforceList**
> SalesforceKnowledgeListOutput KnowledgeSalesforceList(ctx, optional)
List Knowledge Articles

This endpoint returns a list of knowledge articles hosted in salesforce.  ``` Sample Request curl -X GET https://console.jumpcloud.com/api/v2/knowledge/articles \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**[]string**](string.md)|  | 
 **filter** | [**[]string**](string.md)| Supported operators are: eq, ne, gt, ge, lt, le, between, search, in | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 

### Return type

[**SalesforceKnowledgeListOutput**](salesforce-knowledge-list-output.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

