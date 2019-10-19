// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// MobileAppRequestBuilder is request builder for MobileApp
type MobileAppRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobileAppRequest
func (b *MobileAppRequestBuilder) Request() *MobileAppRequest {
	return &MobileAppRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobileAppRequest is request for MobileApp
type MobileAppRequest struct{ BaseRequest }

// Do performs HTTP request for MobileApp
func (r *MobileAppRequest) Do(method, path string, reqObj interface{}) (resObj *MobileApp, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for MobileApp
func (r *MobileAppRequest) Get() (*MobileApp, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for MobileApp
func (r *MobileAppRequest) Update(reqObj *MobileApp) (*MobileApp, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for MobileApp
func (r *MobileAppRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Assignments returns request builder for MobileAppAssignment collection
func (b *MobileAppRequestBuilder) Assignments() *MobileAppAssignmentsCollectionRequestBuilder {
	bb := &MobileAppAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// MobileAppAssignmentsCollectionRequestBuilder is request builder for MobileAppAssignment collection
type MobileAppAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MobileAppAssignment collection
func (b *MobileAppAssignmentsCollectionRequestBuilder) Request() *MobileAppAssignmentsCollectionRequest {
	return &MobileAppAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MobileAppAssignment item
func (b *MobileAppAssignmentsCollectionRequestBuilder) ID(id string) *MobileAppAssignmentRequestBuilder {
	bb := &MobileAppAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MobileAppAssignmentsCollectionRequest is request for MobileAppAssignment collection
type MobileAppAssignmentsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for MobileAppAssignment collection
func (r *MobileAppAssignmentsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *MobileAppAssignment, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for MobileAppAssignment collection
func (r *MobileAppAssignmentsCollectionRequest) Paging(method, path string, obj interface{}) ([]MobileAppAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []MobileAppAssignment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []MobileAppAssignment
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

// Get performs GET request for MobileAppAssignment collection
func (r *MobileAppAssignmentsCollectionRequest) Get() ([]MobileAppAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for MobileAppAssignment collection
func (r *MobileAppAssignmentsCollectionRequest) Add(reqObj *MobileAppAssignment) (*MobileAppAssignment, error) {
	return r.Do("POST", "", reqObj)
}

// Categories returns request builder for MobileAppCategory collection
func (b *MobileAppRequestBuilder) Categories() *MobileAppCategoriesCollectionRequestBuilder {
	bb := &MobileAppCategoriesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/categories"
	return bb
}

// MobileAppCategoriesCollectionRequestBuilder is request builder for MobileAppCategory collection
type MobileAppCategoriesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MobileAppCategory collection
func (b *MobileAppCategoriesCollectionRequestBuilder) Request() *MobileAppCategoriesCollectionRequest {
	return &MobileAppCategoriesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MobileAppCategory item
func (b *MobileAppCategoriesCollectionRequestBuilder) ID(id string) *MobileAppCategoryRequestBuilder {
	bb := &MobileAppCategoryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MobileAppCategoriesCollectionRequest is request for MobileAppCategory collection
type MobileAppCategoriesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for MobileAppCategory collection
func (r *MobileAppCategoriesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *MobileAppCategory, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for MobileAppCategory collection
func (r *MobileAppCategoriesCollectionRequest) Paging(method, path string, obj interface{}) ([]MobileAppCategory, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []MobileAppCategory
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []MobileAppCategory
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

// Get performs GET request for MobileAppCategory collection
func (r *MobileAppCategoriesCollectionRequest) Get() ([]MobileAppCategory, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for MobileAppCategory collection
func (r *MobileAppCategoriesCollectionRequest) Add(reqObj *MobileAppCategory) (*MobileAppCategory, error) {
	return r.Do("POST", "", reqObj)
}

// DeviceStatuses returns request builder for MobileAppInstallStatus collection
func (b *MobileAppRequestBuilder) DeviceStatuses() *MobileAppDeviceStatusesCollectionRequestBuilder {
	bb := &MobileAppDeviceStatusesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceStatuses"
	return bb
}

// MobileAppDeviceStatusesCollectionRequestBuilder is request builder for MobileAppInstallStatus collection
type MobileAppDeviceStatusesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MobileAppInstallStatus collection
func (b *MobileAppDeviceStatusesCollectionRequestBuilder) Request() *MobileAppDeviceStatusesCollectionRequest {
	return &MobileAppDeviceStatusesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MobileAppInstallStatus item
func (b *MobileAppDeviceStatusesCollectionRequestBuilder) ID(id string) *MobileAppInstallStatusRequestBuilder {
	bb := &MobileAppInstallStatusRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MobileAppDeviceStatusesCollectionRequest is request for MobileAppInstallStatus collection
type MobileAppDeviceStatusesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for MobileAppInstallStatus collection
func (r *MobileAppDeviceStatusesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *MobileAppInstallStatus, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for MobileAppInstallStatus collection
func (r *MobileAppDeviceStatusesCollectionRequest) Paging(method, path string, obj interface{}) ([]MobileAppInstallStatus, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []MobileAppInstallStatus
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []MobileAppInstallStatus
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

// Get performs GET request for MobileAppInstallStatus collection
func (r *MobileAppDeviceStatusesCollectionRequest) Get() ([]MobileAppInstallStatus, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for MobileAppInstallStatus collection
func (r *MobileAppDeviceStatusesCollectionRequest) Add(reqObj *MobileAppInstallStatus) (*MobileAppInstallStatus, error) {
	return r.Do("POST", "", reqObj)
}

// InstallSummary is navigation property
func (b *MobileAppRequestBuilder) InstallSummary() *MobileAppInstallSummaryRequestBuilder {
	bb := &MobileAppInstallSummaryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/installSummary"
	return bb
}

// Relationships returns request builder for MobileAppRelationship collection
func (b *MobileAppRequestBuilder) Relationships() *MobileAppRelationshipsCollectionRequestBuilder {
	bb := &MobileAppRelationshipsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/relationships"
	return bb
}

// MobileAppRelationshipsCollectionRequestBuilder is request builder for MobileAppRelationship collection
type MobileAppRelationshipsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MobileAppRelationship collection
func (b *MobileAppRelationshipsCollectionRequestBuilder) Request() *MobileAppRelationshipsCollectionRequest {
	return &MobileAppRelationshipsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MobileAppRelationship item
func (b *MobileAppRelationshipsCollectionRequestBuilder) ID(id string) *MobileAppRelationshipRequestBuilder {
	bb := &MobileAppRelationshipRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MobileAppRelationshipsCollectionRequest is request for MobileAppRelationship collection
type MobileAppRelationshipsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for MobileAppRelationship collection
func (r *MobileAppRelationshipsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *MobileAppRelationship, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for MobileAppRelationship collection
func (r *MobileAppRelationshipsCollectionRequest) Paging(method, path string, obj interface{}) ([]MobileAppRelationship, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []MobileAppRelationship
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []MobileAppRelationship
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

// Get performs GET request for MobileAppRelationship collection
func (r *MobileAppRelationshipsCollectionRequest) Get() ([]MobileAppRelationship, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for MobileAppRelationship collection
func (r *MobileAppRelationshipsCollectionRequest) Add(reqObj *MobileAppRelationship) (*MobileAppRelationship, error) {
	return r.Do("POST", "", reqObj)
}

// UserStatuses returns request builder for UserAppInstallStatus collection
func (b *MobileAppRequestBuilder) UserStatuses() *MobileAppUserStatusesCollectionRequestBuilder {
	bb := &MobileAppUserStatusesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/userStatuses"
	return bb
}

// MobileAppUserStatusesCollectionRequestBuilder is request builder for UserAppInstallStatus collection
type MobileAppUserStatusesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for UserAppInstallStatus collection
func (b *MobileAppUserStatusesCollectionRequestBuilder) Request() *MobileAppUserStatusesCollectionRequest {
	return &MobileAppUserStatusesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for UserAppInstallStatus item
func (b *MobileAppUserStatusesCollectionRequestBuilder) ID(id string) *UserAppInstallStatusRequestBuilder {
	bb := &UserAppInstallStatusRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MobileAppUserStatusesCollectionRequest is request for UserAppInstallStatus collection
type MobileAppUserStatusesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for UserAppInstallStatus collection
func (r *MobileAppUserStatusesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *UserAppInstallStatus, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for UserAppInstallStatus collection
func (r *MobileAppUserStatusesCollectionRequest) Paging(method, path string, obj interface{}) ([]UserAppInstallStatus, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []UserAppInstallStatus
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []UserAppInstallStatus
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

// Get performs GET request for UserAppInstallStatus collection
func (r *MobileAppUserStatusesCollectionRequest) Get() ([]UserAppInstallStatus, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for UserAppInstallStatus collection
func (r *MobileAppUserStatusesCollectionRequest) Add(reqObj *UserAppInstallStatus) (*UserAppInstallStatus, error) {
	return r.Do("POST", "", reqObj)
}
