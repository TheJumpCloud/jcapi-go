# ApplicationtemplateOidc

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantTypes** | **[]string** | The grant types allowed. Currently only authorization_code is allowed. | [optional] [default to null]
**RedirectUris** | **[]string** | List of allowed redirectUris | [optional] [default to null]
**SsoUrl** | **string** | The relying party url to trigger an oidc login. | [optional] [default to null]
**TokenEndpointAuthMethod** | **string** | Method that the client uses to authenticate when requesting a token. If &#x27;none&#x27;, then the client must use PKCE. If &#x27;client_secret_post&#x27;, then the secret is passed in the post body when requesting the token. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

