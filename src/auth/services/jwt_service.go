package services

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JwtService struct{}

func ProvideJwtService() JwtService {
	return JwtService{}
}

func (service *JwtService) GenerateToken(userId uint) (string, error) {
	secret := []byte(os.Getenv("SECRET_KEY"))

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		Issuer:    "testt",
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString(secret)

	if err != nil {
		panic(err)
	}

	return token, nil
}
