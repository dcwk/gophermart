package repositories

import (
	"context"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WithdrawalRepository interface {
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

func (r *withdrawalRepository) FindUserWithdrawals(ctx context.Context, userID int64) ([]*models.Withdrawal, error) {
	rows, err := r.DB.Query(ctx, `SELECT id, order_id, value, created_at FROM withdrawal WHERE user_id=$1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*models.Order
	for rows.Next() {
		order := models.Order{}
		err := rows.Scan(&order.ID, &order.UserID, &order.Number, &order.CreatedAt)
		if err != nil {
			return nil, err
		}

		orders = append(orders, &order)
	}

	return orders, nil
}
