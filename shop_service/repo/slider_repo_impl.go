package repo

import (
	"context"
	"fiber-app/dto"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
)

type sliderRepo struct {
	db *pgxpool.Pool
}

func NewSliderRepo(db *pgxpool.Pool) ISliderRepo {
	return &sliderRepo{db: db}
}

// defaultLimit defines the maximum number of products returned per slider query.
// Aligned with UI slider component capacity (11 items).
const defaultLimit uint8 = 11

func (r sliderRepo) GetProductsByDiscount(ctx context.Context) ([]dto.PreviewProductDto, error) {
	query := `
		SELECT 
			p.id,
			p.name,
			p.image_url,
			p.price,
			p.discount,
			(p.price * (100 - p.discount) / 100)::BIGINT AS final_price
		FROM products p
		WHERE p.deleted_at IS NULL AND p.count > 0 AND p.discount > 0
		ORDER BY p.discount DESC
		LIMIT $1
	`

	rows, err := r.db.Query(ctx, query, defaultLimit)
	if err != nil {
		return nil, fmt.Errorf("failed to query discounted products: %v", err)
	}
	defer rows.Close()

	products := make([]dto.PreviewProductDto, 0, defaultLimit)
	for rows.Next() {
		var p dto.PreviewProductDto
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.ImgUrl,
			&p.Price,
			&p.Discount,
			&p.FinalPrice,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan discounted product row: %v", err)
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error for discounted products: %v", err)
	}

	if len(products) < 1 {
		log.Debug("Discounted product len is under zero !")
	}

	return products, nil
}

func (r sliderRepo) GetCheapestProducts(ctx context.Context) ([]dto.PreviewProductDto, error) {

	query := `
		SELECT 
			p.id,
			p.name,
			p.image_url,
			p.price,
			p.discount,
			(p.price * (100 - p.discount) / 100)::BIGINT AS final_price
		FROM products p
		WHERE p.deleted_at IS NULL AND p.count > 0
		ORDER BY final_price ASC
		LIMIT $1
	`

	rows, err := r.db.Query(ctx, query, defaultLimit)
	if err != nil {
		return nil, fmt.Errorf("failed to query cheapest products: %v", err)
	}
	defer rows.Close()

	products := make([]dto.PreviewProductDto, 0, defaultLimit)
	for rows.Next() {
		var p dto.PreviewProductDto
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.ImgUrl,
			&p.Price,
			&p.Discount,
			&p.FinalPrice,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan cheapest product row: %v", err)
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error for cheapest products: %v", err)
	}

	if len(products) < 1 {
		log.Debug("Cheapest product len is under zero !")
	}

	return products, nil
}

func (r sliderRepo) GetBestSellingProducts(ctx context.Context) ([]dto.PreviewProductDto, error) {

	query := `
		SELECT 
			p.id,
			p.name,
			p.image_url,
			p.price,
			p.discount,
			(p.price * (100 - p.discount) / 100)::BIGINT AS final_price
		FROM products p
		WHERE p.deleted_at IS NULL AND p.count > 0
		ORDER BY p.sales_count DESC
		LIMIT $1
	`

	rows, err := r.db.Query(ctx, query, defaultLimit)
	if err != nil {
		return nil, fmt.Errorf("failed to query best selling products: %v", err)
	}
	defer rows.Close()

	products := make([]dto.PreviewProductDto, 0, defaultLimit)
	for rows.Next() {
		var p dto.PreviewProductDto
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.ImgUrl,
			&p.Price,
			&p.Discount,
			&p.FinalPrice,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan best selling product row: %v", err)
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error for best selling products: %v", err)
	}

	if len(products) < 1 {
		log.Debug("Best selling product len is under zero !")
	}

	return products, nil
}

func (r sliderRepo) GetCheapestProductsByFilter(ctx context.Context,
	filter dto.ProductFilterDto) ([]dto.PreviewProductDto, error) {

	query := `
		SELECT 
			p.id,
			p.name,
			p.image_url,
			p.price,
			p.discount,
			(p.price * (100 - p.discount) / 100)::BIGINT AS final_price
		FROM products p
		WHERE p.deleted_at IS NULL
			AND (p.tags @> ARRAY[$1] OR p.kind = $1) AND p.count > 0
		ORDER BY final_price ASC
		LIMIT $2
	`

	rows, err := r.db.Query(ctx, query, filter, defaultLimit)
	if err != nil {
		return nil, fmt.Errorf("failed to query cheapest products by filter '%s': %v", filter, err)
	}
	defer rows.Close()

	products := make([]dto.PreviewProductDto, 0, defaultLimit)
	for rows.Next() {
		var p dto.PreviewProductDto
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.ImgUrl,
			&p.Price,
			&p.Discount,
			&p.FinalPrice,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan cheapest product by filter '%s' row: %v", filter, err)
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error for cheapest products by filter '%s': %v", filter, err)
	}

	if len(products) < 1 {
		log.Debugf("Cheapest product by filter len is under zero !")
	}

	return products, nil
}
