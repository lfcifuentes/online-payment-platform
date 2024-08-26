package presentation

import (
	"github.com/lfcifuentes/online-payment-platform/api/app/infrastructure"
	"github.com/lfcifuentes/online-payment-platform/api/app/infrastructure/http/middleware"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/auth/domain/services"
)

func RegisterAuthRoutes(a *infrastructure.Application) {
	authServices := services.NewAuthServices(a.DbAdapter, a.Validator)

	authGroup := a.Router.Group("/auth")

	authGroup.POST("/login", authServices.Login)
	authGroup.POST("/register", authServices.Register)

	// Middleware JWT
	authGroup.Use(middleware.JWTMiddleware(a.DbAdapter))
	authGroup.POST("/logout", authServices.Logout)

}
