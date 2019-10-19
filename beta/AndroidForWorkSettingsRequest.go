// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidForWorkSettingsRequestBuilder is request builder for AndroidForWorkSettings
type AndroidForWorkSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidForWorkSettingsRequest
func (b *AndroidForWorkSettingsRequestBuilder) Request() *AndroidForWorkSettingsRequest {
	return &AndroidForWorkSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AndroidForWorkSettingsRequest is request for AndroidForWorkSettings
type AndroidForWorkSettingsRequest struct{ BaseRequest }

// Do performs HTTP request for AndroidForWorkSettings
func (r *AndroidForWorkSettingsRequest) Do(method, path string, reqObj interface{}) (resObj *AndroidForWorkSettings, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for AndroidForWorkSettings
func (r *AndroidForWorkSettingsRequest) Get() (*AndroidForWorkSettings, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for AndroidForWorkSettings
func (r *AndroidForWorkSettingsRequest) Update(reqObj *AndroidForWorkSettings) (*AndroidForWorkSettings, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for AndroidForWorkSettings
func (r *AndroidForWorkSettingsRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
