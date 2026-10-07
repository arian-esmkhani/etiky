package repo

import (
	"context"
	"fiber-app/dto"
)

// ISliderRepo provides data access for fetching product lists.
type ISliderRepo interface {
	GetProductsByDiscount(ctx context.Context) ([]dto.PreviewProductDto, error)
	GetCheapestProducts(ctx context.Context) ([]dto.PreviewProductDto, error)
	GetBestSellingProducts(ctx context.Context) ([]dto.PreviewProductDto, error)
	GetCheapestProductsByFilter(ctx context.Context, filter dto.ProductFilterDto) ([]dto.PreviewProductDto, error)
}
