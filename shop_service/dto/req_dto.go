package dto

// For Fetching products from db witch filter
type ProductFilterDto struct {
	Mark      string `json:"mark,omitempty"`
	Feature   string `json:"feature,omitempty"`
	Tag   	  string `json:"tag,omitempty"`
}