package usecase

import (
	"context"
	"fmt"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/dcwk/gophermart/internal/repositories"
)

type GetWithdrawalsHandler struct {
	UserRepository       repositories.UserRepository
	WithdrawalRepository repositories.WithdrawalRepository
}

func NewGetWithdrawalsHandler(
	userRepository repositories.UserRepository,
	WithdrawalRepository repositories.WithdrawalRepository,
) *GetWithdrawalsHandler {
	return &GetWithdrawalsHandler{
		UserRepository:       userRepository,
		WithdrawalRepository: WithdrawalRepository,
	}
}

func (s *GetWithdrawalsHandler) Handle(ctx context.Context, userID int64) ([]*models.Withdrawal, error) {
	user, err := s.UserRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user %d not found", userID)
	}

	withdrawals, err := s.WithdrawalRepository.FindUserWithdrawals(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return withdrawals, nil
}
