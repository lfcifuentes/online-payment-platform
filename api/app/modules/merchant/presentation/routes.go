package presentation

import (
	"github.com/lfcifuentes/online-payment-platform/api/app/infrastructure"
	"github.com/lfcifuentes/online-payment-platform/api/app/infrastructure/http/middleware"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/merchant/domain/services"
)

func RegisterMerchantRoutes(a *infrastructure.Application) {
	merchantService := services.NewMerchantService(a.DbAdapter, a.Validator)

	merchantGroup := a.Router.Group("/merchant")

	// Middleware JWT
	merchantGroup.Use(middleware.JWTMiddleware(a.DbAdapter))
	merchantGroup.POST("", merchantService.CreateMerchant)

}
