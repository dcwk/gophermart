package services

import (
	"context"
	"fmt"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/dcwk/gophermart/internal/repositories"
	"github.com/dcwk/gophermart/internal/utils/auth"
)

type AuthUserService struct {
	UserRepository repositories.UserRepository
}

func NewAuthService(userRepository repositories.UserRepository) *AuthUserService {
	return &AuthUserService{
		UserRepository: userRepository,
	}
}

func (aus *AuthUserService) Authenticate(ctx context.Context, user *models.User) (string, error) {
	currentUser, err := aus.UserRepository.FindUserByLogin(ctx, user.Login)
	if err != nil || currentUser.ID == 0 {
		return "", fmt.Errorf("failed to find user by login: %w", err)
	}

	token, err := auth.BuildJWTString(currentUser.ID)
	if err != nil {
		return "", fmt.Errorf("failed to build JWT token: %w", err)
	}

	return token, nil
}
