package service

import (
	"context"
	"fiber-app/dto"
	"fiber-app/repo"

	"github.com/gofiber/fiber/v2/log"
)

type sliderService struct {
	repo repo.ISliderRepo
}

func NewSliderService(repo repo.ISliderRepo) ISliderService {
	return &sliderService{repo: repo}
}

func (s *sliderService) GetProductsByDiscount(ctx context.Context) ([]dto.PreviewProductDto, error) {
	slider, err := s.repo.GetProductsByDiscount(ctx)
	if err != nil {
		log.Errorf("Failed to get discounted products: %v", err)
		return nil, err
	}
	return slider, nil
}

func (s *sliderService) GetCheapestProducts(ctx context.Context) ([]dto.PreviewProductDto, error) {
	slider, err := s.repo.GetCheapestProducts(ctx)
	if err != nil {
		log.Errorf("Failed to get cheapest products: %v", err)
		return nil, err
	}
	return slider, nil
}

func (s *sliderService) GetBestSellingProducts(ctx context.Context) ([]dto.PreviewProductDto, error) {
	slider, err := s.repo.GetBestSellingProducts(ctx)
	if err != nil {
		log.Errorf("Failed to get best selling products: %v", err)
		return nil, err
	}
	return slider, nil
}

func (s *sliderService) GetCheapestProductsByFilter(ctx context.Context,
	filter dto.ProductFilterDto) ([]dto.PreviewProductDto, error) {
	slider, err := s.repo.GetCheapestProductsByFilter(ctx, filter)
	if err != nil {
		log.Errorf("Failed to get products by filter: %v", err)
		return nil, err
	}
	return slider, nil
}
