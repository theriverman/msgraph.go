// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DimensionRequestBuilder is request builder for Dimension
type DimensionRequestBuilder struct{ BaseRequestBuilder }

// Request returns DimensionRequest
func (b *DimensionRequestBuilder) Request() *DimensionRequest {
	return &DimensionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DimensionRequest is request for Dimension
type DimensionRequest struct{ BaseRequest }

// Do performs HTTP request for Dimension
func (r *DimensionRequest) Do(method, path string, reqObj interface{}) (resObj *Dimension, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for Dimension
func (r *DimensionRequest) Get() (*Dimension, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for Dimension
func (r *DimensionRequest) Update(reqObj *Dimension) (*Dimension, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for Dimension
func (r *DimensionRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// DimensionValues returns request builder for DimensionValue collection
func (b *DimensionRequestBuilder) DimensionValues() *DimensionDimensionValuesCollectionRequestBuilder {
	bb := &DimensionDimensionValuesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/dimensionValues"
	return bb
}

// DimensionDimensionValuesCollectionRequestBuilder is request builder for DimensionValue collection
type DimensionDimensionValuesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DimensionValue collection
func (b *DimensionDimensionValuesCollectionRequestBuilder) Request() *DimensionDimensionValuesCollectionRequest {
	return &DimensionDimensionValuesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DimensionValue item
func (b *DimensionDimensionValuesCollectionRequestBuilder) ID(id string) *DimensionValueRequestBuilder {
	bb := &DimensionValueRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DimensionDimensionValuesCollectionRequest is request for DimensionValue collection
type DimensionDimensionValuesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for DimensionValue collection
func (r *DimensionDimensionValuesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *DimensionValue, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for DimensionValue collection
func (r *DimensionDimensionValuesCollectionRequest) Paging(method, path string, obj interface{}) ([]DimensionValue, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DimensionValue
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DimensionValue
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

// Get performs GET request for DimensionValue collection
func (r *DimensionDimensionValuesCollectionRequest) Get() ([]DimensionValue, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DimensionValue collection
func (r *DimensionDimensionValuesCollectionRequest) Add(reqObj *DimensionValue) (*DimensionValue, error) {
	return r.Do("POST", "", reqObj)
}
