## JCAPI-Go

### Description ###

This repository contains the Go client code for the JumpCloud API v2. This code is automatically generated using swagger-codegen.
It also provides the tools to generate the client code from the API yaml files, using swagger-codegen.
For detailed instructions on how to generate the code, see the [Contributing](CONTRIBUTING.md) section.

For the API v1 Go Client, please refer to [JCAPI](https://github.com/TheJumpCloud/jcapi).


#### Usage Examples

```
import (
  jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"
)
...
// instantiate the API object for the group of endpoints you need to use
// for instance for user groups API:
userGroupsAPI = jcapiv2.NewUserGroupsApi()

// set up the API key:
userGroupsAPI.Configuration.APIKey["x-api-key"] = YOUR_API_KEY

// make an API call to retrieve a specific user group by id:
userGroup, apiResponse, err := userGroupsAPI.GroupsUserGet("<YOUR_GROUP_ID>", "application/json", "application/json")
```
