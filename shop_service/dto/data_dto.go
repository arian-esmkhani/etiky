package dto

// Getting data for cafe preview (slider, cart, etc)
type PreviewProductDto struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	ImgUrl      string `json:"img_url"`
	Price       uint32 `json:"price"`
	Discount    uint8  `json:"discount"`
	FinalPrice  uint32 `json:"final_price"`
}