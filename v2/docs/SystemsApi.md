# \SystemsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphSystemAssociationsList**](SystemsApi.md#GraphSystemAssociationsList) | **Get** /systems/{system_id}/associations | List the associations of a System
[**GraphSystemAssociationsPost**](SystemsApi.md#GraphSystemAssociationsPost) | **Post** /systems/{system_id}/associations | Manage associations of a System
[**GraphSystemMemberOf**](SystemsApi.md#GraphSystemMemberOf) | **Get** /systems/{system_id}/memberof | List the parent Groups of a System
[**GraphSystemTraverseCommand**](SystemsApi.md#GraphSystemTraverseCommand) | **Get** /systems/{system_id}/commands | List the Commands bound to a System
[**GraphSystemTraversePolicy**](SystemsApi.md#GraphSystemTraversePolicy) | **Get** /systems/{system_id}/policies | List the Policies bound to a System
[**GraphSystemTraverseUser**](SystemsApi.md#GraphSystemTraverseUser) | **Get** /systems/{system_id}/users | List the Users bound to a System
[**GraphSystemTraverseUserGroup**](SystemsApi.md#GraphSystemTraverseUserGroup) | **Get** /systems/{system_id}/usergroups | List the User Groups bound to a System
[**SystemsGetFDEKey**](SystemsApi.md#SystemsGetFDEKey) | **Get** /systems/{system_id}/fdekey | Get System FDE Key


# **GraphSystemAssociationsList**
> []GraphConnection GraphSystemAssociationsList(ctx, systemId, contentType, accept, targets, optional)
List the associations of a System

This endpoint returns the _direct_ associations of a System.  A direct association can be a non-homogeneous relationship between 2 different objects, for example Systems and Users.   #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/systems/{System_ID}/associations?targets=user \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'  ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **systemId** | **string**| ObjectID of the System. | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
  **targets** | [**[]string**](string.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemId** | **string**| ObjectID of the System. | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **targets** | [**[]string**](string.md)|  | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **date** | **string**| Current date header for the System Context API | 
 **authorization** | **string**| Authorization header for the System Context API | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]GraphConnection**](GraphConnection.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphSystemAssociationsPost**
> GraphSystemAssociationsPost(ctx, systemId, contentType, accept, optional)
Manage associations of a System

This endpoint allows you to manage the _direct_ associations of a System.  A direct association can be a non-homogeneous relationship between 2 different objects, for example Systems and Users.   #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/systems/{System_ID}/associations \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{    \"attributes\": {       \"sudo\": {          \"enabled\": true,          \"withoutPassword\": false       }    },     \"op\": \"add\",     \"type\": \"user\",     \"id\": \"UserID\" }'  ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **systemId** | **string**| ObjectID of the System. | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemId** | **string**| ObjectID of the System. | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **body** | [**SystemGraphManagementReq**](SystemGraphManagementReq.md)|  | 
 **date** | **string**| Current date header for the System Context API | 
 **authorization** | **string**| Authorization header for the System Context API | 
 **xOrgId** | **string**|  | [default to ]

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphSystemMemberOf**
> []GraphObjectWithPaths GraphSystemMemberOf(ctx, systemId, contentType, accept, optional)
List the parent Groups of a System

This endpoint returns all the System Groups a System is a member of.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/systems/{System_ID}/memberof \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'  ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **systemId** | **string**| ObjectID of the System. | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemId** | **string**| ObjectID of the System. | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **filter** | [**[]string**](string.md)| Supported operators are: eq, ne, gt, ge, lt, le, between, search, in | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **date** | **string**| Current date header for the System Context API | 
 **authorization** | **string**| Authorization header for the System Context API | 
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]GraphObjectWithPaths**](GraphObjectWithPaths.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphSystemTraverseCommand**
> []GraphObjectWithPaths GraphSystemTraverseCommand(ctx, systemId, contentType, accept, optional)
List the Commands bound to a System

This endpoint will return all Commands bound to a System, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this System to the corresponding Command; this array represents all grouping and/or associations that would have to be removed to deprovision the Command from this System.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/systems/{System_ID}/commands \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'  ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **systemId** | **string**| ObjectID of the System. | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemId** | **string**| ObjectID of the System. | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **string**|  | [default to ]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq, ne, gt, ge, lt, le, between, search, in | 

### Return type

[**[]GraphObjectWithPaths**](GraphObjectWithPaths.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphSystemTraversePolicy**
> []GraphObjectWithPaths GraphSystemTraversePolicy(ctx, systemId, contentType, accept, optional)
List the Policies bound to a System

This endpoint will return all Policies bound to a System, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this System to the corresponding Policy; this array represents all grouping and/or associations that would have to be removed to deprovision the Policy from this System.  See `/members` and `/associations` endpoints to manage those collections.  This endpoint is not yet public as we have finish the code.  ##### Sample Request  ``` curl -X GET https://console.jumpcloud.com/api/v2/{System_ID}/policies \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'  ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **systemId** | **string**| ObjectID of the System. | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemId** | **string**| ObjectID of the System. | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **string**|  | [default to ]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq, ne, gt, ge, lt, le, between, search, in | 

### Return type

[**[]GraphObjectWithPaths**](GraphObjectWithPaths.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphSystemTraverseUser**
> []GraphObjectWithPaths GraphSystemTraverseUser(ctx, systemId, contentType, accept, optional)
List the Users bound to a System

This endpoint will return all Users bound to a System, either directly or indirectly essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this System to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this System.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/systems/{System_ID}/users \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'  ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **systemId** | **string**| ObjectID of the System. | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemId** | **string**| ObjectID of the System. | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **string**|  | [default to ]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **date** | **string**| Current date header for the System Context API | 
 **authorization** | **string**| Authorization header for the System Context API | 
 **filter** | [**[]string**](string.md)| Supported operators are: eq, ne, gt, ge, lt, le, between, search, in | 

### Return type

[**[]GraphObjectWithPaths**](GraphObjectWithPaths.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphSystemTraverseUserGroup**
> []GraphObjectWithPaths GraphSystemTraverseUserGroup(ctx, systemId, contentType, accept, optional)
List the User Groups bound to a System

This endpoint will return all User Groups bound to a System, either directly or indirectly essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this System to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this System.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/systems/{System_ID}/usergroups \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'  ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **systemId** | **string**| ObjectID of the System. | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemId** | **string**| ObjectID of the System. | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **string**|  | [default to ]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **date** | **string**| Current date header for the System Context API | 
 **authorization** | **string**| Authorization header for the System Context API | 
 **filter** | [**[]string**](string.md)| Supported operators are: eq, ne, gt, ge, lt, le, between, search, in | 

### Return type

[**[]GraphObjectWithPaths**](GraphObjectWithPaths.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemsGetFDEKey**
> Systemfdekey SystemsGetFDEKey(ctx, systemId, optional)
Get System FDE Key

This endpoint will return the current (latest) fde key saved for a system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **systemId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemId** | **string**|  | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**Systemfdekey**](systemfdekey.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

