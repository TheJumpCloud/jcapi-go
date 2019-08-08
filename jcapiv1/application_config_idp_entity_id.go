/*
 * JumpCloud APIs
 *
 *  JumpCloud's V1 API. This set of endpoints allows JumpCloud customers to manage commands, systems, & system users.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v1

type ApplicationConfigIdpEntityId struct {

	Label string `json:"label,omitempty"`

	ReadOnly bool `json:"readOnly,omitempty"`

	Tooltip *ApplicationConfigIdpEntityIdTooltip `json:"tooltip,omitempty"`

	Type_ string `json:"type,omitempty"`

	Value string `json:"value,omitempty"`

	Visible bool `json:"visible,omitempty"`

	Required bool `json:"required,omitempty"`

	Position int32 `json:"position,omitempty"`
}
