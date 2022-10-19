# {{classname}}

All URIs are relative to *https://console.jumpcloud.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RadiusServersDelete**](RadiusServersApi.md#RadiusServersDelete) | **Delete** /radiusservers/{id} | Delete Radius Server
[**RadiusServersGet**](RadiusServersApi.md#RadiusServersGet) | **Get** /radiusservers/{id} | Get Radius Server
[**RadiusServersList**](RadiusServersApi.md#RadiusServersList) | **Get** /radiusservers | List Radius Servers
[**RadiusServersPost**](RadiusServersApi.md#RadiusServersPost) | **Post** /radiusservers | Create a Radius Server
[**RadiusServersPut**](RadiusServersApi.md#RadiusServersPut) | **Put** /radiusservers/{id} | Update Radius Servers

# **RadiusServersDelete**
> Radiusserverput RadiusServersDelete(ctx, id, optional)
Delete Radius Server

This endpoint allows you to delete RADIUS servers in your organization. ``` curl -X DELETE https://console.jumpcloud.com/api/radiusservers/{ServerID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\ ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***RadiusServersApiRadiusServersDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RadiusServersApiRadiusServersDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **optional.String**|  | 

### Return type

[**Radiusserverput**](radiusserverput.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RadiusServersGet**
> Radiusserver RadiusServersGet(ctx, id, optional)
Get Radius Server

This endpoint allows you to get a RADIUS server in your organization.  #### ``` curl -X PUT https://console.jumpcloud.com/api/radiusservers/{ServerID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\ ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***RadiusServersApiRadiusServersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RadiusServersApiRadiusServersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **optional.String**|  | 

### Return type

[**Radiusserver**](radiusserver.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RadiusServersList**
> Radiusserverslist RadiusServersList(ctx, optional)
List Radius Servers

This endpoint allows you to get a list of all RADIUS servers in your organization.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/radiusservers/ \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\ ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RadiusServersApiRadiusServersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RadiusServersApiRadiusServersListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Use a space seperated string of field parameters to include the data in the response. If omitted, the default list of fields will be returned.  | 
 **filter** | **optional.String**| A filter to apply to the query. See the supported operators below. For more complex searches, see the related &#x60;/search/&lt;domain&gt;&#x60; endpoints, e.g. &#x60;/search/systems&#x60;.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D; Supported operators are: - &#x60;$eq&#x60; (equals) - &#x60;$ne&#x60; (does not equal) - &#x60;$gt&#x60; (is greater than) - &#x60;$gte&#x60; (is greater than or equal to) - &#x60;$lt&#x60; (is less than) - &#x60;$lte&#x60; (is less than or equal to)  _Note: v1 operators differ from v2 operators._  _Note: For v1 operators, excluding the &#x60;$&#x60; will result in undefined behavior._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive.  **Examples** - &#x60;GET /users?filter&#x3D;username:$eq:testuser&#x60; - &#x60;GET /systemusers?filter&#x3D;password_expiration_date:$lte:2021-10-24&#x60; - &#x60;GET /systemusers?filter&#x3D;department:$ne:Accounting&#x60; - &#x60;GET /systems?filter[0]&#x3D;firstname:$eq:foo&amp;filter[1]&#x3D;lastname:$eq:bar&#x60; - this will    AND the filters together. - &#x60;GET /systems?filter[or][0]&#x3D;lastname:$eq:foo&amp;filter[or][1]&#x3D;lastname:$eq:bar&#x60; - this will    OR the filters together. | 
 **limit** | **optional.Int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **optional.Int32**| The offset into the records to return. | [default to 0]
 **sort** | **optional.String**| Use space separated sort parameters to sort the collection. Default sort is ascending. Prefix with &#x60;-&#x60; to sort descending.  | 
 **xOrgId** | **optional.String**|  | 

### Return type

[**Radiusserverslist**](radiusserverslist.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RadiusServersPost**
> Radiusserver RadiusServersPost(ctx, optional)
Create a Radius Server

This endpoint allows you to create RADIUS servers in your organization.  #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/radiusservers/ \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{     \"name\": \"{test_radius}\",     \"networkSourceIp\": \"{0.0.0.0}\",     \"sharedSecret\":\"{secretpassword}\",     \"userLockoutAction\": \"REMOVE\",     \"userPasswordExpirationAction\": \"MAINTAIN\" }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RadiusServersApiRadiusServersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RadiusServersApiRadiusServersPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Radiusserverpost**](Radiusserverpost.md)|  | 
 **xOrgId** | **optional.**|  | 

### Return type

[**Radiusserver**](radiusserver.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RadiusServersPut**
> Radiusserverput RadiusServersPut(ctx, id, optional)
Update Radius Servers

This endpoint allows you to update RADIUS servers in your organization.  #### ``` curl -X PUT https://console.jumpcloud.com/api/radiusservers/{ServerID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{     \"name\": \"{name_update}\",     \"networkSourceIp\": \"{0.0.0.0}\",     \"sharedSecret\": \"{secret_password}\",     \"userLockoutAction\": \"REMOVE\",     \"userPasswordExpirationAction\": \"MAINTAIN\" }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***RadiusServersApiRadiusServersPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RadiusServersApiRadiusServersPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RadiusserversIdBody**](RadiusserversIdBody.md)|  | 
 **xOrgId** | **optional.**|  | 

### Return type

[**Radiusserverput**](radiusserverput.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

