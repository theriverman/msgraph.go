// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// IOSLobAppProvisioningConfigurationCollectionHasPayloadLinksRequestParameter undocumented
type IOSLobAppProvisioningConfigurationCollectionHasPayloadLinksRequestParameter struct {
	// PayloadIDs undocumented
	PayloadIDs []string `json:"payloadIds,omitempty"`
}

// IOSLobAppProvisioningConfigurationAssignRequestParameter undocumented
type IOSLobAppProvisioningConfigurationAssignRequestParameter struct {
	// AppProvisioningConfigurationGroupAssignments undocumented
	AppProvisioningConfigurationGroupAssignments []MobileAppProvisioningConfigGroupAssignment `json:"appProvisioningConfigurationGroupAssignments,omitempty"`
	// IOSLobAppProvisioningConfigAssignments undocumented
	IOSLobAppProvisioningConfigAssignments []IOSLobAppProvisioningConfigurationAssignment `json:"iOSLobAppProvisioningConfigAssignments,omitempty"`
}

//
type IOSLobAppProvisioningConfigurationCollectionHasPayloadLinksRequestBuilder struct{ BaseRequestBuilder }

// HasPayloadLinks action undocumented
func (b *DeviceAppManagementIOSLobAppProvisioningConfigurationsCollectionRequestBuilder) HasPayloadLinks(reqObj *IOSLobAppProvisioningConfigurationCollectionHasPayloadLinksRequestParameter) *IOSLobAppProvisioningConfigurationCollectionHasPayloadLinksRequestBuilder {
	bb := &IOSLobAppProvisioningConfigurationCollectionHasPayloadLinksRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/hasPayloadLinks"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type IOSLobAppProvisioningConfigurationCollectionHasPayloadLinksRequest struct{ BaseRequest }

//
func (b *IOSLobAppProvisioningConfigurationCollectionHasPayloadLinksRequestBuilder) Request() *IOSLobAppProvisioningConfigurationCollectionHasPayloadLinksRequest {
	return &IOSLobAppProvisioningConfigurationCollectionHasPayloadLinksRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *IOSLobAppProvisioningConfigurationCollectionHasPayloadLinksRequest) Do(method, path string, reqObj interface{}) (resObj *[]HasPayloadLinkResultItem, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

//
func (r *IOSLobAppProvisioningConfigurationCollectionHasPayloadLinksRequest) Paging(method, path string, obj interface{}) ([][]HasPayloadLinkResultItem, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]HasPayloadLinkResultItem
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]HasPayloadLinkResultItem
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
func (r *IOSLobAppProvisioningConfigurationCollectionHasPayloadLinksRequest) Get() ([][]HasPayloadLinkResultItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

//
type IOSLobAppProvisioningConfigurationAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *IOSLobAppProvisioningConfigurationRequestBuilder) Assign(reqObj *IOSLobAppProvisioningConfigurationAssignRequestParameter) *IOSLobAppProvisioningConfigurationAssignRequestBuilder {
	bb := &IOSLobAppProvisioningConfigurationAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type IOSLobAppProvisioningConfigurationAssignRequest struct{ BaseRequest }

//
func (b *IOSLobAppProvisioningConfigurationAssignRequestBuilder) Request() *IOSLobAppProvisioningConfigurationAssignRequest {
	return &IOSLobAppProvisioningConfigurationAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *IOSLobAppProvisioningConfigurationAssignRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *IOSLobAppProvisioningConfigurationAssignRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}
