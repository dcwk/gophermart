package usecase

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"

	"github.com/dcwk/gophermart/internal/models"
)

func TestCreateOrderHandler_Handle(t *testing.T) {
	testDate := time.Now()
	tests := []struct {
		Name         string
		PrepareMocks func(s *suite)
		UserID       int64
		OrderNumber  string
		Want         string
		Err          error
	}{
		{
			Name: "test can create order",
			PrepareMocks: func(s *suite) {
				s.UserRepository.EXPECT().
					GetUserByID(gomock.Any(), gomock.Any()).
					Return(&models.User{ID: 1, Login: "test", Password: "test"}, nil)
				s.OrderRepository.EXPECT().
					FindOrderByNumber(gomock.Any(), gomock.Any()).
					Return(nil, fmt.Errorf("test"))
				s.OrderRepository.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(
						&models.Order{
							ID:        1,
							UserID:    1,
							Number:    "110135840886155",
							CreatedAt: testDate,
						},
						nil,
					)
				s.AccrualRepository.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(
						&models.Accrual{
							OrderID:   1,
							Status:    models.Processing,
							CreatedAt: testDate,
						},
						nil,
					)
				s.AccrualRepository.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Return(
						&models.Accrual{
							OrderID:   1,
							Status:    models.Processed,
							Value:     100,
							CreatedAt: testDate,
						},
						nil,
					)
				s.UserBalanceRepository.EXPECT().
					GetUserBalanceByID(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(
						&models.UserBalance{
							UserID:     1,
							Accrual:    4830,
							Withdrawal: 0,
						},
						nil,
					)
				s.UserBalanceRepository.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			UserID:      1,
			OrderNumber: "110135840886155",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			s := defaultSuite(t)
			acrrualOrderChan := make(chan models.AccrualOrder)
			tt.PrepareMocks(s)
			logger := zaptest.NewLogger(t)
			handler := NewCreateOrderHandler(
				logger,
				s.UserRepository,
				s.OrderRepository,
				s.AccrualRepository,
				s.UserBalanceRepository,
			)

			go func() {
				acrrualOrderChan <- models.AccrualOrder{
					Order:   "110135840886155",
					Status:  models.Processed,
					Accrual: 100,
				}
			}()
			res, err := handler.Handle(
				context.Background(),
				acrrualOrderChan,
				tt.OrderNumber,
				tt.UserID,
			)

			if err != nil {
				assert.Error(t, tt.Err, err)
			} else {
				assert.Equal(t, tt.Want, res)
				assert.NoError(t, err)
			}
		})
	}
}
