package usecases

import (
	"time"

	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/payments/data/models"
	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/payments/data/repositories"
	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/payments/data/valueobjects"
	"golang.org/x/exp/rand"
)

type PayUseCase struct {
	repo *repositories.PaymentRepository
}

func NewPayUseCase(repo *repositories.PaymentRepository) *PayUseCase {
	return &PayUseCase{
		repo: repo,
	}
}

func (u *PayUseCase) Execute(request valueobjects.PayParams) (*models.Payment, error) {

	payment := models.Payment{
		BankID:          request.BankID,
		PaymentMethodID: request.PaymentMethodID,
		ReceiverID:      request.ReceiverID,
		Amount:          request.Amount,
		Status:          models.PENDING,
	}
	err := u.repo.CreatePayment(&payment)
	if err != nil {
		return nil, err
	}
	// Approve or reject payment randomly
	random := rand.Intn(2)

	if random == 0 {
		payment.Status = models.APPROVED
		payment.UpdatedAt = time.Now().Add(2 * time.Minute)
	} else {
		payment.Status = models.REJECTED
		payment.UpdatedAt = time.Now().Add(2 * time.Minute)
	}

	err = u.repo.UpdatePaymentStatus(&payment)

	if err != nil {
		return nil, err
	}

	return &payment, nil
}
