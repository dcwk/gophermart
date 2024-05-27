package repositories

import "github.com/jackc/pgx/v5/pgxpool"

type AccrualRepository interface{}

type accrualRepository struct {
	DB *pgxpool.Pool
}

func NewAccrualRepository(db *pgxpool.Pool) AccrualRepository {
	return &accrualRepository{
		DB: db,
	}
}
