// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// WorkbookCommentRequestBuilder is request builder for WorkbookComment
type WorkbookCommentRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookCommentRequest
func (b *WorkbookCommentRequestBuilder) Request() *WorkbookCommentRequest {
	return &WorkbookCommentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookCommentRequest is request for WorkbookComment
type WorkbookCommentRequest struct{ BaseRequest }

// Do performs HTTP request for WorkbookComment
func (r *WorkbookCommentRequest) Do(method, path string, reqObj interface{}) (resObj *WorkbookComment, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for WorkbookComment
func (r *WorkbookCommentRequest) Get() (*WorkbookComment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for WorkbookComment
func (r *WorkbookCommentRequest) Update(reqObj *WorkbookComment) (*WorkbookComment, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for WorkbookComment
func (r *WorkbookCommentRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Replies returns request builder for WorkbookCommentReply collection
func (b *WorkbookCommentRequestBuilder) Replies() *WorkbookCommentRepliesCollectionRequestBuilder {
	bb := &WorkbookCommentRepliesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/replies"
	return bb
}

// WorkbookCommentRepliesCollectionRequestBuilder is request builder for WorkbookCommentReply collection
type WorkbookCommentRepliesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WorkbookCommentReply collection
func (b *WorkbookCommentRepliesCollectionRequestBuilder) Request() *WorkbookCommentRepliesCollectionRequest {
	return &WorkbookCommentRepliesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WorkbookCommentReply item
func (b *WorkbookCommentRepliesCollectionRequestBuilder) ID(id string) *WorkbookCommentReplyRequestBuilder {
	bb := &WorkbookCommentReplyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WorkbookCommentRepliesCollectionRequest is request for WorkbookCommentReply collection
type WorkbookCommentRepliesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for WorkbookCommentReply collection
func (r *WorkbookCommentRepliesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *WorkbookCommentReply, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for WorkbookCommentReply collection
func (r *WorkbookCommentRepliesCollectionRequest) Paging(method, path string, obj interface{}) ([]WorkbookCommentReply, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []WorkbookCommentReply
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []WorkbookCommentReply
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

// Get performs GET request for WorkbookCommentReply collection
func (r *WorkbookCommentRepliesCollectionRequest) Get() ([]WorkbookCommentReply, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for WorkbookCommentReply collection
func (r *WorkbookCommentRepliesCollectionRequest) Add(reqObj *WorkbookCommentReply) (*WorkbookCommentReply, error) {
	return r.Do("POST", "", reqObj)
}
