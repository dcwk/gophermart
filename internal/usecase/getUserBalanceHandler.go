package usecase

import (
	"context"
	"fmt"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/dcwk/gophermart/internal/repositories"
)

type GetUserBalanceHandler struct {
	UserRepository        repositories.UserRepository
	UserBalanceRepository repositories.UserBalanceRepository
}

func NewGetUserBalanceHandler(
	UserRepository repositories.UserRepository,
	UserBalanceRepository repositories.UserBalanceRepository,
) *GetUserBalanceHandler {
	return &GetUserBalanceHandler{
		UserRepository:        UserRepository,
		UserBalanceRepository: UserBalanceRepository,
	}
}

func (s *GetUserBalanceHandler) Handle(ctx context.Context, userID int64) (*models.UserBalance, error) {
	user, err := s.UserRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user %d not found", userID)
	}

	userBalance, err := s.UserBalanceRepository.GetUserBalanceByID(ctx, user.ID, false)
	if err != nil {
		return nil, fmt.Errorf("user balance not found for user %d", userID)
	}

	return userBalance, nil
}
