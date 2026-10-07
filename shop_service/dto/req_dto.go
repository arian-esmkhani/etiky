package dto

// For Fetching products from db witch filter
type ProductFilterDto struct {
	Mark      string `json:"mark,omitempty"`
	Features  string `json:"features,omitempty"`
	Tags   	  string `json:"tags,omitempty"`
}