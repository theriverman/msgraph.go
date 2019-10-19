// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ImportedAppleDeviceIdentityRequestBuilder is request builder for ImportedAppleDeviceIdentity
type ImportedAppleDeviceIdentityRequestBuilder struct{ BaseRequestBuilder }

// Request returns ImportedAppleDeviceIdentityRequest
func (b *ImportedAppleDeviceIdentityRequestBuilder) Request() *ImportedAppleDeviceIdentityRequest {
	return &ImportedAppleDeviceIdentityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ImportedAppleDeviceIdentityRequest is request for ImportedAppleDeviceIdentity
type ImportedAppleDeviceIdentityRequest struct{ BaseRequest }

// Do performs HTTP request for ImportedAppleDeviceIdentity
func (r *ImportedAppleDeviceIdentityRequest) Do(method, path string, reqObj interface{}) (resObj *ImportedAppleDeviceIdentity, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for ImportedAppleDeviceIdentity
func (r *ImportedAppleDeviceIdentityRequest) Get() (*ImportedAppleDeviceIdentity, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for ImportedAppleDeviceIdentity
func (r *ImportedAppleDeviceIdentityRequest) Update(reqObj *ImportedAppleDeviceIdentity) (*ImportedAppleDeviceIdentity, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for ImportedAppleDeviceIdentity
func (r *ImportedAppleDeviceIdentityRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
