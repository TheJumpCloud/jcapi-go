# {{classname}}

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DuoAccountDelete**](DuoApi.md#DuoAccountDelete) | **Delete** /duo/accounts/{id} | Delete a Duo Account
[**DuoAccountGet**](DuoApi.md#DuoAccountGet) | **Get** /duo/accounts/{id} | Get a Duo Acount
[**DuoAccountList**](DuoApi.md#DuoAccountList) | **Get** /duo/accounts | List Duo Accounts
[**DuoAccountPost**](DuoApi.md#DuoAccountPost) | **Post** /duo/accounts | Create Duo Account
[**DuoApplicationDelete**](DuoApi.md#DuoApplicationDelete) | **Delete** /duo/accounts/{account_id}/applications/{application_id} | Delete a Duo Application
[**DuoApplicationGet**](DuoApi.md#DuoApplicationGet) | **Get** /duo/accounts/{account_id}/applications/{application_id} | Get a Duo application
[**DuoApplicationList**](DuoApi.md#DuoApplicationList) | **Get** /duo/accounts/{account_id}/applications | List Duo Applications
[**DuoApplicationPost**](DuoApi.md#DuoApplicationPost) | **Post** /duo/accounts/{account_id}/applications | Create Duo Application
[**DuoApplicationUpdate**](DuoApi.md#DuoApplicationUpdate) | **Put** /duo/accounts/{account_id}/applications/{application_id} | Update Duo Application

# **DuoAccountDelete**
> DuoAccount DuoAccountDelete(ctx, id, optional)
Delete a Duo Account

Removes the specified Duo account, an error will be returned if the account has some Duo application used in a protected resource.  #### Sample Request ``` curl -X DELETE https://console.jumpcloud.com/api/v2/duo/accounts/{id} \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ObjectID of the Duo Account | 
 **optional** | ***DuoApiDuoAccountDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DuoApiDuoAccountDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

[**DuoAccount**](DuoAccount.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DuoAccountGet**
> DuoAccount DuoAccountGet(ctx, id, optional)
Get a Duo Acount

This endpoint returns a specific Duo account.  #### Sample Request ``` curl https://console.jumpcloud.com/api/v2/duo/accounts/{id} \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ObjectID of the Duo Account | 
 **optional** | ***DuoApiDuoAccountGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DuoApiDuoAccountGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

[**DuoAccount**](DuoAccount.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DuoAccountList**
> []DuoAccount DuoAccountList(ctx, optional)
List Duo Accounts

This endpoint returns all the Duo accounts for your organization. Note: There can currently only be one Duo account for your organization.  #### Sample Request ``` curl https://console.jumpcloud.com/api/v2/duo/accounts \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DuoApiDuoAccountListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DuoApiDuoAccountListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

[**[]DuoAccount**](DuoAccount.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DuoAccountPost**
> DuoAccount DuoAccountPost(ctx, optional)
Create Duo Account

Registers a Duo account for an organization. Only one Duo account will be allowed, in case an organization has a Duo account already a 409 (Conflict) code will be returned.  #### Sample Request ```   curl -X POST https://console.jumpcloud.com/api/v2/duo/accounts \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DuoApiDuoAccountPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DuoApiDuoAccountPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

[**DuoAccount**](DuoAccount.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DuoApplicationDelete**
> DuoApplication DuoApplicationDelete(ctx, accountId, applicationId, optional)
Delete a Duo Application

Deletes the specified Duo application, an error will be returned if the application is used in a protected resource.  #### Sample Request ```   curl -X DELETE https://console.jumpcloud.com/api/v2/duo/accounts/{ACCOUNT_ID}/applications/{APPLICATION_ID} \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}'' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
  **applicationId** | **string**|  | 
 **optional** | ***DuoApiDuoApplicationDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DuoApiDuoApplicationDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

[**DuoApplication**](DuoApplication.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DuoApplicationGet**
> DuoApplication DuoApplicationGet(ctx, accountId, applicationId, optional)
Get a Duo application

This endpoint returns a specific Duo application that is associated with the specified Duo account.  #### Sample Request ```   curl https://console.jumpcloud.com/api/v2/duo/accounts/{ACCOUNT_ID}/applications/{APPLICATION_ID} \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
  **applicationId** | **string**|  | 
 **optional** | ***DuoApiDuoApplicationGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DuoApiDuoApplicationGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

[**DuoApplication**](DuoApplication.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DuoApplicationList**
> []DuoApplication DuoApplicationList(ctx, accountId, optional)
List Duo Applications

This endpoint returns all the Duo applications for the specified Duo account. Note: There can currently only be one Duo application for your organization.  #### Sample Request ```   curl https://console.jumpcloud.com/api/v2/duo/accounts/{ACCOUNT_ID}/applications \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
 **optional** | ***DuoApiDuoApplicationListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DuoApiDuoApplicationListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

[**[]DuoApplication**](DuoApplication.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DuoApplicationPost**
> DuoApplication DuoApplicationPost(ctx, accountId, optional)
Create Duo Application

Creates a Duo application for your organization and the specified account.  #### Sample Request ```   curl -X POST https://console.jumpcloud.com/api/v2/duo/accounts/{ACCOUNT_ID}/applications \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{     \"name\": \"Application Name\",     \"apiHost\": \"api-1234.duosecurity.com\",     \"integrationKey\": \"1234\",     \"secretKey\": \"5678\"   }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
 **optional** | ***DuoApiDuoApplicationPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DuoApiDuoApplicationPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DuoApplicationReq**](DuoApplicationReq.md)|  | 
 **xOrgId** | **optional.**| Organization identifier that can be obtained from console settings. | 

### Return type

[**DuoApplication**](DuoApplication.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DuoApplicationUpdate**
> DuoApplication DuoApplicationUpdate(ctx, accountId, applicationId, optional)
Update Duo Application

Updates the specified Duo application.  #### Sample Request ```   curl -X PUT https://console.jumpcloud.com/api/v2/duo/accounts/{ACCOUNT_ID}/applications/{APPLICATION_ID} \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{     \"name\": \"Application Name\",     \"apiHost\": \"api-1234.duosecurity.com\",     \"integrationKey\": \"1234\",     \"secretKey\": \"5678\"   }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
  **applicationId** | **string**|  | 
 **optional** | ***DuoApiDuoApplicationUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DuoApiDuoApplicationUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DuoApplicationUpdateReq**](DuoApplicationUpdateReq.md)|  | 
 **xOrgId** | **optional.**| Organization identifier that can be obtained from console settings. | 

### Return type

[**DuoApplication**](DuoApplication.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

