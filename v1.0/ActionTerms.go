// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/theriverman/msgraph.go/jsonx"
)

// AcceptanceStatuses returns request builder for TermsAndConditionsAcceptanceStatus collection
func (b *TermsAndConditionsRequestBuilder) AcceptanceStatuses() *TermsAndConditionsAcceptanceStatusesCollectionRequestBuilder {
	bb := &TermsAndConditionsAcceptanceStatusesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/acceptanceStatuses"
	return bb
}

// TermsAndConditionsAcceptanceStatusesCollectionRequestBuilder is request builder for TermsAndConditionsAcceptanceStatus collection
type TermsAndConditionsAcceptanceStatusesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TermsAndConditionsAcceptanceStatus collection
func (b *TermsAndConditionsAcceptanceStatusesCollectionRequestBuilder) Request() *TermsAndConditionsAcceptanceStatusesCollectionRequest {
	return &TermsAndConditionsAcceptanceStatusesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TermsAndConditionsAcceptanceStatus item
func (b *TermsAndConditionsAcceptanceStatusesCollectionRequestBuilder) ID(id string) *TermsAndConditionsAcceptanceStatusRequestBuilder {
	bb := &TermsAndConditionsAcceptanceStatusRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TermsAndConditionsAcceptanceStatusesCollectionRequest is request for TermsAndConditionsAcceptanceStatus collection
type TermsAndConditionsAcceptanceStatusesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TermsAndConditionsAcceptanceStatus collection
func (r *TermsAndConditionsAcceptanceStatusesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TermsAndConditionsAcceptanceStatus, error) {
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
	var values []TermsAndConditionsAcceptanceStatus
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
			value  []TermsAndConditionsAcceptanceStatus
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

// GetN performs GET request for TermsAndConditionsAcceptanceStatus collection, max N pages
func (r *TermsAndConditionsAcceptanceStatusesCollectionRequest) GetN(ctx context.Context, n int) ([]TermsAndConditionsAcceptanceStatus, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TermsAndConditionsAcceptanceStatus collection
func (r *TermsAndConditionsAcceptanceStatusesCollectionRequest) Get(ctx context.Context) ([]TermsAndConditionsAcceptanceStatus, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TermsAndConditionsAcceptanceStatus collection
func (r *TermsAndConditionsAcceptanceStatusesCollectionRequest) Add(ctx context.Context, reqObj *TermsAndConditionsAcceptanceStatus) (resObj *TermsAndConditionsAcceptanceStatus, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Assignments returns request builder for TermsAndConditionsAssignment collection
func (b *TermsAndConditionsRequestBuilder) Assignments() *TermsAndConditionsAssignmentsCollectionRequestBuilder {
	bb := &TermsAndConditionsAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// TermsAndConditionsAssignmentsCollectionRequestBuilder is request builder for TermsAndConditionsAssignment collection
type TermsAndConditionsAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TermsAndConditionsAssignment collection
func (b *TermsAndConditionsAssignmentsCollectionRequestBuilder) Request() *TermsAndConditionsAssignmentsCollectionRequest {
	return &TermsAndConditionsAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TermsAndConditionsAssignment item
func (b *TermsAndConditionsAssignmentsCollectionRequestBuilder) ID(id string) *TermsAndConditionsAssignmentRequestBuilder {
	bb := &TermsAndConditionsAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TermsAndConditionsAssignmentsCollectionRequest is request for TermsAndConditionsAssignment collection
type TermsAndConditionsAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TermsAndConditionsAssignment collection
func (r *TermsAndConditionsAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TermsAndConditionsAssignment, error) {
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
	var values []TermsAndConditionsAssignment
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
			value  []TermsAndConditionsAssignment
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

// GetN performs GET request for TermsAndConditionsAssignment collection, max N pages
func (r *TermsAndConditionsAssignmentsCollectionRequest) GetN(ctx context.Context, n int) ([]TermsAndConditionsAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TermsAndConditionsAssignment collection
func (r *TermsAndConditionsAssignmentsCollectionRequest) Get(ctx context.Context) ([]TermsAndConditionsAssignment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TermsAndConditionsAssignment collection
func (r *TermsAndConditionsAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *TermsAndConditionsAssignment) (resObj *TermsAndConditionsAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// TermsAndConditions is navigation property
func (b *TermsAndConditionsAcceptanceStatusRequestBuilder) TermsAndConditions() *TermsAndConditionsRequestBuilder {
	bb := &TermsAndConditionsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/termsAndConditions"
	return bb
}
