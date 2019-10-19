// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SecurityActionRequestBuilder is request builder for SecurityAction
type SecurityActionRequestBuilder struct{ BaseRequestBuilder }

// Request returns SecurityActionRequest
func (b *SecurityActionRequestBuilder) Request() *SecurityActionRequest {
	return &SecurityActionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SecurityActionRequest is request for SecurityAction
type SecurityActionRequest struct{ BaseRequest }

// Do performs HTTP request for SecurityAction
func (r *SecurityActionRequest) Do(method, path string, reqObj interface{}) (resObj *SecurityAction, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for SecurityAction
func (r *SecurityActionRequest) Get() (*SecurityAction, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for SecurityAction
func (r *SecurityActionRequest) Update(reqObj *SecurityAction) (*SecurityAction, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for SecurityAction
func (r *SecurityActionRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
