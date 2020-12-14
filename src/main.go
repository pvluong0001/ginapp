package main

import (
	"ginapp/auth"
	models "ginapp/core"
	"ginapp/user"
	"github.com/gin-gonic/gin"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// connect database
	db := models.ConnectDatabase()
	userController := initUserController(db)
	authController := initAuthController(db)

	server := gin.Default()

	api := server.Group("/api")
	{
		user.Routes(api, userController)
		auth.Routes(api, authController)
	}

	server.Run(":" + os.Getenv("PORT"))
}
