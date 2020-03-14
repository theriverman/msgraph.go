// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EdgeSearchEngine undocumented
type EdgeSearchEngine struct {
	// EdgeSearchEngineBase is the base model of EdgeSearchEngine
	EdgeSearchEngineBase
	// EdgeSearchEngineType Allows IT admins to set a predefined default search engine for MDM-Controlled devices.
	EdgeSearchEngineType *EdgeSearchEngineType `json:"edgeSearchEngineType,omitempty"`
}

// EdgeSearchEngineBase undocumented
type EdgeSearchEngineBase struct {
	// Object is the base model of EdgeSearchEngineBase
	Object
}

// EdgeSearchEngineCustom undocumented
type EdgeSearchEngineCustom struct {
	// EdgeSearchEngineBase is the base model of EdgeSearchEngineCustom
	EdgeSearchEngineBase
	// EdgeSearchEngineOpenSearchXMLURL Points to a https link containing the OpenSearch xml file that contains, at minimum, the short name and the URL to the search Engine.
	EdgeSearchEngineOpenSearchXMLURL *string `json:"edgeSearchEngineOpenSearchXmlUrl,omitempty"`
}
