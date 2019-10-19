// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DataPolicyOperationRequestBuilder is request builder for DataPolicyOperation
type DataPolicyOperationRequestBuilder struct{ BaseRequestBuilder }

// Request returns DataPolicyOperationRequest
func (b *DataPolicyOperationRequestBuilder) Request() *DataPolicyOperationRequest {
	return &DataPolicyOperationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DataPolicyOperationRequest is request for DataPolicyOperation
type DataPolicyOperationRequest struct{ BaseRequest }

// Do performs HTTP request for DataPolicyOperation
func (r *DataPolicyOperationRequest) Do(method, path string, reqObj interface{}) (resObj *DataPolicyOperation, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for DataPolicyOperation
func (r *DataPolicyOperationRequest) Get() (*DataPolicyOperation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for DataPolicyOperation
func (r *DataPolicyOperationRequest) Update(reqObj *DataPolicyOperation) (*DataPolicyOperation, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for DataPolicyOperation
func (r *DataPolicyOperationRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
