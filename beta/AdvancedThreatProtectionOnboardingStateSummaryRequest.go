// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder is request builder for AdvancedThreatProtectionOnboardingStateSummary
type AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns AdvancedThreatProtectionOnboardingStateSummaryRequest
func (b *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) Request() *AdvancedThreatProtectionOnboardingStateSummaryRequest {
	return &AdvancedThreatProtectionOnboardingStateSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AdvancedThreatProtectionOnboardingStateSummaryRequest is request for AdvancedThreatProtectionOnboardingStateSummary
type AdvancedThreatProtectionOnboardingStateSummaryRequest struct{ BaseRequest }

// Do performs HTTP request for AdvancedThreatProtectionOnboardingStateSummary
func (r *AdvancedThreatProtectionOnboardingStateSummaryRequest) Do(method, path string, reqObj interface{}) (resObj *AdvancedThreatProtectionOnboardingStateSummary, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for AdvancedThreatProtectionOnboardingStateSummary
func (r *AdvancedThreatProtectionOnboardingStateSummaryRequest) Get() (*AdvancedThreatProtectionOnboardingStateSummary, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for AdvancedThreatProtectionOnboardingStateSummary
func (r *AdvancedThreatProtectionOnboardingStateSummaryRequest) Update(reqObj *AdvancedThreatProtectionOnboardingStateSummary) (*AdvancedThreatProtectionOnboardingStateSummary, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for AdvancedThreatProtectionOnboardingStateSummary
func (r *AdvancedThreatProtectionOnboardingStateSummaryRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// AdvancedThreatProtectionOnboardingDeviceSettingStates returns request builder for AdvancedThreatProtectionOnboardingDeviceSettingState collection
func (b *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) AdvancedThreatProtectionOnboardingDeviceSettingStates() *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequestBuilder {
	bb := &AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/advancedThreatProtectionOnboardingDeviceSettingStates"
	return bb
}

// AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequestBuilder is request builder for AdvancedThreatProtectionOnboardingDeviceSettingState collection
type AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AdvancedThreatProtectionOnboardingDeviceSettingState collection
func (b *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequestBuilder) Request() *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequest {
	return &AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AdvancedThreatProtectionOnboardingDeviceSettingState item
func (b *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequestBuilder) ID(id string) *AdvancedThreatProtectionOnboardingDeviceSettingStateRequestBuilder {
	bb := &AdvancedThreatProtectionOnboardingDeviceSettingStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequest is request for AdvancedThreatProtectionOnboardingDeviceSettingState collection
type AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for AdvancedThreatProtectionOnboardingDeviceSettingState collection
func (r *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *AdvancedThreatProtectionOnboardingDeviceSettingState, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for AdvancedThreatProtectionOnboardingDeviceSettingState collection
func (r *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequest) Paging(method, path string, obj interface{}) ([]AdvancedThreatProtectionOnboardingDeviceSettingState, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []AdvancedThreatProtectionOnboardingDeviceSettingState
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []AdvancedThreatProtectionOnboardingDeviceSettingState
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

// Get performs GET request for AdvancedThreatProtectionOnboardingDeviceSettingState collection
func (r *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequest) Get() ([]AdvancedThreatProtectionOnboardingDeviceSettingState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for AdvancedThreatProtectionOnboardingDeviceSettingState collection
func (r *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequest) Add(reqObj *AdvancedThreatProtectionOnboardingDeviceSettingState) (*AdvancedThreatProtectionOnboardingDeviceSettingState, error) {
	return r.Do("POST", "", reqObj)
}
