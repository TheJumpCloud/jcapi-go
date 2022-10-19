# {{classname}}

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphPolicyGroupAssociationsList**](PolicyGroupsApi.md#GraphPolicyGroupAssociationsList) | **Get** /policygroups/{group_id}/associations | List the associations of a Policy Group.
[**GraphPolicyGroupAssociationsPost**](PolicyGroupsApi.md#GraphPolicyGroupAssociationsPost) | **Post** /policygroups/{group_id}/associations | Manage the associations of a Policy Group
[**GraphPolicyGroupMembersList**](PolicyGroupsApi.md#GraphPolicyGroupMembersList) | **Get** /policygroups/{group_id}/members | List the members of a Policy Group
[**GraphPolicyGroupMembersPost**](PolicyGroupsApi.md#GraphPolicyGroupMembersPost) | **Post** /policygroups/{group_id}/members | Manage the members of a Policy Group
[**GraphPolicyGroupMembership**](PolicyGroupsApi.md#GraphPolicyGroupMembership) | **Get** /policygroups/{group_id}/membership | List the Policy Group&#x27;s membership
[**GraphPolicyGroupTraverseSystem**](PolicyGroupsApi.md#GraphPolicyGroupTraverseSystem) | **Get** /policygroups/{group_id}/systems | List the Systems bound to a Policy Group
[**GraphPolicyGroupTraverseSystemGroup**](PolicyGroupsApi.md#GraphPolicyGroupTraverseSystemGroup) | **Get** /policygroups/{group_id}/systemgroups | List the System Groups bound to Policy Groups
[**GroupsPolicyDelete**](PolicyGroupsApi.md#GroupsPolicyDelete) | **Delete** /policygroups/{id} | Delete a Policy Group
[**GroupsPolicyGet**](PolicyGroupsApi.md#GroupsPolicyGet) | **Get** /policygroups/{id} | View an individual Policy Group details
[**GroupsPolicyList**](PolicyGroupsApi.md#GroupsPolicyList) | **Get** /policygroups | List all Policy Groups
[**GroupsPolicyPost**](PolicyGroupsApi.md#GroupsPolicyPost) | **Post** /policygroups | Create a new Policy Group
[**GroupsPolicyPut**](PolicyGroupsApi.md#GroupsPolicyPut) | **Put** /policygroups/{id} | Update a Policy Group

# **GraphPolicyGroupAssociationsList**
> []GraphConnection GraphPolicyGroupAssociationsList(ctx, groupId, targets, optional)
List the associations of a Policy Group.

This endpoint returns the _direct_ associations of this Policy Group.  A direct association can be a non-homogeneous relationship between 2 different objects, for example Policy Groups and Policies.   #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/policygroups/{GroupID}/associations?targets=system \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**| ObjectID of the Policy Group. | 
  **targets** | [**[]string**](string.md)| Targets which a \&quot;policy_group\&quot; can be associated to. | 
 **optional** | ***PolicyGroupsApiGraphPolicyGroupAssociationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyGroupsApiGraphPolicyGroupAssociationsListOpts struct
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

# **GraphPolicyGroupAssociationsPost**
> GraphPolicyGroupAssociationsPost(ctx, groupId, optional)
Manage the associations of a Policy Group

This endpoint manages the _direct_ associations of this Policy Group.  A direct association can be a non-homogeneous relationship between 2 different objects, for example Policy Groups and Policies.   #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/policygroups/{GroupID}/associations \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{     \"op\": \"add\",     \"type\": \"system\",     \"id\": \"{SystemID}\"   }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**| ObjectID of the Policy Group. | 
 **optional** | ***PolicyGroupsApiGraphPolicyGroupAssociationsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyGroupsApiGraphPolicyGroupAssociationsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GraphOperationPolicyGroup**](GraphOperationPolicyGroup.md)|  | 
 **xOrgId** | **optional.**| Organization identifier that can be obtained from console settings. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphPolicyGroupMembersList**
> []GraphConnection GraphPolicyGroupMembersList(ctx, groupId, optional)
List the members of a Policy Group

This endpoint returns the Policy members of a Policy Group.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/policygroups/{GroupID}/members \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**| ObjectID of the Policy Group. | 
 **optional** | ***PolicyGroupsApiGraphPolicyGroupMembersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyGroupsApiGraphPolicyGroupMembersListOpts struct
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

# **GraphPolicyGroupMembersPost**
> GraphPolicyGroupMembersPost(ctx, groupId, optional)
Manage the members of a Policy Group

This endpoint allows you to manage the Policy members of a Policy Group.  #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/policygroups/{GroupID}/members \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{     \"op\": \"add\",     \"type\": \"policy\",     \"id\": \"{Policy_ID}\"   }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**| ObjectID of the Policy Group. | 
 **optional** | ***PolicyGroupsApiGraphPolicyGroupMembersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyGroupsApiGraphPolicyGroupMembersPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GraphOperationPolicyGroupMember**](GraphOperationPolicyGroupMember.md)|  | 
 **xOrgId** | **optional.**| Organization identifier that can be obtained from console settings. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphPolicyGroupMembership**
> []GraphObjectWithPaths GraphPolicyGroupMembership(ctx, groupId, optional)
List the Policy Group's membership

This endpoint returns all Policy members that are a member of this Policy Group.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/policygroups/{GroupID}/membership \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**| ObjectID of the Policy Group. | 
 **optional** | ***PolicyGroupsApiGraphPolicyGroupMembershipOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyGroupsApiGraphPolicyGroupMembershipOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | [**optional.Interface of []string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **limit** | **optional.Int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **optional.Int32**| The offset into the records to return. | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 
 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

[**[]GraphObjectWithPaths**](GraphObjectWithPaths.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphPolicyGroupTraverseSystem**
> []GraphObjectWithPaths GraphPolicyGroupTraverseSystem(ctx, groupId, optional)
List the Systems bound to a Policy Group

This endpoint will return all Systems bound to a Policy Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this Policy Group to the corresponding System; this array represents all grouping and/or associations that would have to be removed to deprovision the System from this Policy Group.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/policygroups/{GroupID}/systems \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**| ObjectID of the Policy Group. | 
 **optional** | ***PolicyGroupsApiGraphPolicyGroupTraverseSystemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyGroupsApiGraphPolicyGroupTraverseSystemOpts struct
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

# **GraphPolicyGroupTraverseSystemGroup**
> []GraphObjectWithPaths GraphPolicyGroupTraverseSystemGroup(ctx, groupId, optional)
List the System Groups bound to Policy Groups

This endpoint will return all System Groups bound to a Policy Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.  Each element will contain the type, id, attributes and paths.  The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.  The `paths` array enumerates each path from this Policy Group to the corresponding System Group; this array represents all grouping and/or associations that would have to be removed to deprovision the System Group from this Policy Group.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/policygroups/{GroupID}/systemgroups \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**| ObjectID of the Policy Group. | 
 **optional** | ***PolicyGroupsApiGraphPolicyGroupTraverseSystemGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyGroupsApiGraphPolicyGroupTraverseSystemGroupOpts struct
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

# **GroupsPolicyDelete**
> PolicyGroup GroupsPolicyDelete(ctx, id, optional)
Delete a Policy Group

This endpoint allows you to delete a Policy Group.  #### Sample Request ``` curl -X DELETE https://console.jumpcloud.com/api/v2/policygroups/{GroupID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'  ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ObjectID of the Policy Group. | 
 **optional** | ***PolicyGroupsApiGroupsPolicyDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyGroupsApiGroupsPolicyDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

[**PolicyGroup**](PolicyGroup.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsPolicyGet**
> PolicyGroup GroupsPolicyGet(ctx, id, optional)
View an individual Policy Group details

This endpoint returns the details of a Policy Group.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/policygroups/{GroupID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ObjectID of the Policy Group. | 
 **optional** | ***PolicyGroupsApiGroupsPolicyGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyGroupsApiGroupsPolicyGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

[**PolicyGroup**](PolicyGroup.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsPolicyList**
> []PolicyGroup GroupsPolicyList(ctx, optional)
List all Policy Groups

This endpoint returns all Policy Groups.  Available filter fields:   - `name`   - `disabled`   - `type`  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/policygroups \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyGroupsApiGroupsPolicyListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyGroupsApiGroupsPolicyListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | 
 **filter** | [**optional.Interface of []string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **limit** | **optional.Int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **optional.Int32**| The offset into the records to return. | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 
 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

[**[]PolicyGroup**](PolicyGroup.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsPolicyPost**
> PolicyGroup GroupsPolicyPost(ctx, optional)
Create a new Policy Group

This endpoint allows you to create a new Policy Group.  #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/policygroups \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{     \"name\": \"{Group_Name}\"   }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyGroupsApiGroupsPolicyPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyGroupsApiGroupsPolicyPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PolicyGroupData**](PolicyGroupData.md)|  | 
 **xOrgId** | **optional.**| Organization identifier that can be obtained from console settings. | 

### Return type

[**PolicyGroup**](PolicyGroup.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsPolicyPut**
> PolicyGroup GroupsPolicyPut(ctx, id, optional)
Update a Policy Group

This endpoint allows you to do a full update of the Policy Group.  #### Sample Request ``` curl -X PUT https://console.jumpcloud.com/api/v2/policygroups/{Group_ID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{     \"name\": \"group_update\"   }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ObjectID of the Policy Group. | 
 **optional** | ***PolicyGroupsApiGroupsPolicyPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyGroupsApiGroupsPolicyPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PolicyGroupData**](PolicyGroupData.md)|  | 
 **xOrgId** | **optional.**| Organization identifier that can be obtained from console settings. | 

### Return type

[**PolicyGroup**](PolicyGroup.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

