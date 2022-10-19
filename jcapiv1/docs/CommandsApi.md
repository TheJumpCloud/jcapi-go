# {{classname}}

All URIs are relative to *https://console.jumpcloud.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommandFileGet**](CommandsApi.md#CommandFileGet) | **Get** /files/command/{id} | Get a Command File
[**CommandsDelete**](CommandsApi.md#CommandsDelete) | **Delete** /commands/{id} | Delete a Command
[**CommandsGet**](CommandsApi.md#CommandsGet) | **Get** /commands/{id} | List an individual Command
[**CommandsGetResults**](CommandsApi.md#CommandsGetResults) | **Get** /commands/{id}/results | Get results for a specific command
[**CommandsList**](CommandsApi.md#CommandsList) | **Get** /commands | List All Commands
[**CommandsPost**](CommandsApi.md#CommandsPost) | **Post** /commands | Create A Command
[**CommandsPut**](CommandsApi.md#CommandsPut) | **Put** /commands/{id} | Update a Command

# **CommandFileGet**
> Commandfilereturn CommandFileGet(ctx, id, optional)
Get a Command File

This endpoint returns the uploaded file(s) associated with a specific command.  #### Sample Request  ``` curl -X GET https://console.jumpcloud.com/api/files/command/{commandID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CommandsApiCommandFileGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommandsApiCommandFileGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Use a space seperated string of field parameters to include the data in the response. If omitted, the default list of fields will be returned.  | 
 **limit** | **optional.Int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **optional.String**|  | 
 **skip** | **optional.Int32**| The offset into the records to return. | [default to 0]

### Return type

[**Commandfilereturn**](commandfilereturn.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommandsDelete**
> Command CommandsDelete(ctx, id, optional)
Delete a Command

This endpoint deletes a specific command based on the Command ID.  #### Sample Request ``` curl -X DELETE https://console.jumpcloud.com/api/commands/{CommandID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'  ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CommandsApiCommandsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommandsApiCommandsDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **optional.String**|  | 

### Return type

[**Command**](command.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommandsGet**
> Command CommandsGet(ctx, id, optional)
List an individual Command

This endpoint returns a specific command based on the command ID.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/commands/{CommandID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CommandsApiCommandsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommandsApiCommandsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Use a space seperated string of field parameters to include the data in the response. If omitted, the default list of fields will be returned.  | 
 **xOrgId** | **optional.String**|  | 

### Return type

[**Command**](command.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommandsGetResults**
> []Commandresult CommandsGetResults(ctx, id)
Get results for a specific command

This endpoint returns results for a specific command.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/commands/{id}/results \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'   ````

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**[]Commandresult**](commandresult.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommandsList**
> Commandslist CommandsList(ctx, optional)
List All Commands

This endpoint returns all commands.  #### Sample Request ``` curl -X GET https://console.jumpcloud.com/api/commands/ \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'  ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CommandsApiCommandsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommandsApiCommandsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Use a space seperated string of field parameters to include the data in the response. If omitted, the default list of fields will be returned.  | 
 **filter** | **optional.String**| A filter to apply to the query. See the supported operators below. For more complex searches, see the related &#x60;/search/&lt;domain&gt;&#x60; endpoints, e.g. &#x60;/search/systems&#x60;.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D; Supported operators are: - &#x60;$eq&#x60; (equals) - &#x60;$ne&#x60; (does not equal) - &#x60;$gt&#x60; (is greater than) - &#x60;$gte&#x60; (is greater than or equal to) - &#x60;$lt&#x60; (is less than) - &#x60;$lte&#x60; (is less than or equal to)  _Note: v1 operators differ from v2 operators._  _Note: For v1 operators, excluding the &#x60;$&#x60; will result in undefined behavior._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive.  **Examples** - &#x60;GET /users?filter&#x3D;username:$eq:testuser&#x60; - &#x60;GET /systemusers?filter&#x3D;password_expiration_date:$lte:2021-10-24&#x60; - &#x60;GET /systemusers?filter&#x3D;department:$ne:Accounting&#x60; - &#x60;GET /systems?filter[0]&#x3D;firstname:$eq:foo&amp;filter[1]&#x3D;lastname:$eq:bar&#x60; - this will    AND the filters together. - &#x60;GET /systems?filter[or][0]&#x3D;lastname:$eq:foo&amp;filter[or][1]&#x3D;lastname:$eq:bar&#x60; - this will    OR the filters together. | 
 **limit** | **optional.Int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **optional.String**|  | 
 **skip** | **optional.Int32**| The offset into the records to return. | [default to 0]
 **sort** | **optional.String**| Use space separated sort parameters to sort the collection. Default sort is ascending. Prefix with &#x60;-&#x60; to sort descending.  | 

### Return type

[**Commandslist**](commandslist.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommandsPost**
> Command CommandsPost(ctx, optional)
Create A Command

This endpoint allows you to create a new command.  #### Sample Request  ``` curl -X POST https://console.jumpcloud.com/api/commands/ \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{  \"name\":\"Test API Command\",  \"command\":\"String\",  \"user\":\"{UserID}\",  \"schedule\":\"\",  \"timeout\":\"100\" }'  ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CommandsApiCommandsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommandsApiCommandsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Command**](Command.md)|  | 
 **xOrgId** | **optional.**|  | 

### Return type

[**Command**](command.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommandsPut**
> Command CommandsPut(ctx, id, optional)
Update a Command

This endpoint Updates a command based on the command ID and returns the modified command record.  #### Sample Request ``` curl -X PUT https://console.jumpcloud.com/api/commands/{CommandID} \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{  \"name\":\"Test API Command\",  \"command\":\"String\",  \"user\":\"{UserID}\",  \"schedule\":\"\",  \"timeout\":\"100\" }'  ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CommandsApiCommandsPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommandsApiCommandsPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Command**](Command.md)|  | 
 **xOrgId** | **optional.**|  | 

### Return type

[**Command**](command.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

