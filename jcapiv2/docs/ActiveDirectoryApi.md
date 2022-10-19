# {{classname}}

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivedirectoriesAgentsDelete**](ActiveDirectoryApi.md#ActivedirectoriesAgentsDelete) | **Delete** /activedirectories/{activedirectory_id}/agents/{agent_id} | Delete Active Directory Agent
[**ActivedirectoriesAgentsGet**](ActiveDirectoryApi.md#ActivedirectoriesAgentsGet) | **Get** /activedirectories/{activedirectory_id}/agents/{agent_id} | Get Active Directory Agent
[**ActivedirectoriesAgentsList**](ActiveDirectoryApi.md#ActivedirectoriesAgentsList) | **Get** /activedirectories/{activedirectory_id}/agents | List Active Directory Agents
[**ActivedirectoriesAgentsPost**](ActiveDirectoryApi.md#ActivedirectoriesAgentsPost) | **Post** /activedirectories/{activedirectory_id}/agents | Create a new Active Directory Agent
[**ActivedirectoriesDelete**](ActiveDirectoryApi.md#ActivedirectoriesDelete) | **Delete** /activedirectories/{id} | Delete an Active Directory
[**ActivedirectoriesGet**](ActiveDirectoryApi.md#ActivedirectoriesGet) | **Get** /activedirectories/{id} | Get an Active Directory
[**ActivedirectoriesList**](ActiveDirectoryApi.md#ActivedirectoriesList) | **Get** /activedirectories | List Active Directories
[**ActivedirectoriesPost**](ActiveDirectoryApi.md#ActivedirectoriesPost) | **Post** /activedirectories | Create a new Active Directory
[**GraphActiveDirectoryAssociationsList**](ActiveDirectoryApi.md#GraphActiveDirectoryAssociationsList) | **Get** /activedirectories/{activedirectory_id}/associations | List the associations of an Active Directory instance
[**GraphActiveDirectoryAssociationsPost**](ActiveDirectoryApi.md#GraphActiveDirectoryAssociationsPost) | **Post** /activedirectories/{activedirectory_id}/associations | Manage the associations of an Active Directory instance
[**GraphActiveDirectoryTraverseUser**](ActiveDirectoryApi.md#GraphActiveDirectoryTraverseUser) | **Get** /activedirectories/{activedirectory_id}/users | List the Users bound to an Active Directory instance
[**GraphActiveDirectoryTraverseUserGroup**](ActiveDirectoryApi.md#GraphActiveDirectoryTraverseUserGroup) | **Get** /activedirectories/{activedirectory_id}/usergroups | List the User Groups bound to an Active Directory instance

# **ActivedirectoriesAgentsDelete**
> ActivedirectoriesAgentsDelete(ctx, activedirectoryId, agentId, optional)
Delete Active Directory Agent

This endpoint deletes an Active Directory agent.  #### Sample Request ``` curl -X DELETE https://console.jumpcloud.com/api/v2/activedirectories/{activedirectory_id}/agents/{agent_id} \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activedirectoryId** | **string**|  | 
  **agentId** | **string**|  | 
 **optional** | ***ActiveDirectoryApiActivedirectoriesAgentsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActiveDirectoryApiActivedirectoriesAgentsDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActivedirectoriesAgentsGet**
> ActiveDirectoryAgentListOutput ActivedirectoriesAgentsGet(ctx, activedirectoryId, agentId, optional)
Get Active Directory Agent

This endpoint returns an Active Directory agent.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/activedirectories/{activedirectory_id}/agents/{agent_id} \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activedirectoryId** | **string**|  | 
  **agentId** | **string**|  | 
 **optional** | ***ActiveDirectoryApiActivedirectoriesAgentsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActiveDirectoryApiActivedirectoriesAgentsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

[**ActiveDirectoryAgentListOutput**](active-directory-agent-list-output.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActivedirectoriesAgentsList**
> []ActiveDirectoryAgentListOutput ActivedirectoriesAgentsList(ctx, activedirectoryId, optional)
List Active Directory Agents

This endpoint allows you to list all your Active Directory Agents for a given Instance.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/activedirectories/{activedirectory_id}/agents \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}'   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activedirectoryId** | **string**|  | 
 **optional** | ***ActiveDirectoryApiActivedirectoriesAgentsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActiveDirectoryApiActivedirectoriesAgentsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **optional.Int32**| The offset into the records to return. | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 
 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

[**[]ActiveDirectoryAgentListOutput**](active-directory-agent-list-output.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActivedirectoriesAgentsPost**
> ActiveDirectoryAgentGetOutput ActivedirectoriesAgentsPost(ctx, activedirectoryId, optional)
Create a new Active Directory Agent

This endpoint allows you to create a new Active Directory Agent.   #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/activedirectories/{activedirectory_id}/agents \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activedirectoryId** | **string**|  | 
 **optional** | ***ActiveDirectoryApiActivedirectoriesAgentsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActiveDirectoryApiActivedirectoriesAgentsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ActiveDirectoryAgentInput**](ActiveDirectoryAgentInput.md)|  | 
 **xOrgId** | **optional.**| Organization identifier that can be obtained from console settings. | 

### Return type

[**ActiveDirectoryAgentGetOutput**](active-directory-agent-get-output.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActivedirectoriesDelete**
> ActiveDirectoryOutput ActivedirectoriesDelete(ctx, id, optional)
Delete an Active Directory

This endpoint allows you to delete an Active Directory Instance.  #### Sample Request ``` curl -X DELETE https://console.jumpcloud.com/api/v2/activedirectories/{ActiveDirectory_ID} \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}'   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ObjectID of this Active Directory instance. | 
 **optional** | ***ActiveDirectoryApiActivedirectoriesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActiveDirectoryApiActivedirectoriesDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

[**ActiveDirectoryOutput**](active-directory-output.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActivedirectoriesGet**
> ActiveDirectoryOutput ActivedirectoriesGet(ctx, id, optional)
Get an Active Directory

This endpoint returns a specific Active Directory.  #### Sample Request  ``` curl -X GET https://console.jumpcloud.com/api/v2/activedirectories/{ActiveDirectory_ID} \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}'   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ObjectID of this Active Directory instance. | 
 **optional** | ***ActiveDirectoryApiActivedirectoriesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActiveDirectoryApiActivedirectoriesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

[**ActiveDirectoryOutput**](active-directory-output.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActivedirectoriesList**
> []ActiveDirectoryOutput ActivedirectoriesList(ctx, optional)
List Active Directories

This endpoint allows you to list all your Active Directory Instances.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/activedirectories/ \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}'   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ActiveDirectoryApiActivedirectoriesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActiveDirectoryApiActivedirectoriesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | 
 **filter** | [**optional.Interface of []string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **limit** | **optional.Int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **optional.Int32**| The offset into the records to return. | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 
 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

[**[]ActiveDirectoryOutput**](active-directory-output.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActivedirectoriesPost**
> ActiveDirectoryOutput ActivedirectoriesPost(ctx, optional)
Create a new Active Directory

This endpoint allows you to create a new Active Directory.   #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/activedirectories/ \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{     \"domain\": \"{DC=AD_domain_name;DC=com}\"   }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ActiveDirectoryApiActivedirectoriesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActiveDirectoryApiActivedirectoriesPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ActiveDirectoryInput**](ActiveDirectoryInput.md)|  | 
 **xOrgId** | **optional.**| Organization identifier that can be obtained from console settings. | 

### Return type

[**ActiveDirectoryOutput**](active-directory-output.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphActiveDirectoryAssociationsList**
> []GraphConnection GraphActiveDirectoryAssociationsList(ctx, activedirectoryId, targets, optional)
List the associations of an Active Directory instance

This endpoint returns the direct associations of this Active Directory instance.  A direct association can be a non-homogeneous relationship between 2 different objects, for example Active Directory and Users.   #### Sample Request ``` curl -X GET 'https://console.jumpcloud.com/api/v2/activedirectories/{ActiveDirectory_ID}/associations?targets=user \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activedirectoryId** | **string**|  | 
  **targets** | [**[]string**](string.md)| Targets which a \&quot;active_directory\&quot; can be associated to. | 
 **optional** | ***ActiveDirectoryApiGraphActiveDirectoryAssociationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActiveDirectoryApiGraphActiveDirectoryAssociationsListOpts struct
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

# **GraphActiveDirectoryAssociationsPost**
> GraphActiveDirectoryAssociationsPost(ctx, activedirectoryId, optional)
Manage the associations of an Active Directory instance

This endpoint allows you to manage the _direct_ associations of an Active Directory instance.  A direct association can be a non-homogeneous relationship between 2 different objects, for example Active Directory and Users.  #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/activedirectories/{AD_Instance_ID}/associations \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{     \"op\": \"add\",     \"type\": \"user\",     \"id\": \"{User_ID}\"   }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activedirectoryId** | **string**|  | 
 **optional** | ***ActiveDirectoryApiGraphActiveDirectoryAssociationsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActiveDirectoryApiGraphActiveDirectoryAssociationsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GraphOperationActiveDirectory**](GraphOperationActiveDirectory.md)|  | 
 **xOrgId** | **optional.**| Organization identifier that can be obtained from console settings. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphActiveDirectoryTraverseUser**
> []GraphObjectWithPaths GraphActiveDirectoryTraverseUser(ctx, activedirectoryId, optional)
List the Users bound to an Active Directory instance

This endpoint will return all Users bound to an Active Directory instance, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this Active Directory instance to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this Active Directory instance.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/activedirectories/{ActiveDirectory_ID}/users \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activedirectoryId** | **string**| ObjectID of the Active Directory instance. | 
 **optional** | ***ActiveDirectoryApiGraphActiveDirectoryTraverseUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActiveDirectoryApiGraphActiveDirectoryTraverseUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | [**optional.Interface of []string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **limit** | **optional.Int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 
 **skip** | **optional.Int32**| The offset into the records to return. | [default to 0]

### Return type

[**[]GraphObjectWithPaths**](GraphObjectWithPaths.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphActiveDirectoryTraverseUserGroup**
> []GraphObjectWithPaths GraphActiveDirectoryTraverseUserGroup(ctx, activedirectoryId, optional)
List the User Groups bound to an Active Directory instance

This endpoint will return all Users Groups bound to an Active Directory instance, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the group's type, id, attributes and paths.  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this Active Directory instance to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this Active Directory instance.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/activedirectories/{ActiveDirectory_ID}/usergroups \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activedirectoryId** | **string**| ObjectID of the Active Directory instance. | 
 **optional** | ***ActiveDirectoryApiGraphActiveDirectoryTraverseUserGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActiveDirectoryApiGraphActiveDirectoryTraverseUserGroupOpts struct
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

