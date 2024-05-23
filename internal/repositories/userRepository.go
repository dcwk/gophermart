package repositories

import (
	"context"
	"database/sql"

	"github.com/dcwk/gophermart/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	FindUserByLogin(ctx context.Context, login string) (*models.User, error)
}

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	return &models.User{}, nil
}

// TODO: Нужно сделать поведение как GetUser и переименовать метод
func (ur *userRepository) FindUserByLogin(ctx context.Context, login string) (*models.User, error) {
	return &models.User{}, nil
}
