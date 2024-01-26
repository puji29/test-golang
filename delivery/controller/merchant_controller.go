package controller

import (
	"fmt"
	"net/http"
	"simple_payments/config"
	"simple_payments/entity/dto"
	"simple_payments/usecase"

	"github.com/gin-gonic/gin"
)

type MerchantController struct {
	merchantUc usecase.MerchantUseCase
	rg         *gin.RouterGroup
}

func (m *MerchantController) RegisterMerchantHandler(c *gin.Context) {
	var data dto.MerchantDto

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
	}

	err := m.merchantUc.RegisterMer(data.Name, data.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "register failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("merchant register succesfully. Welcome, %s!", data.Name)})
}

func (a *MerchantController) Route() {
	a.rg.POST(config.RegisterMerchant, a.RegisterMerchantHandler)

}

func NewMerchantController(merchantUc usecase.MerchantUseCase, rg *gin.RouterGroup) *MerchantController {
	return &MerchantController{merchantUc: merchantUc, rg: rg}

}
