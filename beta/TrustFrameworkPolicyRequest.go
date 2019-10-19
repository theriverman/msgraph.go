// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// TrustFrameworkPolicyRequestBuilder is request builder for TrustFrameworkPolicy
type TrustFrameworkPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns TrustFrameworkPolicyRequest
func (b *TrustFrameworkPolicyRequestBuilder) Request() *TrustFrameworkPolicyRequest {
	return &TrustFrameworkPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TrustFrameworkPolicyRequest is request for TrustFrameworkPolicy
type TrustFrameworkPolicyRequest struct{ BaseRequest }

// Do performs HTTP request for TrustFrameworkPolicy
func (r *TrustFrameworkPolicyRequest) Do(method, path string, reqObj interface{}) (resObj *TrustFrameworkPolicy, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for TrustFrameworkPolicy
func (r *TrustFrameworkPolicyRequest) Get() (*TrustFrameworkPolicy, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for TrustFrameworkPolicy
func (r *TrustFrameworkPolicyRequest) Update(reqObj *TrustFrameworkPolicy) (*TrustFrameworkPolicy, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for TrustFrameworkPolicy
func (r *TrustFrameworkPolicyRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
