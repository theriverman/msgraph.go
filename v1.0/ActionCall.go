// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// CallAnswerRequestParameter undocumented
type CallAnswerRequestParameter struct {
	// CallbackURI undocumented
	CallbackURI *string `json:"callbackUri,omitempty"`
	// MediaConfig undocumented
	MediaConfig *MediaConfig `json:"mediaConfig,omitempty"`
	// AcceptedModalities undocumented
	AcceptedModalities []Modality `json:"acceptedModalities,omitempty"`
}

// CallChangeScreenSharingRoleRequestParameter undocumented
type CallChangeScreenSharingRoleRequestParameter struct {
	// Role undocumented
	Role *ScreenSharingRole `json:"role,omitempty"`
}

// CallKeepAliveRequestParameter undocumented
type CallKeepAliveRequestParameter struct {
}

// CallMuteRequestParameter undocumented
type CallMuteRequestParameter struct {
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
}

// CallPlayPromptRequestParameter undocumented
type CallPlayPromptRequestParameter struct {
	// Prompts undocumented
	Prompts []Prompt `json:"prompts,omitempty"`
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
}

// CallRecordResponseRequestParameter undocumented
type CallRecordResponseRequestParameter struct {
	// Prompts undocumented
	Prompts []Prompt `json:"prompts,omitempty"`
	// BargeInAllowed undocumented
	BargeInAllowed *bool `json:"bargeInAllowed,omitempty"`
	// InitialSilenceTimeoutInSeconds undocumented
	InitialSilenceTimeoutInSeconds *int `json:"initialSilenceTimeoutInSeconds,omitempty"`
	// MaxSilenceTimeoutInSeconds undocumented
	MaxSilenceTimeoutInSeconds *int `json:"maxSilenceTimeoutInSeconds,omitempty"`
	// MaxRecordDurationInSeconds undocumented
	MaxRecordDurationInSeconds *int `json:"maxRecordDurationInSeconds,omitempty"`
	// PlayBeep undocumented
	PlayBeep *bool `json:"playBeep,omitempty"`
	// StopTones undocumented
	StopTones []string `json:"stopTones,omitempty"`
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
}

// CallRedirectRequestParameter undocumented
type CallRedirectRequestParameter struct {
	// Targets undocumented
	Targets []InvitationParticipantInfo `json:"targets,omitempty"`
	// Timeout undocumented
	Timeout *int `json:"timeout,omitempty"`
	// CallbackURI undocumented
	CallbackURI *string `json:"callbackUri,omitempty"`
}

// CallRejectRequestParameter undocumented
type CallRejectRequestParameter struct {
	// Reason undocumented
	Reason *RejectReason `json:"reason,omitempty"`
	// CallbackURI undocumented
	CallbackURI *string `json:"callbackUri,omitempty"`
}

// CallSubscribeToToneRequestParameter undocumented
type CallSubscribeToToneRequestParameter struct {
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
}

// CallTransferRequestParameter undocumented
type CallTransferRequestParameter struct {
	// TransferTarget undocumented
	TransferTarget *InvitationParticipantInfo `json:"transferTarget,omitempty"`
}

// CallUnmuteRequestParameter undocumented
type CallUnmuteRequestParameter struct {
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
}

// Operations returns request builder for CommsOperation collection
func (b *CallRequestBuilder) Operations() *CallOperationsCollectionRequestBuilder {
	bb := &CallOperationsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/operations"
	return bb
}

// CallOperationsCollectionRequestBuilder is request builder for CommsOperation collection
type CallOperationsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for CommsOperation collection
func (b *CallOperationsCollectionRequestBuilder) Request() *CallOperationsCollectionRequest {
	return &CallOperationsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for CommsOperation item
func (b *CallOperationsCollectionRequestBuilder) ID(id string) *CommsOperationRequestBuilder {
	bb := &CommsOperationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CallOperationsCollectionRequest is request for CommsOperation collection
type CallOperationsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for CommsOperation collection
func (r *CallOperationsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]CommsOperation, error) {
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
	var values []CommsOperation
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []CommsOperation
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
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

// GetN performs GET request for CommsOperation collection, max N pages
func (r *CallOperationsCollectionRequest) GetN(ctx context.Context, n int) ([]CommsOperation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for CommsOperation collection
func (r *CallOperationsCollectionRequest) Get(ctx context.Context) ([]CommsOperation, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for CommsOperation collection
func (r *CallOperationsCollectionRequest) Add(ctx context.Context, reqObj *CommsOperation) (resObj *CommsOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Participants returns request builder for Participant collection
func (b *CallRequestBuilder) Participants() *CallParticipantsCollectionRequestBuilder {
	bb := &CallParticipantsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/participants"
	return bb
}

// CallParticipantsCollectionRequestBuilder is request builder for Participant collection
type CallParticipantsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Participant collection
func (b *CallParticipantsCollectionRequestBuilder) Request() *CallParticipantsCollectionRequest {
	return &CallParticipantsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Participant item
func (b *CallParticipantsCollectionRequestBuilder) ID(id string) *ParticipantRequestBuilder {
	bb := &ParticipantRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CallParticipantsCollectionRequest is request for Participant collection
type CallParticipantsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Participant collection
func (r *CallParticipantsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Participant, error) {
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
	var values []Participant
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []Participant
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
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

// GetN performs GET request for Participant collection, max N pages
func (r *CallParticipantsCollectionRequest) GetN(ctx context.Context, n int) ([]Participant, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Participant collection
func (r *CallParticipantsCollectionRequest) Get(ctx context.Context) ([]Participant, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Participant collection
func (r *CallParticipantsCollectionRequest) Add(ctx context.Context, reqObj *Participant) (resObj *Participant, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
