/*
 * JumpCloud APIs
 *
 *  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v2

type SystemGroup struct {

	// ObjectId uniquely identifying a System Group.
	Id string `json:"id,omitempty"`

	// The type of the group; always 'system' for a System Group.
	Type_ string `json:"type,omitempty"`

	// Display name of a System Group.
	Name string `json:"name,omitempty"`
}
