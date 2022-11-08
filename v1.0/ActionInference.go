// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/theriverman/msgraph.go/jsonx"
)

// Overrides returns request builder for InferenceClassificationOverride collection
func (b *InferenceClassificationRequestBuilder) Overrides() *InferenceClassificationOverridesCollectionRequestBuilder {
	bb := &InferenceClassificationOverridesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/overrides"
	return bb
}

// InferenceClassificationOverridesCollectionRequestBuilder is request builder for InferenceClassificationOverride collection
type InferenceClassificationOverridesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for InferenceClassificationOverride collection
func (b *InferenceClassificationOverridesCollectionRequestBuilder) Request() *InferenceClassificationOverridesCollectionRequest {
	return &InferenceClassificationOverridesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for InferenceClassificationOverride item
func (b *InferenceClassificationOverridesCollectionRequestBuilder) ID(id string) *InferenceClassificationOverrideRequestBuilder {
	bb := &InferenceClassificationOverrideRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// InferenceClassificationOverridesCollectionRequest is request for InferenceClassificationOverride collection
type InferenceClassificationOverridesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for InferenceClassificationOverride collection
func (r *InferenceClassificationOverridesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]InferenceClassificationOverride, error) {
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
	var values []InferenceClassificationOverride
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
			value  []InferenceClassificationOverride
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

// GetN performs GET request for InferenceClassificationOverride collection, max N pages
func (r *InferenceClassificationOverridesCollectionRequest) GetN(ctx context.Context, n int) ([]InferenceClassificationOverride, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for InferenceClassificationOverride collection
func (r *InferenceClassificationOverridesCollectionRequest) Get(ctx context.Context) ([]InferenceClassificationOverride, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for InferenceClassificationOverride collection
func (r *InferenceClassificationOverridesCollectionRequest) Add(ctx context.Context, reqObj *InferenceClassificationOverride) (resObj *InferenceClassificationOverride, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
