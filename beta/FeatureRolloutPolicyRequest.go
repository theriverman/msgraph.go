// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// FeatureRolloutPolicyRequestBuilder is request builder for FeatureRolloutPolicy
type FeatureRolloutPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns FeatureRolloutPolicyRequest
func (b *FeatureRolloutPolicyRequestBuilder) Request() *FeatureRolloutPolicyRequest {
	return &FeatureRolloutPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// FeatureRolloutPolicyRequest is request for FeatureRolloutPolicy
type FeatureRolloutPolicyRequest struct{ BaseRequest }

// Do performs HTTP request for FeatureRolloutPolicy
func (r *FeatureRolloutPolicyRequest) Do(method, path string, reqObj interface{}) (resObj *FeatureRolloutPolicy, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for FeatureRolloutPolicy
func (r *FeatureRolloutPolicyRequest) Get() (*FeatureRolloutPolicy, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for FeatureRolloutPolicy
func (r *FeatureRolloutPolicyRequest) Update(reqObj *FeatureRolloutPolicy) (*FeatureRolloutPolicy, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for FeatureRolloutPolicy
func (r *FeatureRolloutPolicyRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// AppliesTo returns request builder for DirectoryObject collection
func (b *FeatureRolloutPolicyRequestBuilder) AppliesTo() *FeatureRolloutPolicyAppliesToCollectionRequestBuilder {
	bb := &FeatureRolloutPolicyAppliesToCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/appliesTo"
	return bb
}

// FeatureRolloutPolicyAppliesToCollectionRequestBuilder is request builder for DirectoryObject collection
type FeatureRolloutPolicyAppliesToCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *FeatureRolloutPolicyAppliesToCollectionRequestBuilder) Request() *FeatureRolloutPolicyAppliesToCollectionRequest {
	return &FeatureRolloutPolicyAppliesToCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *FeatureRolloutPolicyAppliesToCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// FeatureRolloutPolicyAppliesToCollectionRequest is request for DirectoryObject collection
type FeatureRolloutPolicyAppliesToCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for DirectoryObject collection
func (r *FeatureRolloutPolicyAppliesToCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for DirectoryObject collection
func (r *FeatureRolloutPolicyAppliesToCollectionRequest) Paging(method, path string, obj interface{}) ([]DirectoryObject, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DirectoryObject
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DirectoryObject
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

// Get performs GET request for DirectoryObject collection
func (r *FeatureRolloutPolicyAppliesToCollectionRequest) Get() ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DirectoryObject collection
func (r *FeatureRolloutPolicyAppliesToCollectionRequest) Add(reqObj *DirectoryObject) (*DirectoryObject, error) {
	return r.Do("POST", "", reqObj)
}
