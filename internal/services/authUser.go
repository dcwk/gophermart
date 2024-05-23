package services

import "github.com/dcwk/gophermart/internal/repositories"

type AuthUserService struct {
	UserRepository repositories.UserRepository
}

func NewAuthService(userRepository repositories.UserRepository) *AuthUserService {
	return &AuthUserService{
		UserRepository: userRepository,
	}
}
