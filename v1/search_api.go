/*
 * JumpCloud APIs
 *
 *  JumpCloud's V1 API. This set of endpoints allows JumpCloud customers to manage commands, systems, & system users.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v1

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
)

// Linger please
var (
	_ context.Context
)

type SearchApiService service


/* SearchApiService Search Organizations
 This endpoint will return Organization data based on your search parameters. This endpoint WILL NOT allow you to add a new Organization.  You can use the supported parameters and pass those in the body of request.   The parameters must be passed as Content-Type application/json.   #### Sample Request &#x60;&#x60;&#x60; curl -X POST https://console.jumpcloud.com/api/search/organizations \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{   \&quot;search\&quot;:{     \&quot;fields\&quot; : [\&quot;settings.name\&quot;],     \&quot;searchTerm\&quot;: \&quot;Second\&quot;     },   \&quot;fields\&quot;: [\&quot;_id\&quot;, \&quot;displayName\&quot;, \&quot;logoUrl\&quot;],   \&quot;limit\&quot; : 0,   \&quot;skip\&quot; : 0 }&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param contentType 
 @param accept 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (Search) 
     @param "fields" (string) Use a space seperated string of field parameters to include the data in the response. If omitted the default list of fields will be returned. 
     @param "filter" (string) A filter to apply to the query.
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
 @return Organizationslist*/
func (a *SearchApiService) SearchOrganizationsPost(ctx context.Context, contentType string, accept string, localVarOptionals map[string]interface{}) (Organizationslist,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Organizationslist
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/search/organizations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["fields"], "string", "fields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["filter"], "string", "filter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["fields"].(string); localVarOk {
		localVarQueryParams.Add("fields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].(string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json; charset=utf-8",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["Accept"] = parameterToString(accept, "")
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["body"].(Search); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* SearchApiService Search Systems
 Return Systems in multi-record format allowing for the passing of the &#x60;filter&#x60; and &#x60;searchFilter&#x60; parameters. This WILL NOT allow you to add a new system.  To support advanced filtering you can use the &#x60;filter&#x60; and &#x60;searchFilter&#x60; parameters that can only be passed in the body of POST /api/search/_* routes. The &#x60;filter&#x60; and &#x60;searchFilter&#x60; parameters must be passed as Content-Type application/json.  The &#x60;filter&#x60; parameter is an object with a single property, either &#x60;and&#x60; or &#x60;or&#x60; with the value of the property being an array of query expressions.   This allows you to filter records using the logic of matching ALL or ANY records in the array of query expressions. If the &#x60;and&#x60; or &#x60;or&#x60; are not included the default behavior is to match ALL query expressions.  The &#x60;searchFilter&#x60; parameter allows text searching on supported fields by specifying a &#x60;searchTerm&#x60; and a list of &#x60;fields&#x60; to query on. If any &#x60;field&#x60; has a partial text match on the&#x60;searchTerm&#x60; the record will be returned.   #### Sample Request  Exact search for a list of hostnames &#x60;&#x60;&#x60; curl -X POST https://console.jumpcloud.com/api/search/systems \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{   \&quot;filter\&quot;: {     \&quot;or\&quot;: [       {\&quot;hostname\&quot; : \&quot;my-hostname\&quot;},       {\&quot;hostname\&quot; : \&quot;other-hostname\&quot;}     ]   },   \&quot;fields\&quot; : \&quot;os hostname displayName\&quot; }&#39; &#x60;&#x60;&#x60;  Text search for a hostname or display name &#x60;&#x60;&#x60; curl -X POST https://console.jumpcloud.com/api/search/systems \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{   \&quot;searchFilter\&quot;: {     \&quot;searchTerm\&quot;: \&quot;my-host\&quot;,     \&quot;fields\&quot;: [\&quot;hostname\&quot;, \&quot;displayName\&quot;]   },   \&quot;fields\&quot;: \&quot;os hostname displayName\&quot; }&#39; &#x60;&#x60;&#x60;  Combining &#x60;filter&#x60; and &#x60;searchFilter&#x60; to search for names that match a given OS &#x60;&#x60;&#x60; curl -X POST https://console.jumpcloud.com/api/search/systems \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{   \&quot;searchFilter\&quot;: {     \&quot;searchTerm\&quot;: \&quot;my-host\&quot;,     \&quot;fields\&quot;: [\&quot;hostname\&quot;, \&quot;displayName\&quot;]   },   \&quot;filter\&quot;: {     \&quot;or\&quot;: [       {\&quot;os\&quot; : \&quot;Ubuntu\&quot;},       {\&quot;os\&quot; : \&quot;Mac OS X\&quot;}     ]   },   \&quot;fields\&quot;: \&quot;os hostname displayName\&quot; }&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param contentType 
 @param accept 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (Search) 
     @param "fields" (string) Use a space seperated string of field parameters to include the data in the response. If omitted the default list of fields will be returned. 
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "xOrgId" (string) 
     @param "skip" (int32) The offset into the records to return.
     @param "filter" (string) A filter to apply to the query.
 @return Systemslist*/
func (a *SearchApiService) SearchSystemsPost(ctx context.Context, contentType string, accept string, localVarOptionals map[string]interface{}) (Systemslist,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Systemslist
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/search/systems"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["fields"], "string", "fields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["filter"], "string", "filter"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["fields"].(string); localVarOk {
		localVarQueryParams.Add("fields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].(string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json; charset=utf-8",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xOrgId"].(string); localVarOk {
		localVarHeaderParams["x-org-id"] = parameterToString(localVarTempParam, "")
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["Accept"] = parameterToString(accept, "")
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["body"].(Search); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* SearchApiService Search System Users
 Return System Users in multi-record format allowing for the passing of the &#x60;filter&#x60; and &#x60;searchFilter&#x60; parameters. This WILL NOT allow you to add a new system user.  To support advanced filtering you can use the &#x60;filter&#x60; and &#x60;searchFilter&#x60; parameters that can only be passed in the body of POST /api/search/_* routes. The &#x60;filter&#x60; and &#x60;searchFilter&#x60; parameters must be passed as Content-Type application/json.  The &#x60;filter&#x60; parameter is an object with a single property, either &#x60;and&#x60; or &#x60;or&#x60; with the value of the property being an array of query expressions.   This allows you to filter records using the logic of matching ALL or ANY records in the array of query expressions. If the &#x60;and&#x60; or &#x60;or&#x60; are not included the default behavior is to match ALL query expressions.  The &#x60;searchFilter&#x60; parameter allows text searching on supported fields by specifying a &#x60;searchTerm&#x60; and a list of &#x60;fields&#x60; to query on. If any &#x60;field&#x60; has a partial text match on the&#x60;searchTerm&#x60; the record will be returned.   #### Sample Request  Exact search for a list of system users in a department &#x60;&#x60;&#x60; curl -X POST https://console.jumpcloud.com/api/search/systemusers \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{   \&quot;filter\&quot; : [{\&quot;department\&quot; : \&quot;IT\&quot;}],   \&quot;fields\&quot; : \&quot;email username sudo\&quot; }&#39; &#x60;&#x60;&#x60;  Text search for system users with and email on a domain &#x60;&#x60;&#x60; curl -X POST https://console.jumpcloud.com/api/search/systemusers \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{   \&quot;searchFilter\&quot; : {     \&quot;searchTerm\&quot;: \&quot;@jumpcloud.com\&quot;,     \&quot;fields\&quot;: [\&quot;email\&quot;]   },   \&quot;fields\&quot; : \&quot;email username sudo\&quot; }&#39; &#x60;&#x60;&#x60;  Combining &#x60;filter&#x60; and &#x60;searchFilter&#x60; to text search for system users with and email on a domain who are in a list of departments &#x60;&#x60;&#x60; curl -X POST https://console.jumpcloud.com/api/search/systemusers \\   -H &#39;Accept: application/json&#39; \\   -H &#39;Content-Type: application/json&#39; \\   -H &#39;x-api-key: {API_KEY}&#39; \\   -d &#39;{   \&quot;searchFilter\&quot;: {     \&quot;searchTerm\&quot;: \&quot;@jumpcloud.com\&quot;,     \&quot;fields\&quot;: [\&quot;email\&quot;]   },   \&quot;filter\&quot;: {     \&quot;or\&quot;: [       {\&quot;department\&quot; : \&quot;IT\&quot;},       {\&quot;department\&quot; : \&quot;Sales\&quot;}     ]   },   \&quot;fields\&quot; : \&quot;email username sudo\&quot; }&#39; &#x60;&#x60;&#x60;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param contentType 
 @param accept 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (Search) 
     @param "fields" (string) Use a space seperated string of field parameters to include the data in the response. If omitted the default list of fields will be returned. 
     @param "filter" (string) A filter to apply to the query.
     @param "limit" (int32) The number of records to return at once. Limited to 100.
     @param "skip" (int32) The offset into the records to return.
     @param "xOrgId" (string) 
 @return Systemuserslist*/
func (a *SearchApiService) SearchSystemusersPost(ctx context.Context, contentType string, accept string, localVarOptionals map[string]interface{}) (Systemuserslist,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Systemuserslist
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/search/systemusers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["fields"], "string", "fields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["filter"], "string", "filter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skip"], "int32", "skip"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xOrgId"], "string", "xOrgId"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["fields"].(string); localVarOk {
		localVarQueryParams.Add("fields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].(string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skip"].(int32); localVarOk {
		localVarQueryParams.Add("skip", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json; charset=utf-8",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	localVarHeaderParams["Accept"] = parameterToString(accept, "")
	if localVarTempParam, localVarOk := localVarOptionals["xOrgId"].(string); localVarOk {
		localVarHeaderParams["x-org-id"] = parameterToString(localVarTempParam, "")
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["body"].(Search); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

