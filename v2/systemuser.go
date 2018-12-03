/*
 * JumpCloud APIs
 *
 *  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v2

type Systemuser struct {

	Email string `json:"email,omitempty"`

	Username string `json:"username,omitempty"`

	AllowPublicKey bool `json:"allow_public_key,omitempty"`

	PublicKey string `json:"public_key,omitempty"`

	SshKeys []Sshkeylist `json:"ssh_keys,omitempty"`

	Sudo bool `json:"sudo,omitempty"`

	EnableManagedUid bool `json:"enable_managed_uid,omitempty"`

	UnixUid int32 `json:"unix_uid,omitempty"`

	UnixGuid int32 `json:"unix_guid,omitempty"`

	Activated bool `json:"activated,omitempty"`

	Tags []string `json:"tags,omitempty"`

	PasswordExpired bool `json:"password_expired,omitempty"`

	AccountLocked bool `json:"account_locked,omitempty"`

	PasswordlessSudo bool `json:"passwordless_sudo,omitempty"`

	ExternallyManaged bool `json:"externally_managed,omitempty"`

	ExternalDn string `json:"external_dn,omitempty"`

	ExternalSourceType string `json:"external_source_type,omitempty"`

	Firstname string `json:"firstname,omitempty"`

	Lastname string `json:"lastname,omitempty"`

	LdapBindingUser bool `json:"ldap_binding_user,omitempty"`

	EnableUserPortalMultifactor bool `json:"enable_user_portal_multifactor,omitempty"`

	AssociatedTagCount int32 `json:"associatedTagCount,omitempty"`

	TotpEnabled bool `json:"totp_enabled,omitempty"`

	PasswordExpirationDate string `json:"password_expiration_date,omitempty"`

	Attributes []interface{} `json:"attributes,omitempty"`

	Created string `json:"created,omitempty"`

	SambaServiceUser bool `json:"samba_service_user,omitempty"`

	PasswordNeverExpires bool `json:"password_never_expires,omitempty"`

	Id string `json:"_id,omitempty"`

	Middlename string `json:"middlename,omitempty"`

	Displayname string `json:"displayname,omitempty"`

	Description string `json:"description,omitempty"`

	Location string `json:"location,omitempty"`

	CostCenter string `json:"costCenter,omitempty"`

	EmployeeType string `json:"employeeType,omitempty"`

	Company string `json:"company,omitempty"`

	// Must be unique per user. 
	EmployeeIdentifier string `json:"employeeIdentifier,omitempty"`

	JobTitle string `json:"jobTitle,omitempty"`

	Department string `json:"department,omitempty"`

	MfaData *SystemuserMfaData `json:"mfaData,omitempty"`
}
