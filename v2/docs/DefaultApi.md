# \DefaultApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JcEnrollmentProfilesDelete**](DefaultApi.md#JcEnrollmentProfilesDelete) | **Delete** /enrollmentprofiles/{enrollment_profile_id} | Delete Enrollment Profile
[**JcEnrollmentProfilesGet**](DefaultApi.md#JcEnrollmentProfilesGet) | **Get** /enrollmentprofiles/{enrollment_profile_id} | Get Enrollment Profile
[**JcEnrollmentProfilesList**](DefaultApi.md#JcEnrollmentProfilesList) | **Get** /enrollmentprofiles | List Enrollment Profiles
[**JcEnrollmentProfilesPost**](DefaultApi.md#JcEnrollmentProfilesPost) | **Post** /enrollmentprofiles | Create new Enrollment Profile
[**JcEnrollmentProfilesPut**](DefaultApi.md#JcEnrollmentProfilesPut) | **Put** /enrollmentprofiles/{enrollment_profile_id} | Update Enrollment Profile


# **JcEnrollmentProfilesDelete**
> JcEnrollmentProfile JcEnrollmentProfilesDelete(ctx, enrollmentProfileId)
Delete Enrollment Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **enrollmentProfileId** | **string**|  | 

### Return type

[**JcEnrollmentProfile**](jc-enrollment-profile.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JcEnrollmentProfilesGet**
> JcEnrollmentProfilesGet(ctx, enrollmentProfileId, optional)
Get Enrollment Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **enrollmentProfileId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enrollmentProfileId** | **string**|  | 
 **body** | [**JcEnrollmentProfile**](JcEnrollmentProfile.md)|  | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JcEnrollmentProfilesList**
> []JcEnrollmentProfile JcEnrollmentProfilesList(ctx, optional)
List Enrollment Profiles

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

### Return type

[**[]JcEnrollmentProfile**](jc-enrollment-profile.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JcEnrollmentProfilesPost**
> JcEnrollmentProfile JcEnrollmentProfilesPost(ctx, optional)
Create new Enrollment Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Body1**](Body1.md)|  | 

### Return type

[**JcEnrollmentProfile**](jc-enrollment-profile.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JcEnrollmentProfilesPut**
> JcEnrollmentProfile JcEnrollmentProfilesPut(ctx, enrollmentProfileId, optional)
Update Enrollment Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **enrollmentProfileId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enrollmentProfileId** | **string**|  | 
 **body** | [**Body**](Body.md)|  | 

### Return type

[**JcEnrollmentProfile**](jc-enrollment-profile.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

