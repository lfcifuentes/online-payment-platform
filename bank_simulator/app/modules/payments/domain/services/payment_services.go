package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfcifuentes/online-payment-platform/bank/app/internal/adapters/pgsql"
	"github.com/lfcifuentes/online-payment-platform/bank/app/internal/adapters/validator"
	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/payments/application/usecases"
	_ "github.com/lfcifuentes/online-payment-platform/bank/app/modules/payments/data/models"
	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/payments/data/repositories"
	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/payments/data/valueobjects"
)

type PaymentService struct {
	Repo          *repositories.PaymentRepository
	Validator     *validator.ValidatorAdapter
	PayUseCase    *usecases.PayUseCase
	RefundUseCase *usecases.RefundUseCase
}

func NewPaymentService(db *pgsql.DBAdapter, validator *validator.ValidatorAdapter) *PaymentService {
	repo := repositories.NewPaymentRepository(db)
	return &PaymentService{
		Repo:          repo,
		Validator:     validator,
		PayUseCase:    usecases.NewPayUseCase(repo),
		RefundUseCase: usecases.NewRefundUseCase(repo),
	}
}

// Pay
// @Summary Pay
// @Description Pay
// @Tags Payments
// @Accept  json
// @Produce  json
// @Param pay body valueobjects.PayParams true "Pay"
// @Success 200 {object} models.Payment
// @Router /pay [post]
func (s *PaymentService) Pay(c *gin.Context) {
	var request valueobjects.PayParams
	_ = c.ShouldBindJSON(&request)

	errStruct := s.Validator.ValidateStruct(request)
	if errStruct != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": errStruct},
		)
		return
	}
	payment, err := s.PayUseCase.Execute(request)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"payment": payment})
}

// Refund
// @Summary Refund
// @Description Refund
// @Tags Payments
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} models.Payment
// @Router /refund/{id} [post]
func (s *PaymentService) Refund(c *gin.Context) {
	id := c.Param("id")

	payment, err := s.RefundUseCase.Execute(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"payment": payment})
}
