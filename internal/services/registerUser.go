package services

import (
	"context"
	"fmt"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/dcwk/gophermart/internal/repositories"
	"github.com/dcwk/gophermart/internal/utils/auth"
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

func (s *RegisterUserService) Handle(ctx context.Context, login string, password string) (string, error) {
	_, err := s.UserRepository.GetUserByLogin(ctx, login)
	if err == nil {
		return "", fmt.Errorf("user with login %s already exists", login)
	}

	user := models.NewUser(login, password)
	if err := user.HashPassword(); err != nil {
		return "", fmt.Errorf("failed to hash password: %v", err)
	}

	user, err = s.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return "", fmt.Errorf("could not create user: %v", err)
	}

	userBalance := models.NewUserBalance(user.ID)
	_, err = s.UserBalanceRepository.Create(ctx, userBalance)
	if err != nil {
		return "", fmt.Errorf("could not create user balance: %v", err)
	}

	token, err := auth.BuildJWTString(user.ID)
	if err != nil {
		return "", fmt.Errorf("failed to build JWT token: %w", err)
	}

	return token, nil
}
