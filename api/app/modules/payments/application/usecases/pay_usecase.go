package usecases

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	merchantsRepo "github.com/lfcifuentes/online-payment-platform/api/app/modules/merchant/data/repositories"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/paymentmethods/data/models"
	paymentsMethodsRepo "github.com/lfcifuentes/online-payment-platform/api/app/modules/paymentmethods/data/repositories"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/payments/data/repositories"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/payments/data/valueobjects"
	"github.com/lfcifuentes/online-payment-platform/api/app/pkg"
)

type PayUseCase struct {
	merchantRepo       *merchantsRepo.MerchantRepository
	paymentMethodsRepo *paymentsMethodsRepo.PaymentMethodRepository
	repo               *repositories.PaymentRepository
	passwordHasher     pkg.PasswordHasher
	jwt                pkg.ApiJWT
	bankApi            pkg.BankApi
}

func NewPayUseCase(repo *repositories.PaymentRepository) *PayUseCase {
	return &PayUseCase{
		repo:               repo,
		merchantRepo:       merchantsRepo.NewMerchantRepository(repo.DBAdapter),
		paymentMethodsRepo: paymentsMethodsRepo.NewPaymentMethodRepository(repo.DBAdapter),
		passwordHasher:     *pkg.NewPasswordHasher(),
		jwt:                *pkg.NewApiJWT(),
		bankApi:            *pkg.NewBankApi(),
	}
}

type PaymentResponse struct {
	ID        int64     `json:"id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseData struct {
	Payment PaymentResponse `json:"payment"`
}

func (u *PayUseCase) Pay(userID string, request valueobjects.PayParams) (*models.Payment, error) {
	// check if merchant exists
	merchant, err := u.merchantRepo.GetMerchantById(request.MerchantID)
	if err != nil {
		return nil, errors.New("merchant not found")
	}
	// check if merchant payment method exists
	paymentMethod, err := u.paymentMethodsRepo.GetPaymentMethodByID(
		userID,
		strconv.FormatInt(request.PaymentMethodID, 10),
	)
	if err != nil {
		return nil, errors.New("payment method not found")
	}
	// check if merchant payment method exists
	paymentMethodMerchant, err := u.paymentMethodsRepo.GetPaymentMethodByID(
		strconv.FormatInt(merchant.UserID, 10),
		strconv.FormatInt(request.MerchantPaymentMethodID, 10),
	)
	if err != nil {
		return nil, errors.New("payment method not found")
	}
	// send payment to bank
	payload := []byte(`{
		"bank_id": ` + u.bankApi.GetBankID() + `,
		"amount": ` + strconv.FormatFloat(request.Amount, 'f', -1, 64) + `,
		"payment_method_id": ` + strconv.FormatInt(paymentMethod.BankCode, 10) + `,
		"receiver_id": ` + strconv.FormatInt(paymentMethodMerchant.BankCode, 10) + `
	}`)

	dataString, err := u.bankApi.Post(
		"/payments/pay",
		payload,
	)
	if err != nil {
		return nil, err
	}
	// actualizo el id del cliente en la base de datos
	var response ResponseData
	err = json.Unmarshal([]byte(dataString), &response)
	if err != nil {
		return nil, errors.New("error parsing bank response")
	}
	// create payment
	err = u.repo.CreatePayment(
		userID,
		request,
		response.Payment.ID,
		response.Payment.Status,
	)
	if err != nil {
		return nil, err
	}
	payment, err := u.repo.GetPaymentByBankPayID(strconv.FormatInt(response.Payment.ID, 10))
	if err != nil {
		return nil, err
	}
	// Add your code here
	return payment, nil
}
