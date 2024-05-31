package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/dcwk/gophermart/internal/models"
)

func TestGetOrdersService_Handle(t *testing.T) {
	testDate := time.Now()
	tests := []struct {
		Name         string
		PrepareMocks func(s *suite)
		Want         []*models.Order
		Err          error
	}{
		{
			Name: "Test can get orders",
			PrepareMocks: func(s *suite) {
				s.UserRepository.EXPECT().
					GetUserByID(gomock.Any(), gomock.Any()).
					Return(&models.User{ID: 1, Login: "test", Password: "test"}, nil)
				s.OrderRepository.EXPECT().
					FindUserOrders(gomock.Any(), gomock.Any()).
					Return(
						[]*models.Order{
							{
								UserID:    1,
								Number:    "12344453",
								Status:    "PROCESSED",
								Accrual:   5451,
								CreatedAt: testDate,
							},
						},
						nil,
					)
			},
			Want: []*models.Order{
				{
					UserID:    1,
					Number:    "12344453",
					Status:    "PROCESSED",
					Accrual:   5451,
					CreatedAt: testDate,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			s := defaultSuite(t)
			test.PrepareMocks(s)

			service := NewGetOrdersHandler(
				s.UserRepository,
				s.OrderRepository,
			)

			res, err := service.Handle(context.Background(), 1)
			if err != nil {
				require.Error(t, test.Err, err)
			}

			require.Equal(t, test.Want, res)
			require.NoError(t, err)
		})
	}
}
