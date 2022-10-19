# {{classname}}

All URIs are relative to *https://console.jumpcloud.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemsCommandBuiltinErase**](SystemsApi.md#SystemsCommandBuiltinErase) | **Post** /systems/{system_id}/command/builtin/erase | Erase a System
[**SystemsCommandBuiltinLock**](SystemsApi.md#SystemsCommandBuiltinLock) | **Post** /systems/{system_id}/command/builtin/lock | Lock a System
[**SystemsCommandBuiltinRestart**](SystemsApi.md#SystemsCommandBuiltinRestart) | **Post** /systems/{system_id}/command/builtin/restart | Restart a System
[**SystemsCommandBuiltinShutdown**](SystemsApi.md#SystemsCommandBuiltinShutdown) | **Post** /systems/{system_id}/command/builtin/shutdown | Shutdown a System
[**SystemsDelete**](SystemsApi.md#SystemsDelete) | **Delete** /systems/{id} | Delete a System
[**SystemsGet**](SystemsApi.md#SystemsGet) | **Get** /systems/{id} | List an individual system
[**SystemsList**](SystemsApi.md#SystemsList) | **Get** /systems | List All Systems
[**SystemsPut**](SystemsApi.md#SystemsPut) | **Put** /systems/{id} | Update a system

# **SystemsCommandBuiltinErase**
> SystemsCommandBuiltinErase(ctx, systemId, optional)
Erase a System

This endpoint allows you to run the erase command on the specified device. If a device is offline, the command will be run when the device becomes available.  #### Sample Request ``` curl -X POST \\   https://console.jumpcloud.com/api/systems/{system_id}/command/builtin/erase \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d {} ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **systemId** | **string**|  | 
 **optional** | ***SystemsApiSystemsCommandBuiltinEraseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemsApiSystemsCommandBuiltinEraseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemsCommandBuiltinLock**
> SystemsCommandBuiltinLock(ctx, systemId, optional)
Lock a System

This endpoint allows you to run the lock command on the specified device. If a device is offline, the command will be run when the device becomes available.  #### Sample Request ``` curl -X POST \\   https://console.jumpcloud.com/api/systems/{system_id}/command/builtin/lock \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d {} ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **systemId** | **string**|  | 
 **optional** | ***SystemsApiSystemsCommandBuiltinLockOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemsApiSystemsCommandBuiltinLockOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemsCommandBuiltinRestart**
> SystemsCommandBuiltinRestart(ctx, systemId, optional)
Restart a System

This endpoint allows you to run the restart command on the specified device. If a device is offline, the command will be run when the device becomes available.  #### Sample Request ``` curl -X POST \\   https://console.jumpcloud.com/api/systems/{system_id}/command/builtin/restart \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d {} ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **systemId** | **string**|  | 
 **optional** | ***SystemsApiSystemsCommandBuiltinRestartOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemsApiSystemsCommandBuiltinRestartOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemsCommandBuiltinShutdown**
> SystemsCommandBuiltinShutdown(ctx, systemId, optional)
Shutdown a System

This endpoint allows you to run the shutdown command on the specified device. If a device is offline, the command will be run when the device becomes available.  #### Sample Request ``` curl -X POST \\   https://console.jumpcloud.com/api/systems/{system_id}/command/builtin/shutdown \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d {} ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **systemId** | **string**|  | 
 **optional** | ***SystemsApiSystemsCommandBuiltinShutdownOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemsApiSystemsCommandBuiltinShutdownOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemsDelete**
> System SystemsDelete(ctx, id, optional)
Delete a System

This endpoint allows you to delete a system. This command will cause the system to uninstall the JumpCloud agent from its self which can can take about a minute. If the system is not connected to JumpCloud the system record will simply be removed.  #### Sample Request ``` curl -X DELETE https://console.jumpcloud.com/api/systems/{SystemID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SystemsApiSystemsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemsApiSystemsDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **optional.String**| Current date header for the System Context API | 
 **authorization** | **optional.String**| Authorization header for the System Context API | 
 **xOrgId** | **optional.String**|  | 

### Return type

[**System**](system.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemsGet**
> System SystemsGet(ctx, id, optional)
List an individual system

This endpoint returns an individual system.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/systems/{SystemID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SystemsApiSystemsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemsApiSystemsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Use a space seperated string of field parameters to include the data in the response. If omitted, the default list of fields will be returned.  | 
 **filter** | **optional.String**| A filter to apply to the query. See the supported operators below. For more complex searches, see the related &#x60;/search/&lt;domain&gt;&#x60; endpoints, e.g. &#x60;/search/systems&#x60;.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D; Supported operators are: - &#x60;$eq&#x60; (equals) - &#x60;$ne&#x60; (does not equal) - &#x60;$gt&#x60; (is greater than) - &#x60;$gte&#x60; (is greater than or equal to) - &#x60;$lt&#x60; (is less than) - &#x60;$lte&#x60; (is less than or equal to)  _Note: v1 operators differ from v2 operators._  _Note: For v1 operators, excluding the &#x60;$&#x60; will result in undefined behavior._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive.  **Examples** - &#x60;GET /users?filter&#x3D;username:$eq:testuser&#x60; - &#x60;GET /systemusers?filter&#x3D;password_expiration_date:$lte:2021-10-24&#x60; - &#x60;GET /systemusers?filter&#x3D;department:$ne:Accounting&#x60; - &#x60;GET /systems?filter[0]&#x3D;firstname:$eq:foo&amp;filter[1]&#x3D;lastname:$eq:bar&#x60; - this will    AND the filters together. - &#x60;GET /systems?filter[or][0]&#x3D;lastname:$eq:foo&amp;filter[or][1]&#x3D;lastname:$eq:bar&#x60; - this will    OR the filters together. | 
 **date** | **optional.String**| Current date header for the System Context API | 
 **authorization** | **optional.String**| Authorization header for the System Context API | 
 **xOrgId** | **optional.String**|  | 

### Return type

[**System**](system.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemsList**
> Systemslist SystemsList(ctx, optional)
List All Systems

This endpoint returns all Systems.  #### Sample Requests ``` curl -X GET https://console.jumpcloud.com/api/systems \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'  ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemsApiSystemsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemsApiSystemsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Use a space seperated string of field parameters to include the data in the response. If omitted, the default list of fields will be returned.  | 
 **limit** | **optional.Int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **optional.String**|  | 
 **search** | **optional.String**| A nested object containing a &#x60;searchTerm&#x60; string or array of strings and a list of &#x60;fields&#x60; to search on. | 
 **skip** | **optional.Int32**| The offset into the records to return. | [default to 0]
 **sort** | **optional.String**| Use space separated sort parameters to sort the collection. Default sort is ascending. Prefix with &#x60;-&#x60; to sort descending.  | 
 **filter** | **optional.String**| A filter to apply to the query. See the supported operators below. For more complex searches, see the related &#x60;/search/&lt;domain&gt;&#x60; endpoints, e.g. &#x60;/search/systems&#x60;.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D; Supported operators are: - &#x60;$eq&#x60; (equals) - &#x60;$ne&#x60; (does not equal) - &#x60;$gt&#x60; (is greater than) - &#x60;$gte&#x60; (is greater than or equal to) - &#x60;$lt&#x60; (is less than) - &#x60;$lte&#x60; (is less than or equal to)  _Note: v1 operators differ from v2 operators._  _Note: For v1 operators, excluding the &#x60;$&#x60; will result in undefined behavior._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive.  **Examples** - &#x60;GET /users?filter&#x3D;username:$eq:testuser&#x60; - &#x60;GET /systemusers?filter&#x3D;password_expiration_date:$lte:2021-10-24&#x60; - &#x60;GET /systemusers?filter&#x3D;department:$ne:Accounting&#x60; - &#x60;GET /systems?filter[0]&#x3D;firstname:$eq:foo&amp;filter[1]&#x3D;lastname:$eq:bar&#x60; - this will    AND the filters together. - &#x60;GET /systems?filter[or][0]&#x3D;lastname:$eq:foo&amp;filter[or][1]&#x3D;lastname:$eq:bar&#x60; - this will    OR the filters together. | 

### Return type

[**Systemslist**](systemslist.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemsPut**
> System SystemsPut(ctx, id, optional)
Update a system

This endpoint allows you to update a system.  #### Sample Request  ``` curl -X PUT https://console.jumpcloud.com/api/systems/{SystemID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{  \"displayName\":\"Name_Update\",  \"allowSshPasswordAuthentication\":\"true\",  \"allowSshRootLogin\":\"true\",  \"allowMultiFactorAuthentication\":\"true\",  \"allowPublicKeyAuthentication\":\"false\" }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SystemsApiSystemsPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemsApiSystemsPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Systemput**](Systemput.md)|  | 
 **date** | **optional.**| Current date header for the System Context API | 
 **authorization** | **optional.**| Authorization header for the System Context API | 
 **xOrgId** | **optional.**|  | 

### Return type

[**System**](system.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

