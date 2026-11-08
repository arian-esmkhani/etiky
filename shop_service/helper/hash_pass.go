package helper

import (
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost int = 10

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		log.Errorf("Error hashing password: %v", err)
		return "", err
	}
	return string(bytes), nil
}

func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}