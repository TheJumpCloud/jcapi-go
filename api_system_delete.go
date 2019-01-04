package main

import (
	"context"
	"fmt"

	jcapiv1 "github.com/TheJumpCloud/jcapi-go/v1"
)

func main() {
	apiKey := "c967cf539189cae968f55b6acf8605c57c686815"
	systemID := "5c2fc265ad2d81603d115833"

	contentType := "application/json"
	accept := "application/json"

	// Instantiate the API client
	config := jcapiv1.NewConfiguration()
	config.BasePath = "https://console.awsstg.jumpcloud.com/api"
	client := jcapiv1.NewAPIClient(config)

	// Set up the API key via context
	auth := context.WithValue(context.TODO(), jcapiv1.ContextAPIKey, jcapiv1.APIKey{
		Key: apiKey,
	})

	// Make an API call to delete a specific system
	system, res, err := client.SystemsApi.SystemsDelete(auth, systemID, contentType, accept, nil)
	if err != nil {
		fmt.Printf("Error retrieving system %s: %s - response = %+v\n", systemID, err, res)
	} else {
		fmt.Printf("Details for System %s: %+v\n", systemID, system)
	}
}
