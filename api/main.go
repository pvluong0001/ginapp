package main

import (
	"ginapp/auth"
	"ginapp/auth/middlewares"
	"ginapp/auth/services"
	models "ginapp/core"
	"ginapp/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"net/http"
	"os"
)

func main() {
	// connect database
	db := models.ConnectDatabase()

	userController := initUserController(db)
	authController := initAuthController(db)
	service := services.ProvideJwtService()

	server := gin.Default()

	// apply cors
	server.Use(cors.Default())

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
