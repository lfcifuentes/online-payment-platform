package usecases

import (
	"strconv"

	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/paymentmethods/data/repositories"
)

type DeletePaymentMethodUseCase struct {
	repo *repositories.PaymentMethodRepository
}

func NewDeletePaymentMethodUseCase(repo *repositories.PaymentMethodRepository) *DeletePaymentMethodUseCase {
	return &DeletePaymentMethodUseCase{
		repo: repo,
	}
}

func (u *DeletePaymentMethodUseCase) Execute(userId, paymentMethodId string) error {
	userID, err := strconv.Atoi(userId)
	if err != nil {
		return err
	}
	paymentMethodID, err := strconv.Atoi(paymentMethodId)
	if err != nil {
		return err
	}

	// check if the payment method exists
	_, err = u.repo.GetPaymentMethodByID(paymentMethodID)
	if err != nil {
		return err
	}

	err = u.repo.DeletePaymentMethod(userID, paymentMethodID)
	if err != nil {
		return err
	}

	return nil
}
