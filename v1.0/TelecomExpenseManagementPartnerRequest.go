// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// TelecomExpenseManagementPartnerRequestBuilder is request builder for TelecomExpenseManagementPartner
type TelecomExpenseManagementPartnerRequestBuilder struct{ BaseRequestBuilder }

// Request returns TelecomExpenseManagementPartnerRequest
func (b *TelecomExpenseManagementPartnerRequestBuilder) Request() *TelecomExpenseManagementPartnerRequest {
	return &TelecomExpenseManagementPartnerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TelecomExpenseManagementPartnerRequest is request for TelecomExpenseManagementPartner
type TelecomExpenseManagementPartnerRequest struct{ BaseRequest }

// Do performs HTTP request for TelecomExpenseManagementPartner
func (r *TelecomExpenseManagementPartnerRequest) Do(method, path string, reqObj interface{}) (resObj *TelecomExpenseManagementPartner, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for TelecomExpenseManagementPartner
func (r *TelecomExpenseManagementPartnerRequest) Get() (*TelecomExpenseManagementPartner, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for TelecomExpenseManagementPartner
func (r *TelecomExpenseManagementPartnerRequest) Update(reqObj *TelecomExpenseManagementPartner) (*TelecomExpenseManagementPartner, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for TelecomExpenseManagementPartner
func (r *TelecomExpenseManagementPartnerRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
