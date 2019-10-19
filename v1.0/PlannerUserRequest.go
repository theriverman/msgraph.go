// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// PlannerUserRequestBuilder is request builder for PlannerUser
type PlannerUserRequestBuilder struct{ BaseRequestBuilder }

// Request returns PlannerUserRequest
func (b *PlannerUserRequestBuilder) Request() *PlannerUserRequest {
	return &PlannerUserRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PlannerUserRequest is request for PlannerUser
type PlannerUserRequest struct{ BaseRequest }

// Do performs HTTP request for PlannerUser
func (r *PlannerUserRequest) Do(method, path string, reqObj interface{}) (resObj *PlannerUser, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for PlannerUser
func (r *PlannerUserRequest) Get() (*PlannerUser, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for PlannerUser
func (r *PlannerUserRequest) Update(reqObj *PlannerUser) (*PlannerUser, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for PlannerUser
func (r *PlannerUserRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Plans returns request builder for PlannerPlan collection
func (b *PlannerUserRequestBuilder) Plans() *PlannerUserPlansCollectionRequestBuilder {
	bb := &PlannerUserPlansCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/plans"
	return bb
}

// PlannerUserPlansCollectionRequestBuilder is request builder for PlannerPlan collection
type PlannerUserPlansCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PlannerPlan collection
func (b *PlannerUserPlansCollectionRequestBuilder) Request() *PlannerUserPlansCollectionRequest {
	return &PlannerUserPlansCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PlannerPlan item
func (b *PlannerUserPlansCollectionRequestBuilder) ID(id string) *PlannerPlanRequestBuilder {
	bb := &PlannerPlanRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PlannerUserPlansCollectionRequest is request for PlannerPlan collection
type PlannerUserPlansCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for PlannerPlan collection
func (r *PlannerUserPlansCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *PlannerPlan, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for PlannerPlan collection
func (r *PlannerUserPlansCollectionRequest) Paging(method, path string, obj interface{}) ([]PlannerPlan, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []PlannerPlan
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []PlannerPlan
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

// Get performs GET request for PlannerPlan collection
func (r *PlannerUserPlansCollectionRequest) Get() ([]PlannerPlan, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for PlannerPlan collection
func (r *PlannerUserPlansCollectionRequest) Add(reqObj *PlannerPlan) (*PlannerPlan, error) {
	return r.Do("POST", "", reqObj)
}

// Tasks returns request builder for PlannerTask collection
func (b *PlannerUserRequestBuilder) Tasks() *PlannerUserTasksCollectionRequestBuilder {
	bb := &PlannerUserTasksCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/tasks"
	return bb
}

// PlannerUserTasksCollectionRequestBuilder is request builder for PlannerTask collection
type PlannerUserTasksCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PlannerTask collection
func (b *PlannerUserTasksCollectionRequestBuilder) Request() *PlannerUserTasksCollectionRequest {
	return &PlannerUserTasksCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PlannerTask item
func (b *PlannerUserTasksCollectionRequestBuilder) ID(id string) *PlannerTaskRequestBuilder {
	bb := &PlannerTaskRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PlannerUserTasksCollectionRequest is request for PlannerTask collection
type PlannerUserTasksCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for PlannerTask collection
func (r *PlannerUserTasksCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *PlannerTask, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for PlannerTask collection
func (r *PlannerUserTasksCollectionRequest) Paging(method, path string, obj interface{}) ([]PlannerTask, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []PlannerTask
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []PlannerTask
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

// Get performs GET request for PlannerTask collection
func (r *PlannerUserTasksCollectionRequest) Get() ([]PlannerTask, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for PlannerTask collection
func (r *PlannerUserTasksCollectionRequest) Add(reqObj *PlannerTask) (*PlannerTask, error) {
	return r.Do("POST", "", reqObj)
}
