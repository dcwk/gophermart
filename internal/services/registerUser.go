package services

import (
	"context"
	"fmt"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/dcwk/gophermart/internal/repositories"
)

type RegisterUserService struct {
	UserRepository repositories.UserRepository
}

func NewRegisterUserService(userRepository repositories.UserRepository) *RegisterUserService {
	return &RegisterUserService{
		UserRepository: userRepository,
	}
}

func (regUs *RegisterUserService) CreateUser(ctx context.Context, login string, password string) (*models.User, error) {
	_, err := regUs.UserRepository.FindUserByLogin(ctx, login)
	if err == nil {
		return nil, fmt.Errorf("user with login %s already exists", login)

	}

	user := &models.User{
		Login:    login,
		Password: password,
	}
	if err := user.HashPassword(); err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	user, err = regUs.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("could not create user: %v", err)
	}

	return user, nil
}
