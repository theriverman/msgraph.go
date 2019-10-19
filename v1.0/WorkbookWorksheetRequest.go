// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// WorkbookWorksheetRequestBuilder is request builder for WorkbookWorksheet
type WorkbookWorksheetRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookWorksheetRequest
func (b *WorkbookWorksheetRequestBuilder) Request() *WorkbookWorksheetRequest {
	return &WorkbookWorksheetRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookWorksheetRequest is request for WorkbookWorksheet
type WorkbookWorksheetRequest struct{ BaseRequest }

// Do performs HTTP request for WorkbookWorksheet
func (r *WorkbookWorksheetRequest) Do(method, path string, reqObj interface{}) (resObj *WorkbookWorksheet, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for WorkbookWorksheet
func (r *WorkbookWorksheetRequest) Get() (*WorkbookWorksheet, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for WorkbookWorksheet
func (r *WorkbookWorksheetRequest) Update(reqObj *WorkbookWorksheet) (*WorkbookWorksheet, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for WorkbookWorksheet
func (r *WorkbookWorksheetRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Charts returns request builder for WorkbookChart collection
func (b *WorkbookWorksheetRequestBuilder) Charts() *WorkbookWorksheetChartsCollectionRequestBuilder {
	bb := &WorkbookWorksheetChartsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/charts"
	return bb
}

// WorkbookWorksheetChartsCollectionRequestBuilder is request builder for WorkbookChart collection
type WorkbookWorksheetChartsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WorkbookChart collection
func (b *WorkbookWorksheetChartsCollectionRequestBuilder) Request() *WorkbookWorksheetChartsCollectionRequest {
	return &WorkbookWorksheetChartsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WorkbookChart item
func (b *WorkbookWorksheetChartsCollectionRequestBuilder) ID(id string) *WorkbookChartRequestBuilder {
	bb := &WorkbookChartRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WorkbookWorksheetChartsCollectionRequest is request for WorkbookChart collection
type WorkbookWorksheetChartsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for WorkbookChart collection
func (r *WorkbookWorksheetChartsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *WorkbookChart, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for WorkbookChart collection
func (r *WorkbookWorksheetChartsCollectionRequest) Paging(method, path string, obj interface{}) ([]WorkbookChart, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []WorkbookChart
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []WorkbookChart
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

// Get performs GET request for WorkbookChart collection
func (r *WorkbookWorksheetChartsCollectionRequest) Get() ([]WorkbookChart, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for WorkbookChart collection
func (r *WorkbookWorksheetChartsCollectionRequest) Add(reqObj *WorkbookChart) (*WorkbookChart, error) {
	return r.Do("POST", "", reqObj)
}

// Names returns request builder for WorkbookNamedItem collection
func (b *WorkbookWorksheetRequestBuilder) Names() *WorkbookWorksheetNamesCollectionRequestBuilder {
	bb := &WorkbookWorksheetNamesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/names"
	return bb
}

// WorkbookWorksheetNamesCollectionRequestBuilder is request builder for WorkbookNamedItem collection
type WorkbookWorksheetNamesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WorkbookNamedItem collection
func (b *WorkbookWorksheetNamesCollectionRequestBuilder) Request() *WorkbookWorksheetNamesCollectionRequest {
	return &WorkbookWorksheetNamesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WorkbookNamedItem item
func (b *WorkbookWorksheetNamesCollectionRequestBuilder) ID(id string) *WorkbookNamedItemRequestBuilder {
	bb := &WorkbookNamedItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WorkbookWorksheetNamesCollectionRequest is request for WorkbookNamedItem collection
type WorkbookWorksheetNamesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for WorkbookNamedItem collection
func (r *WorkbookWorksheetNamesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *WorkbookNamedItem, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for WorkbookNamedItem collection
func (r *WorkbookWorksheetNamesCollectionRequest) Paging(method, path string, obj interface{}) ([]WorkbookNamedItem, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []WorkbookNamedItem
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []WorkbookNamedItem
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

// Get performs GET request for WorkbookNamedItem collection
func (r *WorkbookWorksheetNamesCollectionRequest) Get() ([]WorkbookNamedItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for WorkbookNamedItem collection
func (r *WorkbookWorksheetNamesCollectionRequest) Add(reqObj *WorkbookNamedItem) (*WorkbookNamedItem, error) {
	return r.Do("POST", "", reqObj)
}

// PivotTables returns request builder for WorkbookPivotTable collection
func (b *WorkbookWorksheetRequestBuilder) PivotTables() *WorkbookWorksheetPivotTablesCollectionRequestBuilder {
	bb := &WorkbookWorksheetPivotTablesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/pivotTables"
	return bb
}

// WorkbookWorksheetPivotTablesCollectionRequestBuilder is request builder for WorkbookPivotTable collection
type WorkbookWorksheetPivotTablesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WorkbookPivotTable collection
func (b *WorkbookWorksheetPivotTablesCollectionRequestBuilder) Request() *WorkbookWorksheetPivotTablesCollectionRequest {
	return &WorkbookWorksheetPivotTablesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WorkbookPivotTable item
func (b *WorkbookWorksheetPivotTablesCollectionRequestBuilder) ID(id string) *WorkbookPivotTableRequestBuilder {
	bb := &WorkbookPivotTableRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WorkbookWorksheetPivotTablesCollectionRequest is request for WorkbookPivotTable collection
type WorkbookWorksheetPivotTablesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for WorkbookPivotTable collection
func (r *WorkbookWorksheetPivotTablesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *WorkbookPivotTable, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for WorkbookPivotTable collection
func (r *WorkbookWorksheetPivotTablesCollectionRequest) Paging(method, path string, obj interface{}) ([]WorkbookPivotTable, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []WorkbookPivotTable
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []WorkbookPivotTable
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

// Get performs GET request for WorkbookPivotTable collection
func (r *WorkbookWorksheetPivotTablesCollectionRequest) Get() ([]WorkbookPivotTable, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for WorkbookPivotTable collection
func (r *WorkbookWorksheetPivotTablesCollectionRequest) Add(reqObj *WorkbookPivotTable) (*WorkbookPivotTable, error) {
	return r.Do("POST", "", reqObj)
}

// Protection is navigation property
func (b *WorkbookWorksheetRequestBuilder) Protection() *WorkbookWorksheetProtectionRequestBuilder {
	bb := &WorkbookWorksheetProtectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/protection"
	return bb
}

// Tables returns request builder for WorkbookTable collection
func (b *WorkbookWorksheetRequestBuilder) Tables() *WorkbookWorksheetTablesCollectionRequestBuilder {
	bb := &WorkbookWorksheetTablesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/tables"
	return bb
}

// WorkbookWorksheetTablesCollectionRequestBuilder is request builder for WorkbookTable collection
type WorkbookWorksheetTablesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WorkbookTable collection
func (b *WorkbookWorksheetTablesCollectionRequestBuilder) Request() *WorkbookWorksheetTablesCollectionRequest {
	return &WorkbookWorksheetTablesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WorkbookTable item
func (b *WorkbookWorksheetTablesCollectionRequestBuilder) ID(id string) *WorkbookTableRequestBuilder {
	bb := &WorkbookTableRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WorkbookWorksheetTablesCollectionRequest is request for WorkbookTable collection
type WorkbookWorksheetTablesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for WorkbookTable collection
func (r *WorkbookWorksheetTablesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *WorkbookTable, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for WorkbookTable collection
func (r *WorkbookWorksheetTablesCollectionRequest) Paging(method, path string, obj interface{}) ([]WorkbookTable, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []WorkbookTable
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []WorkbookTable
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

// Get performs GET request for WorkbookTable collection
func (r *WorkbookWorksheetTablesCollectionRequest) Get() ([]WorkbookTable, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for WorkbookTable collection
func (r *WorkbookWorksheetTablesCollectionRequest) Add(reqObj *WorkbookTable) (*WorkbookTable, error) {
	return r.Do("POST", "", reqObj)
}
