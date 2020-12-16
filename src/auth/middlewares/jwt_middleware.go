package middlewares

import (
	"fmt"
	"ginapp/auth/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CheckJwt(service services.JwtService) gin.HandlerFunc {
	return func(context *gin.Context) {
		const BearerSchema = "Bearer"
		authHeader := context.GetHeader("Authorization")
		if len(authHeader) <= len(BearerSchema) {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		tokenString := strings.Trim(authHeader[len(BearerSchema):], " ")
		token, err := service.ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
		}
	}
}
