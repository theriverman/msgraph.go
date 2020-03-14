// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// ContactRequestBuilder is request builder for Contact
type ContactRequestBuilder struct{ BaseRequestBuilder }

// Request returns ContactRequest
func (b *ContactRequestBuilder) Request() *ContactRequest {
	return &ContactRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ContactRequest is request for Contact
type ContactRequest struct{ BaseRequest }

// Get performs GET request for Contact
func (r *ContactRequest) Get(ctx context.Context) (resObj *Contact, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Contact
func (r *ContactRequest) Update(ctx context.Context, reqObj *Contact) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Contact
func (r *ContactRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ContactFolderRequestBuilder is request builder for ContactFolder
type ContactFolderRequestBuilder struct{ BaseRequestBuilder }

// Request returns ContactFolderRequest
func (b *ContactFolderRequestBuilder) Request() *ContactFolderRequest {
	return &ContactFolderRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ContactFolderRequest is request for ContactFolder
type ContactFolderRequest struct{ BaseRequest }

// Get performs GET request for ContactFolder
func (r *ContactFolderRequest) Get(ctx context.Context) (resObj *ContactFolder, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ContactFolder
func (r *ContactFolderRequest) Update(ctx context.Context, reqObj *ContactFolder) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ContactFolder
func (r *ContactFolderRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
