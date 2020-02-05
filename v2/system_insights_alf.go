/*
 * JumpCloud APIs
 *
 *  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v2

type SystemInsightsAlf struct {

	AllowSignedEnabled int32 `json:"allow_signed_enabled	,omitempty"`

	CollectionTime string `json:"collection_time,omitempty"`

	FirewallUnload int32 `json:"firewall_unload,omitempty"`

	GlobalState int32 `json:"global_state,omitempty"`

	LoggingEnabled int32 `json:"logging_enabled,omitempty"`

	LoggingOption int32 `json:"logging_option,omitempty"`

	StealthEnabled int32 `json:"stealth_enabled,omitempty"`

	SystemId string `json:"system_id,omitempty"`

	Version string `json:"version,omitempty"`
}
