// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WorkbookCommentReplyRequestBuilder is request builder for WorkbookCommentReply
type WorkbookCommentReplyRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookCommentReplyRequest
func (b *WorkbookCommentReplyRequestBuilder) Request() *WorkbookCommentReplyRequest {
	return &WorkbookCommentReplyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookCommentReplyRequest is request for WorkbookCommentReply
type WorkbookCommentReplyRequest struct{ BaseRequest }

// Do performs HTTP request for WorkbookCommentReply
func (r *WorkbookCommentReplyRequest) Do(method, path string, reqObj interface{}) (resObj *WorkbookCommentReply, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for WorkbookCommentReply
func (r *WorkbookCommentReplyRequest) Get() (*WorkbookCommentReply, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for WorkbookCommentReply
func (r *WorkbookCommentReplyRequest) Update(reqObj *WorkbookCommentReply) (*WorkbookCommentReply, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for WorkbookCommentReply
func (r *WorkbookCommentReplyRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
