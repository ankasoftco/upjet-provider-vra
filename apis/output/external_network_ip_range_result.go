package models 
type ExternalNetworkIPRangeResult struct {

	// List of content items
	// Read Only: true
	Content []*ExternalNetworkIPRange `json:"content"`

	// Number of elements in the current page
	// Example: 1
	// Read Only: true
	NumberOfElements int64 `json:"numberOfElements,omitempty"`

	// Total number of elements. In some cases the field may not be populated
	// Example: 1
	// Read Only: true
	TotalElements int64 `json:"totalElements,omitempty"`
}

