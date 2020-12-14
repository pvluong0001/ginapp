package user

import (
	"ginapp/user/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(group *gin.RouterGroup, controller controllers.UserController) {
	group.GET("/users", controller.GetUserListAction())
}
