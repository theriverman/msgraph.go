// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsPhone81TrustedRootCertificateRequestBuilder is request builder for WindowsPhone81TrustedRootCertificate
type WindowsPhone81TrustedRootCertificateRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsPhone81TrustedRootCertificateRequest
func (b *WindowsPhone81TrustedRootCertificateRequestBuilder) Request() *WindowsPhone81TrustedRootCertificateRequest {
	return &WindowsPhone81TrustedRootCertificateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsPhone81TrustedRootCertificateRequest is request for WindowsPhone81TrustedRootCertificate
type WindowsPhone81TrustedRootCertificateRequest struct{ BaseRequest }

// Do performs HTTP request for WindowsPhone81TrustedRootCertificate
func (r *WindowsPhone81TrustedRootCertificateRequest) Do(method, path string, reqObj interface{}) (resObj *WindowsPhone81TrustedRootCertificate, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for WindowsPhone81TrustedRootCertificate
func (r *WindowsPhone81TrustedRootCertificateRequest) Get() (*WindowsPhone81TrustedRootCertificate, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for WindowsPhone81TrustedRootCertificate
func (r *WindowsPhone81TrustedRootCertificateRequest) Update(reqObj *WindowsPhone81TrustedRootCertificate) (*WindowsPhone81TrustedRootCertificate, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for WindowsPhone81TrustedRootCertificate
func (r *WindowsPhone81TrustedRootCertificateRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
