package service

import (
	"context"
	"fiber-app/dto"
)

// ISliderService provides product lists for slider widgets.
type ISliderService interface {
	GetProductsByDiscount(ctx context.Context) ([]dto.PreviewProductDto, error)
	GetCheapestProducts(ctx context.Context) ([]dto.PreviewProductDto, error)
	GetBestSellingProducts(ctx context.Context) ([]dto.PreviewProductDto, error)
	GetCheapestProductsByFilter(ctx context.Context, filter dto.ProductFilterDto) ([]dto.PreviewProductDto, error)
}