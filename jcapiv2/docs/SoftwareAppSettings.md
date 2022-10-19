# SoftwareAppSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUpdateDelay** | **bool** |  | [optional] [default to false]
**AppleVpp** | [***SoftwareAppAppleVpp**](software-app-apple-vpp.md) |  | [optional] [default to null]
**AssetKind** | **string** | The manifest asset kind (ex: software). | [optional] [default to null]
**AssetSha256Size** | **int32** | The incremental size to use for summing the package as it is downloaded. | [optional] [default to null]
**AssetSha256Strings** | **[]string** | The array of checksums, one each for the hash size up to the total size of the package. | [optional] [default to null]
**AutoUpdate** | **bool** |  | [optional] [default to false]
**Description** | **string** | The software app description. | [optional] [default to null]
**DesiredState** | **string** | State of Install or Uninstall | [optional] [default to null]
**Location** | **string** | Repository where the app is located within the package manager | [optional] [default to null]
**LocationObjectId** | **string** | ID of the repository where the app is located within the package manager | [optional] [default to null]
**PackageId** | **string** |  | [optional] [default to null]
**PackageKind** | **string** | The package manifest kind (ex: software-package). | [optional] [default to null]
**PackageManager** | **string** | App store serving the app: APPLE_VPP, CHOCOLATEY, etc. | [optional] [default to null]
**PackageSubtitle** | **string** | The package manifest subtitle. | [optional] [default to null]
**PackageVersion** | **string** | The package manifest version. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

