package repositories

import (
	"context"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WithdrawalRepository interface {
	Create(ctx context.Context, withdrawal *models.Withdrawal) (*models.Withdrawal, error)
	FindUserWithdrawals(ctx context.Context, userID int64) ([]*models.Withdrawal, error)
}

type withdrawalRepository struct {
	DB *pgxpool.Pool
}

func NewWithdrawalRepository(db *pgxpool.Pool) WithdrawalRepository {
	return &withdrawalRepository{
		DB: db,
	}
}

func (r *withdrawalRepository) Create(ctx context.Context, withdrawal *models.Withdrawal) (*models.Withdrawal, error) {
	row := r.DB.QueryRow(
		ctx,
		`INSERT INTO withdrawal (user_id, order_id, value, created_at) VALUES ($1, $2, $3, NOW()) RETURNING ("id")`,
		withdrawal.UserID,
		withdrawal.OrderID,
		withdrawal.Value,
	)
	err := row.Scan(&withdrawal.ID)
	if err != nil {
		return nil, err
	}

	return withdrawal, nil
}

func (r *withdrawalRepository) FindUserWithdrawals(ctx context.Context, userID int64) ([]*models.Withdrawal, error) {
	rows, err := r.DB.Query(ctx, `SELECT id, user_id, order_id, value, created_at FROM withdrawal WHERE user_id=$1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var withdrawals []*models.Withdrawal
	for rows.Next() {
		withdrawal := models.Withdrawal{}
		err := rows.Scan(&withdrawal.ID, &withdrawal.UserID, &withdrawal.OrderID, &withdrawal.Value)
		if err != nil {
			return nil, err
		}

		withdrawals = append(withdrawals, &withdrawal)
	}

	return withdrawals, nil
}
