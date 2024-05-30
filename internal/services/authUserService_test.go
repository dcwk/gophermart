package services

import (
	"context"
	"testing"

	"github.com/dcwk/gophermart/internal/models"
	mock_repositories "github.com/dcwk/gophermart/internal/repositories/mocks"
	"github.com/dcwk/gophermart/internal/utils/auth"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAuthUserService_Handle(t *testing.T) {
	tests := []struct {
		Name     string
		Login    string
		Password string
		UserID   int64
		Want     int64
		Err      error
	}{
		{
			Name:     "Test can auth user",
			Login:    "test",
			Password: "test",
			UserID:   1,
			Want:     1,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var err error
			assert.NoError(t, err)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			userRepository := mock_repositories.NewMockUserRepository(ctrl)
			user := models.User{
				ID:       test.UserID,
				Login:    test.Login,
				Password: test.Password,
			}
			err = user.HashPassword()
			assert.NoError(t, err)
			userRepository.EXPECT().GetUserByLogin(gomock.Any(), gomock.Any()).Return(&user, nil)
			service := NewAuthService(userRepository)

			token, err := service.Handle(context.Background(), test.Login, test.Password)

			assert.Equal(t, test.Want, auth.GetUserID(token))
			if test.Err != nil {
				assert.Equal(t, test.Err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
