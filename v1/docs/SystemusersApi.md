# \SystemusersApi

All URIs are relative to *https://console.jumpcloud.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SshkeyDelete**](SystemusersApi.md#SshkeyDelete) | **Delete** /systemusers/{id}/sshkeys/{id} | Delete a system user&#39;s Public SSH Keys
[**SshkeyList**](SystemusersApi.md#SshkeyList) | **Get** /systemusers/{id}/sshkeys | List a system user&#39;s public SSH keys
[**SshkeyPost**](SystemusersApi.md#SshkeyPost) | **Post** /systemusers/{id}/sshkeys | Create a system user&#39;s Public SSH Key
[**SystemusersDelete**](SystemusersApi.md#SystemusersDelete) | **Delete** /systemusers/{id} | Delete a system user
[**SystemusersGet**](SystemusersApi.md#SystemusersGet) | **Get** /systemusers/{id} | List a system user
[**SystemusersList**](SystemusersApi.md#SystemusersList) | **Get** /systemusers | List all system users
[**SystemusersPost**](SystemusersApi.md#SystemusersPost) | **Post** /systemusers | Create a system user
[**SystemusersPut**](SystemusersApi.md#SystemusersPut) | **Put** /systemusers/{id} | Update a system user
[**SystemusersResetmfa**](SystemusersApi.md#SystemusersResetmfa) | **Post** /systemusers/{id}/resetmfa | Reset a system user&#39;s MFA token
[**SystemusersSystemsBindingList**](SystemusersApi.md#SystemusersSystemsBindingList) | **Get** /systemusers/{id}/systems | List system user binding
[**SystemusersSystemsBindingPut**](SystemusersApi.md#SystemusersSystemsBindingPut) | **Put** /systemusers/{id}/systems | Update a system user binding


# **SshkeyDelete**
> SshkeyDelete(ctx, id, contentType, accept, optional)
Delete a system user's Public SSH Keys

This endpoint will delete a specific System User's SSH Key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **xOrgId** | **string**|  | [default to ]

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SshkeyList**
> Sshkeylist SshkeyList(ctx, id, contentType, accept, optional)
List a system user's public SSH keys

This endpoint will return a specific System User's public SSH key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **xOrgId** | **string**|  | [default to ]

### Return type

[**Sshkeylist**](sshkeylist.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SshkeyPost**
> Sshkeylist SshkeyPost(ctx, id, contentType, accept, optional)
Create a system user's Public SSH Key

This endpoint will create a specific System User's Public SSH Key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **body** | [**Sshkeypost**](Sshkeypost.md)|  | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**Sshkeylist**](sshkeylist.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemusersDelete**
> Systemuserreturn SystemusersDelete(ctx, id, contentType, accept, optional)
Delete a system user

This endpoint allows you to delete a particular system user.  #### Sample Request ``` curl -X DELETE https://console.jumpcloud.com/api/systemusers/{UserID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **xOrgId** | **string**|  | [default to ]

### Return type

[**Systemuserreturn**](systemuserreturn.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemusersGet**
> Systemuserreturn SystemusersGet(ctx, id, contentType, accept, optional)
List a system user

This endpoint returns a particular System User.  #### Sample Request  ``` curl -X GET https://console.jumpcloud.com/api/systemusers/{UserID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **fields** | **string**| Use a space seperated string of field parameters to include the data in the response. If omitted the default list of fields will be returned.  | [default to ]
 **xOrgId** | **string**|  | [default to ]

### Return type

[**Systemuserreturn**](systemuserreturn.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemusersList**
> Systemuserslist SystemusersList(ctx, contentType, accept, optional)
List all system users

This endpoint returns all systemusers.  #### Sample Request  ``` curl -X GET https://console.jumpcloud.com/api/systemusers \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **limit** | **int32**| The number of records to return at once. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | **string**| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to ]
 **fields** | **string**| The comma separated fields included in the returned records. If omitted the default list of fields will be returned.  | [default to ]
 **filter** | **string**|  | [default to ]
 **xOrgId** | **string**|  | [default to ]

### Return type

[**Systemuserslist**](systemuserslist.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemusersPost**
> Systemuserreturn SystemusersPost(ctx, contentType, accept, optional)
Create a system user

This endpoint allows you to create a new system user.  #### Sample Request  ``` curl -X POST https://console.jumpcloud.com/api/systemusers \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{  \"username\":\"{username}\",  \"email\":\"{email_address}\",  \"firstname\":\"{Name}\",  \"lastname\":\"{Name}\" }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **body** | [**Systemuserputpost**](Systemuserputpost.md)|  | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**Systemuserreturn**](systemuserreturn.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemusersPut**
> Systemuserreturn SystemusersPut(ctx, id, contentType, accept, optional)
Update a system user

This endpoint allows you to update a system user.  #### Sample Request  ``` curl -X PUT https://console.jumpcloud.com/api/systemusers/{UserID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{  \"email\":\"{email_address}\",  \"firstname\":\"{Name}\",  \"lastname\":\"{Name}\" }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **body** | [**Systemuserput**](Systemuserput.md)|  | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**Systemuserreturn**](systemuserreturn.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemusersResetmfa**
> string SystemusersResetmfa(ctx, id, xApiKey, optional)
Reset a system user's MFA token

This endpoint allows you to reset the MFA TOTP token for a specified system user and put them in an MFA enrollment period. This will result in the user being prompted to setup MFA when logging into userportal. Please be aware that if the user does not complete MFA setup before the `exclusionUntil` date, they will be locked out of any resources that require MFA.  Please refer to our [Knowledge Base Article](https://support.jumpcloud.com/customer/en/portal/articles/2959138-using-multifactor-authentication-with-jumpcloud) on setting up MFA for more information.  #### Sample Request ``` curl -X POST \\   https://console.jumpcloud.com/api/systemusers/{UserID}/resetmfa \\   -H 'x-api-key: {API_KEY}' \\   -H 'Content-Type: application/json' \\   -d '{\"exclusion\": true, \"exclusionUntil\": \"{date-time}\"}'   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 
  **xApiKey** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **xApiKey** | **string**|  | 
 **body** | [**Body1**](Body1.md)|  | 
 **xOrgId** | **string**|  | [default to ]

### Return type

**string**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemusersSystemsBindingList**
> interface{} SystemusersSystemsBindingList(ctx, id, contentType, accept, optional)
List system user binding

Hidden as Tags is deprecated  Adds or removes a system binding for a user.  This endpoint is only used for users still using JumpCloud Tags. If you are using JumpCloud Groups please refer to the documentation found [here](https://docs.jumpcloud.com/2.0/systems/manage-associations-of-a-system).   List system bindings for a specific system user in a system and user binding format.  ### Examples  #### List system bindings for specific system user  ``` curl \\   -H 'Content-Type: application/json' \\   -H \"x-api-key: [YOUR_API_KEY_HERE]\" \\   \"https://console.jumpcloud.com/api/systemusers/[SYSTEM_USER_ID_HERE]/systems\" ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **fields** | **string**| Use a space seperated string of field parameters to include the data in the response. If omitted the default list of fields will be returned.  | [default to ]
 **limit** | **int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | **string**| Use space separated sort parameters to sort the collection. Default sort is ascending. Prefix with &#x60;-&#x60; to sort descending.  | [default to ]
 **xOrgId** | **string**|  | [default to ]

### Return type

[**interface{}**](interface{}.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemusersSystemsBindingPut**
> Usersystembinding SystemusersSystemsBindingPut(ctx, id, contentType, accept, optional)
Update a system user binding

Hidden as Tags is deprecated  Adds or removes a system binding for a user.  This endpoint is only used for users still using JumpCloud Tags. If you are using JumpCloud Groups please refer to the documentation found [here](https://docs.jumpcloud.com/2.0/systems/manage-associations-of-a-system).  ### Example  #### Add (or remove) system to system user  ``` curl \\   -d '{ \"add\": [\"[SYSTEM_ID_TO_ADD_HERE]\"], \"remove\": [\"[SYSTEM_ID_TO_REMOVE_HERE]\"] }' \\   -X PUT \\   -H 'Content-Type: application/json' \\   -H 'Accept: application/json' \\   -H \"x-api-key: [YOUR_API_KEY_HERE]\" \\   \"https://console.jumpcloud.com/api/systemusers/[SYSTEM_USER_ID_HERE]/systems\" ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **body** | [**Usersystembindingsput**](Usersystembindingsput.md)|  | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**Usersystembinding**](usersystembinding.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

