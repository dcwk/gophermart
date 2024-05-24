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

func (regUs *RegisterUserService) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	alreadyExistUser, err := regUs.UserRepository.FindUserByLogin(ctx, user.Login)
	if err != nil {
		return nil, err
	}

	if alreadyExistUser.ID != 0 {
		return nil, fmt.Errorf("user with login %s already exists", user.Login)
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
