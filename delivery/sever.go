package delivery

import (
	"fmt"
	"simple_payments/config"
	"simple_payments/delivery/controller"
	"simple_payments/repository"
	"simple_payments/usecase"

	"github.com/gin-gonic/gin"
)

type Server struct {
	authUc     usecase.AuthUseCase
	paymentUc  usecase.PaymentUseCase
	merchantUc usecase.MerchantUseCase
	engine     *gin.Engine
	port       string
}

func (s *Server) initRoute() {
	rg := s.engine.Group(config.ApiGroup)
	controller.NewAuthController(s.authUc, rg).Route()
	controller.NewMerchantController(s.merchantUc, rg).Route()
	controller.NewPaymentController(s.paymentUc, rg).Route()
}

func (s *Server) Run() {
	s.initRoute()
	if err := s.engine.Run(s.port); err != nil {
		panic(fmt.Errorf("server not running on host %s, becauce error %v", s.port, err.Error()))
	}
}

func NewServer() *Server {

	customerRepo := repository.NewInMemoryCustomerRepository()
	merchantRepo := repository.NewInMemoryMerchantRepository()
	paymentRepo := repository.NewInMemoryPaymentRepository()

	authUseCase := usecase.NewAuthUseCase(customerRepo)
	paymentUC := usecase.NewPaymentUsecase(customerRepo, merchantRepo, paymentRepo)
	merchantUC := usecase.NewMerchantUseCase(merchantRepo)
	engine := gin.Default()
	port := ":8080"
	return &Server{
		authUc:     authUseCase,
		paymentUc:  paymentUC,
		merchantUc: merchantUC,
		engine:     engine,
		port:       port,
	}
}
