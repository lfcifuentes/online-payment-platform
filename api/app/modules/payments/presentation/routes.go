package presentation

import (
	"github.com/lfcifuentes/online-payment-platform/api/app/infrastructure"
	"github.com/lfcifuentes/online-payment-platform/api/app/infrastructure/http/middleware"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/payments/domain/services"
)

func RegisterPaymentsRoutes(a *infrastructure.Application) {
	payServices := services.NewPaymentServices(a.DbAdapter, a.Validator)

	paymentsGroup := a.Router.Group("/payments")
	paymentsGroup.Use(middleware.JWTMiddleware(a.DbAdapter))
	paymentsGroup.POST("pay", payServices.Pay)
	paymentsGroup.GET("", payServices.GetTransactions)
	paymentsGroup.GET("receive", payServices.Receives)
	paymentsGroup.GET(":id", payServices.GetTransaction)
	paymentsGroup.POST(":id/refund", payServices.Refund)
}
