/*
 * JumpCloud APIs
 *
 *  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v2

type SystemGroupMembersReq struct {

	// How to modify the membership connection.
	Op string `json:"op"`

	// The member type.
	Type_ string `json:"type"`

	// The ObjectID of member being added or removed.
	Id string `json:"id"`
}
