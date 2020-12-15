package services

import (
	"fmt"
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

func (service *JwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
}
