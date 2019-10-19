// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// EducationSynchronizationProfileResumeRequestParameter undocumented
type EducationSynchronizationProfileResumeRequestParameter struct {
}

// EducationSynchronizationProfilePauseRequestParameter undocumented
type EducationSynchronizationProfilePauseRequestParameter struct {
}

// EducationSynchronizationProfileResetRequestParameter undocumented
type EducationSynchronizationProfileResetRequestParameter struct {
}

// EducationSynchronizationProfileStartRequestParameter undocumented
type EducationSynchronizationProfileStartRequestParameter struct {
}

//
type EducationSynchronizationProfileResumeRequestBuilder struct{ BaseRequestBuilder }

// Resume action undocumented
func (b *EducationSynchronizationProfileRequestBuilder) Resume(reqObj *EducationSynchronizationProfileResumeRequestParameter) *EducationSynchronizationProfileResumeRequestBuilder {
	bb := &EducationSynchronizationProfileResumeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/resume"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type EducationSynchronizationProfileResumeRequest struct{ BaseRequest }

//
func (b *EducationSynchronizationProfileResumeRequestBuilder) Request() *EducationSynchronizationProfileResumeRequest {
	return &EducationSynchronizationProfileResumeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *EducationSynchronizationProfileResumeRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *EducationSynchronizationProfileResumeRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type EducationSynchronizationProfilePauseRequestBuilder struct{ BaseRequestBuilder }

// Pause action undocumented
func (b *EducationSynchronizationProfileRequestBuilder) Pause(reqObj *EducationSynchronizationProfilePauseRequestParameter) *EducationSynchronizationProfilePauseRequestBuilder {
	bb := &EducationSynchronizationProfilePauseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/pause"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type EducationSynchronizationProfilePauseRequest struct{ BaseRequest }

//
func (b *EducationSynchronizationProfilePauseRequestBuilder) Request() *EducationSynchronizationProfilePauseRequest {
	return &EducationSynchronizationProfilePauseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *EducationSynchronizationProfilePauseRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *EducationSynchronizationProfilePauseRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type EducationSynchronizationProfileResetRequestBuilder struct{ BaseRequestBuilder }

// Reset action undocumented
func (b *EducationSynchronizationProfileRequestBuilder) Reset(reqObj *EducationSynchronizationProfileResetRequestParameter) *EducationSynchronizationProfileResetRequestBuilder {
	bb := &EducationSynchronizationProfileResetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/reset"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type EducationSynchronizationProfileResetRequest struct{ BaseRequest }

//
func (b *EducationSynchronizationProfileResetRequestBuilder) Request() *EducationSynchronizationProfileResetRequest {
	return &EducationSynchronizationProfileResetRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *EducationSynchronizationProfileResetRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *EducationSynchronizationProfileResetRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type EducationSynchronizationProfileStartRequestBuilder struct{ BaseRequestBuilder }

// Start action undocumented
func (b *EducationSynchronizationProfileRequestBuilder) Start(reqObj *EducationSynchronizationProfileStartRequestParameter) *EducationSynchronizationProfileStartRequestBuilder {
	bb := &EducationSynchronizationProfileStartRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/start"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type EducationSynchronizationProfileStartRequest struct{ BaseRequest }

//
func (b *EducationSynchronizationProfileStartRequestBuilder) Request() *EducationSynchronizationProfileStartRequest {
	return &EducationSynchronizationProfileStartRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *EducationSynchronizationProfileStartRequest) Do(method, path string, reqObj interface{}) (resObj *[]EducationFileSynchronizationVerificationMessage, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

//
func (r *EducationSynchronizationProfileStartRequest) Paging(method, path string, obj interface{}) ([][]EducationFileSynchronizationVerificationMessage, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]EducationFileSynchronizationVerificationMessage
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]EducationFileSynchronizationVerificationMessage
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

//
func (r *EducationSynchronizationProfileStartRequest) Get() ([][]EducationFileSynchronizationVerificationMessage, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}
