package dto

import (
	"time"
)

// Getting data for cafe preview (slider, cart, etc)
type PreviewProductDto struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	ImgUrl      string `json:"img_url"`
	Price       uint32 `json:"price"`
	Discount    uint8  `json:"discount"`
	FinalPrice  uint32 `json:"final_price"`
}

// Finding Article data for pages
type ArticleDto struct {
	ID               uint64       `json:"id"`
	Title            string       `json:"title"`
	Description      string       `json:"description"`
	ViewsCount       uint64       `json:"views_count"`
	AuthorID         *uint64      `json:"author_id,omitempty"`
	Status           string       `json:"status"`
	ImgUrl           string       `json:"img_url,omitempty"`
	VideoUrl         string       `json:"video_url,omitempty"`
	Tags         	 []string     `json:"tags,omitempty"`
	// SEO
	Slug             string       `json:"slug"`
    MetaTitle        string       `json:"meta_title"`
    MetaDescription  string       `json:"meta_description"`
    MetaRobots       string       `json:"meta_robots"`
	CanonicalMode    string       `json:"canonical_mode"`
	CanonicalUrl     string       `json:"canonical_url"`
	BreadcrumbTitle  string       `json:"breadcrumb_title"`
	BreadcrumbParent string       `json:"breadcrumb_parent"`
	BreadcrumbUrl 	 string       `json:"breadcrumb_url"`
	SchemaType 	 	 string       `json:"schema_type"`
	// Date
	PublishedAt    	 time.Time    `json:"published_at,omitempty"`
	UpdatedAt        *time.Time   `json:"updated_at,omitempty"`
	// Content
	Content			 []ContentDto `json:"content"`
}

type ContentDto struct{
	ContentTitle     string      `json:"content_title"`
	TitleTag         string      `json:"title_tag"`
	Paragraph        []string    `json:"paragraph"`
	ParagraphTag     []string    `json:"paragraph_tag"`
	NewLine          []bool      `json:"new_line"`
	ContentImgUrl    *string     `json:"content_img_url,omitempty"`
	Order			 uint8       `json:"order"`
}

type NewProductDto struct {
	Name        string
	Description string
	ImageUrl    string
	Mark        string
	Price       uint32
	Discount    uint8
	SalesCount  uint64
	Count       uint16
	Images      []string
	Features    []string
	Tags        []string
}

type ProductDto struct {
	ID			uint64     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ImageUrl    string     `json:"image_url"`
	Mark        string     `json:"mark"`
	Price       uint32     `json:"price"`
	FinalPrice  uint32     `json:"final_price"`
	Discount    uint8      `json:"discount"`
	SalesCount  uint64     `json:"sales_count"`
	Count       uint16     `json:"count"`
	Images      []string   `json:"images"`
	Features    []string   `json:"features"`
	Tags        []string   `json:"tags"`
}

type CartDto struct {
	ID         uint64  `json:"id"`
	ProductID  uint64  `json:"product_id"`
	Quantity   uint16  `json:"quantity"`
    Price      uint32  `json:"price"`
	Discount   uint8   `json:"discount"`
	FinalPrice uint32  `json:"final_price"`
	Name       string  `json:"name"`
	ImgUrl     string  `json:"img_url"`
	Count      uint16  `json:"count"`
}