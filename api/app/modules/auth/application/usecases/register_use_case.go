package usecases

import (
	"encoding/json"
	"time"

	"github.com/lfcifuentes/online-payment-platform/api/app/modules/auth/data/repositories"
	"github.com/lfcifuentes/online-payment-platform/api/app/pkg"
)

type RegisterUseCase struct {
	repo           *repositories.AuthRepository
	passwordHasher pkg.PasswordHasher
	jwt            pkg.ApiJWT
	bankApi        pkg.BankApi
}

func NewRegisterUseCase(repo *repositories.AuthRepository) *RegisterUseCase {
	return &RegisterUseCase{
		repo:           repo,
		passwordHasher: *pkg.NewPasswordHasher(),
		jwt:            *pkg.NewApiJWT(),
		bankApi:        *pkg.NewBankApi(),
	}
}

type RegisterBankResponseClient struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Code      string    `json:"code"`
	BankID    int64     `json:"bank_id"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterBankResponse struct {
	Client RegisterBankResponseClient `json:"client"`
}

func (uc *RegisterUseCase) Register(username, email, password string) error {
	hashedPassword, err := uc.passwordHasher.Make(password)
	if err != nil {
		return err
	}

	err = uc.repo.Create(username, email, hashedPassword)
	if err != nil {
		return err
	}

	dataString, err := uc.bankApi.Post("/clients", []byte(`{"name":"`+username+`","email":"`+email+`","bank_id":`+uc.bankApi.GetBankID()+`}`))

	if err != nil {
		return err
	}

	// actualizo el id del cliente en la base de datos
	var response RegisterBankResponse
	err = json.Unmarshal([]byte(dataString), &response)
	if err != nil {
		return err
	}
	uc.repo.UpdateClientID(email, response.Client.ID)

	return nil
}
