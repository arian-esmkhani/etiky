package service

import (
	"context"
	"fiber-app/dto"
)

// IFinderService provides data discovery
type IFinderService interface {
	// product
	FindCheapestProducts(ctx context.Context,
		page uint8) (dto.PreviewProductRes, error)
	FindProductsByDiscount(ctx context.Context, category string,
		page uint8) (dto.PreviewProductRes, error)
	FindBestSellingProducts(ctx context.Context, category string,
		page uint8) (dto.PreviewProductRes, error)
	FindCheapestProductsByFilter(ctx context.Context, filter dto.ProductFilterDto,
		page uint8) (dto.PreviewProductRes, error)
	SearchProducts(ctx context.Context, queryStr, category string, lowPrice, maxPrice uint32, 
		page uint8) (dto.PreviewProductRes, error)
	GetSearchSuggestions(ctx context.Context, queryStr string) ([]string, error)
	// article
	FindArticleBySlug(ctx context.Context, url string) (*dto.ArticleDto, error)
	FindArticleByTitle(ctx context.Context, title string) (*dto.ArticleDto, error)
}