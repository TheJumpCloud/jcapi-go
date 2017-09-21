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
> string LdapserversSambaDomainsDelete($ldapserverId, $id, $contentType, $accept)

Delete Samba Domain

This endpoint allows you to delete a samba domain from an LDAP server.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapserverId** | **string**| Unique identifier o f the LDAP server. | 
 **id** | **string**| Unique identifier of the samba domain. | 
 **contentType** | **string**|  | [optional] [default to application/json]
 **accept** | **string**|  | [optional] [default to application/json]

### Return type

**string**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapserversSambaDomainsGet**
> SambaDomainOutput LdapserversSambaDomainsGet($ldapserverId, $id, $contentType, $accept)

Get Samba Domain

This endpoint returns a specific samba domain for an LDAP server.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapserverId** | **string**| Unique identifier o f the LDAP server. | 
 **id** | **string**| Unique identifier of the samba domain. | 
 **contentType** | **string**|  | [optional] [default to application/json]
 **accept** | **string**|  | [optional] [default to application/json]

### Return type

[**SambaDomainOutput**](samba-domain-output.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapserversSambaDomainsList**
> []SambaDomainOutput LdapserversSambaDomainsList($ldapserverId, $contentType, $accept, $fields, $filter, $limit, $skip, $sort)

List Samba Domains

This endpoint returns all samba domains for an LDAP server.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapserverId** | **string**| Unique identifier of the LDAP server. | 
 **contentType** | **string**|  | [optional] [default to application/json]
 **accept** | **string**|  | [optional] [default to application/json]
 **fields** | **string**| The comma separated fields included in the returned records. If omitted the default list of fields will be returned.  | [optional] [default to ]
 **filter** | **string**| Supported operators are: eq, ne, gt, ge, lt, le, between, search | [optional] [default to ]
 **limit** | **int32**| The number of records to return at once. | [optional] [default to 10]
 **skip** | **int32**| The offset into the records to return. | [optional] [default to 0]
 **sort** | **string**| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [optional] [default to ]

### Return type

[**[]SambaDomainOutput**](samba-domain-output.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapserversSambaDomainsPost**
> SambaDomainOutput LdapserversSambaDomainsPost($ldapserverId, $body, $contentType, $accept)

Create Samba Domain

This endpoint allows you to create a samba domain for an LDAP server.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapserverId** | **string**| Unique identifier of the LDAP server. | 
 **body** | [**SambaDomainInput**](SambaDomainInput.md)|  | [optional] 
 **contentType** | **string**|  | [optional] [default to application/json]
 **accept** | **string**|  | [optional] [default to application/json]

### Return type

[**SambaDomainOutput**](samba-domain-output.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapserversSambaDomainsPut**
> SambaDomainOutput LdapserversSambaDomainsPut($ldapserverId, $id, $body, $contentType, $accept)

Update Samba Domain

This endpoint allows you to update the samba domain information for an LDAP server.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapserverId** | **string**| Unique identifier o f the LDAP server. | 
 **id** | **string**| Unique identifier of the samba domain. | 
 **body** | [**SambaDomainInput**](SambaDomainInput.md)|  | [optional] 
 **contentType** | **string**|  | [optional] [default to application/json]
 **accept** | **string**|  | [optional] [default to application/json]

### Return type

[**SambaDomainOutput**](samba-domain-output.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

