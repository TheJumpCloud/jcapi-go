# {{classname}}

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphRadiusServerAssociationsList**](RADIUSServersApi.md#GraphRadiusServerAssociationsList) | **Get** /radiusservers/{radiusserver_id}/associations | List the associations of a RADIUS  Server
[**GraphRadiusServerAssociationsPost**](RADIUSServersApi.md#GraphRadiusServerAssociationsPost) | **Post** /radiusservers/{radiusserver_id}/associations | Manage the associations of a RADIUS Server
[**GraphRadiusServerTraverseUser**](RADIUSServersApi.md#GraphRadiusServerTraverseUser) | **Get** /radiusservers/{radiusserver_id}/users | List the Users bound to a RADIUS  Server
[**GraphRadiusServerTraverseUserGroup**](RADIUSServersApi.md#GraphRadiusServerTraverseUserGroup) | **Get** /radiusservers/{radiusserver_id}/usergroups | List the User Groups bound to a RADIUS  Server

# **GraphRadiusServerAssociationsList**
> []GraphConnection GraphRadiusServerAssociationsList(ctx, radiusserverId, targets, optional)
List the associations of a RADIUS  Server

This endpoint returns the _direct_ associations of a Radius Server.  A direct association can be a non-homogeneous relationship between 2 different objects, for example Radius Servers and Users.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/radiusservers/{RADIUS_ID}/associations?targets=user_group \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **radiusserverId** | **string**| ObjectID of the Radius Server. | 
  **targets** | [**[]string**](string.md)| Targets which a \&quot;radius_server\&quot; can be associated to. | 
 **optional** | ***RADIUSServersApiGraphRadiusServerAssociationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RADIUSServersApiGraphRadiusServerAssociationsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **optional.Int32**| The offset into the records to return. | [default to 0]
 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

[**[]GraphConnection**](GraphConnection.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphRadiusServerAssociationsPost**
> GraphRadiusServerAssociationsPost(ctx, radiusserverId, optional)
Manage the associations of a RADIUS Server

This endpoint allows you to manage the _direct_ associations of a Radius Server.  A direct association can be a non-homogeneous relationship between 2 different objects, for example Radius Servers and Users.  #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/radiusservers/{RADIUS_ID}/associations \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   \"type\":\"user\",  \"id\":\"{USER_ID}\",  \"op\":\"add\"   }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **radiusserverId** | **string**| ObjectID of the Radius Server. | 
 **optional** | ***RADIUSServersApiGraphRadiusServerAssociationsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RADIUSServersApiGraphRadiusServerAssociationsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GraphOperationRadiusServer**](GraphOperationRadiusServer.md)|  | 
 **xOrgId** | **optional.**| Organization identifier that can be obtained from console settings. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphRadiusServerTraverseUser**
> []GraphObjectWithPaths GraphRadiusServerTraverseUser(ctx, radiusserverId, optional)
List the Users bound to a RADIUS  Server

This endpoint will return all Users bound to a RADIUS Server, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this RADIUS server instance to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this RADIUS server instance.  See `/members` and `/associations` endpoints to manage those collections.   #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID}/users \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **radiusserverId** | **string**| ObjectID of the Radius Server. | 
 **optional** | ***RADIUSServersApiGraphRadiusServerTraverseUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RADIUSServersApiGraphRadiusServerTraverseUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 
 **skip** | **optional.Int32**| The offset into the records to return. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 

### Return type

[**[]GraphObjectWithPaths**](GraphObjectWithPaths.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphRadiusServerTraverseUserGroup**
> []GraphObjectWithPaths GraphRadiusServerTraverseUserGroup(ctx, radiusserverId, optional)
List the User Groups bound to a RADIUS  Server

This endpoint will return all Users Groups bound to a RADIUS Server, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the group's type, id, attributes and paths.  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this RADIUS server instance to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this RADIUS server instance.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/radiusservers/{RADIUS_ID}/usergroups \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **radiusserverId** | **string**| ObjectID of the Radius Server. | 
 **optional** | ***RADIUSServersApiGraphRadiusServerTraverseUserGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RADIUSServersApiGraphRadiusServerTraverseUserGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 
 **skip** | **optional.Int32**| The offset into the records to return. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 

### Return type

[**[]GraphObjectWithPaths**](GraphObjectWithPaths.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

