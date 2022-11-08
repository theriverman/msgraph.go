// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/theriverman/msgraph.go/jsonx"
)

// ColumnLinks returns request builder for ColumnLink collection
func (b *ContentTypeRequestBuilder) ColumnLinks() *ContentTypeColumnLinksCollectionRequestBuilder {
	bb := &ContentTypeColumnLinksCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/columnLinks"
	return bb
}

// ContentTypeColumnLinksCollectionRequestBuilder is request builder for ColumnLink collection
type ContentTypeColumnLinksCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ColumnLink collection
func (b *ContentTypeColumnLinksCollectionRequestBuilder) Request() *ContentTypeColumnLinksCollectionRequest {
	return &ContentTypeColumnLinksCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ColumnLink item
func (b *ContentTypeColumnLinksCollectionRequestBuilder) ID(id string) *ColumnLinkRequestBuilder {
	bb := &ColumnLinkRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ContentTypeColumnLinksCollectionRequest is request for ColumnLink collection
type ContentTypeColumnLinksCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ColumnLink collection
func (r *ContentTypeColumnLinksCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ColumnLink, error) {
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
	var values []ColumnLink
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
			value  []ColumnLink
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

// GetN performs GET request for ColumnLink collection, max N pages
func (r *ContentTypeColumnLinksCollectionRequest) GetN(ctx context.Context, n int) ([]ColumnLink, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ColumnLink collection
func (r *ContentTypeColumnLinksCollectionRequest) Get(ctx context.Context) ([]ColumnLink, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ColumnLink collection
func (r *ContentTypeColumnLinksCollectionRequest) Add(ctx context.Context, reqObj *ColumnLink) (resObj *ColumnLink, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
