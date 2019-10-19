// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PrivilegedRoleAssignmentRequestObjectCancelRequestParameter undocumented
type PrivilegedRoleAssignmentRequestObjectCancelRequestParameter struct {
}

//
type PrivilegedRoleAssignmentRequestObjectCancelRequestBuilder struct{ BaseRequestBuilder }

// Cancel action undocumented
func (b *PrivilegedRoleAssignmentRequestObjectRequestBuilder) Cancel(reqObj *PrivilegedRoleAssignmentRequestObjectCancelRequestParameter) *PrivilegedRoleAssignmentRequestObjectCancelRequestBuilder {
	bb := &PrivilegedRoleAssignmentRequestObjectCancelRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/cancel"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type PrivilegedRoleAssignmentRequestObjectCancelRequest struct{ BaseRequest }

//
func (b *PrivilegedRoleAssignmentRequestObjectCancelRequestBuilder) Request() *PrivilegedRoleAssignmentRequestObjectCancelRequest {
	return &PrivilegedRoleAssignmentRequestObjectCancelRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *PrivilegedRoleAssignmentRequestObjectCancelRequest) Do(method, path string, reqObj interface{}) (resObj *PrivilegedRoleAssignmentRequestObject, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

//
func (r *PrivilegedRoleAssignmentRequestObjectCancelRequest) Post() (*PrivilegedRoleAssignmentRequestObject, error) {
	return r.Do("POST", "", r.requestObject)
}
