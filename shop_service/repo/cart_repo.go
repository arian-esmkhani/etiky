package repo

import (
	"context"
	"fiber-app/dto"
)

// CartRepo defines the interface for cart repository operations including add, update, delete, and retrieval of cart items
type ICartRepo interface {
	Add(ctx context.Context, userId, productID uint64) error
	Update(ctx context.Context, userId, productID uint64, statutes bool) error
	HardDelete(ctx context.Context, userId, productID uint64) error
	Get(ctx context.Context, userId uint64) ([]dto.CartDto, error)
	DeleteExpiredCards(ctx context.Context) (int64, error)
	GetProductQuantity(ctx context.Context, userId, productID uint64) (uint16, error)
}