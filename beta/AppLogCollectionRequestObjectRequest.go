// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AppLogCollectionRequestObjectRequestBuilder is request builder for AppLogCollectionRequestObject
type AppLogCollectionRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns AppLogCollectionRequestObjectRequest
func (b *AppLogCollectionRequestObjectRequestBuilder) Request() *AppLogCollectionRequestObjectRequest {
	return &AppLogCollectionRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AppLogCollectionRequestObjectRequest is request for AppLogCollectionRequestObject
type AppLogCollectionRequestObjectRequest struct{ BaseRequest }

// Do performs HTTP request for AppLogCollectionRequestObject
func (r *AppLogCollectionRequestObjectRequest) Do(method, path string, reqObj interface{}) (resObj *AppLogCollectionRequestObject, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for AppLogCollectionRequestObject
func (r *AppLogCollectionRequestObjectRequest) Get() (*AppLogCollectionRequestObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for AppLogCollectionRequestObject
func (r *AppLogCollectionRequestObjectRequest) Update(reqObj *AppLogCollectionRequestObject) (*AppLogCollectionRequestObject, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for AppLogCollectionRequestObject
func (r *AppLogCollectionRequestObjectRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
