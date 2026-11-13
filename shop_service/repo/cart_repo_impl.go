package repo

import (
	"context"
	"errors"
	"fiber-app/dto"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type cartRepo struct {
	db *pgxpool.Pool
}

func NewCartRepo(db *pgxpool.Pool) ICartRepo {
	return &cartRepo{db: db}
}

const defaultQuantity uint8 = 1

func (r *cartRepo) Add(ctx context.Context, userId, productID uint64) error {
	var price uint32
    err := r.db.QueryRow(ctx,
        "SELECT p.price FROM products p WHERE id = $1",
        productID,
    ).Scan(&price)
    if err != nil {
        return fmt.Errorf("failed to get price when create cart '%s': %v", productID, err)
    }

	query := `
		INSERT INTO carts (user_id, product_id, quantity, price, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err = r.db.Exec(ctx, query, userId, productID, defaultQuantity, price, time.Now())
	if err != nil {
		return fmt.Errorf("failed to create cart '%s': %v", productID, err)
	}

	return nil
}

func (r *cartRepo) Update(ctx context.Context, userId, productID uint64, statutes bool) error {
    var count uint16
    var quantity uint16
    
    
    err := r.db.QueryRow(ctx,
        `SELECT p.count , c.quantity
        FROM carts c 
        JOIN products p ON c.product_id = p.id
        WHERE product_id = $1 AND user_id = $2`,
        productID, userId,
    ).Scan(&count, &quantity)
    if err != nil {
        return fmt.Errorf("failed to query cart data: user_id=%d, product_id=%d, error=%v", userId, productID, err)
    }
    

    query := `
        UPDATE carts 
        SET 
            quantity = quantity + $1,
            created_at = $2
        WHERE product_id = $3 AND user_id = $4
    `

    var add int8
    if statutes == true {
        add = 1
    } else {
        add = -1
    }

    if !((quantity + uint16(add)) <= count) {
        log.Warnf("Update blocked: insufficient stock. user_id=%d, product_id=%d, quantity=%d, add=%d, count=%d", 
            userId, productID, quantity, add, count)
        return nil
    }

    _, err = r.db.Exec(
        ctx,
        query,
        add,
        time.Now(),
        productID,
        userId,
    )

    if err != nil {
        return fmt.Errorf("failed to update cart: user_id=%d, product_id=%d, error=%v", userId, productID, err)
    }
    
    return nil
}

func (r *cartRepo) HardDelete(ctx context.Context, userId, productID uint64) error {    
    query := `
        DELETE FROM carts 
        WHERE product_id = $1 AND user_id = $2
    `

    _, err := r.db.Exec(
        ctx,
        query,
        productID,
        userId,
    )

    if err != nil {
        return fmt.Errorf("failed to hard delete cart item: user_id=%d, product_id=%d, error=%v", userId, productID, err)
    }
    
    return nil
}

func (r cartRepo) Get(ctx context.Context, userId uint64) ([]dto.CartDto, error) {
	query := `
		SELECT 
			p.id AS product_id,
			p.name,
			p.image_url,
			c.id,
			c.quantity,
			(c.price  * c.quantity) AS price,
			p.discount,
			(((p.price * (100 - p.discount) / 100)::BIGINT) * c.quantity) AS final_price,
			p.count
		FROM carts c
		JOIN products p ON c.product_id = p.id
		WHERE c.user_id = $1 
	`

	rows, err := r.db.Query(ctx, query, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to query cart products: %v", err)
	}
	defer rows.Close()

	var cart []dto.CartDto
	for rows.Next() {
		var c dto.CartDto
		err := rows.Scan(
			&c.ProductID,
			&c.Name,
			&c.ImgUrl,
			&c.ID,
			&c.Quantity,
			&c.Price,
			&c.Discount,
			&c.FinalPrice,
			&c.Count,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan discounted product row: %v", err)
		}
		cart = append(cart, c)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error for discounted products: %v", err)
	}

	return cart, nil
}

func (r *cartRepo) DeleteExpiredCards(ctx context.Context) (int64, error) {
	query := `
		DELETE FROM carts 
		WHERE created_at < NOW() - INTERVAL '12 hours'
		AND deleted_at IS NULL
		RETURNING COUNT(*)
	`

	var deletedCount int64
	err := r.db.QueryRow(ctx, query).Scan(&deletedCount)
	if err != nil {
		return 0, fmt.Errorf("Failed to delete expired cards: %v", err)
	}

	return deletedCount, nil
}

func (r *cartRepo) GetProductQuantity(ctx context.Context, userId, productID uint64) (uint16, error) {
    var quantity uint16
    query := `SELECT COALESCE(quantity, 0) FROM carts WHERE user_id = $1 AND product_id = $2`
    err := r.db.QueryRow(ctx, query, userId, productID).Scan(&quantity)
    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            return 0, nil
        }
        return 0, fmt.Errorf("Failed to get product quantity from cart: user_id=%d, product_id=%d, error=%v", userId, productID, err)
    }
    return quantity, nil
}