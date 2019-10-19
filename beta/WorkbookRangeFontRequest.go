// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WorkbookRangeFontRequestBuilder is request builder for WorkbookRangeFont
type WorkbookRangeFontRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookRangeFontRequest
func (b *WorkbookRangeFontRequestBuilder) Request() *WorkbookRangeFontRequest {
	return &WorkbookRangeFontRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookRangeFontRequest is request for WorkbookRangeFont
type WorkbookRangeFontRequest struct{ BaseRequest }

// Do performs HTTP request for WorkbookRangeFont
func (r *WorkbookRangeFontRequest) Do(method, path string, reqObj interface{}) (resObj *WorkbookRangeFont, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for WorkbookRangeFont
func (r *WorkbookRangeFontRequest) Get() (*WorkbookRangeFont, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for WorkbookRangeFont
func (r *WorkbookRangeFontRequest) Update(reqObj *WorkbookRangeFont) (*WorkbookRangeFont, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for WorkbookRangeFont
func (r *WorkbookRangeFontRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
