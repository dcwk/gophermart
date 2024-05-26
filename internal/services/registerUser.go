package services

import (
	"context"
	"fmt"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/dcwk/gophermart/internal/repositories"
)

type RegisterUserService struct {
	UserRepository        repositories.UserRepository
	UserBalanceRepository repositories.UserBalanceRepository
}

func NewRegisterUserService(
	userRepository repositories.UserRepository,
	userBalanceRepository repositories.UserBalanceRepository,
) *RegisterUserService {
	return &RegisterUserService{
		UserRepository:        userRepository,
		UserBalanceRepository: userBalanceRepository,
	}
}

func (regUs *RegisterUserService) CreateUser(ctx context.Context, login string, password string) (*models.User, error) {
	_, err := regUs.UserRepository.GetUserByLogin(ctx, login)
	if err == nil {
		return nil, fmt.Errorf("user with login %s already exists", login)
	}

	user := models.NewUser(login, password)
	if err := user.HashPassword(); err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	user, err = regUs.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("could not create user: %v", err)
	}

	userBalance := models.NewUserBalance(user.ID)
	_, err = regUs.UserBalanceRepository.Create(ctx, userBalance)
	if err != nil {
		return nil, fmt.Errorf("could not create user balance: %v", err)
	}

	return user, nil
}
