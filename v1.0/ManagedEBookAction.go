// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ManagedEBookAssignRequestParameter undocumented
type ManagedEBookAssignRequestParameter struct {
	// ManagedEBookAssignments undocumented
	ManagedEBookAssignments []ManagedEBookAssignment `json:"managedEBookAssignments,omitempty"`
}

//
type ManagedEBookAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *ManagedEBookRequestBuilder) Assign(reqObj *ManagedEBookAssignRequestParameter) *ManagedEBookAssignRequestBuilder {
	bb := &ManagedEBookAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedEBookAssignRequest struct{ BaseRequest }

//
func (b *ManagedEBookAssignRequestBuilder) Request() *ManagedEBookAssignRequest {
	return &ManagedEBookAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedEBookAssignRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *ManagedEBookAssignRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}
