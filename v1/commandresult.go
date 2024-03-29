/*
 * JumpCloud APIs
 *
 *  JumpCloud's V1 API. This set of endpoints allows JumpCloud customers to manage commands, systems, & system users.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v1

type Commandresult struct {

	// The ID of the command.
	Id string `json:"_id,omitempty"`

	// The command that was executed on the system.
	Command string `json:"command,omitempty"`

	// An array of file ids that were included in the command
	Files []string `json:"files,omitempty"`

	// The name of the command.
	Name string `json:"name,omitempty"`

	// The ID of the organization.
	Organization string `json:"organization,omitempty"`

	// The time that the command was sent.
	RequestTime string `json:"requestTime,omitempty"`

	Response *CommandresultResponse `json:"response,omitempty"`

	// The time that the command was completed.
	ResponseTime string `json:"responseTime,omitempty"`

	// If the user had sudo rights
	Sudo bool `json:"sudo,omitempty"`

	// The name of the system the command was executed on.
	System string `json:"system,omitempty"`

	// The id of the system the command was executed on.
	SystemId string `json:"systemId,omitempty"`

	// The user the command ran as.
	User string `json:"user,omitempty"`

	WorkflowId string `json:"workflowId,omitempty"`

	WorkflowInstanceId string `json:"workflowInstanceId,omitempty"`
}
