# {{classname}}

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportUsers**](SCIMImportApi.md#ImportUsers) | **Get** /applications/{application_id}/import/users | Get a list of users to import from an Application IdM service provider

# **ImportUsers**
> ImportUsersResponse ImportUsers(ctx, applicationId, optional)
Get a list of users to import from an Application IdM service provider

Get a list of users to import from an Application IdM service provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**| ObjectID of the Application. | 
 **optional** | ***SCIMImportApiImportUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SCIMImportApiImportUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| Filter users by a search term | 
 **query** | **optional.String**| URL query to merge with the service provider request | 
 **sort** | **optional.String**| Sort users by supported fields | 
 **sortOrder** | **optional.String**|  | [default to asc]
 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 
 **limit** | **optional.Int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **optional.Int32**| The offset into the records to return. | [default to 0]

### Return type

[**ImportUsersResponse**](importUsersResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

