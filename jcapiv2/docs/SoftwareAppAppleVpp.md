# SoftwareAppAppleVpp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppConfiguration** | **string** | Text sent to configure the application, the text should be a valid plist.  Returned only by &#x27;GET /softwareapps/{id}&#x27;. | [optional] [default to null]
**AssignedLicenses** | **int32** |  | [optional] [default to null]
**AvailableLicenses** | **int32** |  | [optional] [default to null]
**Details** | [***interface{}**](interface{}.md) | App details returned by iTunes API. See example. The properties in this field are out of our control and we cannot guarantee consistency, so it should be checked by the client and manage the details accordingly. | [optional] [default to null]
**IsConfigEnabled** | **bool** | Denotes if configuration has been enabled for the application.  Returned only by &#x27;&#x27;GET /softwareapps/{id}&#x27;&#x27;. | [optional] [default to null]
**SupportedDeviceFamilies** | **[]string** | The supported device families for this VPP Application. | [optional] [default to null]
**TotalLicenses** | **int32** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

