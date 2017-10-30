# \SystemsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphSystemAssociationsList**](SystemsApi.md#GraphSystemAssociationsList) | **Get** /systems/{system_id}/associations | List the associations of a System
[**GraphSystemAssociationsPost**](SystemsApi.md#GraphSystemAssociationsPost) | **Post** /systems/{system_id}/associations | Manage associations of a System
[**GraphSystemMemberOf**](SystemsApi.md#GraphSystemMemberOf) | **Get** /systems/{system_id}/memberof | List the parent Groups of a System
[**GraphSystemTraversePolicy**](SystemsApi.md#GraphSystemTraversePolicy) | **Get** /systems/{system_id}/policies | List the Policies associated with a System
[**GraphSystemTraverseUser**](SystemsApi.md#GraphSystemTraverseUser) | **Get** /systems/{system_id}/users | List the Users associated with a System


# **GraphSystemAssociationsList**
> []GraphConnection GraphSystemAssociationsList(ctx, systemId, targets, contentType, accept, optional)
List the associations of a System

This endpoint returns the _direct_ associations of a System.  A direct association can be a non-homogenous relationship between 2 different objects. for example Systems and Users.   #### Sample Request ``` https://console.jumpcloud.com/api/v2/systems/{system_id}/associations?targets=user ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **systemId** | **string**| ObjectID of the System. | 
  **targets** | [**[]string**](string.md)|  | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemId** | **string**| ObjectID of the System. | 
 **targets** | [**[]string**](string.md)|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **limit** | **int32**| The number of records to return at once. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **date** | **string**| Current date header for the System Context API | 
 **authorization** | **string**| Authorization header for the System Context API | 

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

This endpoint allows you to manage the _direct_ associations of a System.  A direct association can be a non-homogenous relationship between 2 different objects. for example Systems and Users.   #### Sample Request ``` https://console.jumpcloud.com/api/v2/systems/{system_id}/associations ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
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
 **body** | [**GraphManagementReq**](GraphManagementReq.md)|  | 
 **date** | **string**| Current date header for the System Context API | 
 **authorization** | **string**| Authorization header for the System Context API | 

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

This endpoint returns all the System Groups a System is a member of.  #### Sample Request ``` https://console.jumpcloud.com/api/v2/systems/{system_id}/memberof ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
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
 **limit** | **int32**| The number of records to return at once. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **date** | **string**| Current date header for the System Context API | 
 **authorization** | **string**| Authorization header for the System Context API | 

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
List the Policies associated with a System

This endpoint will return Policies associated with a System. Each element will contain the type, id, attributes and paths.  The `attributes` object is a key/value hash of attributes specifically set for this group.  The `paths` array enumerates each path from this System to the corresponding Policy; this array represents all grouping and/or associations that would have to be removed to deprovision the Policy from this System.  See `/members` and `/associations` endpoints to manage those collections.  This endpoint is not yet public as we have finish the code.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
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
 **limit** | **int32**| The number of records to return at once. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]

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
List the Users associated with a System

This endpoint will return Users associated with a System. Each element will contain the type, id, attributes and paths.  The `attributes` object is a key/value hash of attributes specifically set for this group.  The `paths` array enumerates each path from this System to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this System.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` https://console.jumpcloud.com/api/v2/systems/{system_id}/users ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
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
 **limit** | **int32**| The number of records to return at once. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **date** | **string**| Current date header for the System Context API | 
 **authorization** | **string**| Authorization header for the System Context API | 

### Return type

[**[]GraphObjectWithPaths**](GraphObjectWithPaths.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

