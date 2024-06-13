package usecase

import (
	"context"
	"fmt"

	"github.com/dcwk/gophermart/internal/repositories"
	"github.com/dcwk/gophermart/internal/utils/auth"
)

type AuthUserHandler struct {
	UserRepository repositories.UserRepository
}

func NewAuthHandler(userRepository repositories.UserRepository) *AuthUserHandler {
	return &AuthUserHandler{
		UserRepository: userRepository,
	}
}

func (s *AuthUserHandler) Handle(ctx context.Context, login string, password string) (string, error) {
	currentUser, err := s.UserRepository.GetUserByLogin(ctx, login)
	if err != nil || currentUser.ID == 0 {
		return "", fmt.Errorf("failed to find user by login: %w", err)
	}

	if !currentUser.VerifyPassword(password) {
		return "", fmt.Errorf("invalid password")
	}

	token, err := auth.BuildJWTString(currentUser.ID)
	if err != nil {
		return "", fmt.Errorf("failed to build JWT token: %w", err)
	}

	return token, nil
}
