package usecase

import (
	"context"
	"fmt"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/dcwk/gophermart/internal/repositories"
)

type GetOrdersHandler struct {
	UserRepository  repositories.UserRepository
	OrderRepository repositories.OrderRepository
}

func NewGetOrdersHandler(
	userRepository repositories.UserRepository,
	orderRepository repositories.OrderRepository,
) *GetOrdersHandler {
	return &GetOrdersHandler{
		UserRepository:  userRepository,
		OrderRepository: orderRepository,
	}
}

func (s *GetOrdersHandler) Handle(ctx context.Context, userID int64) ([]*models.Order, error) {
	user, err := s.UserRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user %d not found", userID)
	}

	orders, err := s.OrderRepository.FindUserOrders(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
