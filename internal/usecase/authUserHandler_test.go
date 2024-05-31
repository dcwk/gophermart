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

func TestAuthUserHandler_Handle(t *testing.T) {
	tests := []struct {
		Name     string
		Login    string
		Password string
		MockUser *models.User
		MockErr  error
		UserID   int64
		Want     int64
		Err      error
	}{
		{
			Name:     "Test can auth user",
			Login:    "test",
			Password: "test",
			MockUser: &models.User{
				ID:       1,
				Login:    "test",
				Password: "test",
			},
			MockErr: nil,
			UserID:  1,
			Want:    1,
		},
		{
			Name:     "Test fail user not found",
			Login:    "test",
			Password: "test",
			MockUser: nil,
			MockErr:  fmt.Errorf("no rows"),
			UserID:   0,
			Want:     1,
			Err:      fmt.Errorf("failed to find user by login: no rows"),
		},
		{
			Name:     "Test fail invalid password",
			Login:    "test",
			Password: "test1",
			MockUser: &models.User{
				ID:       1,
				Login:    "test",
				Password: "test",
			},
			MockErr: nil,
			UserID:  0,
			Want:    1,
			Err:     fmt.Errorf("invalid password"),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var err error
			assert.NoError(t, err)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			userRepository := mock_repositories.NewMockUserRepository(ctrl)
			if test.MockUser != nil {
				err = test.MockUser.HashPassword()
				assert.NoError(t, err)
			}
			userRepository.EXPECT().GetUserByLogin(gomock.Any(), gomock.Any()).Return(test.MockUser, test.MockErr)
			service := NewAuthHandler(userRepository)

			token, err := service.Handle(context.Background(), test.Login, test.Password)

			if test.Err != nil {
				assert.Error(t, test.Err, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.Want, auth.GetUserID(token))
			}
		})
	}
}
