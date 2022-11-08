// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/theriverman/msgraph.go/jsonx"
)

// TargetedManagedAppConfigurationAssignRequestParameter undocumented
type TargetedManagedAppConfigurationAssignRequestParameter struct {
	// Assignments undocumented
	Assignments []TargetedManagedAppPolicyAssignment `json:"assignments,omitempty"`
}

// TargetedManagedAppConfigurationTargetAppsRequestParameter undocumented
type TargetedManagedAppConfigurationTargetAppsRequestParameter struct {
	// Apps undocumented
	Apps []ManagedMobileApp `json:"apps,omitempty"`
}

// TargetedManagedAppProtectionAssignRequestParameter undocumented
type TargetedManagedAppProtectionAssignRequestParameter struct {
	// Assignments undocumented
	Assignments []TargetedManagedAppPolicyAssignment `json:"assignments,omitempty"`
}

// Apps returns request builder for ManagedMobileApp collection
func (b *TargetedManagedAppConfigurationRequestBuilder) Apps() *TargetedManagedAppConfigurationAppsCollectionRequestBuilder {
	bb := &TargetedManagedAppConfigurationAppsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/apps"
	return bb
}

// TargetedManagedAppConfigurationAppsCollectionRequestBuilder is request builder for ManagedMobileApp collection
type TargetedManagedAppConfigurationAppsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagedMobileApp collection
func (b *TargetedManagedAppConfigurationAppsCollectionRequestBuilder) Request() *TargetedManagedAppConfigurationAppsCollectionRequest {
	return &TargetedManagedAppConfigurationAppsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagedMobileApp item
func (b *TargetedManagedAppConfigurationAppsCollectionRequestBuilder) ID(id string) *ManagedMobileAppRequestBuilder {
	bb := &ManagedMobileAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TargetedManagedAppConfigurationAppsCollectionRequest is request for ManagedMobileApp collection
type TargetedManagedAppConfigurationAppsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ManagedMobileApp collection
func (r *TargetedManagedAppConfigurationAppsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ManagedMobileApp, error) {
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
	var values []ManagedMobileApp
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
			value  []ManagedMobileApp
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

// GetN performs GET request for ManagedMobileApp collection, max N pages
func (r *TargetedManagedAppConfigurationAppsCollectionRequest) GetN(ctx context.Context, n int) ([]ManagedMobileApp, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ManagedMobileApp collection
func (r *TargetedManagedAppConfigurationAppsCollectionRequest) Get(ctx context.Context) ([]ManagedMobileApp, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ManagedMobileApp collection
func (r *TargetedManagedAppConfigurationAppsCollectionRequest) Add(ctx context.Context, reqObj *ManagedMobileApp) (resObj *ManagedMobileApp, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Assignments returns request builder for TargetedManagedAppPolicyAssignment collection
func (b *TargetedManagedAppConfigurationRequestBuilder) Assignments() *TargetedManagedAppConfigurationAssignmentsCollectionRequestBuilder {
	bb := &TargetedManagedAppConfigurationAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// TargetedManagedAppConfigurationAssignmentsCollectionRequestBuilder is request builder for TargetedManagedAppPolicyAssignment collection
type TargetedManagedAppConfigurationAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TargetedManagedAppPolicyAssignment collection
func (b *TargetedManagedAppConfigurationAssignmentsCollectionRequestBuilder) Request() *TargetedManagedAppConfigurationAssignmentsCollectionRequest {
	return &TargetedManagedAppConfigurationAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TargetedManagedAppPolicyAssignment item
func (b *TargetedManagedAppConfigurationAssignmentsCollectionRequestBuilder) ID(id string) *TargetedManagedAppPolicyAssignmentRequestBuilder {
	bb := &TargetedManagedAppPolicyAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TargetedManagedAppConfigurationAssignmentsCollectionRequest is request for TargetedManagedAppPolicyAssignment collection
type TargetedManagedAppConfigurationAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TargetedManagedAppPolicyAssignment collection
func (r *TargetedManagedAppConfigurationAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TargetedManagedAppPolicyAssignment, error) {
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
	var values []TargetedManagedAppPolicyAssignment
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
			value  []TargetedManagedAppPolicyAssignment
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

// GetN performs GET request for TargetedManagedAppPolicyAssignment collection, max N pages
func (r *TargetedManagedAppConfigurationAssignmentsCollectionRequest) GetN(ctx context.Context, n int) ([]TargetedManagedAppPolicyAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TargetedManagedAppPolicyAssignment collection
func (r *TargetedManagedAppConfigurationAssignmentsCollectionRequest) Get(ctx context.Context) ([]TargetedManagedAppPolicyAssignment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TargetedManagedAppPolicyAssignment collection
func (r *TargetedManagedAppConfigurationAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *TargetedManagedAppPolicyAssignment) (resObj *TargetedManagedAppPolicyAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// DeploymentSummary is navigation property
func (b *TargetedManagedAppConfigurationRequestBuilder) DeploymentSummary() *ManagedAppPolicyDeploymentSummaryRequestBuilder {
	bb := &ManagedAppPolicyDeploymentSummaryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deploymentSummary"
	return bb
}

// Assignments returns request builder for TargetedManagedAppPolicyAssignment collection
func (b *TargetedManagedAppProtectionRequestBuilder) Assignments() *TargetedManagedAppProtectionAssignmentsCollectionRequestBuilder {
	bb := &TargetedManagedAppProtectionAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// TargetedManagedAppProtectionAssignmentsCollectionRequestBuilder is request builder for TargetedManagedAppPolicyAssignment collection
type TargetedManagedAppProtectionAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TargetedManagedAppPolicyAssignment collection
func (b *TargetedManagedAppProtectionAssignmentsCollectionRequestBuilder) Request() *TargetedManagedAppProtectionAssignmentsCollectionRequest {
	return &TargetedManagedAppProtectionAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TargetedManagedAppPolicyAssignment item
func (b *TargetedManagedAppProtectionAssignmentsCollectionRequestBuilder) ID(id string) *TargetedManagedAppPolicyAssignmentRequestBuilder {
	bb := &TargetedManagedAppPolicyAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TargetedManagedAppProtectionAssignmentsCollectionRequest is request for TargetedManagedAppPolicyAssignment collection
type TargetedManagedAppProtectionAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TargetedManagedAppPolicyAssignment collection
func (r *TargetedManagedAppProtectionAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TargetedManagedAppPolicyAssignment, error) {
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
	var values []TargetedManagedAppPolicyAssignment
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
			value  []TargetedManagedAppPolicyAssignment
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

// GetN performs GET request for TargetedManagedAppPolicyAssignment collection, max N pages
func (r *TargetedManagedAppProtectionAssignmentsCollectionRequest) GetN(ctx context.Context, n int) ([]TargetedManagedAppPolicyAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TargetedManagedAppPolicyAssignment collection
func (r *TargetedManagedAppProtectionAssignmentsCollectionRequest) Get(ctx context.Context) ([]TargetedManagedAppPolicyAssignment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TargetedManagedAppPolicyAssignment collection
func (r *TargetedManagedAppProtectionAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *TargetedManagedAppPolicyAssignment) (resObj *TargetedManagedAppPolicyAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
