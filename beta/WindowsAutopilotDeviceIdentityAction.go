// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsAutopilotDeviceIdentityAssignUserToDeviceRequestParameter undocumented
type WindowsAutopilotDeviceIdentityAssignUserToDeviceRequestParameter struct {
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// AddressableUserName undocumented
	AddressableUserName *string `json:"addressableUserName,omitempty"`
}

// WindowsAutopilotDeviceIdentityUnassignUserFromDeviceRequestParameter undocumented
type WindowsAutopilotDeviceIdentityUnassignUserFromDeviceRequestParameter struct {
}

// WindowsAutopilotDeviceIdentityAssignResourceAccountToDeviceRequestParameter undocumented
type WindowsAutopilotDeviceIdentityAssignResourceAccountToDeviceRequestParameter struct {
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// AddressableUserName undocumented
	AddressableUserName *string `json:"addressableUserName,omitempty"`
	// ResourceAccountName undocumented
	ResourceAccountName *string `json:"resourceAccountName,omitempty"`
}

// WindowsAutopilotDeviceIdentityUnassignResourceAccountFromDeviceRequestParameter undocumented
type WindowsAutopilotDeviceIdentityUnassignResourceAccountFromDeviceRequestParameter struct {
}

//
type WindowsAutopilotDeviceIdentityAssignUserToDeviceRequestBuilder struct{ BaseRequestBuilder }

// AssignUserToDevice action undocumented
func (b *WindowsAutopilotDeviceIdentityRequestBuilder) AssignUserToDevice(reqObj *WindowsAutopilotDeviceIdentityAssignUserToDeviceRequestParameter) *WindowsAutopilotDeviceIdentityAssignUserToDeviceRequestBuilder {
	bb := &WindowsAutopilotDeviceIdentityAssignUserToDeviceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assignUserToDevice"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WindowsAutopilotDeviceIdentityAssignUserToDeviceRequest struct{ BaseRequest }

//
func (b *WindowsAutopilotDeviceIdentityAssignUserToDeviceRequestBuilder) Request() *WindowsAutopilotDeviceIdentityAssignUserToDeviceRequest {
	return &WindowsAutopilotDeviceIdentityAssignUserToDeviceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WindowsAutopilotDeviceIdentityAssignUserToDeviceRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *WindowsAutopilotDeviceIdentityAssignUserToDeviceRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type WindowsAutopilotDeviceIdentityUnassignUserFromDeviceRequestBuilder struct{ BaseRequestBuilder }

// UnassignUserFromDevice action undocumented
func (b *WindowsAutopilotDeviceIdentityRequestBuilder) UnassignUserFromDevice(reqObj *WindowsAutopilotDeviceIdentityUnassignUserFromDeviceRequestParameter) *WindowsAutopilotDeviceIdentityUnassignUserFromDeviceRequestBuilder {
	bb := &WindowsAutopilotDeviceIdentityUnassignUserFromDeviceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/unassignUserFromDevice"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WindowsAutopilotDeviceIdentityUnassignUserFromDeviceRequest struct{ BaseRequest }

//
func (b *WindowsAutopilotDeviceIdentityUnassignUserFromDeviceRequestBuilder) Request() *WindowsAutopilotDeviceIdentityUnassignUserFromDeviceRequest {
	return &WindowsAutopilotDeviceIdentityUnassignUserFromDeviceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WindowsAutopilotDeviceIdentityUnassignUserFromDeviceRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *WindowsAutopilotDeviceIdentityUnassignUserFromDeviceRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type WindowsAutopilotDeviceIdentityAssignResourceAccountToDeviceRequestBuilder struct{ BaseRequestBuilder }

// AssignResourceAccountToDevice action undocumented
func (b *WindowsAutopilotDeviceIdentityRequestBuilder) AssignResourceAccountToDevice(reqObj *WindowsAutopilotDeviceIdentityAssignResourceAccountToDeviceRequestParameter) *WindowsAutopilotDeviceIdentityAssignResourceAccountToDeviceRequestBuilder {
	bb := &WindowsAutopilotDeviceIdentityAssignResourceAccountToDeviceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assignResourceAccountToDevice"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WindowsAutopilotDeviceIdentityAssignResourceAccountToDeviceRequest struct{ BaseRequest }

//
func (b *WindowsAutopilotDeviceIdentityAssignResourceAccountToDeviceRequestBuilder) Request() *WindowsAutopilotDeviceIdentityAssignResourceAccountToDeviceRequest {
	return &WindowsAutopilotDeviceIdentityAssignResourceAccountToDeviceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WindowsAutopilotDeviceIdentityAssignResourceAccountToDeviceRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *WindowsAutopilotDeviceIdentityAssignResourceAccountToDeviceRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type WindowsAutopilotDeviceIdentityUnassignResourceAccountFromDeviceRequestBuilder struct{ BaseRequestBuilder }

// UnassignResourceAccountFromDevice action undocumented
func (b *WindowsAutopilotDeviceIdentityRequestBuilder) UnassignResourceAccountFromDevice(reqObj *WindowsAutopilotDeviceIdentityUnassignResourceAccountFromDeviceRequestParameter) *WindowsAutopilotDeviceIdentityUnassignResourceAccountFromDeviceRequestBuilder {
	bb := &WindowsAutopilotDeviceIdentityUnassignResourceAccountFromDeviceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/unassignResourceAccountFromDevice"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WindowsAutopilotDeviceIdentityUnassignResourceAccountFromDeviceRequest struct{ BaseRequest }

//
func (b *WindowsAutopilotDeviceIdentityUnassignResourceAccountFromDeviceRequestBuilder) Request() *WindowsAutopilotDeviceIdentityUnassignResourceAccountFromDeviceRequest {
	return &WindowsAutopilotDeviceIdentityUnassignResourceAccountFromDeviceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WindowsAutopilotDeviceIdentityUnassignResourceAccountFromDeviceRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *WindowsAutopilotDeviceIdentityUnassignResourceAccountFromDeviceRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}
