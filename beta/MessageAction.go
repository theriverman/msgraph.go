// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MessageCreateReplyRequestParameter undocumented
type MessageCreateReplyRequestParameter struct {
	// Message undocumented
	Message *Message `json:"Message,omitempty"`
	// Comment undocumented
	Comment *string `json:"Comment,omitempty"`
}

// MessageCreateReplyAllRequestParameter undocumented
type MessageCreateReplyAllRequestParameter struct {
	// Message undocumented
	Message *Message `json:"Message,omitempty"`
	// Comment undocumented
	Comment *string `json:"Comment,omitempty"`
}

// MessageCreateForwardRequestParameter undocumented
type MessageCreateForwardRequestParameter struct {
	// ToRecipients undocumented
	ToRecipients []Recipient `json:"ToRecipients,omitempty"`
	// Message undocumented
	Message *Message `json:"Message,omitempty"`
	// Comment undocumented
	Comment *string `json:"Comment,omitempty"`
}

// MessageSendRequestParameter undocumented
type MessageSendRequestParameter struct {
}

// MessageCopyRequestParameter undocumented
type MessageCopyRequestParameter struct {
	// DestinationID undocumented
	DestinationID *string `json:"DestinationId,omitempty"`
}

// MessageMoveRequestParameter undocumented
type MessageMoveRequestParameter struct {
	// DestinationID undocumented
	DestinationID *string `json:"DestinationId,omitempty"`
}

// MessageReplyRequestParameter undocumented
type MessageReplyRequestParameter struct {
	// Message undocumented
	Message *Message `json:"Message,omitempty"`
	// Comment undocumented
	Comment *string `json:"Comment,omitempty"`
}

// MessageReplyAllRequestParameter undocumented
type MessageReplyAllRequestParameter struct {
	// Message undocumented
	Message *Message `json:"Message,omitempty"`
	// Comment undocumented
	Comment *string `json:"Comment,omitempty"`
}

// MessageForwardRequestParameter undocumented
type MessageForwardRequestParameter struct {
	// ToRecipients undocumented
	ToRecipients []Recipient `json:"ToRecipients,omitempty"`
	// Message undocumented
	Message *Message `json:"Message,omitempty"`
	// Comment undocumented
	Comment *string `json:"Comment,omitempty"`
}

// MessageUnsubscribeRequestParameter undocumented
type MessageUnsubscribeRequestParameter struct {
}

//
type MessageCreateReplyRequestBuilder struct{ BaseRequestBuilder }

// CreateReply action undocumented
func (b *MessageRequestBuilder) CreateReply(reqObj *MessageCreateReplyRequestParameter) *MessageCreateReplyRequestBuilder {
	bb := &MessageCreateReplyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/createReply"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type MessageCreateReplyRequest struct{ BaseRequest }

//
func (b *MessageCreateReplyRequestBuilder) Request() *MessageCreateReplyRequest {
	return &MessageCreateReplyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *MessageCreateReplyRequest) Do(method, path string, reqObj interface{}) (resObj *Message, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

//
func (r *MessageCreateReplyRequest) Post() (*Message, error) {
	return r.Do("POST", "", r.requestObject)
}

//
type MessageCreateReplyAllRequestBuilder struct{ BaseRequestBuilder }

// CreateReplyAll action undocumented
func (b *MessageRequestBuilder) CreateReplyAll(reqObj *MessageCreateReplyAllRequestParameter) *MessageCreateReplyAllRequestBuilder {
	bb := &MessageCreateReplyAllRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/createReplyAll"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type MessageCreateReplyAllRequest struct{ BaseRequest }

//
func (b *MessageCreateReplyAllRequestBuilder) Request() *MessageCreateReplyAllRequest {
	return &MessageCreateReplyAllRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *MessageCreateReplyAllRequest) Do(method, path string, reqObj interface{}) (resObj *Message, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

//
func (r *MessageCreateReplyAllRequest) Post() (*Message, error) {
	return r.Do("POST", "", r.requestObject)
}

//
type MessageCreateForwardRequestBuilder struct{ BaseRequestBuilder }

// CreateForward action undocumented
func (b *MessageRequestBuilder) CreateForward(reqObj *MessageCreateForwardRequestParameter) *MessageCreateForwardRequestBuilder {
	bb := &MessageCreateForwardRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/createForward"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type MessageCreateForwardRequest struct{ BaseRequest }

//
func (b *MessageCreateForwardRequestBuilder) Request() *MessageCreateForwardRequest {
	return &MessageCreateForwardRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *MessageCreateForwardRequest) Do(method, path string, reqObj interface{}) (resObj *Message, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

//
func (r *MessageCreateForwardRequest) Post() (*Message, error) {
	return r.Do("POST", "", r.requestObject)
}

//
type MessageSendRequestBuilder struct{ BaseRequestBuilder }

// Send action undocumented
func (b *MessageRequestBuilder) Send(reqObj *MessageSendRequestParameter) *MessageSendRequestBuilder {
	bb := &MessageSendRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/send"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type MessageSendRequest struct{ BaseRequest }

//
func (b *MessageSendRequestBuilder) Request() *MessageSendRequest {
	return &MessageSendRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *MessageSendRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *MessageSendRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type MessageCopyRequestBuilder struct{ BaseRequestBuilder }

// Copy action undocumented
func (b *MessageRequestBuilder) Copy(reqObj *MessageCopyRequestParameter) *MessageCopyRequestBuilder {
	bb := &MessageCopyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/copy"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type MessageCopyRequest struct{ BaseRequest }

//
func (b *MessageCopyRequestBuilder) Request() *MessageCopyRequest {
	return &MessageCopyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *MessageCopyRequest) Do(method, path string, reqObj interface{}) (resObj *Message, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

//
func (r *MessageCopyRequest) Post() (*Message, error) {
	return r.Do("POST", "", r.requestObject)
}

//
type MessageMoveRequestBuilder struct{ BaseRequestBuilder }

// Move action undocumented
func (b *MessageRequestBuilder) Move(reqObj *MessageMoveRequestParameter) *MessageMoveRequestBuilder {
	bb := &MessageMoveRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/move"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type MessageMoveRequest struct{ BaseRequest }

//
func (b *MessageMoveRequestBuilder) Request() *MessageMoveRequest {
	return &MessageMoveRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *MessageMoveRequest) Do(method, path string, reqObj interface{}) (resObj *Message, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

//
func (r *MessageMoveRequest) Post() (*Message, error) {
	return r.Do("POST", "", r.requestObject)
}

//
type MessageReplyRequestBuilder struct{ BaseRequestBuilder }

// Reply action undocumented
func (b *MessageRequestBuilder) Reply(reqObj *MessageReplyRequestParameter) *MessageReplyRequestBuilder {
	bb := &MessageReplyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/reply"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type MessageReplyRequest struct{ BaseRequest }

//
func (b *MessageReplyRequestBuilder) Request() *MessageReplyRequest {
	return &MessageReplyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *MessageReplyRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *MessageReplyRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type MessageReplyAllRequestBuilder struct{ BaseRequestBuilder }

// ReplyAll action undocumented
func (b *MessageRequestBuilder) ReplyAll(reqObj *MessageReplyAllRequestParameter) *MessageReplyAllRequestBuilder {
	bb := &MessageReplyAllRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/replyAll"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type MessageReplyAllRequest struct{ BaseRequest }

//
func (b *MessageReplyAllRequestBuilder) Request() *MessageReplyAllRequest {
	return &MessageReplyAllRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *MessageReplyAllRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *MessageReplyAllRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type MessageForwardRequestBuilder struct{ BaseRequestBuilder }

// Forward action undocumented
func (b *MessageRequestBuilder) Forward(reqObj *MessageForwardRequestParameter) *MessageForwardRequestBuilder {
	bb := &MessageForwardRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/forward"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type MessageForwardRequest struct{ BaseRequest }

//
func (b *MessageForwardRequestBuilder) Request() *MessageForwardRequest {
	return &MessageForwardRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *MessageForwardRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *MessageForwardRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type MessageUnsubscribeRequestBuilder struct{ BaseRequestBuilder }

// Unsubscribe action undocumented
func (b *MessageRequestBuilder) Unsubscribe(reqObj *MessageUnsubscribeRequestParameter) *MessageUnsubscribeRequestBuilder {
	bb := &MessageUnsubscribeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/unsubscribe"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type MessageUnsubscribeRequest struct{ BaseRequest }

//
func (b *MessageUnsubscribeRequestBuilder) Request() *MessageUnsubscribeRequest {
	return &MessageUnsubscribeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *MessageUnsubscribeRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *MessageUnsubscribeRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}
