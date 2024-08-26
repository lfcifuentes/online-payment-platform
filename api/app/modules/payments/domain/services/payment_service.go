package services

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lfcifuentes/online-payment-platform/api/app/internal/adapters/pgsql"
	"github.com/lfcifuentes/online-payment-platform/api/app/internal/adapters/validator"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/payments/application/usecases"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/payments/data/repositories"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/payments/data/valueobjects"
)

type PaymentServices struct {
	Repository                  *repositories.PaymentRepository
	Validator                   *validator.ValidatorAdapter
	PayUseCase                  *usecases.PayUseCase
	TransactionsUseCase         *usecases.TransactionsUseCase
	TransactionsReceivesUseCase *usecases.TransactionsReceivesUseCase
	TransactionUseCase          *usecases.TransactionUseCase
	TransactionRefundUseCase    *usecases.TransactionRefundUseCase
}

func NewPaymentServices(db *pgsql.DBAdapter, validator *validator.ValidatorAdapter) *PaymentServices {
	repo := repositories.NewPaymentRepository(db)
	// Add your code here
	return &PaymentServices{
		Validator:                   validator,
		PayUseCase:                  usecases.NewPayUseCase(repo),
		TransactionsUseCase:         usecases.NewTransactionsUseCase(repo),
		TransactionUseCase:          usecases.NewTransactionUseCase(repo),
		TransactionsReceivesUseCase: usecases.NewTransactionsReceivesUseCase(repo),
		TransactionRefundUseCase:    usecases.NewTransactionRefundUseCase(repo),
	}
}

func GetUserFromContext(c *gin.Context) (int64, error) {
	value, exists := c.Get("userID")
	if !exists {
		return 0, errors.New("failed to read user ID from context")
	}
	idInt, err := strconv.ParseInt(fmt.Sprintf("%v", value), 10, 64)
	if err != nil {
		return 0, err
	}
	return idInt, nil
}

// Pay Pagar un producto
//
//	@Summary		Pagar un producto
//	@Description	Pagar un producto
//	@Tags			Payments
//	@Param			payment	body	valueobjects.PayParams	true	"Payment"
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	{object}	gin.H{data=nil}
//	@Failure		400	{object}	gin.H{data=nil}
//	@Failure		401	{object}	gin.H{data=nil}
//	@Failure		404	{object}	gin.H{data=nil}
//	@Failure		500	{object}	gin.H{data=nil}
//	@Router			/payments/pay [post]
func (s *PaymentServices) Pay(c *gin.Context) {
	var query valueobjects.PayParams
	_ = c.ShouldBindJSON(&query)
	userID, err := GetUserFromContext(c)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err},
		)
		return
	}

	errStruct := s.Validator.ValidateStruct(query)
	if errStruct != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": errStruct},
		)
		return
	}

	data, err := s.PayUseCase.Pay(
		strconv.FormatInt(userID, 10),
		query,
	)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{"data": data},
	)
}

// GetTransactions Obtener transacciones
//
//	@Summary		Obtener transacciones
//	@Description	Obtener transacciones
//	@Tags			Payments
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	{object}	gin.H{data=[]}
//	@Failure		400	{object}	gin.H{data=nil}
//	@Failure		401	{object}	gin.H{data=nil}
//	@Failure		404	{object}	gin.H{data=nil}
//	@Failure		500	{object}	gin.H{data=nil}\
//	@Router			/payments [get]
func (s *PaymentServices) GetTransactions(c *gin.Context) {
	userID, err := GetUserFromContext(c)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err},
		)
		return
	}

	data, err := s.TransactionsUseCase.Execute(
		userID,
	)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{"data": data},
	)
}

// GetTransaction Obtener transacción
//
//	@Summary		Obtener transacción
//	@Description	Obtener transacción
//	@Tags			Payments
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			id	path	int	true	"ID"
//	@Success		200	{object}	gin.H{data=nil}
//	@Failure		400	{object}	gin.H{data=nil}
//	@Failure		401	{object}	gin.H{data=nil}
//	@Failure		404	{object}	gin.H{data=nil}
//	@Failure		500	{object}	gin.H{data=nil}
//	@Router			/payments/{id} [get]
func (s *PaymentServices) GetTransaction(c *gin.Context) {
	userID, err := GetUserFromContext(c)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err},
		)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	data, err := s.TransactionUseCase.Execute(
		userID,
		id,
	)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{"data": data},
	)
}

func (s *PaymentServices) Receives(c *gin.Context) {
	userID, err := GetUserFromContext(c)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err},
		)
		return
	}

	data, err := s.TransactionsReceivesUseCase.Execute(
		userID,
	)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{"data": data},
	)
}

// Refund Reembolsar
func (s *PaymentServices) Refund(c *gin.Context) {
	userID, err := GetUserFromContext(c)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err},
		)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	data, err := s.TransactionRefundUseCase.Execute(
		userID,
		id,
	)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{"data": data},
	)
}
