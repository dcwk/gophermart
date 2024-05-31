package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock_repositories "github.com/dcwk/gophermart/internal/repositories/mocks"
)

type suite struct {
	UserRepository        *mock_repositories.MockUserRepository
	UserBalanceRepository *mock_repositories.MockUserBalanceRepository
	OrderRepository       *mock_repositories.MockOrderRepository
	AccrualRepository     *mock_repositories.MockAccrualRepository
	WithdrawalRepository  *mock_repositories.MockWithdrawalRepository
}

func defaultSuite(t *testing.T) *suite {
	ctrl := gomock.NewController(t)
	s := suite{}
	s.UserRepository = mock_repositories.NewMockUserRepository(ctrl)
	s.UserBalanceRepository = mock_repositories.NewMockUserBalanceRepository(ctrl)
	s.OrderRepository = mock_repositories.NewMockOrderRepository(ctrl)
	s.AccrualRepository = mock_repositories.NewMockAccrualRepository(ctrl)
	s.WithdrawalRepository = mock_repositories.NewMockWithdrawalRepository(ctrl)

	return &s
}