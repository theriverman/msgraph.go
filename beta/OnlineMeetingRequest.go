// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// OnlineMeetingRequestBuilder is request builder for OnlineMeeting
type OnlineMeetingRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnlineMeetingRequest
func (b *OnlineMeetingRequestBuilder) Request() *OnlineMeetingRequest {
	return &OnlineMeetingRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnlineMeetingRequest is request for OnlineMeeting
type OnlineMeetingRequest struct{ BaseRequest }

// Do performs HTTP request for OnlineMeeting
func (r *OnlineMeetingRequest) Do(method, path string, reqObj interface{}) (resObj *OnlineMeeting, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for OnlineMeeting
func (r *OnlineMeetingRequest) Get() (*OnlineMeeting, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for OnlineMeeting
func (r *OnlineMeetingRequest) Update(reqObj *OnlineMeeting) (*OnlineMeeting, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for OnlineMeeting
func (r *OnlineMeetingRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
