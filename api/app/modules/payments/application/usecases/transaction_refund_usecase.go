package usecases

import (
	"errors"
	"fmt"

	"github.com/lfcifuentes/online-payment-platform/api/app/modules/paymentmethods/data/models"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/payments/data/repositories"
	"github.com/lfcifuentes/online-payment-platform/api/app/pkg"
)

type TransactionRefundUseCase struct {
	repo    *repositories.PaymentRepository
	bankApi pkg.BankApi
}

func NewTransactionRefundUseCase(repo *repositories.PaymentRepository) *TransactionRefundUseCase {
	return &TransactionRefundUseCase{
		repo:    repo,
		bankApi: *pkg.NewBankApi(),
	}
}

func (u *TransactionRefundUseCase) Execute(userID, transactionID int64) (*models.Payment, error) {
	transactions, err := u.repo.GetTransactionReceive(userID, transactionID)
	if err != nil {
		return nil, err
	}
	if transactions.Status == models.PaymentStatusRefunded {
		return nil, errors.New("transaction already refunded")
	}
	// send refund to bank
	_, err = u.bankApi.Post(
		fmt.Sprintf("/payments/refund/%v", transactions.BankReference),
		nil,
	)
	if err != nil {
		return nil, err
	}
	//  update transaction status
	err = u.repo.RefundTransaction(transactions)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
