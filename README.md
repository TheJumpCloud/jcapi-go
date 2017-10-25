## JCAPI-Go

### Description ###

This repository contains the Go client code for the JumpCloud API v2. This code is automatically generated using swagger-codegen.
It also provides the tools to generate the client code from the API yaml files, using swagger-codegen.
For detailed instructions on how to generate the code, see the [Contributing](CONTRIBUTING.md) section.

For the API v1 Go Client, please refer to [JCAPI](https://github.com/TheJumpCloud/jcapi).


### Usage Examples

```
import (
  jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"
  "fmt"
)

func main() {

  contentType := "application/json"
  accept := "application/json"
  userGroupId := "<YOUR_GROUP_ID>"

  // instantiate the API object for the group of endpoints you need to use
  // for instance for user groups API:
  userGroupsAPI := jcapiv2.NewUserGroupsApi()

  // set up the API key:
  userGroupsAPI.Configuration.APIKey["x-api-key"] = "<YOUR_API_KEY>"

  // make an API call to retrieve a specific user group by id:
  userGroup, _, err := userGroupsAPI.GroupsUserGet(userGroupId, contentType, accept)

  if err != nil {
    fmt.Printf("Error retrieving user group %s: %s\n", userGroupId, err)
  } else {
    fmt.Printf("Details for User group %s: %+v\n", userGroupId, userGroup)
  }

}
```

System Context Authorization example:

```
import (
  jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"
  "fmt"
)

func main() {

  contentType := "application/json"
  accept := "application/json"

  systemId := "<YOUR_SYSTEM_ID>"

  // set headers for the System Context Authorization:
  // for detailed instructions on how to generate these headers,
  // refer to: https://docs.jumpcloud.com/2.0/authentication-and-authorization/system-context
  sysContextAuth := `Signature keyId="system/<YOUR_SYSTEM_ID>",headers="request-line date",algorithm="rsa-sha256",signature="<YOUR_SYSTEM_SIGNATURE>"`
  sysContextDate := "Thu, 19 Oct 2017 17:27:57 GMT" // the current date on the system

  // instantiate object for the Systems API:
  systemsAPI := jcapiv2.NewSystemsApi()

  // list the system groups this system is a member of using the System Context Authorization parameters:
  groups, _, err := systemsAPI.GraphSystemMemberOf(systemId, contentType, accept, 100, 100, sysContextDate, sysContextAuth)

  if err != nil {
    fmt.Printf("Error retrieving system groups for system %s: %s\n", systemId, err)
  } else {
    // print the system groups we just retrieved:
    for _, group : range groups {
      fmt.Printf("System group ID %s\n", group.Id)
    }
  }
}

```
