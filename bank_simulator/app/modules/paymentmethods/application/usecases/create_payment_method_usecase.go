package usecases

import (
	"strconv"

	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/paymentmethods/data/models"
	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/paymentmethods/data/repositories"
	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/paymentmethods/data/valueobjects"
)

type CreatePaymentMethodsUseCase struct {
	repo *repositories.PaymentMethodRepository
}

func NewCreatePaymentMethodsUseCase(repo *repositories.PaymentMethodRepository) *CreatePaymentMethodsUseCase {
	return &CreatePaymentMethodsUseCase{
		repo: repo,
	}
}

func (u *CreatePaymentMethodsUseCase) Execute(userId string, request valueobjects.NewPaymentMethodRequest) (*models.PaymentMethod, error) {
	userID, err := strconv.Atoi(userId)
	if err != nil {
		return nil, err
	}
	last4 := request.CardNumber[len(request.CardNumber)-4:]
	paymentmethod := models.PaymentMethod{
		UserID:     userID,
		Brand:      request.Brand,
		Last4:      last4,
		ExpMonth:   request.ExpMonth,
		ExpYear:    request.ExpYear,
		BankID:     request.BankID,
		CardNumber: request.CardNumber,
	}
	err = u.repo.CreatePaymentMethod(&paymentmethod)
	if err != nil {
		return nil, err
	}

	return &paymentmethod, nil
}
