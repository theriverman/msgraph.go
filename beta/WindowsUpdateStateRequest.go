// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsUpdateStateRequestBuilder is request builder for WindowsUpdateState
type WindowsUpdateStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsUpdateStateRequest
func (b *WindowsUpdateStateRequestBuilder) Request() *WindowsUpdateStateRequest {
	return &WindowsUpdateStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsUpdateStateRequest is request for WindowsUpdateState
type WindowsUpdateStateRequest struct{ BaseRequest }

// Do performs HTTP request for WindowsUpdateState
func (r *WindowsUpdateStateRequest) Do(method, path string, reqObj interface{}) (resObj *WindowsUpdateState, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for WindowsUpdateState
func (r *WindowsUpdateStateRequest) Get() (*WindowsUpdateState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for WindowsUpdateState
func (r *WindowsUpdateStateRequest) Update(reqObj *WindowsUpdateState) (*WindowsUpdateState, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for WindowsUpdateState
func (r *WindowsUpdateStateRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
