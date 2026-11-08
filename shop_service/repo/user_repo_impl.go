package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) IUserRepo {
	return &UserRepo{db: db}
}

const (
	roleUser  = "customer"
	roleAdmin = "admin"
)

func (r *UserRepo) CreateUser(ctx context.Context, number, password string) error {
	query := `
		INSERT INTO users (number, password, role, status, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(ctx, query, number, password, roleUser, true, time.Now())
	if err != nil {
		return fmt.Errorf("Failed to create user '%s': %v", number, err)
	}

	log.Infof("User '%s' created successfully with role '%s'", number, roleUser)
	return nil
}

func (r *UserRepo) CreateAdminUser(ctx context.Context, number, password string) error {
	query := `
		INSERT INTO users (number, password, role, status, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(ctx, query, number, password, roleAdmin, true, time.Now())
	if err != nil {
		return fmt.Errorf("Failed to create admin user '%s': %v", number, err)
	}

	log.Infof("Admin user '%s' created successfully with role '%s'", number, roleAdmin)
	return nil
}

func (r *UserRepo) IsUserExist(ctx context.Context, number string) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1 FROM users 
			WHERE number = $1 AND deleted_at IS NULL
		)
	`

	var exists bool
	err := r.db.QueryRow(ctx, query, number).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("Failed to check user existence '%s': %v", number, err)
	}

	return exists, nil
}

func (r *UserRepo) FindUserByNumber(ctx context.Context, number string) (id uint64, password, role string, err error) {
	query := `
		SELECT u.id, u.password, u.role 
		FROM users u
		WHERE number = $1 AND deleted_at IS NULL
	`

	var userID uint64
	err = r.db.QueryRow(ctx, query, number).Scan(&userID, &password, &role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, "", "", errors.New("user not found")
		}
		return 0, "", "", err
	}

	return userID, password, role, nil
}