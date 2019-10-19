// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EducationCategoryRequestBuilder is request builder for EducationCategory
type EducationCategoryRequestBuilder struct{ BaseRequestBuilder }

// Request returns EducationCategoryRequest
func (b *EducationCategoryRequestBuilder) Request() *EducationCategoryRequest {
	return &EducationCategoryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EducationCategoryRequest is request for EducationCategory
type EducationCategoryRequest struct{ BaseRequest }

// Do performs HTTP request for EducationCategory
func (r *EducationCategoryRequest) Do(method, path string, reqObj interface{}) (resObj *EducationCategory, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for EducationCategory
func (r *EducationCategoryRequest) Get() (*EducationCategory, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for EducationCategory
func (r *EducationCategoryRequest) Update(reqObj *EducationCategory) (*EducationCategory, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for EducationCategory
func (r *EducationCategoryRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
