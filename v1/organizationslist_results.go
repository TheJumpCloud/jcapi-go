/*
 * JumpCloud APIs
 *
 *  JumpCloud's V1 API. This set of endpoints allows JumpCloud customers to manage commands, systems, & system users.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v1

type OrganizationslistResults struct {

	// the ID of the organization.
	Id string `json:"_id,omitempty"`

	// The name of the organization.
	DisplayName string `json:"displayName,omitempty"`

	// The organization logo image URL. 
	LogoUrl string `json:"logoUrl,omitempty"`
}
