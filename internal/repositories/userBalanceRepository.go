package repositories

import (
	"context"
	"fmt"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserBalanceRepository interface {
	Create(ctx context.Context, userBalance *models.UserBalance) (*models.UserBalance, error)
	Update(ctx context.Context, userBalance *models.UserBalance) error
	GetUserBalanceByID(ctx context.Context, userID int64, withLock bool) (*models.UserBalance, error)
}

type userBalanceRepository struct {
	DB *pgxpool.Pool
}

func NewUserBalanceRepository(db *pgxpool.Pool) UserBalanceRepository {
	return &userBalanceRepository{
		DB: db,
	}
}

func (r *userBalanceRepository) Create(ctx context.Context, userBalance *models.UserBalance) (*models.UserBalance, error) {
	_, err := r.DB.Query(
		ctx,
		`INSERT INTO user_balance (user_id, accrual, withdrawal) VALUES ($1, $2, $3)`,
		userBalance.UserId,
		userBalance.Accrual,
		userBalance.Withdrawal,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create user balance: %w", err)
	}

	return userBalance, nil
}

func (r *userBalanceRepository) Update(ctx context.Context, userBalance *models.UserBalance) error {
	_, err := r.DB.Query(
		ctx,
		`UPDATE user_balance SET accrual = $2, withdrawal = $3 WHERE user_id = $1`,
		userBalance.Accrual,
		userBalance.Withdrawal,
		userBalance.UserId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *userBalanceRepository) GetUserBalanceByID(ctx context.Context, userID int64, withLock bool) (*models.UserBalance, error) {
	var (
		userBalance models.UserBalance
		row         pgx.Row
	)

	if withLock {
		row = r.DB.QueryRow(
			ctx,
			`SELECT user_id, accrual, withdrawal FROM public.user_balance WHERE user_id = $1 FOR UPDATE`,
			userID,
		)
	} else {
		row = r.DB.QueryRow(
			ctx,
			`SELECT user_id, accrual, withdrawal FROM public.user_balance WHERE user_id = $1`,
			userID,
		)
	}

	err := row.Scan(&userBalance.UserId, &userBalance.Accrual, &userBalance.Withdrawal)
	if err != nil {
		return nil, err
	}

	return &userBalance, nil
}
