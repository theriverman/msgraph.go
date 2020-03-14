// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EncryptContent undocumented
type EncryptContent struct {
	// LabelActionBase is the base model of EncryptContent
	LabelActionBase
	// EncryptWith undocumented
	EncryptWith *EncryptWith `json:"encryptWith,omitempty"`
}

// EncryptWithTemplate undocumented
type EncryptWithTemplate struct {
	// EncryptContent is the base model of EncryptWithTemplate
	EncryptContent
	// TemplateID undocumented
	TemplateID *string `json:"templateId,omitempty"`
	// AvailableForEncryption undocumented
	AvailableForEncryption *bool `json:"availableForEncryption,omitempty"`
}

// EncryptWithUserDefinedRights undocumented
type EncryptWithUserDefinedRights struct {
	// EncryptContent is the base model of EncryptWithUserDefinedRights
	EncryptContent
	// DecryptionRightsManagementTemplateID undocumented
	DecryptionRightsManagementTemplateID *string `json:"decryptionRightsManagementTemplateId,omitempty"`
	// AllowMailForwarding undocumented
	AllowMailForwarding *bool `json:"allowMailForwarding,omitempty"`
	// AllowAdHocPermissions undocumented
	AllowAdHocPermissions *bool `json:"allowAdHocPermissions,omitempty"`
}
