// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// TeamsAsyncOperationRequestBuilder is request builder for TeamsAsyncOperation
type TeamsAsyncOperationRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamsAsyncOperationRequest
func (b *TeamsAsyncOperationRequestBuilder) Request() *TeamsAsyncOperationRequest {
	return &TeamsAsyncOperationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamsAsyncOperationRequest is request for TeamsAsyncOperation
type TeamsAsyncOperationRequest struct{ BaseRequest }

// Do performs HTTP request for TeamsAsyncOperation
func (r *TeamsAsyncOperationRequest) Do(method, path string, reqObj interface{}) (resObj *TeamsAsyncOperation, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for TeamsAsyncOperation
func (r *TeamsAsyncOperationRequest) Get() (*TeamsAsyncOperation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for TeamsAsyncOperation
func (r *TeamsAsyncOperationRequest) Update(reqObj *TeamsAsyncOperation) (*TeamsAsyncOperation, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for TeamsAsyncOperation
func (r *TeamsAsyncOperationRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
