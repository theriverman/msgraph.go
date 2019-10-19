// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementScriptDeviceStateRequestBuilder is request builder for DeviceManagementScriptDeviceState
type DeviceManagementScriptDeviceStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceManagementScriptDeviceStateRequest
func (b *DeviceManagementScriptDeviceStateRequestBuilder) Request() *DeviceManagementScriptDeviceStateRequest {
	return &DeviceManagementScriptDeviceStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceManagementScriptDeviceStateRequest is request for DeviceManagementScriptDeviceState
type DeviceManagementScriptDeviceStateRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceManagementScriptDeviceState
func (r *DeviceManagementScriptDeviceStateRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceManagementScriptDeviceState, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for DeviceManagementScriptDeviceState
func (r *DeviceManagementScriptDeviceStateRequest) Get() (*DeviceManagementScriptDeviceState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for DeviceManagementScriptDeviceState
func (r *DeviceManagementScriptDeviceStateRequest) Update(reqObj *DeviceManagementScriptDeviceState) (*DeviceManagementScriptDeviceState, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for DeviceManagementScriptDeviceState
func (r *DeviceManagementScriptDeviceStateRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// ManagedDevice is navigation property
func (b *DeviceManagementScriptDeviceStateRequestBuilder) ManagedDevice() *ManagedDeviceRequestBuilder {
	bb := &ManagedDeviceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/managedDevice"
	return bb
}
