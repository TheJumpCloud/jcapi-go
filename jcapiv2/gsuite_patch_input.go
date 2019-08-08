/*
 * JumpCloud APIs
 *
 *  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v2

type GsuitePatchInput struct {

	UserLockoutAction string `json:"userLockoutAction,omitempty"`

	UserPasswordExpirationAction string `json:"userPasswordExpirationAction,omitempty"`
}
