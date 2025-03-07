// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/theriverman/msgraph.go/jsonx"
)

// ConversationThreadReplyRequestParameter undocumented
type ConversationThreadReplyRequestParameter struct {
	// Post undocumented
	Post *Post `json:"Post,omitempty"`
}

// Threads returns request builder for ConversationThread collection
func (b *ConversationRequestBuilder) Threads() *ConversationThreadsCollectionRequestBuilder {
	bb := &ConversationThreadsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/threads"
	return bb
}

// ConversationThreadsCollectionRequestBuilder is request builder for ConversationThread collection
type ConversationThreadsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ConversationThread collection
func (b *ConversationThreadsCollectionRequestBuilder) Request() *ConversationThreadsCollectionRequest {
	return &ConversationThreadsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ConversationThread item
func (b *ConversationThreadsCollectionRequestBuilder) ID(id string) *ConversationThreadRequestBuilder {
	bb := &ConversationThreadRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ConversationThreadsCollectionRequest is request for ConversationThread collection
type ConversationThreadsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ConversationThread collection
func (r *ConversationThreadsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ConversationThread, error) {
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
	var values []ConversationThread
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
			value  []ConversationThread
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

// GetN performs GET request for ConversationThread collection, max N pages
func (r *ConversationThreadsCollectionRequest) GetN(ctx context.Context, n int) ([]ConversationThread, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ConversationThread collection
func (r *ConversationThreadsCollectionRequest) Get(ctx context.Context) ([]ConversationThread, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ConversationThread collection
func (r *ConversationThreadsCollectionRequest) Add(ctx context.Context, reqObj *ConversationThread) (resObj *ConversationThread, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Posts returns request builder for Post collection
func (b *ConversationThreadRequestBuilder) Posts() *ConversationThreadPostsCollectionRequestBuilder {
	bb := &ConversationThreadPostsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/posts"
	return bb
}

// ConversationThreadPostsCollectionRequestBuilder is request builder for Post collection
type ConversationThreadPostsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Post collection
func (b *ConversationThreadPostsCollectionRequestBuilder) Request() *ConversationThreadPostsCollectionRequest {
	return &ConversationThreadPostsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Post item
func (b *ConversationThreadPostsCollectionRequestBuilder) ID(id string) *PostRequestBuilder {
	bb := &PostRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ConversationThreadPostsCollectionRequest is request for Post collection
type ConversationThreadPostsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Post collection
func (r *ConversationThreadPostsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Post, error) {
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
	var values []Post
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
			value  []Post
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

// GetN performs GET request for Post collection, max N pages
func (r *ConversationThreadPostsCollectionRequest) GetN(ctx context.Context, n int) ([]Post, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Post collection
func (r *ConversationThreadPostsCollectionRequest) Get(ctx context.Context) ([]Post, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Post collection
func (r *ConversationThreadPostsCollectionRequest) Add(ctx context.Context, reqObj *Post) (resObj *Post, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
