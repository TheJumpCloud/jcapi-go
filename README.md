## JCAPI-Go

### Description ###

This repository contains the Go client code for the JumpCloud API v1 & v2. This code is automatically generated using swagger-codegen.
It also provides the tools to generate the client code from the API yaml files, using swagger-codegen.
For detailed instructions on how to generate the code, see the [Contributing](CONTRIBUTING.md) section.

To access our old API v1 Go Client, please refer to [JCAPI](https://github.com/TheJumpCloud/jcapi). However this code repository is not up to date with current API functionality. 


### Usage Examples

```golang
import (
	"fmt"
	jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"
)

func main() {

	contentType := "application/json"
	accept := "application/json"
	userGroupId := "<YOUR_GROUP_ID>"

	// instantiate the API client:
	client := jcapiv2.NewAPIClient(jcapiv2.NewConfiguration())

	// set up the API key via context:
	auth := context.WithValue(context.TODO(), jcapiv2.ContextAPIKey, jcapiv2.APIKey{
		Key: "<YOUR_API_KEY>",
	})

	// make an API call to retrieve a specific user group by id:
	userGroup, res, err := client.UserGroupsApi.GroupsUserGet(auth, userGroupId, contentType, accept, nil)

	if err != nil {
		fmt.Printf("Error retrieving user group %s: %s - response = %+v\n", userGroupId, err, res)
	} else {
		fmt.Printf("Details for User group %s: %+v\n", userGroupId, userGroup)
	}
}
```

System Context Authorization example:

```golang
import (
	"context"
	"fmt"
	jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"
)

func main() {

	// instantiate the API client:
	client := jcapiv2.NewAPIClient(jcapiv2.NewConfiguration())

	contentType := "application/json"
	accept := "application/json"

	systemId := "<YOUR_SYSTEM_ID>"

	// set headers for the System Context Authorization:
	// for detailed instructions on how to generate these headers,
	// refer to: https://docs.jumpcloud.com/2.0/authentication-and-authorization/system-context
	sysContextAuth := `Signature keyId="system/<YOUR_SYSTEM_ID>",headers="request-line date",algorithm="rsa-sha256",signature="<YOUR_SYSTEM_SIGNATURE>"`
	sysContextDate := "Thu, 19 Oct 2017 17:27:57 GMT" // the current date on the system

	// add date and authorization to the list of optional parameters:
	optParams := map[string]interface{}{
		"date":          sysContextDate,
		"authorization": sysContextAuth,
	}

	// list the system groups this system is a member of using the System Context Authorization parameters:
	// since we authenticate via system context parameters, no need to set the API key in the context
	groups, res, err := client.SystemsApi.GraphSystemMemberOf(context.TODO(), systemId, contentType, accept, optParams)

	if err != nil {
		fmt.Printf("Error retrieving system groups for system %s: %s - response = %+v\n", systemId, err, res)
		return
	}
	// print the system groups we just retrieved:
	for _, group := range groups {
		fmt.Printf("System group ID %s\n", group.Id)
	}

	return
}

```
