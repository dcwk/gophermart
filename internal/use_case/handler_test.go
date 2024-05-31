package use_case

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock_repositories "github.com/dcwk/gophermart/internal/repositories/mocks"
)

type suite struct {
	UserRepository  *mock_repositories.MockUserRepository
	OrderRepository *mock_repositories.MockOrderRepository
}

func defaultSuite(t *testing.T) *suite {
	ctrl := gomock.NewController(t)
	s := suite{}
	s.UserRepository = mock_repositories.NewMockUserRepository(ctrl)
	s.OrderRepository = mock_repositories.NewMockOrderRepository(ctrl)

	return &s
}
