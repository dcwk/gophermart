package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/dcwk/gophermart/internal/models"
	mock_repositories "github.com/dcwk/gophermart/internal/repositories/mocks"
	"github.com/dcwk/gophermart/internal/utils/auth"
)

func TestRegisterUserService_Handle(t *testing.T) {
	tests := []struct {
		Name      string
		Login     string
		Password  string
		MockUser  *models.User
		MockErr   error
		MockUBErr error
		UserID    int64
		Want      int64
		Err       error
	}{
		{
			Name:     "Can register user",
			Login:    "test",
			Password: "test",
			MockUser: &models.User{
				ID:       1,
				Login:    "test",
				Password: "test",
			},
			MockErr:   nil,
			MockUBErr: nil,
			UserID:    1,
			Want:      1,
			Err:       nil,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userRepository := mock_repositories.NewMockUserRepository(ctrl)
			userRepository.EXPECT().
				GetUserByLogin(gomock.Any(), gomock.Any()).
				Return(nil, fmt.Errorf("err"))
			userRepository.EXPECT().
				Create(gomock.Any(), gomock.Any()).
				Return(test.MockUser, test.MockErr)
			userBalanceRepository := mock_repositories.NewMockUserBalanceRepository(ctrl)
			userBalanceRepository.EXPECT().
				Create(gomock.Any(), gomock.Any()).
				Return(nil, test.MockUBErr)
			service := NewRegisterUserHandler(userRepository, userBalanceRepository)

			token, err := service.Handle(context.Background(), test.Login, test.Password)

			if test.Err != nil {
				assert.Equal(t, test.Err, err)
			}

			assert.Equal(t, test.Want, auth.GetUserID(token))
			assert.NoError(t, err)
		})
	}
}
