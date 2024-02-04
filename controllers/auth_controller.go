package controllers

import (
	"gin-fleamarket/dto"
	"gin-fleamarket/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	Signup(ctx *gin.Context)
}

type AuthController struct {
	service services.IAuthService
}

func NewAuthController(service services.IAuthService) IAuthController {
	return &AuthController{
		service: service,
	}
}

func (c *AuthController) Signup(ctx *gin.Context) {
	var input dto.SignupInput
	// リクエストデータをバインド
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unexpected error"})
	}

	err := c.service.Signup(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to creaate user"})
	}
	ctx.Status(http.StatusOK)
}
