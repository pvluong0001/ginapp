package controllers

import (
	"ginapp/auth/services"
	"ginapp/core/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	AuthService services.AuthService
}

func ProvideAuthController(service services.AuthService) AuthController {
	return AuthController{
		AuthService: service,
	}
}

type userRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (controller AuthController) RegisterAction() gin.HandlerFunc {
	return func(context *gin.Context) {
		var request userRequest
		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		user := models.User{
			Email:    request.Email,
			Password: request.Password,
		}
		controller.AuthService.RegisterUserHandle(&user)

		context.JSON(http.StatusOK, gin.H{
			"message": "Register success",
			"data":    user,
		})
	}
}
