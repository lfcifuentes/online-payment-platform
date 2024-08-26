package presentation

import (
	"github.com/lfcifuentes/online-payment-platform/bank/app/infrastructure"
	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/clients/domain/services"
)

func RegisterClientRoutes(a *infrastructure.Application) {
	clientSevice := services.NewClientService(a.DbAdapter, a.Validator)

	a.Router.POST("/clients", clientSevice.CreateClient)
}
