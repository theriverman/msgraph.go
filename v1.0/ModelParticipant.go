// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Participant undocumented
type Participant struct {
	// Entity is the base model of Participant
	Entity
	// Info undocumented
	Info *ParticipantInfo `json:"info,omitempty"`
	// MediaStreams undocumented
	MediaStreams []MediaStream `json:"mediaStreams,omitempty"`
	// IsMuted undocumented
	IsMuted *bool `json:"isMuted,omitempty"`
	// IsInLobby undocumented
	IsInLobby *bool `json:"isInLobby,omitempty"`
}

// ParticipantInfo undocumented
type ParticipantInfo struct {
	// Object is the base model of ParticipantInfo
	Object
	// Identity undocumented
	Identity *IdentitySet `json:"identity,omitempty"`
	// Region undocumented
	Region *string `json:"region,omitempty"`
	// LanguageID undocumented
	LanguageID *string `json:"languageId,omitempty"`
}
