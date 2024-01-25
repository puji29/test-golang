package controller

import (
	"fmt"
	"net/http"
	"simple_payments/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	paymentUseCase usecase.PaymentUseCase
}

func (p *PaymentController) PaymentHandler(c *gin.Context) {

	from := c.PostForm("from")
	to := c.PostForm("to")
	amount := c.PostForm("amount")
	fmt.Println(from)
	fmt.Println(to)
	fmt.Println(amount)

	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid amount"})
		return
	}
	err = p.paymentUseCase.MakePayment(from, to, amountFloat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("payment succesfully. %s tranfer %f to %s", from, amountFloat, to)})
}

func (p *PaymentController) TransactionHistoryHandle(c *gin.Context) {

	transactions, err := p.paymentUseCase.GetTransactionHistory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func NewPaymentController(paymentUseCase usecase.PaymentUseCase) *PaymentController {
	return &PaymentController{paymentUseCase: paymentUseCase}

}
