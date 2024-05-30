package services

import (
	"context"
	"fmt"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/dcwk/gophermart/internal/repositories"
	"go.uber.org/zap"
)

const (
	NotFound             = "NotFound"
	OrderAlreadyExists   = "OrderAlreadyExists"
	IncorrectOrderNumber = "IncorrectOrderNumber"
	ForbiddenOrder       = "ForbiddenOrder"
	InternalError        = "InternalError"
	InvalidOrder         = "InvalidOrder"
)

type CreateOrderService struct {
	Logger                *zap.Logger
	UserRepository        repositories.UserRepository
	OrderRepository       repositories.OrderRepository
	AccrualRepository     repositories.AccrualRepository
	UserBalanceRepository repositories.UserBalanceRepository
}

func NewCreateOrderService(
	logger *zap.Logger,
	userRepository repositories.UserRepository,
	orderRepository repositories.OrderRepository,
	accrualRepository repositories.AccrualRepository,
	userBalanceRepository repositories.UserBalanceRepository,
) *CreateOrderService {
	return &CreateOrderService{
		Logger:                logger,
		UserRepository:        userRepository,
		OrderRepository:       orderRepository,
		AccrualRepository:     accrualRepository,
		UserBalanceRepository: userBalanceRepository,
	}
}

func (s *CreateOrderService) Handle(
	ctx context.Context,
	orderChannel <-chan models.AccrualOrder,
	orderNumber string,
	userID int64,
) (string, error) {
	user, err := s.UserRepository.GetUserByID(ctx, userID)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("user %d not found", userID))
		return NotFound, nil
	}

	existingOrder, _ := s.OrderRepository.FindOrderByNumber(ctx, orderNumber)
	if existingOrder != nil && existingOrder.UserID == user.ID {
		return OrderAlreadyExists, nil
	} else if existingOrder != nil && existingOrder.UserID != user.ID {
		return ForbiddenOrder, nil
	}

	order := models.NewOrder(user.ID, orderNumber)
	if !order.IsValid() {
		return IncorrectOrderNumber, nil
	}
	order, err = s.OrderRepository.Create(ctx, order)
	if err != nil {
		return InternalError, fmt.Errorf("could not create order: %v", err)
	}
	accrual := models.NewAccrual(order.ID)
	accrual, err = s.AccrualRepository.Create(ctx, accrual)
	if err != nil {
		return "", fmt.Errorf("could not create accrual: %v", err)
	}

	accrualOrder := <-orderChannel
	if accrualOrder.Status == "" {
		return InternalError, fmt.Errorf("empty response from accrual service")
	}
	accrual.UpdateStatus(accrualOrder.Status, accrualOrder.Accrual)
	accrual, err = s.AccrualRepository.Update(ctx, accrual)
	if err != nil {
		return "", fmt.Errorf("could not update accrual: %v", err)
	}
	if accrual.Value == 0 {
		return "", nil
	}

	userBalance, err := s.UserBalanceRepository.GetUserBalanceByID(ctx, user.ID, true)
	if err != nil {
		return "", fmt.Errorf("could not get user balance: %v", err)
	}
	userBalance.DoAccrual(accrual.Value)
	err = s.UserBalanceRepository.Update(ctx, userBalance)
	if err != nil {
		return "", fmt.Errorf("could not update user balance: %v", err)
	}

	return "", nil
}
