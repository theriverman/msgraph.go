// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Contact undocumented
type Contact struct {
	// OutlookItem is the base model of Contact
	OutlookItem
	// ParentFolderID undocumented
	ParentFolderID *string `json:"parentFolderId,omitempty"`
	// Birthday undocumented
	Birthday *time.Time `json:"birthday,omitempty"`
	// FileAs undocumented
	FileAs *string `json:"fileAs,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// GivenName undocumented
	GivenName *string `json:"givenName,omitempty"`
	// Initials undocumented
	Initials *string `json:"initials,omitempty"`
	// MiddleName undocumented
	MiddleName *string `json:"middleName,omitempty"`
	// NickName undocumented
	NickName *string `json:"nickName,omitempty"`
	// Surname undocumented
	Surname *string `json:"surname,omitempty"`
	// Title undocumented
	Title *string `json:"title,omitempty"`
	// YomiGivenName undocumented
	YomiGivenName *string `json:"yomiGivenName,omitempty"`
	// YomiSurname undocumented
	YomiSurname *string `json:"yomiSurname,omitempty"`
	// YomiCompanyName undocumented
	YomiCompanyName *string `json:"yomiCompanyName,omitempty"`
	// Generation undocumented
	Generation *string `json:"generation,omitempty"`
	// EmailAddresses undocumented
	EmailAddresses []EmailAddress `json:"emailAddresses,omitempty"`
	// ImAddresses undocumented
	ImAddresses []string `json:"imAddresses,omitempty"`
	// JobTitle undocumented
	JobTitle *string `json:"jobTitle,omitempty"`
	// CompanyName undocumented
	CompanyName *string `json:"companyName,omitempty"`
	// Department undocumented
	Department *string `json:"department,omitempty"`
	// OfficeLocation undocumented
	OfficeLocation *string `json:"officeLocation,omitempty"`
	// Profession undocumented
	Profession *string `json:"profession,omitempty"`
	// BusinessHomePage undocumented
	BusinessHomePage *string `json:"businessHomePage,omitempty"`
	// AssistantName undocumented
	AssistantName *string `json:"assistantName,omitempty"`
	// Manager undocumented
	Manager *string `json:"manager,omitempty"`
	// HomePhones undocumented
	HomePhones []string `json:"homePhones,omitempty"`
	// MobilePhone undocumented
	MobilePhone *string `json:"mobilePhone,omitempty"`
	// BusinessPhones undocumented
	BusinessPhones []string `json:"businessPhones,omitempty"`
	// HomeAddress undocumented
	HomeAddress *PhysicalAddress `json:"homeAddress,omitempty"`
	// BusinessAddress undocumented
	BusinessAddress *PhysicalAddress `json:"businessAddress,omitempty"`
	// OtherAddress undocumented
	OtherAddress *PhysicalAddress `json:"otherAddress,omitempty"`
	// SpouseName undocumented
	SpouseName *string `json:"spouseName,omitempty"`
	// PersonalNotes undocumented
	PersonalNotes *string `json:"personalNotes,omitempty"`
	// Children undocumented
	Children []string `json:"children,omitempty"`
	// SingleValueExtendedProperties undocumented
	SingleValueExtendedProperties []SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
	// MultiValueExtendedProperties undocumented
	MultiValueExtendedProperties []MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`
	// Photo undocumented
	Photo *ProfilePhoto `json:"photo,omitempty"`
	// Extensions undocumented
	Extensions []Extension `json:"extensions,omitempty"`
}

// ContactFolder undocumented
type ContactFolder struct {
	// Entity is the base model of ContactFolder
	Entity
	// ParentFolderID undocumented
	ParentFolderID *string `json:"parentFolderId,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// SingleValueExtendedProperties undocumented
	SingleValueExtendedProperties []SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
	// MultiValueExtendedProperties undocumented
	MultiValueExtendedProperties []MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`
	// Contacts undocumented
	Contacts []Contact `json:"contacts,omitempty"`
	// ChildFolders undocumented
	ChildFolders []ContactFolder `json:"childFolders,omitempty"`
}
