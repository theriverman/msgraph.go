// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AppleEnrollmentProfileAssignmentRequestBuilder is request builder for AppleEnrollmentProfileAssignment
type AppleEnrollmentProfileAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns AppleEnrollmentProfileAssignmentRequest
func (b *AppleEnrollmentProfileAssignmentRequestBuilder) Request() *AppleEnrollmentProfileAssignmentRequest {
	return &AppleEnrollmentProfileAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AppleEnrollmentProfileAssignmentRequest is request for AppleEnrollmentProfileAssignment
type AppleEnrollmentProfileAssignmentRequest struct{ BaseRequest }

// Do performs HTTP request for AppleEnrollmentProfileAssignment
func (r *AppleEnrollmentProfileAssignmentRequest) Do(method, path string, reqObj interface{}) (resObj *AppleEnrollmentProfileAssignment, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for AppleEnrollmentProfileAssignment
func (r *AppleEnrollmentProfileAssignmentRequest) Get() (*AppleEnrollmentProfileAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for AppleEnrollmentProfileAssignment
func (r *AppleEnrollmentProfileAssignmentRequest) Update(reqObj *AppleEnrollmentProfileAssignment) (*AppleEnrollmentProfileAssignment, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for AppleEnrollmentProfileAssignment
func (r *AppleEnrollmentProfileAssignmentRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
