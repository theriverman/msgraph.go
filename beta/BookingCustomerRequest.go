// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// BookingCustomerRequestBuilder is request builder for BookingCustomer
type BookingCustomerRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingCustomerRequest
func (b *BookingCustomerRequestBuilder) Request() *BookingCustomerRequest {
	return &BookingCustomerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingCustomerRequest is request for BookingCustomer
type BookingCustomerRequest struct{ BaseRequest }

// Do performs HTTP request for BookingCustomer
func (r *BookingCustomerRequest) Do(method, path string, reqObj interface{}) (resObj *BookingCustomer, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for BookingCustomer
func (r *BookingCustomerRequest) Get() (*BookingCustomer, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for BookingCustomer
func (r *BookingCustomerRequest) Update(reqObj *BookingCustomer) (*BookingCustomer, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for BookingCustomer
func (r *BookingCustomerRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
