package services

import (
	"context"

	"github.com/dcwk/gophermart/internal/repositories"
	"go.uber.org/zap"
)

type WithdrawRequestService struct {
	Logger                *zap.Logger
	UserRepository        repositories.UserRepository
	UserBalanceRepository repositories.UserBalanceRepository
	OrderRepository       repositories.OrderRepository
	WithdrawalRepository  repositories.WithdrawalRepository
}

func NewWithdrawRequestService(
	logger *zap.Logger,
	userRepository repositories.UserRepository,
	userBalanceRepository repositories.UserBalanceRepository,
	orderRepository repositories.OrderRepository,
	withdrawalRepository repositories.WithdrawalRepository,
) *WithdrawRequestService {
	return &WithdrawRequestService{
		Logger:                logger,
		UserRepository:        userRepository,
		UserBalanceRepository: userBalanceRepository,
		OrderRepository:       orderRepository,
		WithdrawalRepository:  withdrawalRepository,
	}
}

func (s *WithdrawRequestService) Handle(
	ctx context.Context,
	userID int64,
	orderNumber string,
	sum float64,
) (string, error) {
	return "", nil
}
