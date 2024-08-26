package usecases

import (
	"fmt"

	"github.com/lfcifuentes/online-payment-platform/api/app/modules/paymentmethods/data/repositories"
	"github.com/lfcifuentes/online-payment-platform/api/app/pkg"
)

type DeletePaymentMethodUseCase struct {
	repo    *repositories.PaymentMethodRepository
	bankApi pkg.BankApi
}

func NewDeletePaymentMethodUseCase(repo *repositories.PaymentMethodRepository) *DeletePaymentMethodUseCase {
	return &DeletePaymentMethodUseCase{
		repo:    repo,
		bankApi: *pkg.NewBankApi(),
	}
}

func (u *DeletePaymentMethodUseCase) Execute(userID, payment_method_id string) error {
	// check if the user payment method exists
	data, err := u.repo.GetPaymentMethodByID(userID, payment_method_id)
	if err != nil {
		return err
	}

	_, err = u.bankApi.Delete(fmt.Sprintf("/payment-methods/%v/%v", userID, data.BankCode))
	if err != nil {
		return err
	}
	// delete the user payment method
	return u.repo.DeletePaymentMethod(userID, payment_method_id)
}
