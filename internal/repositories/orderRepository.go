package repositories

import (
	"context"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository interface {
	Create(ctx context.Context, order *models.Order) (*models.Order, error)
	FindOrderByNumber(ctx context.Context, orderNumber string) (*models.Order, error)
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

func (r *orderRepository) Create(ctx context.Context, order *models.Order) (*models.Order, error) {
	row := r.DB.QueryRow(
		ctx,
		`INSERT INTO "order" (user_id, number, created_at) VALUES ($1, $2, NOW()) RETURNING ("id")`,
		order.UserID,
		order.Number,
	)
	err := row.Scan(&order.ID)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *orderRepository) FindUserOrders(ctx context.Context, userID int64) ([]*models.Order, error) {
	rows, err := r.DB.Query(
		ctx,
		`SELECT o.id, o.user_id, o.number, a.status, a.value, o.created_at FROM "order" AS o 
				INNER JOIN "accrual" AS a ON o.id = a.order_id
			   WHERE o.user_id=$1`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*models.Order
	for rows.Next() {
		order := models.Order{}
		err := rows.Scan(
			&order.ID,
			&order.UserID,
			&order.Number,
			&order.Status,
			&order.Accrual,
			&order.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		orders = append(orders, &order)
	}

	return orders, nil
}

func (r *orderRepository) FindOrderByNumber(ctx context.Context, orderNumber string) (*models.Order, error) {
	var order models.Order
	row := r.DB.QueryRow(
		ctx,
		`SELECT id, user_id, number, created_at FROM "order" WHERE number=$1`,
		orderNumber,
	)
	err := row.Scan(&order.ID, &order.UserID, &order.Number, &order.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
