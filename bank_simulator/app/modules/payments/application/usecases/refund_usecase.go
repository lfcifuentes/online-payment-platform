package usecases

import (
	"errors"
	"strconv"
	"time"

	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/payments/data/models"
	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/payments/data/repositories"
)

type RefundUseCase struct {
	repo *repositories.PaymentRepository
}

func NewRefundUseCase(repo *repositories.PaymentRepository) *RefundUseCase {
	return &RefundUseCase{
		repo: repo,
	}
}

func (u *RefundUseCase) Execute(id string) (*models.Payment, error) {
	idInt64, _ := strconv.ParseInt(id, 10, 64)
	payment, err := u.repo.GetPaymentByID(idInt64)

	if err != nil {
		return nil, err
	}

	if payment.Status == models.REFUND {
		return nil, errors.New("payment already refunded")
	}

	payment.Status = models.REFUND
	payment.UpdatedAt = time.Now().Add(2 * time.Minute)

	err = u.repo.UpdatePaymentStatus(payment)

	if err != nil {
		return nil, err
	}
	return u.repo.GetPaymentByID(idInt64)
}
