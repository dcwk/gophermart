package services

import (
	"context"
	"fmt"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/dcwk/gophermart/internal/repositories"
)

type GetUserBalanceService struct {
	UserRepository        repositories.UserRepository
	UserBalanceRepository repositories.UserBalanceRepository
}

func NewGetUserBalanceService(
	UserRepository repositories.UserRepository,
	UserBalanceRepository repositories.UserBalanceRepository,
) *GetUserBalanceService {
	return &GetUserBalanceService{
		UserRepository:        UserRepository,
		UserBalanceRepository: UserBalanceRepository,
	}
}

func (s *GetUserBalanceService) Handle(ctx context.Context, userID int64) (*models.UserBalance, error) {
	user, err := s.UserRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user %d not found", userID)
	}

	userBalance, err := s.UserBalanceRepository.GetUserBalanceByID(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("user balance not found for user %d", userID)
	}

	return userBalance, nil
}
