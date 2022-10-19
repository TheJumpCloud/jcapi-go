# {{classname}}

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomEmailsCreate**](CustomEmailsApi.md#CustomEmailsCreate) | **Post** /customemails | Create custom email configuration
[**CustomEmailsDestroy**](CustomEmailsApi.md#CustomEmailsDestroy) | **Delete** /customemails/{custom_email_type} | Delete custom email configuration
[**CustomEmailsGetTemplates**](CustomEmailsApi.md#CustomEmailsGetTemplates) | **Get** /customemail/templates | List custom email templates
[**CustomEmailsRead**](CustomEmailsApi.md#CustomEmailsRead) | **Get** /customemails/{custom_email_type} | Get custom email configuration
[**CustomEmailsUpdate**](CustomEmailsApi.md#CustomEmailsUpdate) | **Put** /customemails/{custom_email_type} | Update custom email configuration

# **CustomEmailsCreate**
> CustomEmail CustomEmailsCreate(ctx, optional)
Create custom email configuration

Create the custom email configuration for the specified custom email type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomEmailsApiCustomEmailsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomEmailsApiCustomEmailsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CustomEmail**](CustomEmail.md)|  | 
 **xOrgId** | **optional.**| Organization identifier that can be obtained from console settings. | 

### Return type

[**CustomEmail**](CustomEmail.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomEmailsDestroy**
> CustomEmailsDestroy(ctx, customEmailType, optional)
Delete custom email configuration

Delete the custom email configuration for the specified custom email type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customEmailType** | **string**|  | 
 **optional** | ***CustomEmailsApiCustomEmailsDestroyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomEmailsApiCustomEmailsDestroyOpts struct
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

# **CustomEmailsGetTemplates**
> []CustomEmailTemplate CustomEmailsGetTemplates(ctx, )
List custom email templates

Get the list of custom email templates

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]CustomEmailTemplate**](CustomEmailTemplate.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomEmailsRead**
> CustomEmail CustomEmailsRead(ctx, customEmailType, optional)
Get custom email configuration

Get the custom email configuration for the specified custom email type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customEmailType** | **string**|  | 
 **optional** | ***CustomEmailsApiCustomEmailsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomEmailsApiCustomEmailsReadOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

[**CustomEmail**](CustomEmail.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomEmailsUpdate**
> CustomEmail CustomEmailsUpdate(ctx, customEmailType, optional)
Update custom email configuration

Update the custom email configuration for the specified custom email type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customEmailType** | **string**|  | 
 **optional** | ***CustomEmailsApiCustomEmailsUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomEmailsApiCustomEmailsUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CustomEmail**](CustomEmail.md)|  | 
 **xOrgId** | **optional.**| Organization identifier that can be obtained from console settings. | 

### Return type

[**CustomEmail**](CustomEmail.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

