package models

import "time"

type Cart struct {
    ID         uint64 	  `db:"id"`
    UserID     uint64     `db:"user_id"`
    ProductID  uint64     `db:"product_id"`
    Quantity   uint16     `db:"quantity"`
    Price      uint32     `db:"price"`
    CreatedAt  *time.Time `db:"created_at"`
}