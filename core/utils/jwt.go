package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/nade-harlow/E-commerce/core/models"
	"os"
	"time"
)

func GenerateToken(user *models.User) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": int64(user.ID),
		"exp":     time.Now().Add(time.Hour * 2400000).Unix(),
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}
