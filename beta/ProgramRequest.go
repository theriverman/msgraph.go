// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ProgramRequestBuilder is request builder for Program
type ProgramRequestBuilder struct{ BaseRequestBuilder }

// Request returns ProgramRequest
func (b *ProgramRequestBuilder) Request() *ProgramRequest {
	return &ProgramRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ProgramRequest is request for Program
type ProgramRequest struct{ BaseRequest }

// Do performs HTTP request for Program
func (r *ProgramRequest) Do(method, path string, reqObj interface{}) (resObj *Program, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for Program
func (r *ProgramRequest) Get() (*Program, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for Program
func (r *ProgramRequest) Update(reqObj *Program) (*Program, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for Program
func (r *ProgramRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Controls returns request builder for ProgramControl collection
func (b *ProgramRequestBuilder) Controls() *ProgramControlsCollectionRequestBuilder {
	bb := &ProgramControlsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/controls"
	return bb
}

// ProgramControlsCollectionRequestBuilder is request builder for ProgramControl collection
type ProgramControlsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ProgramControl collection
func (b *ProgramControlsCollectionRequestBuilder) Request() *ProgramControlsCollectionRequest {
	return &ProgramControlsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ProgramControl item
func (b *ProgramControlsCollectionRequestBuilder) ID(id string) *ProgramControlRequestBuilder {
	bb := &ProgramControlRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ProgramControlsCollectionRequest is request for ProgramControl collection
type ProgramControlsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for ProgramControl collection
func (r *ProgramControlsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *ProgramControl, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for ProgramControl collection
func (r *ProgramControlsCollectionRequest) Paging(method, path string, obj interface{}) ([]ProgramControl, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ProgramControl
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []ProgramControl
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

// Get performs GET request for ProgramControl collection
func (r *ProgramControlsCollectionRequest) Get() ([]ProgramControl, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ProgramControl collection
func (r *ProgramControlsCollectionRequest) Add(reqObj *ProgramControl) (*ProgramControl, error) {
	return r.Do("POST", "", reqObj)
}
