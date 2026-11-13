package service

import (
	"context"
	"fiber-app/dto"

	"fiber-app/repo"

	"github.com/gofiber/fiber/v2/log"
)

type cartService struct {
	repo repo.ICartRepo
}


func NewCartService(repo repo.ICartRepo) ICartService {
	return &cartService{repo: repo}
}

func (s *cartService) Add(ctx context.Context, userId uint64, productID uint64) error {
    err := s.repo.Add(ctx, userId, productID)
    if err != nil {
        log.Errorf("Failed to add item to cart: user_id=%d, product_id=%d, error=%v", userId, productID, err)
        return err
    }

    return nil
}

func (s *cartService) Get(ctx context.Context, userId uint64) (*dto.CartRes, error) {
    cart, err := s.repo.Get(ctx, userId)
    if err != nil {
        log.Errorf("Failed to get cart items: user_id=%d, error=%v", userId, err)
        return nil, err
    }

    var finalPrice uint32
    for _, item := range cart {
        finalPrice += item.FinalPrice
    }

    return &dto.CartRes{
        CartDto:    cart,
        FinalPrice: finalPrice,
    }, nil
}

func (s *cartService) HardDelete(ctx context.Context, userId uint64, productID uint64) error {
    err := s.repo.HardDelete(ctx, userId, productID)
    if err != nil {
        log.Errorf("Failed to hard delete item from cart: user_id=%d, product_id=%d, error=%v", userId, productID, err)
        return err
    }

    return nil
}

func (s *cartService) Update(ctx context.Context, userId uint64, productID uint64, statutes bool) error {
    err := s.repo.Update(ctx, userId, productID, statutes)
    if err != nil {
        log.Errorf("Failed to update cart item: user_id=%d, product_id=%d, error=%v", userId, productID, err)
        return err
    }

    return nil
}

func (s *cartService) GetProductQuantity(ctx context.Context, userId uint64, productID uint64) (uint16, error) {
    quantity, err := s.repo.GetProductQuantity(ctx, userId, productID)
    if err != nil {
        log.Errorf("Failed to get product quantity from cart: user_id=%d, product_id=%d, error=%v", userId, productID, err)
        return 0, err
    }
    return quantity, nil
}