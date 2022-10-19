# {{classname}}

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplemdmsCsrget**](AppleMDMApi.md#ApplemdmsCsrget) | **Get** /applemdms/{apple_mdm_id}/csr | Get Apple MDM CSR Plist
[**ApplemdmsDelete**](AppleMDMApi.md#ApplemdmsDelete) | **Delete** /applemdms/{id} | Delete an Apple MDM
[**ApplemdmsDeletedevice**](AppleMDMApi.md#ApplemdmsDeletedevice) | **Delete** /applemdms/{apple_mdm_id}/devices/{device_id} | Remove an Apple MDM Device&#x27;s Enrollment
[**ApplemdmsDepkeyget**](AppleMDMApi.md#ApplemdmsDepkeyget) | **Get** /applemdms/{apple_mdm_id}/depkey | Get Apple MDM DEP Public Key
[**ApplemdmsDevicesClearActivationLock**](AppleMDMApi.md#ApplemdmsDevicesClearActivationLock) | **Post** /applemdms/{apple_mdm_id}/devices/{device_id}/clearActivationLock | Clears the Activation Lock for a Device
[**ApplemdmsDevicesRefreshActivationLockInformation**](AppleMDMApi.md#ApplemdmsDevicesRefreshActivationLockInformation) | **Post** /applemdms/{apple_mdm_id}/devices/{device_id}/refreshActivationLockInformation | Refresh activation lock information for a device
[**ApplemdmsDeviceserase**](AppleMDMApi.md#ApplemdmsDeviceserase) | **Post** /applemdms/{apple_mdm_id}/devices/{device_id}/erase | Erase Device
[**ApplemdmsDeviceslist**](AppleMDMApi.md#ApplemdmsDeviceslist) | **Get** /applemdms/{apple_mdm_id}/devices | List AppleMDM Devices
[**ApplemdmsDeviceslock**](AppleMDMApi.md#ApplemdmsDeviceslock) | **Post** /applemdms/{apple_mdm_id}/devices/{device_id}/lock | Lock Device
[**ApplemdmsDevicesrestart**](AppleMDMApi.md#ApplemdmsDevicesrestart) | **Post** /applemdms/{apple_mdm_id}/devices/{device_id}/restart | Restart Device
[**ApplemdmsDevicesshutdown**](AppleMDMApi.md#ApplemdmsDevicesshutdown) | **Post** /applemdms/{apple_mdm_id}/devices/{device_id}/shutdown | Shut Down Device
[**ApplemdmsEnrollmentprofilesget**](AppleMDMApi.md#ApplemdmsEnrollmentprofilesget) | **Get** /applemdms/{apple_mdm_id}/enrollmentprofiles/{id} | Get an Apple MDM Enrollment Profile
[**ApplemdmsEnrollmentprofileslist**](AppleMDMApi.md#ApplemdmsEnrollmentprofileslist) | **Get** /applemdms/{apple_mdm_id}/enrollmentprofiles | List Apple MDM Enrollment Profiles
[**ApplemdmsGetdevice**](AppleMDMApi.md#ApplemdmsGetdevice) | **Get** /applemdms/{apple_mdm_id}/devices/{device_id} | Details of an AppleMDM Device
[**ApplemdmsList**](AppleMDMApi.md#ApplemdmsList) | **Get** /applemdms | List Apple MDMs
[**ApplemdmsPut**](AppleMDMApi.md#ApplemdmsPut) | **Put** /applemdms/{id} | Update an Apple MDM
[**ApplemdmsRefreshdepdevices**](AppleMDMApi.md#ApplemdmsRefreshdepdevices) | **Post** /applemdms/{apple_mdm_id}/refreshdepdevices | Refresh DEP Devices

# **ApplemdmsCsrget**
> string ApplemdmsCsrget(ctx, appleMdmId, optional)
Get Apple MDM CSR Plist

Retrieves an Apple MDM signed CSR Plist for an organization.  The user must supply the returned plist to Apple for signing, and then provide the certificate provided by Apple back into the PUT API.  #### Sample Request ```   curl -X GET https://console.jumpcloud.com/api/v2/applemdms/{APPLE_MDM_ID}/csr \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appleMdmId** | **string**|  | 
 **optional** | ***AppleMDMApiApplemdmsCsrgetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppleMDMApiApplemdmsCsrgetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

**string**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplemdmsDelete**
> AppleMdm ApplemdmsDelete(ctx, id, optional)
Delete an Apple MDM

Removes an Apple MDM configuration.  Warning: This is a destructive operation and will remove your Apple Push Certificates.  We will no longer be able to manage your devices and the only recovery option is to re-register all devices into MDM.  #### Sample Request ``` curl -X DELETE https://console.jumpcloud.com/api/v2/applemdms/{id} \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AppleMDMApiApplemdmsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppleMDMApiApplemdmsDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

[**AppleMdm**](AppleMDM.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplemdmsDeletedevice**
> AppleMdmDevice ApplemdmsDeletedevice(ctx, appleMdmId, deviceId, optional)
Remove an Apple MDM Device's Enrollment

Remove a single Apple MDM device from MDM enrollment.  #### Sample Request ```   curl -X DELETE https://console.jumpcloud.com/api/v2/applemdms/{apple_mdm_id}/devices/{device_id} \\   -H 'accept: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appleMdmId** | **string**|  | 
  **deviceId** | **string**|  | 
 **optional** | ***AppleMDMApiApplemdmsDeletedeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppleMDMApiApplemdmsDeletedeviceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

[**AppleMdmDevice**](apple-mdm-device.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplemdmsDepkeyget**
> string ApplemdmsDepkeyget(ctx, appleMdmId, optional)
Get Apple MDM DEP Public Key

Retrieves an Apple MDM DEP Public Key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appleMdmId** | **string**|  | 
 **optional** | ***AppleMDMApiApplemdmsDepkeygetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppleMDMApiApplemdmsDepkeygetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

**string**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-pem-file

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplemdmsDevicesClearActivationLock**
> ApplemdmsDevicesClearActivationLock(ctx, appleMdmId, deviceId, optional)
Clears the Activation Lock for a Device

Clears the activation lock on the specified device.  #### Sample Request ```   curl -X POST https://console.jumpcloud.com/api/v2/applemdms/{apple_mdm_id}/devices/{device_id}/clearActivationLock \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appleMdmId** | **string**|  | 
  **deviceId** | **string**|  | 
 **optional** | ***AppleMDMApiApplemdmsDevicesClearActivationLockOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppleMDMApiApplemdmsDevicesClearActivationLockOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplemdmsDevicesRefreshActivationLockInformation**
> ApplemdmsDevicesRefreshActivationLockInformation(ctx, appleMdmId, deviceId, optional)
Refresh activation lock information for a device

Refreshes the activation lock information for a device  #### Sample Request  ``` curl -X POST https://console.jumpcloud.com/api/v2/applemdms/{apple_mdm_id}/devices/{device_id}/refreshActivationLockInformation \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appleMdmId** | **string**|  | 
  **deviceId** | **string**|  | 
 **optional** | ***AppleMDMApiApplemdmsDevicesRefreshActivationLockInformationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppleMDMApiApplemdmsDevicesRefreshActivationLockInformationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplemdmsDeviceserase**
> ApplemdmsDeviceserase(ctx, appleMdmId, deviceId, optional)
Erase Device

Erases a DEP-enrolled device.  #### Sample Request ```   curl -X POST https://console.jumpcloud.com/api/v2/applemdms/{apple_mdm_id}/devices/{device_id}/erase \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appleMdmId** | **string**|  | 
  **deviceId** | **string**|  | 
 **optional** | ***AppleMDMApiApplemdmsDeviceseraseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppleMDMApiApplemdmsDeviceseraseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdEraseBody**](DeviceIdEraseBody.md)|  | 
 **xOrgId** | **optional.**| Organization identifier that can be obtained from console settings. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplemdmsDeviceslist**
> []AppleMdmDevice ApplemdmsDeviceslist(ctx, appleMdmId, optional)
List AppleMDM Devices

Lists all Apple MDM devices.  The filter and sort queries will allow the following fields: `createdAt` `depRegistered` `enrolled` `id` `osVersion` `serialNumber` `udid`  #### Sample Request ```   curl -X GET https://console.jumpcloud.com/api/v2/applemdms/{apple_mdm_id}/devices \\   -H 'accept: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appleMdmId** | **string**|  | 
 **optional** | ***AppleMDMApiApplemdmsDeviceslistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppleMDMApiApplemdmsDeviceslistOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 
 **skip** | **optional.Int32**| The offset into the records to return. | [default to 0]
 **filter** | [**optional.Interface of []string**](string.md)| A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | 
 **sort** | [**optional.Interface of []string**](string.md)| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 
 **xTotalCount** | **optional.Int32**|  | 

### Return type

[**[]AppleMdmDevice**](apple-mdm-device.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplemdmsDeviceslock**
> ApplemdmsDeviceslock(ctx, appleMdmId, deviceId, optional)
Lock Device

Locks a DEP-enrolled device.  #### Sample Request ```   curl -X POST https://console.jumpcloud.com/api/v2/applemdms/{apple_mdm_id}/devices/{device_id}/lock \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appleMdmId** | **string**|  | 
  **deviceId** | **string**|  | 
 **optional** | ***AppleMDMApiApplemdmsDeviceslockOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppleMDMApiApplemdmsDeviceslockOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdLockBody**](DeviceIdLockBody.md)|  | 
 **xOrgId** | **optional.**| Organization identifier that can be obtained from console settings. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplemdmsDevicesrestart**
> ApplemdmsDevicesrestart(ctx, appleMdmId, deviceId, optional)
Restart Device

Restarts a DEP-enrolled device.  #### Sample Request ```   curl -X POST https://console.jumpcloud.com/api/v2/applemdms/{apple_mdm_id}/devices/{device_id}/restart \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{\"kextPaths\": [\"Path1\", \"Path2\"]}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appleMdmId** | **string**|  | 
  **deviceId** | **string**|  | 
 **optional** | ***AppleMDMApiApplemdmsDevicesrestartOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppleMDMApiApplemdmsDevicesrestartOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdRestartBody**](DeviceIdRestartBody.md)|  | 
 **xOrgId** | **optional.**| Organization identifier that can be obtained from console settings. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplemdmsDevicesshutdown**
> ApplemdmsDevicesshutdown(ctx, appleMdmId, deviceId, optional)
Shut Down Device

Shuts down a DEP-enrolled device.  #### Sample Request ```   curl -X POST https://console.jumpcloud.com/api/v2/applemdms/{apple_mdm_id}/devices/{device_id}/shutdown \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appleMdmId** | **string**|  | 
  **deviceId** | **string**|  | 
 **optional** | ***AppleMDMApiApplemdmsDevicesshutdownOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppleMDMApiApplemdmsDevicesshutdownOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplemdmsEnrollmentprofilesget**
> string ApplemdmsEnrollmentprofilesget(ctx, appleMdmId, id, optional)
Get an Apple MDM Enrollment Profile

Get an enrollment profile  Currently only requesting the mobileconfig is supported.  #### Sample Request  ``` curl https://console.jumpcloud.com/api/v2/applemdms/{APPLE_MDM_ID}/enrollmentprofiles/{ID} \\   -H 'accept: application/x-apple-aspen-config' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appleMdmId** | **string**|  | 
  **id** | **string**|  | 
 **optional** | ***AppleMDMApiApplemdmsEnrollmentprofilesgetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppleMDMApiApplemdmsEnrollmentprofilesgetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

**string**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-apple-aspen-config

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplemdmsEnrollmentprofileslist**
> []AppleMdm ApplemdmsEnrollmentprofileslist(ctx, appleMdmId, optional)
List Apple MDM Enrollment Profiles

Get a list of enrollment profiles for an apple mdm.  Note: currently only one enrollment profile is supported.  #### Sample Request ```  curl https://console.jumpcloud.com/api/v2/applemdms/{APPLE_MDM_ID}/enrollmentprofiles \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appleMdmId** | **string**|  | 
 **optional** | ***AppleMDMApiApplemdmsEnrollmentprofileslistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppleMDMApiApplemdmsEnrollmentprofileslistOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

[**[]AppleMdm**](AppleMDM.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplemdmsGetdevice**
> AppleMdmDevice ApplemdmsGetdevice(ctx, appleMdmId, deviceId, optional)
Details of an AppleMDM Device

Gets a single Apple MDM device.  #### Sample Request ```   curl -X GET https://console.jumpcloud.com/api/v2/applemdms/{apple_mdm_id}/devices/{device_id} \\   -H 'accept: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appleMdmId** | **string**|  | 
  **deviceId** | **string**|  | 
 **optional** | ***AppleMDMApiApplemdmsGetdeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppleMDMApiApplemdmsGetdeviceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

[**AppleMdmDevice**](apple-mdm-device.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplemdmsList**
> []AppleMdm ApplemdmsList(ctx, optional)
List Apple MDMs

Get a list of all Apple MDM configurations.  An empty topic indicates that a signed certificate from Apple has not been provided to the PUT endpoint yet.  Note: currently only one MDM configuration per organization is supported.  #### Sample Request ``` curl https://console.jumpcloud.com/api/v2/applemdms \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppleMDMApiApplemdmsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppleMDMApiApplemdmsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

[**[]AppleMdm**](AppleMDM.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplemdmsPut**
> AppleMdm ApplemdmsPut(ctx, id, optional)
Update an Apple MDM

Updates an Apple MDM configuration.  This endpoint is used to supply JumpCloud with a signed certificate from Apple in order to finalize the setup and allow JumpCloud to manage your devices.  It may also be used to update the DEP Settings.  #### Sample Request ```   curl -X PUT https://console.jumpcloud.com/api/v2/applemdms/{ID} \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{     \"name\": \"MDM name\",     \"appleSignedCert\": \"{CERTIFICATE}\",     \"encryptedDepServerToken\": \"{SERVER_TOKEN}\",     \"dep\": {       \"welcomeScreen\": {         \"title\": \"Welcome\",         \"paragraph\": \"In just a few steps, you will be working securely from your Mac.\",         \"button\": \"continue\",       },     },   }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***AppleMDMApiApplemdmsPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppleMDMApiApplemdmsPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AppleMdmPatchInput**](AppleMdmPatchInput.md)|  | 
 **xOrgId** | **optional.**| Organization identifier that can be obtained from console settings. | 

### Return type

[**AppleMdm**](AppleMDM.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplemdmsRefreshdepdevices**
> ApplemdmsRefreshdepdevices(ctx, appleMdmId, optional)
Refresh DEP Devices

Refreshes the list of devices that a JumpCloud admin has added to their virtual MDM in Apple Business Manager - ABM so that they can be DEP enrolled with JumpCloud.  #### Sample Request ```   curl -X POST https://console.jumpcloud.com/api/v2/applemdms/{apple_mdm_id}/refreshdepdevices \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -d '{}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appleMdmId** | **string**|  | 
 **optional** | ***AppleMDMApiApplemdmsRefreshdepdevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppleMDMApiApplemdmsRefreshdepdevicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **optional.String**| Organization identifier that can be obtained from console settings. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

