package usecases

import (
	"errors"

	"github.com/lfcifuentes/online-payment-platform/api/app/modules/merchant/data/repositories"
)

type CreateMerchantUseCase struct {
	repo *repositories.MerchantRepository
}

func NewCreateMerchantUseCase(repo *repositories.MerchantRepository) *CreateMerchantUseCase {
	return &CreateMerchantUseCase{
		repo: repo,
	}
}

func (u *CreateMerchantUseCase) Execute(userID string) error {
	// check if user exists
	_, err := u.repo.GetMerchantByUserID(userID)

	if err == nil {
		return errors.New("Merchant already exists")
	}

	err = u.repo.CreateMerchant(userID)

	if err != nil {
		return err
	}
	// get merchant id
	_, err = u.repo.GetMerchantByUserID(userID)

	if err != nil {
		return err
	}

	return nil
}
