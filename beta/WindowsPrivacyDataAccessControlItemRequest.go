// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsPrivacyDataAccessControlItemRequestBuilder is request builder for WindowsPrivacyDataAccessControlItem
type WindowsPrivacyDataAccessControlItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsPrivacyDataAccessControlItemRequest
func (b *WindowsPrivacyDataAccessControlItemRequestBuilder) Request() *WindowsPrivacyDataAccessControlItemRequest {
	return &WindowsPrivacyDataAccessControlItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsPrivacyDataAccessControlItemRequest is request for WindowsPrivacyDataAccessControlItem
type WindowsPrivacyDataAccessControlItemRequest struct{ BaseRequest }

// Do performs HTTP request for WindowsPrivacyDataAccessControlItem
func (r *WindowsPrivacyDataAccessControlItemRequest) Do(method, path string, reqObj interface{}) (resObj *WindowsPrivacyDataAccessControlItem, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for WindowsPrivacyDataAccessControlItem
func (r *WindowsPrivacyDataAccessControlItemRequest) Get() (*WindowsPrivacyDataAccessControlItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for WindowsPrivacyDataAccessControlItem
func (r *WindowsPrivacyDataAccessControlItemRequest) Update(reqObj *WindowsPrivacyDataAccessControlItem) (*WindowsPrivacyDataAccessControlItem, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for WindowsPrivacyDataAccessControlItem
func (r *WindowsPrivacyDataAccessControlItemRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
