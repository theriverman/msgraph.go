// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ResourceOperationRequestBuilder is request builder for ResourceOperation
type ResourceOperationRequestBuilder struct{ BaseRequestBuilder }

// Request returns ResourceOperationRequest
func (b *ResourceOperationRequestBuilder) Request() *ResourceOperationRequest {
	return &ResourceOperationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ResourceOperationRequest is request for ResourceOperation
type ResourceOperationRequest struct{ BaseRequest }

// Do performs HTTP request for ResourceOperation
func (r *ResourceOperationRequest) Do(method, path string, reqObj interface{}) (resObj *ResourceOperation, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for ResourceOperation
func (r *ResourceOperationRequest) Get() (*ResourceOperation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for ResourceOperation
func (r *ResourceOperationRequest) Update(reqObj *ResourceOperation) (*ResourceOperation, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for ResourceOperation
func (r *ResourceOperationRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
