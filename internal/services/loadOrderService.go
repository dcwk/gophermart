package services

import "github.com/dcwk/gophermart/internal/repositories"

type LoadOrderService struct {
	UserRepository        repositories.UserRepository
	OrderRepository       repositories.OrderRepository
	AccrualRepository     repositories.AccrualRepository
	UserBalanceRepository repositories.UserBalanceRepository
}

func NewLoadOrderService(
	userRepository repositories.UserRepository,
	orderRepository repositories.OrderRepository,
	accrualRepository repositories.AccrualRepository,
	userBalanceRepository repositories.UserBalanceRepository,
) *LoadOrderService {
	return &LoadOrderService{
		UserRepository:        userRepository,
		OrderRepository:       orderRepository,
		AccrualRepository:     accrualRepository,
		UserBalanceRepository: userBalanceRepository,
	}
}

func (s *LoadOrderService) Handle() {
}
