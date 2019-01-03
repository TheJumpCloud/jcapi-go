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
	config := jcapiv2.NewConfiguration()
	config.BasePath = "http://localhost/api/v2"
	client := jcapiv2.NewAPIClient(config)

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
