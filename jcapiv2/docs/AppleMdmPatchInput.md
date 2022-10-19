# AppleMdmPatchInput

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ades** | [***Ades**](ADES.md) |  | [optional] [default to null]
**AllowMobileUserEnrollment** | **bool** | A toggle to allow mobile device enrollment for an organization. | [optional] [default to null]
**AppleSignedCert** | **string** | A signed certificate obtained from Apple after providing Apple with the plist file provided on POST. | [optional] [default to null]
**DefaultIosUserEnrollmentDeviceGroupID** | **string** | ObjectId uniquely identifying the MDM default iOS user enrollment device group. | [optional] [default to null]
**DefaultSystemGroupID** | **string** | ObjectId uniquely identifying the MDM default System Group. | [optional] [default to null]
**Dep** | [***Dep**](DEP.md) |  | [optional] [default to null]
**EncryptedDepServerToken** | **string** | The S/MIME encoded DEP Server Token returned by Apple Business Manager when creating an MDM instance. | [optional] [default to null]
**Name** | **string** | A new name for the Apple MDM configuration. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

