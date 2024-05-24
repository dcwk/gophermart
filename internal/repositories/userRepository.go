package repositories

import (
	"context"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	FindUserByLogin(ctx context.Context, login string) (*models.User, error)
}

type userRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	row := ur.DB.QueryRow(
		ctx,
		`INSERT INTO public.user (login, password, created_at, updated_at) VALUES ($1, $2, NOW(), NOW()) RETURNING ("id")`,
		user.Login,
		user.Password,
	)
	err := row.Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) FindUserByLogin(ctx context.Context, login string) (*models.User, error) {
	var user models.User
	row := ur.DB.QueryRow(
		ctx,
		`SELECT id, login, password FROM public.user WHERE login = $1`,
		login,
	)
	err := row.Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}