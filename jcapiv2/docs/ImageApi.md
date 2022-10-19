# {{classname}}

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationsDeleteLogo**](ImageApi.md#ApplicationsDeleteLogo) | **Delete** /applications/{application_id}/logo | Delete application image

# **ApplicationsDeleteLogo**
> ApplicationsDeleteLogo(ctx, applicationId, optional)
Delete application image

Deletes the specified image from an application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**|  | 
 **optional** | ***ImageApiApplicationsDeleteLogoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImageApiApplicationsDeleteLogoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

