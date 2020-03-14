// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Permission undocumented
type Permission struct {
	// Entity is the base model of Permission
	Entity
	// GrantedTo undocumented
	GrantedTo *IdentitySet `json:"grantedTo,omitempty"`
	// InheritedFrom undocumented
	InheritedFrom *ItemReference `json:"inheritedFrom,omitempty"`
	// Invitation undocumented
	Invitation *SharingInvitation `json:"invitation,omitempty"`
	// Link undocumented
	Link *SharingLink `json:"link,omitempty"`
	// Roles undocumented
	Roles []string `json:"roles,omitempty"`
	// ShareID undocumented
	ShareID *string `json:"shareId,omitempty"`
}

// PermissionScope undocumented
type PermissionScope struct {
	// Object is the base model of PermissionScope
	Object
	// AdminConsentDescription undocumented
	AdminConsentDescription *string `json:"adminConsentDescription,omitempty"`
	// AdminConsentDisplayName undocumented
	AdminConsentDisplayName *string `json:"adminConsentDisplayName,omitempty"`
	// ID undocumented
	ID *UUID `json:"id,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Origin undocumented
	Origin *string `json:"origin,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
	// UserConsentDescription undocumented
	UserConsentDescription *string `json:"userConsentDescription,omitempty"`
	// UserConsentDisplayName undocumented
	UserConsentDisplayName *string `json:"userConsentDisplayName,omitempty"`
	// Value undocumented
	Value *string `json:"value,omitempty"`
}
