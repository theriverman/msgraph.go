// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// FilterOperatorSchemaRequestBuilder is request builder for FilterOperatorSchema
type FilterOperatorSchemaRequestBuilder struct{ BaseRequestBuilder }

// Request returns FilterOperatorSchemaRequest
func (b *FilterOperatorSchemaRequestBuilder) Request() *FilterOperatorSchemaRequest {
	return &FilterOperatorSchemaRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// FilterOperatorSchemaRequest is request for FilterOperatorSchema
type FilterOperatorSchemaRequest struct{ BaseRequest }

// Do performs HTTP request for FilterOperatorSchema
func (r *FilterOperatorSchemaRequest) Do(method, path string, reqObj interface{}) (resObj *FilterOperatorSchema, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for FilterOperatorSchema
func (r *FilterOperatorSchemaRequest) Get() (*FilterOperatorSchema, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for FilterOperatorSchema
func (r *FilterOperatorSchemaRequest) Update(reqObj *FilterOperatorSchema) (*FilterOperatorSchema, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for FilterOperatorSchema
func (r *FilterOperatorSchemaRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
