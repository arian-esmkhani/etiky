package service

import (
	"context"
	"fiber-app/dto"
	"fiber-app/helper"
	"fiber-app/repo"
	"strings"

	"github.com/gofiber/fiber/v2/log"
)

type finderService struct {
	repo repo.IFinderRepo
}

func NewFinderService(repo repo.IFinderRepo) IFinderService {
	return &finderService{repo: repo}
}

const limitPage uint8 = 10

func (s *finderService) FindCheapestProducts(ctx context.Context,
	page uint8) (dto.PreviewProductRes, error) {

	total, err := s.repo.CountCheapestProducts(ctx)
	if err != nil {
		log.Errorf("Failed to count cheapest products: %v", err)
		return dto.PreviewProductRes{}, err
	}

	var hasNext bool = helper.HasNext(page, limitPage, uint8(total))

	products, err := s.repo.FindCheapestProducts(ctx, limitPage, (page * limitPage))
	if err != nil {
		log.Errorf("Failed to find cheapest products (page: %d): %v", page, err)
		return dto.PreviewProductRes{}, err
	}

	return dto.PreviewProductRes{Products: products, HasNext: hasNext}, nil
}

func (s *finderService) FindProductsByDiscount(ctx context.Context, category string,
	page uint8) (dto.PreviewProductRes, error) {

	total, err := s.repo.CountProductsByDiscount(ctx, category)
	if err != nil {
		log.Errorf("Failed to count discounted products (category: %s): %v", category, err)
		return dto.PreviewProductRes{}, err
	}

	var hasNext bool = helper.HasNext(page, limitPage, uint8(total))

	products, err := s.repo.FindProductsByDiscount(ctx, category, limitPage, (page * limitPage))
	if err != nil {
		log.Errorf("Failed to find discounted products (category: %s, page: %d): %v", category, page, err)
		return dto.PreviewProductRes{}, err
	}

	return dto.PreviewProductRes{Products: products, HasNext: hasNext}, nil
}

func (s *finderService) FindBestSellingProducts(ctx context.Context, category string,
	page uint8) (dto.PreviewProductRes, error) {

	total, err := s.repo.CountBestSellingProducts(ctx, category)
	if err != nil {
		log.Errorf("Failed to count best selling products (category: %s): %v", category, err)
		return dto.PreviewProductRes{}, err
	}

	var hasNext bool = helper.HasNext(page, limitPage, uint8(total))

	products, err := s.repo.FindBestSellingProducts(ctx, category, limitPage, (page * limitPage))
	if err != nil {
		log.Errorf("Failed to find best selling products (category: %s, page: %d): %v", category, page, err)
		return dto.PreviewProductRes{}, err
	}

	return dto.PreviewProductRes{Products: products, HasNext: hasNext}, nil
}

func (s *finderService) FindCheapestProductsByFilter(ctx context.Context, filter dto.ProductFilterDto,
	page uint8) (dto.PreviewProductRes, error) {

	total, err := s.repo.CountCheapestProductsByFilter(ctx, filter)
	if err != nil {
		log.Errorf("Failed to count cheapest products by filter (filter: %+v): %v", filter, err)
		return dto.PreviewProductRes{}, err
	}

	var hasNext bool = helper.HasNext(page, limitPage, uint8(total))

	products, err := s.repo.FindCheapestProductsByFilter(ctx, filter, limitPage, (page * limitPage))
	if err != nil {
		log.Errorf("Failed to find cheapest products by filter (filter: %+v, page: %d): %v", filter, page, err)
		return dto.PreviewProductRes{}, err
	}

	return dto.PreviewProductRes{Products: products, HasNext: hasNext}, nil
}

func (s *finderService) SearchProducts(ctx context.Context, queryStr, category string,
	lowPrice, maxPrice uint32, page uint8) (dto.PreviewProductRes, error) {

	queryStr = strings.TrimSpace(queryStr)

	total, err := s.repo.CountSearch(ctx, queryStr, category, lowPrice, maxPrice)
	if err != nil {
		log.Errorf("Failed to count search results (query: %s, category: %s, price range: %d-%d): %v",
			queryStr, category, lowPrice, maxPrice, err)
		return dto.PreviewProductRes{}, err
	}

	var hasNext bool = helper.HasNext(page, limitPage, uint8(total))

	products, err := s.repo.SearchProducts(ctx, queryStr, category, lowPrice, maxPrice,
		limitPage, (page * limitPage))
	if err != nil {
		log.Errorf("Failed to search products (query: %s, category: %s, price range: %d-%d, page: %d): %v",
			queryStr, category, lowPrice, maxPrice, page, err)
		return dto.PreviewProductRes{}, err
	}

	return dto.PreviewProductRes{Products: products, HasNext: hasNext}, nil
}

func (s *finderService) GetSearchSuggestions(ctx context.Context, queryStr string) ([]string, error) {
	queryStr = strings.TrimSpace(queryStr)

	if queryStr == "" {
		return []string{}, nil
	}

	suggestions, err := s.repo.GetSearchSuggestions(ctx, queryStr)
	if err != nil {
		log.Errorf("Failed to get search suggestions for query '%s': %v", queryStr, err)
		return nil, err
	}

	return suggestions, nil
}

func (s *finderService) FindArticleBySlug(ctx context.Context, url string) (*dto.ArticleDto, error) {
	article, err := s.repo.FindArticleBySlug(ctx, url)
	if err != nil {
		log.Errorf("Failed to find article by slug (url: %s): %v", url, err)
		return nil, err
	}

	return article, nil
}

func (s *finderService) FindArticleByTitle(ctx context.Context, title string) (*dto.ArticleDto, error) {
	title = strings.TrimSpace(title)

	article, err := s.repo.FindArticleByTitle(ctx, title)
	if err != nil {
		log.Errorf("Failed to find article by title (title: %s): %v", title, err)
		return nil, err
	}

	return article, nil
}