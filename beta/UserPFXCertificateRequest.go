// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// UserPFXCertificateRequestBuilder is request builder for UserPFXCertificate
type UserPFXCertificateRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserPFXCertificateRequest
func (b *UserPFXCertificateRequestBuilder) Request() *UserPFXCertificateRequest {
	return &UserPFXCertificateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserPFXCertificateRequest is request for UserPFXCertificate
type UserPFXCertificateRequest struct{ BaseRequest }

// Do performs HTTP request for UserPFXCertificate
func (r *UserPFXCertificateRequest) Do(method, path string, reqObj interface{}) (resObj *UserPFXCertificate, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for UserPFXCertificate
func (r *UserPFXCertificateRequest) Get() (*UserPFXCertificate, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for UserPFXCertificate
func (r *UserPFXCertificateRequest) Update(reqObj *UserPFXCertificate) (*UserPFXCertificate, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for UserPFXCertificate
func (r *UserPFXCertificateRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
