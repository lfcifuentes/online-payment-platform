package infrastructure

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/lfcifuentes/online-payment-platform/bank/app/infrastructure/http/middleware"
	"github.com/lfcifuentes/online-payment-platform/bank/app/internal/adapters/pgsql"
	"github.com/lfcifuentes/online-payment-platform/bank/app/internal/adapters/validator"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Application encapsula el router y el adaptador de la base de datos.
type Application struct {
	Router    *gin.Engine
	DbAdapter *pgsql.DBAdapter
	Validator *validator.ValidatorAdapter
}

func NewApplication(dbAdapter *pgsql.DBAdapter) *Application {
	slog.Info("Creating application instance", "event", "creating_application_instance")
	router := gin.New()
	router.UseRawPath = false
	router.RedirectTrailingSlash = false

	// add middlewares
	router = middleware.StartAllMiddlewares(router)

	router.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	slog.Info("Creating validator instance", "event", "creating_validator_instance")
	validatorInstance := validator.NewValidatorAdapter()
	slog.Info("Validator instance created successfully", "event", "created_validator_instance")

	slog.Info("Application instance created successfully", "event", "created_application_instance")

	return &Application{
		Router:    router,
		DbAdapter: dbAdapter,
		Validator: validatorInstance,
	}
}
