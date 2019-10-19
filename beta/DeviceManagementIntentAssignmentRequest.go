// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementIntentAssignmentRequestBuilder is request builder for DeviceManagementIntentAssignment
type DeviceManagementIntentAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceManagementIntentAssignmentRequest
func (b *DeviceManagementIntentAssignmentRequestBuilder) Request() *DeviceManagementIntentAssignmentRequest {
	return &DeviceManagementIntentAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceManagementIntentAssignmentRequest is request for DeviceManagementIntentAssignment
type DeviceManagementIntentAssignmentRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceManagementIntentAssignment
func (r *DeviceManagementIntentAssignmentRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceManagementIntentAssignment, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for DeviceManagementIntentAssignment
func (r *DeviceManagementIntentAssignmentRequest) Get() (*DeviceManagementIntentAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for DeviceManagementIntentAssignment
func (r *DeviceManagementIntentAssignmentRequest) Update(reqObj *DeviceManagementIntentAssignment) (*DeviceManagementIntentAssignment, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for DeviceManagementIntentAssignment
func (r *DeviceManagementIntentAssignmentRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
