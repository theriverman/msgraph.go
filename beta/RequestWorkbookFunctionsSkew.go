// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsSkewRequestBuilder struct{ BaseRequestBuilder }

// Skew action undocumented
func (b *WorkbookFunctionsRequestBuilder) Skew(reqObj *WorkbookFunctionsSkewRequestParameter) *WorkbookFunctionsSkewRequestBuilder {
	bb := &WorkbookFunctionsSkewRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/skew"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsSkewRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsSkewRequestBuilder) Request() *WorkbookFunctionsSkewRequest {
	return &WorkbookFunctionsSkewRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsSkewRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
