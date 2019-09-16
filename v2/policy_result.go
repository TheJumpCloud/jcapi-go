/*
 * JumpCloud APIs
 *
 *  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v2

import (
	"time"
)

type PolicyResult struct {

	// Details pertaining to the policy result.
	Detail string `json:"detail,omitempty"`

	// The end of the policy application.
	EndedAt time.Time `json:"endedAt,omitempty"`

	// The 32-bit unsigned exit status from the applying the policy.
	ExitStatus int32 `json:"exitStatus,omitempty"`

	// ObjectId uniquely identifying a Policy Result.
	Id string `json:"id,omitempty"`

	// ObjectId uniquely identifying the parent Policy.
	PolicyID string `json:"policyID,omitempty"`

	// The start of the policy application.
	StartedAt time.Time `json:"startedAt,omitempty"`

	// Enumeration describing the state of the policy. Success, failed, or pending.
	State string `json:"state,omitempty"`

	// The STDERR output from applying the policy.
	StdErr string `json:"stdErr,omitempty"`

	// The STDOUT output from applying the policy.
	StdOut string `json:"stdOut,omitempty"`

	// True if the policy was successfully applied; false otherwise.
	Success bool `json:"success,omitempty"`

	// ObjectId uniquely identifying the parent System.
	SystemID string `json:"systemID,omitempty"`
}
