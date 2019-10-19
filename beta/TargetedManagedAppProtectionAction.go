// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// TargetedManagedAppProtectionAssignRequestParameter undocumented
type TargetedManagedAppProtectionAssignRequestParameter struct {
	// Assignments undocumented
	Assignments []TargetedManagedAppPolicyAssignment `json:"assignments,omitempty"`
}

//
type TargetedManagedAppProtectionAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *TargetedManagedAppProtectionRequestBuilder) Assign(reqObj *TargetedManagedAppProtectionAssignRequestParameter) *TargetedManagedAppProtectionAssignRequestBuilder {
	bb := &TargetedManagedAppProtectionAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type TargetedManagedAppProtectionAssignRequest struct{ BaseRequest }

//
func (b *TargetedManagedAppProtectionAssignRequestBuilder) Request() *TargetedManagedAppProtectionAssignRequest {
	return &TargetedManagedAppProtectionAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *TargetedManagedAppProtectionAssignRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *TargetedManagedAppProtectionAssignRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}
