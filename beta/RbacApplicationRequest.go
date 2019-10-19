// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// RbacApplicationRequestBuilder is request builder for RbacApplication
type RbacApplicationRequestBuilder struct{ BaseRequestBuilder }

// Request returns RbacApplicationRequest
func (b *RbacApplicationRequestBuilder) Request() *RbacApplicationRequest {
	return &RbacApplicationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RbacApplicationRequest is request for RbacApplication
type RbacApplicationRequest struct{ BaseRequest }

// Do performs HTTP request for RbacApplication
func (r *RbacApplicationRequest) Do(method, path string, reqObj interface{}) (resObj *RbacApplication, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for RbacApplication
func (r *RbacApplicationRequest) Get() (*RbacApplication, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for RbacApplication
func (r *RbacApplicationRequest) Update(reqObj *RbacApplication) (*RbacApplication, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for RbacApplication
func (r *RbacApplicationRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// RoleAssignments returns request builder for UnifiedRoleAssignment collection
func (b *RbacApplicationRequestBuilder) RoleAssignments() *RbacApplicationRoleAssignmentsCollectionRequestBuilder {
	bb := &RbacApplicationRoleAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleAssignments"
	return bb
}

// RbacApplicationRoleAssignmentsCollectionRequestBuilder is request builder for UnifiedRoleAssignment collection
type RbacApplicationRoleAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for UnifiedRoleAssignment collection
func (b *RbacApplicationRoleAssignmentsCollectionRequestBuilder) Request() *RbacApplicationRoleAssignmentsCollectionRequest {
	return &RbacApplicationRoleAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for UnifiedRoleAssignment item
func (b *RbacApplicationRoleAssignmentsCollectionRequestBuilder) ID(id string) *UnifiedRoleAssignmentRequestBuilder {
	bb := &UnifiedRoleAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// RbacApplicationRoleAssignmentsCollectionRequest is request for UnifiedRoleAssignment collection
type RbacApplicationRoleAssignmentsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for UnifiedRoleAssignment collection
func (r *RbacApplicationRoleAssignmentsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *UnifiedRoleAssignment, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for UnifiedRoleAssignment collection
func (r *RbacApplicationRoleAssignmentsCollectionRequest) Paging(method, path string, obj interface{}) ([]UnifiedRoleAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []UnifiedRoleAssignment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []UnifiedRoleAssignment
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

// Get performs GET request for UnifiedRoleAssignment collection
func (r *RbacApplicationRoleAssignmentsCollectionRequest) Get() ([]UnifiedRoleAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for UnifiedRoleAssignment collection
func (r *RbacApplicationRoleAssignmentsCollectionRequest) Add(reqObj *UnifiedRoleAssignment) (*UnifiedRoleAssignment, error) {
	return r.Do("POST", "", reqObj)
}

// RoleDefinitions returns request builder for UnifiedRoleDefinition collection
func (b *RbacApplicationRequestBuilder) RoleDefinitions() *RbacApplicationRoleDefinitionsCollectionRequestBuilder {
	bb := &RbacApplicationRoleDefinitionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleDefinitions"
	return bb
}

// RbacApplicationRoleDefinitionsCollectionRequestBuilder is request builder for UnifiedRoleDefinition collection
type RbacApplicationRoleDefinitionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for UnifiedRoleDefinition collection
func (b *RbacApplicationRoleDefinitionsCollectionRequestBuilder) Request() *RbacApplicationRoleDefinitionsCollectionRequest {
	return &RbacApplicationRoleDefinitionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for UnifiedRoleDefinition item
func (b *RbacApplicationRoleDefinitionsCollectionRequestBuilder) ID(id string) *UnifiedRoleDefinitionRequestBuilder {
	bb := &UnifiedRoleDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// RbacApplicationRoleDefinitionsCollectionRequest is request for UnifiedRoleDefinition collection
type RbacApplicationRoleDefinitionsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for UnifiedRoleDefinition collection
func (r *RbacApplicationRoleDefinitionsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *UnifiedRoleDefinition, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for UnifiedRoleDefinition collection
func (r *RbacApplicationRoleDefinitionsCollectionRequest) Paging(method, path string, obj interface{}) ([]UnifiedRoleDefinition, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []UnifiedRoleDefinition
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []UnifiedRoleDefinition
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

// Get performs GET request for UnifiedRoleDefinition collection
func (r *RbacApplicationRoleDefinitionsCollectionRequest) Get() ([]UnifiedRoleDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for UnifiedRoleDefinition collection
func (r *RbacApplicationRoleDefinitionsCollectionRequest) Add(reqObj *UnifiedRoleDefinition) (*UnifiedRoleDefinition, error) {
	return r.Do("POST", "", reqObj)
}
