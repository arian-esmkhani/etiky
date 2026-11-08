package service

import (
	"fiber-app/dto"
	"context"
)

type IUserService interface {
	CreateUser(ctx context.Context, loginData dto.LoginDto) error
	CreateAdminUser(ctx context.Context, loginData dto.LoginDto) error
	Login(ctx context.Context, loginData dto.LoginDto) (string, error)
}