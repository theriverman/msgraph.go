// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsDgetRequestBuilder struct{ BaseRequestBuilder }

// Dget action undocumented
func (b *WorkbookFunctionsRequestBuilder) Dget(reqObj *WorkbookFunctionsDgetRequestParameter) *WorkbookFunctionsDgetRequestBuilder {
	bb := &WorkbookFunctionsDgetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/dget"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsDgetRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsDgetRequestBuilder) Request() *WorkbookFunctionsDgetRequest {
	return &WorkbookFunctionsDgetRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsDgetRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
