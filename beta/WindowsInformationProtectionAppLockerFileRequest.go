// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsInformationProtectionAppLockerFileRequestBuilder is request builder for WindowsInformationProtectionAppLockerFile
type WindowsInformationProtectionAppLockerFileRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsInformationProtectionAppLockerFileRequest
func (b *WindowsInformationProtectionAppLockerFileRequestBuilder) Request() *WindowsInformationProtectionAppLockerFileRequest {
	return &WindowsInformationProtectionAppLockerFileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsInformationProtectionAppLockerFileRequest is request for WindowsInformationProtectionAppLockerFile
type WindowsInformationProtectionAppLockerFileRequest struct{ BaseRequest }

// Do performs HTTP request for WindowsInformationProtectionAppLockerFile
func (r *WindowsInformationProtectionAppLockerFileRequest) Do(method, path string, reqObj interface{}) (resObj *WindowsInformationProtectionAppLockerFile, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for WindowsInformationProtectionAppLockerFile
func (r *WindowsInformationProtectionAppLockerFileRequest) Get() (*WindowsInformationProtectionAppLockerFile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for WindowsInformationProtectionAppLockerFile
func (r *WindowsInformationProtectionAppLockerFileRequest) Update(reqObj *WindowsInformationProtectionAppLockerFile) (*WindowsInformationProtectionAppLockerFile, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for WindowsInformationProtectionAppLockerFile
func (r *WindowsInformationProtectionAppLockerFileRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
