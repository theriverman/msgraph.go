// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ChannelRequestBuilder is request builder for Channel
type ChannelRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChannelRequest
func (b *ChannelRequestBuilder) Request() *ChannelRequest {
	return &ChannelRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChannelRequest is request for Channel
type ChannelRequest struct{ BaseRequest }

// Do performs HTTP request for Channel
func (r *ChannelRequest) Do(method, path string, reqObj interface{}) (resObj *Channel, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for Channel
func (r *ChannelRequest) Get() (*Channel, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for Channel
func (r *ChannelRequest) Update(reqObj *Channel) (*Channel, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for Channel
func (r *ChannelRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Tabs returns request builder for TeamsTab collection
func (b *ChannelRequestBuilder) Tabs() *ChannelTabsCollectionRequestBuilder {
	bb := &ChannelTabsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/tabs"
	return bb
}

// ChannelTabsCollectionRequestBuilder is request builder for TeamsTab collection
type ChannelTabsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamsTab collection
func (b *ChannelTabsCollectionRequestBuilder) Request() *ChannelTabsCollectionRequest {
	return &ChannelTabsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamsTab item
func (b *ChannelTabsCollectionRequestBuilder) ID(id string) *TeamsTabRequestBuilder {
	bb := &TeamsTabRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ChannelTabsCollectionRequest is request for TeamsTab collection
type ChannelTabsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for TeamsTab collection
func (r *ChannelTabsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *TeamsTab, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for TeamsTab collection
func (r *ChannelTabsCollectionRequest) Paging(method, path string, obj interface{}) ([]TeamsTab, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []TeamsTab
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []TeamsTab
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

// Get performs GET request for TeamsTab collection
func (r *ChannelTabsCollectionRequest) Get() ([]TeamsTab, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for TeamsTab collection
func (r *ChannelTabsCollectionRequest) Add(reqObj *TeamsTab) (*TeamsTab, error) {
	return r.Do("POST", "", reqObj)
}
