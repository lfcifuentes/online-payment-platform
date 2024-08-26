package usecases

import (
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/paymentmethods/data/models"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/payments/data/repositories"
	"github.com/lfcifuentes/online-payment-platform/api/app/pkg"
)

type TransactionsReceivesUseCase struct {
	repo           *repositories.PaymentRepository
	passwordHasher pkg.PasswordHasher
}

func NewTransactionsReceivesUseCase(repo *repositories.PaymentRepository) *TransactionsReceivesUseCase {
	return &TransactionsReceivesUseCase{
		repo:           repo,
		passwordHasher: *pkg.NewPasswordHasher(),
	}
}

func (u *TransactionsReceivesUseCase) Execute(userID int64) ([]models.Payment, error) {
	transactions, err := u.repo.GetTransactionsReceives(userID)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
