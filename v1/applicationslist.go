/*
 * JumpCloud APIs
 *
 *  JumpCloud's V1 API. This set of endpoints allows JumpCloud customers to manage commands, systems, & system users.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v1

type Applicationslist struct {

	Name string `json:"name,omitempty"`

	// The list of applications.
	Results []Application `json:"results,omitempty"`

	// The total number of applications.
	TotalCount int32 `json:"totalCount,omitempty"`
}
