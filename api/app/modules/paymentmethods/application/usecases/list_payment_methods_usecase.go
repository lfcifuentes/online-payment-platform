package usecases

import (
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/paymentmethods/data/models"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/paymentmethods/data/repositories"
)

type ListPaymentMethodsUseCase struct {
	repo *repositories.PaymentMethodRepository
}

func NewListPaymentMethodsUseCase(repo *repositories.PaymentMethodRepository) *ListPaymentMethodsUseCase {
	return &ListPaymentMethodsUseCase{
		repo: repo,
	}
}

func (u *ListPaymentMethodsUseCase) Execute(userID string) ([]*models.PaymentMethod, error) {
	return u.repo.GetPaymentMethodsByUserCode(userID)
}
