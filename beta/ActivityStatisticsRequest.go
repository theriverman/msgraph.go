// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ActivityStatisticsRequestBuilder is request builder for ActivityStatistics
type ActivityStatisticsRequestBuilder struct{ BaseRequestBuilder }

// Request returns ActivityStatisticsRequest
func (b *ActivityStatisticsRequestBuilder) Request() *ActivityStatisticsRequest {
	return &ActivityStatisticsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ActivityStatisticsRequest is request for ActivityStatistics
type ActivityStatisticsRequest struct{ BaseRequest }

// Do performs HTTP request for ActivityStatistics
func (r *ActivityStatisticsRequest) Do(method, path string, reqObj interface{}) (resObj *ActivityStatistics, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for ActivityStatistics
func (r *ActivityStatisticsRequest) Get() (*ActivityStatistics, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for ActivityStatistics
func (r *ActivityStatisticsRequest) Update(reqObj *ActivityStatistics) (*ActivityStatistics, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for ActivityStatistics
func (r *ActivityStatisticsRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
