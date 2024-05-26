package models

type UserBalance struct {
	ID         int     `json:"id"`
	Accrual    float64 `json:"accrual"`
	Withdrawal float64 `json:"withdrawal"`
}
