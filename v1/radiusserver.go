/*
 * JumpCloud APIs
 *
 *  JumpCloud's V1 API. This set of endpoints allows JumpCloud customers to manage commands, systems, & system users.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v1

type Radiusserver struct {

	Id string `json:"_id,omitempty"`

	Mfa string `json:"mfa,omitempty"`

	Name string `json:"name,omitempty"`

	NetworkSourceIp string `json:"networkSourceIp,omitempty"`

	Organization string `json:"organization,omitempty"`

	SharedSecret string `json:"sharedSecret,omitempty"`

	TagNames []string `json:"tagNames,omitempty"`

	Tags []string `json:"tags,omitempty"`

	UserLockoutAction string `json:"userLockoutAction,omitempty"`

	UserPasswordExpirationAction string `json:"userPasswordExpirationAction,omitempty"`
}
