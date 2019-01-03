package main

import (
	"context"
	"fmt"

	jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"
)

func main() {
	apiKey := "a23b348c24f238d6fc5de2f2317015d955146024"
	userGroupID := "5bedae547a36e106ab525677"

	contentType := "application/json"
	accept := "application/json"

	// Instantiate the API client
	config := jcapiv2.NewConfiguration()
	config.BasePath = "http://localhost/api/v2"
	client := jcapiv2.NewAPIClient(config)

	// Set up the API key via context
	auth := context.WithValue(context.TODO(), jcapiv2.ContextAPIKey, jcapiv2.APIKey{
		Key: apiKey,
	})

	// Make an API call to retrieve a specific user group by ID
	userGroup, res, err := client.UserGroupsApi.GroupsUserGet(auth, userGroupID, contentType, accept, nil)
	if err != nil {
		fmt.Printf("Error retrieving user group %s: %s - response = %+v\n", userGroupID, err, res)
	} else {
		fmt.Printf("Details for User group %s: %+v\n", userGroupID, userGroup)
	}
}
