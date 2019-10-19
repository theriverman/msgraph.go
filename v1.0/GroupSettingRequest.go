// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// GroupSettingRequestBuilder is request builder for GroupSetting
type GroupSettingRequestBuilder struct{ BaseRequestBuilder }

// Request returns GroupSettingRequest
func (b *GroupSettingRequestBuilder) Request() *GroupSettingRequest {
	return &GroupSettingRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// GroupSettingRequest is request for GroupSetting
type GroupSettingRequest struct{ BaseRequest }

// Do performs HTTP request for GroupSetting
func (r *GroupSettingRequest) Do(method, path string, reqObj interface{}) (resObj *GroupSetting, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for GroupSetting
func (r *GroupSettingRequest) Get() (*GroupSetting, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for GroupSetting
func (r *GroupSettingRequest) Update(reqObj *GroupSetting) (*GroupSetting, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for GroupSetting
func (r *GroupSettingRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
