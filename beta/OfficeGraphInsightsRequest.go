// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// OfficeGraphInsightsRequestBuilder is request builder for OfficeGraphInsights
type OfficeGraphInsightsRequestBuilder struct{ BaseRequestBuilder }

// Request returns OfficeGraphInsightsRequest
func (b *OfficeGraphInsightsRequestBuilder) Request() *OfficeGraphInsightsRequest {
	return &OfficeGraphInsightsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OfficeGraphInsightsRequest is request for OfficeGraphInsights
type OfficeGraphInsightsRequest struct{ BaseRequest }

// Do performs HTTP request for OfficeGraphInsights
func (r *OfficeGraphInsightsRequest) Do(method, path string, reqObj interface{}) (resObj *OfficeGraphInsights, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for OfficeGraphInsights
func (r *OfficeGraphInsightsRequest) Get() (*OfficeGraphInsights, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for OfficeGraphInsights
func (r *OfficeGraphInsightsRequest) Update(reqObj *OfficeGraphInsights) (*OfficeGraphInsights, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for OfficeGraphInsights
func (r *OfficeGraphInsightsRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Shared returns request builder for SharedInsight collection
func (b *OfficeGraphInsightsRequestBuilder) Shared() *OfficeGraphInsightsSharedCollectionRequestBuilder {
	bb := &OfficeGraphInsightsSharedCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/shared"
	return bb
}

// OfficeGraphInsightsSharedCollectionRequestBuilder is request builder for SharedInsight collection
type OfficeGraphInsightsSharedCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SharedInsight collection
func (b *OfficeGraphInsightsSharedCollectionRequestBuilder) Request() *OfficeGraphInsightsSharedCollectionRequest {
	return &OfficeGraphInsightsSharedCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SharedInsight item
func (b *OfficeGraphInsightsSharedCollectionRequestBuilder) ID(id string) *SharedInsightRequestBuilder {
	bb := &SharedInsightRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OfficeGraphInsightsSharedCollectionRequest is request for SharedInsight collection
type OfficeGraphInsightsSharedCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for SharedInsight collection
func (r *OfficeGraphInsightsSharedCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *SharedInsight, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for SharedInsight collection
func (r *OfficeGraphInsightsSharedCollectionRequest) Paging(method, path string, obj interface{}) ([]SharedInsight, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []SharedInsight
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []SharedInsight
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

// Get performs GET request for SharedInsight collection
func (r *OfficeGraphInsightsSharedCollectionRequest) Get() ([]SharedInsight, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for SharedInsight collection
func (r *OfficeGraphInsightsSharedCollectionRequest) Add(reqObj *SharedInsight) (*SharedInsight, error) {
	return r.Do("POST", "", reqObj)
}

// Trending returns request builder for Trending collection
func (b *OfficeGraphInsightsRequestBuilder) Trending() *OfficeGraphInsightsTrendingCollectionRequestBuilder {
	bb := &OfficeGraphInsightsTrendingCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/trending"
	return bb
}

// OfficeGraphInsightsTrendingCollectionRequestBuilder is request builder for Trending collection
type OfficeGraphInsightsTrendingCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Trending collection
func (b *OfficeGraphInsightsTrendingCollectionRequestBuilder) Request() *OfficeGraphInsightsTrendingCollectionRequest {
	return &OfficeGraphInsightsTrendingCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Trending item
func (b *OfficeGraphInsightsTrendingCollectionRequestBuilder) ID(id string) *TrendingRequestBuilder {
	bb := &TrendingRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OfficeGraphInsightsTrendingCollectionRequest is request for Trending collection
type OfficeGraphInsightsTrendingCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for Trending collection
func (r *OfficeGraphInsightsTrendingCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *Trending, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for Trending collection
func (r *OfficeGraphInsightsTrendingCollectionRequest) Paging(method, path string, obj interface{}) ([]Trending, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Trending
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []Trending
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

// Get performs GET request for Trending collection
func (r *OfficeGraphInsightsTrendingCollectionRequest) Get() ([]Trending, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Trending collection
func (r *OfficeGraphInsightsTrendingCollectionRequest) Add(reqObj *Trending) (*Trending, error) {
	return r.Do("POST", "", reqObj)
}

// Used returns request builder for UsedInsight collection
func (b *OfficeGraphInsightsRequestBuilder) Used() *OfficeGraphInsightsUsedCollectionRequestBuilder {
	bb := &OfficeGraphInsightsUsedCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/used"
	return bb
}

// OfficeGraphInsightsUsedCollectionRequestBuilder is request builder for UsedInsight collection
type OfficeGraphInsightsUsedCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for UsedInsight collection
func (b *OfficeGraphInsightsUsedCollectionRequestBuilder) Request() *OfficeGraphInsightsUsedCollectionRequest {
	return &OfficeGraphInsightsUsedCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for UsedInsight item
func (b *OfficeGraphInsightsUsedCollectionRequestBuilder) ID(id string) *UsedInsightRequestBuilder {
	bb := &UsedInsightRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OfficeGraphInsightsUsedCollectionRequest is request for UsedInsight collection
type OfficeGraphInsightsUsedCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for UsedInsight collection
func (r *OfficeGraphInsightsUsedCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *UsedInsight, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for UsedInsight collection
func (r *OfficeGraphInsightsUsedCollectionRequest) Paging(method, path string, obj interface{}) ([]UsedInsight, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []UsedInsight
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []UsedInsight
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

// Get performs GET request for UsedInsight collection
func (r *OfficeGraphInsightsUsedCollectionRequest) Get() ([]UsedInsight, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for UsedInsight collection
func (r *OfficeGraphInsightsUsedCollectionRequest) Add(reqObj *UsedInsight) (*UsedInsight, error) {
	return r.Do("POST", "", reqObj)
}
