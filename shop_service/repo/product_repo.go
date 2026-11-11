package repo

import (
	"context"
	"fiber-app/dto"
)

// IProductRepo provides data access methods for product management operations.
type IProductRepo interface {
	Create(ctx context.Context, product dto.NewProductDto) error
	Update(ctx context.Context, id uint64, product dto.NewProductDto) error
	SoftDelete(ctx context.Context, id uint64) error
	HardDelete(ctx context.Context, id uint64) error
	Restore(ctx context.Context, id uint64) error
	FindByID(ctx context.Context, id uint64) (*dto.ProductDto, error)
	GetProductImages(ctx context.Context, id uint64) (string, []string, error)
	UpdateGallery(ctx context.Context, id uint64, images []string) error
	UpdatePricesByPercent(ctx context.Context, percent float64) error
}