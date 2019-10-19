// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementIntentUserStateSummaryRequestBuilder is request builder for DeviceManagementIntentUserStateSummary
type DeviceManagementIntentUserStateSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceManagementIntentUserStateSummaryRequest
func (b *DeviceManagementIntentUserStateSummaryRequestBuilder) Request() *DeviceManagementIntentUserStateSummaryRequest {
	return &DeviceManagementIntentUserStateSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceManagementIntentUserStateSummaryRequest is request for DeviceManagementIntentUserStateSummary
type DeviceManagementIntentUserStateSummaryRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceManagementIntentUserStateSummary
func (r *DeviceManagementIntentUserStateSummaryRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceManagementIntentUserStateSummary, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for DeviceManagementIntentUserStateSummary
func (r *DeviceManagementIntentUserStateSummaryRequest) Get() (*DeviceManagementIntentUserStateSummary, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for DeviceManagementIntentUserStateSummary
func (r *DeviceManagementIntentUserStateSummaryRequest) Update(reqObj *DeviceManagementIntentUserStateSummary) (*DeviceManagementIntentUserStateSummary, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for DeviceManagementIntentUserStateSummary
func (r *DeviceManagementIntentUserStateSummaryRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
