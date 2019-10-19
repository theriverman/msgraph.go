// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MacOSEnterpriseWiFiConfigurationRequestBuilder is request builder for MacOSEnterpriseWiFiConfiguration
type MacOSEnterpriseWiFiConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacOSEnterpriseWiFiConfigurationRequest
func (b *MacOSEnterpriseWiFiConfigurationRequestBuilder) Request() *MacOSEnterpriseWiFiConfigurationRequest {
	return &MacOSEnterpriseWiFiConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MacOSEnterpriseWiFiConfigurationRequest is request for MacOSEnterpriseWiFiConfiguration
type MacOSEnterpriseWiFiConfigurationRequest struct{ BaseRequest }

// Do performs HTTP request for MacOSEnterpriseWiFiConfiguration
func (r *MacOSEnterpriseWiFiConfigurationRequest) Do(method, path string, reqObj interface{}) (resObj *MacOSEnterpriseWiFiConfiguration, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for MacOSEnterpriseWiFiConfiguration
func (r *MacOSEnterpriseWiFiConfigurationRequest) Get() (*MacOSEnterpriseWiFiConfiguration, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for MacOSEnterpriseWiFiConfiguration
func (r *MacOSEnterpriseWiFiConfigurationRequest) Update(reqObj *MacOSEnterpriseWiFiConfiguration) (*MacOSEnterpriseWiFiConfiguration, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for MacOSEnterpriseWiFiConfiguration
func (r *MacOSEnterpriseWiFiConfigurationRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// IdentityCertificateForClientAuthentication is navigation property
func (b *MacOSEnterpriseWiFiConfigurationRequestBuilder) IdentityCertificateForClientAuthentication() *MacOSCertificateProfileBaseRequestBuilder {
	bb := &MacOSCertificateProfileBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/identityCertificateForClientAuthentication"
	return bb
}

// RootCertificateForServerValidation is navigation property
func (b *MacOSEnterpriseWiFiConfigurationRequestBuilder) RootCertificateForServerValidation() *MacOSTrustedRootCertificateRequestBuilder {
	bb := &MacOSTrustedRootCertificateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rootCertificateForServerValidation"
	return bb
}
