package repositories

import (
	"context"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository interface {
	FindUserOrders(ctx context.Context, userID int64) ([]*models.Order, error)
}

type orderRepository struct {
	DB *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) OrderRepository {
	return &orderRepository{
		DB: db,
	}
}

func (or *orderRepository) FindUserOrders(ctx context.Context, userID int64) ([]*models.Order, error) {
	rows, err := or.DB.Query(ctx, `SELECT id, user_id, number, created_at FROM "order" WHERE user_id=$1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*models.Order
	for rows.Next() {
		order := &models.Order{}
		err := rows.Scan(order.ID, &order.UserID, &order.Number, &order.CreatedAt)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}
