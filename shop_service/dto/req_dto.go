package dto

import "mime/multipart"

// For Fetching products from db witch filter
type ProductFilterDto struct {
	Mark      string `json:"mark,omitempty"`
	Feature   string `json:"feature,omitempty"`
	Tag   	  string `json:"tag,omitempty"`
}

type LoginDto struct {
	Number   string `json:"number"`
	Password string `json:"password"`
}

type CreateProductReq struct {
	Name        string
	Description string
	Mark        string
	Price       uint32
	Discount    uint8
	SalesCount  uint64
	Count       uint16
	Features    []string
	Tags        []string
	ImageFile   *multipart.FileHeader
	ImageFiles  []*multipart.FileHeader
}
