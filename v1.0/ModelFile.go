// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// File undocumented
type File struct {
	// Object is the base model of File
	Object
	// Hashes undocumented
	Hashes *Hashes `json:"hashes,omitempty"`
	// MimeType undocumented
	MimeType *string `json:"mimeType,omitempty"`
	// ProcessingMetadata undocumented
	ProcessingMetadata *bool `json:"processingMetadata,omitempty"`
}

// FileAttachment undocumented
type FileAttachment struct {
	// Attachment is the base model of FileAttachment
	Attachment
	// ContentID undocumented
	ContentID *string `json:"contentId,omitempty"`
	// ContentLocation undocumented
	ContentLocation *string `json:"contentLocation,omitempty"`
	// ContentBytes undocumented
	ContentBytes *Binary `json:"contentBytes,omitempty"`
}

// FileEncryptionInfo undocumented
type FileEncryptionInfo struct {
	// Object is the base model of FileEncryptionInfo
	Object
	// EncryptionKey The key used to encrypt the file content.
	EncryptionKey *Binary `json:"encryptionKey,omitempty"`
	// InitializationVector The initialization vector used for the encryption algorithm.
	InitializationVector *Binary `json:"initializationVector,omitempty"`
	// Mac The hash of the encrypted file content + IV (content hash).
	Mac *Binary `json:"mac,omitempty"`
	// MacKey The key used to get mac.
	MacKey *Binary `json:"macKey,omitempty"`
	// ProfileIdentifier The the profile identifier.
	ProfileIdentifier *string `json:"profileIdentifier,omitempty"`
	// FileDigest The file digest prior to encryption.
	FileDigest *Binary `json:"fileDigest,omitempty"`
	// FileDigestAlgorithm The file digest algorithm.
	FileDigestAlgorithm *string `json:"fileDigestAlgorithm,omitempty"`
}

// FileHash undocumented
type FileHash struct {
	// Object is the base model of FileHash
	Object
	// HashType undocumented
	HashType *FileHashType `json:"hashType,omitempty"`
	// HashValue undocumented
	HashValue *string `json:"hashValue,omitempty"`
}

// FileSecurityState undocumented
type FileSecurityState struct {
	// Object is the base model of FileSecurityState
	Object
	// FileHash undocumented
	FileHash *FileHash `json:"fileHash,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Path undocumented
	Path *string `json:"path,omitempty"`
	// RiskScore undocumented
	RiskScore *string `json:"riskScore,omitempty"`
}

// FileSystemInfo undocumented
type FileSystemInfo struct {
	// Object is the base model of FileSystemInfo
	Object
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// LastAccessedDateTime undocumented
	LastAccessedDateTime *time.Time `json:"lastAccessedDateTime,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}
