// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WorkbookChartLineFormatRequestBuilder is request builder for WorkbookChartLineFormat
type WorkbookChartLineFormatRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookChartLineFormatRequest
func (b *WorkbookChartLineFormatRequestBuilder) Request() *WorkbookChartLineFormatRequest {
	return &WorkbookChartLineFormatRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookChartLineFormatRequest is request for WorkbookChartLineFormat
type WorkbookChartLineFormatRequest struct{ BaseRequest }

// Do performs HTTP request for WorkbookChartLineFormat
func (r *WorkbookChartLineFormatRequest) Do(method, path string, reqObj interface{}) (resObj *WorkbookChartLineFormat, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for WorkbookChartLineFormat
func (r *WorkbookChartLineFormatRequest) Get() (*WorkbookChartLineFormat, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for WorkbookChartLineFormat
func (r *WorkbookChartLineFormatRequest) Update(reqObj *WorkbookChartLineFormat) (*WorkbookChartLineFormat, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for WorkbookChartLineFormat
func (r *WorkbookChartLineFormatRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
