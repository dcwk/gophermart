package models

type UserBalance struct {
	UserId     int64   `json:"user_id"`
	Accrual    float64 `json:"accrual"`
	Withdrawal float64 `json:"withdrawal"`
}
