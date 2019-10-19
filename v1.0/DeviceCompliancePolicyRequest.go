// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DeviceCompliancePolicyRequestBuilder is request builder for DeviceCompliancePolicy
type DeviceCompliancePolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceCompliancePolicyRequest
func (b *DeviceCompliancePolicyRequestBuilder) Request() *DeviceCompliancePolicyRequest {
	return &DeviceCompliancePolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceCompliancePolicyRequest is request for DeviceCompliancePolicy
type DeviceCompliancePolicyRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceCompliancePolicy
func (r *DeviceCompliancePolicyRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceCompliancePolicy, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for DeviceCompliancePolicy
func (r *DeviceCompliancePolicyRequest) Get() (*DeviceCompliancePolicy, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for DeviceCompliancePolicy
func (r *DeviceCompliancePolicyRequest) Update(reqObj *DeviceCompliancePolicy) (*DeviceCompliancePolicy, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for DeviceCompliancePolicy
func (r *DeviceCompliancePolicyRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Assignments returns request builder for DeviceCompliancePolicyAssignment collection
func (b *DeviceCompliancePolicyRequestBuilder) Assignments() *DeviceCompliancePolicyAssignmentsCollectionRequestBuilder {
	bb := &DeviceCompliancePolicyAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// DeviceCompliancePolicyAssignmentsCollectionRequestBuilder is request builder for DeviceCompliancePolicyAssignment collection
type DeviceCompliancePolicyAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceCompliancePolicyAssignment collection
func (b *DeviceCompliancePolicyAssignmentsCollectionRequestBuilder) Request() *DeviceCompliancePolicyAssignmentsCollectionRequest {
	return &DeviceCompliancePolicyAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceCompliancePolicyAssignment item
func (b *DeviceCompliancePolicyAssignmentsCollectionRequestBuilder) ID(id string) *DeviceCompliancePolicyAssignmentRequestBuilder {
	bb := &DeviceCompliancePolicyAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceCompliancePolicyAssignmentsCollectionRequest is request for DeviceCompliancePolicyAssignment collection
type DeviceCompliancePolicyAssignmentsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceCompliancePolicyAssignment collection
func (r *DeviceCompliancePolicyAssignmentsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceCompliancePolicyAssignment, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for DeviceCompliancePolicyAssignment collection
func (r *DeviceCompliancePolicyAssignmentsCollectionRequest) Paging(method, path string, obj interface{}) ([]DeviceCompliancePolicyAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DeviceCompliancePolicyAssignment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DeviceCompliancePolicyAssignment
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

// Get performs GET request for DeviceCompliancePolicyAssignment collection
func (r *DeviceCompliancePolicyAssignmentsCollectionRequest) Get() ([]DeviceCompliancePolicyAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DeviceCompliancePolicyAssignment collection
func (r *DeviceCompliancePolicyAssignmentsCollectionRequest) Add(reqObj *DeviceCompliancePolicyAssignment) (*DeviceCompliancePolicyAssignment, error) {
	return r.Do("POST", "", reqObj)
}

// DeviceSettingStateSummaries returns request builder for SettingStateDeviceSummary collection
func (b *DeviceCompliancePolicyRequestBuilder) DeviceSettingStateSummaries() *DeviceCompliancePolicyDeviceSettingStateSummariesCollectionRequestBuilder {
	bb := &DeviceCompliancePolicyDeviceSettingStateSummariesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceSettingStateSummaries"
	return bb
}

// DeviceCompliancePolicyDeviceSettingStateSummariesCollectionRequestBuilder is request builder for SettingStateDeviceSummary collection
type DeviceCompliancePolicyDeviceSettingStateSummariesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SettingStateDeviceSummary collection
func (b *DeviceCompliancePolicyDeviceSettingStateSummariesCollectionRequestBuilder) Request() *DeviceCompliancePolicyDeviceSettingStateSummariesCollectionRequest {
	return &DeviceCompliancePolicyDeviceSettingStateSummariesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SettingStateDeviceSummary item
func (b *DeviceCompliancePolicyDeviceSettingStateSummariesCollectionRequestBuilder) ID(id string) *SettingStateDeviceSummaryRequestBuilder {
	bb := &SettingStateDeviceSummaryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceCompliancePolicyDeviceSettingStateSummariesCollectionRequest is request for SettingStateDeviceSummary collection
type DeviceCompliancePolicyDeviceSettingStateSummariesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for SettingStateDeviceSummary collection
func (r *DeviceCompliancePolicyDeviceSettingStateSummariesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *SettingStateDeviceSummary, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for SettingStateDeviceSummary collection
func (r *DeviceCompliancePolicyDeviceSettingStateSummariesCollectionRequest) Paging(method, path string, obj interface{}) ([]SettingStateDeviceSummary, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []SettingStateDeviceSummary
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []SettingStateDeviceSummary
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

// Get performs GET request for SettingStateDeviceSummary collection
func (r *DeviceCompliancePolicyDeviceSettingStateSummariesCollectionRequest) Get() ([]SettingStateDeviceSummary, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for SettingStateDeviceSummary collection
func (r *DeviceCompliancePolicyDeviceSettingStateSummariesCollectionRequest) Add(reqObj *SettingStateDeviceSummary) (*SettingStateDeviceSummary, error) {
	return r.Do("POST", "", reqObj)
}

// DeviceStatusOverview is navigation property
func (b *DeviceCompliancePolicyRequestBuilder) DeviceStatusOverview() *DeviceComplianceDeviceOverviewRequestBuilder {
	bb := &DeviceComplianceDeviceOverviewRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceStatusOverview"
	return bb
}

// DeviceStatuses returns request builder for DeviceComplianceDeviceStatus collection
func (b *DeviceCompliancePolicyRequestBuilder) DeviceStatuses() *DeviceCompliancePolicyDeviceStatusesCollectionRequestBuilder {
	bb := &DeviceCompliancePolicyDeviceStatusesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceStatuses"
	return bb
}

// DeviceCompliancePolicyDeviceStatusesCollectionRequestBuilder is request builder for DeviceComplianceDeviceStatus collection
type DeviceCompliancePolicyDeviceStatusesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceComplianceDeviceStatus collection
func (b *DeviceCompliancePolicyDeviceStatusesCollectionRequestBuilder) Request() *DeviceCompliancePolicyDeviceStatusesCollectionRequest {
	return &DeviceCompliancePolicyDeviceStatusesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceComplianceDeviceStatus item
func (b *DeviceCompliancePolicyDeviceStatusesCollectionRequestBuilder) ID(id string) *DeviceComplianceDeviceStatusRequestBuilder {
	bb := &DeviceComplianceDeviceStatusRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceCompliancePolicyDeviceStatusesCollectionRequest is request for DeviceComplianceDeviceStatus collection
type DeviceCompliancePolicyDeviceStatusesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceComplianceDeviceStatus collection
func (r *DeviceCompliancePolicyDeviceStatusesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceComplianceDeviceStatus, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for DeviceComplianceDeviceStatus collection
func (r *DeviceCompliancePolicyDeviceStatusesCollectionRequest) Paging(method, path string, obj interface{}) ([]DeviceComplianceDeviceStatus, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DeviceComplianceDeviceStatus
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DeviceComplianceDeviceStatus
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

// Get performs GET request for DeviceComplianceDeviceStatus collection
func (r *DeviceCompliancePolicyDeviceStatusesCollectionRequest) Get() ([]DeviceComplianceDeviceStatus, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DeviceComplianceDeviceStatus collection
func (r *DeviceCompliancePolicyDeviceStatusesCollectionRequest) Add(reqObj *DeviceComplianceDeviceStatus) (*DeviceComplianceDeviceStatus, error) {
	return r.Do("POST", "", reqObj)
}

// ScheduledActionsForRule returns request builder for DeviceComplianceScheduledActionForRule collection
func (b *DeviceCompliancePolicyRequestBuilder) ScheduledActionsForRule() *DeviceCompliancePolicyScheduledActionsForRuleCollectionRequestBuilder {
	bb := &DeviceCompliancePolicyScheduledActionsForRuleCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/scheduledActionsForRule"
	return bb
}

// DeviceCompliancePolicyScheduledActionsForRuleCollectionRequestBuilder is request builder for DeviceComplianceScheduledActionForRule collection
type DeviceCompliancePolicyScheduledActionsForRuleCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceComplianceScheduledActionForRule collection
func (b *DeviceCompliancePolicyScheduledActionsForRuleCollectionRequestBuilder) Request() *DeviceCompliancePolicyScheduledActionsForRuleCollectionRequest {
	return &DeviceCompliancePolicyScheduledActionsForRuleCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceComplianceScheduledActionForRule item
func (b *DeviceCompliancePolicyScheduledActionsForRuleCollectionRequestBuilder) ID(id string) *DeviceComplianceScheduledActionForRuleRequestBuilder {
	bb := &DeviceComplianceScheduledActionForRuleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceCompliancePolicyScheduledActionsForRuleCollectionRequest is request for DeviceComplianceScheduledActionForRule collection
type DeviceCompliancePolicyScheduledActionsForRuleCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceComplianceScheduledActionForRule collection
func (r *DeviceCompliancePolicyScheduledActionsForRuleCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceComplianceScheduledActionForRule, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for DeviceComplianceScheduledActionForRule collection
func (r *DeviceCompliancePolicyScheduledActionsForRuleCollectionRequest) Paging(method, path string, obj interface{}) ([]DeviceComplianceScheduledActionForRule, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DeviceComplianceScheduledActionForRule
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DeviceComplianceScheduledActionForRule
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

// Get performs GET request for DeviceComplianceScheduledActionForRule collection
func (r *DeviceCompliancePolicyScheduledActionsForRuleCollectionRequest) Get() ([]DeviceComplianceScheduledActionForRule, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DeviceComplianceScheduledActionForRule collection
func (r *DeviceCompliancePolicyScheduledActionsForRuleCollectionRequest) Add(reqObj *DeviceComplianceScheduledActionForRule) (*DeviceComplianceScheduledActionForRule, error) {
	return r.Do("POST", "", reqObj)
}

// UserStatusOverview is navigation property
func (b *DeviceCompliancePolicyRequestBuilder) UserStatusOverview() *DeviceComplianceUserOverviewRequestBuilder {
	bb := &DeviceComplianceUserOverviewRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/userStatusOverview"
	return bb
}

// UserStatuses returns request builder for DeviceComplianceUserStatus collection
func (b *DeviceCompliancePolicyRequestBuilder) UserStatuses() *DeviceCompliancePolicyUserStatusesCollectionRequestBuilder {
	bb := &DeviceCompliancePolicyUserStatusesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/userStatuses"
	return bb
}

// DeviceCompliancePolicyUserStatusesCollectionRequestBuilder is request builder for DeviceComplianceUserStatus collection
type DeviceCompliancePolicyUserStatusesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceComplianceUserStatus collection
func (b *DeviceCompliancePolicyUserStatusesCollectionRequestBuilder) Request() *DeviceCompliancePolicyUserStatusesCollectionRequest {
	return &DeviceCompliancePolicyUserStatusesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceComplianceUserStatus item
func (b *DeviceCompliancePolicyUserStatusesCollectionRequestBuilder) ID(id string) *DeviceComplianceUserStatusRequestBuilder {
	bb := &DeviceComplianceUserStatusRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceCompliancePolicyUserStatusesCollectionRequest is request for DeviceComplianceUserStatus collection
type DeviceCompliancePolicyUserStatusesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceComplianceUserStatus collection
func (r *DeviceCompliancePolicyUserStatusesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceComplianceUserStatus, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for DeviceComplianceUserStatus collection
func (r *DeviceCompliancePolicyUserStatusesCollectionRequest) Paging(method, path string, obj interface{}) ([]DeviceComplianceUserStatus, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DeviceComplianceUserStatus
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DeviceComplianceUserStatus
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

// Get performs GET request for DeviceComplianceUserStatus collection
func (r *DeviceCompliancePolicyUserStatusesCollectionRequest) Get() ([]DeviceComplianceUserStatus, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DeviceComplianceUserStatus collection
func (r *DeviceCompliancePolicyUserStatusesCollectionRequest) Add(reqObj *DeviceComplianceUserStatus) (*DeviceComplianceUserStatus, error) {
	return r.Do("POST", "", reqObj)
}
