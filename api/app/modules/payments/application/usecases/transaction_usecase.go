package usecases

import (
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/paymentmethods/data/models"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/payments/data/repositories"
	"github.com/lfcifuentes/online-payment-platform/api/app/pkg"
)

type TransactionUseCase struct {
	repo           *repositories.PaymentRepository
	passwordHasher pkg.PasswordHasher
}

func NewTransactionUseCase(repo *repositories.PaymentRepository) *TransactionUseCase {
	return &TransactionUseCase{
		repo:           repo,
		passwordHasher: *pkg.NewPasswordHasher(),
	}
}

func (u *TransactionUseCase) Execute(userID, transactionID int64) (*models.Payment, error) {
	transactions, err := u.repo.GetTransaction(userID, transactionID)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
