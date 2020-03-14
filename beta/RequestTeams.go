// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// TeamsAppRequestBuilder is request builder for TeamsApp
type TeamsAppRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamsAppRequest
func (b *TeamsAppRequestBuilder) Request() *TeamsAppRequest {
	return &TeamsAppRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamsAppRequest is request for TeamsApp
type TeamsAppRequest struct{ BaseRequest }

// Get performs GET request for TeamsApp
func (r *TeamsAppRequest) Get(ctx context.Context) (resObj *TeamsApp, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamsApp
func (r *TeamsAppRequest) Update(ctx context.Context, reqObj *TeamsApp) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamsApp
func (r *TeamsAppRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamsAppDefinitionRequestBuilder is request builder for TeamsAppDefinition
type TeamsAppDefinitionRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamsAppDefinitionRequest
func (b *TeamsAppDefinitionRequestBuilder) Request() *TeamsAppDefinitionRequest {
	return &TeamsAppDefinitionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamsAppDefinitionRequest is request for TeamsAppDefinition
type TeamsAppDefinitionRequest struct{ BaseRequest }

// Get performs GET request for TeamsAppDefinition
func (r *TeamsAppDefinitionRequest) Get(ctx context.Context) (resObj *TeamsAppDefinition, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamsAppDefinition
func (r *TeamsAppDefinitionRequest) Update(ctx context.Context, reqObj *TeamsAppDefinition) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamsAppDefinition
func (r *TeamsAppDefinitionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamsAppInstallationRequestBuilder is request builder for TeamsAppInstallation
type TeamsAppInstallationRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamsAppInstallationRequest
func (b *TeamsAppInstallationRequestBuilder) Request() *TeamsAppInstallationRequest {
	return &TeamsAppInstallationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamsAppInstallationRequest is request for TeamsAppInstallation
type TeamsAppInstallationRequest struct{ BaseRequest }

// Get performs GET request for TeamsAppInstallation
func (r *TeamsAppInstallationRequest) Get(ctx context.Context) (resObj *TeamsAppInstallation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamsAppInstallation
func (r *TeamsAppInstallationRequest) Update(ctx context.Context, reqObj *TeamsAppInstallation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamsAppInstallation
func (r *TeamsAppInstallationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamsAsyncOperationRequestBuilder is request builder for TeamsAsyncOperation
type TeamsAsyncOperationRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamsAsyncOperationRequest
func (b *TeamsAsyncOperationRequestBuilder) Request() *TeamsAsyncOperationRequest {
	return &TeamsAsyncOperationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamsAsyncOperationRequest is request for TeamsAsyncOperation
type TeamsAsyncOperationRequest struct{ BaseRequest }

// Get performs GET request for TeamsAsyncOperation
func (r *TeamsAsyncOperationRequest) Get(ctx context.Context) (resObj *TeamsAsyncOperation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamsAsyncOperation
func (r *TeamsAsyncOperationRequest) Update(ctx context.Context, reqObj *TeamsAsyncOperation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamsAsyncOperation
func (r *TeamsAsyncOperationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamsCatalogAppRequestBuilder is request builder for TeamsCatalogApp
type TeamsCatalogAppRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamsCatalogAppRequest
func (b *TeamsCatalogAppRequestBuilder) Request() *TeamsCatalogAppRequest {
	return &TeamsCatalogAppRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamsCatalogAppRequest is request for TeamsCatalogApp
type TeamsCatalogAppRequest struct{ BaseRequest }

// Get performs GET request for TeamsCatalogApp
func (r *TeamsCatalogAppRequest) Get(ctx context.Context) (resObj *TeamsCatalogApp, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamsCatalogApp
func (r *TeamsCatalogAppRequest) Update(ctx context.Context, reqObj *TeamsCatalogApp) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamsCatalogApp
func (r *TeamsCatalogAppRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamsTabRequestBuilder is request builder for TeamsTab
type TeamsTabRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamsTabRequest
func (b *TeamsTabRequestBuilder) Request() *TeamsTabRequest {
	return &TeamsTabRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamsTabRequest is request for TeamsTab
type TeamsTabRequest struct{ BaseRequest }

// Get performs GET request for TeamsTab
func (r *TeamsTabRequest) Get(ctx context.Context) (resObj *TeamsTab, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamsTab
func (r *TeamsTabRequest) Update(ctx context.Context, reqObj *TeamsTab) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamsTab
func (r *TeamsTabRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamsTemplateRequestBuilder is request builder for TeamsTemplate
type TeamsTemplateRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamsTemplateRequest
func (b *TeamsTemplateRequestBuilder) Request() *TeamsTemplateRequest {
	return &TeamsTemplateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamsTemplateRequest is request for TeamsTemplate
type TeamsTemplateRequest struct{ BaseRequest }

// Get performs GET request for TeamsTemplate
func (r *TeamsTemplateRequest) Get(ctx context.Context) (resObj *TeamsTemplate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamsTemplate
func (r *TeamsTemplateRequest) Update(ctx context.Context, reqObj *TeamsTemplate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamsTemplate
func (r *TeamsTemplateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type TeamsAppInstallationUpgradeRequestBuilder struct{ BaseRequestBuilder }

// Upgrade action undocumented
func (b *TeamsAppInstallationRequestBuilder) Upgrade(reqObj *TeamsAppInstallationUpgradeRequestParameter) *TeamsAppInstallationUpgradeRequestBuilder {
	bb := &TeamsAppInstallationUpgradeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/upgrade"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type TeamsAppInstallationUpgradeRequest struct{ BaseRequest }

//
func (b *TeamsAppInstallationUpgradeRequestBuilder) Request() *TeamsAppInstallationUpgradeRequest {
	return &TeamsAppInstallationUpgradeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *TeamsAppInstallationUpgradeRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
