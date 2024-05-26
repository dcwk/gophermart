package models

import "time"

type Withdrawal struct {
	ID        int64     `json:"id"`
	OrderID   int64     `json:"order_id"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}
