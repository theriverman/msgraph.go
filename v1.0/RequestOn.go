// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// OnPremisesConditionalAccessSettingsRequestBuilder is request builder for OnPremisesConditionalAccessSettings
type OnPremisesConditionalAccessSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnPremisesConditionalAccessSettingsRequest
func (b *OnPremisesConditionalAccessSettingsRequestBuilder) Request() *OnPremisesConditionalAccessSettingsRequest {
	return &OnPremisesConditionalAccessSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnPremisesConditionalAccessSettingsRequest is request for OnPremisesConditionalAccessSettings
type OnPremisesConditionalAccessSettingsRequest struct{ BaseRequest }

// Get performs GET request for OnPremisesConditionalAccessSettings
func (r *OnPremisesConditionalAccessSettingsRequest) Get(ctx context.Context) (resObj *OnPremisesConditionalAccessSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnPremisesConditionalAccessSettings
func (r *OnPremisesConditionalAccessSettingsRequest) Update(ctx context.Context, reqObj *OnPremisesConditionalAccessSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnPremisesConditionalAccessSettings
func (r *OnPremisesConditionalAccessSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
