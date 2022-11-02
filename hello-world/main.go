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
