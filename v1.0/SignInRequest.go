// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SignInRequestBuilder is request builder for SignIn
type SignInRequestBuilder struct{ BaseRequestBuilder }

// Request returns SignInRequest
func (b *SignInRequestBuilder) Request() *SignInRequest {
	return &SignInRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SignInRequest is request for SignIn
type SignInRequest struct{ BaseRequest }

// Do performs HTTP request for SignIn
func (r *SignInRequest) Do(method, path string, reqObj interface{}) (resObj *SignIn, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for SignIn
func (r *SignInRequest) Get() (*SignIn, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for SignIn
func (r *SignInRequest) Update(reqObj *SignIn) (*SignIn, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for SignIn
func (r *SignInRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
