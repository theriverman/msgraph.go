// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// TiIndicatorRequestBuilder is request builder for TiIndicator
type TiIndicatorRequestBuilder struct{ BaseRequestBuilder }

// Request returns TiIndicatorRequest
func (b *TiIndicatorRequestBuilder) Request() *TiIndicatorRequest {
	return &TiIndicatorRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TiIndicatorRequest is request for TiIndicator
type TiIndicatorRequest struct{ BaseRequest }

// Get performs GET request for TiIndicator
func (r *TiIndicatorRequest) Get(ctx context.Context) (resObj *TiIndicator, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TiIndicator
func (r *TiIndicatorRequest) Update(ctx context.Context, reqObj *TiIndicator) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TiIndicator
func (r *TiIndicatorRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type TiIndicatorCollectionSubmitTiIndicatorsRequestBuilder struct{ BaseRequestBuilder }

// SubmitTiIndicators action undocumented
func (b *SecurityTiIndicatorsCollectionRequestBuilder) SubmitTiIndicators(reqObj *TiIndicatorCollectionSubmitTiIndicatorsRequestParameter) *TiIndicatorCollectionSubmitTiIndicatorsRequestBuilder {
	bb := &TiIndicatorCollectionSubmitTiIndicatorsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/submitTiIndicators"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type TiIndicatorCollectionSubmitTiIndicatorsRequest struct{ BaseRequest }

//
func (b *TiIndicatorCollectionSubmitTiIndicatorsRequestBuilder) Request() *TiIndicatorCollectionSubmitTiIndicatorsRequest {
	return &TiIndicatorCollectionSubmitTiIndicatorsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *TiIndicatorCollectionSubmitTiIndicatorsRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TiIndicator, error) {
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
	var values []TiIndicator
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
			value  []TiIndicator
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

//
func (r *TiIndicatorCollectionSubmitTiIndicatorsRequest) PostN(ctx context.Context, n int) ([]TiIndicator, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *TiIndicatorCollectionSubmitTiIndicatorsRequest) Post(ctx context.Context) ([]TiIndicator, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

//
type TiIndicatorCollectionUpdateTiIndicatorsRequestBuilder struct{ BaseRequestBuilder }

// UpdateTiIndicators action undocumented
func (b *SecurityTiIndicatorsCollectionRequestBuilder) UpdateTiIndicators(reqObj *TiIndicatorCollectionUpdateTiIndicatorsRequestParameter) *TiIndicatorCollectionUpdateTiIndicatorsRequestBuilder {
	bb := &TiIndicatorCollectionUpdateTiIndicatorsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/updateTiIndicators"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type TiIndicatorCollectionUpdateTiIndicatorsRequest struct{ BaseRequest }

//
func (b *TiIndicatorCollectionUpdateTiIndicatorsRequestBuilder) Request() *TiIndicatorCollectionUpdateTiIndicatorsRequest {
	return &TiIndicatorCollectionUpdateTiIndicatorsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *TiIndicatorCollectionUpdateTiIndicatorsRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TiIndicator, error) {
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
	var values []TiIndicator
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
			value  []TiIndicator
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

//
func (r *TiIndicatorCollectionUpdateTiIndicatorsRequest) PostN(ctx context.Context, n int) ([]TiIndicator, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *TiIndicatorCollectionUpdateTiIndicatorsRequest) Post(ctx context.Context) ([]TiIndicator, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

//
type TiIndicatorCollectionDeleteTiIndicatorsRequestBuilder struct{ BaseRequestBuilder }

// DeleteTiIndicators action undocumented
func (b *SecurityTiIndicatorsCollectionRequestBuilder) DeleteTiIndicators(reqObj *TiIndicatorCollectionDeleteTiIndicatorsRequestParameter) *TiIndicatorCollectionDeleteTiIndicatorsRequestBuilder {
	bb := &TiIndicatorCollectionDeleteTiIndicatorsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/deleteTiIndicators"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type TiIndicatorCollectionDeleteTiIndicatorsRequest struct{ BaseRequest }

//
func (b *TiIndicatorCollectionDeleteTiIndicatorsRequestBuilder) Request() *TiIndicatorCollectionDeleteTiIndicatorsRequest {
	return &TiIndicatorCollectionDeleteTiIndicatorsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *TiIndicatorCollectionDeleteTiIndicatorsRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ResultInfo, error) {
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
	var values []ResultInfo
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
			value  []ResultInfo
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

//
func (r *TiIndicatorCollectionDeleteTiIndicatorsRequest) PostN(ctx context.Context, n int) ([]ResultInfo, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *TiIndicatorCollectionDeleteTiIndicatorsRequest) Post(ctx context.Context) ([]ResultInfo, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

//
type TiIndicatorCollectionDeleteTiIndicatorsByExternalIDRequestBuilder struct{ BaseRequestBuilder }

// DeleteTiIndicatorsByExternalID action undocumented
func (b *SecurityTiIndicatorsCollectionRequestBuilder) DeleteTiIndicatorsByExternalID(reqObj *TiIndicatorCollectionDeleteTiIndicatorsByExternalIDRequestParameter) *TiIndicatorCollectionDeleteTiIndicatorsByExternalIDRequestBuilder {
	bb := &TiIndicatorCollectionDeleteTiIndicatorsByExternalIDRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/deleteTiIndicatorsByExternalId"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type TiIndicatorCollectionDeleteTiIndicatorsByExternalIDRequest struct{ BaseRequest }

//
func (b *TiIndicatorCollectionDeleteTiIndicatorsByExternalIDRequestBuilder) Request() *TiIndicatorCollectionDeleteTiIndicatorsByExternalIDRequest {
	return &TiIndicatorCollectionDeleteTiIndicatorsByExternalIDRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *TiIndicatorCollectionDeleteTiIndicatorsByExternalIDRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ResultInfo, error) {
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
	var values []ResultInfo
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
			value  []ResultInfo
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

//
func (r *TiIndicatorCollectionDeleteTiIndicatorsByExternalIDRequest) PostN(ctx context.Context, n int) ([]ResultInfo, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *TiIndicatorCollectionDeleteTiIndicatorsByExternalIDRequest) Post(ctx context.Context) ([]ResultInfo, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}
