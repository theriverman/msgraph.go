// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidDeviceOwnerEnrollmentProfileRevokeTokenRequestParameter undocumented
type AndroidDeviceOwnerEnrollmentProfileRevokeTokenRequestParameter struct {
}

// AndroidDeviceOwnerEnrollmentProfileCreateTokenRequestParameter undocumented
type AndroidDeviceOwnerEnrollmentProfileCreateTokenRequestParameter struct {
	// TokenValidityInSeconds undocumented
	TokenValidityInSeconds *int `json:"tokenValidityInSeconds,omitempty"`
}

//
type AndroidDeviceOwnerEnrollmentProfileRevokeTokenRequestBuilder struct{ BaseRequestBuilder }

// RevokeToken action undocumented
func (b *AndroidDeviceOwnerEnrollmentProfileRequestBuilder) RevokeToken(reqObj *AndroidDeviceOwnerEnrollmentProfileRevokeTokenRequestParameter) *AndroidDeviceOwnerEnrollmentProfileRevokeTokenRequestBuilder {
	bb := &AndroidDeviceOwnerEnrollmentProfileRevokeTokenRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/revokeToken"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type AndroidDeviceOwnerEnrollmentProfileRevokeTokenRequest struct{ BaseRequest }

//
func (b *AndroidDeviceOwnerEnrollmentProfileRevokeTokenRequestBuilder) Request() *AndroidDeviceOwnerEnrollmentProfileRevokeTokenRequest {
	return &AndroidDeviceOwnerEnrollmentProfileRevokeTokenRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *AndroidDeviceOwnerEnrollmentProfileRevokeTokenRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *AndroidDeviceOwnerEnrollmentProfileRevokeTokenRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type AndroidDeviceOwnerEnrollmentProfileCreateTokenRequestBuilder struct{ BaseRequestBuilder }

// CreateToken action undocumented
func (b *AndroidDeviceOwnerEnrollmentProfileRequestBuilder) CreateToken(reqObj *AndroidDeviceOwnerEnrollmentProfileCreateTokenRequestParameter) *AndroidDeviceOwnerEnrollmentProfileCreateTokenRequestBuilder {
	bb := &AndroidDeviceOwnerEnrollmentProfileCreateTokenRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/createToken"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type AndroidDeviceOwnerEnrollmentProfileCreateTokenRequest struct{ BaseRequest }

//
func (b *AndroidDeviceOwnerEnrollmentProfileCreateTokenRequestBuilder) Request() *AndroidDeviceOwnerEnrollmentProfileCreateTokenRequest {
	return &AndroidDeviceOwnerEnrollmentProfileCreateTokenRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *AndroidDeviceOwnerEnrollmentProfileCreateTokenRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *AndroidDeviceOwnerEnrollmentProfileCreateTokenRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}
