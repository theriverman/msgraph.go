// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeleteAction undocumented
type DeleteAction struct {
	// Object is the base model of DeleteAction
	Object
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// ObjectType undocumented
	ObjectType *string `json:"objectType,omitempty"`
}

// DeleteUserFromSharedAppleDeviceActionResult undocumented
type DeleteUserFromSharedAppleDeviceActionResult struct {
	// DeviceActionResult is the base model of DeleteUserFromSharedAppleDeviceActionResult
	DeviceActionResult
	// UserPrincipalName User principal name of the user to be deleted
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}
