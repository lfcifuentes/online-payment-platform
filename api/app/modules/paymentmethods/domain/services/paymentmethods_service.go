package services

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lfcifuentes/online-payment-platform/api/app/internal/adapters/pgsql"
	"github.com/lfcifuentes/online-payment-platform/api/app/internal/adapters/validator"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/paymentmethods/application/usecases"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/paymentmethods/data/repositories"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/paymentmethods/data/valueobjects"
)

type PaymentMethodService struct {
	Repo                       *repositories.PaymentMethodRepository
	Validator                  *validator.ValidatorAdapter
	ListPaymentMethodsUseCase  *usecases.ListPaymentMethodsUseCase
	CreatePaymentMethodUseCase *usecases.CreatePaymentMethodsUseCase
	DeletePaymentMethodUseCase *usecases.DeletePaymentMethodUseCase
}

// NewPaymentMethodService creates a new PaymentMethodService
func NewPaymentMethodService(db *pgsql.DBAdapter, validator *validator.ValidatorAdapter) *PaymentMethodService {
	repo := repositories.NewPaymentMethodRepository(db)
	return &PaymentMethodService{
		Repo:                       repo,
		Validator:                  validator,
		ListPaymentMethodsUseCase:  usecases.NewListPaymentMethodsUseCase(repo),
		CreatePaymentMethodUseCase: usecases.NewCreatePaymentMethodsUseCase(repo),
		DeletePaymentMethodUseCase: usecases.NewDeletePaymentMethodUseCase(repo),
	}
}

// ListPaymentMethods lists payment methods
// @Summary List payment methods
// @Description List payment methods
// @Tags Payment Methods
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} []models.PaymentMethod
// @Failure 404 {object} string
// @Router /payment-methods [get]
func (s *PaymentMethodService) ListPaymentMethods(c *gin.Context) {
	// get user claim from JWT
	userID, _ := c.Get("userID")
	userIDStr := fmt.Sprintf("%v", userID)

	data, err := s.ListPaymentMethodsUseCase.Execute(userIDStr)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, data)
}

// CreatePaymentMethod creates a payment method
// @Summary Create a payment method
// @Description Create a payment method
// @Tags Payment Methods
// @Accept json
// @Produce json
// @Security Bearer
// @Param paymentMethod body valueobjects.NewPaymentMethodRequest true "Payment Method"
// @Success 201 {object} string
// @Failure 400 {object} string
// @Router /payment-methods [post]
func (s *PaymentMethodService) CreatePaymentMethod(c *gin.Context) {
	// get user claim from JWT
	userID, _ := c.Get("userID")
	userIDStr := fmt.Sprintf("%v", userID)

	var request valueobjects.NewPaymentMethodRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	errStruct := s.Validator.ValidateStruct(request)
	if errStruct != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": errStruct},
		)
		return
	}

	userIDInt64, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = s.CreatePaymentMethodUseCase.CreatePaymentMethod(
		request.Name,
		request.Brand,
		request.CardNumber,
		userIDInt64,
		request.ExpMonth,
		request.ExpYear,
	)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "Payment method created"})
}

// DeletePaymentMethod deletes a payment method
// @Summary Delete a payment method
// @Description Delete a payment method
// @Tags Payment Methods
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Payment Method ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /payment-methods/{id} [delete]
func (s *PaymentMethodService) DeletePaymentMethod(c *gin.Context) {
	// get user claim from JWT
	userID, _ := c.Get("userID")
	userIDStr := fmt.Sprintf("%v", userID)

	id := c.Param("id")

	err := s.DeletePaymentMethodUseCase.Execute(userIDStr, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Payment method deleted"})
}
