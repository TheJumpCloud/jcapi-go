## JCAPI-Go

### Description

This repository contains the Go client code for the JumpCloud API v1 and v2.
This code is automatically generated using Swagger Codegen. It also provides
the tools to generate the client code from the API YAML files, using Swagger
Codegen. For detailed instructions on how to generate the code, see the
[Contributing](CONTRIBUTING.md) section.

To access our old API v1 Go client, please refer to
[JCAPI](https://github.com/TheJumpCloud/jcapi). However this code repository
is not up to date with current API functionality.

### Authentication and Authorization

All endpoints support authentication via API key: see the
[Authentication & Authorization](https://docs.jumpcloud.com/2.0/authentication-and-authorization/authentication-and-authorization-overview)
section in our API documentation.

Some systems endpoints (in both API v1 and v2) also support
[System Context Authorization](https://docs.jumpcloud.com/2.0/authentication-and-authorization/system-context)
which allows an individual system to manage its information and resource
associations.

### Usage Examples

For more detailed instructions, refer to each API version's respective README
file ([README for API v1](v1/README.md) and [README for API v2](v2/README.md))
and the generated documentation under each folder.

#### API v2 Example

```go
package main

import (
	"context"
	"fmt"

	jcapiv2 "github.com/TheJumpCloud/jcapi-go/jcapiv2"
	"github.com/antihax/optional"
)

func main() {
	apiKey := "YOUR_API_KEY"
	XOrgId := "YOUR_ORG_ID"
	userGroupID := "YOUR_GROUP_ID"

	// Instantiate the API client
	client := jcapiv2.NewAPIClient(jcapiv2.NewConfiguration())

	// Set up the API key via context
	auth := context.WithValue(context.TODO(), jcapiv2.ContextAPIKey, jcapiv2.APIKey{
		Key: apiKey,
	})

	// Make an API call to retrieve a specific user group by ID
	userGroup, res, err := client.UserGroupsApi.GroupsUserGet(auth, userGroupID, &jcapiv2.UserGroupsApiGroupsUserGetOpts{XOrgId: optional.NewString(XOrgId)})
	if err != nil {
		fmt.Printf("Error retrieving user group %s: %s - response = %+v\n", userGroupID, err, res)
	} else {
		fmt.Printf("Details for User group %s: %+v\n", userGroupID, userGroup)
	}
}

```

#### System Context Authorization Example

```go
package main

import (
	"context"
	"fmt"

	jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"
)

func main() {
	// Set headers for System Context Authorization. For detailed instructions on
	// how to generate these headers, refer to:
	// https://docs.jumpcloud.com/2.0/authentication-and-authorization/system-context
	systemID := "YOUR_SYSTEM_ID"
	systemDate := "YOUR_SYSTEM_DATE" // The current date on the system, e.g. "Tue, 10 Nov 2009 23:00:00 GMT"
	systemSignature := "YOUR_SYSTEM_SIGNATURE"
	systemContextAuth := fmt.Sprintf(`Signature keyId="system/%s",headers="request-line date",algorithm="rsa-sha256",signature="%s"`, systemID, systemSignature)

	contentType := "application/json"
	accept := "application/json"

	// Instantiate the API client
	client := jcapiv2.NewAPIClient(jcapiv2.NewConfiguration())

	// Add date and authorization to the list of optional parameters
	optParams := map[string]interface{}{
		"date":          systemDate,
		"authorization": systemContextAuth,
	}

	// List the system groups this system is a member of using the
	// System Context Authorization parameters. Since we authenticate via the
	// parameters, there is no need to set the API key in the context.
	groups, res, err := client.SystemsApi.GraphSystemMemberOf(context.TODO(), systemID, contentType, accept, optParams)
	if err != nil {
		fmt.Printf("Error retrieving system groups for system %s: %s - response = %+v\n", systemID, err, res)
		return
	}

	// Print the system groups we just retrieved
	for _, group := range groups {
		fmt.Printf("System group ID %s\n", group.Id)
	}
}

```
