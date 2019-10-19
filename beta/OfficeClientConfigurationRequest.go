// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// OfficeClientConfigurationRequestBuilder is request builder for OfficeClientConfiguration
type OfficeClientConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns OfficeClientConfigurationRequest
func (b *OfficeClientConfigurationRequestBuilder) Request() *OfficeClientConfigurationRequest {
	return &OfficeClientConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OfficeClientConfigurationRequest is request for OfficeClientConfiguration
type OfficeClientConfigurationRequest struct{ BaseRequest }

// Do performs HTTP request for OfficeClientConfiguration
func (r *OfficeClientConfigurationRequest) Do(method, path string, reqObj interface{}) (resObj *OfficeClientConfiguration, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for OfficeClientConfiguration
func (r *OfficeClientConfigurationRequest) Get() (*OfficeClientConfiguration, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for OfficeClientConfiguration
func (r *OfficeClientConfigurationRequest) Update(reqObj *OfficeClientConfiguration) (*OfficeClientConfiguration, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for OfficeClientConfiguration
func (r *OfficeClientConfigurationRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Assignments returns request builder for OfficeClientConfigurationAssignment collection
func (b *OfficeClientConfigurationRequestBuilder) Assignments() *OfficeClientConfigurationAssignmentsCollectionRequestBuilder {
	bb := &OfficeClientConfigurationAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// OfficeClientConfigurationAssignmentsCollectionRequestBuilder is request builder for OfficeClientConfigurationAssignment collection
type OfficeClientConfigurationAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OfficeClientConfigurationAssignment collection
func (b *OfficeClientConfigurationAssignmentsCollectionRequestBuilder) Request() *OfficeClientConfigurationAssignmentsCollectionRequest {
	return &OfficeClientConfigurationAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OfficeClientConfigurationAssignment item
func (b *OfficeClientConfigurationAssignmentsCollectionRequestBuilder) ID(id string) *OfficeClientConfigurationAssignmentRequestBuilder {
	bb := &OfficeClientConfigurationAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OfficeClientConfigurationAssignmentsCollectionRequest is request for OfficeClientConfigurationAssignment collection
type OfficeClientConfigurationAssignmentsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for OfficeClientConfigurationAssignment collection
func (r *OfficeClientConfigurationAssignmentsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *OfficeClientConfigurationAssignment, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for OfficeClientConfigurationAssignment collection
func (r *OfficeClientConfigurationAssignmentsCollectionRequest) Paging(method, path string, obj interface{}) ([]OfficeClientConfigurationAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []OfficeClientConfigurationAssignment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []OfficeClientConfigurationAssignment
		)
		err := json.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for OfficeClientConfigurationAssignment collection
func (r *OfficeClientConfigurationAssignmentsCollectionRequest) Get() ([]OfficeClientConfigurationAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for OfficeClientConfigurationAssignment collection
func (r *OfficeClientConfigurationAssignmentsCollectionRequest) Add(reqObj *OfficeClientConfigurationAssignment) (*OfficeClientConfigurationAssignment, error) {
	return r.Do("POST", "", reqObj)
}
