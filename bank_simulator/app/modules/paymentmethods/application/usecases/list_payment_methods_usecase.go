package usecases

import (
	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/paymentmethods/data/models"
	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/paymentmethods/data/repositories"
)

type ListPaymentMethodsUseCase struct {
	repo *repositories.PaymentMethodRepository
}

func NewListPaymentMethodsUseCase(repo *repositories.PaymentMethodRepository) *ListPaymentMethodsUseCase {
	return &ListPaymentMethodsUseCase{
		repo: repo,
	}
}

func (u *ListPaymentMethodsUseCase) Execute(userId string) ([]models.PaymentMethod, error) {
	paymentMethods, err := u.repo.GetPaymentMethodsByUserCode(userId)
	if err != nil {
		return nil, err
	}

	return paymentMethods, nil
}
