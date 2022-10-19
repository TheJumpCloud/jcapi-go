# Go API client for jcapiv1

# Overview  JumpCloud's V1 API. This set of endpoints allows JumpCloud customers to manage commands, systems, and system users.  # API Key  ## Access Your API Key  To locate your API Key:  1. Log into the [JumpCloud Admin Console](https://console.jumpcloud.com/). 2. Go to the username drop down located in the top-right of the Console. 3. Retrieve your API key from API Settings.  ## API Key Considerations  This API key is associated to the currently logged in administrator. Other admins will have different API keys.  **WARNING** Please keep this API key secret, as it grants full access to any data accessible via your JumpCloud console account.  You can also reset your API key in the same location in the JumpCloud Admin Console.  ## Recycling or Resetting Your API Key  In order to revoke access with the current API key, simply reset your API key. This will render all calls using the previous API key inaccessible.  Your API key will be passed in as a header with the header name \"x-api-key\".  ```bash curl -H \"x-api-key: [YOUR_API_KEY_HERE]\" \"https://console.jumpcloud.com/api/systemusers\" ```  # System Context  * [Introduction](#introduction) * [Supported endpoints](#supported-endpoints) * [Response codes](#response-codes) * [Authentication](#authentication) * [Additional examples](#additional-examples) * [Third party](#third-party)  ## Introduction  JumpCloud System Context Authorization is an alternative way to authenticate with a subset of JumpCloud's REST APIs. Using this method, a system can manage its information and resource associations, allowing modern auto provisioning environments to scale as needed.  **Notes:**   * The following documentation applies to Linux Operating Systems only.  * Systems that have been automatically enrolled using Apple's Device Enrollment Program (DEP) or systems enrolled using the User Portal install are not eligible to use the System Context API to prevent unauthorized access to system groups and resources. If a script that utilizes the System Context API is invoked on a system enrolled in this way, it will display an error.  ## Supported Endpoints  JumpCloud System Context Authorization can be used in conjunction with Systems endpoints found in the V1 API and certain System Group endpoints found in the v2 API.  * A system may fetch, alter, and delete metadata about itself, including manipulating a system's Group and Systemuser associations,   * `/api/systems/{system_id}` | [`GET`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_get) [`PUT`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_put) * A system may delete itself from your JumpCloud organization   * `/api/systems/{system_id}` | [`DELETE`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_delete) * A system may fetch its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/memberof` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembership)   * `/api/v2/systems/{system_id}/associations` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsList)   * `/api/v2/systems/{system_id}/users` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemTraverseUser) * A system may alter its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/associations` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsPost) * A system may alter its System Group associations   * `/api/v2/systemgroups/{group_id}/members` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembersPost)     * _NOTE_ If a system attempts to alter the system group membership of a different system the request will be rejected  ## Response Codes  If endpoints other than those described above are called using the System Context API, the server will return a `401` response.  ## Authentication  To allow for secure access to our APIs, you must authenticate each API request. JumpCloud System Context Authorization uses [HTTP Signatures](https://tools.ietf.org/html/draft-cavage-http-signatures-00) to authenticate API requests. The HTTP Signatures sent with each request are similar to the signatures used by the Amazon Web Services REST API. To help with the request-signing process, we have provided an [example bash script](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh). This example API request simply requests the entire system record. You must be root, or have permissions to access the contents of the `/opt/jc` directory to generate a signature.  Here is a breakdown of the example script with explanations.  First, the script extracts the systemKey from the JSON formatted `/opt/jc/jcagent.conf` file.  ```bash #!/bin/bash conf=\"`cat /opt/jc/jcagent.conf`\" regex=\"systemKey\\\":\\\"(\\w+)\\\"\"  if [[ $conf =~ $regex ]] ; then   systemKey=\"${BASH_REMATCH[1]}\" fi ```  Then, the script retrieves the current date in the correct format.  ```bash now=`date -u \"+%a, %d %h %Y %H:%M:%S GMT\"`; ```  Next, we build a signing string to demonstrate the expected signature format. The signed string must consist of the [request-line](https://tools.ietf.org/html/rfc2616#page-35) and the date header, separated by a newline character.  ```bash signstr=\"GET /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" ```  The next step is to calculate and apply the signature. This is a two-step process:  1. Create a signature from the signing string using the JumpCloud Agent private key: ``printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key`` 2. Then Base64-encode the signature string and trim off the newline characters: ``| openssl enc -e -a | tr -d '\\n'``  The combined steps above result in:  ```bash signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ; ```  Finally, we make sure the API call sending the signature has the same Authorization and Date header values, HTTP method, and URL that were used in the signing string.  ```bash curl -iq \\   -H \"Accept: application/json\" \\   -H \"Content-Type: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Input Data  All PUT and POST methods should use the HTTP Content-Type header with a value of 'application/json'. PUT methods are used for updating a record. POST methods are used to create a record.  The following example demonstrates how to update the `displayName` of the system.  ```bash signstr=\"PUT /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ;  curl -iq \\   -d \"{\\\"displayName\\\" : \\\"updated-system-name-1\\\"}\" \\   -X \"PUT\" \\   -H \"Content-Type: application/json\" \\   -H \"Accept: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Output Data  All results will be formatted as JSON.  Here is an abbreviated example of response output:  ```json {   \"_id\": \"525ee96f52e144993e000015\",   \"agentServer\": \"lappy386\",   \"agentVersion\": \"0.9.42\",   \"arch\": \"x86_64\",   \"connectionKey\": \"127.0.0.1_51812\",   \"displayName\": \"ubuntu-1204\",   \"firstContact\": \"2013-10-16T19:30:55.611Z\",   \"hostname\": \"ubuntu-1204\"   ... ```  ## Additional Examples  ### Signing Authentication Example  This example demonstrates how to make an authenticated request to fetch the JumpCloud record for this system.  [SigningExample.sh](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh)  ### Shutdown Hook  This example demonstrates how to make an authenticated request on system shutdown. Using an init.d script registered at run level 0, you can call the System Context API as the system is shutting down.  [Instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) is an example of an init.d script that only runs at system shutdown.  After customizing the [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) script, you should install it on the system(s) running the JumpCloud agent.  1. Copy the modified [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) to `/etc/init.d/instance-shutdown`. 2. On Ubuntu systems, run `update-rc.d instance-shutdown defaults`. On RedHat/CentOS systems, run `chkconfig --add instance-shutdown`.  ## Third Party  ### Chef Cookbooks  [https://github.com/nshenry03/jumpcloud](https://github.com/nshenry03/jumpcloud)  [https://github.com/cjs226/jumpcloud](https://github.com/cjs226/jumpcloud)  # Multi-Tenant Portal Headers  Multi-Tenant Organization API Headers are available for JumpCloud Admins to use when making API requests from Organizations that have multiple managed organizations.  The `x-org-id` is a required header for all multi-tenant admins when making API requests to JumpCloud. This header will define to which organization you would like to make the request.  **NOTE** Single Tenant Admins do not need to provide this header when making an API request.  ## Header Value  `x-org-id`  ## API Response Codes  * `400` Malformed ID. * `400` x-org-id and Organization path ID do not match. * `401` ID not included for multi-tenant admin * `403` ID included on unsupported route. * `404` Organization ID Not Found.  ```bash curl -X GET https://console.jumpcloud.com/api/v2/directories \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -H 'x-org-id: {ORG_ID}'  ```  ## To Obtain an Individual Organization ID via the UI  As a prerequisite, your Primary Organization will need to be setup for Multi-Tenancy. This provides access to the Multi-Tenant Organization Admin Portal.  1. Log into JumpCloud [Admin Console](https://console.jumpcloud.com). If you are a multi-tenant Admin, you will automatically be routed to the Multi-Tenant Admin Portal. 2. From the Multi-Tenant Portal's primary navigation bar, select the Organization you'd like to access. 3. You will automatically be routed to that Organization's Admin Console. 4. Go to Settings in the sub-tenant's primary navigation. 5. You can obtain your Organization ID below your Organization's Contact Information on the Settings page.  ## To Obtain All Organization IDs via the API  * You can make an API request to this endpoint using the API key of your Primary Organization.  `https://console.jumpcloud.com/api/organizations/` This will return all your managed organizations.  ```bash curl -X GET \\   https://console.jumpcloud.com/api/organizations/ \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```  # SDKs  You can find language specific SDKs that can help you kickstart your Integration with JumpCloud in the following GitHub repositories:  * [Python](https://github.com/TheJumpCloud/jcapi-python) * [Go](https://github.com/TheJumpCloud/jcapi-go) * [Ruby](https://github.com/TheJumpCloud/jcapi-ruby) * [Java](https://github.com/TheJumpCloud/jcapi-java) 

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 5.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen
For more information, please visit [https://support.jumpcloud.com/support/s/](https://support.jumpcloud.com/support/s/)

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./jcapiv1"
```

## Documentation for API Endpoints

All URIs are relative to *https://console.jumpcloud.com/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApplicationTemplatesApi* | [**ApplicationTemplatesGet**](docs/ApplicationTemplatesApi.md#applicationtemplatesget) | **Get** /application-templates/{id} | Get an Application Template
*ApplicationTemplatesApi* | [**ApplicationTemplatesList**](docs/ApplicationTemplatesApi.md#applicationtemplateslist) | **Get** /application-templates | List Application Templates
*ApplicationsApi* | [**ApplicationsDelete**](docs/ApplicationsApi.md#applicationsdelete) | **Delete** /applications/{id} | Delete an Application
*ApplicationsApi* | [**ApplicationsGet**](docs/ApplicationsApi.md#applicationsget) | **Get** /applications/{id} | Get an Application
*ApplicationsApi* | [**ApplicationsList**](docs/ApplicationsApi.md#applicationslist) | **Get** /applications | Applications
*ApplicationsApi* | [**ApplicationsPost**](docs/ApplicationsApi.md#applicationspost) | **Post** /applications | Create an Application
*ApplicationsApi* | [**ApplicationsPut**](docs/ApplicationsApi.md#applicationsput) | **Put** /applications/{id} | Update an Application
*CommandResultsApi* | [**CommandResultsDelete**](docs/CommandResultsApi.md#commandresultsdelete) | **Delete** /commandresults/{id} | Delete a Command result
*CommandResultsApi* | [**CommandResultsGet**](docs/CommandResultsApi.md#commandresultsget) | **Get** /commandresults/{id} | List an individual Command result
*CommandResultsApi* | [**CommandResultsList**](docs/CommandResultsApi.md#commandresultslist) | **Get** /commandresults | List all Command Results
*CommandTriggersApi* | [**CommandTriggerWebhookPost**](docs/CommandTriggersApi.md#commandtriggerwebhookpost) | **Post** /command/trigger/{triggername} | Launch a command via a Trigger
*CommandsApi* | [**CommandFileGet**](docs/CommandsApi.md#commandfileget) | **Get** /files/command/{id} | Get a Command File
*CommandsApi* | [**CommandsDelete**](docs/CommandsApi.md#commandsdelete) | **Delete** /commands/{id} | Delete a Command
*CommandsApi* | [**CommandsGet**](docs/CommandsApi.md#commandsget) | **Get** /commands/{id} | List an individual Command
*CommandsApi* | [**CommandsGetResults**](docs/CommandsApi.md#commandsgetresults) | **Get** /commands/{id}/results | Get results for a specific command
*CommandsApi* | [**CommandsList**](docs/CommandsApi.md#commandslist) | **Get** /commands | List All Commands
*CommandsApi* | [**CommandsPost**](docs/CommandsApi.md#commandspost) | **Post** /commands | Create A Command
*CommandsApi* | [**CommandsPut**](docs/CommandsApi.md#commandsput) | **Put** /commands/{id} | Update a Command
*ManagedServiceProviderApi* | [**AdminTotpresetBegin**](docs/ManagedServiceProviderApi.md#admintotpresetbegin) | **Post** /users/resettotp/{id} | Administrator TOTP Reset Initiation
*ManagedServiceProviderApi* | [**OrganizationList**](docs/ManagedServiceProviderApi.md#organizationlist) | **Get** /organizations | Get Organization Details
*ManagedServiceProviderApi* | [**UsersPut**](docs/ManagedServiceProviderApi.md#usersput) | **Put** /users/{id} | Update a user
*ManagedServiceProviderApi* | [**UsersReactivateGet**](docs/ManagedServiceProviderApi.md#usersreactivateget) | **Get** /users/reactivate/{id} | Administrator Password Reset Initiation
*OrganizationsApi* | [**OrganizationList**](docs/OrganizationsApi.md#organizationlist) | **Get** /organizations | Get Organization Details
*OrganizationsApi* | [**OrganizationPut**](docs/OrganizationsApi.md#organizationput) | **Put** /organizations/{id} | Update an Organization
*OrganizationsApi* | [**OrganizationsGet**](docs/OrganizationsApi.md#organizationsget) | **Get** /organizations/{id} | Get an Organization
*RadiusServersApi* | [**RadiusServersDelete**](docs/RadiusServersApi.md#radiusserversdelete) | **Delete** /radiusservers/{id} | Delete Radius Server
*RadiusServersApi* | [**RadiusServersGet**](docs/RadiusServersApi.md#radiusserversget) | **Get** /radiusservers/{id} | Get Radius Server
*RadiusServersApi* | [**RadiusServersList**](docs/RadiusServersApi.md#radiusserverslist) | **Get** /radiusservers | List Radius Servers
*RadiusServersApi* | [**RadiusServersPost**](docs/RadiusServersApi.md#radiusserverspost) | **Post** /radiusservers | Create a Radius Server
*RadiusServersApi* | [**RadiusServersPut**](docs/RadiusServersApi.md#radiusserversput) | **Put** /radiusservers/{id} | Update Radius Servers
*SearchApi* | [**SearchCommandresultsPost**](docs/SearchApi.md#searchcommandresultspost) | **Post** /search/commandresults | Search Commands Results
*SearchApi* | [**SearchCommandsPost**](docs/SearchApi.md#searchcommandspost) | **Post** /search/commands | Search Commands
*SearchApi* | [**SearchOrganizationsPost**](docs/SearchApi.md#searchorganizationspost) | **Post** /search/organizations | Search Organizations
*SearchApi* | [**SearchSystemsPost**](docs/SearchApi.md#searchsystemspost) | **Post** /search/systems | Search Systems
*SearchApi* | [**SearchSystemusersPost**](docs/SearchApi.md#searchsystemuserspost) | **Post** /search/systemusers | Search System Users
*SystemsApi* | [**SystemsCommandBuiltinErase**](docs/SystemsApi.md#systemscommandbuiltinerase) | **Post** /systems/{system_id}/command/builtin/erase | Erase a System
*SystemsApi* | [**SystemsCommandBuiltinLock**](docs/SystemsApi.md#systemscommandbuiltinlock) | **Post** /systems/{system_id}/command/builtin/lock | Lock a System
*SystemsApi* | [**SystemsCommandBuiltinRestart**](docs/SystemsApi.md#systemscommandbuiltinrestart) | **Post** /systems/{system_id}/command/builtin/restart | Restart a System
*SystemsApi* | [**SystemsCommandBuiltinShutdown**](docs/SystemsApi.md#systemscommandbuiltinshutdown) | **Post** /systems/{system_id}/command/builtin/shutdown | Shutdown a System
*SystemsApi* | [**SystemsDelete**](docs/SystemsApi.md#systemsdelete) | **Delete** /systems/{id} | Delete a System
*SystemsApi* | [**SystemsGet**](docs/SystemsApi.md#systemsget) | **Get** /systems/{id} | List an individual system
*SystemsApi* | [**SystemsList**](docs/SystemsApi.md#systemslist) | **Get** /systems | List All Systems
*SystemsApi* | [**SystemsPut**](docs/SystemsApi.md#systemsput) | **Put** /systems/{id} | Update a system
*SystemusersApi* | [**SshkeyDelete**](docs/SystemusersApi.md#sshkeydelete) | **Delete** /systemusers/{systemuser_id}/sshkeys/{id} | Delete a system user&#x27;s Public SSH Keys
*SystemusersApi* | [**SshkeyList**](docs/SystemusersApi.md#sshkeylist) | **Get** /systemusers/{id}/sshkeys | List a system user&#x27;s public SSH keys
*SystemusersApi* | [**SshkeyPost**](docs/SystemusersApi.md#sshkeypost) | **Post** /systemusers/{id}/sshkeys | Create a system user&#x27;s Public SSH Key
*SystemusersApi* | [**SystemusersDelete**](docs/SystemusersApi.md#systemusersdelete) | **Delete** /systemusers/{id} | Delete a system user
*SystemusersApi* | [**SystemusersExpire**](docs/SystemusersApi.md#systemusersexpire) | **Post** /systemusers/{id}/expire | Expire a system user&#x27;s password
*SystemusersApi* | [**SystemusersGet**](docs/SystemusersApi.md#systemusersget) | **Get** /systemusers/{id} | List a system user
*SystemusersApi* | [**SystemusersList**](docs/SystemusersApi.md#systemuserslist) | **Get** /systemusers | List all system users
*SystemusersApi* | [**SystemusersMfasync**](docs/SystemusersApi.md#systemusersmfasync) | **Post** /systemusers/{id}/mfasync | Sync a systemuser&#x27;s mfa enrollment status
*SystemusersApi* | [**SystemusersPost**](docs/SystemusersApi.md#systemuserspost) | **Post** /systemusers | Create a system user
*SystemusersApi* | [**SystemusersPut**](docs/SystemusersApi.md#systemusersput) | **Put** /systemusers/{id} | Update a system user
*SystemusersApi* | [**SystemusersResetmfa**](docs/SystemusersApi.md#systemusersresetmfa) | **Post** /systemusers/{id}/resetmfa | Reset a system user&#x27;s MFA token
*SystemusersApi* | [**SystemusersStateActivate**](docs/SystemusersApi.md#systemusersstateactivate) | **Post** /systemusers/{id}/state/activate | Activate System User
*SystemusersApi* | [**SystemusersUnlock**](docs/SystemusersApi.md#systemusersunlock) | **Post** /systemusers/{id}/unlock | Unlock a system user
*UsersApi* | [**AdminTotpresetBegin**](docs/UsersApi.md#admintotpresetbegin) | **Post** /users/resettotp/{id} | Administrator TOTP Reset Initiation
*UsersApi* | [**UsersPut**](docs/UsersApi.md#usersput) | **Put** /users/{id} | Update a user
*UsersApi* | [**UsersReactivateGet**](docs/UsersApi.md#usersreactivateget) | **Get** /users/reactivate/{id} | Administrator Password Reset Initiation

## Documentation For Models

 - [Application](docs/Application.md)
 - [ApplicationConfig](docs/ApplicationConfig.md)
 - [ApplicationConfigAcsUrl](docs/ApplicationConfigAcsUrl.md)
 - [ApplicationConfigAcsUrlTooltip](docs/ApplicationConfigAcsUrlTooltip.md)
 - [ApplicationConfigAcsUrlTooltipVariables](docs/ApplicationConfigAcsUrlTooltipVariables.md)
 - [ApplicationConfigConstantAttributes](docs/ApplicationConfigConstantAttributes.md)
 - [ApplicationConfigConstantAttributesValue](docs/ApplicationConfigConstantAttributesValue.md)
 - [ApplicationConfigDatabaseAttributes](docs/ApplicationConfigDatabaseAttributes.md)
 - [ApplicationLogo](docs/ApplicationLogo.md)
 - [Applicationslist](docs/Applicationslist.md)
 - [Applicationtemplate](docs/Applicationtemplate.md)
 - [ApplicationtemplateJit](docs/ApplicationtemplateJit.md)
 - [ApplicationtemplateLogo](docs/ApplicationtemplateLogo.md)
 - [ApplicationtemplateOidc](docs/ApplicationtemplateOidc.md)
 - [ApplicationtemplateProvision](docs/ApplicationtemplateProvision.md)
 - [Applicationtemplateslist](docs/Applicationtemplateslist.md)
 - [Command](docs/Command.md)
 - [Commandfilereturn](docs/Commandfilereturn.md)
 - [CommandfilereturnResults](docs/CommandfilereturnResults.md)
 - [Commandresult](docs/Commandresult.md)
 - [CommandresultResponse](docs/CommandresultResponse.md)
 - [CommandresultResponseData](docs/CommandresultResponseData.md)
 - [Commandresultslist](docs/Commandresultslist.md)
 - [CommandresultslistResults](docs/CommandresultslistResults.md)
 - [Commandslist](docs/Commandslist.md)
 - [CommandslistResults](docs/CommandslistResults.md)
 - [ErrorDetails](docs/ErrorDetails.md)
 - [Fde](docs/Fde.md)
 - [IdResetmfaBody](docs/IdResetmfaBody.md)
 - [Mfa](docs/Mfa.md)
 - [MfaEnrollment](docs/MfaEnrollment.md)
 - [MfaEnrollmentStatus](docs/MfaEnrollmentStatus.md)
 - [ModelError](docs/ModelError.md)
 - [Organization](docs/Organization.md)
 - [Organizationentitlement](docs/Organizationentitlement.md)
 - [OrganizationentitlementEntitlementProducts](docs/OrganizationentitlementEntitlementProducts.md)
 - [OrganizationsIdBody](docs/OrganizationsIdBody.md)
 - [Organizationsettings](docs/Organizationsettings.md)
 - [OrganizationsettingsDisplayPreferences](docs/OrganizationsettingsDisplayPreferences.md)
 - [OrganizationsettingsDisplayPreferencesOrgInsights](docs/OrganizationsettingsDisplayPreferencesOrgInsights.md)
 - [OrganizationsettingsDisplayPreferencesOrgInsightsApplicationsUsage](docs/OrganizationsettingsDisplayPreferencesOrgInsightsApplicationsUsage.md)
 - [OrganizationsettingsDisplayPreferencesOrgInsightsConsoleStats](docs/OrganizationsettingsDisplayPreferencesOrgInsightsConsoleStats.md)
 - [OrganizationsettingsDisplayPreferencesOrgInsightsDeviceNotifications](docs/OrganizationsettingsDisplayPreferencesOrgInsightsDeviceNotifications.md)
 - [OrganizationsettingsDisplayPreferencesOrgInsightsUserNotifications](docs/OrganizationsettingsDisplayPreferencesOrgInsightsUserNotifications.md)
 - [OrganizationsettingsFeatures](docs/OrganizationsettingsFeatures.md)
 - [OrganizationsettingsFeaturesDirectoryInsights](docs/OrganizationsettingsFeaturesDirectoryInsights.md)
 - [OrganizationsettingsFeaturesDirectoryInsightsPremium](docs/OrganizationsettingsFeaturesDirectoryInsightsPremium.md)
 - [OrganizationsettingsFeaturesSystemInsights](docs/OrganizationsettingsFeaturesSystemInsights.md)
 - [OrganizationsettingsNewSystemUserStateDefaults](docs/OrganizationsettingsNewSystemUserStateDefaults.md)
 - [OrganizationsettingsPasswordPolicy](docs/OrganizationsettingsPasswordPolicy.md)
 - [OrganizationsettingsUserPortal](docs/OrganizationsettingsUserPortal.md)
 - [Organizationsettingsput](docs/Organizationsettingsput.md)
 - [OrganizationsettingsputNewSystemUserStateDefaults](docs/OrganizationsettingsputNewSystemUserStateDefaults.md)
 - [OrganizationsettingsputPasswordPolicy](docs/OrganizationsettingsputPasswordPolicy.md)
 - [Organizationslist](docs/Organizationslist.md)
 - [OrganizationslistResults](docs/OrganizationslistResults.md)
 - [Radiusserver](docs/Radiusserver.md)
 - [Radiusserverpost](docs/Radiusserverpost.md)
 - [Radiusserverput](docs/Radiusserverput.md)
 - [RadiusserversIdBody](docs/RadiusserversIdBody.md)
 - [Radiusserverslist](docs/Radiusserverslist.md)
 - [Search](docs/Search.md)
 - [Sshkeylist](docs/Sshkeylist.md)
 - [Sshkeypost](docs/Sshkeypost.md)
 - [Sso](docs/Sso.md)
 - [StateActivateBody](docs/StateActivateBody.md)
 - [System](docs/System.md)
 - [SystemBuiltInCommands](docs/SystemBuiltInCommands.md)
 - [SystemDomainInfo](docs/SystemDomainInfo.md)
 - [SystemMdm](docs/SystemMdm.md)
 - [SystemMdmInternal](docs/SystemMdmInternal.md)
 - [SystemNetworkInterfaces](docs/SystemNetworkInterfaces.md)
 - [SystemOsVersionDetail](docs/SystemOsVersionDetail.md)
 - [SystemProvisionMetadata](docs/SystemProvisionMetadata.md)
 - [SystemProvisionMetadataProvisioner](docs/SystemProvisionMetadataProvisioner.md)
 - [SystemServiceAccountState](docs/SystemServiceAccountState.md)
 - [SystemSshdParams](docs/SystemSshdParams.md)
 - [SystemSystemInsights](docs/SystemSystemInsights.md)
 - [SystemUserMetrics](docs/SystemUserMetrics.md)
 - [Systemput](docs/Systemput.md)
 - [SystemputAgentBoundMessages](docs/SystemputAgentBoundMessages.md)
 - [Systemslist](docs/Systemslist.md)
 - [Systemuserput](docs/Systemuserput.md)
 - [SystemuserputAddresses](docs/SystemuserputAddresses.md)
 - [SystemuserputAttributes](docs/SystemuserputAttributes.md)
 - [SystemuserputPhoneNumbers](docs/SystemuserputPhoneNumbers.md)
 - [SystemuserputRelationships](docs/SystemuserputRelationships.md)
 - [Systemuserputpost](docs/Systemuserputpost.md)
 - [SystemuserputpostAddresses](docs/SystemuserputpostAddresses.md)
 - [SystemuserputpostPhoneNumbers](docs/SystemuserputpostPhoneNumbers.md)
 - [SystemuserputpostRecoveryEmail](docs/SystemuserputpostRecoveryEmail.md)
 - [Systemuserreturn](docs/Systemuserreturn.md)
 - [SystemuserreturnAddresses](docs/SystemuserreturnAddresses.md)
 - [SystemuserreturnPhoneNumbers](docs/SystemuserreturnPhoneNumbers.md)
 - [SystemuserreturnRecoveryEmail](docs/SystemuserreturnRecoveryEmail.md)
 - [Systemuserslist](docs/Systemuserslist.md)
 - [Triggerreturn](docs/Triggerreturn.md)
 - [TrustedappConfigGet](docs/TrustedappConfigGet.md)
 - [TrustedappConfigGetTrustedApps](docs/TrustedappConfigGetTrustedApps.md)
 - [TrustedappConfigPut](docs/TrustedappConfigPut.md)
 - [Userput](docs/Userput.md)
 - [Userreturn](docs/Userreturn.md)
 - [UserreturnGrowthData](docs/UserreturnGrowthData.md)

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
