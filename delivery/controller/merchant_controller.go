package controller

import (
	"fmt"
	"net/http"
	"simple_payments/config"
	"simple_payments/usecase"

	"github.com/gin-gonic/gin"
)

type MerchantController struct {
	merchantUc usecase.MerchantUseCase
	rg         *gin.RouterGroup
}

func (m *MerchantController) RegisterMerchantHandler(c *gin.Context) {

	name := c.PostForm("name")
	password := c.PostForm("password")

	err := m.merchantUc.RegisterMer(name, password)
	fmt.Println(err, "reg")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "register failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("merchant register succesfully. Welcome, %s!", name)})
}

func (a *MerchantController) Route() {
	a.rg.POST(config.RegisterMerchant, a.RegisterMerchantHandler)

}

func NewMerchantController(merchantUc usecase.MerchantUseCase, rg *gin.RouterGroup) *MerchantController {
	return &MerchantController{merchantUc: merchantUc, rg: rg}

}
