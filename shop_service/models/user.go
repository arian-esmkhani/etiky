package models

import (
	"time"
)

type User struct {
	Id        uint64    		 `db:"_id,omitempty"`
	Number    string             `db:"number"`
	Password  string             `db:"password"`
	Role      string             `db:"role"`
	Status    bool               `db:"status"`
	CreatedAt time.Time    		 `db:"created_at"`
	UpdatedAt *time.Time   		 `db:"updated_at"`
	DeletedAt *time.Time   		 `db:"deleted_at"`
}
