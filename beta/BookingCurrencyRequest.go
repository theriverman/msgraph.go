// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// BookingCurrencyRequestBuilder is request builder for BookingCurrency
type BookingCurrencyRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingCurrencyRequest
func (b *BookingCurrencyRequestBuilder) Request() *BookingCurrencyRequest {
	return &BookingCurrencyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingCurrencyRequest is request for BookingCurrency
type BookingCurrencyRequest struct{ BaseRequest }

// Do performs HTTP request for BookingCurrency
func (r *BookingCurrencyRequest) Do(method, path string, reqObj interface{}) (resObj *BookingCurrency, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for BookingCurrency
func (r *BookingCurrencyRequest) Get() (*BookingCurrency, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for BookingCurrency
func (r *BookingCurrencyRequest) Update(reqObj *BookingCurrency) (*BookingCurrency, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for BookingCurrency
func (r *BookingCurrencyRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
