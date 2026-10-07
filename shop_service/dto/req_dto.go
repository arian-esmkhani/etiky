package dto

// For Fetching products from db witch filter
type ProductFilterDto struct {
	Mark      string `json:"mark"`
	Features  string `json:"features"`
	Tags   	  string `json:"tags"`
}