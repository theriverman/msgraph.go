// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsDays360RequestBuilder struct{ BaseRequestBuilder }

// Days360 action undocumented
func (b *WorkbookFunctionsRequestBuilder) Days360(reqObj *WorkbookFunctionsDays360RequestParameter) *WorkbookFunctionsDays360RequestBuilder {
	bb := &WorkbookFunctionsDays360RequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/days360"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsDays360Request struct{ BaseRequest }

//
func (b *WorkbookFunctionsDays360RequestBuilder) Request() *WorkbookFunctionsDays360Request {
	return &WorkbookFunctionsDays360Request{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsDays360Request) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
