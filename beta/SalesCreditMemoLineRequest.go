// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SalesCreditMemoLineRequestBuilder is request builder for SalesCreditMemoLine
type SalesCreditMemoLineRequestBuilder struct{ BaseRequestBuilder }

// Request returns SalesCreditMemoLineRequest
func (b *SalesCreditMemoLineRequestBuilder) Request() *SalesCreditMemoLineRequest {
	return &SalesCreditMemoLineRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SalesCreditMemoLineRequest is request for SalesCreditMemoLine
type SalesCreditMemoLineRequest struct{ BaseRequest }

// Do performs HTTP request for SalesCreditMemoLine
func (r *SalesCreditMemoLineRequest) Do(method, path string, reqObj interface{}) (resObj *SalesCreditMemoLine, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for SalesCreditMemoLine
func (r *SalesCreditMemoLineRequest) Get() (*SalesCreditMemoLine, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for SalesCreditMemoLine
func (r *SalesCreditMemoLineRequest) Update(reqObj *SalesCreditMemoLine) (*SalesCreditMemoLine, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for SalesCreditMemoLine
func (r *SalesCreditMemoLineRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Account is navigation property
func (b *SalesCreditMemoLineRequestBuilder) Account() *AccountRequestBuilder {
	bb := &AccountRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/account"
	return bb
}

// Item is navigation property
func (b *SalesCreditMemoLineRequestBuilder) Item() *ItemRequestBuilder {
	bb := &ItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/item"
	return bb
}
