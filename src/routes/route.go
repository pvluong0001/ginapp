package auth

import (
	"ginapp/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup) {
	auth := route.Group("/auth")
	{
		auth.POST("/register", authController.RegisterHandler)
	}
}
