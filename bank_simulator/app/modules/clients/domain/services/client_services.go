package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfcifuentes/online-payment-platform/bank/app/internal/adapters/pgsql"
	"github.com/lfcifuentes/online-payment-platform/bank/app/internal/adapters/validator"
	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/clients/application/usecases"
	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/clients/data/repositories"
	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/clients/data/valueobjects"
)

type ClientService struct {
	CreateClientUseCase *usecases.CreateClientUseCase
	Validator           *validator.ValidatorAdapter
}

func NewClientService(db *pgsql.DBAdapter, validator *validator.ValidatorAdapter) *ClientService {
	repository := repositories.NewClientRepository(db)
	return &ClientService{
		Validator:           validator,
		CreateClientUseCase: usecases.NewCreateClientUseCase(repository),
	}
}

// CreateClient
//
//	@Summary Create a new client
//	@Description Create a new client
//	@Tags clients
//	@Accept json
//	@Produce json
//	@Param client body valueobjects.CreateClientRequest true "Client data"
//	@Success		200	{object}	gin.H{data=nil}
//	@Failure		400	{object}	gin.H{data=nil}
//	@Failure		401	{object}	gin.H{data=nil}
//	@Failure		404	{object}	gin.H{data=nil}
//	@Failure		500	{object}	gin.H{data=nil}
func (cs *ClientService) CreateClient(c *gin.Context) {
	var request valueobjects.CreateClientRequest
	_ = c.ShouldBindJSON(&request)
	errStruct := cs.Validator.ValidateStruct(request)
	if errStruct != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": errStruct},
		)
		return
	}

	client, err := cs.CreateClientUseCase.Execute(request)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "Client created successfully",
			"client":  client,
		},
	)
}
