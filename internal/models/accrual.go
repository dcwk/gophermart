package models

import "time"

const (
	New        = "NEW"
	Processing = "PROCESSING"
	Invalid    = "INVALID"
	Processed  = "PROCESSED"
)

type Accrual struct {
	OrderID   int64     `json:"order_id"`
	Status    string    `json:"status"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewAccrual(orderID int64) *Accrual {
	return &Accrual{
		OrderID:   orderID,
		Status:    Processing,
		Value:     0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (a *Accrual) UpdateStatus(status string, value float64) {
	a.Status = status
	a.Value = value
}
