package usecase

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/dcwk/gophermart/internal/repositories"
)

const NotEnoughPoints = "NotEnoughPoints"

type WithdrawRequestHandler struct {
	Logger                *zap.Logger
	UserRepository        repositories.UserRepository
	UserBalanceRepository repositories.UserBalanceRepository
	OrderRepository       repositories.OrderRepository
	WithdrawalRepository  repositories.WithdrawalRepository
}

func NewWithdrawRequestHandler(
	logger *zap.Logger,
	userRepository repositories.UserRepository,
	userBalanceRepository repositories.UserBalanceRepository,
	orderRepository repositories.OrderRepository,
	withdrawalRepository repositories.WithdrawalRepository,
) *WithdrawRequestHandler {
	return &WithdrawRequestHandler{
		Logger:                logger,
		UserRepository:        userRepository,
		UserBalanceRepository: userBalanceRepository,
		OrderRepository:       orderRepository,
		WithdrawalRepository:  withdrawalRepository,
	}
}

func (s *WithdrawRequestHandler) Handle(
	ctx context.Context,
	userID int64,
	orderNumber string,
	sum float64,
) (string, error) {
	user, err := s.UserRepository.GetUserByID(ctx, userID)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("user %d not found", userID))
		return NotFound, nil
	}

	order := models.NewOrder(user.ID, orderNumber)
	if !order.IsValid() {
		return IncorrectOrderNumber, nil
	}
	order, err = s.OrderRepository.Create(ctx, order)
	if err != nil {
		return InternalError, err
	}

	userBalance, err := s.UserBalanceRepository.GetUserBalanceByID(ctx, user.ID, true)
	if err != nil {
		return InternalError, err
	}
	if !userBalance.IsWithdrawPossible(sum) {
		return NotEnoughPoints, nil
	}
	withdrawal := models.NewWithdrawal(user.ID, order.ID, sum)
	withdrawal, err = s.WithdrawalRepository.Create(ctx, withdrawal)
	if err != nil {
		return InternalError, err
	}
	userBalance.DoWithdrawal(withdrawal.Value)
	err = s.UserBalanceRepository.Update(ctx, userBalance)
	if err != nil {
		return InternalError, err
	}

	return "", nil
}
