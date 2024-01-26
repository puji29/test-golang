package controller

import (
	"fmt"
	"net/http"
	"simple_payments/config"
	"simple_payments/entity/dto"
	"simple_payments/usecase"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	paymentUseCase usecase.PaymentUseCase
	rg             *gin.RouterGroup
}

func (p *PaymentController) PaymentHandler(c *gin.Context) {
	var data dto.PaymentDto

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
	}

	err := p.paymentUseCase.MakePayment(data.From, data.To, data.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("payment succesfully. %s tranfer %f to %s", data.From, data.Amount, data.To)})
}

func (p *PaymentController) TransactionHistoryHandle(c *gin.Context) {

	transactions, err := p.paymentUseCase.GetTransactionHistory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (p *PaymentController) Route() {
	p.rg.POST(config.Transaction, p.PaymentHandler)
	p.rg.GET(config.History, p.TransactionHistoryHandle)
}

func NewPaymentController(paymentUseCase usecase.PaymentUseCase, rg *gin.RouterGroup) *PaymentController {
	return &PaymentController{paymentUseCase: paymentUseCase, rg: rg}

}
