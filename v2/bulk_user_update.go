/*
 * JumpCloud APIs
 *
 *  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v2

// See [V1 system user update](https://docs.jumpcloud.com/1.0/systemusers/update-a-system-user) for full list of attributes.
type BulkUserUpdate struct {

	// Object ID of the systemuser being updated
	Id string `json:"id,omitempty"`

	// Map of additional attributes.
	Attributes []interface{} `json:"attributes,omitempty"`
}