package presentation

import (
	"github.com/lfcifuentes/online-payment-platform/bank/app/infrastructure"
	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/payments/domain/services"
)

func RegisterPayRoutes(a *infrastructure.Application) {
	paymentService := services.NewPaymentService(a.DbAdapter, a.Validator)

	paymentMethodsRoutes := a.Router.Group("payments")

	paymentMethodsRoutes.POST("pay", paymentService.Pay)
	paymentMethodsRoutes.POST("refund/:id", paymentService.Refund)
}
