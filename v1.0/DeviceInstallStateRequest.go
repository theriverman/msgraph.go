// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceInstallStateRequestBuilder is request builder for DeviceInstallState
type DeviceInstallStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceInstallStateRequest
func (b *DeviceInstallStateRequestBuilder) Request() *DeviceInstallStateRequest {
	return &DeviceInstallStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceInstallStateRequest is request for DeviceInstallState
type DeviceInstallStateRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceInstallState
func (r *DeviceInstallStateRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceInstallState, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for DeviceInstallState
func (r *DeviceInstallStateRequest) Get() (*DeviceInstallState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for DeviceInstallState
func (r *DeviceInstallStateRequest) Update(reqObj *DeviceInstallState) (*DeviceInstallState, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for DeviceInstallState
func (r *DeviceInstallStateRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
