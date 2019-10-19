// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ProvisioningObjectSummaryRequestBuilder is request builder for ProvisioningObjectSummary
type ProvisioningObjectSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns ProvisioningObjectSummaryRequest
func (b *ProvisioningObjectSummaryRequestBuilder) Request() *ProvisioningObjectSummaryRequest {
	return &ProvisioningObjectSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ProvisioningObjectSummaryRequest is request for ProvisioningObjectSummary
type ProvisioningObjectSummaryRequest struct{ BaseRequest }

// Do performs HTTP request for ProvisioningObjectSummary
func (r *ProvisioningObjectSummaryRequest) Do(method, path string, reqObj interface{}) (resObj *ProvisioningObjectSummary, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for ProvisioningObjectSummary
func (r *ProvisioningObjectSummaryRequest) Get() (*ProvisioningObjectSummary, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for ProvisioningObjectSummary
func (r *ProvisioningObjectSummaryRequest) Update(reqObj *ProvisioningObjectSummary) (*ProvisioningObjectSummary, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for ProvisioningObjectSummary
func (r *ProvisioningObjectSummaryRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
