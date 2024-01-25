package main

import (
	"simple_payments/delivery"
)

func main() {

	delivery.NewServer().Run()

	// customerRepo := repository.NewInMemoryCustomerRepository()
	// merchantRepo := repository.NewInMemoryMerchantRepository()
	// paymentRepo := repository.NewInMemoryPaymentRepository()

	// authUseCase := usecase.NewAuthUseCase(customerRepo)
	// paymentUC := usecase.NewPaymentUsecase(customerRepo, merchantRepo, paymentRepo)
	// merchantUC := usecase.NewMerchantUseCase(merchantRepo)

	// authController := controller.NewAuthController(authUseCase)
	// paymentController := controller.NewPaymentController(paymentUC)
	// merchantController := controller.NewMerchantController(merchantUC)

	// router := gin.Default()

	// router.POST("/login", authController.LoginHandler)
	// router.POST("/logout", authController.LogoutHander)
	// router.POST("/register", authController.RegisterCustomerHandle)
	// router.POST("/registerMerchant", merchantController.RegisterMerchantHandler)
	// router.POST("/payment", paymentController.PaymentHandler)
	// router.GET("/history", paymentController.TransactionHistoryHandle)
	// // router.GET("/merchant/:name", merchantController.FindByNameHandle)

	// log.Fatal(router.Run(":8080"))
}
