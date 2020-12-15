package main

import (
	"ginapp/auth"
	"ginapp/auth/middlewares"
	"ginapp/auth/services"
	models "ginapp/core"
	"ginapp/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// connect database
	db := models.ConnectDatabase()
	userController := initUserController(db)
	authController := initAuthController(db)
	service := services.ProvideJwtService()

	server := gin.Default()

	api := server.Group("/api")
	{
		user.Routes(api, userController)
		auth.Routes(api, authController)

		api.GET("/test", middlewares.CheckJwt(service), func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "good",
			})
		})
	}

	server.Run(":" + os.Getenv("PORT"))
}
