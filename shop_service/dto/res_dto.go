package dto

// Getting data for cafe preview (slider, cart, etc) with pagination
type PreviewProductRes struct {
	Products []PreviewProductDto  `json:"products"`
	HasNext  bool      		 	  `json:"has_next"`
}

type CartRes struct {
	CartDto    []CartDto `json:"cart"`
	FinalPrice uint32    `json:"final_price"`
}