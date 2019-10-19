// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DriveItemRequestBuilder is request builder for DriveItem
type DriveItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns DriveItemRequest
func (b *DriveItemRequestBuilder) Request() *DriveItemRequest {
	return &DriveItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DriveItemRequest is request for DriveItem
type DriveItemRequest struct{ BaseRequest }

// Do performs HTTP request for DriveItem
func (r *DriveItemRequest) Do(method, path string, reqObj interface{}) (resObj *DriveItem, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for DriveItem
func (r *DriveItemRequest) Get() (*DriveItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for DriveItem
func (r *DriveItemRequest) Update(reqObj *DriveItem) (*DriveItem, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for DriveItem
func (r *DriveItemRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Activities returns request builder for ItemActivityOLD collection
func (b *DriveItemRequestBuilder) Activities() *DriveItemActivitiesCollectionRequestBuilder {
	bb := &DriveItemActivitiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/activities"
	return bb
}

// DriveItemActivitiesCollectionRequestBuilder is request builder for ItemActivityOLD collection
type DriveItemActivitiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ItemActivityOLD collection
func (b *DriveItemActivitiesCollectionRequestBuilder) Request() *DriveItemActivitiesCollectionRequest {
	return &DriveItemActivitiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ItemActivityOLD item
func (b *DriveItemActivitiesCollectionRequestBuilder) ID(id string) *ItemActivityOLDRequestBuilder {
	bb := &ItemActivityOLDRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DriveItemActivitiesCollectionRequest is request for ItemActivityOLD collection
type DriveItemActivitiesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for ItemActivityOLD collection
func (r *DriveItemActivitiesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *ItemActivityOLD, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for ItemActivityOLD collection
func (r *DriveItemActivitiesCollectionRequest) Paging(method, path string, obj interface{}) ([]ItemActivityOLD, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ItemActivityOLD
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []ItemActivityOLD
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

// Get performs GET request for ItemActivityOLD collection
func (r *DriveItemActivitiesCollectionRequest) Get() ([]ItemActivityOLD, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ItemActivityOLD collection
func (r *DriveItemActivitiesCollectionRequest) Add(reqObj *ItemActivityOLD) (*ItemActivityOLD, error) {
	return r.Do("POST", "", reqObj)
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

// Do performs HTTP request for DriveItem collection
func (r *DriveItemChildrenCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *DriveItem, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for DriveItem collection
func (r *DriveItemChildrenCollectionRequest) Paging(method, path string, obj interface{}) ([]DriveItem, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DriveItem
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DriveItem
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

// Get performs GET request for DriveItem collection
func (r *DriveItemChildrenCollectionRequest) Get() ([]DriveItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DriveItem collection
func (r *DriveItemChildrenCollectionRequest) Add(reqObj *DriveItem) (*DriveItem, error) {
	return r.Do("POST", "", reqObj)
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

// Do performs HTTP request for Permission collection
func (r *DriveItemPermissionsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *Permission, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for Permission collection
func (r *DriveItemPermissionsCollectionRequest) Paging(method, path string, obj interface{}) ([]Permission, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Permission
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []Permission
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

// Get performs GET request for Permission collection
func (r *DriveItemPermissionsCollectionRequest) Get() ([]Permission, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Permission collection
func (r *DriveItemPermissionsCollectionRequest) Add(reqObj *Permission) (*Permission, error) {
	return r.Do("POST", "", reqObj)
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

// Do performs HTTP request for Subscription collection
func (r *DriveItemSubscriptionsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *Subscription, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for Subscription collection
func (r *DriveItemSubscriptionsCollectionRequest) Paging(method, path string, obj interface{}) ([]Subscription, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Subscription
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []Subscription
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

// Get performs GET request for Subscription collection
func (r *DriveItemSubscriptionsCollectionRequest) Get() ([]Subscription, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Subscription collection
func (r *DriveItemSubscriptionsCollectionRequest) Add(reqObj *Subscription) (*Subscription, error) {
	return r.Do("POST", "", reqObj)
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

// Do performs HTTP request for ThumbnailSet collection
func (r *DriveItemThumbnailsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *ThumbnailSet, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for ThumbnailSet collection
func (r *DriveItemThumbnailsCollectionRequest) Paging(method, path string, obj interface{}) ([]ThumbnailSet, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ThumbnailSet
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []ThumbnailSet
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

// Get performs GET request for ThumbnailSet collection
func (r *DriveItemThumbnailsCollectionRequest) Get() ([]ThumbnailSet, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ThumbnailSet collection
func (r *DriveItemThumbnailsCollectionRequest) Add(reqObj *ThumbnailSet) (*ThumbnailSet, error) {
	return r.Do("POST", "", reqObj)
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

// Do performs HTTP request for DriveItemVersion collection
func (r *DriveItemVersionsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *DriveItemVersion, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for DriveItemVersion collection
func (r *DriveItemVersionsCollectionRequest) Paging(method, path string, obj interface{}) ([]DriveItemVersion, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DriveItemVersion
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DriveItemVersion
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

// Get performs GET request for DriveItemVersion collection
func (r *DriveItemVersionsCollectionRequest) Get() ([]DriveItemVersion, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DriveItemVersion collection
func (r *DriveItemVersionsCollectionRequest) Add(reqObj *DriveItemVersion) (*DriveItemVersion, error) {
	return r.Do("POST", "", reqObj)
}

// Workbook is navigation property
func (b *DriveItemRequestBuilder) Workbook() *WorkbookRequestBuilder {
	bb := &WorkbookRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/workbook"
	return bb
}
