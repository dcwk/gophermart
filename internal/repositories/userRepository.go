package repositories

import (
	"context"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByID(ctx context.Context, ID int64) (*models.User, error)
	GetUserByLogin(ctx context.Context, login string) (*models.User, error)
}

type userRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	row := r.DB.QueryRow(
		ctx,
		`INSERT INTO "user" (login, password, created_at, updated_at) VALUES ($1, $2, NOW(), NOW()) RETURNING ("id")`,
		user.Login,
		user.Password,
	)
	err := row.Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, userID int64) (*models.User, error) {
	var user models.User
	row := r.DB.QueryRow(
		ctx,
		`SELECT id, login, password FROM "user" WHERE id = $1`,
		userID,
	)
	err := row.Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUserByLogin(ctx context.Context, login string) (*models.User, error) {
	var user models.User
	row := r.DB.QueryRow(
		ctx,
		`SELECT id, login, password FROM "user" WHERE login = $1`,
		login,
	)
	err := row.Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
