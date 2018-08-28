/*
 * JumpCloud APIs
 *
 *  JumpCloud's V1 API. This set of endpoints allows JumpCloud customers to manage commands, systems, & system users.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v1

type Organizationslist struct {

	// The total of organizations. 
	TotalCount int32 `json:"totalCount,omitempty"`

	// The list of organizations.
	Results []OrganizationslistResults `json:"results,omitempty"`
}
