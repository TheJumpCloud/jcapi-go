# {{classname}}

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionsGet**](SubscriptionsApi.md#SubscriptionsGet) | **Get** /subscriptions | Lists all the Pricing &amp; Packaging Subscriptions

# **SubscriptionsGet**
> []Subscription SubscriptionsGet(ctx, )
Lists all the Pricing & Packaging Subscriptions

This endpoint returns all pricing & packaging subscriptions.  ##### Sample Request  ```  curl -X GET  https://console.jumpcloud.com/api/v2/subscriptions \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}'   ```

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Subscription**](subscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

