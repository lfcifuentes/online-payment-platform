package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lfcifuentes/online-payment-platform/bank/app/infrastructure"
	"github.com/lfcifuentes/online-payment-platform/bank/app/internal/adapters/pgsql"
	ClientsPresentation "github.com/lfcifuentes/online-payment-platform/bank/app/modules/clients/presentation"
	PaymentMethodsPresentation "github.com/lfcifuentes/online-payment-platform/bank/app/modules/paymentmethods/presentation"
	PaymentsPresentation "github.com/lfcifuentes/online-payment-platform/bank/app/modules/payments/presentation"
	"github.com/lfcifuentes/online-payment-platform/bank/docs"
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

	app.Router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "ok",
		})
	})

	ClientsPresentation.RegisterClientRoutes(app)
	PaymentMethodsPresentation.RegisterPaymentMethodsRoutes(app)
	PaymentsPresentation.RegisterPayRoutes(app)

	return app.Router

}
