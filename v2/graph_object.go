/*
 * JumpCloud APIs
 *
 *  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v2

type GraphObject struct {

	// The ObjectID of the graph object.
	Id string `json:"id"`

	// The type of graph object.
	Type_ string `json:"type"`
}