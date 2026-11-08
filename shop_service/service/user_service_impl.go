package service

import (
	"fiber-app/dto"
	"fiber-app/helper"
	"fiber-app/repo"
	"context"
	"errors"

	"github.com/gofiber/fiber/v2/log"
)

type UserService struct {
	user       repo.IUserRepo
	jwt *helper.JWT
}

func NewUserService(user repo.IUserRepo, jwt *helper.JWT) IUserService {
	return &UserService{
		user: user,
		jwt: jwt,
	}
}

func (s *UserService) CreateUser(ctx context.Context, loginData dto.LoginDto) error {
	hashedPass, err := helper.HashPassword(loginData.Password)
	if err != nil {
		log.Errorf("Failed to hash password for '%s': %v", loginData.Number, err)
		return err
	}

	exist, err := s.user.IsUserExist(ctx, loginData.Number)
	if err != nil {
		log.Errorf("Failed to check user existence for '%s': %v", loginData.Number, err)
		return err
	}

	if exist {
		log.Warnf("User creation failed: number '%s' already exists", loginData.Number)
		return errors.New("user already exists")
	}

	err = s.user.CreateUser(ctx, loginData.Number, hashedPass)
	if err != nil {
		log.Errorf("Failed to create user '%s': %v", loginData.Number, err)
		return err
	}

	return nil
}

func (s *UserService) CreateAdminUser(ctx context.Context, loginData dto.LoginDto) error {
	hashedPass, err := helper.HashPassword(loginData.Password)
	if err != nil {
		log.Errorf("Failed to hash password for admin '%s': %v", loginData.Number, err)
		return err
	}

	exist, err := s.user.IsUserExist(ctx, loginData.Number)
	if err != nil {
		log.Errorf("Failed to check admin user existence for '%s': %v", loginData.Number, err)
		return err
	}

	if exist {
		log.Warnf("Admin user creation failed: number '%s' already exists", loginData.Number)
		return errors.New("admin user already exists")
	}

	err = s.user.CreateAdminUser(ctx, loginData.Number, hashedPass)
	if err != nil {
		log.Errorf("Failed to create admin user '%s': %v", loginData.Number, err)
		return err
	}

	return nil
}

func (s *UserService) Login(ctx context.Context, loginData dto.LoginDto) (string, error) {
	id, hashedPass, role, err := s.user.FindUserByNumber(ctx, loginData.Number)
	if err != nil {
		if err.Error() == "user not found" {
			log.Errorf("Failed to find user '%s': %v", loginData.Number, err)
			return "", errors.New("user not found")
		}
		log.Errorf("Failed to find user by name '%s': %v", loginData.Number, err)
		return "", err
	}

	err = helper.ComparePassword(hashedPass, loginData.Password)
	if err != nil {
		log.Warnf("Login failed: invalid password for user '%s'", loginData.Number)
		return  "", errors.New("invalid password")
	}

	accessToken, err := s.jwt.GenerateAccessToken(id, role)
	if err != nil {
		log.Errorf("Failed to generate access token for user '%s': %v", loginData.Number, err)
		return "", errors.New("failed to generate access token")
	}

	return accessToken, nil
}