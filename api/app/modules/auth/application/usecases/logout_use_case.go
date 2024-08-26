package usecases

import (
	"errors"

	"github.com/lfcifuentes/online-payment-platform/api/app/modules/auth/data/repositories"
	"github.com/lfcifuentes/online-payment-platform/api/app/pkg"
)

type LogoutUseCase struct {
	repo *repositories.AuthRepository
	jwt  pkg.ApiJWT
}

func NewLogoutUseCase(repo *repositories.AuthRepository) *LogoutUseCase {
	// Add your code here
	return &LogoutUseCase{
		repo: repo,
		jwt:  *pkg.NewApiJWT(),
	}
}

func (uc *LogoutUseCase) Logout(token string) error {
	_, token, err := uc.jwt.ValidateJWT(token)
	if err != nil {
		return err
	}

	tokenStatus, err := uc.repo.ValidateToken(token)
	if err != nil {
		return err
	}

	if !tokenStatus {
		return errors.New("invalid token")
	}

	err = uc.repo.InvalidateToken(token)
	if err != nil {
		return err
	}
	// Add your code here
	return nil
}
