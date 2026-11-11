package repo

import (
	"context"
	"database/sql"
	"errors"
	"fiber-app/dto"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
)

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) IProductRepo {
	return &productRepo{db: db}
}

var errProductNotFound error = errors.New("Product not found or already deleted")

func (r *productRepo) Create(ctx context.Context, product dto.NewProductDto) error {
	query := `
		INSERT INTO products (
			name, 
			description, 
			image_url, 
			mark, 
			price, 
			discount, 
			sales_count, 
			count, 
			images, 
			features, 
			tags,
			created_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
		) RETURNING id, created_at
	`

	var id uint64
	var createdAt time.Time

	err := r.db.QueryRow(
		ctx,
		query,
		product.Name,
		product.Description,
		product.ImageUrl,
		product.Mark,
		product.Price,
		product.Discount,
		product.SalesCount,
		product.Count,
		product.Images,
		product.Features,
		product.Tags,
		time.Now(),
	).Scan(&id, &createdAt)

	if err != nil {
		return fmt.Errorf("Failed to create product: %v", err)
	}

	log.Infof("Product created successfully - ID: %d, Name: %s, CreatedAt: %s", 
		id, 
		product.Name, 
		createdAt.Format("2006-01-02 15:04:05"),
	)
	return nil
}

func (r *productRepo) Update(ctx context.Context, id uint64, product dto.NewProductDto) error {
	query := `
		UPDATE products 
		SET 
			name = COALESCE($1, name),
			description = COALESCE($2, description),
			image_url = COALESCE($3, image_url),
			mark = COALESCE($4, mark),
			price = COALESCE($5, price),
			discount = COALESCE($6, discount),
			sales_count = COALESCE($7, sales_count),
			count = COALESCE($8, count),
			images = COALESCE($9, images),
			features = COALESCE($10, features),
			tags = COALESCE($11, tags),
			updated_at = $12
		WHERE id = $13 AND deleted_at IS NULL
		RETURNING id, updated_at
	`

	var updatedId uint64
	var updatedAt time.Time

	err := r.db.QueryRow(
		ctx,
		query,
		product.Name,
		product.Description,
		product.ImageUrl,
		product.Mark,
		product.Price,
		product.Discount,
		product.SalesCount,
		product.Count,
		product.Images,
		product.Features,
		product.Tags,
		time.Now(),
		id,
	).Scan(&updatedId, &updatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Warnf("Product not found or already deleted - ID: %d", id)
			return errProductNotFound
		}
		return fmt.Errorf("Failed to update product (ID: %d): %v", id, err)
	}

	log.Infof("Product updated successfully - ID: %d, UpdatedAt: %s", 
		updatedId, 
		updatedAt.Format("2006-01-02 15:04:05"),
	)
	return nil
}

func (r *productRepo) SoftDelete(ctx context.Context, id uint64) error {
	query := `
		UPDATE products 
		SET deleted_at = $1
		WHERE id = $2 AND deleted_at IS NULL
		RETURNING id, deleted_at
	`

	var deletedId uint64
	var deletedAt time.Time

	err := r.db.QueryRow(
		ctx,
		query,
		time.Now(),
		id,
	).Scan(&deletedId, &deletedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Warnf("Product not found or already deleted - ID: %d", id)
			return errProductNotFound
		}
		return fmt.Errorf("Failed to soft delete product (ID: %d): %v", id, err)
	}

	log.Infof("Product soft deleted successfully - ID: %d, DeletedAt: %s", 
		deletedId, 
		deletedAt.Format("2006-01-02 15:04:05"),
	)
	return nil
}

func (r *productRepo) HardDelete(ctx context.Context, id uint64) error {
	query := `
		DELETE FROM products 
		WHERE id = $1
		RETURNING id, name
	`

	var deletedId uint64
	var deletedName string

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(&deletedId, &deletedName)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Warnf("Product not found - ID: %d", id)
			return errProductNotFound
		}
		return fmt.Errorf("Failed to hard delete product (ID: %d): %v", id, err)
	}

	log.Infof("Product hard deleted permanently - ID: %d, Name: %s", 
		deletedId, 
		deletedName,
	)
	return nil
}

func (r *productRepo) Restore(ctx context.Context, id uint64) error {
	query := `
		UPDATE products 
		SET deleted_at = NULL
		WHERE id = $1 AND deleted_at IS NOT NULL
		RETURNING id, name
	`

	var restoredId uint64
	var restoredName string

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(&restoredId, &restoredName)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Warnf("Product not found or not deleted - ID: %d", id)
			return errProductNotFound
		}
		return fmt.Errorf("Failed to restore product (ID: %d): %v", id, err)
	}

	log.Infof("Product restored successfully - ID: %d, Name: %s", 
		restoredId, 
		restoredName,
	)
	return nil
}

func (r *productRepo) FindByID(ctx context.Context, id uint64) (*dto.ProductDto, error) {
	query := `
		SELECT 
			p.id, 
			p.name, 
			p.description, 
			p.image_url, 
			p.mark, 
			p.price, 
			p.discount, 
			p.sales_count, 
			p.count, 
			p.images, 
			p.features, 
			p.tags,
			(p.price * (100 - p.discount) / 100)::BIGINT AS final_price
		FROM products p
		WHERE id = $1 AND p.deleted_at IS NULL 
	`

	var product dto.ProductDto

	err := r.db.QueryRow(ctx, query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.ImageUrl,
		&product.Mark,
		&product.Price,
		&product.Discount,
		&product.SalesCount,
		&product.Count,
		&product.Images,
		&product.Features,
		&product.Tags,
		&product.FinalPrice,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errProductNotFound
		}
		return nil, fmt.Errorf("Failed to find product by ID %d: %v", id, err)
	}

	return &product, nil
}

func (r *productRepo) GetProductImages(ctx context.Context, id uint64) (string, []string, error) {
	query := `SELECT image_url, images FROM products WHERE id = $1`
	
	var mainImage string
	var galleryImages []string
	
	err := r.db.QueryRow(ctx, query, id).Scan(&mainImage, &galleryImages)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil, errProductNotFound
		}
		log.Errorf("Failed to find product images %d: %v", id, err)
		return "", nil, err
	}
	
	return mainImage, galleryImages, nil
}

func (r *productRepo) UpdateGallery(ctx context.Context, id uint64, images []string) error {
	query := `
		UPDATE products 
		SET images = $1, updated_at = $2
		WHERE id = $3 AND deleted_at IS NULL
		RETURNING id
	`

	var updatedId uint64
	err := r.db.QueryRow(ctx, query, images, time.Now(), id).Scan(&updatedId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errProductNotFound
		}
		return fmt.Errorf("Failed to update gallery for product %d: %v", id, err)
	}

	log.Infof("Gallery updated successfully for product ID: %d", id)
	return nil
}

func (r *productRepo) UpdatePricesByPercent(ctx context.Context, percent float64) error {
	query := `
		UPDATE products 
		SET 
			price = price * (1 + $1 / 100),
			updated_at = $2
		WHERE deleted_at IS NULL
	`

	_, err := r.db.Exec(
		ctx,
		query,
		percent,
		time.Now(),
	)

	if err != nil {
		return fmt.Errorf("Failed to update products prices: %v", err)
	}

	log.Infof("All products prices updated successfully - Percent: %.2f%%", percent)
	return nil
}