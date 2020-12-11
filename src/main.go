package main

import (
	"ginapp/models"
	"ginapp/routes"
	"github.com/gin-gonic/gin"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// connect database
	models.ConnectDatabase()

	server := gin.Default()

	api := server.Group("/api")
	{
		auth.Routes(api)
	}

	server.Run(":" + os.Getenv("PORT"))
}
