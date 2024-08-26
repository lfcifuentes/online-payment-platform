package usecases

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/lfcifuentes/online-payment-platform/api/app/modules/paymentmethods/data/models"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/paymentmethods/data/repositories"
	"github.com/lfcifuentes/online-payment-platform/api/app/pkg"
)

type CreatePaymentMethodsUseCase struct {
	repo    *repositories.PaymentMethodRepository
	bankApi pkg.BankApi
}

func NewCreatePaymentMethodsUseCase(repo *repositories.PaymentMethodRepository) *CreatePaymentMethodsUseCase {
	return &CreatePaymentMethodsUseCase{
		repo:    repo,
		bankApi: *pkg.NewBankApi(),
	}
}

type ResponseData struct {
	ID int64 `json:"id"`
}

type CreatePaymentMethodResponse struct {
	Data ResponseData `json:"data"`
}

func (uc *CreatePaymentMethodsUseCase) CreatePaymentMethod(name, brand, cardNumber string, userID int64, expMonth, expYear int) error {
	// last 4 digits
	last4 := cardNumber[len(cardNumber)-4:]
	payload := []byte(`{
		"bank_id": ` + uc.bankApi.GetBankID() + `,
		"brand": "` + brand + `",
		"card_number": "` + cardNumber + `",
		"exp_month": ` + strconv.Itoa(expMonth) + `,
		"exp_year": ` + strconv.Itoa(expYear) + `
	}`)

	dataString, err := uc.bankApi.Post(
		fmt.Sprintf("/payment-methods/%v", userID),
		payload,
	)
	if err != nil {
		return err
	}

	// actualizo el id del cliente en la base de datos
	var response CreatePaymentMethodResponse
	err = json.Unmarshal([]byte(dataString), &response)
	if err != nil {
		return err
	}

	last4Int, err := strconv.Atoi(last4)
	if err != nil {
		return err
	}

	bankID, err := strconv.ParseInt(uc.bankApi.GetBankID(), 10, 64)
	if err != nil {
		return err
	}

	paymentMethod := models.PaymentMethod{
		Name:     name,
		UserID:   userID,
		Brand:    brand,
		Last4:    last4Int,
		ExpMonth: expMonth,
		ExpYear:  expYear,
		BankID:   bankID,
		BankCode: response.Data.ID,
	}

	err = uc.repo.CreatePaymentMethod(paymentMethod)
	if err != nil {
		return err
	}

	return nil
}
