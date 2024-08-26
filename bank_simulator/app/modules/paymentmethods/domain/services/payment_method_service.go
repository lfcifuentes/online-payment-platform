package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfcifuentes/online-payment-platform/bank/app/internal/adapters/pgsql"
	"github.com/lfcifuentes/online-payment-platform/bank/app/internal/adapters/validator"
	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/paymentmethods/application/usecases"
	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/paymentmethods/data/repositories"
	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/paymentmethods/data/valueobjects"
)

type PaymentMethodService struct {
	Repo                       *repositories.PaymentMethodRepository
	Validator                  *validator.ValidatorAdapter
	ListPaymentMethodsUseCase  *usecases.ListPaymentMethodsUseCase
	CreatePaymentMethodUseCase *usecases.CreatePaymentMethodsUseCase
	DeletePaymentMethodUseCase *usecases.DeletePaymentMethodUseCase
}

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

// ListPaymentMethods
// @Summary List payment methods
// @Description List payment methods
// @Tags Payment Methods
// @Accept json
// @Produce json
// @Param user_code path string true "User code"
// @Success 200 {object} []models.PaymentMethod
// @Failure 404 {object} map[string]string
// @Router /payment-methods/{user_code} [get]
func (s *PaymentMethodService) ListPaymentMethods(c *gin.Context) {
	userID := c.Param("user_code")

	data, err := s.ListPaymentMethodsUseCase.Execute(userID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Not Found"})
		return
	}
	c.JSON(200, data)
}

// CreatePaymentMethod
// @Summary Create payment method
// @Description Create payment method
// @Tags Payment Methods
// @Accept json
// @Produce json
// @Param user_code path string true "User code"
// @Param payment_method body models.PaymentMethod true "Payment Method"
// @Success 201 {object} models.PaymentMethod
// @Failure 400 {object} map[string]string
// @Router /payment-methods/{user_code} [post]
func (s *PaymentMethodService) CreatePaymentMethod(c *gin.Context) {
	userID := c.Param("user_code")
	var request valueobjects.NewPaymentMethodRequest
	_ = c.ShouldBindJSON(&request)

	errStruct := s.Validator.ValidateStruct(request)
	if errStruct != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": errStruct},
		)
		return
	}

	data, err := s.CreatePaymentMethodUseCase.Execute(userID, request)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Payment method created",
		"data":    data,
	})
}

// DeletePaymentMethod
// @Summary Delete payment method
// @Description Delete payment method
// @Tags Payment Methods
// @Accept json
// @Produce json
// @Param user_code path string true "User code"
// @Param payment_method_id path string true "Payment Method ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /payment-methods/{user_code}/{payment_method_id} [delete]
func (s *PaymentMethodService) DeletePaymentMethod(c *gin.Context) {
	userID := c.Param("user_code")
	paymentMethodID := c.Param("payment_method_id")

	err := s.DeletePaymentMethodUseCase.Execute(userID, paymentMethodID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, gin.H{
		"message": "Payment method deleted",
	})
}
