// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DepMacOSEnrollmentProfileRequestBuilder is request builder for DepMacOSEnrollmentProfile
type DepMacOSEnrollmentProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns DepMacOSEnrollmentProfileRequest
func (b *DepMacOSEnrollmentProfileRequestBuilder) Request() *DepMacOSEnrollmentProfileRequest {
	return &DepMacOSEnrollmentProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DepMacOSEnrollmentProfileRequest is request for DepMacOSEnrollmentProfile
type DepMacOSEnrollmentProfileRequest struct{ BaseRequest }

// Do performs HTTP request for DepMacOSEnrollmentProfile
func (r *DepMacOSEnrollmentProfileRequest) Do(method, path string, reqObj interface{}) (resObj *DepMacOSEnrollmentProfile, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for DepMacOSEnrollmentProfile
func (r *DepMacOSEnrollmentProfileRequest) Get() (*DepMacOSEnrollmentProfile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for DepMacOSEnrollmentProfile
func (r *DepMacOSEnrollmentProfileRequest) Update(reqObj *DepMacOSEnrollmentProfile) (*DepMacOSEnrollmentProfile, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for DepMacOSEnrollmentProfile
func (r *DepMacOSEnrollmentProfileRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
