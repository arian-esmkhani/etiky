package dto

// Getting data for cafe preview (slider, cart, etc) with pagination
type PreviewProductRes struct {
	Products []PreviewProductDto  `json:"products"`
	HasNext  bool      		 	  `json:"has_next"`
}