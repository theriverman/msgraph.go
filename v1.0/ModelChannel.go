// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Channel undocumented
type Channel struct {
	// Entity is the base model of Channel
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// Email undocumented
	Email *string `json:"email,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
	// Tabs undocumented
	Tabs []TeamsTab `json:"tabs,omitempty"`
}
