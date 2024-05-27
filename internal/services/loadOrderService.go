package services

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/dcwk/gophermart/internal/repositories"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type LoadOrderService struct {
	AccrualSystemAddress  string
	Logger                *zap.Logger
	UserRepository        repositories.UserRepository
	OrderRepository       repositories.OrderRepository
	AccrualRepository     repositories.AccrualRepository
	UserBalanceRepository repositories.UserBalanceRepository
}

type bonusSystemResponse struct {
	Order   string  `json:"order"`
	Status  string  `json:"status"`
	Accrual float64 `json:"accrual"`
}

func NewLoadOrderService(
	accrualSystemAddress string,
	logger *zap.Logger,
	userRepository repositories.UserRepository,
	orderRepository repositories.OrderRepository,
	accrualRepository repositories.AccrualRepository,
	userBalanceRepository repositories.UserBalanceRepository,
) *LoadOrderService {
	return &LoadOrderService{
		AccrualSystemAddress:  accrualSystemAddress,
		Logger:                logger,
		UserRepository:        userRepository,
		OrderRepository:       orderRepository,
		AccrualRepository:     accrualRepository,
		UserBalanceRepository: userBalanceRepository,
	}
}

func (s *LoadOrderService) Handle(ctx context.Context, orderNumber string, userID int64) error {
	wg := new(sync.WaitGroup)
	user, err := s.UserRepository.GetUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("user %d not found", userID)
	}

	order := models.NewOrder(user.ID, orderNumber)
	order, err = s.OrderRepository.Create(ctx, order)
	if err != nil {
		return fmt.Errorf("could not create order: %v", err)
	}
	accrual := models.NewAccrual(order.ID)
	accrual, err = s.AccrualRepository.Create(ctx, accrual)
	if err != nil {
		return fmt.Errorf("could not create accrual: %v", err)
	}

	wg.Add(1)
	var bonusSystemResponse bonusSystemResponse
	go s.getOrderDataByNumber(wg, order.Number, &bonusSystemResponse)
	wg.Wait()
	if &bonusSystemResponse == nil {
		return fmt.Errorf("could not get order info from bonus system")
	}

	accrual.UpdateStatus(bonusSystemResponse.Status, bonusSystemResponse.Accrual)
	accrual, err = s.AccrualRepository.Update(ctx, accrual)
	if err != nil {
		return fmt.Errorf("could not update accrual: %v", err)
	}
	if accrual.Value == 0 {
		return nil
	}

	userBalance, err := s.UserBalanceRepository.GetUserBalanceByID(ctx, user.ID, true)
	if err != nil {
		return fmt.Errorf("could not get user balance: %v", err)
	}
	userBalance.UpdateAccrual(accrual.Value)
	err = s.UserBalanceRepository.Update(ctx, userBalance)
	if err != nil {
		return fmt.Errorf("could not update user balance: %v", err)
	}

	return nil
}

func (s *LoadOrderService) getOrderDataByNumber(wg *sync.WaitGroup, orderNumber string, response *bonusSystemResponse) {
	path := fmt.Sprintf("http://%s/api/orders/%s", s.AccrualSystemAddress, orderNumber)
	client := resty.New()
	resp, err := client.R().Get(path)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("could not get order info from bonus system: %v", err))
		return
	}

	if err := json.Unmarshal(resp.Body(), response); err != nil {
		s.Logger.Error(fmt.Sprintf("could not unmarshal order info from bonus system: %v", err))
		return
	}

	wg.Done()
	return
}
