// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AppleSubjectNameFormat undocumented
type AppleSubjectNameFormat string

const (
	// AppleSubjectNameFormatVCommonName undocumented
	AppleSubjectNameFormatVCommonName AppleSubjectNameFormat = "commonName"
	// AppleSubjectNameFormatVCommonNameAsEmail undocumented
	AppleSubjectNameFormatVCommonNameAsEmail AppleSubjectNameFormat = "commonNameAsEmail"
	// AppleSubjectNameFormatVCustom undocumented
	AppleSubjectNameFormatVCustom AppleSubjectNameFormat = "custom"
	// AppleSubjectNameFormatVCommonNameIncludingEmail undocumented
	AppleSubjectNameFormatVCommonNameIncludingEmail AppleSubjectNameFormat = "commonNameIncludingEmail"
	// AppleSubjectNameFormatVCommonNameAsIMEI undocumented
	AppleSubjectNameFormatVCommonNameAsIMEI AppleSubjectNameFormat = "commonNameAsIMEI"
	// AppleSubjectNameFormatVCommonNameAsSerialNumber undocumented
	AppleSubjectNameFormatVCommonNameAsSerialNumber AppleSubjectNameFormat = "commonNameAsSerialNumber"
)

var (
	// AppleSubjectNameFormatPCommonName is a pointer to AppleSubjectNameFormatVCommonName
	AppleSubjectNameFormatPCommonName = &_AppleSubjectNameFormatPCommonName
	// AppleSubjectNameFormatPCommonNameAsEmail is a pointer to AppleSubjectNameFormatVCommonNameAsEmail
	AppleSubjectNameFormatPCommonNameAsEmail = &_AppleSubjectNameFormatPCommonNameAsEmail
	// AppleSubjectNameFormatPCustom is a pointer to AppleSubjectNameFormatVCustom
	AppleSubjectNameFormatPCustom = &_AppleSubjectNameFormatPCustom
	// AppleSubjectNameFormatPCommonNameIncludingEmail is a pointer to AppleSubjectNameFormatVCommonNameIncludingEmail
	AppleSubjectNameFormatPCommonNameIncludingEmail = &_AppleSubjectNameFormatPCommonNameIncludingEmail
	// AppleSubjectNameFormatPCommonNameAsIMEI is a pointer to AppleSubjectNameFormatVCommonNameAsIMEI
	AppleSubjectNameFormatPCommonNameAsIMEI = &_AppleSubjectNameFormatPCommonNameAsIMEI
	// AppleSubjectNameFormatPCommonNameAsSerialNumber is a pointer to AppleSubjectNameFormatVCommonNameAsSerialNumber
	AppleSubjectNameFormatPCommonNameAsSerialNumber = &_AppleSubjectNameFormatPCommonNameAsSerialNumber
)

var (
	_AppleSubjectNameFormatPCommonName               = AppleSubjectNameFormatVCommonName
	_AppleSubjectNameFormatPCommonNameAsEmail        = AppleSubjectNameFormatVCommonNameAsEmail
	_AppleSubjectNameFormatPCustom                   = AppleSubjectNameFormatVCustom
	_AppleSubjectNameFormatPCommonNameIncludingEmail = AppleSubjectNameFormatVCommonNameIncludingEmail
	_AppleSubjectNameFormatPCommonNameAsIMEI         = AppleSubjectNameFormatVCommonNameAsIMEI
	_AppleSubjectNameFormatPCommonNameAsSerialNumber = AppleSubjectNameFormatVCommonNameAsSerialNumber
)

// AppleUserInitiatedEnrollmentType undocumented
type AppleUserInitiatedEnrollmentType string

const (
	// AppleUserInitiatedEnrollmentTypeVUnknown undocumented
	AppleUserInitiatedEnrollmentTypeVUnknown AppleUserInitiatedEnrollmentType = "unknown"
	// AppleUserInitiatedEnrollmentTypeVDevice undocumented
	AppleUserInitiatedEnrollmentTypeVDevice AppleUserInitiatedEnrollmentType = "device"
	// AppleUserInitiatedEnrollmentTypeVUser undocumented
	AppleUserInitiatedEnrollmentTypeVUser AppleUserInitiatedEnrollmentType = "user"
)

var (
	// AppleUserInitiatedEnrollmentTypePUnknown is a pointer to AppleUserInitiatedEnrollmentTypeVUnknown
	AppleUserInitiatedEnrollmentTypePUnknown = &_AppleUserInitiatedEnrollmentTypePUnknown
	// AppleUserInitiatedEnrollmentTypePDevice is a pointer to AppleUserInitiatedEnrollmentTypeVDevice
	AppleUserInitiatedEnrollmentTypePDevice = &_AppleUserInitiatedEnrollmentTypePDevice
	// AppleUserInitiatedEnrollmentTypePUser is a pointer to AppleUserInitiatedEnrollmentTypeVUser
	AppleUserInitiatedEnrollmentTypePUser = &_AppleUserInitiatedEnrollmentTypePUser
)

var (
	_AppleUserInitiatedEnrollmentTypePUnknown = AppleUserInitiatedEnrollmentTypeVUnknown
	_AppleUserInitiatedEnrollmentTypePDevice  = AppleUserInitiatedEnrollmentTypeVDevice
	_AppleUserInitiatedEnrollmentTypePUser    = AppleUserInitiatedEnrollmentTypeVUser
)

// AppleVPNConnectionType undocumented
type AppleVPNConnectionType string

const (
	// AppleVPNConnectionTypeVCiscoAnyConnect undocumented
	AppleVPNConnectionTypeVCiscoAnyConnect AppleVPNConnectionType = "ciscoAnyConnect"
	// AppleVPNConnectionTypeVPulseSecure undocumented
	AppleVPNConnectionTypeVPulseSecure AppleVPNConnectionType = "pulseSecure"
	// AppleVPNConnectionTypeVF5EdgeClient undocumented
	AppleVPNConnectionTypeVF5EdgeClient AppleVPNConnectionType = "f5EdgeClient"
	// AppleVPNConnectionTypeVDellSonicWallMobileConnect undocumented
	AppleVPNConnectionTypeVDellSonicWallMobileConnect AppleVPNConnectionType = "dellSonicWallMobileConnect"
	// AppleVPNConnectionTypeVCheckPointCapsuleVPN undocumented
	AppleVPNConnectionTypeVCheckPointCapsuleVPN AppleVPNConnectionType = "checkPointCapsuleVpn"
	// AppleVPNConnectionTypeVCustomVPN undocumented
	AppleVPNConnectionTypeVCustomVPN AppleVPNConnectionType = "customVpn"
	// AppleVPNConnectionTypeVCiscoIPSec undocumented
	AppleVPNConnectionTypeVCiscoIPSec AppleVPNConnectionType = "ciscoIPSec"
	// AppleVPNConnectionTypeVCitrix undocumented
	AppleVPNConnectionTypeVCitrix AppleVPNConnectionType = "citrix"
	// AppleVPNConnectionTypeVCiscoAnyConnectV2 undocumented
	AppleVPNConnectionTypeVCiscoAnyConnectV2 AppleVPNConnectionType = "ciscoAnyConnectV2"
	// AppleVPNConnectionTypeVPaloAltoGlobalProtect undocumented
	AppleVPNConnectionTypeVPaloAltoGlobalProtect AppleVPNConnectionType = "paloAltoGlobalProtect"
	// AppleVPNConnectionTypeVZscalerPrivateAccess undocumented
	AppleVPNConnectionTypeVZscalerPrivateAccess AppleVPNConnectionType = "zscalerPrivateAccess"
	// AppleVPNConnectionTypeVF5Access2018 undocumented
	AppleVPNConnectionTypeVF5Access2018 AppleVPNConnectionType = "f5Access2018"
	// AppleVPNConnectionTypeVCitrixSso undocumented
	AppleVPNConnectionTypeVCitrixSso AppleVPNConnectionType = "citrixSso"
	// AppleVPNConnectionTypeVPaloAltoGlobalProtectV2 undocumented
	AppleVPNConnectionTypeVPaloAltoGlobalProtectV2 AppleVPNConnectionType = "paloAltoGlobalProtectV2"
	// AppleVPNConnectionTypeVIkEv2 undocumented
	AppleVPNConnectionTypeVIkEv2 AppleVPNConnectionType = "ikEv2"
)

var (
	// AppleVPNConnectionTypePCiscoAnyConnect is a pointer to AppleVPNConnectionTypeVCiscoAnyConnect
	AppleVPNConnectionTypePCiscoAnyConnect = &_AppleVPNConnectionTypePCiscoAnyConnect
	// AppleVPNConnectionTypePPulseSecure is a pointer to AppleVPNConnectionTypeVPulseSecure
	AppleVPNConnectionTypePPulseSecure = &_AppleVPNConnectionTypePPulseSecure
	// AppleVPNConnectionTypePF5EdgeClient is a pointer to AppleVPNConnectionTypeVF5EdgeClient
	AppleVPNConnectionTypePF5EdgeClient = &_AppleVPNConnectionTypePF5EdgeClient
	// AppleVPNConnectionTypePDellSonicWallMobileConnect is a pointer to AppleVPNConnectionTypeVDellSonicWallMobileConnect
	AppleVPNConnectionTypePDellSonicWallMobileConnect = &_AppleVPNConnectionTypePDellSonicWallMobileConnect
	// AppleVPNConnectionTypePCheckPointCapsuleVPN is a pointer to AppleVPNConnectionTypeVCheckPointCapsuleVPN
	AppleVPNConnectionTypePCheckPointCapsuleVPN = &_AppleVPNConnectionTypePCheckPointCapsuleVPN
	// AppleVPNConnectionTypePCustomVPN is a pointer to AppleVPNConnectionTypeVCustomVPN
	AppleVPNConnectionTypePCustomVPN = &_AppleVPNConnectionTypePCustomVPN
	// AppleVPNConnectionTypePCiscoIPSec is a pointer to AppleVPNConnectionTypeVCiscoIPSec
	AppleVPNConnectionTypePCiscoIPSec = &_AppleVPNConnectionTypePCiscoIPSec
	// AppleVPNConnectionTypePCitrix is a pointer to AppleVPNConnectionTypeVCitrix
	AppleVPNConnectionTypePCitrix = &_AppleVPNConnectionTypePCitrix
	// AppleVPNConnectionTypePCiscoAnyConnectV2 is a pointer to AppleVPNConnectionTypeVCiscoAnyConnectV2
	AppleVPNConnectionTypePCiscoAnyConnectV2 = &_AppleVPNConnectionTypePCiscoAnyConnectV2
	// AppleVPNConnectionTypePPaloAltoGlobalProtect is a pointer to AppleVPNConnectionTypeVPaloAltoGlobalProtect
	AppleVPNConnectionTypePPaloAltoGlobalProtect = &_AppleVPNConnectionTypePPaloAltoGlobalProtect
	// AppleVPNConnectionTypePZscalerPrivateAccess is a pointer to AppleVPNConnectionTypeVZscalerPrivateAccess
	AppleVPNConnectionTypePZscalerPrivateAccess = &_AppleVPNConnectionTypePZscalerPrivateAccess
	// AppleVPNConnectionTypePF5Access2018 is a pointer to AppleVPNConnectionTypeVF5Access2018
	AppleVPNConnectionTypePF5Access2018 = &_AppleVPNConnectionTypePF5Access2018
	// AppleVPNConnectionTypePCitrixSso is a pointer to AppleVPNConnectionTypeVCitrixSso
	AppleVPNConnectionTypePCitrixSso = &_AppleVPNConnectionTypePCitrixSso
	// AppleVPNConnectionTypePPaloAltoGlobalProtectV2 is a pointer to AppleVPNConnectionTypeVPaloAltoGlobalProtectV2
	AppleVPNConnectionTypePPaloAltoGlobalProtectV2 = &_AppleVPNConnectionTypePPaloAltoGlobalProtectV2
	// AppleVPNConnectionTypePIkEv2 is a pointer to AppleVPNConnectionTypeVIkEv2
	AppleVPNConnectionTypePIkEv2 = &_AppleVPNConnectionTypePIkEv2
)

var (
	_AppleVPNConnectionTypePCiscoAnyConnect            = AppleVPNConnectionTypeVCiscoAnyConnect
	_AppleVPNConnectionTypePPulseSecure                = AppleVPNConnectionTypeVPulseSecure
	_AppleVPNConnectionTypePF5EdgeClient               = AppleVPNConnectionTypeVF5EdgeClient
	_AppleVPNConnectionTypePDellSonicWallMobileConnect = AppleVPNConnectionTypeVDellSonicWallMobileConnect
	_AppleVPNConnectionTypePCheckPointCapsuleVPN       = AppleVPNConnectionTypeVCheckPointCapsuleVPN
	_AppleVPNConnectionTypePCustomVPN                  = AppleVPNConnectionTypeVCustomVPN
	_AppleVPNConnectionTypePCiscoIPSec                 = AppleVPNConnectionTypeVCiscoIPSec
	_AppleVPNConnectionTypePCitrix                     = AppleVPNConnectionTypeVCitrix
	_AppleVPNConnectionTypePCiscoAnyConnectV2          = AppleVPNConnectionTypeVCiscoAnyConnectV2
	_AppleVPNConnectionTypePPaloAltoGlobalProtect      = AppleVPNConnectionTypeVPaloAltoGlobalProtect
	_AppleVPNConnectionTypePZscalerPrivateAccess       = AppleVPNConnectionTypeVZscalerPrivateAccess
	_AppleVPNConnectionTypePF5Access2018               = AppleVPNConnectionTypeVF5Access2018
	_AppleVPNConnectionTypePCitrixSso                  = AppleVPNConnectionTypeVCitrixSso
	_AppleVPNConnectionTypePPaloAltoGlobalProtectV2    = AppleVPNConnectionTypeVPaloAltoGlobalProtectV2
	_AppleVPNConnectionTypePIkEv2                      = AppleVPNConnectionTypeVIkEv2
)
