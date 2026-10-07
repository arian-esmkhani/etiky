package models

import (
	"time"
)

type Product struct {
	ID            uint64       `db:"id"`
	Name      	  string       `db:"name"`
	Description   string       `db:"description"`
	ImgUrl        string       `db:"image_url"`
	Mark          string       `db:"mark"`
	Price         uint32       `db:"price"`
	Discount      uint8        `db:"discount"`
	SalesCount    uint64       `db:"sales_count"`
	Count	      uint16       `db:"count"`
	Images        []string     `db:"images"`
	Features      []string     `db:"features"`
	Tags   	      []string     `db:"tags"`
	CreatedAt     time.Time    `db:"created_at"`
	UpdatedAt     *time.Time   `db:"updated_at"`
	DeletedAt     *time.Time   `db:"deleted_at"`
}