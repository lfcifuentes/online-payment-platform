package presentation

import (
	"github.com/lfcifuentes/online-payment-platform/api/app/infrastructure"
	"github.com/lfcifuentes/online-payment-platform/api/app/infrastructure/http/middleware"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/paymentmethods/domain/services"
)

func RegisterPaymentMethodsRoutes(a *infrastructure.Application) {
	paymentMethodServices := services.NewPaymentMethodService(a.DbAdapter, a.Validator)

	paymentMethodsGroup := a.Router.Group("/payment-methods")
	paymentMethodsGroup.Use(middleware.JWTMiddleware(a.DbAdapter))
	paymentMethodsGroup.GET("/", paymentMethodServices.ListPaymentMethods)
	paymentMethodsGroup.POST("/", paymentMethodServices.CreatePaymentMethod)
	paymentMethodsGroup.DELETE("/:id", paymentMethodServices.DeletePaymentMethod)
}
