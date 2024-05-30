package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type LoadOrderService struct {
	AccrualSystemAddress string
	Logger               *zap.Logger
}

func NewLoadOrderService(accrualSystemAddress string, logger *zap.Logger) *LoadOrderService {
	return &LoadOrderService{
		AccrualSystemAddress: accrualSystemAddress,
		Logger:               logger,
	}
}

func (s *LoadOrderService) Handle(ctx context.Context, orderChannel chan models.AccrualOrder, orderNumber string) {
	path := fmt.Sprintf("http://%s/api/orders/%s", s.AccrualSystemAddress, orderNumber)
	var accrualOrder models.AccrualOrder

	for i := 0; i < 10; i++ {
		if i != 0 {
			time.Sleep(2 * time.Second)
		}

		client := resty.New().SetTimeout(10 * time.Second)
		resp, err := client.R().
			Get(path)
		if err != nil {
			s.Logger.Error(fmt.Sprintf("could not get order info from bonus system: %v", err))
			continue
		}

		if err := json.Unmarshal(resp.Body(), &accrualOrder); err != nil {
			s.Logger.Error(fmt.Sprintf("could not unmarshal order info from bonus system: %v", err))
			continue
		}

		s.Logger.Info(fmt.Sprintf("order info from bonus system: %v", accrualOrder))
		if accrualOrder.Status == models.Invalid || accrualOrder.Status == models.Processed {
			break
		}
	}

	orderChannel <- accrualOrder
}
