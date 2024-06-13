package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/dcwk/gophermart/internal/models"
)

func TestGetWithdrawalsHandler_Handle(t *testing.T) {
	testDate := time.Now()
	tests := []struct {
		Name         string
		PrepareMocks func(s *suite)
		Want         []*models.Withdrawal
		Err          error
	}{
		{
			Name: "test can get withdrawals",
			PrepareMocks: func(s *suite) {
				s.UserRepository.EXPECT().
					GetUserByID(gomock.Any(), gomock.Any()).
					Return(&models.User{ID: 1, Login: "test", Password: "test"}, nil)
				s.WithdrawalRepository.EXPECT().
					FindUserWithdrawals(gomock.Any(), gomock.Any()).
					Return(
						[]*models.Withdrawal{
							{
								ID:          1,
								UserID:      1,
								OrderID:     1,
								OrderNumber: "110135840886155",
								Value:       100,
								CreatedAt:   testDate,
							},
						},
						nil,
					)
			},
			Want: []*models.Withdrawal{
				{
					ID:          1,
					UserID:      1,
					OrderID:     1,
					OrderNumber: "110135840886155",
					Value:       100,
					CreatedAt:   testDate,
				},
			},
			Err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			s := defaultSuite(t)
			test.PrepareMocks(s)
			userID := int64(1)
			handler := NewGetWithdrawalsHandler(s.UserRepository, s.WithdrawalRepository)

			res, err := handler.Handle(context.Background(), userID)

			if err != nil {
				assert.Error(t, test.Err, err)
			} else {
				assert.Equal(t, test.Want, res)
				assert.NoError(t, test.Err)
			}
		})
	}
}
