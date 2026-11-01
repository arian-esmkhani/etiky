package repo

import (
	"context"
	"fiber-app/dto"
)

// IFinderRepo provides data access for fetching and finding data.
type IFinderRepo interface {
	// product
	FindCheapestProducts(ctx context.Context,
		limit, offset uint8) ([]dto.PreviewProductDto, error)
	CountCheapestProducts(ctx context.Context) (uint64, error)

	FindProductsByDiscount(ctx context.Context, category string,
		limit, offset uint8) ([]dto.PreviewProductDto, error)
	CountProductsByDiscount(ctx context.Context, category string) (uint64, error)

	FindBestSellingProducts(ctx context.Context, category string,
		limit, offset uint8) ([]dto.PreviewProductDto, error)
	CountBestSellingProducts(ctx context.Context, category string) (uint64, error)

	FindCheapestProductsByFilter(ctx context.Context, filter dto.ProductFilterDto,
		limit, offset uint8) ([]dto.PreviewProductDto, error)
	CountCheapestProductsByFilter(ctx context.Context, filter dto.ProductFilterDto) (uint64, error)

	SearchProducts(ctx context.Context, queryStr, category string, lowPrice, maxPrice uint32, limit,
					offset uint8) ([]dto.PreviewProductDto, error)
	CountSearch(ctx context.Context, queryStr, category string, lowPrice,
					maxPrice uint32) (uint64, error)

	GetSearchSuggestions(ctx context.Context, queryStr string) ([]string, error)
	// article
	FindArticleBySlug(ctx context.Context, url string) (*dto.ArticleDto, error)
	FindArticleByTitle(ctx context.Context, title string) (*dto.ArticleDto, error)
}
