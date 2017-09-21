# \PoliciesApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphPolicyAssociationsList**](PoliciesApi.md#GraphPolicyAssociationsList) | **Get** /policies/{policy_id}/associations | List the associations of a Policy
[**GraphPolicyAssociationsPost**](PoliciesApi.md#GraphPolicyAssociationsPost) | **Post** /policies/{policy_id}/associations | Manage the associations of a Policy
[**GraphPolicyTraverseSystem**](PoliciesApi.md#GraphPolicyTraverseSystem) | **Get** /policies/{policy_id}/systems | List the Systems associated with a Policy
[**GraphPolicyTraverseSystemGroup**](PoliciesApi.md#GraphPolicyTraverseSystemGroup) | **Get** /policies/{policy_id}/systemgroups | List the System Groups associated with a Policy
[**PoliciesDelete**](PoliciesApi.md#PoliciesDelete) | **Delete** /policies/{id} | Deletes a Policy
[**PoliciesGet**](PoliciesApi.md#PoliciesGet) | **Get** /policies/{id} | Gets a specific Policy.
[**PoliciesList**](PoliciesApi.md#PoliciesList) | **Get** /policies | Lists all the Policies
[**PoliciesPost**](PoliciesApi.md#PoliciesPost) | **Post** /policies | Create a new Policy
[**PoliciesPut**](PoliciesApi.md#PoliciesPut) | **Put** /policies/{id} | Update an existing Policy
[**PolicyresultsGet**](PoliciesApi.md#PolicyresultsGet) | **Get** /policyresults/{id} | Get a specific Policy Result.
[**PolicyresultsList**](PoliciesApi.md#PolicyresultsList) | **Get** /policies/{policy_id}/policyresults | Lists all the policy results of a given policy.
[**PolicyresultsList_0**](PoliciesApi.md#PolicyresultsList_0) | **Get** /policyresults | Lists all the policy results for an organization.
[**PolicytemplatesGet**](PoliciesApi.md#PolicytemplatesGet) | **Get** /policytemplates/{id} | Get a specific Policy Template
[**PolicytemplatesList**](PoliciesApi.md#PolicytemplatesList) | **Get** /policytemplates | Lists all of the Policy Templates


# **GraphPolicyAssociationsList**
> []GraphConnection GraphPolicyAssociationsList($policyId, $targets, $contentType, $accept, $limit, $skip)

List the associations of a Policy

This endpoint returns the _direct_ associations of a Policy.  A direct association can be a non-homogenous relationship between 2 different objects. for example Policies and Systems.  #### Sample Request ``` https://console.jumpcloud.com/api/v2/policies/{policy_id}/associations?targets=system ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyId** | **string**| ObjectID of the Policy. | 
 **targets** | [**[]string**](string.md)|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **limit** | **int32**| The number of records to return at once. | [optional] [default to 10]
 **skip** | **int32**| The offset into the records to return. | [optional] [default to 0]

### Return type

[**[]GraphConnection**](GraphConnection.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphPolicyAssociationsPost**
> GraphPolicyAssociationsPost($policyId, $contentType, $accept, $body)

Manage the associations of a Policy

This endpoint allows you to manage the _direct_ associations of a Policy.  A direct association can be a non-homogenous relationship between 2 different objects. for example Policies and Systems.  #### Sample Request ``` https://console.jumpcloud.com/api/v2/policies/{policy_id}/associations ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyId** | **string**| ObjectID of the Policy. | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **body** | [**GraphManagementReq**](GraphManagementReq.md)|  | [optional] 

### Return type

void (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphPolicyTraverseSystem**
> []GraphObjectWithPaths GraphPolicyTraverseSystem($policyId, $contentType, $accept, $limit, $skip)

List the Systems associated with a Policy

This endpoint will return Systems associated with a Policy. Each element will contain the type, id, attributes and paths.  The `attributes` object is a key/value hash of attributes specifically set for this group.  The `paths` array enumerates each path from this Policy to the corresponding System; this array represents all grouping and/or associations that would have to be removed to deprovision the System from this Policy.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` https://console.jumpcloud.com/api/v2/policies/{policy_id}/systems ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyId** | **string**| ObjectID of the Command. | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **limit** | **int32**| The number of records to return at once. | [optional] [default to 10]
 **skip** | **int32**| The offset into the records to return. | [optional] [default to 0]

### Return type

[**[]GraphObjectWithPaths**](GraphObjectWithPaths.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphPolicyTraverseSystemGroup**
> []GraphObjectWithPaths GraphPolicyTraverseSystemGroup($policyId, $contentType, $accept, $limit, $skip)

List the System Groups associated with a Policy

This endpoint will return System Groups associated with a Policy. Each element will contain the group's type, id, attributes and paths.  The `attributes` object is a key/value hash of attributes specifically set for this group.  The `paths` array enumerates each path from this Policy to the corresponding System Group; this array represents all grouping and/or associations that would have to be removed to deprovision the System Group from this Policy.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` https://console.jumpcloud.com/api/v2/policies/{policy_id}/systemsgroups ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyId** | **string**| ObjectID of the Command. | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **limit** | **int32**| The number of records to return at once. | [optional] [default to 10]
 **skip** | **int32**| The offset into the records to return. | [optional] [default to 0]

### Return type

[**[]GraphObjectWithPaths**](GraphObjectWithPaths.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoliciesDelete**
> PoliciesDelete($id, $contentType, $accept)

Deletes a Policy

This endpoint allows you to delete a policy.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ObjectID of the Policy object. | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]

### Return type

void (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoliciesGet**
> PolicyWithDetails PoliciesGet($id, $contentType, $accept)

Gets a specific Policy.

This endpoint returns a specific policy.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ObjectID of the Policy object. | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]

### Return type

[**PolicyWithDetails**](PolicyWithDetails.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoliciesList**
> []Policy PoliciesList($contentType, $accept, $fields, $filter, $limit, $skip, $sort)

Lists all the Policies

This endpoint returns all policies.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **fields** | **string**| The comma separated fields included in the returned records. If omitted the default list of fields will be returned.  | [optional] [default to ]
 **filter** | **string**| Supported operators are: eq, ne, gt, ge, lt, le, between, search | [optional] [default to ]
 **limit** | **int32**| The number of records to return at once. | [optional] [default to 10]
 **skip** | **int32**| The offset into the records to return. | [optional] [default to 0]
 **sort** | **string**| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [optional] [default to ]

### Return type

[**[]Policy**](Policy.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoliciesPost**
> PolicyWithDetails PoliciesPost($contentType, $accept, $body)

Create a new Policy

This endpoint allows you to create a policy.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **body** | [**PolicyRequest**](PolicyRequest.md)|  | [optional] 

### Return type

[**PolicyWithDetails**](PolicyWithDetails.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoliciesPut**
> Policy PoliciesPut($id, $body)

Update an existing Policy

This endpoint allows you to update a policy.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ObjectID of the Policy object. | 
 **body** | [**PolicyRequest**](PolicyRequest.md)|  | [optional] 

### Return type

[**Policy**](Policy.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PolicyresultsGet**
> PolicyResult PolicyresultsGet($id, $contentType, $accept)

Get a specific Policy Result.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ObjectID of the Policy Result. | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]

### Return type

[**PolicyResult**](PolicyResult.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PolicyresultsList**
> []PolicyResult PolicyresultsList($policyId, $contentType, $accept, $fields, $filter, $limit, $skip, $sort, $aggregate)

Lists all the policy results of a given policy.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyId** | **string**|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **fields** | **string**| The comma separated fields included in the returned records. If omitted the default list of fields will be returned.  | [optional] [default to ]
 **filter** | **string**| Supported operators are: eq, ne, gt, ge, lt, le, between, search | [optional] [default to ]
 **limit** | **int32**| The number of records to return at once. | [optional] [default to 10]
 **skip** | **int32**| The offset into the records to return. | [optional] [default to 0]
 **sort** | **string**| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [optional] [default to ]
 **aggregate** | **string**|  | [optional] [default to ]

### Return type

[**[]PolicyResult**](PolicyResult.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PolicyresultsList_0**
> []PolicyResult PolicyresultsList_0($contentType, $accept, $fields, $filter, $limit, $skip, $sort, $aggregate)

Lists all the policy results for an organization.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **fields** | **string**| The comma separated fields included in the returned records. If omitted the default list of fields will be returned.  | [optional] [default to ]
 **filter** | **string**| Supported operators are: eq, ne, gt, ge, lt, le, between, search | [optional] [default to ]
 **limit** | **int32**| The number of records to return at once. | [optional] [default to 10]
 **skip** | **int32**| The offset into the records to return. | [optional] [default to 0]
 **sort** | **string**| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [optional] [default to ]
 **aggregate** | **string**|  | [optional] [default to ]

### Return type

[**[]PolicyResult**](PolicyResult.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PolicytemplatesGet**
> PolicyTemplateWithDetails PolicytemplatesGet($id, $contentType, $accept)

Get a specific Policy Template

This endpoint returns a specific policy template.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ObjectID of the Policy Template. | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]

### Return type

[**PolicyTemplateWithDetails**](PolicyTemplateWithDetails.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PolicytemplatesList**
> []PolicyTemplate PolicytemplatesList($contentType, $accept, $fields, $filter, $limit, $skip, $sort)

Lists all of the Policy Templates

This endpoint returns all policy templates.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **fields** | **string**| The comma separated fields included in the returned records. If omitted the default list of fields will be returned.  | [optional] [default to ]
 **filter** | **string**| Supported operators are: eq, ne, gt, ge, lt, le, between, search | [optional] [default to ]
 **limit** | **int32**| The number of records to return at once. | [optional] [default to 10]
 **skip** | **int32**| The offset into the records to return. | [optional] [default to 0]
 **sort** | **string**| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [optional] [default to ]

### Return type

[**[]PolicyTemplate**](PolicyTemplate.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

