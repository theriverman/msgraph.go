// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// EducationAssignmentRequestBuilder is request builder for EducationAssignment
type EducationAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns EducationAssignmentRequest
func (b *EducationAssignmentRequestBuilder) Request() *EducationAssignmentRequest {
	return &EducationAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EducationAssignmentRequest is request for EducationAssignment
type EducationAssignmentRequest struct{ BaseRequest }

// Do performs HTTP request for EducationAssignment
func (r *EducationAssignmentRequest) Do(method, path string, reqObj interface{}) (resObj *EducationAssignment, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for EducationAssignment
func (r *EducationAssignmentRequest) Get() (*EducationAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for EducationAssignment
func (r *EducationAssignmentRequest) Update(reqObj *EducationAssignment) (*EducationAssignment, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for EducationAssignment
func (r *EducationAssignmentRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Categories returns request builder for EducationCategory collection
func (b *EducationAssignmentRequestBuilder) Categories() *EducationAssignmentCategoriesCollectionRequestBuilder {
	bb := &EducationAssignmentCategoriesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/categories"
	return bb
}

// EducationAssignmentCategoriesCollectionRequestBuilder is request builder for EducationCategory collection
type EducationAssignmentCategoriesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationCategory collection
func (b *EducationAssignmentCategoriesCollectionRequestBuilder) Request() *EducationAssignmentCategoriesCollectionRequest {
	return &EducationAssignmentCategoriesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationCategory item
func (b *EducationAssignmentCategoriesCollectionRequestBuilder) ID(id string) *EducationCategoryRequestBuilder {
	bb := &EducationCategoryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationAssignmentCategoriesCollectionRequest is request for EducationCategory collection
type EducationAssignmentCategoriesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for EducationCategory collection
func (r *EducationAssignmentCategoriesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *EducationCategory, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for EducationCategory collection
func (r *EducationAssignmentCategoriesCollectionRequest) Paging(method, path string, obj interface{}) ([]EducationCategory, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []EducationCategory
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []EducationCategory
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

// Get performs GET request for EducationCategory collection
func (r *EducationAssignmentCategoriesCollectionRequest) Get() ([]EducationCategory, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for EducationCategory collection
func (r *EducationAssignmentCategoriesCollectionRequest) Add(reqObj *EducationCategory) (*EducationCategory, error) {
	return r.Do("POST", "", reqObj)
}

// Resources returns request builder for EducationAssignmentResource collection
func (b *EducationAssignmentRequestBuilder) Resources() *EducationAssignmentResourcesCollectionRequestBuilder {
	bb := &EducationAssignmentResourcesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/resources"
	return bb
}

// EducationAssignmentResourcesCollectionRequestBuilder is request builder for EducationAssignmentResource collection
type EducationAssignmentResourcesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationAssignmentResource collection
func (b *EducationAssignmentResourcesCollectionRequestBuilder) Request() *EducationAssignmentResourcesCollectionRequest {
	return &EducationAssignmentResourcesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationAssignmentResource item
func (b *EducationAssignmentResourcesCollectionRequestBuilder) ID(id string) *EducationAssignmentResourceRequestBuilder {
	bb := &EducationAssignmentResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationAssignmentResourcesCollectionRequest is request for EducationAssignmentResource collection
type EducationAssignmentResourcesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for EducationAssignmentResource collection
func (r *EducationAssignmentResourcesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *EducationAssignmentResource, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for EducationAssignmentResource collection
func (r *EducationAssignmentResourcesCollectionRequest) Paging(method, path string, obj interface{}) ([]EducationAssignmentResource, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []EducationAssignmentResource
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []EducationAssignmentResource
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

// Get performs GET request for EducationAssignmentResource collection
func (r *EducationAssignmentResourcesCollectionRequest) Get() ([]EducationAssignmentResource, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for EducationAssignmentResource collection
func (r *EducationAssignmentResourcesCollectionRequest) Add(reqObj *EducationAssignmentResource) (*EducationAssignmentResource, error) {
	return r.Do("POST", "", reqObj)
}

// Rubric is navigation property
func (b *EducationAssignmentRequestBuilder) Rubric() *EducationRubricRequestBuilder {
	bb := &EducationRubricRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rubric"
	return bb
}

// Submissions returns request builder for EducationSubmission collection
func (b *EducationAssignmentRequestBuilder) Submissions() *EducationAssignmentSubmissionsCollectionRequestBuilder {
	bb := &EducationAssignmentSubmissionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/submissions"
	return bb
}

// EducationAssignmentSubmissionsCollectionRequestBuilder is request builder for EducationSubmission collection
type EducationAssignmentSubmissionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationSubmission collection
func (b *EducationAssignmentSubmissionsCollectionRequestBuilder) Request() *EducationAssignmentSubmissionsCollectionRequest {
	return &EducationAssignmentSubmissionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationSubmission item
func (b *EducationAssignmentSubmissionsCollectionRequestBuilder) ID(id string) *EducationSubmissionRequestBuilder {
	bb := &EducationSubmissionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationAssignmentSubmissionsCollectionRequest is request for EducationSubmission collection
type EducationAssignmentSubmissionsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for EducationSubmission collection
func (r *EducationAssignmentSubmissionsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *EducationSubmission, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for EducationSubmission collection
func (r *EducationAssignmentSubmissionsCollectionRequest) Paging(method, path string, obj interface{}) ([]EducationSubmission, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []EducationSubmission
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []EducationSubmission
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

// Get performs GET request for EducationSubmission collection
func (r *EducationAssignmentSubmissionsCollectionRequest) Get() ([]EducationSubmission, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for EducationSubmission collection
func (r *EducationAssignmentSubmissionsCollectionRequest) Add(reqObj *EducationSubmission) (*EducationSubmission, error) {
	return r.Do("POST", "", reqObj)
}
