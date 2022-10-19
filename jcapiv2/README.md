# Go API client for jcapiv2

# Overview  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.  # Directory Objects  This API offers the ability to interact with some of our core features; otherwise known as Directory Objects. The Directory Objects are:  * Commands * Policies * Policy Groups * Applications * Systems * Users * User Groups * System Groups * Radius Servers * Directories: Office 365, LDAP,G-Suite, Active Directory * Duo accounts and applications.  The Directory Object is an important concept to understand in order to successfully use JumpCloud API.  ## JumpCloud Graph  We've also introduced the concept of the JumpCloud Graph along with  Directory Objects. The Graph is a powerful aspect of our platform which will enable you to associate objects with each other, or establish membership for certain objects to become members of other objects.  Specific `GET` endpoints will allow you to traverse the JumpCloud Graph to return all indirect and directly bound objects in your organization.  | ![alt text](https://s3.amazonaws.com/jumpcloud-kb/Knowledge+Base+Photos/API+Docs/jumpcloud_graph.png \"JumpCloud Graph Model Example\") | |:--:| | **This diagram highlights our association and membership model as it relates to Directory Objects.** |  # API Key  ## Access Your API Key  To locate your API Key:  1. Log into the [JumpCloud Admin Console](https://console.jumpcloud.com/). 2. Go to the username drop down located in the top-right of the Console. 3. Retrieve your API key from API Settings.  ## API Key Considerations  This API key is associated to the currently logged in administrator. Other admins will have different API keys.  **WARNING** Please keep this API key secret, as it grants full access to any data accessible via your JumpCloud console account.  You can also reset your API key in the same location in the JumpCloud Admin Console.  ## Recycling or Resetting Your API Key  In order to revoke access with the current API key, simply reset your API key. This will render all calls using the previous API key inaccessible.  Your API key will be passed in as a header with the header name \"x-api-key\".  ```bash curl -H \"x-api-key: [YOUR_API_KEY_HERE]\" \"https://console.jumpcloud.com/api/v2/systemgroups\" ```  # System Context  * [Introduction](#introduction) * [Supported endpoints](#supported-endpoints) * [Response codes](#response-codes) * [Authentication](#authentication) * [Additional examples](#additional-examples) * [Third party](#third-party)  ## Introduction  JumpCloud System Context Authorization is an alternative way to authenticate with a subset of JumpCloud's REST APIs. Using this method, a system can manage its information and resource associations, allowing modern auto provisioning environments to scale as needed.  **Notes:**   * The following documentation applies to Linux Operating Systems only.  * Systems that have been automatically enrolled using Apple's Device Enrollment Program (DEP) or systems enrolled using the User Portal install are not eligible to use the System Context API to prevent unauthorized access to system groups and resources. If a script that utilizes the System Context API is invoked on a system enrolled in this way, it will display an error.  ## Supported Endpoints  JumpCloud System Context Authorization can be used in conjunction with Systems endpoints found in the V1 API and certain System Group endpoints found in the v2 API.  * A system may fetch, alter, and delete metadata about itself, including manipulating a system's Group and Systemuser associations,   * `/api/systems/{system_id}` | [`GET`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_get) [`PUT`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_put) * A system may delete itself from your JumpCloud organization   * `/api/systems/{system_id}` | [`DELETE`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_delete) * A system may fetch its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/memberof` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembership)   * `/api/v2/systems/{system_id}/associations` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsList)   * `/api/v2/systems/{system_id}/users` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemTraverseUser) * A system may alter its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/associations` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsPost) * A system may alter its System Group associations   * `/api/v2/systemgroups/{group_id}/members` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembersPost)     * _NOTE_ If a system attempts to alter the system group membership of a different system the request will be rejected  ## Response Codes  If endpoints other than those described above are called using the System Context API, the server will return a `401` response.  ## Authentication  To allow for secure access to our APIs, you must authenticate each API request. JumpCloud System Context Authorization uses [HTTP Signatures](https://tools.ietf.org/html/draft-cavage-http-signatures-00) to authenticate API requests. The HTTP Signatures sent with each request are similar to the signatures used by the Amazon Web Services REST API. To help with the request-signing process, we have provided an [example bash script](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh). This example API request simply requests the entire system record. You must be root, or have permissions to access the contents of the `/opt/jc` directory to generate a signature.  Here is a breakdown of the example script with explanations.  First, the script extracts the systemKey from the JSON formatted `/opt/jc/jcagent.conf` file.  ```bash #!/bin/bash conf=\"`cat /opt/jc/jcagent.conf`\" regex=\"systemKey\\\":\\\"(\\w+)\\\"\"  if [[ $conf =~ $regex ]] ; then   systemKey=\"${BASH_REMATCH[1]}\" fi ```  Then, the script retrieves the current date in the correct format.  ```bash now=`date -u \"+%a, %d %h %Y %H:%M:%S GMT\"`; ```  Next, we build a signing string to demonstrate the expected signature format. The signed string must consist of the [request-line](https://tools.ietf.org/html/rfc2616#page-35) and the date header, separated by a newline character.  ```bash signstr=\"GET /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" ```  The next step is to calculate and apply the signature. This is a two-step process:  1. Create a signature from the signing string using the JumpCloud Agent private key: ``printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key`` 2. Then Base64-encode the signature string and trim off the newline characters: ``| openssl enc -e -a | tr -d '\\n'``  The combined steps above result in:  ```bash signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ; ```  Finally, we make sure the API call sending the signature has the same Authorization and Date header values, HTTP method, and URL that were used in the signing string.  ```bash curl -iq \\   -H \"Accept: application/json\" \\   -H \"Content-Type: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Input Data  All PUT and POST methods should use the HTTP Content-Type header with a value of 'application/json'. PUT methods are used for updating a record. POST methods are used to create a record.  The following example demonstrates how to update the `displayName` of the system.  ```bash signstr=\"PUT /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ;  curl -iq \\   -d \"{\\\"displayName\\\" : \\\"updated-system-name-1\\\"}\" \\   -X \"PUT\" \\   -H \"Content-Type: application/json\" \\   -H \"Accept: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Output Data  All results will be formatted as JSON.  Here is an abbreviated example of response output:  ```json {   \"_id\": \"525ee96f52e144993e000015\",   \"agentServer\": \"lappy386\",   \"agentVersion\": \"0.9.42\",   \"arch\": \"x86_64\",   \"connectionKey\": \"127.0.0.1_51812\",   \"displayName\": \"ubuntu-1204\",   \"firstContact\": \"2013-10-16T19:30:55.611Z\",   \"hostname\": \"ubuntu-1204\"   ... ```  ## Additional Examples  ### Signing Authentication Example  This example demonstrates how to make an authenticated request to fetch the JumpCloud record for this system.  [SigningExample.sh](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh)  ### Shutdown Hook  This example demonstrates how to make an authenticated request on system shutdown. Using an init.d script registered at run level 0, you can call the System Context API as the system is shutting down.  [Instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) is an example of an init.d script that only runs at system shutdown.  After customizing the [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) script, you should install it on the system(s) running the JumpCloud agent.  1. Copy the modified [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) to `/etc/init.d/instance-shutdown`. 2. On Ubuntu systems, run `update-rc.d instance-shutdown defaults`. On RedHat/CentOS systems, run `chkconfig --add instance-shutdown`.  ## Third Party  ### Chef Cookbooks  [https://github.com/nshenry03/jumpcloud](https://github.com/nshenry03/jumpcloud)  [https://github.com/cjs226/jumpcloud](https://github.com/cjs226/jumpcloud)  # Multi-Tenant Portal Headers  Multi-Tenant Organization API Headers are available for JumpCloud Admins to use when making API requests from Organizations that have multiple managed organizations.  The `x-org-id` is a required header for all multi-tenant admins when making API requests to JumpCloud. This header will define to which organization you would like to make the request.  **NOTE** Single Tenant Admins do not need to provide this header when making an API request.  ## Header Value  `x-org-id`  ## API Response Codes  * `400` Malformed ID. * `400` x-org-id and Organization path ID do not match. * `401` ID not included for multi-tenant admin * `403` ID included on unsupported route. * `404` Organization ID Not Found.  ```bash curl -X GET https://console.jumpcloud.com/api/v2/directories \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -H 'x-org-id: {ORG_ID}'  ```  ## To Obtain an Individual Organization ID via the UI  As a prerequisite, your Primary Organization will need to be setup for Multi-Tenancy. This provides access to the Multi-Tenant Organization Admin Portal.  1. Log into JumpCloud [Admin Console](https://console.jumpcloud.com). If you are a multi-tenant Admin, you will automatically be routed to the Multi-Tenant Admin Portal. 2. From the Multi-Tenant Portal's primary navigation bar, select the Organization you'd like to access. 3. You will automatically be routed to that Organization's Admin Console. 4. Go to Settings in the sub-tenant's primary navigation. 5. You can obtain your Organization ID below your Organization's Contact Information on the Settings page.  ## To Obtain All Organization IDs via the API  * You can make an API request to this endpoint using the API key of your Primary Organization.  `https://console.jumpcloud.com/api/organizations/` This will return all your managed organizations.  ```bash curl -X GET \\   https://console.jumpcloud.com/api/organizations/ \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```  # SDKs  You can find language specific SDKs that can help you kickstart your Integration with JumpCloud in the following GitHub repositories:  * [Python](https://github.com/TheJumpCloud/jcapi-python) * [Go](https://github.com/TheJumpCloud/jcapi-go) * [Ruby](https://github.com/TheJumpCloud/jcapi-ruby) * [Java](https://github.com/TheJumpCloud/jcapi-java) 

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2.0
- Package version: 5.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen
For more information, please visit [https://support.jumpcloud.com/support/s/](https://support.jumpcloud.com/support/s/)

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./jcapiv2"
```

## Documentation for API Endpoints

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ActiveDirectoryApi* | [**ActivedirectoriesAgentsDelete**](docs/ActiveDirectoryApi.md#activedirectoriesagentsdelete) | **Delete** /activedirectories/{activedirectory_id}/agents/{agent_id} | Delete Active Directory Agent
*ActiveDirectoryApi* | [**ActivedirectoriesAgentsGet**](docs/ActiveDirectoryApi.md#activedirectoriesagentsget) | **Get** /activedirectories/{activedirectory_id}/agents/{agent_id} | Get Active Directory Agent
*ActiveDirectoryApi* | [**ActivedirectoriesAgentsList**](docs/ActiveDirectoryApi.md#activedirectoriesagentslist) | **Get** /activedirectories/{activedirectory_id}/agents | List Active Directory Agents
*ActiveDirectoryApi* | [**ActivedirectoriesAgentsPost**](docs/ActiveDirectoryApi.md#activedirectoriesagentspost) | **Post** /activedirectories/{activedirectory_id}/agents | Create a new Active Directory Agent
*ActiveDirectoryApi* | [**ActivedirectoriesDelete**](docs/ActiveDirectoryApi.md#activedirectoriesdelete) | **Delete** /activedirectories/{id} | Delete an Active Directory
*ActiveDirectoryApi* | [**ActivedirectoriesGet**](docs/ActiveDirectoryApi.md#activedirectoriesget) | **Get** /activedirectories/{id} | Get an Active Directory
*ActiveDirectoryApi* | [**ActivedirectoriesList**](docs/ActiveDirectoryApi.md#activedirectorieslist) | **Get** /activedirectories | List Active Directories
*ActiveDirectoryApi* | [**ActivedirectoriesPost**](docs/ActiveDirectoryApi.md#activedirectoriespost) | **Post** /activedirectories | Create a new Active Directory
*ActiveDirectoryApi* | [**GraphActiveDirectoryAssociationsList**](docs/ActiveDirectoryApi.md#graphactivedirectoryassociationslist) | **Get** /activedirectories/{activedirectory_id}/associations | List the associations of an Active Directory instance
*ActiveDirectoryApi* | [**GraphActiveDirectoryAssociationsPost**](docs/ActiveDirectoryApi.md#graphactivedirectoryassociationspost) | **Post** /activedirectories/{activedirectory_id}/associations | Manage the associations of an Active Directory instance
*ActiveDirectoryApi* | [**GraphActiveDirectoryTraverseUser**](docs/ActiveDirectoryApi.md#graphactivedirectorytraverseuser) | **Get** /activedirectories/{activedirectory_id}/users | List the Users bound to an Active Directory instance
*ActiveDirectoryApi* | [**GraphActiveDirectoryTraverseUserGroup**](docs/ActiveDirectoryApi.md#graphactivedirectorytraverseusergroup) | **Get** /activedirectories/{activedirectory_id}/usergroups | List the User Groups bound to an Active Directory instance
*AdministratorsApi* | [**AdministratorOrganizationsCreateByAdministrator**](docs/AdministratorsApi.md#administratororganizationscreatebyadministrator) | **Post** /administrators/{id}/organizationlinks | Allow Adminstrator access to an Organization.
*AdministratorsApi* | [**AdministratorOrganizationsListByAdministrator**](docs/AdministratorsApi.md#administratororganizationslistbyadministrator) | **Get** /administrators/{id}/organizationlinks | List the association links between an Administrator and Organizations.
*AdministratorsApi* | [**AdministratorOrganizationsListByOrganization**](docs/AdministratorsApi.md#administratororganizationslistbyorganization) | **Get** /organizations/{id}/administratorlinks | List the association links between an Organization and Administrators.
*AdministratorsApi* | [**AdministratorOrganizationsRemoveByAdministrator**](docs/AdministratorsApi.md#administratororganizationsremovebyadministrator) | **Delete** /administrators/{administrator_id}/organizationlinks/{id} | Remove association between an Administrator and an Organization.
*AppleMDMApi* | [**ApplemdmsCsrget**](docs/AppleMDMApi.md#applemdmscsrget) | **Get** /applemdms/{apple_mdm_id}/csr | Get Apple MDM CSR Plist
*AppleMDMApi* | [**ApplemdmsDelete**](docs/AppleMDMApi.md#applemdmsdelete) | **Delete** /applemdms/{id} | Delete an Apple MDM
*AppleMDMApi* | [**ApplemdmsDeletedevice**](docs/AppleMDMApi.md#applemdmsdeletedevice) | **Delete** /applemdms/{apple_mdm_id}/devices/{device_id} | Remove an Apple MDM Device&#x27;s Enrollment
*AppleMDMApi* | [**ApplemdmsDepkeyget**](docs/AppleMDMApi.md#applemdmsdepkeyget) | **Get** /applemdms/{apple_mdm_id}/depkey | Get Apple MDM DEP Public Key
*AppleMDMApi* | [**ApplemdmsDevicesClearActivationLock**](docs/AppleMDMApi.md#applemdmsdevicesclearactivationlock) | **Post** /applemdms/{apple_mdm_id}/devices/{device_id}/clearActivationLock | Clears the Activation Lock for a Device
*AppleMDMApi* | [**ApplemdmsDevicesRefreshActivationLockInformation**](docs/AppleMDMApi.md#applemdmsdevicesrefreshactivationlockinformation) | **Post** /applemdms/{apple_mdm_id}/devices/{device_id}/refreshActivationLockInformation | Refresh activation lock information for a device
*AppleMDMApi* | [**ApplemdmsDeviceserase**](docs/AppleMDMApi.md#applemdmsdeviceserase) | **Post** /applemdms/{apple_mdm_id}/devices/{device_id}/erase | Erase Device
*AppleMDMApi* | [**ApplemdmsDeviceslist**](docs/AppleMDMApi.md#applemdmsdeviceslist) | **Get** /applemdms/{apple_mdm_id}/devices | List AppleMDM Devices
*AppleMDMApi* | [**ApplemdmsDeviceslock**](docs/AppleMDMApi.md#applemdmsdeviceslock) | **Post** /applemdms/{apple_mdm_id}/devices/{device_id}/lock | Lock Device
*AppleMDMApi* | [**ApplemdmsDevicesrestart**](docs/AppleMDMApi.md#applemdmsdevicesrestart) | **Post** /applemdms/{apple_mdm_id}/devices/{device_id}/restart | Restart Device
*AppleMDMApi* | [**ApplemdmsDevicesshutdown**](docs/AppleMDMApi.md#applemdmsdevicesshutdown) | **Post** /applemdms/{apple_mdm_id}/devices/{device_id}/shutdown | Shut Down Device
*AppleMDMApi* | [**ApplemdmsEnrollmentprofilesget**](docs/AppleMDMApi.md#applemdmsenrollmentprofilesget) | **Get** /applemdms/{apple_mdm_id}/enrollmentprofiles/{id} | Get an Apple MDM Enrollment Profile
*AppleMDMApi* | [**ApplemdmsEnrollmentprofileslist**](docs/AppleMDMApi.md#applemdmsenrollmentprofileslist) | **Get** /applemdms/{apple_mdm_id}/enrollmentprofiles | List Apple MDM Enrollment Profiles
*AppleMDMApi* | [**ApplemdmsGetdevice**](docs/AppleMDMApi.md#applemdmsgetdevice) | **Get** /applemdms/{apple_mdm_id}/devices/{device_id} | Details of an AppleMDM Device
*AppleMDMApi* | [**ApplemdmsList**](docs/AppleMDMApi.md#applemdmslist) | **Get** /applemdms | List Apple MDMs
*AppleMDMApi* | [**ApplemdmsPut**](docs/AppleMDMApi.md#applemdmsput) | **Put** /applemdms/{id} | Update an Apple MDM
*AppleMDMApi* | [**ApplemdmsRefreshdepdevices**](docs/AppleMDMApi.md#applemdmsrefreshdepdevices) | **Post** /applemdms/{apple_mdm_id}/refreshdepdevices | Refresh DEP Devices
*ApplicationsApi* | [**ApplicationsDeleteLogo**](docs/ApplicationsApi.md#applicationsdeletelogo) | **Delete** /applications/{application_id}/logo | Delete application image
*ApplicationsApi* | [**ApplicationsGet**](docs/ApplicationsApi.md#applicationsget) | **Get** /applications/{application_id} | Get an Application
*ApplicationsApi* | [**ApplicationsPostLogo**](docs/ApplicationsApi.md#applicationspostlogo) | **Post** /applications/{application_id}/logo | 
*ApplicationsApi* | [**GraphApplicationAssociationsList**](docs/ApplicationsApi.md#graphapplicationassociationslist) | **Get** /applications/{application_id}/associations | List the associations of an Application
*ApplicationsApi* | [**GraphApplicationAssociationsPost**](docs/ApplicationsApi.md#graphapplicationassociationspost) | **Post** /applications/{application_id}/associations | Manage the associations of an Application
*ApplicationsApi* | [**GraphApplicationTraverseUser**](docs/ApplicationsApi.md#graphapplicationtraverseuser) | **Get** /applications/{application_id}/users | List the Users bound to an Application
*ApplicationsApi* | [**GraphApplicationTraverseUserGroup**](docs/ApplicationsApi.md#graphapplicationtraverseusergroup) | **Get** /applications/{application_id}/usergroups | List the User Groups bound to an Application
*ApplicationsApi* | [**ImportUsers**](docs/ApplicationsApi.md#importusers) | **Get** /applications/{application_id}/import/users | Get a list of users to import from an Application IdM service provider
*AuthenticationPoliciesApi* | [**AuthnpoliciesDelete**](docs/AuthenticationPoliciesApi.md#authnpoliciesdelete) | **Delete** /authn/policies/{id} | Delete Authentication Policy
*AuthenticationPoliciesApi* | [**AuthnpoliciesGet**](docs/AuthenticationPoliciesApi.md#authnpoliciesget) | **Get** /authn/policies/{id} | Get an authentication policy
*AuthenticationPoliciesApi* | [**AuthnpoliciesList**](docs/AuthenticationPoliciesApi.md#authnpolicieslist) | **Get** /authn/policies | List Authentication Policies
*AuthenticationPoliciesApi* | [**AuthnpoliciesPatch**](docs/AuthenticationPoliciesApi.md#authnpoliciespatch) | **Patch** /authn/policies/{id} | Patch Authentication Policy
*AuthenticationPoliciesApi* | [**AuthnpoliciesPost**](docs/AuthenticationPoliciesApi.md#authnpoliciespost) | **Post** /authn/policies | Create an Authentication Policy
*BulkJobRequestsApi* | [**BulkUserStatesCreate**](docs/BulkJobRequestsApi.md#bulkuserstatescreate) | **Post** /bulk/userstates | Create Scheduled Userstate Job
*BulkJobRequestsApi* | [**BulkUserStatesDelete**](docs/BulkJobRequestsApi.md#bulkuserstatesdelete) | **Delete** /bulk/userstates/{id} | Delete Scheduled Userstate Job
*BulkJobRequestsApi* | [**BulkUserStatesGetNextScheduled**](docs/BulkJobRequestsApi.md#bulkuserstatesgetnextscheduled) | **Get** /bulk/userstates/eventlist/next | Gets the next scheduled state change for each user in a list of system users
*BulkJobRequestsApi* | [**BulkUserStatesList**](docs/BulkJobRequestsApi.md#bulkuserstateslist) | **Get** /bulk/userstates | List Scheduled Userstate Change Jobs
*BulkJobRequestsApi* | [**BulkUsersCreate**](docs/BulkJobRequestsApi.md#bulkuserscreate) | **Post** /bulk/users | Bulk Users Create
*BulkJobRequestsApi* | [**BulkUsersCreateResults**](docs/BulkJobRequestsApi.md#bulkuserscreateresults) | **Get** /bulk/users/{job_id}/results | List Bulk Users Results
*BulkJobRequestsApi* | [**BulkUsersUpdate**](docs/BulkJobRequestsApi.md#bulkusersupdate) | **Patch** /bulk/users | Bulk Users Update
*CommandResultsApi* | [**CommandsListResultsByWorkflow**](docs/CommandResultsApi.md#commandslistresultsbyworkflow) | **Get** /commandresult/workflows | List all Command Results by Workflow
*CommandsApi* | [**CommandsCancelQueuedCommandsByWorkflowInstanceId**](docs/CommandsApi.md#commandscancelqueuedcommandsbyworkflowinstanceid) | **Delete** /commandqueue/{workflow_instance_id} | Cancel all queued commands for an organization by workflow instance Id
*CommandsApi* | [**CommandsGetQueuedCommandsByWorkflow**](docs/CommandsApi.md#commandsgetqueuedcommandsbyworkflow) | **Get** /queuedcommand/workflows | Fetch the queued Commands for an Organization
*CommandsApi* | [**GraphCommandAssociationsList**](docs/CommandsApi.md#graphcommandassociationslist) | **Get** /commands/{command_id}/associations | List the associations of a Command
*CommandsApi* | [**GraphCommandAssociationsPost**](docs/CommandsApi.md#graphcommandassociationspost) | **Post** /commands/{command_id}/associations | Manage the associations of a Command
*CommandsApi* | [**GraphCommandTraverseSystem**](docs/CommandsApi.md#graphcommandtraversesystem) | **Get** /commands/{command_id}/systems | List the Systems bound to a Command
*CommandsApi* | [**GraphCommandTraverseSystemGroup**](docs/CommandsApi.md#graphcommandtraversesystemgroup) | **Get** /commands/{command_id}/systemgroups | List the System Groups bound to a Command
*CustomEmailsApi* | [**CustomEmailsCreate**](docs/CustomEmailsApi.md#customemailscreate) | **Post** /customemails | Create custom email configuration
*CustomEmailsApi* | [**CustomEmailsDestroy**](docs/CustomEmailsApi.md#customemailsdestroy) | **Delete** /customemails/{custom_email_type} | Delete custom email configuration
*CustomEmailsApi* | [**CustomEmailsGetTemplates**](docs/CustomEmailsApi.md#customemailsgettemplates) | **Get** /customemail/templates | List custom email templates
*CustomEmailsApi* | [**CustomEmailsRead**](docs/CustomEmailsApi.md#customemailsread) | **Get** /customemails/{custom_email_type} | Get custom email configuration
*CustomEmailsApi* | [**CustomEmailsUpdate**](docs/CustomEmailsApi.md#customemailsupdate) | **Put** /customemails/{custom_email_type} | Update custom email configuration
*DirectoriesApi* | [**DirectoriesList**](docs/DirectoriesApi.md#directorieslist) | **Get** /directories | List All Directories
*DuoApi* | [**DuoAccountDelete**](docs/DuoApi.md#duoaccountdelete) | **Delete** /duo/accounts/{id} | Delete a Duo Account
*DuoApi* | [**DuoAccountGet**](docs/DuoApi.md#duoaccountget) | **Get** /duo/accounts/{id} | Get a Duo Acount
*DuoApi* | [**DuoAccountList**](docs/DuoApi.md#duoaccountlist) | **Get** /duo/accounts | List Duo Accounts
*DuoApi* | [**DuoAccountPost**](docs/DuoApi.md#duoaccountpost) | **Post** /duo/accounts | Create Duo Account
*DuoApi* | [**DuoApplicationDelete**](docs/DuoApi.md#duoapplicationdelete) | **Delete** /duo/accounts/{account_id}/applications/{application_id} | Delete a Duo Application
*DuoApi* | [**DuoApplicationGet**](docs/DuoApi.md#duoapplicationget) | **Get** /duo/accounts/{account_id}/applications/{application_id} | Get a Duo application
*DuoApi* | [**DuoApplicationList**](docs/DuoApi.md#duoapplicationlist) | **Get** /duo/accounts/{account_id}/applications | List Duo Applications
*DuoApi* | [**DuoApplicationPost**](docs/DuoApi.md#duoapplicationpost) | **Post** /duo/accounts/{account_id}/applications | Create Duo Application
*DuoApi* | [**DuoApplicationUpdate**](docs/DuoApi.md#duoapplicationupdate) | **Put** /duo/accounts/{account_id}/applications/{application_id} | Update Duo Application
*FdeApi* | [**SystemsGetFDEKey**](docs/FdeApi.md#systemsgetfdekey) | **Get** /systems/{system_id}/fdekey | Get System FDE Key
*GSuiteApi* | [**GraphGSuiteAssociationsList**](docs/GSuiteApi.md#graphgsuiteassociationslist) | **Get** /gsuites/{gsuite_id}/associations | List the associations of a G Suite instance
*GSuiteApi* | [**GraphGSuiteAssociationsPost**](docs/GSuiteApi.md#graphgsuiteassociationspost) | **Post** /gsuites/{gsuite_id}/associations | Manage the associations of a G Suite instance
*GSuiteApi* | [**GraphGSuiteTraverseUser**](docs/GSuiteApi.md#graphgsuitetraverseuser) | **Get** /gsuites/{gsuite_id}/users | List the Users bound to a G Suite instance
*GSuiteApi* | [**GraphGSuiteTraverseUserGroup**](docs/GSuiteApi.md#graphgsuitetraverseusergroup) | **Get** /gsuites/{gsuite_id}/usergroups | List the User Groups bound to a G Suite instance
*GSuiteApi* | [**GsuitesGet**](docs/GSuiteApi.md#gsuitesget) | **Get** /gsuites/{id} | Get G Suite
*GSuiteApi* | [**GsuitesListImportJumpcloudUsers**](docs/GSuiteApi.md#gsuiteslistimportjumpcloudusers) | **Get** /gsuites/{gsuite_id}/import/jumpcloudusers | Get a list of users in Jumpcloud format to import from a Google Workspace account.
*GSuiteApi* | [**GsuitesListImportUsers**](docs/GSuiteApi.md#gsuiteslistimportusers) | **Get** /gsuites/{gsuite_id}/import/users | Get a list of users to import from a G Suite instance
*GSuiteApi* | [**GsuitesPatch**](docs/GSuiteApi.md#gsuitespatch) | **Patch** /gsuites/{id} | Update existing G Suite
*GSuiteApi* | [**TranslationRulesGSuiteDelete**](docs/GSuiteApi.md#translationrulesgsuitedelete) | **Delete** /gsuites/{gsuite_id}/translationrules/{id} | Deletes a G Suite translation rule
*GSuiteApi* | [**TranslationRulesGSuiteGet**](docs/GSuiteApi.md#translationrulesgsuiteget) | **Get** /gsuites/{gsuite_id}/translationrules/{id} | Gets a specific G Suite translation rule
*GSuiteApi* | [**TranslationRulesGSuiteList**](docs/GSuiteApi.md#translationrulesgsuitelist) | **Get** /gsuites/{gsuite_id}/translationrules | List all the G Suite Translation Rules
*GSuiteApi* | [**TranslationRulesGSuitePost**](docs/GSuiteApi.md#translationrulesgsuitepost) | **Post** /gsuites/{gsuite_id}/translationrules | Create a new G Suite Translation Rule
*GSuiteImportApi* | [**GsuitesListImportJumpcloudUsers**](docs/GSuiteImportApi.md#gsuiteslistimportjumpcloudusers) | **Get** /gsuites/{gsuite_id}/import/jumpcloudusers | Get a list of users in Jumpcloud format to import from a Google Workspace account.
*GSuiteImportApi* | [**GsuitesListImportUsers**](docs/GSuiteImportApi.md#gsuiteslistimportusers) | **Get** /gsuites/{gsuite_id}/import/users | Get a list of users to import from a G Suite instance
*GraphApi* | [**GraphActiveDirectoryAssociationsList**](docs/GraphApi.md#graphactivedirectoryassociationslist) | **Get** /activedirectories/{activedirectory_id}/associations | List the associations of an Active Directory instance
*GraphApi* | [**GraphActiveDirectoryAssociationsPost**](docs/GraphApi.md#graphactivedirectoryassociationspost) | **Post** /activedirectories/{activedirectory_id}/associations | Manage the associations of an Active Directory instance
*GraphApi* | [**GraphActiveDirectoryTraverseUser**](docs/GraphApi.md#graphactivedirectorytraverseuser) | **Get** /activedirectories/{activedirectory_id}/users | List the Users bound to an Active Directory instance
*GraphApi* | [**GraphActiveDirectoryTraverseUserGroup**](docs/GraphApi.md#graphactivedirectorytraverseusergroup) | **Get** /activedirectories/{activedirectory_id}/usergroups | List the User Groups bound to an Active Directory instance
*GraphApi* | [**GraphApplicationAssociationsList**](docs/GraphApi.md#graphapplicationassociationslist) | **Get** /applications/{application_id}/associations | List the associations of an Application
*GraphApi* | [**GraphApplicationAssociationsPost**](docs/GraphApi.md#graphapplicationassociationspost) | **Post** /applications/{application_id}/associations | Manage the associations of an Application
*GraphApi* | [**GraphApplicationTraverseUser**](docs/GraphApi.md#graphapplicationtraverseuser) | **Get** /applications/{application_id}/users | List the Users bound to an Application
*GraphApi* | [**GraphApplicationTraverseUserGroup**](docs/GraphApi.md#graphapplicationtraverseusergroup) | **Get** /applications/{application_id}/usergroups | List the User Groups bound to an Application
*GraphApi* | [**GraphCommandAssociationsList**](docs/GraphApi.md#graphcommandassociationslist) | **Get** /commands/{command_id}/associations | List the associations of a Command
*GraphApi* | [**GraphCommandAssociationsPost**](docs/GraphApi.md#graphcommandassociationspost) | **Post** /commands/{command_id}/associations | Manage the associations of a Command
*GraphApi* | [**GraphCommandTraverseSystem**](docs/GraphApi.md#graphcommandtraversesystem) | **Get** /commands/{command_id}/systems | List the Systems bound to a Command
*GraphApi* | [**GraphCommandTraverseSystemGroup**](docs/GraphApi.md#graphcommandtraversesystemgroup) | **Get** /commands/{command_id}/systemgroups | List the System Groups bound to a Command
*GraphApi* | [**GraphGSuiteAssociationsList**](docs/GraphApi.md#graphgsuiteassociationslist) | **Get** /gsuites/{gsuite_id}/associations | List the associations of a G Suite instance
*GraphApi* | [**GraphGSuiteAssociationsPost**](docs/GraphApi.md#graphgsuiteassociationspost) | **Post** /gsuites/{gsuite_id}/associations | Manage the associations of a G Suite instance
*GraphApi* | [**GraphGSuiteTraverseUser**](docs/GraphApi.md#graphgsuitetraverseuser) | **Get** /gsuites/{gsuite_id}/users | List the Users bound to a G Suite instance
*GraphApi* | [**GraphGSuiteTraverseUserGroup**](docs/GraphApi.md#graphgsuitetraverseusergroup) | **Get** /gsuites/{gsuite_id}/usergroups | List the User Groups bound to a G Suite instance
*GraphApi* | [**GraphLdapServerAssociationsList**](docs/GraphApi.md#graphldapserverassociationslist) | **Get** /ldapservers/{ldapserver_id}/associations | List the associations of a LDAP Server
*GraphApi* | [**GraphLdapServerAssociationsPost**](docs/GraphApi.md#graphldapserverassociationspost) | **Post** /ldapservers/{ldapserver_id}/associations | Manage the associations of a LDAP Server
*GraphApi* | [**GraphLdapServerTraverseUser**](docs/GraphApi.md#graphldapservertraverseuser) | **Get** /ldapservers/{ldapserver_id}/users | List the Users bound to a LDAP Server
*GraphApi* | [**GraphLdapServerTraverseUserGroup**](docs/GraphApi.md#graphldapservertraverseusergroup) | **Get** /ldapservers/{ldapserver_id}/usergroups | List the User Groups bound to a LDAP Server
*GraphApi* | [**GraphOffice365AssociationsList**](docs/GraphApi.md#graphoffice365associationslist) | **Get** /office365s/{office365_id}/associations | List the associations of an Office 365 instance
*GraphApi* | [**GraphOffice365AssociationsPost**](docs/GraphApi.md#graphoffice365associationspost) | **Post** /office365s/{office365_id}/associations | Manage the associations of an Office 365 instance
*GraphApi* | [**GraphOffice365TraverseUser**](docs/GraphApi.md#graphoffice365traverseuser) | **Get** /office365s/{office365_id}/users | List the Users bound to an Office 365 instance
*GraphApi* | [**GraphOffice365TraverseUserGroup**](docs/GraphApi.md#graphoffice365traverseusergroup) | **Get** /office365s/{office365_id}/usergroups | List the User Groups bound to an Office 365 instance
*GraphApi* | [**GraphPolicyAssociationsList**](docs/GraphApi.md#graphpolicyassociationslist) | **Get** /policies/{policy_id}/associations | List the associations of a Policy
*GraphApi* | [**GraphPolicyAssociationsPost**](docs/GraphApi.md#graphpolicyassociationspost) | **Post** /policies/{policy_id}/associations | Manage the associations of a Policy
*GraphApi* | [**GraphPolicyGroupAssociationsList**](docs/GraphApi.md#graphpolicygroupassociationslist) | **Get** /policygroups/{group_id}/associations | List the associations of a Policy Group.
*GraphApi* | [**GraphPolicyGroupAssociationsPost**](docs/GraphApi.md#graphpolicygroupassociationspost) | **Post** /policygroups/{group_id}/associations | Manage the associations of a Policy Group
*GraphApi* | [**GraphPolicyGroupMembersList**](docs/GraphApi.md#graphpolicygroupmemberslist) | **Get** /policygroups/{group_id}/members | List the members of a Policy Group
*GraphApi* | [**GraphPolicyGroupMembersPost**](docs/GraphApi.md#graphpolicygroupmemberspost) | **Post** /policygroups/{group_id}/members | Manage the members of a Policy Group
*GraphApi* | [**GraphPolicyGroupMembership**](docs/GraphApi.md#graphpolicygroupmembership) | **Get** /policygroups/{group_id}/membership | List the Policy Group&#x27;s membership
*GraphApi* | [**GraphPolicyGroupTraverseSystem**](docs/GraphApi.md#graphpolicygrouptraversesystem) | **Get** /policygroups/{group_id}/systems | List the Systems bound to a Policy Group
*GraphApi* | [**GraphPolicyGroupTraverseSystemGroup**](docs/GraphApi.md#graphpolicygrouptraversesystemgroup) | **Get** /policygroups/{group_id}/systemgroups | List the System Groups bound to Policy Groups
*GraphApi* | [**GraphPolicyMemberOf**](docs/GraphApi.md#graphpolicymemberof) | **Get** /policies/{policy_id}/memberof | List the parent Groups of a Policy
*GraphApi* | [**GraphPolicyTraverseSystem**](docs/GraphApi.md#graphpolicytraversesystem) | **Get** /policies/{policy_id}/systems | List the Systems bound to a Policy
*GraphApi* | [**GraphPolicyTraverseSystemGroup**](docs/GraphApi.md#graphpolicytraversesystemgroup) | **Get** /policies/{policy_id}/systemgroups | List the System Groups bound to a Policy
*GraphApi* | [**GraphRadiusServerAssociationsList**](docs/GraphApi.md#graphradiusserverassociationslist) | **Get** /radiusservers/{radiusserver_id}/associations | List the associations of a RADIUS  Server
*GraphApi* | [**GraphRadiusServerAssociationsPost**](docs/GraphApi.md#graphradiusserverassociationspost) | **Post** /radiusservers/{radiusserver_id}/associations | Manage the associations of a RADIUS Server
*GraphApi* | [**GraphRadiusServerTraverseUser**](docs/GraphApi.md#graphradiusservertraverseuser) | **Get** /radiusservers/{radiusserver_id}/users | List the Users bound to a RADIUS  Server
*GraphApi* | [**GraphRadiusServerTraverseUserGroup**](docs/GraphApi.md#graphradiusservertraverseusergroup) | **Get** /radiusservers/{radiusserver_id}/usergroups | List the User Groups bound to a RADIUS  Server
*GraphApi* | [**GraphSoftwareappsAssociationsList**](docs/GraphApi.md#graphsoftwareappsassociationslist) | **Get** /softwareapps/{software_app_id}/associations | List the associations of a Software Application
*GraphApi* | [**GraphSoftwareappsAssociationsPost**](docs/GraphApi.md#graphsoftwareappsassociationspost) | **Post** /softwareapps/{software_app_id}/associations | Manage the associations of a software application.
*GraphApi* | [**GraphSoftwareappsTraverseSystem**](docs/GraphApi.md#graphsoftwareappstraversesystem) | **Get** /softwareapps/{software_app_id}/systems | List the Systems bound to a Software App.
*GraphApi* | [**GraphSoftwareappsTraverseSystemGroup**](docs/GraphApi.md#graphsoftwareappstraversesystemgroup) | **Get** /softwareapps/{software_app_id}/systemgroups | List the System Groups bound to a Software App.
*GraphApi* | [**GraphSystemAssociationsList**](docs/GraphApi.md#graphsystemassociationslist) | **Get** /systems/{system_id}/associations | List the associations of a System
*GraphApi* | [**GraphSystemAssociationsPost**](docs/GraphApi.md#graphsystemassociationspost) | **Post** /systems/{system_id}/associations | Manage associations of a System
*GraphApi* | [**GraphSystemGroupAssociationsList**](docs/GraphApi.md#graphsystemgroupassociationslist) | **Get** /systemgroups/{group_id}/associations | List the associations of a System Group
*GraphApi* | [**GraphSystemGroupAssociationsPost**](docs/GraphApi.md#graphsystemgroupassociationspost) | **Post** /systemgroups/{group_id}/associations | Manage the associations of a System Group
*GraphApi* | [**GraphSystemGroupMembersList**](docs/GraphApi.md#graphsystemgroupmemberslist) | **Get** /systemgroups/{group_id}/members | List the members of a System Group
*GraphApi* | [**GraphSystemGroupMembersPost**](docs/GraphApi.md#graphsystemgroupmemberspost) | **Post** /systemgroups/{group_id}/members | Manage the members of a System Group
*GraphApi* | [**GraphSystemGroupMembership**](docs/GraphApi.md#graphsystemgroupmembership) | **Get** /systemgroups/{group_id}/membership | List the System Group&#x27;s membership
*GraphApi* | [**GraphSystemGroupTraverseCommand**](docs/GraphApi.md#graphsystemgrouptraversecommand) | **Get** /systemgroups/{group_id}/commands | List the Commands bound to a System Group
*GraphApi* | [**GraphSystemGroupTraversePolicy**](docs/GraphApi.md#graphsystemgrouptraversepolicy) | **Get** /systemgroups/{group_id}/policies | List the Policies bound to a System Group
*GraphApi* | [**GraphSystemGroupTraversePolicyGroup**](docs/GraphApi.md#graphsystemgrouptraversepolicygroup) | **Get** /systemgroups/{group_id}/policygroups | List the Policy Groups bound to a System Group
*GraphApi* | [**GraphSystemGroupTraverseUser**](docs/GraphApi.md#graphsystemgrouptraverseuser) | **Get** /systemgroups/{group_id}/users | List the Users bound to a System Group
*GraphApi* | [**GraphSystemGroupTraverseUserGroup**](docs/GraphApi.md#graphsystemgrouptraverseusergroup) | **Get** /systemgroups/{group_id}/usergroups | List the User Groups bound to a System Group
*GraphApi* | [**GraphSystemMemberOf**](docs/GraphApi.md#graphsystemmemberof) | **Get** /systems/{system_id}/memberof | List the parent Groups of a System
*GraphApi* | [**GraphSystemTraverseCommand**](docs/GraphApi.md#graphsystemtraversecommand) | **Get** /systems/{system_id}/commands | List the Commands bound to a System
*GraphApi* | [**GraphSystemTraversePolicy**](docs/GraphApi.md#graphsystemtraversepolicy) | **Get** /systems/{system_id}/policies | List the Policies bound to a System
*GraphApi* | [**GraphSystemTraversePolicyGroup**](docs/GraphApi.md#graphsystemtraversepolicygroup) | **Get** /systems/{system_id}/policygroups | List the Policy Groups bound to a System
*GraphApi* | [**GraphSystemTraverseUser**](docs/GraphApi.md#graphsystemtraverseuser) | **Get** /systems/{system_id}/users | List the Users bound to a System
*GraphApi* | [**GraphSystemTraverseUserGroup**](docs/GraphApi.md#graphsystemtraverseusergroup) | **Get** /systems/{system_id}/usergroups | List the User Groups bound to a System
*GraphApi* | [**GraphUserAssociationsList**](docs/GraphApi.md#graphuserassociationslist) | **Get** /users/{user_id}/associations | List the associations of a User
*GraphApi* | [**GraphUserAssociationsPost**](docs/GraphApi.md#graphuserassociationspost) | **Post** /users/{user_id}/associations | Manage the associations of a User
*GraphApi* | [**GraphUserGroupAssociationsList**](docs/GraphApi.md#graphusergroupassociationslist) | **Get** /usergroups/{group_id}/associations | List the associations of a User Group.
*GraphApi* | [**GraphUserGroupAssociationsPost**](docs/GraphApi.md#graphusergroupassociationspost) | **Post** /usergroups/{group_id}/associations | Manage the associations of a User Group
*GraphApi* | [**GraphUserGroupMembersList**](docs/GraphApi.md#graphusergroupmemberslist) | **Get** /usergroups/{group_id}/members | List the members of a User Group
*GraphApi* | [**GraphUserGroupMembersPost**](docs/GraphApi.md#graphusergroupmemberspost) | **Post** /usergroups/{group_id}/members | Manage the members of a User Group
*GraphApi* | [**GraphUserGroupMembership**](docs/GraphApi.md#graphusergroupmembership) | **Get** /usergroups/{group_id}/membership | List the User Group&#x27;s membership
*GraphApi* | [**GraphUserGroupTraverseActiveDirectory**](docs/GraphApi.md#graphusergrouptraverseactivedirectory) | **Get** /usergroups/{group_id}/activedirectories | List the Active Directories bound to a User Group
*GraphApi* | [**GraphUserGroupTraverseApplication**](docs/GraphApi.md#graphusergrouptraverseapplication) | **Get** /usergroups/{group_id}/applications | List the Applications bound to a User Group
*GraphApi* | [**GraphUserGroupTraverseDirectory**](docs/GraphApi.md#graphusergrouptraversedirectory) | **Get** /usergroups/{group_id}/directories | List the Directories bound to a User Group
*GraphApi* | [**GraphUserGroupTraverseGSuite**](docs/GraphApi.md#graphusergrouptraversegsuite) | **Get** /usergroups/{group_id}/gsuites | List the G Suite instances bound to a User Group
*GraphApi* | [**GraphUserGroupTraverseLdapServer**](docs/GraphApi.md#graphusergrouptraverseldapserver) | **Get** /usergroups/{group_id}/ldapservers | List the LDAP Servers bound to a User Group
*GraphApi* | [**GraphUserGroupTraverseOffice365**](docs/GraphApi.md#graphusergrouptraverseoffice365) | **Get** /usergroups/{group_id}/office365s | List the Office 365 instances bound to a User Group
*GraphApi* | [**GraphUserGroupTraverseRadiusServer**](docs/GraphApi.md#graphusergrouptraverseradiusserver) | **Get** /usergroups/{group_id}/radiusservers | List the RADIUS Servers bound to a User Group
*GraphApi* | [**GraphUserGroupTraverseSystem**](docs/GraphApi.md#graphusergrouptraversesystem) | **Get** /usergroups/{group_id}/systems | List the Systems bound to a User Group
*GraphApi* | [**GraphUserGroupTraverseSystemGroup**](docs/GraphApi.md#graphusergrouptraversesystemgroup) | **Get** /usergroups/{group_id}/systemgroups | List the System Groups bound to User Groups
*GraphApi* | [**GraphUserMemberOf**](docs/GraphApi.md#graphusermemberof) | **Get** /users/{user_id}/memberof | List the parent Groups of a User
*GraphApi* | [**GraphUserTraverseActiveDirectory**](docs/GraphApi.md#graphusertraverseactivedirectory) | **Get** /users/{user_id}/activedirectories | List the Active Directory instances bound to a User
*GraphApi* | [**GraphUserTraverseApplication**](docs/GraphApi.md#graphusertraverseapplication) | **Get** /users/{user_id}/applications | List the Applications bound to a User
*GraphApi* | [**GraphUserTraverseDirectory**](docs/GraphApi.md#graphusertraversedirectory) | **Get** /users/{user_id}/directories | List the Directories bound to a User
*GraphApi* | [**GraphUserTraverseGSuite**](docs/GraphApi.md#graphusertraversegsuite) | **Get** /users/{user_id}/gsuites | List the G Suite instances bound to a User
*GraphApi* | [**GraphUserTraverseLdapServer**](docs/GraphApi.md#graphusertraverseldapserver) | **Get** /users/{user_id}/ldapservers | List the LDAP servers bound to a User
*GraphApi* | [**GraphUserTraverseOffice365**](docs/GraphApi.md#graphusertraverseoffice365) | **Get** /users/{user_id}/office365s | List the Office 365 instances bound to a User
*GraphApi* | [**GraphUserTraverseRadiusServer**](docs/GraphApi.md#graphusertraverseradiusserver) | **Get** /users/{user_id}/radiusservers | List the RADIUS Servers bound to a User
*GraphApi* | [**GraphUserTraverseSystem**](docs/GraphApi.md#graphusertraversesystem) | **Get** /users/{user_id}/systems | List the Systems bound to a User
*GraphApi* | [**GraphUserTraverseSystemGroup**](docs/GraphApi.md#graphusertraversesystemgroup) | **Get** /users/{user_id}/systemgroups | List the System Groups bound to a User
*GraphApi* | [**PolicystatusesSystemsList**](docs/GraphApi.md#policystatusessystemslist) | **Get** /systems/{system_id}/policystatuses | List the policy statuses for a system
*GroupsApi* | [**GroupsList**](docs/GroupsApi.md#groupslist) | **Get** /groups | List All Groups
*IPListsApi* | [**IplistsDelete**](docs/IPListsApi.md#iplistsdelete) | **Delete** /iplists/{id} | Delete an IP list
*IPListsApi* | [**IplistsGet**](docs/IPListsApi.md#iplistsget) | **Get** /iplists/{id} | Get an IP list
*IPListsApi* | [**IplistsList**](docs/IPListsApi.md#iplistslist) | **Get** /iplists | List IP Lists
*IPListsApi* | [**IplistsPatch**](docs/IPListsApi.md#iplistspatch) | **Patch** /iplists/{id} | Update an IP list
*IPListsApi* | [**IplistsPost**](docs/IPListsApi.md#iplistspost) | **Post** /iplists | Create IP List
*IPListsApi* | [**IplistsPut**](docs/IPListsApi.md#iplistsput) | **Put** /iplists/{id} | Replace an IP list
*ImageApi* | [**ApplicationsDeleteLogo**](docs/ImageApi.md#applicationsdeletelogo) | **Delete** /applications/{application_id}/logo | Delete application image
*LDAPServersApi* | [**GraphLdapServerAssociationsList**](docs/LDAPServersApi.md#graphldapserverassociationslist) | **Get** /ldapservers/{ldapserver_id}/associations | List the associations of a LDAP Server
*LDAPServersApi* | [**GraphLdapServerAssociationsPost**](docs/LDAPServersApi.md#graphldapserverassociationspost) | **Post** /ldapservers/{ldapserver_id}/associations | Manage the associations of a LDAP Server
*LDAPServersApi* | [**GraphLdapServerTraverseUser**](docs/LDAPServersApi.md#graphldapservertraverseuser) | **Get** /ldapservers/{ldapserver_id}/users | List the Users bound to a LDAP Server
*LDAPServersApi* | [**GraphLdapServerTraverseUserGroup**](docs/LDAPServersApi.md#graphldapservertraverseusergroup) | **Get** /ldapservers/{ldapserver_id}/usergroups | List the User Groups bound to a LDAP Server
*LDAPServersApi* | [**LdapserversGet**](docs/LDAPServersApi.md#ldapserversget) | **Get** /ldapservers/{id} | Get LDAP Server
*LDAPServersApi* | [**LdapserversList**](docs/LDAPServersApi.md#ldapserverslist) | **Get** /ldapservers | List LDAP Servers
*LDAPServersApi* | [**LdapserversPatch**](docs/LDAPServersApi.md#ldapserverspatch) | **Patch** /ldapservers/{id} | Update existing LDAP server
*LogosApi* | [**LogosGet**](docs/LogosApi.md#logosget) | **Get** /logos/{id} | Get the logo associated with the specified id
*ManagedServiceProviderApi* | [**AdministratorOrganizationsCreateByAdministrator**](docs/ManagedServiceProviderApi.md#administratororganizationscreatebyadministrator) | **Post** /administrators/{id}/organizationlinks | Allow Adminstrator access to an Organization.
*ManagedServiceProviderApi* | [**AdministratorOrganizationsListByAdministrator**](docs/ManagedServiceProviderApi.md#administratororganizationslistbyadministrator) | **Get** /administrators/{id}/organizationlinks | List the association links between an Administrator and Organizations.
*ManagedServiceProviderApi* | [**AdministratorOrganizationsListByOrganization**](docs/ManagedServiceProviderApi.md#administratororganizationslistbyorganization) | **Get** /organizations/{id}/administratorlinks | List the association links between an Organization and Administrators.
*ManagedServiceProviderApi* | [**AdministratorOrganizationsRemoveByAdministrator**](docs/ManagedServiceProviderApi.md#administratororganizationsremovebyadministrator) | **Delete** /administrators/{administrator_id}/organizationlinks/{id} | Remove association between an Administrator and an Organization.
*ManagedServiceProviderApi* | [**ProviderOrganizationsUpdateOrg**](docs/ManagedServiceProviderApi.md#providerorganizationsupdateorg) | **Put** /providers/{provider_id}/organizations/{id} | Update Provider Organization
*ManagedServiceProviderApi* | [**ProvidersGetProvider**](docs/ManagedServiceProviderApi.md#providersgetprovider) | **Get** /providers/{provider_id} | Retrieve Provider
*ManagedServiceProviderApi* | [**ProvidersListAdministrators**](docs/ManagedServiceProviderApi.md#providerslistadministrators) | **Get** /providers/{provider_id}/administrators | List Provider Administrators
*ManagedServiceProviderApi* | [**ProvidersListOrganizations**](docs/ManagedServiceProviderApi.md#providerslistorganizations) | **Get** /providers/{provider_id}/organizations | List Provider Organizations
*ManagedServiceProviderApi* | [**ProvidersPostAdmins**](docs/ManagedServiceProviderApi.md#providerspostadmins) | **Post** /providers/{provider_id}/administrators | Create a new Provider Administrator
*ManagedServiceProviderApi* | [**ProvidersRetrieveInvoice**](docs/ManagedServiceProviderApi.md#providersretrieveinvoice) | **Get** /providers/{provider_id}/invoices/{ID} | Download a provider&#x27;s invoice.
*ManagedServiceProviderApi* | [**ProvidersRetrieveInvoices**](docs/ManagedServiceProviderApi.md#providersretrieveinvoices) | **Get** /providers/{provider_id}/invoices | List a provider&#x27;s invoices.
*Office365Api* | [**GraphOffice365AssociationsList**](docs/Office365Api.md#graphoffice365associationslist) | **Get** /office365s/{office365_id}/associations | List the associations of an Office 365 instance
*Office365Api* | [**GraphOffice365AssociationsPost**](docs/Office365Api.md#graphoffice365associationspost) | **Post** /office365s/{office365_id}/associations | Manage the associations of an Office 365 instance
*Office365Api* | [**GraphOffice365TraverseUser**](docs/Office365Api.md#graphoffice365traverseuser) | **Get** /office365s/{office365_id}/users | List the Users bound to an Office 365 instance
*Office365Api* | [**GraphOffice365TraverseUserGroup**](docs/Office365Api.md#graphoffice365traverseusergroup) | **Get** /office365s/{office365_id}/usergroups | List the User Groups bound to an Office 365 instance
*Office365Api* | [**Office365sGet**](docs/Office365Api.md#office365sget) | **Get** /office365s/{office365_id} | Get Office 365 instance
*Office365Api* | [**Office365sListImportUsers**](docs/Office365Api.md#office365slistimportusers) | **Get** /office365s/{office365_id}/import/users | Get a list of users to import from an Office 365 instance
*Office365Api* | [**Office365sPatch**](docs/Office365Api.md#office365spatch) | **Patch** /office365s/{office365_id} | Update existing Office 365 instance.
*Office365Api* | [**TranslationRulesOffice365Delete**](docs/Office365Api.md#translationrulesoffice365delete) | **Delete** /office365s/{office365_id}/translationrules/{id} | Deletes a Office 365 translation rule
*Office365Api* | [**TranslationRulesOffice365Get**](docs/Office365Api.md#translationrulesoffice365get) | **Get** /office365s/{office365_id}/translationrules/{id} | Gets a specific Office 365 translation rule
*Office365Api* | [**TranslationRulesOffice365List**](docs/Office365Api.md#translationrulesoffice365list) | **Get** /office365s/{office365_id}/translationrules | List all the Office 365 Translation Rules
*Office365Api* | [**TranslationRulesOffice365Post**](docs/Office365Api.md#translationrulesoffice365post) | **Post** /office365s/{office365_id}/translationrules | Create a new Office 365 Translation Rule
*Office365ImportApi* | [**Office365sListImportUsers**](docs/Office365ImportApi.md#office365slistimportusers) | **Get** /office365s/{office365_id}/import/users | Get a list of users to import from an Office 365 instance
*OrganizationsApi* | [**AdministratorOrganizationsCreateByAdministrator**](docs/OrganizationsApi.md#administratororganizationscreatebyadministrator) | **Post** /administrators/{id}/organizationlinks | Allow Adminstrator access to an Organization.
*OrganizationsApi* | [**AdministratorOrganizationsListByAdministrator**](docs/OrganizationsApi.md#administratororganizationslistbyadministrator) | **Get** /administrators/{id}/organizationlinks | List the association links between an Administrator and Organizations.
*OrganizationsApi* | [**AdministratorOrganizationsListByOrganization**](docs/OrganizationsApi.md#administratororganizationslistbyorganization) | **Get** /organizations/{id}/administratorlinks | List the association links between an Organization and Administrators.
*OrganizationsApi* | [**AdministratorOrganizationsRemoveByAdministrator**](docs/OrganizationsApi.md#administratororganizationsremovebyadministrator) | **Delete** /administrators/{administrator_id}/organizationlinks/{id} | Remove association between an Administrator and an Organization.
*OrganizationsApi* | [**OrganizationsListCases**](docs/OrganizationsApi.md#organizationslistcases) | **Get** /organizations/cases | Get all cases (Support/Feature requests) for organization
*PoliciesApi* | [**GraphPolicyAssociationsList**](docs/PoliciesApi.md#graphpolicyassociationslist) | **Get** /policies/{policy_id}/associations | List the associations of a Policy
*PoliciesApi* | [**GraphPolicyAssociationsPost**](docs/PoliciesApi.md#graphpolicyassociationspost) | **Post** /policies/{policy_id}/associations | Manage the associations of a Policy
*PoliciesApi* | [**GraphPolicyMemberOf**](docs/PoliciesApi.md#graphpolicymemberof) | **Get** /policies/{policy_id}/memberof | List the parent Groups of a Policy
*PoliciesApi* | [**GraphPolicyTraverseSystem**](docs/PoliciesApi.md#graphpolicytraversesystem) | **Get** /policies/{policy_id}/systems | List the Systems bound to a Policy
*PoliciesApi* | [**GraphPolicyTraverseSystemGroup**](docs/PoliciesApi.md#graphpolicytraversesystemgroup) | **Get** /policies/{policy_id}/systemgroups | List the System Groups bound to a Policy
*PoliciesApi* | [**PoliciesDelete**](docs/PoliciesApi.md#policiesdelete) | **Delete** /policies/{id} | Deletes a Policy
*PoliciesApi* | [**PoliciesGet**](docs/PoliciesApi.md#policiesget) | **Get** /policies/{id} | Gets a specific Policy.
*PoliciesApi* | [**PoliciesList**](docs/PoliciesApi.md#policieslist) | **Get** /policies | Lists all the Policies
*PoliciesApi* | [**PoliciesPost**](docs/PoliciesApi.md#policiespost) | **Post** /policies | Create a new Policy
*PoliciesApi* | [**PoliciesPut**](docs/PoliciesApi.md#policiesput) | **Put** /policies/{id} | Update an existing Policy
*PoliciesApi* | [**PolicyresultsGet**](docs/PoliciesApi.md#policyresultsget) | **Get** /policyresults/{id} | Get a specific Policy Result.
*PoliciesApi* | [**PolicyresultsList**](docs/PoliciesApi.md#policyresultslist) | **Get** /policies/{policy_id}/policyresults | Lists all the policy results of a policy.
*PoliciesApi* | [**PolicyresultsOrgList**](docs/PoliciesApi.md#policyresultsorglist) | **Get** /policyresults | Lists all of the policy results for an organization.
*PoliciesApi* | [**PolicystatusesPoliciesList**](docs/PoliciesApi.md#policystatusespolicieslist) | **Get** /policies/{policy_id}/policystatuses | Lists the latest policy results of a policy.
*PoliciesApi* | [**PolicystatusesSystemsList**](docs/PoliciesApi.md#policystatusessystemslist) | **Get** /systems/{system_id}/policystatuses | List the policy statuses for a system
*PoliciesApi* | [**PolicytemplatesGet**](docs/PoliciesApi.md#policytemplatesget) | **Get** /policytemplates/{id} | Get a specific Policy Template
*PoliciesApi* | [**PolicytemplatesList**](docs/PoliciesApi.md#policytemplateslist) | **Get** /policytemplates | Lists all of the Policy Templates
*PolicyGroupAssociationsApi* | [**GraphPolicyGroupAssociationsList**](docs/PolicyGroupAssociationsApi.md#graphpolicygroupassociationslist) | **Get** /policygroups/{group_id}/associations | List the associations of a Policy Group.
*PolicyGroupAssociationsApi* | [**GraphPolicyGroupAssociationsPost**](docs/PolicyGroupAssociationsApi.md#graphpolicygroupassociationspost) | **Post** /policygroups/{group_id}/associations | Manage the associations of a Policy Group
*PolicyGroupAssociationsApi* | [**GraphPolicyGroupTraverseSystem**](docs/PolicyGroupAssociationsApi.md#graphpolicygrouptraversesystem) | **Get** /policygroups/{group_id}/systems | List the Systems bound to a Policy Group
*PolicyGroupAssociationsApi* | [**GraphPolicyGroupTraverseSystemGroup**](docs/PolicyGroupAssociationsApi.md#graphpolicygrouptraversesystemgroup) | **Get** /policygroups/{group_id}/systemgroups | List the System Groups bound to Policy Groups
*PolicyGroupMembersMembershipApi* | [**GraphPolicyGroupMembersList**](docs/PolicyGroupMembersMembershipApi.md#graphpolicygroupmemberslist) | **Get** /policygroups/{group_id}/members | List the members of a Policy Group
*PolicyGroupMembersMembershipApi* | [**GraphPolicyGroupMembersPost**](docs/PolicyGroupMembersMembershipApi.md#graphpolicygroupmemberspost) | **Post** /policygroups/{group_id}/members | Manage the members of a Policy Group
*PolicyGroupMembersMembershipApi* | [**GraphPolicyGroupMembership**](docs/PolicyGroupMembersMembershipApi.md#graphpolicygroupmembership) | **Get** /policygroups/{group_id}/membership | List the Policy Group&#x27;s membership
*PolicyGroupsApi* | [**GraphPolicyGroupAssociationsList**](docs/PolicyGroupsApi.md#graphpolicygroupassociationslist) | **Get** /policygroups/{group_id}/associations | List the associations of a Policy Group.
*PolicyGroupsApi* | [**GraphPolicyGroupAssociationsPost**](docs/PolicyGroupsApi.md#graphpolicygroupassociationspost) | **Post** /policygroups/{group_id}/associations | Manage the associations of a Policy Group
*PolicyGroupsApi* | [**GraphPolicyGroupMembersList**](docs/PolicyGroupsApi.md#graphpolicygroupmemberslist) | **Get** /policygroups/{group_id}/members | List the members of a Policy Group
*PolicyGroupsApi* | [**GraphPolicyGroupMembersPost**](docs/PolicyGroupsApi.md#graphpolicygroupmemberspost) | **Post** /policygroups/{group_id}/members | Manage the members of a Policy Group
*PolicyGroupsApi* | [**GraphPolicyGroupMembership**](docs/PolicyGroupsApi.md#graphpolicygroupmembership) | **Get** /policygroups/{group_id}/membership | List the Policy Group&#x27;s membership
*PolicyGroupsApi* | [**GraphPolicyGroupTraverseSystem**](docs/PolicyGroupsApi.md#graphpolicygrouptraversesystem) | **Get** /policygroups/{group_id}/systems | List the Systems bound to a Policy Group
*PolicyGroupsApi* | [**GraphPolicyGroupTraverseSystemGroup**](docs/PolicyGroupsApi.md#graphpolicygrouptraversesystemgroup) | **Get** /policygroups/{group_id}/systemgroups | List the System Groups bound to Policy Groups
*PolicyGroupsApi* | [**GroupsPolicyDelete**](docs/PolicyGroupsApi.md#groupspolicydelete) | **Delete** /policygroups/{id} | Delete a Policy Group
*PolicyGroupsApi* | [**GroupsPolicyGet**](docs/PolicyGroupsApi.md#groupspolicyget) | **Get** /policygroups/{id} | View an individual Policy Group details
*PolicyGroupsApi* | [**GroupsPolicyList**](docs/PolicyGroupsApi.md#groupspolicylist) | **Get** /policygroups | List all Policy Groups
*PolicyGroupsApi* | [**GroupsPolicyPost**](docs/PolicyGroupsApi.md#groupspolicypost) | **Post** /policygroups | Create a new Policy Group
*PolicyGroupsApi* | [**GroupsPolicyPut**](docs/PolicyGroupsApi.md#groupspolicyput) | **Put** /policygroups/{id} | Update a Policy Group
*PolicytemplatesApi* | [**PolicytemplatesGet**](docs/PolicytemplatesApi.md#policytemplatesget) | **Get** /policytemplates/{id} | Get a specific Policy Template
*PolicytemplatesApi* | [**PolicytemplatesList**](docs/PolicytemplatesApi.md#policytemplateslist) | **Get** /policytemplates | Lists all of the Policy Templates
*ProvidersApi* | [**AutotaskCreateConfiguration**](docs/ProvidersApi.md#autotaskcreateconfiguration) | **Post** /providers/{provider_id}/integrations/autotask | Creates a new Autotask integration for the provider
*ProvidersApi* | [**AutotaskDeleteConfiguration**](docs/ProvidersApi.md#autotaskdeleteconfiguration) | **Delete** /integrations/autotask/{UUID} | Delete Autotask Integration
*ProvidersApi* | [**AutotaskGetConfiguration**](docs/ProvidersApi.md#autotaskgetconfiguration) | **Get** /integrations/autotask/{UUID} | Retrieve Autotask Integration Configuration
*ProvidersApi* | [**AutotaskPatchMappings**](docs/ProvidersApi.md#autotaskpatchmappings) | **Patch** /integrations/autotask/{UUID}/mappings | Create, edit, and/or delete Autotask Mappings
*ProvidersApi* | [**AutotaskPatchSettings**](docs/ProvidersApi.md#autotaskpatchsettings) | **Patch** /integrations/autotask/{UUID}/settings | Create, edit, and/or delete Autotask Integration settings
*ProvidersApi* | [**AutotaskRetrieveAllAlertConfigurationOptions**](docs/ProvidersApi.md#autotaskretrieveallalertconfigurationoptions) | **Get** /providers/{provider_id}/integrations/autotask/alerts/configuration/options | Get all Autotask ticketing alert configuration options for a provider
*ProvidersApi* | [**AutotaskRetrieveAllAlertConfigurations**](docs/ProvidersApi.md#autotaskretrieveallalertconfigurations) | **Get** /providers/{provider_id}/integrations/autotask/alerts/configuration | Get all Autotask ticketing alert configurations for a provider
*ProvidersApi* | [**AutotaskRetrieveCompanies**](docs/ProvidersApi.md#autotaskretrievecompanies) | **Get** /integrations/autotask/{UUID}/companies | Retrieve Autotask Companies
*ProvidersApi* | [**AutotaskRetrieveCompanyTypes**](docs/ProvidersApi.md#autotaskretrievecompanytypes) | **Get** /integrations/autotask/{UUID}/companytypes | Retrieve Autotask Company Types
*ProvidersApi* | [**AutotaskRetrieveContracts**](docs/ProvidersApi.md#autotaskretrievecontracts) | **Get** /integrations/autotask/{UUID}/contracts | Retrieve Autotask Contracts
*ProvidersApi* | [**AutotaskRetrieveContractsFields**](docs/ProvidersApi.md#autotaskretrievecontractsfields) | **Get** /integrations/autotask/{UUID}/contracts/fields | Retrieve Autotask Contract Fields
*ProvidersApi* | [**AutotaskRetrieveMappings**](docs/ProvidersApi.md#autotaskretrievemappings) | **Get** /integrations/autotask/{UUID}/mappings | Retrieve Autotask mappings
*ProvidersApi* | [**AutotaskRetrieveServices**](docs/ProvidersApi.md#autotaskretrieveservices) | **Get** /integrations/autotask/{UUID}/contracts/services | Retrieve Autotask Contract Services
*ProvidersApi* | [**AutotaskRetrieveSettings**](docs/ProvidersApi.md#autotaskretrievesettings) | **Get** /integrations/autotask/{UUID}/settings | Retrieve Autotask Integration settings
*ProvidersApi* | [**AutotaskUpdateAlertConfiguration**](docs/ProvidersApi.md#autotaskupdatealertconfiguration) | **Put** /providers/{provider_id}/integrations/autotask/alerts/{alert_UUID}/configuration | Update an Autotask ticketing alert&#x27;s configuration
*ProvidersApi* | [**AutotaskUpdateConfiguration**](docs/ProvidersApi.md#autotaskupdateconfiguration) | **Patch** /integrations/autotask/{UUID} | Update Autotask Integration configuration
*ProvidersApi* | [**ConnectwiseCreateConfiguration**](docs/ProvidersApi.md#connectwisecreateconfiguration) | **Post** /providers/{provider_id}/integrations/connectwise | Creates a new ConnectWise integration for the provider
*ProvidersApi* | [**ConnectwiseDeleteConfiguration**](docs/ProvidersApi.md#connectwisedeleteconfiguration) | **Delete** /integrations/connectwise/{UUID} | Delete ConnectWise Integration
*ProvidersApi* | [**ConnectwiseGetConfiguration**](docs/ProvidersApi.md#connectwisegetconfiguration) | **Get** /integrations/connectwise/{UUID} | Retrieve ConnectWise Integration Configuration
*ProvidersApi* | [**ConnectwisePatchMappings**](docs/ProvidersApi.md#connectwisepatchmappings) | **Patch** /integrations/connectwise/{UUID}/mappings | Create, edit, and/or delete ConnectWise Mappings
*ProvidersApi* | [**ConnectwisePatchSettings**](docs/ProvidersApi.md#connectwisepatchsettings) | **Patch** /integrations/connectwise/{UUID}/settings | Create, edit, and/or delete ConnectWise Integration settings
*ProvidersApi* | [**ConnectwiseRetrieveAdditions**](docs/ProvidersApi.md#connectwiseretrieveadditions) | **Get** /integrations/connectwise/{UUID}/agreements/{agreement_ID}/additions | Retrieve ConnectWise Additions
*ProvidersApi* | [**ConnectwiseRetrieveAgreements**](docs/ProvidersApi.md#connectwiseretrieveagreements) | **Get** /integrations/connectwise/{UUID}/agreements | Retrieve ConnectWise Agreements
*ProvidersApi* | [**ConnectwiseRetrieveAllAlertConfigurationOptions**](docs/ProvidersApi.md#connectwiseretrieveallalertconfigurationoptions) | **Get** /providers/{provider_id}/integrations/connectwise/alerts/configuration/options | Get all ConnectWise ticketing alert configuration options for a provider
*ProvidersApi* | [**ConnectwiseRetrieveAllAlertConfigurations**](docs/ProvidersApi.md#connectwiseretrieveallalertconfigurations) | **Get** /providers/{provider_id}/integrations/connectwise/alerts/configuration | Get all ConnectWise ticketing alert configurations for a provider
*ProvidersApi* | [**ConnectwiseRetrieveCompanies**](docs/ProvidersApi.md#connectwiseretrievecompanies) | **Get** /integrations/connectwise/{UUID}/companies | Retrieve ConnectWise Companies
*ProvidersApi* | [**ConnectwiseRetrieveCompanyTypes**](docs/ProvidersApi.md#connectwiseretrievecompanytypes) | **Get** /integrations/connectwise/{UUID}/companytypes | Retrieve ConnectWise Company Types
*ProvidersApi* | [**ConnectwiseRetrieveMappings**](docs/ProvidersApi.md#connectwiseretrievemappings) | **Get** /integrations/connectwise/{UUID}/mappings | Retrieve ConnectWise mappings
*ProvidersApi* | [**ConnectwiseRetrieveSettings**](docs/ProvidersApi.md#connectwiseretrievesettings) | **Get** /integrations/connectwise/{UUID}/settings | Retrieve ConnectWise Integration settings
*ProvidersApi* | [**ConnectwiseUpdateAlertConfiguration**](docs/ProvidersApi.md#connectwiseupdatealertconfiguration) | **Put** /providers/{provider_id}/integrations/connectwise/alerts/{alert_UUID}/configuration | Update a ConnectWise ticketing alert&#x27;s configuration
*ProvidersApi* | [**ConnectwiseUpdateConfiguration**](docs/ProvidersApi.md#connectwiseupdateconfiguration) | **Patch** /integrations/connectwise/{UUID} | Update ConnectWise Integration configuration
*ProvidersApi* | [**MtpIntegrationRetrieveAlerts**](docs/ProvidersApi.md#mtpintegrationretrievealerts) | **Get** /providers/{provider_id}/integrations/ticketing/alerts | Get all ticketing alerts available for a provider&#x27;s ticketing integration.
*ProvidersApi* | [**MtpIntegrationRetrieveSyncErrors**](docs/ProvidersApi.md#mtpintegrationretrievesyncerrors) | **Get** /integrations/{integration_type}/{UUID}/errors | Retrieve Recent Integration Sync Errors
*ProvidersApi* | [**ProviderOrganizationsUpdateOrg**](docs/ProvidersApi.md#providerorganizationsupdateorg) | **Put** /providers/{provider_id}/organizations/{id} | Update Provider Organization
*ProvidersApi* | [**ProvidersGetProvider**](docs/ProvidersApi.md#providersgetprovider) | **Get** /providers/{provider_id} | Retrieve Provider
*ProvidersApi* | [**ProvidersListAdministrators**](docs/ProvidersApi.md#providerslistadministrators) | **Get** /providers/{provider_id}/administrators | List Provider Administrators
*ProvidersApi* | [**ProvidersListOrganizations**](docs/ProvidersApi.md#providerslistorganizations) | **Get** /providers/{provider_id}/organizations | List Provider Organizations
*ProvidersApi* | [**ProvidersPostAdmins**](docs/ProvidersApi.md#providerspostadmins) | **Post** /providers/{provider_id}/administrators | Create a new Provider Administrator
*ProvidersApi* | [**ProvidersRemoveAdministrator**](docs/ProvidersApi.md#providersremoveadministrator) | **Delete** /providers/{provider_id}/administrators/{id} | Delete Provider Administrator
*ProvidersApi* | [**ProvidersRetrieveIntegrations**](docs/ProvidersApi.md#providersretrieveintegrations) | **Get** /providers/{provider_id}/integrations | Retrieve Integrations for Provider
*ProvidersApi* | [**ProvidersRetrieveInvoice**](docs/ProvidersApi.md#providersretrieveinvoice) | **Get** /providers/{provider_id}/invoices/{ID} | Download a provider&#x27;s invoice.
*ProvidersApi* | [**ProvidersRetrieveInvoices**](docs/ProvidersApi.md#providersretrieveinvoices) | **Get** /providers/{provider_id}/invoices | List a provider&#x27;s invoices.
*RADIUSServersApi* | [**GraphRadiusServerAssociationsList**](docs/RADIUSServersApi.md#graphradiusserverassociationslist) | **Get** /radiusservers/{radiusserver_id}/associations | List the associations of a RADIUS  Server
*RADIUSServersApi* | [**GraphRadiusServerAssociationsPost**](docs/RADIUSServersApi.md#graphradiusserverassociationspost) | **Post** /radiusservers/{radiusserver_id}/associations | Manage the associations of a RADIUS Server
*RADIUSServersApi* | [**GraphRadiusServerTraverseUser**](docs/RADIUSServersApi.md#graphradiusservertraverseuser) | **Get** /radiusservers/{radiusserver_id}/users | List the Users bound to a RADIUS  Server
*RADIUSServersApi* | [**GraphRadiusServerTraverseUserGroup**](docs/RADIUSServersApi.md#graphradiusservertraverseusergroup) | **Get** /radiusservers/{radiusserver_id}/usergroups | List the User Groups bound to a RADIUS  Server
*SCIMImportApi* | [**ImportUsers**](docs/SCIMImportApi.md#importusers) | **Get** /applications/{application_id}/import/users | Get a list of users to import from an Application IdM service provider
*SambaDomainsApi* | [**LdapserversSambaDomainsDelete**](docs/SambaDomainsApi.md#ldapserverssambadomainsdelete) | **Delete** /ldapservers/{ldapserver_id}/sambadomains/{id} | Delete Samba Domain
*SambaDomainsApi* | [**LdapserversSambaDomainsGet**](docs/SambaDomainsApi.md#ldapserverssambadomainsget) | **Get** /ldapservers/{ldapserver_id}/sambadomains/{id} | Get Samba Domain
*SambaDomainsApi* | [**LdapserversSambaDomainsList**](docs/SambaDomainsApi.md#ldapserverssambadomainslist) | **Get** /ldapservers/{ldapserver_id}/sambadomains | List Samba Domains
*SambaDomainsApi* | [**LdapserversSambaDomainsPost**](docs/SambaDomainsApi.md#ldapserverssambadomainspost) | **Post** /ldapservers/{ldapserver_id}/sambadomains | Create Samba Domain
*SambaDomainsApi* | [**LdapserversSambaDomainsPut**](docs/SambaDomainsApi.md#ldapserverssambadomainsput) | **Put** /ldapservers/{ldapserver_id}/sambadomains/{id} | Update Samba Domain
*SoftwareAppsApi* | [**GraphSoftwareappsAssociationsList**](docs/SoftwareAppsApi.md#graphsoftwareappsassociationslist) | **Get** /softwareapps/{software_app_id}/associations | List the associations of a Software Application
*SoftwareAppsApi* | [**GraphSoftwareappsAssociationsPost**](docs/SoftwareAppsApi.md#graphsoftwareappsassociationspost) | **Post** /softwareapps/{software_app_id}/associations | Manage the associations of a software application.
*SoftwareAppsApi* | [**GraphSoftwareappsTraverseSystem**](docs/SoftwareAppsApi.md#graphsoftwareappstraversesystem) | **Get** /softwareapps/{software_app_id}/systems | List the Systems bound to a Software App.
*SoftwareAppsApi* | [**GraphSoftwareappsTraverseSystemGroup**](docs/SoftwareAppsApi.md#graphsoftwareappstraversesystemgroup) | **Get** /softwareapps/{software_app_id}/systemgroups | List the System Groups bound to a Software App.
*SoftwareAppsApi* | [**SoftwareAppStatusesList**](docs/SoftwareAppsApi.md#softwareappstatuseslist) | **Get** /softwareapps/{software_app_id}/statuses | Get the status of the provided Software Application
*SoftwareAppsApi* | [**SoftwareAppsDelete**](docs/SoftwareAppsApi.md#softwareappsdelete) | **Delete** /softwareapps/{id} | Delete a configured Software Application
*SoftwareAppsApi* | [**SoftwareAppsGet**](docs/SoftwareAppsApi.md#softwareappsget) | **Get** /softwareapps/{id} | Retrieve a configured Software Application.
*SoftwareAppsApi* | [**SoftwareAppsList**](docs/SoftwareAppsApi.md#softwareappslist) | **Get** /softwareapps | Get all configured Software Applications.
*SoftwareAppsApi* | [**SoftwareAppsPost**](docs/SoftwareAppsApi.md#softwareappspost) | **Post** /softwareapps | Create a Software Application that will be managed by JumpCloud.
*SoftwareAppsApi* | [**SoftwareAppsReclaimLicenses**](docs/SoftwareAppsApi.md#softwareappsreclaimlicenses) | **Post** /softwareapps/{software_app_id}/reclaim-licenses | Reclaim Licenses for a Software Application.
*SoftwareAppsApi* | [**SoftwareAppsRetryInstallation**](docs/SoftwareAppsApi.md#softwareappsretryinstallation) | **Post** /softwareapps/{software_app_id}/retry-installation | Retry Installation for a Software Application
*SoftwareAppsApi* | [**SoftwareAppsUpdate**](docs/SoftwareAppsApi.md#softwareappsupdate) | **Put** /softwareapps/{id} | Update a Software Application Configuration.
*SubscriptionsApi* | [**SubscriptionsGet**](docs/SubscriptionsApi.md#subscriptionsget) | **Get** /subscriptions | Lists all the Pricing &amp; Packaging Subscriptions
*SystemGroupAssociationsApi* | [**GraphSystemGroupAssociationsList**](docs/SystemGroupAssociationsApi.md#graphsystemgroupassociationslist) | **Get** /systemgroups/{group_id}/associations | List the associations of a System Group
*SystemGroupAssociationsApi* | [**GraphSystemGroupAssociationsPost**](docs/SystemGroupAssociationsApi.md#graphsystemgroupassociationspost) | **Post** /systemgroups/{group_id}/associations | Manage the associations of a System Group
*SystemGroupAssociationsApi* | [**GraphSystemGroupTraverseCommand**](docs/SystemGroupAssociationsApi.md#graphsystemgrouptraversecommand) | **Get** /systemgroups/{group_id}/commands | List the Commands bound to a System Group
*SystemGroupAssociationsApi* | [**GraphSystemGroupTraversePolicy**](docs/SystemGroupAssociationsApi.md#graphsystemgrouptraversepolicy) | **Get** /systemgroups/{group_id}/policies | List the Policies bound to a System Group
*SystemGroupAssociationsApi* | [**GraphSystemGroupTraversePolicyGroup**](docs/SystemGroupAssociationsApi.md#graphsystemgrouptraversepolicygroup) | **Get** /systemgroups/{group_id}/policygroups | List the Policy Groups bound to a System Group
*SystemGroupAssociationsApi* | [**GraphSystemGroupTraverseUser**](docs/SystemGroupAssociationsApi.md#graphsystemgrouptraverseuser) | **Get** /systemgroups/{group_id}/users | List the Users bound to a System Group
*SystemGroupAssociationsApi* | [**GraphSystemGroupTraverseUserGroup**](docs/SystemGroupAssociationsApi.md#graphsystemgrouptraverseusergroup) | **Get** /systemgroups/{group_id}/usergroups | List the User Groups bound to a System Group
*SystemGroupMembersMembershipApi* | [**GraphSystemGroupMembersList**](docs/SystemGroupMembersMembershipApi.md#graphsystemgroupmemberslist) | **Get** /systemgroups/{group_id}/members | List the members of a System Group
*SystemGroupMembersMembershipApi* | [**GraphSystemGroupMembersPost**](docs/SystemGroupMembersMembershipApi.md#graphsystemgroupmemberspost) | **Post** /systemgroups/{group_id}/members | Manage the members of a System Group
*SystemGroupMembersMembershipApi* | [**GraphSystemGroupMembership**](docs/SystemGroupMembersMembershipApi.md#graphsystemgroupmembership) | **Get** /systemgroups/{group_id}/membership | List the System Group&#x27;s membership
*SystemGroupsApi* | [**GraphSystemGroupAssociationsList**](docs/SystemGroupsApi.md#graphsystemgroupassociationslist) | **Get** /systemgroups/{group_id}/associations | List the associations of a System Group
*SystemGroupsApi* | [**GraphSystemGroupAssociationsPost**](docs/SystemGroupsApi.md#graphsystemgroupassociationspost) | **Post** /systemgroups/{group_id}/associations | Manage the associations of a System Group
*SystemGroupsApi* | [**GraphSystemGroupMembersList**](docs/SystemGroupsApi.md#graphsystemgroupmemberslist) | **Get** /systemgroups/{group_id}/members | List the members of a System Group
*SystemGroupsApi* | [**GraphSystemGroupMembersPost**](docs/SystemGroupsApi.md#graphsystemgroupmemberspost) | **Post** /systemgroups/{group_id}/members | Manage the members of a System Group
*SystemGroupsApi* | [**GraphSystemGroupMembership**](docs/SystemGroupsApi.md#graphsystemgroupmembership) | **Get** /systemgroups/{group_id}/membership | List the System Group&#x27;s membership
*SystemGroupsApi* | [**GraphSystemGroupTraversePolicy**](docs/SystemGroupsApi.md#graphsystemgrouptraversepolicy) | **Get** /systemgroups/{group_id}/policies | List the Policies bound to a System Group
*SystemGroupsApi* | [**GraphSystemGroupTraversePolicyGroup**](docs/SystemGroupsApi.md#graphsystemgrouptraversepolicygroup) | **Get** /systemgroups/{group_id}/policygroups | List the Policy Groups bound to a System Group
*SystemGroupsApi* | [**GraphSystemGroupTraverseUser**](docs/SystemGroupsApi.md#graphsystemgrouptraverseuser) | **Get** /systemgroups/{group_id}/users | List the Users bound to a System Group
*SystemGroupsApi* | [**GraphSystemGroupTraverseUserGroup**](docs/SystemGroupsApi.md#graphsystemgrouptraverseusergroup) | **Get** /systemgroups/{group_id}/usergroups | List the User Groups bound to a System Group
*SystemGroupsApi* | [**GroupsSystemDelete**](docs/SystemGroupsApi.md#groupssystemdelete) | **Delete** /systemgroups/{id} | Delete a System Group
*SystemGroupsApi* | [**GroupsSystemGet**](docs/SystemGroupsApi.md#groupssystemget) | **Get** /systemgroups/{id} | View an individual System Group details
*SystemGroupsApi* | [**GroupsSystemList**](docs/SystemGroupsApi.md#groupssystemlist) | **Get** /systemgroups | List all System Groups
*SystemGroupsApi* | [**GroupsSystemPost**](docs/SystemGroupsApi.md#groupssystempost) | **Post** /systemgroups | Create a new System Group
*SystemGroupsApi* | [**GroupsSystemPut**](docs/SystemGroupsApi.md#groupssystemput) | **Put** /systemgroups/{id} | Update a System Group
*SystemInsightsApi* | [**SysteminsightsListAlf**](docs/SystemInsightsApi.md#systeminsightslistalf) | **Get** /systeminsights/alf | List System Insights ALF
*SystemInsightsApi* | [**SysteminsightsListAlfExceptions**](docs/SystemInsightsApi.md#systeminsightslistalfexceptions) | **Get** /systeminsights/alf_exceptions | List System Insights ALF Exceptions
*SystemInsightsApi* | [**SysteminsightsListAlfExplicitAuths**](docs/SystemInsightsApi.md#systeminsightslistalfexplicitauths) | **Get** /systeminsights/alf_explicit_auths | List System Insights ALF Explicit Authentications
*SystemInsightsApi* | [**SysteminsightsListAppcompatShims**](docs/SystemInsightsApi.md#systeminsightslistappcompatshims) | **Get** /systeminsights/appcompat_shims | List System Insights Application Compatibility Shims
*SystemInsightsApi* | [**SysteminsightsListApps**](docs/SystemInsightsApi.md#systeminsightslistapps) | **Get** /systeminsights/apps | List System Insights Apps
*SystemInsightsApi* | [**SysteminsightsListAuthorizedKeys**](docs/SystemInsightsApi.md#systeminsightslistauthorizedkeys) | **Get** /systeminsights/authorized_keys | List System Insights Authorized Keys
*SystemInsightsApi* | [**SysteminsightsListAzureInstanceMetadata**](docs/SystemInsightsApi.md#systeminsightslistazureinstancemetadata) | **Get** /systeminsights/azure_instance_metadata | List System Insights Azure Instance Metadata
*SystemInsightsApi* | [**SysteminsightsListAzureInstanceTags**](docs/SystemInsightsApi.md#systeminsightslistazureinstancetags) | **Get** /systeminsights/azure_instance_tags | List System Insights Azure Instance Tags
*SystemInsightsApi* | [**SysteminsightsListBattery**](docs/SystemInsightsApi.md#systeminsightslistbattery) | **Get** /systeminsights/battery | List System Insights Battery
*SystemInsightsApi* | [**SysteminsightsListBitlockerInfo**](docs/SystemInsightsApi.md#systeminsightslistbitlockerinfo) | **Get** /systeminsights/bitlocker_info | List System Insights Bitlocker Info
*SystemInsightsApi* | [**SysteminsightsListBrowserPlugins**](docs/SystemInsightsApi.md#systeminsightslistbrowserplugins) | **Get** /systeminsights/browser_plugins | List System Insights Browser Plugins
*SystemInsightsApi* | [**SysteminsightsListCertificates**](docs/SystemInsightsApi.md#systeminsightslistcertificates) | **Get** /systeminsights/certificates | List System Insights Certificates
*SystemInsightsApi* | [**SysteminsightsListChassisInfo**](docs/SystemInsightsApi.md#systeminsightslistchassisinfo) | **Get** /systeminsights/chassis_info | List System Insights Chassis Info
*SystemInsightsApi* | [**SysteminsightsListChromeExtensions**](docs/SystemInsightsApi.md#systeminsightslistchromeextensions) | **Get** /systeminsights/chrome_extensions | List System Insights Chrome Extensions
*SystemInsightsApi* | [**SysteminsightsListConnectivity**](docs/SystemInsightsApi.md#systeminsightslistconnectivity) | **Get** /systeminsights/connectivity | List System Insights Connectivity
*SystemInsightsApi* | [**SysteminsightsListCrashes**](docs/SystemInsightsApi.md#systeminsightslistcrashes) | **Get** /systeminsights/crashes | List System Insights Crashes
*SystemInsightsApi* | [**SysteminsightsListCupsDestinations**](docs/SystemInsightsApi.md#systeminsightslistcupsdestinations) | **Get** /systeminsights/cups_destinations | List System Insights CUPS Destinations
*SystemInsightsApi* | [**SysteminsightsListDiskEncryption**](docs/SystemInsightsApi.md#systeminsightslistdiskencryption) | **Get** /systeminsights/disk_encryption | List System Insights Disk Encryption
*SystemInsightsApi* | [**SysteminsightsListDiskInfo**](docs/SystemInsightsApi.md#systeminsightslistdiskinfo) | **Get** /systeminsights/disk_info | List System Insights Disk Info
*SystemInsightsApi* | [**SysteminsightsListDnsResolvers**](docs/SystemInsightsApi.md#systeminsightslistdnsresolvers) | **Get** /systeminsights/dns_resolvers | List System Insights DNS Resolvers
*SystemInsightsApi* | [**SysteminsightsListEtcHosts**](docs/SystemInsightsApi.md#systeminsightslistetchosts) | **Get** /systeminsights/etc_hosts | List System Insights Etc Hosts
*SystemInsightsApi* | [**SysteminsightsListFirefoxAddons**](docs/SystemInsightsApi.md#systeminsightslistfirefoxaddons) | **Get** /systeminsights/firefox_addons | List System Insights Firefox Addons
*SystemInsightsApi* | [**SysteminsightsListGroups**](docs/SystemInsightsApi.md#systeminsightslistgroups) | **Get** /systeminsights/groups | List System Insights Groups
*SystemInsightsApi* | [**SysteminsightsListIeExtensions**](docs/SystemInsightsApi.md#systeminsightslistieextensions) | **Get** /systeminsights/ie_extensions | List System Insights IE Extensions
*SystemInsightsApi* | [**SysteminsightsListInterfaceAddresses**](docs/SystemInsightsApi.md#systeminsightslistinterfaceaddresses) | **Get** /systeminsights/interface_addresses | List System Insights Interface Addresses
*SystemInsightsApi* | [**SysteminsightsListInterfaceDetails**](docs/SystemInsightsApi.md#systeminsightslistinterfacedetails) | **Get** /systeminsights/interface_details | List System Insights Interface Details
*SystemInsightsApi* | [**SysteminsightsListKernelInfo**](docs/SystemInsightsApi.md#systeminsightslistkernelinfo) | **Get** /systeminsights/kernel_info | List System Insights Kernel Info
*SystemInsightsApi* | [**SysteminsightsListLaunchd**](docs/SystemInsightsApi.md#systeminsightslistlaunchd) | **Get** /systeminsights/launchd | List System Insights Launchd
*SystemInsightsApi* | [**SysteminsightsListLinuxPackages**](docs/SystemInsightsApi.md#systeminsightslistlinuxpackages) | **Get** /systeminsights/linux_packages | List System Insights Linux Packages
*SystemInsightsApi* | [**SysteminsightsListLoggedInUsers**](docs/SystemInsightsApi.md#systeminsightslistloggedinusers) | **Get** /systeminsights/logged_in_users | List System Insights Logged-In Users
*SystemInsightsApi* | [**SysteminsightsListLogicalDrives**](docs/SystemInsightsApi.md#systeminsightslistlogicaldrives) | **Get** /systeminsights/logical_drives | List System Insights Logical Drives
*SystemInsightsApi* | [**SysteminsightsListManagedPolicies**](docs/SystemInsightsApi.md#systeminsightslistmanagedpolicies) | **Get** /systeminsights/managed_policies | List System Insights Managed Policies
*SystemInsightsApi* | [**SysteminsightsListMounts**](docs/SystemInsightsApi.md#systeminsightslistmounts) | **Get** /systeminsights/mounts | List System Insights Mounts
*SystemInsightsApi* | [**SysteminsightsListOsVersion**](docs/SystemInsightsApi.md#systeminsightslistosversion) | **Get** /systeminsights/os_version | List System Insights OS Version
*SystemInsightsApi* | [**SysteminsightsListPatches**](docs/SystemInsightsApi.md#systeminsightslistpatches) | **Get** /systeminsights/patches | List System Insights Patches
*SystemInsightsApi* | [**SysteminsightsListPrograms**](docs/SystemInsightsApi.md#systeminsightslistprograms) | **Get** /systeminsights/programs | List System Insights Programs
*SystemInsightsApi* | [**SysteminsightsListPythonPackages**](docs/SystemInsightsApi.md#systeminsightslistpythonpackages) | **Get** /systeminsights/python_packages | List System Insights Python Packages
*SystemInsightsApi* | [**SysteminsightsListSafariExtensions**](docs/SystemInsightsApi.md#systeminsightslistsafariextensions) | **Get** /systeminsights/safari_extensions | List System Insights Safari Extensions
*SystemInsightsApi* | [**SysteminsightsListScheduledTasks**](docs/SystemInsightsApi.md#systeminsightslistscheduledtasks) | **Get** /systeminsights/scheduled_tasks | List System Insights Scheduled Tasks
*SystemInsightsApi* | [**SysteminsightsListSecureboot**](docs/SystemInsightsApi.md#systeminsightslistsecureboot) | **Get** /systeminsights/secureboot | List System Insights Secure Boot
*SystemInsightsApi* | [**SysteminsightsListServices**](docs/SystemInsightsApi.md#systeminsightslistservices) | **Get** /systeminsights/services | List System Insights Services
*SystemInsightsApi* | [**SysteminsightsListShadow**](docs/SystemInsightsApi.md#systeminsightslistshadow) | **Get** /systeminsights/shadow | LIst System Insights Shadow
*SystemInsightsApi* | [**SysteminsightsListSharedFolders**](docs/SystemInsightsApi.md#systeminsightslistsharedfolders) | **Get** /systeminsights/shared_folders | List System Insights Shared Folders
*SystemInsightsApi* | [**SysteminsightsListSharedResources**](docs/SystemInsightsApi.md#systeminsightslistsharedresources) | **Get** /systeminsights/shared_resources | List System Insights Shared Resources
*SystemInsightsApi* | [**SysteminsightsListSharingPreferences**](docs/SystemInsightsApi.md#systeminsightslistsharingpreferences) | **Get** /systeminsights/sharing_preferences | List System Insights Sharing Preferences
*SystemInsightsApi* | [**SysteminsightsListSipConfig**](docs/SystemInsightsApi.md#systeminsightslistsipconfig) | **Get** /systeminsights/sip_config | List System Insights SIP Config
*SystemInsightsApi* | [**SysteminsightsListStartupItems**](docs/SystemInsightsApi.md#systeminsightsliststartupitems) | **Get** /systeminsights/startup_items | List System Insights Startup Items
*SystemInsightsApi* | [**SysteminsightsListSystemControls**](docs/SystemInsightsApi.md#systeminsightslistsystemcontrols) | **Get** /systeminsights/system_controls | List System Insights System Control
*SystemInsightsApi* | [**SysteminsightsListSystemInfo**](docs/SystemInsightsApi.md#systeminsightslistsysteminfo) | **Get** /systeminsights/system_info | List System Insights System Info
*SystemInsightsApi* | [**SysteminsightsListTpmInfo**](docs/SystemInsightsApi.md#systeminsightslisttpminfo) | **Get** /systeminsights/tpm_info | List System Insights TPM Info
*SystemInsightsApi* | [**SysteminsightsListUptime**](docs/SystemInsightsApi.md#systeminsightslistuptime) | **Get** /systeminsights/uptime | List System Insights Uptime
*SystemInsightsApi* | [**SysteminsightsListUsbDevices**](docs/SystemInsightsApi.md#systeminsightslistusbdevices) | **Get** /systeminsights/usb_devices | List System Insights USB Devices
*SystemInsightsApi* | [**SysteminsightsListUserGroups**](docs/SystemInsightsApi.md#systeminsightslistusergroups) | **Get** /systeminsights/user_groups | List System Insights User Groups
*SystemInsightsApi* | [**SysteminsightsListUserSshKeys**](docs/SystemInsightsApi.md#systeminsightslistusersshkeys) | **Get** /systeminsights/user_ssh_keys | List System Insights User SSH Keys
*SystemInsightsApi* | [**SysteminsightsListUserassist**](docs/SystemInsightsApi.md#systeminsightslistuserassist) | **Get** /systeminsights/userassist | List System Insights User Assist
*SystemInsightsApi* | [**SysteminsightsListUsers**](docs/SystemInsightsApi.md#systeminsightslistusers) | **Get** /systeminsights/users | List System Insights Users
*SystemInsightsApi* | [**SysteminsightsListWifiNetworks**](docs/SystemInsightsApi.md#systeminsightslistwifinetworks) | **Get** /systeminsights/wifi_networks | List System Insights WiFi Networks
*SystemInsightsApi* | [**SysteminsightsListWifiStatus**](docs/SystemInsightsApi.md#systeminsightslistwifistatus) | **Get** /systeminsights/wifi_status | List System Insights WiFi Status
*SystemInsightsApi* | [**SysteminsightsListWindowsSecurityCenter**](docs/SystemInsightsApi.md#systeminsightslistwindowssecuritycenter) | **Get** /systeminsights/windows_security_center | List System Insights Windows Security Center
*SystemInsightsApi* | [**SysteminsightsListWindowsSecurityProducts**](docs/SystemInsightsApi.md#systeminsightslistwindowssecurityproducts) | **Get** /systeminsights/windows_security_products | List System Insights Windows Security Products
*SystemsApi* | [**GraphSystemAssociationsList**](docs/SystemsApi.md#graphsystemassociationslist) | **Get** /systems/{system_id}/associations | List the associations of a System
*SystemsApi* | [**GraphSystemAssociationsPost**](docs/SystemsApi.md#graphsystemassociationspost) | **Post** /systems/{system_id}/associations | Manage associations of a System
*SystemsApi* | [**GraphSystemMemberOf**](docs/SystemsApi.md#graphsystemmemberof) | **Get** /systems/{system_id}/memberof | List the parent Groups of a System
*SystemsApi* | [**GraphSystemTraverseCommand**](docs/SystemsApi.md#graphsystemtraversecommand) | **Get** /systems/{system_id}/commands | List the Commands bound to a System
*SystemsApi* | [**GraphSystemTraversePolicy**](docs/SystemsApi.md#graphsystemtraversepolicy) | **Get** /systems/{system_id}/policies | List the Policies bound to a System
*SystemsApi* | [**GraphSystemTraversePolicyGroup**](docs/SystemsApi.md#graphsystemtraversepolicygroup) | **Get** /systems/{system_id}/policygroups | List the Policy Groups bound to a System
*SystemsApi* | [**GraphSystemTraverseUser**](docs/SystemsApi.md#graphsystemtraverseuser) | **Get** /systems/{system_id}/users | List the Users bound to a System
*SystemsApi* | [**GraphSystemTraverseUserGroup**](docs/SystemsApi.md#graphsystemtraverseusergroup) | **Get** /systems/{system_id}/usergroups | List the User Groups bound to a System
*SystemsApi* | [**SystemsGetFDEKey**](docs/SystemsApi.md#systemsgetfdekey) | **Get** /systems/{system_id}/fdekey | Get System FDE Key
*SystemsApi* | [**SystemsListSoftwareAppsWithStatuses**](docs/SystemsApi.md#systemslistsoftwareappswithstatuses) | **Get** /systems/{system_id}/softwareappstatuses | List the associated Software Application Statuses of a System
*UserGroupAssociationsApi* | [**GraphUserGroupAssociationsList**](docs/UserGroupAssociationsApi.md#graphusergroupassociationslist) | **Get** /usergroups/{group_id}/associations | List the associations of a User Group.
*UserGroupAssociationsApi* | [**GraphUserGroupAssociationsPost**](docs/UserGroupAssociationsApi.md#graphusergroupassociationspost) | **Post** /usergroups/{group_id}/associations | Manage the associations of a User Group
*UserGroupAssociationsApi* | [**GraphUserGroupTraverseActiveDirectory**](docs/UserGroupAssociationsApi.md#graphusergrouptraverseactivedirectory) | **Get** /usergroups/{group_id}/activedirectories | List the Active Directories bound to a User Group
*UserGroupAssociationsApi* | [**GraphUserGroupTraverseApplication**](docs/UserGroupAssociationsApi.md#graphusergrouptraverseapplication) | **Get** /usergroups/{group_id}/applications | List the Applications bound to a User Group
*UserGroupAssociationsApi* | [**GraphUserGroupTraverseDirectory**](docs/UserGroupAssociationsApi.md#graphusergrouptraversedirectory) | **Get** /usergroups/{group_id}/directories | List the Directories bound to a User Group
*UserGroupAssociationsApi* | [**GraphUserGroupTraverseGSuite**](docs/UserGroupAssociationsApi.md#graphusergrouptraversegsuite) | **Get** /usergroups/{group_id}/gsuites | List the G Suite instances bound to a User Group
*UserGroupAssociationsApi* | [**GraphUserGroupTraverseLdapServer**](docs/UserGroupAssociationsApi.md#graphusergrouptraverseldapserver) | **Get** /usergroups/{group_id}/ldapservers | List the LDAP Servers bound to a User Group
*UserGroupAssociationsApi* | [**GraphUserGroupTraverseOffice365**](docs/UserGroupAssociationsApi.md#graphusergrouptraverseoffice365) | **Get** /usergroups/{group_id}/office365s | List the Office 365 instances bound to a User Group
*UserGroupAssociationsApi* | [**GraphUserGroupTraverseRadiusServer**](docs/UserGroupAssociationsApi.md#graphusergrouptraverseradiusserver) | **Get** /usergroups/{group_id}/radiusservers | List the RADIUS Servers bound to a User Group
*UserGroupAssociationsApi* | [**GraphUserGroupTraverseSystem**](docs/UserGroupAssociationsApi.md#graphusergrouptraversesystem) | **Get** /usergroups/{group_id}/systems | List the Systems bound to a User Group
*UserGroupAssociationsApi* | [**GraphUserGroupTraverseSystemGroup**](docs/UserGroupAssociationsApi.md#graphusergrouptraversesystemgroup) | **Get** /usergroups/{group_id}/systemgroups | List the System Groups bound to User Groups
*UserGroupMembersMembershipApi* | [**GraphUserGroupMembersList**](docs/UserGroupMembersMembershipApi.md#graphusergroupmemberslist) | **Get** /usergroups/{group_id}/members | List the members of a User Group
*UserGroupMembersMembershipApi* | [**GraphUserGroupMembersPost**](docs/UserGroupMembersMembershipApi.md#graphusergroupmemberspost) | **Post** /usergroups/{group_id}/members | Manage the members of a User Group
*UserGroupMembersMembershipApi* | [**GraphUserGroupMembership**](docs/UserGroupMembersMembershipApi.md#graphusergroupmembership) | **Get** /usergroups/{group_id}/membership | List the User Group&#x27;s membership
*UserGroupsApi* | [**GraphUserGroupAssociationsList**](docs/UserGroupsApi.md#graphusergroupassociationslist) | **Get** /usergroups/{group_id}/associations | List the associations of a User Group.
*UserGroupsApi* | [**GraphUserGroupAssociationsPost**](docs/UserGroupsApi.md#graphusergroupassociationspost) | **Post** /usergroups/{group_id}/associations | Manage the associations of a User Group
*UserGroupsApi* | [**GraphUserGroupMembersList**](docs/UserGroupsApi.md#graphusergroupmemberslist) | **Get** /usergroups/{group_id}/members | List the members of a User Group
*UserGroupsApi* | [**GraphUserGroupMembersPost**](docs/UserGroupsApi.md#graphusergroupmemberspost) | **Post** /usergroups/{group_id}/members | Manage the members of a User Group
*UserGroupsApi* | [**GraphUserGroupMembership**](docs/UserGroupsApi.md#graphusergroupmembership) | **Get** /usergroups/{group_id}/membership | List the User Group&#x27;s membership
*UserGroupsApi* | [**GraphUserGroupTraverseActiveDirectory**](docs/UserGroupsApi.md#graphusergrouptraverseactivedirectory) | **Get** /usergroups/{group_id}/activedirectories | List the Active Directories bound to a User Group
*UserGroupsApi* | [**GraphUserGroupTraverseApplication**](docs/UserGroupsApi.md#graphusergrouptraverseapplication) | **Get** /usergroups/{group_id}/applications | List the Applications bound to a User Group
*UserGroupsApi* | [**GraphUserGroupTraverseDirectory**](docs/UserGroupsApi.md#graphusergrouptraversedirectory) | **Get** /usergroups/{group_id}/directories | List the Directories bound to a User Group
*UserGroupsApi* | [**GraphUserGroupTraverseGSuite**](docs/UserGroupsApi.md#graphusergrouptraversegsuite) | **Get** /usergroups/{group_id}/gsuites | List the G Suite instances bound to a User Group
*UserGroupsApi* | [**GraphUserGroupTraverseLdapServer**](docs/UserGroupsApi.md#graphusergrouptraverseldapserver) | **Get** /usergroups/{group_id}/ldapservers | List the LDAP Servers bound to a User Group
*UserGroupsApi* | [**GraphUserGroupTraverseOffice365**](docs/UserGroupsApi.md#graphusergrouptraverseoffice365) | **Get** /usergroups/{group_id}/office365s | List the Office 365 instances bound to a User Group
*UserGroupsApi* | [**GraphUserGroupTraverseRadiusServer**](docs/UserGroupsApi.md#graphusergrouptraverseradiusserver) | **Get** /usergroups/{group_id}/radiusservers | List the RADIUS Servers bound to a User Group
*UserGroupsApi* | [**GraphUserGroupTraverseSystem**](docs/UserGroupsApi.md#graphusergrouptraversesystem) | **Get** /usergroups/{group_id}/systems | List the Systems bound to a User Group
*UserGroupsApi* | [**GraphUserGroupTraverseSystemGroup**](docs/UserGroupsApi.md#graphusergrouptraversesystemgroup) | **Get** /usergroups/{group_id}/systemgroups | List the System Groups bound to User Groups
*UserGroupsApi* | [**GroupsSuggestionsGet**](docs/UserGroupsApi.md#groupssuggestionsget) | **Get** /usergroups/{group_id}/suggestions | List Suggestions for a User Group
*UserGroupsApi* | [**GroupsSuggestionsPost**](docs/UserGroupsApi.md#groupssuggestionspost) | **Post** /usergroups/{group_id}/suggestions | List Suggestions for a User Group
*UserGroupsApi* | [**GroupsUserDelete**](docs/UserGroupsApi.md#groupsuserdelete) | **Delete** /usergroups/{id} | Delete a User Group
*UserGroupsApi* | [**GroupsUserGet**](docs/UserGroupsApi.md#groupsuserget) | **Get** /usergroups/{id} | View an individual User Group details
*UserGroupsApi* | [**GroupsUserList**](docs/UserGroupsApi.md#groupsuserlist) | **Get** /usergroups | List all User Groups
*UserGroupsApi* | [**GroupsUserPost**](docs/UserGroupsApi.md#groupsuserpost) | **Post** /usergroups | Create a new User Group
*UserGroupsApi* | [**GroupsUserPut**](docs/UserGroupsApi.md#groupsuserput) | **Put** /usergroups/{id} | Update a User Group
*UsersApi* | [**GraphUserAssociationsList**](docs/UsersApi.md#graphuserassociationslist) | **Get** /users/{user_id}/associations | List the associations of a User
*UsersApi* | [**GraphUserAssociationsPost**](docs/UsersApi.md#graphuserassociationspost) | **Post** /users/{user_id}/associations | Manage the associations of a User
*UsersApi* | [**GraphUserMemberOf**](docs/UsersApi.md#graphusermemberof) | **Get** /users/{user_id}/memberof | List the parent Groups of a User
*UsersApi* | [**GraphUserTraverseActiveDirectory**](docs/UsersApi.md#graphusertraverseactivedirectory) | **Get** /users/{user_id}/activedirectories | List the Active Directory instances bound to a User
*UsersApi* | [**GraphUserTraverseApplication**](docs/UsersApi.md#graphusertraverseapplication) | **Get** /users/{user_id}/applications | List the Applications bound to a User
*UsersApi* | [**GraphUserTraverseDirectory**](docs/UsersApi.md#graphusertraversedirectory) | **Get** /users/{user_id}/directories | List the Directories bound to a User
*UsersApi* | [**GraphUserTraverseGSuite**](docs/UsersApi.md#graphusertraversegsuite) | **Get** /users/{user_id}/gsuites | List the G Suite instances bound to a User
*UsersApi* | [**GraphUserTraverseLdapServer**](docs/UsersApi.md#graphusertraverseldapserver) | **Get** /users/{user_id}/ldapservers | List the LDAP servers bound to a User
*UsersApi* | [**GraphUserTraverseOffice365**](docs/UsersApi.md#graphusertraverseoffice365) | **Get** /users/{user_id}/office365s | List the Office 365 instances bound to a User
*UsersApi* | [**GraphUserTraverseRadiusServer**](docs/UsersApi.md#graphusertraverseradiusserver) | **Get** /users/{user_id}/radiusservers | List the RADIUS Servers bound to a User
*UsersApi* | [**GraphUserTraverseSystem**](docs/UsersApi.md#graphusertraversesystem) | **Get** /users/{user_id}/systems | List the Systems bound to a User
*UsersApi* | [**GraphUserTraverseSystemGroup**](docs/UsersApi.md#graphusertraversesystemgroup) | **Get** /users/{user_id}/systemgroups | List the System Groups bound to a User
*UsersApi* | [**PushEndpointsDelete**](docs/UsersApi.md#pushendpointsdelete) | **Delete** /users/{user_id}/pushendpoints/{push_endpoint_id} | Delete a Push Endpoint associated with a User
*UsersApi* | [**PushEndpointsGet**](docs/UsersApi.md#pushendpointsget) | **Get** /users/{user_id}/pushendpoints/{push_endpoint_id} | Get a push endpoint associated with a User
*UsersApi* | [**PushEndpointsList**](docs/UsersApi.md#pushendpointslist) | **Get** /users/{user_id}/pushendpoints | List Push Endpoints associated with a User
*UsersApi* | [**PushEndpointsPatch**](docs/UsersApi.md#pushendpointspatch) | **Patch** /users/{user_id}/pushendpoints/{push_endpoint_id} | Update a push endpoint associated with a User
*WorkdayImportApi* | [**WorkdaysAuthorize**](docs/WorkdayImportApi.md#workdaysauthorize) | **Post** /workdays/{workday_id}/auth | Authorize Workday
*WorkdayImportApi* | [**WorkdaysDeauthorize**](docs/WorkdayImportApi.md#workdaysdeauthorize) | **Delete** /workdays/{workday_id}/auth | Deauthorize Workday
*WorkdayImportApi* | [**WorkdaysGet**](docs/WorkdayImportApi.md#workdaysget) | **Get** /workdays/{id} | Get Workday
*WorkdayImportApi* | [**WorkdaysImport**](docs/WorkdayImportApi.md#workdaysimport) | **Post** /workdays/{workday_id}/import | Workday Import
*WorkdayImportApi* | [**WorkdaysImportresults**](docs/WorkdayImportApi.md#workdaysimportresults) | **Get** /workdays/{id}/import/{job_id}/results | List Import Results
*WorkdayImportApi* | [**WorkdaysList**](docs/WorkdayImportApi.md#workdayslist) | **Get** /workdays | List Workdays
*WorkdayImportApi* | [**WorkdaysPost**](docs/WorkdayImportApi.md#workdayspost) | **Post** /workdays | Create new Workday
*WorkdayImportApi* | [**WorkdaysPut**](docs/WorkdayImportApi.md#workdaysput) | **Put** /workdays/{id} | Update Workday
*WorkdayImportApi* | [**WorkdaysWorkers**](docs/WorkdayImportApi.md#workdaysworkers) | **Get** /workdays/{workday_id}/workers | List Workday Workers

## Documentation For Models

 - [ActiveDirectoryAgentGetOutput](docs/ActiveDirectoryAgentGetOutput.md)
 - [ActiveDirectoryAgentInput](docs/ActiveDirectoryAgentInput.md)
 - [ActiveDirectoryAgentListOutput](docs/ActiveDirectoryAgentListOutput.md)
 - [ActiveDirectoryInput](docs/ActiveDirectoryInput.md)
 - [ActiveDirectoryOutput](docs/ActiveDirectoryOutput.md)
 - [Address](docs/Address.md)
 - [Ade](docs/Ade.md)
 - [Ades](docs/Ades.md)
 - [Administrator](docs/Administrator.md)
 - [AdministratorOrganizationLink](docs/AdministratorOrganizationLink.md)
 - [AdministratorOrganizationLinkReq](docs/AdministratorOrganizationLinkReq.md)
 - [AllOfAutotaskTicketingAlertConfigurationListRecordsItems](docs/AllOfAutotaskTicketingAlertConfigurationListRecordsItems.md)
 - [AllOfConnectWiseTicketingAlertConfigurationListRecordsItems](docs/AllOfConnectWiseTicketingAlertConfigurationListRecordsItems.md)
 - [AnyValue](docs/AnyValue.md)
 - [AppleMdm](docs/AppleMdm.md)
 - [AppleMdmDevice](docs/AppleMdmDevice.md)
 - [AppleMdmDeviceInfo](docs/AppleMdmDeviceInfo.md)
 - [AppleMdmDeviceSecurityInfo](docs/AppleMdmDeviceSecurityInfo.md)
 - [AppleMdmPatchInput](docs/AppleMdmPatchInput.md)
 - [ApplicationIdLogoBody](docs/ApplicationIdLogoBody.md)
 - [AuthInfo](docs/AuthInfo.md)
 - [AuthInput](docs/AuthInput.md)
 - [AuthInputObject](docs/AuthInputObject.md)
 - [AuthinputBasic](docs/AuthinputBasic.md)
 - [AuthinputOauth](docs/AuthinputOauth.md)
 - [AuthnPolicy](docs/AuthnPolicy.md)
 - [AuthnPolicyEffect](docs/AuthnPolicyEffect.md)
 - [AuthnPolicyInput](docs/AuthnPolicyInput.md)
 - [AuthnPolicyObligations](docs/AuthnPolicyObligations.md)
 - [AuthnPolicyObligationsMfa](docs/AuthnPolicyObligationsMfa.md)
 - [AuthnPolicyObligationsUserVerification](docs/AuthnPolicyObligationsUserVerification.md)
 - [AuthnPolicyResourceTarget](docs/AuthnPolicyResourceTarget.md)
 - [AuthnPolicyTargets](docs/AuthnPolicyTargets.md)
 - [AuthnPolicyType](docs/AuthnPolicyType.md)
 - [AuthnPolicyUserAttributeFilter](docs/AuthnPolicyUserAttributeFilter.md)
 - [AuthnPolicyUserAttributeTarget](docs/AuthnPolicyUserAttributeTarget.md)
 - [AuthnPolicyUserGroupTarget](docs/AuthnPolicyUserGroupTarget.md)
 - [AuthnPolicyUserTarget](docs/AuthnPolicyUserTarget.md)
 - [AutotaskCompany](docs/AutotaskCompany.md)
 - [AutotaskCompanyResp](docs/AutotaskCompanyResp.md)
 - [AutotaskCompanyTypeResp](docs/AutotaskCompanyTypeResp.md)
 - [AutotaskContract](docs/AutotaskContract.md)
 - [AutotaskContractField](docs/AutotaskContractField.md)
 - [AutotaskContractFieldValues](docs/AutotaskContractFieldValues.md)
 - [AutotaskIntegration](docs/AutotaskIntegration.md)
 - [AutotaskIntegrationPatchReq](docs/AutotaskIntegrationPatchReq.md)
 - [AutotaskIntegrationReq](docs/AutotaskIntegrationReq.md)
 - [AutotaskMappingRequest](docs/AutotaskMappingRequest.md)
 - [AutotaskMappingRequestCompany](docs/AutotaskMappingRequestCompany.md)
 - [AutotaskMappingRequestContract](docs/AutotaskMappingRequestContract.md)
 - [AutotaskMappingRequestData](docs/AutotaskMappingRequestData.md)
 - [AutotaskMappingRequestOrganization](docs/AutotaskMappingRequestOrganization.md)
 - [AutotaskMappingRequestService](docs/AutotaskMappingRequestService.md)
 - [AutotaskMappingResponse](docs/AutotaskMappingResponse.md)
 - [AutotaskMappingResponseCompany](docs/AutotaskMappingResponseCompany.md)
 - [AutotaskMappingResponseContract](docs/AutotaskMappingResponseContract.md)
 - [AutotaskMappingResponseOrganization](docs/AutotaskMappingResponseOrganization.md)
 - [AutotaskMappingResponseService](docs/AutotaskMappingResponseService.md)
 - [AutotaskService](docs/AutotaskService.md)
 - [AutotaskSettings](docs/AutotaskSettings.md)
 - [AutotaskSettingsPatchReq](docs/AutotaskSettingsPatchReq.md)
 - [AutotaskTicketingAlertConfiguration](docs/AutotaskTicketingAlertConfiguration.md)
 - [AutotaskTicketingAlertConfigurationList](docs/AutotaskTicketingAlertConfigurationList.md)
 - [AutotaskTicketingAlertConfigurationOption](docs/AutotaskTicketingAlertConfigurationOption.md)
 - [AutotaskTicketingAlertConfigurationOptionValues](docs/AutotaskTicketingAlertConfigurationOptionValues.md)
 - [AutotaskTicketingAlertConfigurationOptions](docs/AutotaskTicketingAlertConfigurationOptions.md)
 - [AutotaskTicketingAlertConfigurationPriority](docs/AutotaskTicketingAlertConfigurationPriority.md)
 - [AutotaskTicketingAlertConfigurationRequest](docs/AutotaskTicketingAlertConfigurationRequest.md)
 - [AutotaskTicketingAlertConfigurationResource](docs/AutotaskTicketingAlertConfigurationResource.md)
 - [BillingIntegrationCompanyType](docs/BillingIntegrationCompanyType.md)
 - [BulkScheduledStatechangeCreate](docs/BulkScheduledStatechangeCreate.md)
 - [BulkUserCreate](docs/BulkUserCreate.md)
 - [BulkUserUpdate](docs/BulkUserUpdate.md)
 - [CommandResultList](docs/CommandResultList.md)
 - [CommandResultListResults](docs/CommandResultListResults.md)
 - [ConnectWiseMappingRequest](docs/ConnectWiseMappingRequest.md)
 - [ConnectWiseMappingRequestCompany](docs/ConnectWiseMappingRequestCompany.md)
 - [ConnectWiseMappingRequestData](docs/ConnectWiseMappingRequestData.md)
 - [ConnectWiseMappingRequestOrganization](docs/ConnectWiseMappingRequestOrganization.md)
 - [ConnectWiseMappingResponse](docs/ConnectWiseMappingResponse.md)
 - [ConnectWiseMappingResponseAddition](docs/ConnectWiseMappingResponseAddition.md)
 - [ConnectWiseSettings](docs/ConnectWiseSettings.md)
 - [ConnectWiseSettingsPatchReq](docs/ConnectWiseSettingsPatchReq.md)
 - [ConnectWiseTicketingAlertConfiguration](docs/ConnectWiseTicketingAlertConfiguration.md)
 - [ConnectWiseTicketingAlertConfigurationList](docs/ConnectWiseTicketingAlertConfigurationList.md)
 - [ConnectWiseTicketingAlertConfigurationOption](docs/ConnectWiseTicketingAlertConfigurationOption.md)
 - [ConnectWiseTicketingAlertConfigurationOptions](docs/ConnectWiseTicketingAlertConfigurationOptions.md)
 - [ConnectWiseTicketingAlertConfigurationRequest](docs/ConnectWiseTicketingAlertConfigurationRequest.md)
 - [ConnectwiseAddition](docs/ConnectwiseAddition.md)
 - [ConnectwiseAgreement](docs/ConnectwiseAgreement.md)
 - [ConnectwiseCompany](docs/ConnectwiseCompany.md)
 - [ConnectwiseCompanyResp](docs/ConnectwiseCompanyResp.md)
 - [ConnectwiseCompanyTypeResp](docs/ConnectwiseCompanyTypeResp.md)
 - [ConnectwiseIntegration](docs/ConnectwiseIntegration.md)
 - [ConnectwiseIntegrationPatchReq](docs/ConnectwiseIntegrationPatchReq.md)
 - [ConnectwiseIntegrationReq](docs/ConnectwiseIntegrationReq.md)
 - [CustomEmail](docs/CustomEmail.md)
 - [CustomEmailTemplate](docs/CustomEmailTemplate.md)
 - [CustomEmailTemplateField](docs/CustomEmailTemplateField.md)
 - [CustomEmailType](docs/CustomEmailType.md)
 - [Dep](docs/Dep.md)
 - [DepSetupAssistantOption](docs/DepSetupAssistantOption.md)
 - [DepWelcomeScreen](docs/DepWelcomeScreen.md)
 - [DeviceIdEraseBody](docs/DeviceIdEraseBody.md)
 - [DeviceIdLockBody](docs/DeviceIdLockBody.md)
 - [DeviceIdRestartBody](docs/DeviceIdRestartBody.md)
 - [Directory](docs/Directory.md)
 - [DuoAccount](docs/DuoAccount.md)
 - [DuoApplication](docs/DuoApplication.md)
 - [DuoApplicationReq](docs/DuoApplicationReq.md)
 - [DuoApplicationUpdateReq](docs/DuoApplicationUpdateReq.md)
 - [ErrorDetails](docs/ErrorDetails.md)
 - [Feature](docs/Feature.md)
 - [Filter](docs/Filter.md)
 - [FilterQuery](docs/FilterQuery.md)
 - [GSuiteBuiltinTranslation](docs/GSuiteBuiltinTranslation.md)
 - [GSuiteDirectionTranslation](docs/GSuiteDirectionTranslation.md)
 - [GSuiteTranslationRule](docs/GSuiteTranslationRule.md)
 - [GSuiteTranslationRuleRequest](docs/GSuiteTranslationRuleRequest.md)
 - [GraphAttributeLdapGroups](docs/GraphAttributeLdapGroups.md)
 - [GraphAttributePosixGroups](docs/GraphAttributePosixGroups.md)
 - [GraphAttributePosixGroupsPosixGroups](docs/GraphAttributePosixGroupsPosixGroups.md)
 - [GraphAttributeRadius](docs/GraphAttributeRadius.md)
 - [GraphAttributeRadiusRadius](docs/GraphAttributeRadiusRadius.md)
 - [GraphAttributeRadiusRadiusReply](docs/GraphAttributeRadiusRadiusReply.md)
 - [GraphAttributeSambaEnabled](docs/GraphAttributeSambaEnabled.md)
 - [GraphAttributeSudo](docs/GraphAttributeSudo.md)
 - [GraphAttributeSudoSudo](docs/GraphAttributeSudoSudo.md)
 - [GraphAttributes](docs/GraphAttributes.md)
 - [GraphConnection](docs/GraphConnection.md)
 - [GraphObject](docs/GraphObject.md)
 - [GraphObjectWithPaths](docs/GraphObjectWithPaths.md)
 - [GraphOperation](docs/GraphOperation.md)
 - [GraphOperationActiveDirectory](docs/GraphOperationActiveDirectory.md)
 - [GraphOperationApplication](docs/GraphOperationApplication.md)
 - [GraphOperationCommand](docs/GraphOperationCommand.md)
 - [GraphOperationGSuite](docs/GraphOperationGSuite.md)
 - [GraphOperationLdapServer](docs/GraphOperationLdapServer.md)
 - [GraphOperationOffice365](docs/GraphOperationOffice365.md)
 - [GraphOperationPolicy](docs/GraphOperationPolicy.md)
 - [GraphOperationPolicyGroup](docs/GraphOperationPolicyGroup.md)
 - [GraphOperationPolicyGroupMember](docs/GraphOperationPolicyGroupMember.md)
 - [GraphOperationRadiusServer](docs/GraphOperationRadiusServer.md)
 - [GraphOperationSoftwareApp](docs/GraphOperationSoftwareApp.md)
 - [GraphOperationSystem](docs/GraphOperationSystem.md)
 - [GraphOperationSystemGroup](docs/GraphOperationSystemGroup.md)
 - [GraphOperationSystemGroupMember](docs/GraphOperationSystemGroupMember.md)
 - [GraphOperationUser](docs/GraphOperationUser.md)
 - [GraphOperationUserGroup](docs/GraphOperationUserGroup.md)
 - [GraphOperationUserGroupMember](docs/GraphOperationUserGroupMember.md)
 - [GraphType](docs/GraphType.md)
 - [Group](docs/Group.md)
 - [GroupAttributesUserGroup](docs/GroupAttributesUserGroup.md)
 - [GroupIdSuggestionsBody](docs/GroupIdSuggestionsBody.md)
 - [GroupType](docs/GroupType.md)
 - [GsuiteOutput](docs/GsuiteOutput.md)
 - [GsuitePatchInput](docs/GsuitePatchInput.md)
 - [ImportUser](docs/ImportUser.md)
 - [ImportUserAddress](docs/ImportUserAddress.md)
 - [ImportUserPhoneNumber](docs/ImportUserPhoneNumber.md)
 - [ImportUsersResponse](docs/ImportUsersResponse.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse20010](docs/InlineResponse20010.md)
 - [InlineResponse20011](docs/InlineResponse20011.md)
 - [InlineResponse20011Users](docs/InlineResponse20011Users.md)
 - [InlineResponse20012](docs/InlineResponse20012.md)
 - [InlineResponse20013](docs/InlineResponse20013.md)
 - [InlineResponse2002](docs/InlineResponse2002.md)
 - [InlineResponse2002Users](docs/InlineResponse2002Users.md)
 - [InlineResponse2003](docs/InlineResponse2003.md)
 - [InlineResponse2004](docs/InlineResponse2004.md)
 - [InlineResponse2005](docs/InlineResponse2005.md)
 - [InlineResponse2006](docs/InlineResponse2006.md)
 - [InlineResponse2007](docs/InlineResponse2007.md)
 - [InlineResponse2008](docs/InlineResponse2008.md)
 - [InlineResponse2009](docs/InlineResponse2009.md)
 - [InlineResponse201](docs/InlineResponse201.md)
 - [InlineResponse400](docs/InlineResponse400.md)
 - [Integration](docs/Integration.md)
 - [IntegrationSyncError](docs/IntegrationSyncError.md)
 - [IntegrationSyncErrorResp](docs/IntegrationSyncErrorResp.md)
 - [IntegrationType](docs/IntegrationType.md)
 - [IntegrationsResponse](docs/IntegrationsResponse.md)
 - [IpList](docs/IpList.md)
 - [IpListRequest](docs/IpListRequest.md)
 - [JobId](docs/JobId.md)
 - [JobWorkresult](docs/JobWorkresult.md)
 - [LdapGroup](docs/LdapGroup.md)
 - [LdapServerAction](docs/LdapServerAction.md)
 - [LdapServerInput](docs/LdapServerInput.md)
 - [LdapServerOutput](docs/LdapServerOutput.md)
 - [LdapserversIdBody](docs/LdapserversIdBody.md)
 - [MemberSuggestion](docs/MemberSuggestion.md)
 - [MemberSuggestionsPostResult](docs/MemberSuggestionsPostResult.md)
 - [ModelError](docs/ModelError.md)
 - [Office365BuiltinTranslation](docs/Office365BuiltinTranslation.md)
 - [Office365DirectionTranslation](docs/Office365DirectionTranslation.md)
 - [Office365Output](docs/Office365Output.md)
 - [Office365PatchInput](docs/Office365PatchInput.md)
 - [Office365TranslationRule](docs/Office365TranslationRule.md)
 - [Office365TranslationRuleRequest](docs/Office365TranslationRuleRequest.md)
 - [Organization](docs/Organization.md)
 - [OrganizationCase](docs/OrganizationCase.md)
 - [OrganizationCasesResponse](docs/OrganizationCasesResponse.md)
 - [OsRestriction](docs/OsRestriction.md)
 - [OsRestrictionAppleRestrictions](docs/OsRestrictionAppleRestrictions.md)
 - [PhoneNumber](docs/PhoneNumber.md)
 - [Policy](docs/Policy.md)
 - [PolicyGroup](docs/PolicyGroup.md)
 - [PolicyGroupData](docs/PolicyGroupData.md)
 - [PolicyRequest](docs/PolicyRequest.md)
 - [PolicyRequestTemplate](docs/PolicyRequestTemplate.md)
 - [PolicyResult](docs/PolicyResult.md)
 - [PolicyTemplate](docs/PolicyTemplate.md)
 - [PolicyTemplateConfigField](docs/PolicyTemplateConfigField.md)
 - [PolicyTemplateConfigFieldTooltip](docs/PolicyTemplateConfigFieldTooltip.md)
 - [PolicyTemplateConfigFieldTooltipVariables](docs/PolicyTemplateConfigFieldTooltipVariables.md)
 - [PolicyTemplateWithDetails](docs/PolicyTemplateWithDetails.md)
 - [PolicyValue](docs/PolicyValue.md)
 - [PolicyWithDetails](docs/PolicyWithDetails.md)
 - [Provider](docs/Provider.md)
 - [ProviderAdminReq](docs/ProviderAdminReq.md)
 - [ProviderInvoice](docs/ProviderInvoice.md)
 - [ProviderInvoiceResponse](docs/ProviderInvoiceResponse.md)
 - [PushEndpointResponse](docs/PushEndpointResponse.md)
 - [PushEndpointResponseDevice](docs/PushEndpointResponseDevice.md)
 - [PushendpointsPushEndpointIdBody](docs/PushendpointsPushEndpointIdBody.md)
 - [PwmAllUsers](docs/PwmAllUsers.md)
 - [PwmAllUsersGroups](docs/PwmAllUsersGroups.md)
 - [PwmAllUsersResults](docs/PwmAllUsersResults.md)
 - [PwmOverviewAppVersions](docs/PwmOverviewAppVersions.md)
 - [PwmOverviewAppVersionsResults](docs/PwmOverviewAppVersionsResults.md)
 - [PwmOverviewMain](docs/PwmOverviewMain.md)
 - [PwmOverviewMainDevices](docs/PwmOverviewMainDevices.md)
 - [Query](docs/Query.md)
 - [QueuedCommandList](docs/QueuedCommandList.md)
 - [QueuedCommandListResults](docs/QueuedCommandListResults.md)
 - [SambaDomainInput](docs/SambaDomainInput.md)
 - [SambaDomainOutput](docs/SambaDomainOutput.md)
 - [ScheduledUserstateResult](docs/ScheduledUserstateResult.md)
 - [SetupAssistantOption](docs/SetupAssistantOption.md)
 - [SharedFolderAccessLevels](docs/SharedFolderAccessLevels.md)
 - [SharedFolderAccessLevelsResults](docs/SharedFolderAccessLevelsResults.md)
 - [SharedFolderDetails](docs/SharedFolderDetails.md)
 - [SharedFolderUsers](docs/SharedFolderUsers.md)
 - [SharedFolderUsersResults](docs/SharedFolderUsersResults.md)
 - [SharedFoldersList](docs/SharedFoldersList.md)
 - [SharedFoldersListResults](docs/SharedFoldersListResults.md)
 - [SoftwareApp](docs/SoftwareApp.md)
 - [SoftwareAppAppleVpp](docs/SoftwareAppAppleVpp.md)
 - [SoftwareAppReclaimLicenses](docs/SoftwareAppReclaimLicenses.md)
 - [SoftwareAppSettings](docs/SoftwareAppSettings.md)
 - [SoftwareAppStatus](docs/SoftwareAppStatus.md)
 - [SoftwareAppWithStatus](docs/SoftwareAppWithStatus.md)
 - [SoftwareAppsRetryInstallationRequest](docs/SoftwareAppsRetryInstallationRequest.md)
 - [Subscription](docs/Subscription.md)
 - [SuggestionCounts](docs/SuggestionCounts.md)
 - [SystemGroup](docs/SystemGroup.md)
 - [SystemGroupData](docs/SystemGroupData.md)
 - [SystemInsightsAlf](docs/SystemInsightsAlf.md)
 - [SystemInsightsAlfExceptions](docs/SystemInsightsAlfExceptions.md)
 - [SystemInsightsAlfExplicitAuths](docs/SystemInsightsAlfExplicitAuths.md)
 - [SystemInsightsAppcompatShims](docs/SystemInsightsAppcompatShims.md)
 - [SystemInsightsApps](docs/SystemInsightsApps.md)
 - [SystemInsightsAuthorizedKeys](docs/SystemInsightsAuthorizedKeys.md)
 - [SystemInsightsAzureInstanceMetadata](docs/SystemInsightsAzureInstanceMetadata.md)
 - [SystemInsightsAzureInstanceTags](docs/SystemInsightsAzureInstanceTags.md)
 - [SystemInsightsBattery](docs/SystemInsightsBattery.md)
 - [SystemInsightsBitlockerInfo](docs/SystemInsightsBitlockerInfo.md)
 - [SystemInsightsBrowserPlugins](docs/SystemInsightsBrowserPlugins.md)
 - [SystemInsightsCertificates](docs/SystemInsightsCertificates.md)
 - [SystemInsightsChassisInfo](docs/SystemInsightsChassisInfo.md)
 - [SystemInsightsChromeExtensions](docs/SystemInsightsChromeExtensions.md)
 - [SystemInsightsConnectivity](docs/SystemInsightsConnectivity.md)
 - [SystemInsightsCrashes](docs/SystemInsightsCrashes.md)
 - [SystemInsightsCupsDestinations](docs/SystemInsightsCupsDestinations.md)
 - [SystemInsightsDiskEncryption](docs/SystemInsightsDiskEncryption.md)
 - [SystemInsightsDiskInfo](docs/SystemInsightsDiskInfo.md)
 - [SystemInsightsDnsResolvers](docs/SystemInsightsDnsResolvers.md)
 - [SystemInsightsEtcHosts](docs/SystemInsightsEtcHosts.md)
 - [SystemInsightsFirefoxAddons](docs/SystemInsightsFirefoxAddons.md)
 - [SystemInsightsGroups](docs/SystemInsightsGroups.md)
 - [SystemInsightsIeExtensions](docs/SystemInsightsIeExtensions.md)
 - [SystemInsightsInterfaceAddresses](docs/SystemInsightsInterfaceAddresses.md)
 - [SystemInsightsInterfaceDetails](docs/SystemInsightsInterfaceDetails.md)
 - [SystemInsightsKernelInfo](docs/SystemInsightsKernelInfo.md)
 - [SystemInsightsLaunchd](docs/SystemInsightsLaunchd.md)
 - [SystemInsightsLinuxPackages](docs/SystemInsightsLinuxPackages.md)
 - [SystemInsightsLoggedInUsers](docs/SystemInsightsLoggedInUsers.md)
 - [SystemInsightsLogicalDrives](docs/SystemInsightsLogicalDrives.md)
 - [SystemInsightsManagedPolicies](docs/SystemInsightsManagedPolicies.md)
 - [SystemInsightsMounts](docs/SystemInsightsMounts.md)
 - [SystemInsightsOsVersion](docs/SystemInsightsOsVersion.md)
 - [SystemInsightsPatches](docs/SystemInsightsPatches.md)
 - [SystemInsightsPrograms](docs/SystemInsightsPrograms.md)
 - [SystemInsightsPythonPackages](docs/SystemInsightsPythonPackages.md)
 - [SystemInsightsSafariExtensions](docs/SystemInsightsSafariExtensions.md)
 - [SystemInsightsScheduledTasks](docs/SystemInsightsScheduledTasks.md)
 - [SystemInsightsSecureboot](docs/SystemInsightsSecureboot.md)
 - [SystemInsightsServices](docs/SystemInsightsServices.md)
 - [SystemInsightsShadow](docs/SystemInsightsShadow.md)
 - [SystemInsightsSharedFolders](docs/SystemInsightsSharedFolders.md)
 - [SystemInsightsSharedResources](docs/SystemInsightsSharedResources.md)
 - [SystemInsightsSharingPreferences](docs/SystemInsightsSharingPreferences.md)
 - [SystemInsightsSipConfig](docs/SystemInsightsSipConfig.md)
 - [SystemInsightsStartupItems](docs/SystemInsightsStartupItems.md)
 - [SystemInsightsSystemControls](docs/SystemInsightsSystemControls.md)
 - [SystemInsightsSystemInfo](docs/SystemInsightsSystemInfo.md)
 - [SystemInsightsTpmInfo](docs/SystemInsightsTpmInfo.md)
 - [SystemInsightsUptime](docs/SystemInsightsUptime.md)
 - [SystemInsightsUsbDevices](docs/SystemInsightsUsbDevices.md)
 - [SystemInsightsUserGroups](docs/SystemInsightsUserGroups.md)
 - [SystemInsightsUserSshKeys](docs/SystemInsightsUserSshKeys.md)
 - [SystemInsightsUserassist](docs/SystemInsightsUserassist.md)
 - [SystemInsightsUsers](docs/SystemInsightsUsers.md)
 - [SystemInsightsWifiNetworks](docs/SystemInsightsWifiNetworks.md)
 - [SystemInsightsWifiStatus](docs/SystemInsightsWifiStatus.md)
 - [SystemInsightsWindowsSecurityCenter](docs/SystemInsightsWindowsSecurityCenter.md)
 - [SystemInsightsWindowsSecurityProducts](docs/SystemInsightsWindowsSecurityProducts.md)
 - [Systemfdekey](docs/Systemfdekey.md)
 - [TicketingIntegrationAlert](docs/TicketingIntegrationAlert.md)
 - [TicketingIntegrationAlertsResp](docs/TicketingIntegrationAlertsResp.md)
 - [User](docs/User.md)
 - [UserGroup](docs/UserGroup.md)
 - [UserGroupPost](docs/UserGroupPost.md)
 - [UserGroupPut](docs/UserGroupPut.md)
 - [WorkdayFields](docs/WorkdayFields.md)
 - [WorkdayInput](docs/WorkdayInput.md)
 - [WorkdayOutput](docs/WorkdayOutput.md)
 - [WorkdayWorker](docs/WorkdayWorker.md)
 - [WorkdayoutputAuth](docs/WorkdayoutputAuth.md)

## Documentation For Authorization

## x-api-key
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author

support@jumpcloud.com
