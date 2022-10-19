# {{classname}}

All URIs are relative to *https://console.jumpcloud.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchCommandresultsPost**](SearchApi.md#SearchCommandresultsPost) | **Post** /search/commandresults | Search Commands Results
[**SearchCommandsPost**](SearchApi.md#SearchCommandsPost) | **Post** /search/commands | Search Commands
[**SearchOrganizationsPost**](SearchApi.md#SearchOrganizationsPost) | **Post** /search/organizations | Search Organizations
[**SearchSystemsPost**](SearchApi.md#SearchSystemsPost) | **Post** /search/systems | Search Systems
[**SearchSystemusersPost**](SearchApi.md#SearchSystemusersPost) | **Post** /search/systemusers | Search System Users

# **SearchCommandresultsPost**
> Commandresultslist SearchCommandresultsPost(ctx, optional)
Search Commands Results

Return Command Results in multi-record format allowing for the passing of the `filter` parameter.  To support advanced filtering you can use the `filter` and `searchFilter` parameters that can only be passed in the body of POST /api/search/commandresults route. The `filter` parameter must be passed as Content-Type application/json.  The `filter` parameter is an object with a single property, either `and` or `or` with the value of the property being an array of query expressions.  This allows you to filter records using the logic of matching ALL or ANY records in the array of query expressions. If the `and` or `or` are not included the default behavior is to match ALL query expressions.   #### Sample Request  Exact search for a specific command result ``` curl -X POST https://console.jumpcloud.com/api/search/commandresults \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   \"filter\" : \"workflowInstanceId:$eq:62f3c599ec4e928499069c7f\",   \"fields\" : \"name workflowId sudo\" }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiSearchCommandresultsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchCommandresultsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Search**](Search.md)|  | 
 **fields** | **optional.**| Use a space seperated string of field parameters to include the data in the response. If omitted, the default list of fields will be returned.  | 
 **filter** | **optional.**| A filter to apply to the query. See the supported operators below. For more complex searches, see the related &#x60;/search/&lt;domain&gt;&#x60; endpoints, e.g. &#x60;/search/systems&#x60;.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D; Supported operators are: - &#x60;$eq&#x60; (equals) - &#x60;$ne&#x60; (does not equal) - &#x60;$gt&#x60; (is greater than) - &#x60;$gte&#x60; (is greater than or equal to) - &#x60;$lt&#x60; (is less than) - &#x60;$lte&#x60; (is less than or equal to)  _Note: v1 operators differ from v2 operators._  _Note: For v1 operators, excluding the &#x60;$&#x60; will result in undefined behavior._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive.  **Examples** - &#x60;GET /users?filter&#x3D;username:$eq:testuser&#x60; - &#x60;GET /systemusers?filter&#x3D;password_expiration_date:$lte:2021-10-24&#x60; - &#x60;GET /systemusers?filter&#x3D;department:$ne:Accounting&#x60; - &#x60;GET /systems?filter[0]&#x3D;firstname:$eq:foo&amp;filter[1]&#x3D;lastname:$eq:bar&#x60; - this will    AND the filters together. - &#x60;GET /systems?filter[or][0]&#x3D;lastname:$eq:foo&amp;filter[or][1]&#x3D;lastname:$eq:bar&#x60; - this will    OR the filters together. | 
 **limit** | **optional.**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **optional.**| The offset into the records to return. | [default to 0]
 **xOrgId** | **optional.**|  | 

### Return type

[**Commandresultslist**](commandresultslist.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchCommandsPost**
> Commandslist SearchCommandsPost(ctx, optional)
Search Commands

Return Commands in multi-record format allowing for the passing of the `filter` and `searchFilter` parameters. This WILL NOT allow you to add a new command. To support advanced filtering you can use the `filter` and `searchFilter` parameters that can only be passed in the body of POST /api/search/_* routes. The `filter` and `searchFilter` parameters must be passed as Content-Type application/json. The `filter` parameter is an object with a single property, either `and` or `or` with the value of the property being an array of query expressions. This allows you to filter records using the logic of matching ALL or ANY records in the array of query expressions. If the `and` or `or` are not included the default behavior is to match ALL query expressions. The `searchFilter` parameter allows text searching on supported fields by specifying a `searchTerm` and a list of `fields` to query on. If any `field` has a partial text match on the `searchTerm` the record will be returned.  #### Sample Request Exact search for a list of commands in a launchType ``` curl -X POST https://console.jumpcloud.com/api/search/commands \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   \"filter\" : [{\"launchType\" : \"repeated\"}],   \"fields\" : \"name launchType sudo\" }' ``` Text search for commands with name ``` curl -X POST https://console.jumpcloud.com/api/search/commands \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   \"searchFilter\" : {     \"searchTerm\": \"List\",     \"fields\": [\"name\"]   },   \"fields\" : \"name launchType sudo\" }' ``` Text search for multiple commands ``` curl -X POST https://console.jumpcloud.com/api/search/commands \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   \"searchFilter\" : {     \"searchTerm\": [\"List\", \"Log\"],     \"fields\": [\"name\"]   },   \"fields\" : \"name launchType sudo\" }' ``` Combining `filter` and `searchFilter` to text search for commands with name who are in a list of launchType ``` curl -X POST https://console.jumpcloud.com/api/search/commands \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   \"searchFilter\": {     \"searchTerm\": \"List\",     \"fields\": [\"name\"]   },   \"filter\": {     \"or\": [       {\"launchType\" : \"repeated\"},       {\"launchType\" : \"one-time\"}     ]   },   \"fields\" : \"name launchType sudo\" }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiSearchCommandsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchCommandsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Search**](Search.md)|  | 
 **fields** | **optional.**| Use a space seperated string of field parameters to include the data in the response. If omitted, the default list of fields will be returned.  | 
 **filter** | **optional.**| A filter to apply to the query. See the supported operators below. For more complex searches, see the related &#x60;/search/&lt;domain&gt;&#x60; endpoints, e.g. &#x60;/search/systems&#x60;.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D; Supported operators are: - &#x60;$eq&#x60; (equals) - &#x60;$ne&#x60; (does not equal) - &#x60;$gt&#x60; (is greater than) - &#x60;$gte&#x60; (is greater than or equal to) - &#x60;$lt&#x60; (is less than) - &#x60;$lte&#x60; (is less than or equal to)  _Note: v1 operators differ from v2 operators._  _Note: For v1 operators, excluding the &#x60;$&#x60; will result in undefined behavior._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive.  **Examples** - &#x60;GET /users?filter&#x3D;username:$eq:testuser&#x60; - &#x60;GET /systemusers?filter&#x3D;password_expiration_date:$lte:2021-10-24&#x60; - &#x60;GET /systemusers?filter&#x3D;department:$ne:Accounting&#x60; - &#x60;GET /systems?filter[0]&#x3D;firstname:$eq:foo&amp;filter[1]&#x3D;lastname:$eq:bar&#x60; - this will    AND the filters together. - &#x60;GET /systems?filter[or][0]&#x3D;lastname:$eq:foo&amp;filter[or][1]&#x3D;lastname:$eq:bar&#x60; - this will    OR the filters together. | 
 **limit** | **optional.**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **optional.**| The offset into the records to return. | [default to 0]
 **xOrgId** | **optional.**|  | 

### Return type

[**Commandslist**](commandslist.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOrganizationsPost**
> Organizationslist SearchOrganizationsPost(ctx, optional)
Search Organizations

This endpoint will return Organization data based on your search parameters. This endpoint WILL NOT allow you to add a new Organization.  You can use the supported parameters and pass those in the body of request.  The parameters must be passed as Content-Type application/json.   #### Sample Request ``` curl -X POST https://console.jumpcloud.com/api/search/organizations \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   \"search\":{     \"fields\" : [\"settings.name\"],     \"searchTerm\": \"Second\"     },   \"fields\": [\"_id\", \"displayName\", \"logoUrl\"],   \"limit\" : 0,   \"skip\" : 0 }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiSearchOrganizationsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchOrganizationsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Search**](Search.md)|  | 
 **fields** | **optional.**| Use a space seperated string of field parameters to include the data in the response. If omitted, the default list of fields will be returned.  | 
 **filter** | **optional.**| A filter to apply to the query. See the supported operators below. For more complex searches, see the related &#x60;/search/&lt;domain&gt;&#x60; endpoints, e.g. &#x60;/search/systems&#x60;.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D; Supported operators are: - &#x60;$eq&#x60; (equals) - &#x60;$ne&#x60; (does not equal) - &#x60;$gt&#x60; (is greater than) - &#x60;$gte&#x60; (is greater than or equal to) - &#x60;$lt&#x60; (is less than) - &#x60;$lte&#x60; (is less than or equal to)  _Note: v1 operators differ from v2 operators._  _Note: For v1 operators, excluding the &#x60;$&#x60; will result in undefined behavior._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive.  **Examples** - &#x60;GET /users?filter&#x3D;username:$eq:testuser&#x60; - &#x60;GET /systemusers?filter&#x3D;password_expiration_date:$lte:2021-10-24&#x60; - &#x60;GET /systemusers?filter&#x3D;department:$ne:Accounting&#x60; - &#x60;GET /systems?filter[0]&#x3D;firstname:$eq:foo&amp;filter[1]&#x3D;lastname:$eq:bar&#x60; - this will    AND the filters together. - &#x60;GET /systems?filter[or][0]&#x3D;lastname:$eq:foo&amp;filter[or][1]&#x3D;lastname:$eq:bar&#x60; - this will    OR the filters together. | 
 **limit** | **optional.**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **optional.**| The offset into the records to return. | [default to 0]

### Return type

[**Organizationslist**](organizationslist.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSystemsPost**
> Systemslist SearchSystemsPost(ctx, optional)
Search Systems

Return Systems in multi-record format allowing for the passing of the `filter` and `searchFilter` parameters. This WILL NOT allow you to add a new system.  To support advanced filtering you can use the `filter` and `searchFilter` parameters that can only be passed in the body of POST /api/search/_* routes. The `filter` and `searchFilter` parameters must be passed as Content-Type application/json.  The `filter` parameter is an object with a single property, either `and` or `or` with the value of the property being an array of query expressions.  This allows you to filter records using the logic of matching ALL or ANY records in the array of query expressions. If the `and` or `or` are not included the default behavior is to match ALL query expressions.  The `searchFilter` parameter allows text searching on supported fields by specifying a `searchTerm` and a list of `fields` to query on. If any `field` has a partial text match on the `searchTerm` the record will be returned.   #### Sample Request  Exact search for a list of hostnames ``` curl -X POST https://console.jumpcloud.com/api/search/systems \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   \"filter\": {     \"or\": [       {\"hostname\" : \"my-hostname\"},       {\"hostname\" : \"other-hostname\"}     ]   },   \"fields\" : \"os hostname displayName\" }' ```  Text search for a hostname or display name ``` curl -X POST https://console.jumpcloud.com/api/search/systems \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   \"searchFilter\": {     \"searchTerm\": \"my-host\",     \"fields\": [\"hostname\", \"displayName\"]   },   \"fields\": \"os hostname displayName\" }' ```  Text search for a multiple hostnames. ``` curl -X POST https://console.jumpcloud.com/api/search/systems \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   \"searchFilter\": {     \"searchTerm\": [\"my-host\", \"my-other-host\"],     \"fields\": [\"hostname\"]   },   \"fields\": \"os hostname displayName\" }' ```  Combining `filter` and `searchFilter` to search for names that match a given OS ``` curl -X POST https://console.jumpcloud.com/api/search/systems \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   \"searchFilter\": {     \"searchTerm\": \"my-host\",     \"fields\": [\"hostname\", \"displayName\"]   },   \"filter\": {     \"or\": [       {\"os\" : \"Ubuntu\"},       {\"os\" : \"Mac OS X\"}     ]   },   \"fields\": \"os hostname displayName\" }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiSearchSystemsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchSystemsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Search**](Search.md)|  | 
 **fields** | **optional.**| Use a space seperated string of field parameters to include the data in the response. If omitted, the default list of fields will be returned.  | 
 **limit** | **optional.**| The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **optional.**|  | 
 **skip** | **optional.**| The offset into the records to return. | [default to 0]
 **filter** | **optional.**| A filter to apply to the query. See the supported operators below. For more complex searches, see the related &#x60;/search/&lt;domain&gt;&#x60; endpoints, e.g. &#x60;/search/systems&#x60;.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D; Supported operators are: - &#x60;$eq&#x60; (equals) - &#x60;$ne&#x60; (does not equal) - &#x60;$gt&#x60; (is greater than) - &#x60;$gte&#x60; (is greater than or equal to) - &#x60;$lt&#x60; (is less than) - &#x60;$lte&#x60; (is less than or equal to)  _Note: v1 operators differ from v2 operators._  _Note: For v1 operators, excluding the &#x60;$&#x60; will result in undefined behavior._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive.  **Examples** - &#x60;GET /users?filter&#x3D;username:$eq:testuser&#x60; - &#x60;GET /systemusers?filter&#x3D;password_expiration_date:$lte:2021-10-24&#x60; - &#x60;GET /systemusers?filter&#x3D;department:$ne:Accounting&#x60; - &#x60;GET /systems?filter[0]&#x3D;firstname:$eq:foo&amp;filter[1]&#x3D;lastname:$eq:bar&#x60; - this will    AND the filters together. - &#x60;GET /systems?filter[or][0]&#x3D;lastname:$eq:foo&amp;filter[or][1]&#x3D;lastname:$eq:bar&#x60; - this will    OR the filters together. | 

### Return type

[**Systemslist**](systemslist.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSystemusersPost**
> Systemuserslist SearchSystemusersPost(ctx, optional)
Search System Users

Return System Users in multi-record format allowing for the passing of the `filter` and `searchFilter` parameters. This WILL NOT allow you to add a new system user.  To support advanced filtering you can use the `filter` and `searchFilter` parameters that can only be passed in the body of POST /api/search/_* routes. The `filter` and `searchFilter` parameters must be passed as Content-Type application/json.  The `filter` parameter is an object with a single property, either `and` or `or` with the value of the property being an array of query expressions.  This allows you to filter records using the logic of matching ALL or ANY records in the array of query expressions. If the `and` or `or` are not included the default behavior is to match ALL query expressions.  The `searchFilter` parameter allows text searching on supported fields by specifying a `searchTerm` and a list of `fields` to query on. If any `field` has a partial text match on the `searchTerm` the record will be returned.   #### Sample Request  Exact search for a list of system users in a department ``` curl -X POST https://console.jumpcloud.com/api/search/systemusers \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   \"filter\" : [{\"department\" : \"IT\"}],   \"fields\" : \"email username sudo\" }' ```  Text search for system users with and email on a domain ``` curl -X POST https://console.jumpcloud.com/api/search/systemusers \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   \"searchFilter\" : {     \"searchTerm\": \"@jumpcloud.com\",     \"fields\": [\"email\"]   },   \"fields\" : \"email username sudo\" }' ```  Text search for multiple system users ``` curl -X POST https://console.jumpcloud.com/api/search/systemusers \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   \"searchFilter\" : {     \"searchTerm\": [\"john\", \"sarah\"],     \"fields\": [\"username\"]   },   \"fields\" : \"email username sudo\" }' ```  Combining `filter` and `searchFilter` to text search for system users with and email on a domain who are in a list of departments ``` curl -X POST https://console.jumpcloud.com/api/search/systemusers \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{   \"searchFilter\": {     \"searchTerm\": \"@jumpcloud.com\",     \"fields\": [\"email\"]   },   \"filter\": {     \"or\": [       {\"department\" : \"IT\"},       {\"department\" : \"Sales\"}     ]   },   \"fields\" : \"email username sudo\" }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiSearchSystemusersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchSystemusersPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Search**](Search.md)|  | 
 **fields** | **optional.**| Use a space seperated string of field parameters to include the data in the response. If omitted, the default list of fields will be returned.  | 
 **filter** | **optional.**| A filter to apply to the query. See the supported operators below. For more complex searches, see the related &#x60;/search/&lt;domain&gt;&#x60; endpoints, e.g. &#x60;/search/systems&#x60;.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D; Supported operators are: - &#x60;$eq&#x60; (equals) - &#x60;$ne&#x60; (does not equal) - &#x60;$gt&#x60; (is greater than) - &#x60;$gte&#x60; (is greater than or equal to) - &#x60;$lt&#x60; (is less than) - &#x60;$lte&#x60; (is less than or equal to)  _Note: v1 operators differ from v2 operators._  _Note: For v1 operators, excluding the &#x60;$&#x60; will result in undefined behavior._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive.  **Examples** - &#x60;GET /users?filter&#x3D;username:$eq:testuser&#x60; - &#x60;GET /systemusers?filter&#x3D;password_expiration_date:$lte:2021-10-24&#x60; - &#x60;GET /systemusers?filter&#x3D;department:$ne:Accounting&#x60; - &#x60;GET /systems?filter[0]&#x3D;firstname:$eq:foo&amp;filter[1]&#x3D;lastname:$eq:bar&#x60; - this will    AND the filters together. - &#x60;GET /systems?filter[or][0]&#x3D;lastname:$eq:foo&amp;filter[or][1]&#x3D;lastname:$eq:bar&#x60; - this will    OR the filters together. | 
 **limit** | **optional.**| The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **optional.**| The offset into the records to return. | [default to 0]
 **xOrgId** | **optional.**|  | 

### Return type

[**Systemuserslist**](systemuserslist.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

