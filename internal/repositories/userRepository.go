package repositories

import (
	"context"
	"database/sql"

	"github.com/dcwk/gophermart/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	FindUserByID(ctx context.Context, id int64) (models.User, error)
}

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	return models.User{}, nil
}

func (ur *userRepository) FindUserByID(ctx context.Context, id int64) (models.User, error) {
	return models.User{}, nil
}
