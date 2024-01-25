package controller

import (
	"net/http"
	"simple_payments/config"
	"simple_payments/usecase"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUseCase usecase.AuthUseCase
	rg          *gin.RouterGroup
}

func (a *AuthController) LoginHandler(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	message, err := a.authUseCase.Login(username, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func (a *AuthController) LogoutHander(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Logout succesfull"})

}

func (a *AuthController) RegisterCustomerHandle(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	err := a.authUseCase.Register(username, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Register failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "register succcesfully"})

}

func (a *AuthController) Route() {
	a.rg.POST(config.Login, a.LoginHandler)
	a.rg.POST(config.Register, a.RegisterCustomerHandle)
	a.rg.POST(config.Logout, a.LogoutHander)
}

func NewAuthController(authUseCase usecase.AuthUseCase, rg *gin.RouterGroup) *AuthController {
	return &AuthController{authUseCase: authUseCase, rg: rg}
}
