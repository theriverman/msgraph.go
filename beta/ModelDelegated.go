// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// DelegatedPermissionClassification undocumented
type DelegatedPermissionClassification struct {
	// Entity is the base model of DelegatedPermissionClassification
	Entity
	// PermissionID undocumented
	PermissionID *string `json:"permissionId,omitempty"`
	// PermissionName undocumented
	PermissionName *string `json:"permissionName,omitempty"`
	// Classification undocumented
	Classification *PermissionClassificationType `json:"classification,omitempty"`
}
