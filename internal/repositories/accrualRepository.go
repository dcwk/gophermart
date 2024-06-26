package repositories

import (
	"context"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AccrualRepository interface {
	Create(ctx context.Context, accrual *models.Accrual) (*models.Accrual, error)
	Update(ctx context.Context, accrual *models.Accrual) (*models.Accrual, error)
}

type accrualRepository struct {
	DB *pgxpool.Pool
}

func NewAccrualRepository(db *pgxpool.Pool) AccrualRepository {
	return &accrualRepository{
		DB: db,
	}
}

func (r *accrualRepository) Create(ctx context.Context, accrual *models.Accrual) (*models.Accrual, error) {
	_, err := r.DB.Query(
		ctx,
		`INSERT INTO "accrual" (order_id, status, value, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW())`,
		accrual.OrderID,
		accrual.Status,
		accrual.Value,
	)
	if err != nil {
		return nil, err
	}

	return accrual, nil
}

func (r *accrualRepository) Update(ctx context.Context, accrual *models.Accrual) (*models.Accrual, error) {
	_, err := r.DB.Query(
		ctx,
		`UPDATE accrual SET status=$1, value=$2, updated_at=NOW() WHERE order_id=$3`,
		accrual.Status,
		accrual.Value,
		accrual.OrderID,
	)
	if err != nil {
		return nil, err
	}

	return accrual, nil
}
