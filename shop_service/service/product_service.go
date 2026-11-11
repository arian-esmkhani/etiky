package service

import (
	"context"
	"fiber-app/dto"
	"mime/multipart"
)

// IProductService provides product management operations.
type IProductService interface {
	CreateProduct(ctx context.Context, req dto.CreateProductReq) error
	UpdateProduct(ctx context.Context, id uint64, req dto.CreateProductReq) error
	SoftDeleteProduct(ctx context.Context, id uint64) error
	HardDeleteProduct(ctx context.Context, id uint64) error
	RestoreProduct(ctx context.Context, id uint64) error
	GetProductByID(ctx context.Context, id uint64) (*dto.ProductDto, error)
	AddProductImages(ctx context.Context, id uint64, files []*multipart.FileHeader) error
	RemoveProductImages(ctx context.Context, id uint64, imagePaths []string) error
}