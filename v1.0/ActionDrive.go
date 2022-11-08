// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/theriverman/msgraph.go/jsonx"
)

// DriveItemCheckinRequestParameter undocumented
type DriveItemCheckinRequestParameter struct {
	// CheckInAs undocumented
	CheckInAs *string `json:"checkInAs,omitempty"`
	// Comment undocumented
	Comment *string `json:"comment,omitempty"`
}

// DriveItemCheckoutRequestParameter undocumented
type DriveItemCheckoutRequestParameter struct {
}

// DriveItemCopyRequestParameter undocumented
type DriveItemCopyRequestParameter struct {
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// ParentReference undocumented
	ParentReference *ItemReference `json:"parentReference,omitempty"`
}

// DriveItemCreateLinkRequestParameter undocumented
type DriveItemCreateLinkRequestParameter struct {
	// Type undocumented
	Type *string `json:"type,omitempty"`
	// Scope undocumented
	Scope *string `json:"scope,omitempty"`
}

// DriveItemCreateUploadSessionRequestParameter undocumented
type DriveItemCreateUploadSessionRequestParameter struct {
	// Item undocumented
	Item *DriveItemUploadableProperties `json:"item,omitempty"`
}

// DriveItemInviteRequestParameter undocumented
type DriveItemInviteRequestParameter struct {
	// RequireSignIn undocumented
	RequireSignIn *bool `json:"requireSignIn,omitempty"`
	// Roles undocumented
	Roles []string `json:"roles,omitempty"`
	// SendInvitation undocumented
	SendInvitation *bool `json:"sendInvitation,omitempty"`
	// Message undocumented
	Message *string `json:"message,omitempty"`
	// Recipients undocumented
	Recipients []DriveRecipient `json:"recipients,omitempty"`
}

// DriveItemPreviewRequestParameter undocumented
type DriveItemPreviewRequestParameter struct {
	// Page undocumented
	Page *string `json:"page,omitempty"`
	// Zoom undocumented
	Zoom *float64 `json:"zoom,omitempty"`
}

// DriveItemVersionRestoreVersionRequestParameter undocumented
type DriveItemVersionRestoreVersionRequestParameter struct {
}

// Items returns request builder for DriveItem collection
func (b *DriveRequestBuilder) Items() *DriveItemsCollectionRequestBuilder {
	bb := &DriveItemsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/items"
	return bb
}

// DriveItemsCollectionRequestBuilder is request builder for DriveItem collection
type DriveItemsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DriveItem collection
func (b *DriveItemsCollectionRequestBuilder) Request() *DriveItemsCollectionRequest {
	return &DriveItemsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DriveItem item
func (b *DriveItemsCollectionRequestBuilder) ID(id string) *DriveItemRequestBuilder {
	bb := &DriveItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DriveItemsCollectionRequest is request for DriveItem collection
type DriveItemsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DriveItem collection
func (r *DriveItemsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DriveItem, error) {
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
	var values []DriveItem
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
			value  []DriveItem
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

// GetN performs GET request for DriveItem collection, max N pages
func (r *DriveItemsCollectionRequest) GetN(ctx context.Context, n int) ([]DriveItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DriveItem collection
func (r *DriveItemsCollectionRequest) Get(ctx context.Context) ([]DriveItem, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DriveItem collection
func (r *DriveItemsCollectionRequest) Add(ctx context.Context, reqObj *DriveItem) (resObj *DriveItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// List is navigation property
func (b *DriveRequestBuilder) List() *ListRequestBuilder {
	bb := &ListRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/list"
	return bb
}

// Root is navigation property
func (b *DriveRequestBuilder) Root() *DriveItemRequestBuilder {
	bb := &DriveItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/root"
	return bb
}

// Special returns request builder for DriveItem collection
func (b *DriveRequestBuilder) Special() *DriveSpecialCollectionRequestBuilder {
	bb := &DriveSpecialCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/special"
	return bb
}

// DriveSpecialCollectionRequestBuilder is request builder for DriveItem collection
type DriveSpecialCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DriveItem collection
func (b *DriveSpecialCollectionRequestBuilder) Request() *DriveSpecialCollectionRequest {
	return &DriveSpecialCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DriveItem item
func (b *DriveSpecialCollectionRequestBuilder) ID(id string) *DriveItemRequestBuilder {
	bb := &DriveItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DriveSpecialCollectionRequest is request for DriveItem collection
type DriveSpecialCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DriveItem collection
func (r *DriveSpecialCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DriveItem, error) {
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
	var values []DriveItem
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
			value  []DriveItem
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

// GetN performs GET request for DriveItem collection, max N pages
func (r *DriveSpecialCollectionRequest) GetN(ctx context.Context, n int) ([]DriveItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DriveItem collection
func (r *DriveSpecialCollectionRequest) Get(ctx context.Context) ([]DriveItem, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DriveItem collection
func (r *DriveSpecialCollectionRequest) Add(ctx context.Context, reqObj *DriveItem) (resObj *DriveItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Analytics is navigation property
func (b *DriveItemRequestBuilder) Analytics() *ItemAnalyticsRequestBuilder {
	bb := &ItemAnalyticsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/analytics"
	return bb
}

// Children returns request builder for DriveItem collection
func (b *DriveItemRequestBuilder) Children() *DriveItemChildrenCollectionRequestBuilder {
	bb := &DriveItemChildrenCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/children"
	return bb
}

// DriveItemChildrenCollectionRequestBuilder is request builder for DriveItem collection
type DriveItemChildrenCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DriveItem collection
func (b *DriveItemChildrenCollectionRequestBuilder) Request() *DriveItemChildrenCollectionRequest {
	return &DriveItemChildrenCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DriveItem item
func (b *DriveItemChildrenCollectionRequestBuilder) ID(id string) *DriveItemRequestBuilder {
	bb := &DriveItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DriveItemChildrenCollectionRequest is request for DriveItem collection
type DriveItemChildrenCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DriveItem collection
func (r *DriveItemChildrenCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DriveItem, error) {
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
	var values []DriveItem
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
			value  []DriveItem
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

// GetN performs GET request for DriveItem collection, max N pages
func (r *DriveItemChildrenCollectionRequest) GetN(ctx context.Context, n int) ([]DriveItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DriveItem collection
func (r *DriveItemChildrenCollectionRequest) Get(ctx context.Context) ([]DriveItem, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DriveItem collection
func (r *DriveItemChildrenCollectionRequest) Add(ctx context.Context, reqObj *DriveItem) (resObj *DriveItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ListItem is navigation property
func (b *DriveItemRequestBuilder) ListItem() *ListItemRequestBuilder {
	bb := &ListItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/listItem"
	return bb
}

// Permissions returns request builder for Permission collection
func (b *DriveItemRequestBuilder) Permissions() *DriveItemPermissionsCollectionRequestBuilder {
	bb := &DriveItemPermissionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/permissions"
	return bb
}

// DriveItemPermissionsCollectionRequestBuilder is request builder for Permission collection
type DriveItemPermissionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Permission collection
func (b *DriveItemPermissionsCollectionRequestBuilder) Request() *DriveItemPermissionsCollectionRequest {
	return &DriveItemPermissionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Permission item
func (b *DriveItemPermissionsCollectionRequestBuilder) ID(id string) *PermissionRequestBuilder {
	bb := &PermissionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DriveItemPermissionsCollectionRequest is request for Permission collection
type DriveItemPermissionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Permission collection
func (r *DriveItemPermissionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Permission, error) {
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
	var values []Permission
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
			value  []Permission
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

// GetN performs GET request for Permission collection, max N pages
func (r *DriveItemPermissionsCollectionRequest) GetN(ctx context.Context, n int) ([]Permission, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Permission collection
func (r *DriveItemPermissionsCollectionRequest) Get(ctx context.Context) ([]Permission, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Permission collection
func (r *DriveItemPermissionsCollectionRequest) Add(ctx context.Context, reqObj *Permission) (resObj *Permission, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Subscriptions returns request builder for Subscription collection
func (b *DriveItemRequestBuilder) Subscriptions() *DriveItemSubscriptionsCollectionRequestBuilder {
	bb := &DriveItemSubscriptionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/subscriptions"
	return bb
}

// DriveItemSubscriptionsCollectionRequestBuilder is request builder for Subscription collection
type DriveItemSubscriptionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Subscription collection
func (b *DriveItemSubscriptionsCollectionRequestBuilder) Request() *DriveItemSubscriptionsCollectionRequest {
	return &DriveItemSubscriptionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Subscription item
func (b *DriveItemSubscriptionsCollectionRequestBuilder) ID(id string) *SubscriptionRequestBuilder {
	bb := &SubscriptionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DriveItemSubscriptionsCollectionRequest is request for Subscription collection
type DriveItemSubscriptionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Subscription collection
func (r *DriveItemSubscriptionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Subscription, error) {
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
	var values []Subscription
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
			value  []Subscription
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

// GetN performs GET request for Subscription collection, max N pages
func (r *DriveItemSubscriptionsCollectionRequest) GetN(ctx context.Context, n int) ([]Subscription, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Subscription collection
func (r *DriveItemSubscriptionsCollectionRequest) Get(ctx context.Context) ([]Subscription, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Subscription collection
func (r *DriveItemSubscriptionsCollectionRequest) Add(ctx context.Context, reqObj *Subscription) (resObj *Subscription, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Thumbnails returns request builder for ThumbnailSet collection
func (b *DriveItemRequestBuilder) Thumbnails() *DriveItemThumbnailsCollectionRequestBuilder {
	bb := &DriveItemThumbnailsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/thumbnails"
	return bb
}

// DriveItemThumbnailsCollectionRequestBuilder is request builder for ThumbnailSet collection
type DriveItemThumbnailsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ThumbnailSet collection
func (b *DriveItemThumbnailsCollectionRequestBuilder) Request() *DriveItemThumbnailsCollectionRequest {
	return &DriveItemThumbnailsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ThumbnailSet item
func (b *DriveItemThumbnailsCollectionRequestBuilder) ID(id string) *ThumbnailSetRequestBuilder {
	bb := &ThumbnailSetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DriveItemThumbnailsCollectionRequest is request for ThumbnailSet collection
type DriveItemThumbnailsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ThumbnailSet collection
func (r *DriveItemThumbnailsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ThumbnailSet, error) {
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
	var values []ThumbnailSet
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
			value  []ThumbnailSet
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

// GetN performs GET request for ThumbnailSet collection, max N pages
func (r *DriveItemThumbnailsCollectionRequest) GetN(ctx context.Context, n int) ([]ThumbnailSet, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ThumbnailSet collection
func (r *DriveItemThumbnailsCollectionRequest) Get(ctx context.Context) ([]ThumbnailSet, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ThumbnailSet collection
func (r *DriveItemThumbnailsCollectionRequest) Add(ctx context.Context, reqObj *ThumbnailSet) (resObj *ThumbnailSet, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Versions returns request builder for DriveItemVersion collection
func (b *DriveItemRequestBuilder) Versions() *DriveItemVersionsCollectionRequestBuilder {
	bb := &DriveItemVersionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/versions"
	return bb
}

// DriveItemVersionsCollectionRequestBuilder is request builder for DriveItemVersion collection
type DriveItemVersionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DriveItemVersion collection
func (b *DriveItemVersionsCollectionRequestBuilder) Request() *DriveItemVersionsCollectionRequest {
	return &DriveItemVersionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DriveItemVersion item
func (b *DriveItemVersionsCollectionRequestBuilder) ID(id string) *DriveItemVersionRequestBuilder {
	bb := &DriveItemVersionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DriveItemVersionsCollectionRequest is request for DriveItemVersion collection
type DriveItemVersionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DriveItemVersion collection
func (r *DriveItemVersionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DriveItemVersion, error) {
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
	var values []DriveItemVersion
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
			value  []DriveItemVersion
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

// GetN performs GET request for DriveItemVersion collection, max N pages
func (r *DriveItemVersionsCollectionRequest) GetN(ctx context.Context, n int) ([]DriveItemVersion, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DriveItemVersion collection
func (r *DriveItemVersionsCollectionRequest) Get(ctx context.Context) ([]DriveItemVersion, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DriveItemVersion collection
func (r *DriveItemVersionsCollectionRequest) Add(ctx context.Context, reqObj *DriveItemVersion) (resObj *DriveItemVersion, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Workbook is navigation property
func (b *DriveItemRequestBuilder) Workbook() *WorkbookRequestBuilder {
	bb := &WorkbookRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/workbook"
	return bb
}
