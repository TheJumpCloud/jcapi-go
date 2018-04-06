# \SystemGroupAssociationsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphSystemGroupAssociationsList**](SystemGroupAssociationsApi.md#GraphSystemGroupAssociationsList) | **Get** /systemgroups/{group_id}/associations | List the associations of a System Group
[**GraphSystemGroupAssociationsPost**](SystemGroupAssociationsApi.md#GraphSystemGroupAssociationsPost) | **Post** /systemgroups/{group_id}/associations | Manage the associations of a System Group
[**GraphSystemGroupTraverseCommand**](SystemGroupAssociationsApi.md#GraphSystemGroupTraverseCommand) | **Get** /systemgroups/{group_id}/commands | List the Commands bound to a System Group
[**GraphSystemGroupTraversePolicy**](SystemGroupAssociationsApi.md#GraphSystemGroupTraversePolicy) | **Get** /systemgroups/{group_id}/policies | List the Policies bound to a System Group
[**GraphSystemGroupTraverseUser**](SystemGroupAssociationsApi.md#GraphSystemGroupTraverseUser) | **Get** /systemgroups/{group_id}/users | List the Users bound to a System Group
[**GraphSystemGroupTraverseUserGroup**](SystemGroupAssociationsApi.md#GraphSystemGroupTraverseUserGroup) | **Get** /systemgroups/{group_id}/usergroups | List the User Groups bound to a System Group


# **GraphSystemGroupAssociationsList**
> []GraphConnection GraphSystemGroupAssociationsList(ctx, groupId, contentType, accept, targets, optional)
List the associations of a System Group

This endpoint returns the _direct_ associations of a System Group.  A direct association can be a non-homogenous relationship between 2 different objects. for example System Groups and Users.   #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/systemgroups/{GroupID}/associations?targets=user \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **string**| ObjectID of the System Group. | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
  **targets** | [**[]string**](string.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string**| ObjectID of the System Group. | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **targets** | [**[]string**](string.md)|  | 
 **limit** | **int32**| The number of records to return at once. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]

### Return type

[**[]GraphConnection**](GraphConnection.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphSystemGroupAssociationsPost**
> GraphSystemGroupAssociationsPost(ctx, groupId, contentType, accept, optional)
Manage the associations of a System Group

This endpoint allows you to manage the _direct_ associations of a System Group.  A direct association can be a non-homogenous relationship between 2 different objects. for example System Groups and Users.   #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/systemgroups/{GroupID}/associations \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{     \"op\": \"add\",     \"type\": \"user\",     \"id\": \"{UserID}\" }'  ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **string**| ObjectID of the System Group. | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string**| ObjectID of the System Group. | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **body** | [**SystemGroupGraphManagementReq**](SystemGroupGraphManagementReq.md)|  | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphSystemGroupTraverseCommand**
> []GraphObjectWithPaths GraphSystemGroupTraverseCommand(ctx, groupId, contentType, accept, optional)
List the Commands bound to a System Group

This endpoint will return all Commands bound to a System Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the group's type, id, attributes and paths.  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this System Group to the corresponding Command; this array represents all grouping and/or associations that would have to be removed to deprovision the Command from this System Group.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/systemgroups/{GroupID}/commands \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **string**| ObjectID of the System Group. | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string**| ObjectID of the System Group. | 
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

# **GraphSystemGroupTraversePolicy**
> []GraphObjectWithPaths GraphSystemGroupTraversePolicy(ctx, groupId, contentType, accept, optional)
List the Policies bound to a System Group

This endpoint will return all Policies bound to a System Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this System Group to the corresponding Policy; this array represents all grouping and/or associations that would have to be removed to deprovision the Policy from this System Group.  See `/members` and `/associations` endpoints to manage those collections.  This endpoint is not public yet as we haven't finished the code.  ##### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/systemgroups/{GroupID}/policies \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **string**| ObjectID of the System Group. | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string**| ObjectID of the System Group. | 
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

# **GraphSystemGroupTraverseUser**
> []GraphObjectWithPaths GraphSystemGroupTraverseUser(ctx, groupId, contentType, accept, optional)
List the Users bound to a System Group

This endpoint will return all Users bound to a System Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this System Group to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this System Group.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/systemgroups/{GroupID}/users \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **string**| ObjectID of the System Group. | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string**| ObjectID of the System Group. | 
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

# **GraphSystemGroupTraverseUserGroup**
> []GraphObjectWithPaths GraphSystemGroupTraverseUserGroup(ctx, groupId, contentType, accept, optional)
List the User Groups bound to a System Group

This endpoint will return all User Groups bound to a System Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this System Group to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this System Group.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/systemgroups/{GroupID}/usergroups \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'  ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **string**| ObjectID of the System Group. | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string**| ObjectID of the System Group. | 
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

