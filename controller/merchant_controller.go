package controller

import (
	"fmt"
	"net/http"
	"simple_payments/usecase"

	"github.com/gin-gonic/gin"
)

type MerchantController struct {
	merchantUc usecase.MerchantUseCase
}

func (m *MerchantController) RegisterMerchantHandler(c *gin.Context) {

	name := c.PostForm("name")
	password := c.PostForm("password")

	err := m.merchantUc.RegisterMer(name, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "register failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("merchant register succesfully. Welcome, %s!", name)})
}

func (m *MerchantController) FindByNameHandle(c *gin.Context) {
	name := c.Param("name")

	if name == "" {
		c.JSON(http.StatusBadRequest, "invalid name")
		return
	}

	err := m.merchantUc.FindByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "name not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"name": name})
}

func NewMerchantController(merchantUc usecase.MerchantUseCase) *MerchantController {
	return &MerchantController{merchantUc: merchantUc}

}
