// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ItemAnalyticsRequestBuilder is request builder for ItemAnalytics
type ItemAnalyticsRequestBuilder struct{ BaseRequestBuilder }

// Request returns ItemAnalyticsRequest
func (b *ItemAnalyticsRequestBuilder) Request() *ItemAnalyticsRequest {
	return &ItemAnalyticsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ItemAnalyticsRequest is request for ItemAnalytics
type ItemAnalyticsRequest struct{ BaseRequest }

// Do performs HTTP request for ItemAnalytics
func (r *ItemAnalyticsRequest) Do(method, path string, reqObj interface{}) (resObj *ItemAnalytics, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for ItemAnalytics
func (r *ItemAnalyticsRequest) Get() (*ItemAnalytics, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for ItemAnalytics
func (r *ItemAnalyticsRequest) Update(reqObj *ItemAnalytics) (*ItemAnalytics, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for ItemAnalytics
func (r *ItemAnalyticsRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// AllTime is navigation property
func (b *ItemAnalyticsRequestBuilder) AllTime() *ItemActivityStatRequestBuilder {
	bb := &ItemActivityStatRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/allTime"
	return bb
}

// ItemActivityStats returns request builder for ItemActivityStat collection
func (b *ItemAnalyticsRequestBuilder) ItemActivityStats() *ItemAnalyticsItemActivityStatsCollectionRequestBuilder {
	bb := &ItemAnalyticsItemActivityStatsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/itemActivityStats"
	return bb
}

// ItemAnalyticsItemActivityStatsCollectionRequestBuilder is request builder for ItemActivityStat collection
type ItemAnalyticsItemActivityStatsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ItemActivityStat collection
func (b *ItemAnalyticsItemActivityStatsCollectionRequestBuilder) Request() *ItemAnalyticsItemActivityStatsCollectionRequest {
	return &ItemAnalyticsItemActivityStatsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ItemActivityStat item
func (b *ItemAnalyticsItemActivityStatsCollectionRequestBuilder) ID(id string) *ItemActivityStatRequestBuilder {
	bb := &ItemActivityStatRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ItemAnalyticsItemActivityStatsCollectionRequest is request for ItemActivityStat collection
type ItemAnalyticsItemActivityStatsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for ItemActivityStat collection
func (r *ItemAnalyticsItemActivityStatsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *ItemActivityStat, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for ItemActivityStat collection
func (r *ItemAnalyticsItemActivityStatsCollectionRequest) Paging(method, path string, obj interface{}) ([]ItemActivityStat, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ItemActivityStat
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []ItemActivityStat
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

// Get performs GET request for ItemActivityStat collection
func (r *ItemAnalyticsItemActivityStatsCollectionRequest) Get() ([]ItemActivityStat, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ItemActivityStat collection
func (r *ItemAnalyticsItemActivityStatsCollectionRequest) Add(reqObj *ItemActivityStat) (*ItemActivityStat, error) {
	return r.Do("POST", "", reqObj)
}

// LastSevenDays is navigation property
func (b *ItemAnalyticsRequestBuilder) LastSevenDays() *ItemActivityStatRequestBuilder {
	bb := &ItemActivityStatRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/lastSevenDays"
	return bb
}
