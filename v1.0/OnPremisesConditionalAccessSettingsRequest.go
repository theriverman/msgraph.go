// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// OnPremisesConditionalAccessSettingsRequestBuilder is request builder for OnPremisesConditionalAccessSettings
type OnPremisesConditionalAccessSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnPremisesConditionalAccessSettingsRequest
func (b *OnPremisesConditionalAccessSettingsRequestBuilder) Request() *OnPremisesConditionalAccessSettingsRequest {
	return &OnPremisesConditionalAccessSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnPremisesConditionalAccessSettingsRequest is request for OnPremisesConditionalAccessSettings
type OnPremisesConditionalAccessSettingsRequest struct{ BaseRequest }

// Do performs HTTP request for OnPremisesConditionalAccessSettings
func (r *OnPremisesConditionalAccessSettingsRequest) Do(method, path string, reqObj interface{}) (resObj *OnPremisesConditionalAccessSettings, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for OnPremisesConditionalAccessSettings
func (r *OnPremisesConditionalAccessSettingsRequest) Get() (*OnPremisesConditionalAccessSettings, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for OnPremisesConditionalAccessSettings
func (r *OnPremisesConditionalAccessSettingsRequest) Update(reqObj *OnPremisesConditionalAccessSettings) (*OnPremisesConditionalAccessSettings, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for OnPremisesConditionalAccessSettings
func (r *OnPremisesConditionalAccessSettingsRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
