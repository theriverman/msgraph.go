// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ParticipantRequestBuilder is request builder for Participant
type ParticipantRequestBuilder struct{ BaseRequestBuilder }

// Request returns ParticipantRequest
func (b *ParticipantRequestBuilder) Request() *ParticipantRequest {
	return &ParticipantRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ParticipantRequest is request for Participant
type ParticipantRequest struct{ BaseRequest }

// Do performs HTTP request for Participant
func (r *ParticipantRequest) Do(method, path string, reqObj interface{}) (resObj *Participant, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for Participant
func (r *ParticipantRequest) Get() (*Participant, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for Participant
func (r *ParticipantRequest) Update(reqObj *Participant) (*Participant, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for Participant
func (r *ParticipantRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
