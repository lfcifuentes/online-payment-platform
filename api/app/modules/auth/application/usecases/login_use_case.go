package usecases

import (
	"errors"

	"github.com/lfcifuentes/online-payment-platform/api/app/modules/auth/data/repositories"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/auth/data/valueobjects"
	"github.com/lfcifuentes/online-payment-platform/api/app/pkg"
)

type LoginUseCase struct {
	repo           *repositories.AuthRepository
	passwordHasher pkg.PasswordHasher
	jwt            pkg.ApiJWT
}

func NewLoginUseCase(repo *repositories.AuthRepository) *LoginUseCase {
	return &LoginUseCase{
		repo:           repo,
		passwordHasher: *pkg.NewPasswordHasher(),
		jwt:            *pkg.NewApiJWT(),
	}
}

func (uc *LoginUseCase) Login(email, password string) (*valueobjects.LoginResponse, error) {
	user, err := uc.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if !uc.passwordHasher.Check(password, user.Password) {
		return nil, errors.New("invalid password")
	}

	token, expiration, err := uc.jwt.GenerateJWT(user.ID)

	if err != nil {
		return nil, errors.New("error generating token")
	}

	err = uc.repo.SaveToken(user.ID, token, expiration)

	if err != nil {
		return nil, errors.New("error saving token")
	}

	// Return user data or token as needed
	return &valueobjects.LoginResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   expiration,
	}, nil
}
