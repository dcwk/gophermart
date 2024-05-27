package models

import "time"

const (
	NEW        = "NEW"
	PROCESSING = "PROCESSING"
	INVALID    = "INVALID"
	PROCESSED  = "PROCESSED"
)

type Accrual struct {
	OrderID   int64     `json:"order_id"`
	Status    string    `json:"status"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewAccrual(orderId int64) *Accrual {
	return &Accrual{
		OrderID: orderId,
		Status:  PROCESSING,
		Value:   0,
	}
}
