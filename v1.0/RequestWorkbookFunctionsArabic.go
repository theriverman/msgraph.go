// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsArabicRequestBuilder struct{ BaseRequestBuilder }

// Arabic action undocumented
func (b *WorkbookFunctionsRequestBuilder) Arabic(reqObj *WorkbookFunctionsArabicRequestParameter) *WorkbookFunctionsArabicRequestBuilder {
	bb := &WorkbookFunctionsArabicRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/arabic"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsArabicRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsArabicRequestBuilder) Request() *WorkbookFunctionsArabicRequest {
	return &WorkbookFunctionsArabicRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsArabicRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
