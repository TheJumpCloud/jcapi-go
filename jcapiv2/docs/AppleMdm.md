# AppleMdm

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ades** | [***Ades**](ADES.md) |  | [optional] [default to null]
**AllowMobileUserEnrollment** | **bool** | A toggle to allow mobile device enrollment for an organization. | [optional] [default to null]
**ApnsCertExpiry** | **string** | The expiration date and time for the APNS Certificate. | [optional] [default to null]
**ApnsPushTopic** | **string** | The push topic assigned to this enrollment by Apple after uploading the Signed CSR plist. | [optional] [default to null]
**DefaultIosUserEnrollmentDeviceGroupID** | **string** | ObjectId uniquely identifying the MDM default iOS user enrollment device group. | [optional] [default to null]
**DefaultSystemGroupID** | **string** | ObjectId uniquely identifying the MDM default System Group. | [optional] [default to null]
**Dep** | [***Dep**](DEP.md) |  | [optional] [default to null]
**DepAccessTokenExpiry** | **string** | The expiration date and time for the DEP Access Token. This aligns with the DEP Server Token State. | [optional] [default to null]
**DepServerTokenState** | **string** | The state of the dep server token, presence and expiry. | [optional] [default to null]
**Id** | **string** | ObjectId uniquely identifying an MDM Enrollment, | [default to null]
**Name** | **string** | A friendly name to identify this enrollment.  Not required to be unique. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

