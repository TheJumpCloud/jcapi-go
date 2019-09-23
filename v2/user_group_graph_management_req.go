/*
 * JumpCloud APIs
 *
 *  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v2

type UserGroupGraphManagementReq struct {

	// The ObjectID of graph object being added or removed as an association.
	Id string `json:"id"`

	// How to modify the graph connection.
	Op string `json:"op"`

	// The graph type
	Type_ string `json:"type"`
}
