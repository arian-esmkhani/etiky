package service

import (
	"context"
	"fiber-app/dto"
)

// CartService defines the interface for cart business logic operations including add, update, delete, and retrieval of cart items
type ICartService interface {
	Add(ctx context.Context, userId, productID uint64) error
	Update(ctx context.Context, userId, productID uint64, statutes bool) error
	HardDelete(ctx context.Context, userId, productID uint64) error
	Get(ctx context.Context, userId uint64) (*dto.CartRes, error)
	GetProductQuantity(ctx context.Context, userId uint64, productID uint64) (uint16, error)
}