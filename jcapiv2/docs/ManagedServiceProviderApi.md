# {{classname}}

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdministratorOrganizationsCreateByAdministrator**](ManagedServiceProviderApi.md#AdministratorOrganizationsCreateByAdministrator) | **Post** /administrators/{id}/organizationlinks | Allow Adminstrator access to an Organization.
[**AdministratorOrganizationsListByAdministrator**](ManagedServiceProviderApi.md#AdministratorOrganizationsListByAdministrator) | **Get** /administrators/{id}/organizationlinks | List the association links between an Administrator and Organizations.
[**AdministratorOrganizationsListByOrganization**](ManagedServiceProviderApi.md#AdministratorOrganizationsListByOrganization) | **Get** /organizations/{id}/administratorlinks | List the association links between an Organization and Administrators.
[**AdministratorOrganizationsRemoveByAdministrator**](ManagedServiceProviderApi.md#AdministratorOrganizationsRemoveByAdministrator) | **Delete** /administrators/{administrator_id}/organizationlinks/{id} | Remove association between an Administrator and an Organization.
[**ProviderOrganizationsUpdateOrg**](ManagedServiceProviderApi.md#ProviderOrganizationsUpdateOrg) | **Put** /providers/{provider_id}/organizations/{id} | Update Provider Organization
[**ProvidersGetProvider**](ManagedServiceProviderApi.md#ProvidersGetProvider) | **Get** /providers/{provider_id} | Retrieve Provider
[**ProvidersListAdministrators**](ManagedServiceProviderApi.md#ProvidersListAdministrators) | **Get** /providers/{provider_id}/administrators | List Provider Administrators
[**ProvidersListOrganizations**](ManagedServiceProviderApi.md#ProvidersListOrganizations) | **Get** /providers/{provider_id}/organizations | List Provider Organizations
[**ProvidersPostAdmins**](ManagedServiceProviderApi.md#ProvidersPostAdmins) | **Post** /providers/{provider_id}/administrators | Create a new Provider Administrator
[**ProvidersRetrieveInvoice**](ManagedServiceProviderApi.md#ProvidersRetrieveInvoice) | **Get** /providers/{provider_id}/invoices/{ID} | Download a provider&#x27;s invoice.
[**ProvidersRetrieveInvoices**](ManagedServiceProviderApi.md#ProvidersRetrieveInvoices) | **Get** /providers/{provider_id}/invoices | List a provider&#x27;s invoices.

# **AdministratorOrganizationsCreateByAdministrator**
> AdministratorOrganizationLink AdministratorOrganizationsCreateByAdministrator(ctx, id, optional)
Allow Adminstrator access to an Organization.

This endpoint allows you to grant Administrator access to an Organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ManagedServiceProviderApiAdministratorOrganizationsCreateByAdministratorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagedServiceProviderApiAdministratorOrganizationsCreateByAdministratorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AdministratorOrganizationLinkReq**](AdministratorOrganizationLinkReq.md)|  | 

### Return type

[**AdministratorOrganizationLink**](AdministratorOrganizationLink.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdministratorOrganizationsListByAdministrator**
> []AdministratorOrganizationLink AdministratorOrganizationsListByAdministrator(ctx, id, optional)
List the association links between an Administrator and Organizations.

This endpoint returns the association links between an Administrator and Organizations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ManagedServiceProviderApiAdministratorOrganizationsListByAdministratorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagedServiceProviderApiAdministratorOrganizationsListByAdministratorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **optional.Int32**| The offset into the records to return. | [default to 0]

### Return type

[**[]AdministratorOrganizationLink**](AdministratorOrganizationLink.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdministratorOrganizationsListByOrganization**
> []AdministratorOrganizationLink AdministratorOrganizationsListByOrganization(ctx, id, optional)
List the association links between an Organization and Administrators.

This endpoint returns the association links between an Organization and Administrators.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ManagedServiceProviderApiAdministratorOrganizationsListByOrganizationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagedServiceProviderApiAdministratorOrganizationsListByOrganizationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **optional.Int32**| The offset into the records to return. | [default to 0]

### Return type

[**[]AdministratorOrganizationLink**](AdministratorOrganizationLink.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdministratorOrganizationsRemoveByAdministrator**
> AdministratorOrganizationsRemoveByAdministrator(ctx, administratorId, id)
Remove association between an Administrator and an Organization.

This endpoint removes the association link between an Administrator and an Organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **administratorId** | **string**|  | 
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProviderOrganizationsUpdateOrg**
> Organization ProviderOrganizationsUpdateOrg(ctx, providerId, id, optional)
Update Provider Organization

This endpoint updates a provider's organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**|  | 
  **id** | **string**|  | 
 **optional** | ***ManagedServiceProviderApiProviderOrganizationsUpdateOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagedServiceProviderApiProviderOrganizationsUpdateOrgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of Organization**](Organization.md)|  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProvidersGetProvider**
> Provider ProvidersGetProvider(ctx, providerId, optional)
Retrieve Provider

This endpoint returns details about a provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**|  | 
 **optional** | ***ManagedServiceProviderApiProvidersGetProviderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagedServiceProviderApiProvidersGetProviderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | 

### Return type

[**Provider**](Provider.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProvidersListAdministrators**
> InlineResponse20012 ProvidersListAdministrators(ctx, providerId, optional)
List Provider Administrators

This endpoint returns a list of the Administrators associated with the Provider. You must be associated with the provider to use this route.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**|  | 
 **optional** | ***ManagedServiceProviderApiProvidersListAdministratorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagedServiceProviderApiProvidersListAdministratorsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | 
 **filter** | [**optional.Interface of []string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **limit** | **optional.Int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **optional.Int32**| The offset into the records to return. | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProvidersListOrganizations**
> InlineResponse20013 ProvidersListOrganizations(ctx, providerId, optional)
List Provider Organizations

This endpoint returns a list of the Organizations associated with the Provider. You must be associated with the provider to use this route.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**|  | 
 **optional** | ***ManagedServiceProviderApiProvidersListOrganizationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagedServiceProviderApiProvidersListOrganizationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | 
 **filter** | [**optional.Interface of []string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **limit** | **optional.Int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **optional.Int32**| The offset into the records to return. | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProvidersPostAdmins**
> Administrator ProvidersPostAdmins(ctx, providerId, optional)
Create a new Provider Administrator

This endpoint allows you to create a provider administrator. You must be associated with the provider to use this route. You must provide either `role` or `roleName`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**|  | 
 **optional** | ***ManagedServiceProviderApiProvidersPostAdminsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagedServiceProviderApiProvidersPostAdminsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ProviderAdminReq**](ProviderAdminReq.md)|  | 

### Return type

[**Administrator**](Administrator.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProvidersRetrieveInvoice**
> *os.File ProvidersRetrieveInvoice(ctx, providerId, iD)
Download a provider's invoice.

Retrieves an invoice for this provider. You must be associated to the provider to use this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**|  | 
  **iD** | **string**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProvidersRetrieveInvoices**
> ProviderInvoiceResponse ProvidersRetrieveInvoices(ctx, providerId, optional)
List a provider's invoices.

Retrieves a list of invoices for this provider. You must be associated to the provider to use this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**|  | 
 **optional** | ***ManagedServiceProviderApiProvidersRetrieveInvoicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagedServiceProviderApiProvidersRetrieveInvoicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **optional.Int32**| The offset into the records to return. | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 
 **limit** | **optional.Int32**| The number of records to return at once. Limited to 100. | [default to 10]

### Return type

[**ProviderInvoiceResponse**](ProviderInvoiceResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

