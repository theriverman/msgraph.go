// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// LeakedCredentialsRiskEventRequestBuilder is request builder for LeakedCredentialsRiskEvent
type LeakedCredentialsRiskEventRequestBuilder struct{ BaseRequestBuilder }

// Request returns LeakedCredentialsRiskEventRequest
func (b *LeakedCredentialsRiskEventRequestBuilder) Request() *LeakedCredentialsRiskEventRequest {
	return &LeakedCredentialsRiskEventRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// LeakedCredentialsRiskEventRequest is request for LeakedCredentialsRiskEvent
type LeakedCredentialsRiskEventRequest struct{ BaseRequest }

// Do performs HTTP request for LeakedCredentialsRiskEvent
func (r *LeakedCredentialsRiskEventRequest) Do(method, path string, reqObj interface{}) (resObj *LeakedCredentialsRiskEvent, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for LeakedCredentialsRiskEvent
func (r *LeakedCredentialsRiskEventRequest) Get() (*LeakedCredentialsRiskEvent, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for LeakedCredentialsRiskEvent
func (r *LeakedCredentialsRiskEventRequest) Update(reqObj *LeakedCredentialsRiskEvent) (*LeakedCredentialsRiskEvent, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for LeakedCredentialsRiskEvent
func (r *LeakedCredentialsRiskEventRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
