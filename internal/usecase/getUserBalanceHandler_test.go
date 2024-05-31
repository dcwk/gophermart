package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/dcwk/gophermart/internal/models"
)

func TestGetUserBalanceHandler_Handle(t *testing.T) {
	tests := []struct {
		Name         string
		PrepareMocks func(s *suite)
		Want         *models.UserBalance
		Err          error
	}{
		{
			Name: "Test can get user balance",
			PrepareMocks: func(s *suite) {
				s.UserRepository.EXPECT().
					GetUserByID(gomock.Any(), gomock.Any()).
					Return(
						&models.User{
							ID:       1,
							Login:    "test",
							Password: "test",
						},
						nil,
					)
				s.UserBalanceRepository.EXPECT().
					GetUserBalanceByID(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(
						&models.UserBalance{
							UserID:     1,
							Accrual:    4830,
							Withdrawal: 100,
						},
						nil,
					)

			},
			Want: &models.UserBalance{
				UserID:     1,
				Accrual:    4830,
				Withdrawal: 100,
			},
			Err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			s := defaultSuite(t)
			test.PrepareMocks(s)
			userID := 1
			handler := NewGetUserBalanceHandler(s.UserRepository, s.UserBalanceRepository)

			res, err := handler.Handle(context.Background(), int64(userID))

			if err != nil {
				assert.Error(t, test.Err, err)
			} else {
				assert.Equal(t, test.Want, res)
				assert.NoError(t, err)
			}
		})
	}
}
