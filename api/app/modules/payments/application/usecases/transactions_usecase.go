package usecases

import (
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/paymentmethods/data/models"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/payments/data/repositories"
	"github.com/lfcifuentes/online-payment-platform/api/app/pkg"
)

type TransactionsUseCase struct {
	repo           *repositories.PaymentRepository
	passwordHasher pkg.PasswordHasher
}

func NewTransactionsUseCase(repo *repositories.PaymentRepository) *TransactionsUseCase {
	return &TransactionsUseCase{
		repo:           repo,
		passwordHasher: *pkg.NewPasswordHasher(),
	}
}

func (u *TransactionsUseCase) Execute(userID int64) ([]models.Payment, error) {
	transactions, err := u.repo.GetTransactions(userID)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
