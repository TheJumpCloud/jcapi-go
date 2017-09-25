# \GSuiteApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphGSuiteAssociationsList**](GSuiteApi.md#GraphGSuiteAssociationsList) | **Get** /gsuites/{gsuite_id}/associations | List the associations of a G Suite instance
[**GraphGSuiteAssociationsPost**](GSuiteApi.md#GraphGSuiteAssociationsPost) | **Post** /gsuites/{gsuite_id}/associations | Manage the associations of a G Suite instance
[**GraphGSuiteTraverseUser**](GSuiteApi.md#GraphGSuiteTraverseUser) | **Get** /gsuites/{gsuite_id}/users | List the Users associated with a G Suite instance
[**GraphGSuiteTraverseUserGroup**](GSuiteApi.md#GraphGSuiteTraverseUserGroup) | **Get** /gsuites/{gsuite_id}/usergroups | List the User Groups associated with a G Suite instance


# **GraphGSuiteAssociationsList**
> []GraphConnection GraphGSuiteAssociationsList($gsuiteId, $targets, $contentType, $accept, $limit, $skip)

List the associations of a G Suite instance

This endpoint returns the _direct_ associations of this G Suite instance.  A direct association can be a non-homogenous relationship between 2 different objects. for example G Suite and Users.   #### Sample Request ``` https://console.jumpcloud.com/api/v2/gsuites/{gsuite_id}/associations?targets=user ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gsuiteId** | **string**| ObjectID of the G Suite instance. | 
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

# **GraphGSuiteAssociationsPost**
> GraphGSuiteAssociationsPost($gsuiteId, $body)

Manage the associations of a G Suite instance

This endpoint returns the _direct_ associations of this G Suite instance.  A direct association can be a non-homogenous relationship between 2 different objects. for example G Suite and Users.   #### Sample Request ``` https://console.jumpcloud.com/api/v2/gsuites/{gsuite_id}/associations ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gsuiteId** | **string**| ObjectID of the G Suite instance. | 
 **body** | [**GraphManagementReq**](GraphManagementReq.md)|  | [optional] 

### Return type

void (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphGSuiteTraverseUser**
> []GraphObjectWithPaths GraphGSuiteTraverseUser($gsuiteId, $contentType, $accept, $limit, $skip)

List the Users associated with a G Suite instance

This endpoint will return Users associated with a G Suite instance. Each element will contain the type, id, attributes and paths.  The `attributes` object is a key/value hash of attributes specifically set for this group.  The `paths` array enumerates each path from this G Suite instance to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this G Suite instance.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` https://console.jumpcloud.com/api/v2/gsuites/{gsuite_id}/users ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gsuiteId** | **string**| ObjectID of the G Suite instance. | 
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

# **GraphGSuiteTraverseUserGroup**
> []GraphObjectWithPaths GraphGSuiteTraverseUserGroup($gsuiteId, $contentType, $accept, $limit, $skip)

List the User Groups associated with a G Suite instance

This endpoint will return User Groups associated with a G Suite instance. Each element will contain the group's type, id, attributes and paths.  The `attributes` object is a key/value hash of attributes specifically set for this group.  The `paths` array enumerates each path from this G Suite instance to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this G Suite instance.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` https://console.jumpcloud.com/api/v2/gsuites/{gsuite_id}/usersgroups ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gsuiteId** | **string**| ObjectID of the G Suite instance. | 
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

