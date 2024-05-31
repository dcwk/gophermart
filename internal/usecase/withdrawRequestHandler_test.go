package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"

	"github.com/dcwk/gophermart/internal/models"
)

func TestWithdrawRequestHandler_Handle(t *testing.T) {
	testDate := time.Now()
	tests := []struct {
		Name         string
		PrepareMocks func(s *suite)
		UserID       int64
		OrderID      string
		Sum          float64
		Want         string
		Err          error
	}{
		{
			Name: "test can create withdraw request",
			PrepareMocks: func(s *suite) {
				s.UserRepository.EXPECT().
					GetUserByID(gomock.Any(), gomock.Any()).
					Return(&models.User{ID: 1, Login: "test", Password: "test"}, nil)
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
				s.WithdrawalRepository.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(
						&models.Withdrawal{
							ID:          1,
							UserID:      1,
							OrderID:     1,
							OrderNumber: "110135840886155",
							Value:       100,
							CreatedAt:   testDate,
						},
						nil,
					)
				s.UserBalanceRepository.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			UserID:  1,
			OrderID: "110135840886155",
			Sum:     100,
			Want:    "",
			Err:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			s := defaultSuite(t)
			tt.PrepareMocks(s)
			logger := zaptest.NewLogger(t)
			handler := NewWithdrawRequestHandler(
				logger,
				s.UserRepository,
				s.UserBalanceRepository,
				s.OrderRepository,
				s.WithdrawalRepository,
			)

			res, err := handler.Handle(context.Background(), tt.UserID, tt.OrderID, tt.Sum)

			if err != nil {
				assert.Equal(t, tt.Err, err)
			} else {
				assert.Equal(t, tt.Want, res)
				assert.NoError(t, tt.Err, err)
			}
		})
	}
}
