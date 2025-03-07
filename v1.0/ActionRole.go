// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/theriverman/msgraph.go/jsonx"
)

// RoleDefinition is navigation property
func (b *RoleAssignmentRequestBuilder) RoleDefinition() *RoleDefinitionRequestBuilder {
	bb := &RoleDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleDefinition"
	return bb
}

// RoleAssignments returns request builder for RoleAssignment collection
func (b *RoleDefinitionRequestBuilder) RoleAssignments() *RoleDefinitionRoleAssignmentsCollectionRequestBuilder {
	bb := &RoleDefinitionRoleAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleAssignments"
	return bb
}

// RoleDefinitionRoleAssignmentsCollectionRequestBuilder is request builder for RoleAssignment collection
type RoleDefinitionRoleAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for RoleAssignment collection
func (b *RoleDefinitionRoleAssignmentsCollectionRequestBuilder) Request() *RoleDefinitionRoleAssignmentsCollectionRequest {
	return &RoleDefinitionRoleAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for RoleAssignment item
func (b *RoleDefinitionRoleAssignmentsCollectionRequestBuilder) ID(id string) *RoleAssignmentRequestBuilder {
	bb := &RoleAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// RoleDefinitionRoleAssignmentsCollectionRequest is request for RoleAssignment collection
type RoleDefinitionRoleAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for RoleAssignment collection
func (r *RoleDefinitionRoleAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]RoleAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []RoleAssignment
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []RoleAssignment
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for RoleAssignment collection, max N pages
func (r *RoleDefinitionRoleAssignmentsCollectionRequest) GetN(ctx context.Context, n int) ([]RoleAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for RoleAssignment collection
func (r *RoleDefinitionRoleAssignmentsCollectionRequest) Get(ctx context.Context) ([]RoleAssignment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for RoleAssignment collection
func (r *RoleDefinitionRoleAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *RoleAssignment) (resObj *RoleAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
