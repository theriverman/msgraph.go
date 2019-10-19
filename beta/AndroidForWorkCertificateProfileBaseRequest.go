// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidForWorkCertificateProfileBaseRequestBuilder is request builder for AndroidForWorkCertificateProfileBase
type AndroidForWorkCertificateProfileBaseRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidForWorkCertificateProfileBaseRequest
func (b *AndroidForWorkCertificateProfileBaseRequestBuilder) Request() *AndroidForWorkCertificateProfileBaseRequest {
	return &AndroidForWorkCertificateProfileBaseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AndroidForWorkCertificateProfileBaseRequest is request for AndroidForWorkCertificateProfileBase
type AndroidForWorkCertificateProfileBaseRequest struct{ BaseRequest }

// Do performs HTTP request for AndroidForWorkCertificateProfileBase
func (r *AndroidForWorkCertificateProfileBaseRequest) Do(method, path string, reqObj interface{}) (resObj *AndroidForWorkCertificateProfileBase, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for AndroidForWorkCertificateProfileBase
func (r *AndroidForWorkCertificateProfileBaseRequest) Get() (*AndroidForWorkCertificateProfileBase, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for AndroidForWorkCertificateProfileBase
func (r *AndroidForWorkCertificateProfileBaseRequest) Update(reqObj *AndroidForWorkCertificateProfileBase) (*AndroidForWorkCertificateProfileBase, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for AndroidForWorkCertificateProfileBase
func (r *AndroidForWorkCertificateProfileBaseRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// RootCertificate is navigation property
func (b *AndroidForWorkCertificateProfileBaseRequestBuilder) RootCertificate() *AndroidForWorkTrustedRootCertificateRequestBuilder {
	bb := &AndroidForWorkTrustedRootCertificateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rootCertificate"
	return bb
}
