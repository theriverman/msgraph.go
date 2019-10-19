// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AccessReviewDecisionRequestBuilder is request builder for AccessReviewDecision
type AccessReviewDecisionRequestBuilder struct{ BaseRequestBuilder }

// Request returns AccessReviewDecisionRequest
func (b *AccessReviewDecisionRequestBuilder) Request() *AccessReviewDecisionRequest {
	return &AccessReviewDecisionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AccessReviewDecisionRequest is request for AccessReviewDecision
type AccessReviewDecisionRequest struct{ BaseRequest }

// Do performs HTTP request for AccessReviewDecision
func (r *AccessReviewDecisionRequest) Do(method, path string, reqObj interface{}) (resObj *AccessReviewDecision, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for AccessReviewDecision
func (r *AccessReviewDecisionRequest) Get() (*AccessReviewDecision, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for AccessReviewDecision
func (r *AccessReviewDecisionRequest) Update(reqObj *AccessReviewDecision) (*AccessReviewDecision, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for AccessReviewDecision
func (r *AccessReviewDecisionRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
