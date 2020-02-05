# \SystemInsightsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SysteminsightsListApps**](SystemInsightsApi.md#SysteminsightsListApps) | **Get** /systeminsights/apps | List System Insights Apps
[**SysteminsightsListBitlockerInfo**](SystemInsightsApi.md#SysteminsightsListBitlockerInfo) | **Get** /systeminsights/bitlocker_info | List System Insights Bitlocker Info
[**SysteminsightsListBrowserPlugins**](SystemInsightsApi.md#SysteminsightsListBrowserPlugins) | **Get** /systeminsights/browser_plugins | List System Insights Browser Plugins
[**SysteminsightsListChromeExtensions**](SystemInsightsApi.md#SysteminsightsListChromeExtensions) | **Get** /systeminsights/chrome_extensions | List System Insights Chrome Extensions
[**SysteminsightsListDiskEncryption**](SystemInsightsApi.md#SysteminsightsListDiskEncryption) | **Get** /systeminsights/disk_encryption | List System Insights Disk Encryption
[**SysteminsightsListDiskInfo**](SystemInsightsApi.md#SysteminsightsListDiskInfo) | **Get** /systeminsights/disk_info | List System Insights Disk Info
[**SysteminsightsListEtcHosts**](SystemInsightsApi.md#SysteminsightsListEtcHosts) | **Get** /systeminsights/etc_hosts | List System Insights Etc Hosts
[**SysteminsightsListFirefoxAddons**](SystemInsightsApi.md#SysteminsightsListFirefoxAddons) | **Get** /systeminsights/firefox_addons | List System Insights Firefox Addons
[**SysteminsightsListGroups**](SystemInsightsApi.md#SysteminsightsListGroups) | **Get** /systeminsights/groups | List System Insights Groups
[**SysteminsightsListInterfaceAddresses**](SystemInsightsApi.md#SysteminsightsListInterfaceAddresses) | **Get** /systeminsights/interface_addresses | List System Insights Interface Addresses
[**SysteminsightsListKernelInfo**](SystemInsightsApi.md#SysteminsightsListKernelInfo) | **Get** /systeminsights/kernel_info | List System Insights Kernel Info
[**SysteminsightsListLogicalDrives**](SystemInsightsApi.md#SysteminsightsListLogicalDrives) | **Get** /systeminsights/logical_drives | List System Insights Logical Drives
[**SysteminsightsListMounts**](SystemInsightsApi.md#SysteminsightsListMounts) | **Get** /systeminsights/mounts | List System Insights Mounts
[**SysteminsightsListOsVersion**](SystemInsightsApi.md#SysteminsightsListOsVersion) | **Get** /systeminsights/os_version | List System Insights OS Version
[**SysteminsightsListPatches**](SystemInsightsApi.md#SysteminsightsListPatches) | **Get** /systeminsights/patches | List System Insights Patches
[**SysteminsightsListPrograms**](SystemInsightsApi.md#SysteminsightsListPrograms) | **Get** /systeminsights/programs | List System Insights Programs
[**SysteminsightsListSafariExtensions**](SystemInsightsApi.md#SysteminsightsListSafariExtensions) | **Get** /systeminsights/safari_extensions | List System Insights Safari Extensions
[**SysteminsightsListSystemControls**](SystemInsightsApi.md#SysteminsightsListSystemControls) | **Get** /systeminsights/system_controls | List System Insights System Control
[**SysteminsightsListSystemInfo**](SystemInsightsApi.md#SysteminsightsListSystemInfo) | **Get** /systeminsights/system_info | List System Insights System Info
[**SysteminsightsListUptime**](SystemInsightsApi.md#SysteminsightsListUptime) | **Get** /systeminsights/uptime | List System Insights Uptime
[**SysteminsightsListUsers**](SystemInsightsApi.md#SysteminsightsListUsers) | **Get** /systeminsights/users | List System Insights Users


# **SysteminsightsListApps**
> []SystemInsightsApps SysteminsightsListApps(ctx, contentType, accept, optional)
List System Insights Apps

Valid filter fields are `system_id` and `bundle_name`.

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
 **limit** | **int32**|  | [default to 10]
 **xOrgId** | **string**|  | [default to ]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 

### Return type

[**[]SystemInsightsApps**](system-insights-apps.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListBitlockerInfo**
> []SystemInsightsBitlockerInfo SysteminsightsListBitlockerInfo(ctx, contentType, accept, optional)
List System Insights Bitlocker Info

Valid filter fields are `system_id` and `protection_status`.

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
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]SystemInsightsBitlockerInfo**](system-insights-bitlocker-info.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListBrowserPlugins**
> []SystemInsightsBrowserPlugins SysteminsightsListBrowserPlugins(ctx, contentType, accept, optional)
List System Insights Browser Plugins

Valid filter fields are `system_id` and `name`.

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
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]SystemInsightsBrowserPlugins**](system-insights-browser-plugins.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListChromeExtensions**
> []SystemInsightsChromeExtensions SysteminsightsListChromeExtensions(ctx, contentType, accept, optional)
List System Insights Chrome Extensions

Valid filter fields are `system_id` and `name`.

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
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]SystemInsightsChromeExtensions**](system-insights-chrome-extensions.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListDiskEncryption**
> []SystemInsightsDiskEncryption SysteminsightsListDiskEncryption(ctx, contentType, accept, optional)
List System Insights Disk Encryption

Valid filter fields are `system_id` and `encryption_status`.

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
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]SystemInsightsDiskEncryption**](system-insights-disk-encryption.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListDiskInfo**
> []SystemInsightsDiskInfo SysteminsightsListDiskInfo(ctx, contentType, accept, optional)
List System Insights Disk Info

Valid filter fields are `system_id` and `disk_index`.

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
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]SystemInsightsDiskInfo**](system-insights-disk-info.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListEtcHosts**
> []SystemInsightsEtcHosts SysteminsightsListEtcHosts(ctx, contentType, accept, optional)
List System Insights Etc Hosts

Valid filter fields are `system_id` and `address`.

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
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]SystemInsightsEtcHosts**](system-insights-etc-hosts.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListFirefoxAddons**
> []SystemInsightsFirefoxAddons SysteminsightsListFirefoxAddons(ctx, contentType, accept, optional)
List System Insights Firefox Addons

Valid filter fields are `system_id` and `name`.

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
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]SystemInsightsFirefoxAddons**](system-insights-firefox-addons.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListGroups**
> []SystemInsightsGroups SysteminsightsListGroups(ctx, contentType, accept, optional)
List System Insights Groups

Valid filter fields are `system_id` and `groupname`.

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
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]SystemInsightsGroups**](system-insights-groups.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListInterfaceAddresses**
> []SystemInsightsInterfaceAddresses SysteminsightsListInterfaceAddresses(ctx, contentType, accept, optional)
List System Insights Interface Addresses

Valid filter fields are `system_id` and `address`.

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
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]SystemInsightsInterfaceAddresses**](system-insights-interface-addresses.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListKernelInfo**
> []SystemInsightsKernelInfo SysteminsightsListKernelInfo(ctx, contentType, accept, optional)
List System Insights Kernel Info

Valid filter fields are `system_id` and `version`.

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
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]SystemInsightsKernelInfo**](system-insights-kernel-info.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListLogicalDrives**
> []SystemInsightsLogicalDrvies SysteminsightsListLogicalDrives(ctx, contentType, accept, optional)
List System Insights Logical Drives

Valid filter fields are `system_id` and `device_id`.

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
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]SystemInsightsLogicalDrvies**](system-insights-logical-drvies.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListMounts**
> []SystemInsightsMounts SysteminsightsListMounts(ctx, contentType, accept, optional)
List System Insights Mounts

Valid filter fields are `system_id` and `path`.

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
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]SystemInsightsMounts**](system-insights-mounts.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListOsVersion**
> []SystemInsightsOsVersion SysteminsightsListOsVersion(ctx, contentType, accept, optional)
List System Insights OS Version

Valid filter fields are `system_id` and `version`.

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
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]SystemInsightsOsVersion**](system-insights-os-version.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListPatches**
> []SystemInsightsPatches SysteminsightsListPatches(ctx, contentType, accept, optional)
List System Insights Patches

Valid filter fields are `system_id` and `hotfix_id`.

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
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]SystemInsightsPatches**](system-insights-patches.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListPrograms**
> []SystemInsightsPrograms SysteminsightsListPrograms(ctx, contentType, accept, optional)
List System Insights Programs

Valid filter fields are `system_id` and `name`.

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
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]SystemInsightsPrograms**](system-insights-programs.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListSafariExtensions**
> []SystemInsightsSafariExtensions SysteminsightsListSafariExtensions(ctx, contentType, accept, optional)
List System Insights Safari Extensions

Valid filter fields are `system_id` and `name`.

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
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]SystemInsightsSafariExtensions**](system-insights-safari-extensions.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListSystemControls**
> []SystemInsightsSystemControls SysteminsightsListSystemControls(ctx, contentType, accept, optional)
List System Insights System Control

Valid filter fields are `system_id` and `name`.

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
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]SystemInsightsSystemControls**](system-insights-system-controls.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListSystemInfo**
> []SystemInsightsSystemInfo SysteminsightsListSystemInfo(ctx, contentType, accept, optional)
List System Insights System Info

Valid filter fields are `system_id` and `cpu_subtype`.

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
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]SystemInsightsSystemInfo**](system-insights-system-info.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListUptime**
> []SystemInsightsUptime SysteminsightsListUptime(ctx, contentType, accept, optional)
List System Insights Uptime

Valid filter fields are `system_id` and `days`.

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
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]SystemInsightsUptime**](system-insights-uptime.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminsightsListUsers**
> []SystemInsightsUsers SysteminsightsListUsers(ctx, contentType, accept, optional)
List System Insights Users

Valid filter fields are `system_id` and `username`.

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
 **limit** | **int32**|  | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **filter** | [**[]string**](string.md)| Supported operators are: eq | 
 **xOrgId** | **string**|  | [default to ]

### Return type

[**[]SystemInsightsUsers**](system-insights-users.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

