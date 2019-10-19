// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// RestrictedSignInRequestBuilder is request builder for RestrictedSignIn
type RestrictedSignInRequestBuilder struct{ BaseRequestBuilder }

// Request returns RestrictedSignInRequest
func (b *RestrictedSignInRequestBuilder) Request() *RestrictedSignInRequest {
	return &RestrictedSignInRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RestrictedSignInRequest is request for RestrictedSignIn
type RestrictedSignInRequest struct{ BaseRequest }

// Do performs HTTP request for RestrictedSignIn
func (r *RestrictedSignInRequest) Do(method, path string, reqObj interface{}) (resObj *RestrictedSignIn, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for RestrictedSignIn
func (r *RestrictedSignInRequest) Get() (*RestrictedSignIn, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for RestrictedSignIn
func (r *RestrictedSignInRequest) Update(reqObj *RestrictedSignIn) (*RestrictedSignIn, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for RestrictedSignIn
func (r *RestrictedSignInRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
