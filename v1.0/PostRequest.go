// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// PostRequestBuilder is request builder for Post
type PostRequestBuilder struct{ BaseRequestBuilder }

// Request returns PostRequest
func (b *PostRequestBuilder) Request() *PostRequest {
	return &PostRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PostRequest is request for Post
type PostRequest struct{ BaseRequest }

// Do performs HTTP request for Post
func (r *PostRequest) Do(method, path string, reqObj interface{}) (resObj *Post, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for Post
func (r *PostRequest) Get() (*Post, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for Post
func (r *PostRequest) Update(reqObj *Post) (*Post, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for Post
func (r *PostRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Attachments returns request builder for Attachment collection
func (b *PostRequestBuilder) Attachments() *PostAttachmentsCollectionRequestBuilder {
	bb := &PostAttachmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/attachments"
	return bb
}

// PostAttachmentsCollectionRequestBuilder is request builder for Attachment collection
type PostAttachmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Attachment collection
func (b *PostAttachmentsCollectionRequestBuilder) Request() *PostAttachmentsCollectionRequest {
	return &PostAttachmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Attachment item
func (b *PostAttachmentsCollectionRequestBuilder) ID(id string) *AttachmentRequestBuilder {
	bb := &AttachmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PostAttachmentsCollectionRequest is request for Attachment collection
type PostAttachmentsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for Attachment collection
func (r *PostAttachmentsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *Attachment, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for Attachment collection
func (r *PostAttachmentsCollectionRequest) Paging(method, path string, obj interface{}) ([]Attachment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Attachment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []Attachment
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

// Get performs GET request for Attachment collection
func (r *PostAttachmentsCollectionRequest) Get() ([]Attachment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Attachment collection
func (r *PostAttachmentsCollectionRequest) Add(reqObj *Attachment) (*Attachment, error) {
	return r.Do("POST", "", reqObj)
}

// Extensions returns request builder for Extension collection
func (b *PostRequestBuilder) Extensions() *PostExtensionsCollectionRequestBuilder {
	bb := &PostExtensionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/extensions"
	return bb
}

// PostExtensionsCollectionRequestBuilder is request builder for Extension collection
type PostExtensionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Extension collection
func (b *PostExtensionsCollectionRequestBuilder) Request() *PostExtensionsCollectionRequest {
	return &PostExtensionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Extension item
func (b *PostExtensionsCollectionRequestBuilder) ID(id string) *ExtensionRequestBuilder {
	bb := &ExtensionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PostExtensionsCollectionRequest is request for Extension collection
type PostExtensionsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for Extension collection
func (r *PostExtensionsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *Extension, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for Extension collection
func (r *PostExtensionsCollectionRequest) Paging(method, path string, obj interface{}) ([]Extension, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Extension
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []Extension
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

// Get performs GET request for Extension collection
func (r *PostExtensionsCollectionRequest) Get() ([]Extension, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Extension collection
func (r *PostExtensionsCollectionRequest) Add(reqObj *Extension) (*Extension, error) {
	return r.Do("POST", "", reqObj)
}

// InReplyTo is navigation property
func (b *PostRequestBuilder) InReplyTo() *PostRequestBuilder {
	bb := &PostRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/inReplyTo"
	return bb
}

// MultiValueExtendedProperties returns request builder for MultiValueLegacyExtendedProperty collection
func (b *PostRequestBuilder) MultiValueExtendedProperties() *PostMultiValueExtendedPropertiesCollectionRequestBuilder {
	bb := &PostMultiValueExtendedPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/multiValueExtendedProperties"
	return bb
}

// PostMultiValueExtendedPropertiesCollectionRequestBuilder is request builder for MultiValueLegacyExtendedProperty collection
type PostMultiValueExtendedPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MultiValueLegacyExtendedProperty collection
func (b *PostMultiValueExtendedPropertiesCollectionRequestBuilder) Request() *PostMultiValueExtendedPropertiesCollectionRequest {
	return &PostMultiValueExtendedPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MultiValueLegacyExtendedProperty item
func (b *PostMultiValueExtendedPropertiesCollectionRequestBuilder) ID(id string) *MultiValueLegacyExtendedPropertyRequestBuilder {
	bb := &MultiValueLegacyExtendedPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PostMultiValueExtendedPropertiesCollectionRequest is request for MultiValueLegacyExtendedProperty collection
type PostMultiValueExtendedPropertiesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for MultiValueLegacyExtendedProperty collection
func (r *PostMultiValueExtendedPropertiesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *MultiValueLegacyExtendedProperty, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for MultiValueLegacyExtendedProperty collection
func (r *PostMultiValueExtendedPropertiesCollectionRequest) Paging(method, path string, obj interface{}) ([]MultiValueLegacyExtendedProperty, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []MultiValueLegacyExtendedProperty
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []MultiValueLegacyExtendedProperty
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

// Get performs GET request for MultiValueLegacyExtendedProperty collection
func (r *PostMultiValueExtendedPropertiesCollectionRequest) Get() ([]MultiValueLegacyExtendedProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for MultiValueLegacyExtendedProperty collection
func (r *PostMultiValueExtendedPropertiesCollectionRequest) Add(reqObj *MultiValueLegacyExtendedProperty) (*MultiValueLegacyExtendedProperty, error) {
	return r.Do("POST", "", reqObj)
}

// SingleValueExtendedProperties returns request builder for SingleValueLegacyExtendedProperty collection
func (b *PostRequestBuilder) SingleValueExtendedProperties() *PostSingleValueExtendedPropertiesCollectionRequestBuilder {
	bb := &PostSingleValueExtendedPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/singleValueExtendedProperties"
	return bb
}

// PostSingleValueExtendedPropertiesCollectionRequestBuilder is request builder for SingleValueLegacyExtendedProperty collection
type PostSingleValueExtendedPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SingleValueLegacyExtendedProperty collection
func (b *PostSingleValueExtendedPropertiesCollectionRequestBuilder) Request() *PostSingleValueExtendedPropertiesCollectionRequest {
	return &PostSingleValueExtendedPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SingleValueLegacyExtendedProperty item
func (b *PostSingleValueExtendedPropertiesCollectionRequestBuilder) ID(id string) *SingleValueLegacyExtendedPropertyRequestBuilder {
	bb := &SingleValueLegacyExtendedPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PostSingleValueExtendedPropertiesCollectionRequest is request for SingleValueLegacyExtendedProperty collection
type PostSingleValueExtendedPropertiesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for SingleValueLegacyExtendedProperty collection
func (r *PostSingleValueExtendedPropertiesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *SingleValueLegacyExtendedProperty, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for SingleValueLegacyExtendedProperty collection
func (r *PostSingleValueExtendedPropertiesCollectionRequest) Paging(method, path string, obj interface{}) ([]SingleValueLegacyExtendedProperty, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []SingleValueLegacyExtendedProperty
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []SingleValueLegacyExtendedProperty
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

// Get performs GET request for SingleValueLegacyExtendedProperty collection
func (r *PostSingleValueExtendedPropertiesCollectionRequest) Get() ([]SingleValueLegacyExtendedProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for SingleValueLegacyExtendedProperty collection
func (r *PostSingleValueExtendedPropertiesCollectionRequest) Add(reqObj *SingleValueLegacyExtendedProperty) (*SingleValueLegacyExtendedProperty, error) {
	return r.Do("POST", "", reqObj)
}
