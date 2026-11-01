package models

import (
	"time"
)

type Article struct {
	ID               uint64     `db:"id"`
	Title            string     `db:"title"`
	Description      string     `db:"description"`
	ViewsCount       uint64     `db:"views_count"`
	AuthorID         uint64     `db:"author_id"`
	Status           string     `db:"status"`
	ImgUrl           string     `db:"img_url"`
	VideoUrl         string     `db:"video_url"`
	Tags         	 []string   `db:"tags"`
	// SEO
	Slug             string     `db:"slug"`
    MetaTitle        string     `db:"meta_title"`
    MetaDescription  string     `db:"meta_description"`
    MetaRobots       string     `db:"meta_robots"`
	CanonicalMode    string     `db:"canonical_mode"`
	CanonicalUrl     string     `db:"canonical_url"`
	BreadcrumbTitle  string     `db:"breadcrumb_title"`
	BreadcrumbParent string     `db:"breadcrumb_parent"`
	BreadcrumbUrl 	 string     `db:"breadcrumb_url"`
	SchemaType 	 	 string     `db:"schema_type"`
	// Open Graph
	OGTitle       	 string 	`db:"og_title"`
	OGDescription 	 string 	`db:"og_description"`
	OGImage       	 string 	`db:"og_image"`
	// Date
	CreatedAt        time.Time  `db:"created_at"`
	PublishedAt    	 time.Time  `db:"published_at"`
	UpdatedAt        *time.Time `db:"updated_at"`
	DeletedAt        *time.Time `db:"deleted_at"`
}