package models

import "time"

type Withdrawal struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	OrderID   int64     `json:"order_id"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}

func NewWithdrawal(userID int64, orderID int64, value float64) *Withdrawal {
	return &Withdrawal{
		UserID:    userID,
		OrderID:   orderID,
		Value:     value,
		CreatedAt: time.Now(),
	}
}
