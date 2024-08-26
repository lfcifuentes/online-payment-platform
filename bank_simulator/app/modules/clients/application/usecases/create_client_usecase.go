package usecases

import (
	"errors"
	"strconv"

	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/clients/data/models"
	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/clients/data/repositories"
	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/clients/data/valueobjects"
)

const (
	ErrClientCodeExists = "client code already exists"
)

type CreateClientUseCase struct {
	ClientRepository *repositories.ClientRepository
}

func NewCreateClientUseCase(clientRepository *repositories.ClientRepository) *CreateClientUseCase {
	return &CreateClientUseCase{
		ClientRepository: clientRepository,
	}
}

func (ccu *CreateClientUseCase) Execute(request valueobjects.CreateClientRequest) (*models.Client, error) {
	client := &models.Client{
		Name:   request.Name,
		Email:  request.Email,
		BankID: request.BankID,
	}
	// create client code
	code := "CL-" + strconv.FormatInt(client.BankID, 10) + client.Name[:3] + client.Email[:3]
	client.Code = code
	// check if client code exists
	_, err := ccu.ClientRepository.GetClientByCode(client.Code)
	if err == nil {
		return nil, errors.New(ErrClientCodeExists)
	}
	err = ccu.ClientRepository.CreateClient(client)

	if err != nil {
		return nil, err
	}
	client, err = ccu.ClientRepository.GetClientByCode(client.Code)

	if err != nil {
		return nil, err
	}

	return client, nil
}
