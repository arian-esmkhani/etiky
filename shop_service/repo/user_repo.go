package repo

import "context"

// IUserRepo defines the interface for user data access operations.
type IUserRepo interface {
	CreateUser(ctx context.Context, number, password string) error
	CreateAdminUser(ctx context.Context, number, password string) error
	IsUserExist(ctx context.Context, number string) (bool, error)
	FindUserByNumber(ctx context.Context, number string) (uint64, string, string, error)
}