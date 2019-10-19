// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidForWorkTrustedRootCertificateRequestBuilder is request builder for AndroidForWorkTrustedRootCertificate
type AndroidForWorkTrustedRootCertificateRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidForWorkTrustedRootCertificateRequest
func (b *AndroidForWorkTrustedRootCertificateRequestBuilder) Request() *AndroidForWorkTrustedRootCertificateRequest {
	return &AndroidForWorkTrustedRootCertificateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AndroidForWorkTrustedRootCertificateRequest is request for AndroidForWorkTrustedRootCertificate
type AndroidForWorkTrustedRootCertificateRequest struct{ BaseRequest }

// Do performs HTTP request for AndroidForWorkTrustedRootCertificate
func (r *AndroidForWorkTrustedRootCertificateRequest) Do(method, path string, reqObj interface{}) (resObj *AndroidForWorkTrustedRootCertificate, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for AndroidForWorkTrustedRootCertificate
func (r *AndroidForWorkTrustedRootCertificateRequest) Get() (*AndroidForWorkTrustedRootCertificate, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for AndroidForWorkTrustedRootCertificate
func (r *AndroidForWorkTrustedRootCertificateRequest) Update(reqObj *AndroidForWorkTrustedRootCertificate) (*AndroidForWorkTrustedRootCertificate, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for AndroidForWorkTrustedRootCertificate
func (r *AndroidForWorkTrustedRootCertificateRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
