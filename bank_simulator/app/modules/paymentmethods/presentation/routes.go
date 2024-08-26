package presentation

import (
	"github.com/lfcifuentes/online-payment-platform/bank/app/infrastructure"
	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/paymentmethods/domain/services"
)

func RegisterPaymentMethodsRoutes(a *infrastructure.Application) {
	paymentmethodService := services.NewPaymentMethodService(a.DbAdapter, a.Validator)

	paymentMethodsRoutes := a.Router.Group("payment-methods")

	paymentMethodsRoutes.GET(":user_code", paymentmethodService.ListPaymentMethods)
	paymentMethodsRoutes.POST(":user_code", paymentmethodService.CreatePaymentMethod)
	paymentMethodsRoutes.DELETE(":user_code/:payment_method_id", paymentmethodService.DeletePaymentMethod)
}
