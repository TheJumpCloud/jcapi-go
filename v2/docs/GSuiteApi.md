# \GSuiteApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphGSuiteAssociationsList**](GSuiteApi.md#GraphGSuiteAssociationsList) | **Get** /gsuites/{gsuite_id}/associations | List the associations of a G Suite instance
[**GraphGSuiteAssociationsPost**](GSuiteApi.md#GraphGSuiteAssociationsPost) | **Post** /gsuites/{gsuite_id}/associations | Manage the associations of a G Suite instance
[**GraphGSuiteTraverseUser**](GSuiteApi.md#GraphGSuiteTraverseUser) | **Get** /gsuites/{gsuite_id}/users | List the Users bound to a G Suite instance
[**GraphGSuiteTraverseUserGroup**](GSuiteApi.md#GraphGSuiteTraverseUserGroup) | **Get** /gsuites/{gsuite_id}/usergroups | List the User Groups bound to a G Suite instance
[**TranslationRulesGSuiteDelete**](GSuiteApi.md#TranslationRulesGSuiteDelete) | **Delete** /gsuites/{gsuite_id}/translationrules/{id} | Deletes a G Suite translation rule
[**TranslationRulesGSuiteGet**](GSuiteApi.md#TranslationRulesGSuiteGet) | **Get** /gsuites/{gsuite_id}/translationrules/{id} | Gets a specific g suite translation rule
[**TranslationRulesGSuiteList**](GSuiteApi.md#TranslationRulesGSuiteList) | **Get** /gsuites/{gsuite_id}/translationrules | List all the G Suite Translation Rules
[**TranslationRulesGSuitePost**](GSuiteApi.md#TranslationRulesGSuitePost) | **Post** /gsuites/{gsuite_id}/translationrules | Create a new G Suite Translation Rule


# **GraphGSuiteAssociationsList**
> []GraphConnection GraphGSuiteAssociationsList(ctx, gsuiteId, targets, contentType, accept, optional)
List the associations of a G Suite instance

This endpoint returns the _direct_ associations of this G Suite instance.  A direct association can be a non-homogeneous relationship between 2 different objects, for example G Suite and Users.   #### Sample Request ``` curl -X GET 'https://console.jumpcloud.com/api/v2/gsuites/{Gsuite_ID}/associations?targets=user_group \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gsuiteId** | **string**| ObjectID of the G Suite instance. | 
  **targets** | [**[]string**](string.md)|  | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gsuiteId** | **string**| ObjectID of the G Suite instance. | 
 **targets** | [**[]string**](string.md)|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]GraphConnection**](GraphConnection.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphGSuiteAssociationsPost**
> GraphGSuiteAssociationsPost(ctx, gsuiteId, optional)
Manage the associations of a G Suite instance

This endpoint returns the _direct_ associations of this G Suite instance.  A direct association can be a non-homogeneous relationship between 2 different objects, for example G Suite and Users.   #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/gsuites/{Gsuite_ID}/associations \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{     \"op\": \"add\",     \"type\": \"user_group\",     \"id\": \"{Group_ID}\" }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gsuiteId** | **string**| ObjectID of the G Suite instance. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gsuiteId** | **string**| ObjectID of the G Suite instance. | 
 **body** | [**GraphManagementReq**](GraphManagementReq.md)|  | 
 **xOrgId** | **string**|  | [default to ]

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphGSuiteTraverseUser**
> []GraphObjectWithPaths GraphGSuiteTraverseUser(ctx, gsuiteId, contentType, accept, optional)
List the Users bound to a G Suite instance

This endpoint will return all Users bound to a G Suite instance, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this G Suite instance to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this G Suite instance.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ```   curl -X GET https://console.jumpcloud.com/api/v2/gsuites/{Gsuite_ID}/users \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gsuiteId** | **string**| ObjectID of the G Suite instance. | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gsuiteId** | **string**| ObjectID of the G Suite instance. | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]GraphObjectWithPaths**](GraphObjectWithPaths.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphGSuiteTraverseUserGroup**
> []GraphObjectWithPaths GraphGSuiteTraverseUserGroup(ctx, gsuiteId, contentType, accept, optional)
List the User Groups bound to a G Suite instance

This endpoint will return all User Groups bound to an G Suite instance, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the group's type, id, attributes and paths.  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this G Suite instance to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this G Suite instance.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ```   curl -X GET https://console.jumpcloud.com/api/v2/gsuites/{GSuite_ID}/usergroups \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gsuiteId** | **string**| ObjectID of the G Suite instance. | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gsuiteId** | **string**| ObjectID of the G Suite instance. | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]GraphObjectWithPaths**](GraphObjectWithPaths.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TranslationRulesGSuiteDelete**
> TranslationRulesGSuiteDelete(ctx, gsuiteId, id, contentType, accept)
Deletes a G Suite translation rule

This endpoint allows you to delete a translation rule for a specific G Suite instance. These rules specify how JumpCloud attributes translate to [G Suite Admin SDK](https://developers.google.com/admin-sdk/directory/) attributes.  #### Sample Request  ``` curl -X DELETE https://console.jumpcloud.com/api/v2/gsuites/{gsuite_id}/translationrules/{id} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gsuiteId** | **string**|  | 
  **id** | **string**|  | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TranslationRulesGSuiteGet**
> GSuiteTranslationRule TranslationRulesGSuiteGet(ctx, gsuiteId, id, contentType, accept)
Gets a specific g suite translation rule

This endpoint returns a specific translation rule for a specific G Suite instance. These rules specify how JumpCloud attributes translate to [G Suite Admin SDK](https://developers.google.com/admin-sdk/directory/) attributes.  ###### Sample Request  ```   curl -X GET https://console.jumpcloud.com/api/v2/gsuites/{gsuite_id}/translationrules/{id} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gsuiteId** | **string**|  | 
  **id** | **string**|  | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]

### Return type

[**GSuiteTranslationRule**](GSuiteTranslationRule.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TranslationRulesGSuiteList**
> []GSuiteTranslationRule TranslationRulesGSuiteList(ctx, gsuiteId, contentType, accept, optional)
List all the G Suite Translation Rules

This endpoint returns all graph translation rules for a specific G Suite instance. These rules specify how JumpCloud attributes translate to [G Suite Admin SDK](https://developers.google.com/admin-sdk/directory/) attributes.  ##### Sample Request  ```  curl -X GET  https://console.jumpcloud.com/api/v2/gsuites/{gsuite_id}/translationrules \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gsuiteId** | **string**|  | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gsuiteId** | **string**|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **fields** | [**[]string**](string.md)| The comma separated fields included in the returned records. If omitted the default list of fields will be returned.  | 
 **filter** | [**[]string**](string.md)| Supported operators are: eq, ne, gt, ge, lt, le, between, search, in | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 

### Return type

[**[]GSuiteTranslationRule**](GSuiteTranslationRule.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TranslationRulesGSuitePost**
> InlineResponse201 TranslationRulesGSuitePost(ctx, gsuiteId, contentType, accept, optional)
Create a new G Suite Translation Rule

This endpoint allows you to create a translation rule for a specific G Suite instance. These rules specify how JumpCloud attributes translate to [G Suite Admin SDK](https://developers.google.com/admin-sdk/directory/) attributes.  ##### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/gsuites/{gsuite_id}/translationrules \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   {Translation Rule Parameters} }'  ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gsuiteId** | **string**|  | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gsuiteId** | **string**|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **body** | [**GSuiteTranslationRuleRequest**](GSuiteTranslationRuleRequest.md)|  | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

