// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DirectoryRequestBuilder is request builder for Directory
type DirectoryRequestBuilder struct{ BaseRequestBuilder }

// Request returns DirectoryRequest
func (b *DirectoryRequestBuilder) Request() *DirectoryRequest {
	return &DirectoryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DirectoryRequest is request for Directory
type DirectoryRequest struct{ BaseRequest }

// Do performs HTTP request for Directory
func (r *DirectoryRequest) Do(method, path string, reqObj interface{}) (resObj *Directory, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for Directory
func (r *DirectoryRequest) Get() (*Directory, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for Directory
func (r *DirectoryRequest) Update(reqObj *Directory) (*Directory, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for Directory
func (r *DirectoryRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// DeletedItems returns request builder for DirectoryObject collection
func (b *DirectoryRequestBuilder) DeletedItems() *DirectoryDeletedItemsCollectionRequestBuilder {
	bb := &DirectoryDeletedItemsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deletedItems"
	return bb
}

// DirectoryDeletedItemsCollectionRequestBuilder is request builder for DirectoryObject collection
type DirectoryDeletedItemsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *DirectoryDeletedItemsCollectionRequestBuilder) Request() *DirectoryDeletedItemsCollectionRequest {
	return &DirectoryDeletedItemsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *DirectoryDeletedItemsCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DirectoryDeletedItemsCollectionRequest is request for DirectoryObject collection
type DirectoryDeletedItemsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for DirectoryObject collection
func (r *DirectoryDeletedItemsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for DirectoryObject collection
func (r *DirectoryDeletedItemsCollectionRequest) Paging(method, path string, obj interface{}) ([]DirectoryObject, error) {
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
func (r *DirectoryDeletedItemsCollectionRequest) Get() ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DirectoryObject collection
func (r *DirectoryDeletedItemsCollectionRequest) Add(reqObj *DirectoryObject) (*DirectoryObject, error) {
	return r.Do("POST", "", reqObj)
}
