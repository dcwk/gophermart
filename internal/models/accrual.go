package models

import "time"

type Accrual struct {
	OrderID   int64     `json:"order_id"`
	Status    string    `json:"status"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
