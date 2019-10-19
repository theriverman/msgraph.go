// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SalesQuoteLineRequestBuilder is request builder for SalesQuoteLine
type SalesQuoteLineRequestBuilder struct{ BaseRequestBuilder }

// Request returns SalesQuoteLineRequest
func (b *SalesQuoteLineRequestBuilder) Request() *SalesQuoteLineRequest {
	return &SalesQuoteLineRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SalesQuoteLineRequest is request for SalesQuoteLine
type SalesQuoteLineRequest struct{ BaseRequest }

// Do performs HTTP request for SalesQuoteLine
func (r *SalesQuoteLineRequest) Do(method, path string, reqObj interface{}) (resObj *SalesQuoteLine, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for SalesQuoteLine
func (r *SalesQuoteLineRequest) Get() (*SalesQuoteLine, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for SalesQuoteLine
func (r *SalesQuoteLineRequest) Update(reqObj *SalesQuoteLine) (*SalesQuoteLine, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for SalesQuoteLine
func (r *SalesQuoteLineRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Account is navigation property
func (b *SalesQuoteLineRequestBuilder) Account() *AccountRequestBuilder {
	bb := &AccountRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/account"
	return bb
}

// Item is navigation property
func (b *SalesQuoteLineRequestBuilder) Item() *ItemRequestBuilder {
	bb := &ItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/item"
	return bb
}
