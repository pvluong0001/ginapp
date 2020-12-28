package controllers

import (
	"ginapp/auth/services"
	"ginapp/core/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	AuthService services.AuthService
	JwtService  services.JwtService
}

func ProvideAuthController(authService services.AuthService, jwtService services.JwtService) AuthController {
	return AuthController{
		AuthService: authService,
		JwtService:  jwtService,
	}
}

type registerRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (controller AuthController) RegisterAction() gin.HandlerFunc {
	return func(context *gin.Context) {
		var request registerRequest
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
		if err := controller.AuthService.RegisterUserHandle(&user); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "Register success",
			"data":    user,
		})
	}
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (controller AuthController) LoginAction() gin.HandlerFunc {
	return func(context *gin.Context) {
		var request loginRequest
		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		user := models.User{
			Email: request.Email,
		}
		if err := controller.AuthService.GetUserByEmail(&user); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := user.ComparePassword(request.Password); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "The credentials is not correct",
			})
			return
		}

		token, _ := controller.JwtService.GenerateToken(user.ID)

		context.JSON(http.StatusOK, gin.H{
			"message":      "login success",
			"access_token": token,
			"data":         user,
		})
	}
}
