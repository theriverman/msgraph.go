// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ColumnLinkRequestBuilder is request builder for ColumnLink
type ColumnLinkRequestBuilder struct{ BaseRequestBuilder }

// Request returns ColumnLinkRequest
func (b *ColumnLinkRequestBuilder) Request() *ColumnLinkRequest {
	return &ColumnLinkRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ColumnLinkRequest is request for ColumnLink
type ColumnLinkRequest struct{ BaseRequest }

// Do performs HTTP request for ColumnLink
func (r *ColumnLinkRequest) Do(method, path string, reqObj interface{}) (resObj *ColumnLink, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for ColumnLink
func (r *ColumnLinkRequest) Get() (*ColumnLink, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for ColumnLink
func (r *ColumnLinkRequest) Update(reqObj *ColumnLink) (*ColumnLink, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for ColumnLink
func (r *ColumnLinkRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
