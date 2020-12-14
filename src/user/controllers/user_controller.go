package controllers

import (
	"ginapp/core/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func ProvideUserController(service services.UserService) UserController {
	return UserController{
		UserService: service,
	}
}

func (controller *UserController) GetUserListAction() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": controller.UserService.GetAllUserHandle(),
		})
	}
}
