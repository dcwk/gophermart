package services

import (
	"context"
	"fmt"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/dcwk/gophermart/internal/repositories"
)

type GetWithdrawalsService struct {
	UserRepository       repositories.UserRepository
	WithdrawalRepository repositories.WithdrawalRepository
}

func NewGetWithdrawalsService(
	userRepository repositories.UserRepository,
	WithdrawalRepository repositories.WithdrawalRepository,
) *GetWithdrawalsService {
	return &GetWithdrawalsService{
		UserRepository:       userRepository,
		WithdrawalRepository: WithdrawalRepository,
	}
}

func (s *GetWithdrawalsService) Handle(ctx context.Context, userID int64) ([]*models.Withdrawal, error) {
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
