package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lfcifuentes/online-payment-platform/api/app/infrastructure"
	"github.com/lfcifuentes/online-payment-platform/api/app/internal/adapters/pgsql"
	AuthPresentation "github.com/lfcifuentes/online-payment-platform/api/app/modules/auth/presentation"
	MerchantPresentation "github.com/lfcifuentes/online-payment-platform/api/app/modules/merchant/presentation"
	PaymentMethodsPresentation "github.com/lfcifuentes/online-payment-platform/api/app/modules/paymentmethods/presentation"
	PaymentsPresentation "github.com/lfcifuentes/online-payment-platform/api/app/modules/payments/presentation"
	"github.com/lfcifuentes/online-payment-platform/api/docs"
	"github.com/spf13/viper"
)

func NewRouter(db *pgsql.DBAdapter) *gin.Engine {
	// Add swagger info
	docs.SwaggerInfo.Host = viper.GetString("APP_HOST")
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Version = viper.GetString("APP_VERSION")

	// create the application
	app := infrastructure.NewApplication(
		db,
	)

	// Add your code here
	AuthPresentation.RegisterAuthRoutes(app)
	PaymentsPresentation.RegisterPaymentsRoutes(app)
	PaymentMethodsPresentation.RegisterPaymentMethodsRoutes(app)
	MerchantPresentation.RegisterMerchantRoutes(app)

	return app.Router

}
