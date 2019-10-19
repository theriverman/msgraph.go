// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ScopedRoleMembershipRequestBuilder is request builder for ScopedRoleMembership
type ScopedRoleMembershipRequestBuilder struct{ BaseRequestBuilder }

// Request returns ScopedRoleMembershipRequest
func (b *ScopedRoleMembershipRequestBuilder) Request() *ScopedRoleMembershipRequest {
	return &ScopedRoleMembershipRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ScopedRoleMembershipRequest is request for ScopedRoleMembership
type ScopedRoleMembershipRequest struct{ BaseRequest }

// Do performs HTTP request for ScopedRoleMembership
func (r *ScopedRoleMembershipRequest) Do(method, path string, reqObj interface{}) (resObj *ScopedRoleMembership, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for ScopedRoleMembership
func (r *ScopedRoleMembershipRequest) Get() (*ScopedRoleMembership, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for ScopedRoleMembership
func (r *ScopedRoleMembershipRequest) Update(reqObj *ScopedRoleMembership) (*ScopedRoleMembership, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for ScopedRoleMembership
func (r *ScopedRoleMembershipRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
