# \RADIUSServersApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphRadiusServerAssociationsList**](RADIUSServersApi.md#GraphRadiusServerAssociationsList) | **Get** /radiusservers/{radiusserver_id}/associations | List the associations of a Radius Server
[**GraphRadiusServerAssociationsPost**](RADIUSServersApi.md#GraphRadiusServerAssociationsPost) | **Post** /radiusservers/{radiusserver_id}/associations | Manage the associations of a Radius Server
[**GraphRadiusServerTraverseUser**](RADIUSServersApi.md#GraphRadiusServerTraverseUser) | **Get** /radiusservers/{radiusserver_id}/users | List the Users associated with a Radius Server
[**GraphRadiusServerTraverseUserGroup**](RADIUSServersApi.md#GraphRadiusServerTraverseUserGroup) | **Get** /radiusservers/{radiusserver_id}/usergroups | List the User Groups associated with a Radius Server


# **GraphRadiusServerAssociationsList**
> []GraphConnection GraphRadiusServerAssociationsList(ctx, radiusserverId, targets, contentType, accept, optional)
List the associations of a Radius Server

This endpoint returns the _direct_ associations of a Radius Server.  A direct association can be a non-homogenous relationship between 2 different objects. for example Radius Servers and Users.  #### Sample Request ``` https://console.jumpcloud.com/api/v2/radiusservers/{radiusserver_id}/associations?targets=user ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **radiusserverId** | **string**| ObjectID of the Radius Server. | 
  **targets** | [**[]string**](string.md)|  | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **radiusserverId** | **string**| ObjectID of the Radius Server. | 
 **targets** | [**[]string**](string.md)|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
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

# **GraphRadiusServerAssociationsPost**
> GraphRadiusServerAssociationsPost(ctx, radiusserverId, contentType, accept, optional)
Manage the associations of a Radius Server

This endpoint allows you to manage the _direct_ associations of a Radius Server.  A direct association can be a non-homogenous relationship between 2 different objects. for example Radius Servers and Users.  #### Sample Request ``` https://console.jumpcloud.com/api/v2/radiusservers/{radiusserver_id}/associations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **radiusserverId** | **string**| ObjectID of the Radius Server. | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **radiusserverId** | **string**| ObjectID of the Radius Server. | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **body** | [**GraphManagementReq**](GraphManagementReq.md)|  | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphRadiusServerTraverseUser**
> []GraphObjectWithPaths GraphRadiusServerTraverseUser(ctx, radiusserverId, contentType, accept, optional)
List the Users associated with a Radius Server

This endpoint will return Users associated with a RADIUS server instance. Each element will contain the type, id, attributes and paths.  The `attributes` object is a key/value hash of attributes specifically set for this group.  The `paths` array enumerates each path from this RADIUS server instance to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this RADIUS server instance.  See `/members` and `/associations` endpoints to manage those collections.   #### Sample Request ``` https://console.jumpcloud.com/api/v2/radiusservers/{radiusserver_id}/users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **radiusserverId** | **string**| ObjectID of the Radius Server. | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **radiusserverId** | **string**| ObjectID of the Radius Server. | 
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

# **GraphRadiusServerTraverseUserGroup**
> []GraphObjectWithPaths GraphRadiusServerTraverseUserGroup(ctx, radiusserverId, contentType, accept, optional)
List the User Groups associated with a Radius Server

This endpoint will return User Groups associated with a RADIUS server instance. Each element will contain the group's type, id, attributes and paths.  The `attributes` object is a key/value hash of attributes specifically set for this group.  The `paths` array enumerates each path from this RADIUS server instance to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this RADIUS server instance.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` https://console.jumpcloud.com/api/v2/radiusservers/{radiusserver_id}/usergroups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **radiusserverId** | **string**| ObjectID of the Radius Server. | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **radiusserverId** | **string**| ObjectID of the Radius Server. | 
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

