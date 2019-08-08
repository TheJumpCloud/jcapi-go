# \SystemInsightsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SysteminsightsListApps**](SystemInsightsApi.md#SysteminsightsListApps) | **Get** /systeminsights/apps | List System Insights Apps
[**SysteminsightsListApps_0**](SystemInsightsApi.md#SysteminsightsListApps_0) | **Get** /systeminsights/{jc_system_id}/apps | List System Insights System Apps
[**SysteminsightsListBrowserPlugins**](SystemInsightsApi.md#SysteminsightsListBrowserPlugins) | **Get** /systeminsights/{jc_system_id}/browser_plugins | List System Insights System Browser Plugins
[**SysteminsightsListBrowserPlugins_0**](SystemInsightsApi.md#SysteminsightsListBrowserPlugins_0) | **Get** /systeminsights/browser_plugins | List System Insights Browser Plugins
[**SysteminsightsListChromeExtensions**](SystemInsightsApi.md#SysteminsightsListChromeExtensions) | **Get** /systeminsights/{jc_system_id}/chrome_extensions | List System Insights System Chrome Extensions
[**SysteminsightsListChromeExtensions_0**](SystemInsightsApi.md#SysteminsightsListChromeExtensions_0) | **Get** /systeminsights/chrome_extensions | List System Insights Chrome Extensions
[**SysteminsightsListDiskEncryption**](SystemInsightsApi.md#SysteminsightsListDiskEncryption) | **Get** /systeminsights/disk_encryption | List System Insights Disk Encryption
[**SysteminsightsListDiskEncryption_0**](SystemInsightsApi.md#SysteminsightsListDiskEncryption_0) | **Get** /systeminsights/{jc_system_id}/disk_encryption | List System Insights System Disk Encryption
[**SysteminsightsListFirefoxAddons**](SystemInsightsApi.md#SysteminsightsListFirefoxAddons) | **Get** /systeminsights/firefox_addons | List System Insights Firefox Addons
[**SysteminsightsListFirefoxAddons_0**](SystemInsightsApi.md#SysteminsightsListFirefoxAddons_0) | **Get** /systeminsights/{jc_system_id}/firefox_addons | List System Insights System Firefox Addons
[**SysteminsightsListGroups**](SystemInsightsApi.md#SysteminsightsListGroups) | **Get** /systeminsights/groups | List System Insights Groups
[**SysteminsightsListGroups_0**](SystemInsightsApi.md#SysteminsightsListGroups_0) | **Get** /systeminsights/{jc_system_id}/groups | List System Insights System Groups
[**SysteminsightsListInterfaceAddresses**](SystemInsightsApi.md#SysteminsightsListInterfaceAddresses) | **Get** /systeminsights/interface_addresses | List System Insights Interface Addresses
[**SysteminsightsListInterfaceAddresses_0**](SystemInsightsApi.md#SysteminsightsListInterfaceAddresses_0) | **Get** /systeminsights/{jc_system_id}/interface_addresses | List System Insights System Interface Addresses
[**SysteminsightsListMounts**](SystemInsightsApi.md#SysteminsightsListMounts) | **Get** /systeminsights/mounts | List System Insights Mounts
[**SysteminsightsListMounts_0**](SystemInsightsApi.md#SysteminsightsListMounts_0) | **Get** /systeminsights/{jc_system_id}/mounts | List System Insights System Mounts
[**SysteminsightsListOsVersion**](SystemInsightsApi.md#SysteminsightsListOsVersion) | **Get** /systeminsights/{jc_system_id}/os_version | List System Insights System OS Version
[**SysteminsightsListOsVersion_0**](SystemInsightsApi.md#SysteminsightsListOsVersion_0) | **Get** /systeminsights/os_version | List System Insights OS Version
[**SysteminsightsListSafariExtensions**](SystemInsightsApi.md#SysteminsightsListSafariExtensions) | **Get** /systeminsights/{jc_system_id}/safari_extensions | List System Insights System Safari Extensions
[**SysteminsightsListSafariExtensions_0**](SystemInsightsApi.md#SysteminsightsListSafariExtensions_0) | **Get** /systeminsights/safari_extensions | List System Insights Safari Extensions
[**SysteminsightsListSystemInfo**](SystemInsightsApi.md#SysteminsightsListSystemInfo) | **Get** /systeminsights/system_info | List System Insights System Info
[**SysteminsightsListSystemInfo_0**](SystemInsightsApi.md#SysteminsightsListSystemInfo_0) | **Get** /systeminsights/{jc_system_id}/system_info | List System Insights System System Info
[**SysteminsightsListUsers**](SystemInsightsApi.md#SysteminsightsListUsers) | **Get** /systeminsights/users | List System Insights Users
[**SysteminsightsListUsers_0**](SystemInsightsApi.md#SysteminsightsListUsers_0) | **Get** /systeminsights/{jc_system_id}/users | List System Insights System Users


# **SysteminsightsListApps**
> []SystemInsightsApps SysteminsightsListApps(ctx, optional)
List System Insights Apps

Valid filter fields are `jc_system_id` and `bundle_name`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsApps**](system-insights-apps.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListApps_0**
> []SystemInsightsApps SysteminsightsListApps_0(ctx, jcSystemId, optional)
List System Insights System Apps

Valid filter fields are `bundle_name`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jcSystemId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jcSystemId** | **string**|  | 
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsApps**](system-insights-apps.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListBrowserPlugins**
> []SystemInsightsBrowserPlugins SysteminsightsListBrowserPlugins(ctx, jcSystemId, optional)
List System Insights System Browser Plugins

Valid filter fields are `name`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jcSystemId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jcSystemId** | **string**|  | 
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsBrowserPlugins**](system-insights-browser-plugins.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListBrowserPlugins_0**
> []SystemInsightsBrowserPlugins SysteminsightsListBrowserPlugins_0(ctx, optional)
List System Insights Browser Plugins

Valid filter fields are `jc_system_id` and `name`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsBrowserPlugins**](system-insights-browser-plugins.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListChromeExtensions**
> []SystemInsightsChromeExtensions SysteminsightsListChromeExtensions(ctx, jcSystemId, optional)
List System Insights System Chrome Extensions

Valid filter fields are `name`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jcSystemId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jcSystemId** | **string**|  | 
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsChromeExtensions**](system-insights-chrome-extensions.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListChromeExtensions_0**
> []SystemInsightsChromeExtensions SysteminsightsListChromeExtensions_0(ctx, optional)
List System Insights Chrome Extensions

Valid filter fields are `jc_system_id` and `name`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsChromeExtensions**](system-insights-chrome-extensions.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListDiskEncryption**
> []SystemInsightsDiskEncryption SysteminsightsListDiskEncryption(ctx, optional)
List System Insights Disk Encryption

Valid filter fields are `jc_system_id` and `encryption_status`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsDiskEncryption**](system-insights-disk-encryption.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListDiskEncryption_0**
> []SystemInsightsDiskEncryption SysteminsightsListDiskEncryption_0(ctx, jcSystemId, optional)
List System Insights System Disk Encryption

Valid filter fields are `encryption_status`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jcSystemId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jcSystemId** | **string**|  | 
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsDiskEncryption**](system-insights-disk-encryption.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListFirefoxAddons**
> []SystemInsightsFirefoxAddons SysteminsightsListFirefoxAddons(ctx, optional)
List System Insights Firefox Addons

Valid filter fields are `jc_system_id` and `name`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsFirefoxAddons**](system-insights-firefox-addons.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListFirefoxAddons_0**
> []SystemInsightsFirefoxAddons SysteminsightsListFirefoxAddons_0(ctx, jcSystemId, optional)
List System Insights System Firefox Addons

Valid filter fields are `name`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jcSystemId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jcSystemId** | **string**|  | 
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsFirefoxAddons**](system-insights-firefox-addons.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListGroups**
> []SystemInsightsGroups SysteminsightsListGroups(ctx, optional)
List System Insights Groups

Valid filter fields are `jc_system_id` and `groupname`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsGroups**](system-insights-groups.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListGroups_0**
> []SystemInsightsGroups SysteminsightsListGroups_0(ctx, jcSystemId, optional)
List System Insights System Groups

Valid filter fields are `groupname`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jcSystemId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jcSystemId** | **string**|  | 
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsGroups**](system-insights-groups.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListInterfaceAddresses**
> []SystemInsightsInterfaceAddresses SysteminsightsListInterfaceAddresses(ctx, optional)
List System Insights Interface Addresses

Valid filter fields are `jc_system_id` and `address`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsInterfaceAddresses**](system-insights-interface-addresses.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListInterfaceAddresses_0**
> []SystemInsightsInterfaceAddresses SysteminsightsListInterfaceAddresses_0(ctx, jcSystemId, optional)
List System Insights System Interface Addresses

Valid filter fields are `address`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jcSystemId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jcSystemId** | **string**|  | 
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsInterfaceAddresses**](system-insights-interface-addresses.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListMounts**
> []SystemInsightsMounts SysteminsightsListMounts(ctx, optional)
List System Insights Mounts

Valid filter fields are `jc_system_id` and `path`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsMounts**](system-insights-mounts.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListMounts_0**
> []SystemInsightsMounts SysteminsightsListMounts_0(ctx, jcSystemId, optional)
List System Insights System Mounts

Valid filter fields are `path`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jcSystemId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jcSystemId** | **string**|  | 
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsMounts**](system-insights-mounts.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListOsVersion**
> []SystemInsightsOsVersion SysteminsightsListOsVersion(ctx, jcSystemId, optional)
List System Insights System OS Version

Valid filter fields are `version`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jcSystemId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jcSystemId** | **string**|  | 
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsOsVersion**](system-insights-os-version.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListOsVersion_0**
> []SystemInsightsOsVersion SysteminsightsListOsVersion_0(ctx, optional)
List System Insights OS Version

Valid filter fields are `jc_system_id` and `version`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsOsVersion**](system-insights-os-version.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListSafariExtensions**
> []SystemInsightsSafariExtensions SysteminsightsListSafariExtensions(ctx, jcSystemId, optional)
List System Insights System Safari Extensions

Valid filter fields are `name`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jcSystemId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jcSystemId** | **string**|  | 
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsSafariExtensions**](system-insights-safari-extensions.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListSafariExtensions_0**
> []SystemInsightsSafariExtensions SysteminsightsListSafariExtensions_0(ctx, optional)
List System Insights Safari Extensions

Valid filter fields are `jc_system_id` and `name`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsSafariExtensions**](system-insights-safari-extensions.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListSystemInfo**
> []SystemInsightsSystemInfo SysteminsightsListSystemInfo(ctx, optional)
List System Insights System Info

Valid filter fields are `jc_system_id` and `cpu_subtype`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsSystemInfo**](system-insights-system-info.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListSystemInfo_0**
> []SystemInsightsSystemInfo SysteminsightsListSystemInfo_0(ctx, jcSystemId, optional)
List System Insights System System Info

Valid filter fields are `cpu_subtype`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jcSystemId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jcSystemId** | **string**|  | 
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsSystemInfo**](system-insights-system-info.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListUsers**
> []SystemInsightsUsers SysteminsightsListUsers(ctx, optional)
List System Insights Users

Valid filter fields are `jc_system_id` and `username`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsUsers**](system-insights-users.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListUsers_0**
> []SystemInsightsUsers SysteminsightsListUsers_0(ctx, jcSystemId, optional)
List System Insights System Users

Valid filter fields are `username`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jcSystemId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jcSystemId** | **string**|  | 
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsUsers**](system-insights-users.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

