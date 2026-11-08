package helper

import (
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

const accessTTL  = 24 * time.Hour

type Claims struct {
	UserId uint64 `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type Config struct {
	AccessSecret  []byte
	Issuer        string
}

type JWT struct {
	cfg Config
}

func NewJWT(cfg Config) *JWT {
	return &JWT{cfg: cfg}
}

func (j *JWT) GenerateAccessToken(userId uint64,
	role string) (string, error) {
	claims := Claims{
		UserId: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    j.cfg.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(j.cfg.AccessSecret)
	if err != nil {
		log.Errorf("GenerateAccessToken error for user %s: %v", userId, err)
		return "", err
	}
	
	return tokenStr, nil
}

func (j *JWT) ParseAccessToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, jwt.ErrTokenInvalidClaims
		}
		return j.cfg.AccessSecret, nil
	})
	
	if err != nil {
		log.Errorf("ParseAccessToken error: %v", err)
		return nil, err
	}
	
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		log.Errorf("ParseAccessToken: invalid token claims")
		return nil, jwt.ErrTokenInvalidClaims
	}
	
	return claims, nil
}