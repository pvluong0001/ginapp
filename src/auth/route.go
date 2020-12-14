package auth

import (
	"ginapp/auth/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(group *gin.RouterGroup, controller controllers.AuthController) {
	group.POST("/auth/register", controller.RegisterAction())
	group.POST("/auth/login", controller.LoginAction())
}
