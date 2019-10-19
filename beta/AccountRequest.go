// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AccountRequestBuilder is request builder for Account
type AccountRequestBuilder struct{ BaseRequestBuilder }

// Request returns AccountRequest
func (b *AccountRequestBuilder) Request() *AccountRequest {
	return &AccountRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AccountRequest is request for Account
type AccountRequest struct{ BaseRequest }

// Do performs HTTP request for Account
func (r *AccountRequest) Do(method, path string, reqObj interface{}) (resObj *Account, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for Account
func (r *AccountRequest) Get() (*Account, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for Account
func (r *AccountRequest) Update(reqObj *Account) (*Account, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for Account
func (r *AccountRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
