# \SambaDomainsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LdapserversSambaDomainsDelete**](SambaDomainsApi.md#LdapserversSambaDomainsDelete) | **Delete** /ldapservers/{ldapserver_id}/sambadomains/{id} | Delete Samba Domain
[**LdapserversSambaDomainsGet**](SambaDomainsApi.md#LdapserversSambaDomainsGet) | **Get** /ldapservers/{ldapserver_id}/sambadomains/{id} | Get Samba Domain
[**LdapserversSambaDomainsList**](SambaDomainsApi.md#LdapserversSambaDomainsList) | **Get** /ldapservers/{ldapserver_id}/sambadomains | List Samba Domains
[**LdapserversSambaDomainsPost**](SambaDomainsApi.md#LdapserversSambaDomainsPost) | **Post** /ldapservers/{ldapserver_id}/sambadomains | Create Samba Domain
[**LdapserversSambaDomainsPut**](SambaDomainsApi.md#LdapserversSambaDomainsPut) | **Put** /ldapservers/{ldapserver_id}/sambadomains/{id} | Update Samba Domain


# **LdapserversSambaDomainsDelete**
> string LdapserversSambaDomainsDelete(ctx, ldapserverId, id, optional)
Delete Samba Domain

This endpoint allows you to delete a samba domain from an LDAP server.  ##### Sample Request ``` curl -X DELETE https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID}/sambadomains/{SAMBA_ID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ldapserverId** | **string**| Unique identifier of the LDAP server. | 
  **id** | **string**| Unique identifier of the samba domain. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapserverId** | **string**| Unique identifier of the LDAP server. | 
 **id** | **string**| Unique identifier of the samba domain. | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **xOrgId** | **string**|  | [default to ]

### Return type

**string**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapserversSambaDomainsGet**
> SambaDomainOutput LdapserversSambaDomainsGet(ctx, ldapserverId, id, optional)
Get Samba Domain

This endpoint returns a specific samba domain for an LDAP server.  ##### Sample Request ``` curl -X GET \\   https://console.jumpcloud.com/api/v2/ldapservers/ldapservers/{LDAP_ID}/sambadomains/{SAMBA_ID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ldapserverId** | **string**| Unique identifier of the LDAP server. | 
  **id** | **string**| Unique identifier of the samba domain. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapserverId** | **string**| Unique identifier of the LDAP server. | 
 **id** | **string**| Unique identifier of the samba domain. | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **xOrgId** | **string**|  | [default to ]

### Return type

[**SambaDomainOutput**](samba-domain-output.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapserversSambaDomainsList**
> []SambaDomainOutput LdapserversSambaDomainsList(ctx, ldapserverId, optional)
List Samba Domains

This endpoint returns all samba domains for an LDAP server.  ##### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID}/sambadomains \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ldapserverId** | **string**| Unique identifier of the LDAP server. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapserverId** | **string**| Unique identifier of the LDAP server. | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **fields** | [**[]string**](string.md)| The comma separated fields included in the returned records. If omitted the default list of fields will be returned.  | 
 **filter** | [**[]string**](string.md)| Supported operators are: eq, ne, gt, ge, lt, le, between, search, in | 
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | [**[]string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]SambaDomainOutput**](samba-domain-output.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapserversSambaDomainsPost**
> SambaDomainOutput LdapserversSambaDomainsPost(ctx, ldapserverId, optional)
Create Samba Domain

This endpoint allows you to create a samba domain for an LDAP server.  ##### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID}/sambadomains \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{ \"sid\":\"{SID_ID}\", \"name\":\"{WORKGROUP_NAME}\" }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ldapserverId** | **string**| Unique identifier of the LDAP server. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapserverId** | **string**| Unique identifier of the LDAP server. | 
 **body** | [**SambaDomainInput**](SambaDomainInput.md)|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **xOrgId** | **string**|  | [default to ]

### Return type

[**SambaDomainOutput**](samba-domain-output.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapserversSambaDomainsPut**
> SambaDomainOutput LdapserversSambaDomainsPut(ctx, ldapserverId, id, optional)
Update Samba Domain

This endpoint allows you to update the samba domain information for an LDAP server.  ##### Sample Request ``` curl -X PUT https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID}/sambadomains/{SAMBA_ID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{ \"sid\":\"{SID_ID}\", \"name\":\"{WORKGROUP_NAME}\" }'  ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ldapserverId** | **string**| Unique identifier of the LDAP server. | 
  **id** | **string**| Unique identifier of the samba domain. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapserverId** | **string**| Unique identifier of the LDAP server. | 
 **id** | **string**| Unique identifier of the samba domain. | 
 **body** | [**SambaDomainInput**](SambaDomainInput.md)|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **xOrgId** | **string**|  | [default to ]

### Return type

[**SambaDomainOutput**](samba-domain-output.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

